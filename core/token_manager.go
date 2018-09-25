package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// for handling token management
const (
	DefaultIAMURL            = "https://iam.bluemix.net/identity/token"
	DefaultContentType       = "application/x-www-form-urlencoded"
	DefaultAuthorization     = "Basic Yng6Yng="
	RequestTokenGrantType    = "urn:ibm:params:oauth:grant-type:apikey"
	RequestTokenResponseType = "cloud_iam"
	RefreshTokenGrantType    = "refresh_token"
)

// TokenInfo : Response struct from token request
type TokenInfo struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	Expiration   int64  `json:"expiration"`
}

// TokenManager : IAM token information
type TokenManager struct {
	userAccessToken string
	iamAPIkey       string
	iamURL          string
	tokenInfo       *TokenInfo
	client          *http.Client
}

// NewTokenManager : Instantiate TokenManager
func NewTokenManager(iamAPIkey string, iamURL string, userAccessToken string) *TokenManager {
	if iamURL == "" {
		iamURL = DefaultIAMURL
	}

	tokenManager := TokenManager{
		iamAPIkey:       iamAPIkey,
		iamURL:          iamURL,
		userAccessToken: userAccessToken,
		tokenInfo:       &TokenInfo{},

		client: &http.Client{
			// Timeout: time.Second * 30,
		},
	}
	return &tokenManager
}

// makes an HTTP request
func (tm *TokenManager) request(req *http.Request) *TokenInfo {
	resp, err := tm.client.Do(req)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		if resp != nil {
			buff := new(bytes.Buffer)
			buff.ReadFrom(resp.Body)
			panic(fmt.Errorf(buff.String()))
		}
	}

	tokenInfo := TokenInfo{}
	json.NewDecoder(resp.Body).Decode(&tokenInfo)
	defer resp.Body.Close()
	return &tokenInfo
}

// GetToken : Return token set by user or fresh token
// The source of the token is determined by the following logic:
// 1. If user provides their own managed access token, assume it is valid and send it
// 2. If this class is managing tokens and does not yet have one, make a request for one
// 3. If this class is managing tokens and the token has expired refresh it. In case the refresh token is expired, get a new one
// If this class is managing tokens and has a valid token stored, send it
func (tm *TokenManager) GetToken() string {
	if tm.userAccessToken != "" {
		return tm.userAccessToken
	} else if tm.tokenInfo.AccessToken == "" {
		tokenInfo := tm.requestToken()
		tm.saveTokenInfo(tokenInfo)
		return tm.tokenInfo.AccessToken
	} else if tm.isTokenExpired() {
		var tokenInfo *TokenInfo
		if tm.isRefreshTokenExpired() {
			tokenInfo = tm.requestToken()
		} else {
			tokenInfo = tm.refreshToken()
		}
		tm.saveTokenInfo(tokenInfo)
		return tm.tokenInfo.AccessToken
	}
	return tm.tokenInfo.AccessToken
}

// Request an IAM token using an API key
func (tm *TokenManager) requestToken() *TokenInfo {
	builder := NewRequestBuilder(POST).
		ConstructHTTPURL(tm.iamURL, nil, nil)

	builder.AddHeader(ContentType, DefaultContentType).
		AddHeader(Authorization, DefaultAuthorization).
		AddHeader(Accept, ApplicationJSON)

	data := map[string]string{
		"grant_type":    RequestTokenGrantType,
		"apikey":        tm.iamAPIkey,
		"response_type": RequestTokenResponseType,
	}
	builder.SetBodyContentJSON(data)

	req, err := builder.Build()
	if err != nil {
		panic(err)
	}
	return tm.request(req)
}

// Refresh an IAM token using a refresh token
func (tm *TokenManager) refreshToken() *TokenInfo {
	builder := NewRequestBuilder(POST).
		ConstructHTTPURL(tm.iamURL, nil, nil)

	builder.AddHeader(ContentType, DefaultContentType).
		AddHeader(Authorization, DefaultAuthorization).
		AddHeader(Accept, ApplicationJSON)

	data := map[string]string{
		"grant_type":    RefreshTokenGrantType,
		"refresh_token": tm.tokenInfo.RefreshToken,
	}
	builder.SetBodyContentJSON(data)

	req, err := builder.Build()
	if err != nil {
		panic(err)
	}
	return tm.request(req)
}

// SetAccessToken : sets a self-managed IAM access token.
// The access token should be valid and not yet expired.
func (tm *TokenManager) SetAccessToken(userAccessToken string) {
	tm.userAccessToken = userAccessToken
}

// SetIAMAPIKey : Set API key so that SDK manages token
func (tm *TokenManager) SetIAMAPIKey(key string) {
	tm.iamAPIkey = key
}

// Check if currently stored token is expired.
// Using a buffer to prevent the edge case of the
// oken expiring before the request could be made.
// The buffer will be a fraction of the total TTL. Using 80%.
func (tm *TokenManager) isTokenExpired() bool {
	buffer := 0.8
	expiresIn := tm.tokenInfo.ExpiresIn
	expireTime := tm.tokenInfo.Expiration
	refreshTime := expireTime - (expiresIn * int64(1.0-buffer))
	currTime := time.Now().Unix()
	return refreshTime < currTime
}

// Used as a fail-safe to prevent the condition of a refresh token expiring,
// which could happen after around 30 days. This function will return true
// if it has been at least 7 days and 1 hour since the last token was set
func (tm *TokenManager) isRefreshTokenExpired() bool {
	if tm.tokenInfo.Expiration == 0 {
		return true
	}

	sevenDays := int64(7 * 24 * 3600)
	currTime := time.Now().Unix()
	newTokenTime := tm.tokenInfo.Expiration + sevenDays
	return newTokenTime < currTime
}

// Save the response from the IAM service request to the object's state.
func (tm *TokenManager) saveTokenInfo(tokenInfo *TokenInfo) {
	tm.tokenInfo = tokenInfo
}
