package golang_sdk

import (
	"fmt"
	"github.com/cloudfoundry-community/go-cfenv"
)

type Credentials struct {
	ServiceURL string
	Version string
	Username string
	Password string
	APIkey string
	IAMtoken string
}

type Client struct {
	Creds Credentials
	UseTM bool
	TokenManager *TokenManager
}

func NewClient(creds Credentials, serviceName string) (*Client, error) {
	client := Client{
		Creds: creds,
		UseTM: false,
		TokenManager: NewTokenManager(),
	}

	if creds.ServiceURL == "" {
		return nil, fmt.Errorf("you must specify a service API endpoint")
	}

	if creds.Version == "" {
		return nil, fmt.Errorf("you must specify a version")
	}

	// Username/password is default

	if creds.APIkey != "" {
		// SDK manages IAM token
		client.UseTM = true
		client.TokenManager.SetKey(creds.APIkey)
	} else if creds.IAMtoken != "" {
		// User manages IAM token
		client.UseTM = true
		client.TokenManager.SetToken(creds.IAMtoken)
	} else if creds.Username == "" || creds.Password == "" {
		// Try accessing VCAP_SERVICES env variable
		err := accessVCAP(&client, serviceName)

		if err != nil {
			return nil, err
		}
	}

	return &client, nil
}

func accessVCAP(client *Client, serviceName string) error {
	appEnv, envErr := cfenv.Current()

	if envErr != nil {
		return envErr
	}

	service, servErr := appEnv.Services.WithName(serviceName)

	if servErr != nil {
		return servErr
	}

	creds := service.Credentials

	username, userOK := creds["username"]
	password, passOK := creds["password"]
	APIkey, keyOK := creds["apikey"]

	if keyOK {
		client.UseTM = true
		client.TokenManager.SetKey(fmt.Sprint(APIkey))
		return nil
	}

	if userOK && passOK {
		client.Creds.Username = fmt.Sprint(username)
		client.Creds.Password = fmt.Sprint(password)
		return nil
	}

	return fmt.Errorf("you must specify an IAM API key or username and password service credentials")
}
