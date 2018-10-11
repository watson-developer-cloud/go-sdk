package core

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"strings"
	"time"

	cfenv "github.com/cloudfoundry-community/go-cfenv"
)

// common constants for core
const (
	APIKey        = "apikey"
	ICPPrefix     = "icp-"
	UserAgent     = "User-Agent"
	Authorization = "Authorization"
	Bearer        = "Bearer"
)

// ServiceOptions Service options
type ServiceOptions struct {
	Version        string
	URL            string
	Username       string
	Password       string
	IAMApiKey      string
	IAMAccessToken string
	IAMURL         string
}

// WatsonService Base Service
type WatsonService struct {
	Options        *ServiceOptions
	DefaultHeaders http.Header
	tokenManager   *TokenManager
	client         *http.Client
	userAgent      string
}

// NewWatsonService Instantiate a Watson Service
func NewWatsonService(options *ServiceOptions, serviceName string) (*WatsonService, error) {
	service := WatsonService{
		Options: options,

		client: &http.Client{
			Timeout: time.Second * 30,
		},
	}

	const sdkVersion = "0.0.1" // TODO: would there be a bumpversion?
	var userAgent = "watson-apis-go-sdk-" + sdkVersion
	userAgent += "-" + runtime.GOOS
	service.userAgent = userAgent

	if options.Username != "" && options.Password != "" {
		if options.Username == APIKey && !strings.HasPrefix(options.Password, ICPPrefix) {
			service.SetTokenManager(options.IAMApiKey, options.IAMAccessToken, options.IAMURL)
		} else {
			service.SetUsernameAndPassword(options.Username, options.Password)
		}
	} else if options.IAMAccessToken != "" || options.IAMApiKey != "" {
		service.SetTokenManager(options.IAMApiKey, options.IAMAccessToken, options.IAMURL)
	} else {
		// Try accessing VCAP_SERVICES env variable
		err := service.accessVCAP(serviceName)
		if err != nil {
			return nil, err
		}
	}

	return &service, nil
}

// SetUsernameAndPassword Sets the Username and Password
func (service *WatsonService) SetUsernameAndPassword(username string, password string) {
	service.Options.Username = username
	service.Options.Password = password
}

// SetTokenManager Sets the Token Manager for IAM Authentication
func (service *WatsonService) SetTokenManager(iamAPIKey string, iamAccessToken string, iamURL string) {
	service.Options.IAMApiKey = iamAPIKey
	service.Options.IAMAccessToken = iamAccessToken
	service.Options.IAMURL = iamURL
	tokenManager := NewTokenManager(iamAPIKey, iamURL, iamAccessToken)
	service.tokenManager = tokenManager
}

// SetIAMAccessToken Sets the IAM access token
func (service *WatsonService) SetIAMAccessToken(iamAccessToken string) {
	if service.tokenManager != nil {
		service.tokenManager.SetAccessToken(iamAccessToken)
	} else {
		tokenManager := NewTokenManager("", "", iamAccessToken)
		service.tokenManager = tokenManager
	}
	service.Options.IAMAccessToken = iamAccessToken
}

// SetIAMAPIKey Sets the IAM API key
func (service *WatsonService) SetIAMAPIKey(iamAPIKey string) {
	if service.tokenManager != nil {
		service.tokenManager.SetIAMAPIKey(iamAPIKey)
	} else {
		tokenManager := NewTokenManager(iamAPIKey, "", "")
		service.tokenManager = tokenManager
	}
	service.Options.IAMApiKey = iamAPIKey
}

// SetURL sets the service URL
func (service *WatsonService) SetURL(url string) {
	service.Options.URL = url
}

// SetDefaultHeaders sets HTTP headers to be sent in every request.
func (service *WatsonService) SetDefaultHeaders(headers http.Header) {
	service.DefaultHeaders = headers
}

// SetHTTPClient updates the client handling the requests
func (service *WatsonService) SetHTTPClient(client *http.Client) {
	service.client = client
}

// DisableSSLVerification skips SSL verification
func (service *WatsonService) DisableSSLVerification() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	service.client.Transport = tr
}

// Request performs the HTTP request
func (service *WatsonService) Request(req *http.Request, result interface{}) (*DetailedResponse, error) {
	// Add default headers
	if service.DefaultHeaders != nil {
		for k, v := range service.DefaultHeaders {
			req.Header.Add(k, strings.Join(v, ""))
		}
	}

	// Add authentication
	if service.tokenManager != nil {
		token := service.tokenManager.GetToken()
		req.Header.Add(Authorization, fmt.Sprintf(`%s %s`, Bearer, token))
	} else if service.Options.Username != "" && service.Options.Password != "" {
		req.SetBasicAuth(service.Options.Username, service.Options.Password)
	}

	// Perform the request
	resp, err := service.client.Do(req)
	if err != nil {
		return nil, err
	}

	// handle the response
	response := new(DetailedResponse)
	response.Headers = resp.Header
	response.StatusCode = resp.StatusCode
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		if resp != nil {
			buff := new(bytes.Buffer)
			buff.ReadFrom(resp.Body)
			return response, fmt.Errorf(buff.String())
		}
	}

	contentType := resp.Header.Get(ContentType)
	if contentType != "" {
		if IsJSONMimeType(contentType) && result != nil {
			json.NewDecoder(resp.Body).Decode(&result)
			response.Result = result
		}
	}

	if response.Result == nil {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return response, err
		}
		response.Result = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	}

	defer resp.Body.Close()
	return response, nil
}

func (service *WatsonService) accessVCAP(serviceName string) error {
	appEnv, envErr := cfenv.Current()

	if envErr != nil {
		return envErr
	}

	watsonService, servErr := appEnv.Services.WithName(serviceName)

	if servErr != nil {
		return servErr
	}

	creds := watsonService.Credentials

	username, userOK := creds["username"]
	password, passOK := creds["password"]

	if userOK && passOK {
		service.Options.Username = fmt.Sprint(username)
		service.Options.Password = fmt.Sprint(password)
		return nil
	}

	return fmt.Errorf("you must specify an IAM API key or username and password service credentials")
}
