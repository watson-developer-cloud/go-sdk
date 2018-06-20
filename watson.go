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
	creds Credentials
	useTM bool
	tokenManager *TokenManager
}

func NewClient(creds Credentials, serviceName string) (*Client, error) {
	client := Client{
		creds: creds,
		useTM: false,
		tokenManager: NewTokenManager(),
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
		client.useTM = true
		client.tokenManager.SetKey(creds.APIkey)
	} else if creds.IAMtoken != "" {
		// User manages IAM token
		client.useTM = true
		client.tokenManager.SetToken(creds.IAMtoken)
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
		client.useTM = true
		client.tokenManager.SetKey(fmt.Sprint(APIkey))
		return nil
	}

	if userOK && passOK {
		client.creds.Username = fmt.Sprint(username)
		client.creds.Password = fmt.Sprint(password)
		return nil
	}

	return fmt.Errorf("you must specify an IAM API key or username and password service credentials")
}

func (client *Client) Creds() *Credentials {
	return &client.creds
}

func (client *Client) UseTM() bool {
	return client.useTM
}

func (client *Client) TokenManager() *TokenManager {
	return client.tokenManager
}
