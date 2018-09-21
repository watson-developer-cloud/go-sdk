package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
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
func (watsonResponse *WatsonResponse) GetHeaders() http.Header {
	return watsonResponse.Headers
}

// GetStatusCode returns the HTTP status code
func (watsonResponse *WatsonResponse) GetStatusCode() int {
	return watsonResponse.StatusCode
}

// GetResult returns the result from the service
func (watsonResponse *WatsonResponse) GetResult() interface{} {
	return watsonResponse.Result
}

// PrettyPrint pretty prints data
func (watsonResponse *WatsonResponse) PrettyPrint(data interface{}) {
	output, err := json.MarshalIndent(data, "", "    ")

	if err == nil {
		fmt.Printf("%+v\n", string(output))
	} else {
		fmt.Println("Sorry, couldn't pretty print the data")
	}
}

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
	body         io.Reader
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

// SetBodyContentJSON - set the body content from a JSON structure
func (service *WatsonService) SetBodyContentJSON(bodyContent interface{}) error {
	service.body = new(bytes.Buffer)
	err := json.NewEncoder(service.body.(io.Writer)).Encode(bodyContent)
	return err
}

// SetBodyContentStream - set the body content from an io.Reader instance
func (service *WatsonService) SetBodyContentStream(bodyContent io.Reader) error {
	service.body = bodyContent
	return nil
}

// SetBodyContentString - set the body content from a string
func (service *WatsonService) SetBodyContentString(bodyContent string) error {
	service.body = strings.NewReader(bodyContent)
	return nil
}

// SetBodyContent - sets the body content from one of three different sources, based on the content type
func (service *WatsonService) SetBodyContent(contentType string, jsonContent interface{}, jsonPatchContent interface{},
	nonJSONContent interface{}) error {
	if contentType != "" {
		if IsJSONMimeType(contentType) {
			err := service.SetBodyContentJSON(jsonContent)
			if err != nil {
				return err
			}
		} else if IsJSONPatchMimeType(contentType) {
			err := service.SetBodyContentJSON(jsonPatchContent)
			if err != nil {
				return err
			}
		} else {
			// Set the non-JSON body content based on the type of value passed in,
			// which should be either a "string" or an "io.Reader"
			if IsObjectAString(nonJSONContent) {
				service.SetBodyContentString(nonJSONContent.(string))
			} else if IsObjectAReader(nonJSONContent) {
				service.SetBodyContentStream(nonJSONContent.(io.Reader))
			} else {
				return fmt.Errorf("Invalid type for non-JSON body content: %s", reflect.TypeOf(nonJSONContent).String())
			}
		}
	}

	return nil
}

// HandleRequest perform the HTTP request
func (service *WatsonService) HandleRequest(method string, path string, headers map[string]string,
	params map[string]string, result interface{}) (*WatsonResponse, error) {

	fullURL := service.Options.URL + path

	// Create the request
	req, err := http.NewRequest(method, fullURL, service.body)
	if err != nil {
		return nil, err
	}
	service.body = nil

	// Define headers
	headers[userAgent] = service.userAgent
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
	if err != nil {
		panic(err)
	}

	// handle the response
	response := new(WatsonResponse)
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
