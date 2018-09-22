package core

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

// ServiceOptions Service options
type ServiceOptions struct {
	Version        string
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
	Options      *ServiceOptions
	Headers      http.Header
	tokenManager *TokenManager
	client       *http.Client
	userAgent    string
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

// Request performs the HTTP request
func (service *WatsonService) Request(req *http.Request, result interface{}) (*DetailedResponse, error) {
	// TODO: Update the request with default headers and other service-specific parameters.

	// Add authentication
	if service.tokenManager != nil {
		token, _ := service.tokenManager.GetToken()
		req.Header.Add(authorization, bearer+" "+token)
	} else if service.Options.Username != "" && service.Options.Password != "" {
		req.SetBasicAuth(service.Options.Username, service.Options.Password)
	}

	// Perform the request
	resp, err := service.client.Do(req)
	if err != nil {
		panic(err)
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

	// TODO: we should NOT assume the response is JSON just because the operation contains application/json
	// in its "produces" list.
	// Instead, we should interpret the response body according to the Content-Type header returned
	// in the response.
	json.NewDecoder(resp.Body).Decode(&result)

	defer resp.Body.Close()
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
