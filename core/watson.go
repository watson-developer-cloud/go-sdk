package core // TODO: decide on the full path

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
	"strings"
	"time"

	cfenv "github.com/cloudfoundry-community/go-cfenv"
)

const (
	apiKey          = "apikey"
	icpPrefix       = "icp-"
	userAgent       = "User-Agent"
	accept          = "Accept"
	applicationJSON = "application/json"
	authorization   = "Authorization"
	bearer          = "Bearer"
)

// WatsonResponse : Generic response for Watson API
type WatsonResponse struct {
	StatusCode int
	Headers    http.Header
	Result     interface{}
}

// GetHeaders returns the headers
func (watsonresponse *WatsonResponse) GetHeaders() http.Header {
	return watsonresponse.Headers
}

// GetStatusCode returns the HTTP status code
func (watsonresponse *WatsonResponse) GetStatusCode() int {
	return watsonresponse.StatusCode
}

// GetResult returns the result from the service
func (watsonresponse *WatsonResponse) GetResult() interface{} {
	return watsonresponse.Result
}

// PrettyPrint pretty prints data
func (watsonresponse *WatsonResponse) PrettyPrint(data interface{}) {
	output, err := json.MarshalIndent(data, "", "    ")

	if err == nil {
		fmt.Printf("%+v\n", string(output))
	} else {
		fmt.Println("Sorry, couldn't pretty print the data")
	}
}

// ServiceOptions Service options
type ServiceOptions struct {
	URL            string
	Username       string
	Password       string
	APIKey         string
	IAMApiKey      string
	IAMAccessToken string
	IAMURL         string
}

// WatsonService Base Service
type WatsonService struct {
	Version      string
	Options      *ServiceOptions
	Headers      http.Header
	tokenManager *TokenManager
	client       *http.Client
	userAgent    string
}

// NewWatsonService Instantiate a Watson Service
func NewWatsonService(options *ServiceOptions, serviceName string, version string) (*WatsonService, error) {
	service := WatsonService{
		Version: version,
		Options: options,

		client: &http.Client{
			Timeout: time.Second * 30,
		},
	}

	if service.Version == "" {
		return nil, fmt.Errorf("you must specify a version")
	}

	const sdkVersion = "0.0.1" // TODO: would there be a bumpversion?
	var userAgent = "watson-apis-go-sdk-" + sdkVersion
	userAgent += "-" + runtime.GOOS
	service.userAgent = userAgent

	if options.APIKey != "" {
		service.SetAPIKey(options.APIKey)
	} else if options.Username != "" && options.Password != "" {
		if options.Username == apiKey && strings.HasPrefix(options.Password, icpPrefix) {
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

// HandleRequest perform the HTTP request
func (service *WatsonService) HandleRequest(method string, path string, acceptJSON bool,
	body map[string]interface{}, headers map[string]string, params map[string]string, result interface{}) (*WatsonResponse, error) {

	fullURL := service.Options.URL + path

	// Encode body
	buff := new(bytes.Buffer)
	if body != nil {
		json.NewEncoder(buff).Encode(body)
	}

	// Create the request
	req, err := http.NewRequest(method, fullURL, buff)
	if err != nil {
		return nil, err
	}

	// Define headers
	headers[userAgent] = service.userAgent
	if acceptJSON {
		headers[accept] = applicationJSON
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	// Define query params
	query := req.URL.Query()
	for k, v := range params {
		query.Add(k, v)
	}

	// Add authentication
	if service.tokenManager != nil {
		token, _ := service.tokenManager.GetToken()
		req.Header.Add(authorization, bearer+" "+token)
	} else if service.Options.Username != "" && service.Options.Password != "" {
		req.SetBasicAuth(service.Options.Username, service.Options.Password)
	} else if service.Options.APIKey != "" {
		query.Add(apiKey, service.Options.APIKey)
	}

	// Encode query
	req.URL.RawQuery = query.Encode()

	// TODO: handle files

	// Perform the request
	resp, err := service.client.Do(req)
	response := new(WatsonResponse)
	response.Headers = resp.Header
	response.StatusCode = resp.StatusCode
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		if resp != nil {
			buff := new(bytes.Buffer)
			buff.ReadFrom(resp.Body)
			return response, fmt.Errorf(buff.String())
		}
		if err != nil {
			return response, err
		}
	}

	// handle the response
	defer resp.Body.Close()
	if acceptJSON {
		json.NewDecoder(resp.Body).Decode(&result)
	}
	response.Result = result
	return response, nil
}

// SetAPIKey Sets the API key used in Visual Recognition
func (service *WatsonService) SetAPIKey(apiKey string) {
	service.Options.APIKey = apiKey

	// temporary for visual recognition
	if service.Options.URL != "DEFAULT_FOR_VIS_REC" {
		service.Options.URL = "https://gateway-a.watsonplatform.net/visual-recognition/api"
	}
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
	if service.tokenManager == nil {
		service.tokenManager = NewTokenManager()
	}
	// TODO: should even be able to set iam url
	service.tokenManager.SetKey(iamAPIKey)
	service.tokenManager.SetToken(iamAccessToken)
}

// TODO: Expose other IAM token infos

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
	APIkey, keyOK := creds["apikey"]

	if keyOK {
		service.tokenManager.SetKey(fmt.Sprint(APIkey))
		return nil
	}

	if userOK && passOK {
		service.Options.Username = fmt.Sprint(username)
		service.Options.Password = fmt.Sprint(password)
		return nil
	}

	return fmt.Errorf("you must specify an IAM API key or username and password service credentials")
}
