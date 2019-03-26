package core

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path"
	"runtime"
	"strings"
	"time"

	cfenv "github.com/cloudfoundry-community/go-cfenv"
)

// common constants for core
const (
	APIKey                       = "apikey"
	ICPPrefix                    = "icp-"
	UserAgent                    = "User-Agent"
	Authorization                = "Authorization"
	Bearer                       = "Bearer"
	IBM_CREDENTIAL_FILE_ENV      = "IBM_CREDENTIALS_FILE"
	DEFAULT_CREDENTIAL_FILE_NAME = "ibm-credentials.env"
	URL                          = "url"
	USERNAME                     = "username"
	PASSWORD                     = "password"
	IAM_API_KEY                  = "iam_apikey"
	IAM_URL                      = "iam_url"
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
	TokenManager   *TokenManager
	Client         *http.Client
	UserAgent      string
}

// NewWatsonService Instantiate a Watson Service
func NewWatsonService(options *ServiceOptions, serviceName, displayName string) (*WatsonService, error) {
	if HasBadFirstOrLastChar(options.URL) {
		return nil, fmt.Errorf("The URL shouldn't start or end with curly brackets or quotes. Be sure to remove any {} and \" characters surrounding your URL")
	}

	service := WatsonService{
		Options: options,

		Client: &http.Client{
			Timeout: time.Second * 30,
		},
	}

	var userAgent = "watson-apis-go-sdk-" + Version
	userAgent += " " + runtime.GOOS + " " + runtime.Version()
	service.UserAgent = userAgent

	// 1. Credentials are passed in constructor
	if options.Username != "" && options.Password != "" {
		if options.Username == APIKey && !strings.HasPrefix(options.Password, ICPPrefix) {
			if err := service.SetTokenManager(options.IAMApiKey, options.IAMAccessToken, options.IAMURL); err != nil {
				return nil, err
			}
		} else {
			if err := service.SetUsernameAndPassword(options.Username, options.Password); err != nil {
				return nil, err
			}
		}
	} else if options.IAMAccessToken != "" || options.IAMApiKey != "" {
		if options.IAMApiKey != "" && strings.HasPrefix(options.IAMApiKey, ICPPrefix) {
			if err := service.SetUsernameAndPassword(APIKey, options.IAMApiKey); err != nil {
				return nil, err
			}
		} else {
			if err := service.SetTokenManager(options.IAMApiKey, options.IAMAccessToken, options.IAMURL); err != nil {
				return nil, err
			}
		}
	}

	// 2. Credentials from credential file
	if displayName != "" && service.Options.Username == "" && service.TokenManager == nil {
		serviceName := strings.ToLower(strings.Replace(displayName, " ", "_", -1))
		service.loadFromCredentialFile(serviceName, "=")
	}

	// 3. Try accessing VCAP_SERVICES env variable
	if service.Options.Username == "" && service.TokenManager == nil {
		err := service.accessVCAP(serviceName)
		if err != nil {
			return nil, err
		}
	}

	return &service, nil
}

// SetUsernameAndPassword Sets the Username and Password
func (service *WatsonService) SetUsernameAndPassword(username string, password string) error {
	if HasBadFirstOrLastChar(username) {
		return fmt.Errorf("The username shouldn't start or end with curly brackets or quotes. Be sure to remove any {} and \" characters surrounding your username")
	}
	if HasBadFirstOrLastChar(password) {
		return fmt.Errorf("The password shouldn't start or end with curly brackets or quotes. Be sure to remove any {} and \" characters surrounding your password")
	}
	service.Options.Username = username
	service.Options.Password = password
	return nil
}

// SetTokenManager Sets the Token Manager for IAM Authentication
func (service *WatsonService) SetTokenManager(iamAPIKey string, iamAccessToken string, iamURL string) error {
	if HasBadFirstOrLastChar(iamAPIKey) {
		return fmt.Errorf("The credentials shouldn't start or end with curly brackets or quotes. Be sure to remove any {} and \" characters surrounding your credentials")
	}
	service.Options.IAMApiKey = iamAPIKey
	service.Options.IAMAccessToken = iamAccessToken
	service.Options.IAMURL = iamURL
	tokenManager := NewTokenManager(iamAPIKey, iamURL, iamAccessToken)
	service.TokenManager = tokenManager
	return nil
}

// SetIAMAccessToken Sets the IAM access token
func (service *WatsonService) SetIAMAccessToken(iamAccessToken string) {
	if service.TokenManager != nil {
		service.TokenManager.SetAccessToken(iamAccessToken)
	} else {
		tokenManager := NewTokenManager("", "", iamAccessToken)
		service.TokenManager = tokenManager
	}
	service.Options.IAMAccessToken = iamAccessToken
}

// SetIAMAPIKey Sets the IAM API key
func (service *WatsonService) SetIAMAPIKey(iamAPIKey string) error {
	if HasBadFirstOrLastChar(iamAPIKey) {
		return fmt.Errorf("The credentials shouldn't start or end with curly brackets or quotes. Be sure to remove any {} and \" characters surrounding your credentials")
	}
	if service.TokenManager != nil {
		service.TokenManager.SetIAMAPIKey(iamAPIKey)
	} else {
		tokenManager := NewTokenManager(iamAPIKey, "", "")
		service.TokenManager = tokenManager
	}
	service.Options.IAMApiKey = iamAPIKey
	return nil
}

// SetURL sets the service URL
func (service *WatsonService) SetURL(url string) error {
	if HasBadFirstOrLastChar(url) {
		return fmt.Errorf("The URL shouldn't start or end with curly brackets or quotes. Be sure to remove any {} and \" characters surrounding your URL")
	}
	service.Options.URL = url
	return nil
}

// SetDefaultHeaders sets HTTP headers to be sent in every request.
func (service *WatsonService) SetDefaultHeaders(headers http.Header) {
	service.DefaultHeaders = headers
}

// SetHTTPClient updates the client handling the requests
func (service *WatsonService) SetHTTPClient(client *http.Client) {
	service.Client = client
}

// DisableSSLVerification skips SSL verification
func (service *WatsonService) DisableSSLVerification() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	service.Client.Transport = tr
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
	if service.TokenManager != nil {
		token := service.TokenManager.GetToken()
		req.Header.Add(Authorization, fmt.Sprintf(`%s %s`, Bearer, token))
	} else if service.Options.Username != "" && service.Options.Password != "" {
		req.SetBasicAuth(service.Options.Username, service.Options.Password)
	}

	// Perform the request
	resp, err := service.Client.Do(req)
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
			defer resp.Body.Close()
		}
	}

	if response.Result == nil && result != nil {
		response.Result = resp.Body
	}

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

func (service *WatsonService) loadFromCredentialFile(serviceName string, separator string) error {
	// File path specified by env variable
	credentialFilePath := os.Getenv(IBM_CREDENTIAL_FILE_ENV)

	// Home directory
	if credentialFilePath == "" {
		var filePath = path.Join(UserHomeDir(), DEFAULT_CREDENTIAL_FILE_NAME)
		if _, err := os.Stat(filePath); err == nil {
			credentialFilePath = filePath
		}
	}

	// Top-level of project directory
	if credentialFilePath == "" {
		dir, _ := os.Getwd()
		var filePath = path.Join(dir, "..", DEFAULT_CREDENTIAL_FILE_NAME)
		if _, err := os.Stat(filePath); err == nil {
			credentialFilePath = filePath
		}
	}

	if credentialFilePath != "" {
		file, err := os.Open(credentialFilePath)
		if err != nil {
			return err
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			var line = scanner.Text()
			var keyVal = strings.Split(strings.ToLower(line), separator)
			if len(keyVal) == 2 {
				service.setCredentialBasedOnType(serviceName, keyVal[0], keyVal[1])
			}
		}
	}
	return nil
}

func (service *WatsonService) setCredentialBasedOnType(serviceName, key, value string) {
	if strings.Contains(key, serviceName) {
		if strings.Contains(key, APIKey) {
			service.SetIAMAPIKey(value)
		} else if strings.Contains(key, URL) {
			service.SetURL(value)
		} else if strings.Contains(key, USERNAME) {
			service.Options.Username = value
		} else if strings.Contains(key, PASSWORD) {
			service.Options.Password = value
		} else if strings.Contains(key, IAM_API_KEY) {
			service.SetIAMAPIKey(value)
		}
	}
}
