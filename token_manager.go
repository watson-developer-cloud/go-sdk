package golang_sdk

import (
	"fmt"
	"time"
	req "github.com/parnurzeal/gorequest"
)

type TokenInfo struct {
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType string `json:"token_type"`
	ExpiresIn int64 `json:"expires_in"`
	Expiration int64 `json:"expiration"`
}

type TokenManager struct {
	userToken string
	iamAPIkey string
	iamURL string
	tokenInfo TokenInfo
}

func NewTokenManager() *TokenManager {
	return &TokenManager{
		iamURL: "https://iam.bluemix.net/identity/token",
		tokenInfo: TokenInfo{},
	}
}

func (tm *TokenManager) GetToken() (string, []error) {
	if tm.userToken != "" {
		return tm.userToken, nil
	}

	var err []error

	if tm.tokenInfo.AccessToken == "" {
		err = tm.PostToken(tm.RequestTokenBody())
	}

	if tm.IsTokenExpired() {
		if tm.IsRefreshTokenExpired() {
			err = tm.PostToken(tm.RequestTokenBody())
		} else {
			err = tm.PostToken(tm.RefreshTokenBody())
		}
	}

	if err != nil {
		return "", err
	}

	return tm.tokenInfo.AccessToken, nil
}

func (tm *TokenManager) RequestTokenBody() map[string]string {
	return map[string]string {
		"grant_type": "urn:ibm:params:oauth:grant-type:apikey",
		"apikey": tm.iamAPIkey,
		"response_type": "cloud_iam",
	}
}

func (tm *TokenManager) RefreshTokenBody() map[string]string {
	return map[string]string {
		"grant_type": "refresh_token",
		"refresh_token": tm.tokenInfo.RefreshToken,
	}
}

func (tm *TokenManager) PostToken(body map[string]string) []error {
	res, _, err := req.New().Post(tm.iamURL).
		Set("Accept", "application/json").
		Set("Content-Type", "application/x-www-form-urlencoded").
		Set("Authorization", "Basic Yng6Yng=").
		Send(body).
		EndStruct(&tm.tokenInfo)

	if err != nil {
		return err
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		err = append(err, fmt.Errorf(res.Status))
		return err
	}

	return nil
}

func (tm *TokenManager) SetToken(token string) {
	tm.userToken = token
}

func (tm *TokenManager) SetKey(key string) {
	tm.iamAPIkey = key
}

func (tm *TokenManager) IsTokenExpired() bool {
	buffer := 0.8
	expiresIn := tm.tokenInfo.ExpiresIn
	expireTime := tm.tokenInfo.Expiration
	refreshTime := expireTime - (expiresIn * int64(1.0 - buffer))

	currTime := time.Now().Unix()

	return refreshTime < currTime
}

func (tm *TokenManager) IsRefreshTokenExpired() bool {
	expiresIn := int64(7 * 24 * 3600)
	refreshTime := tm.tokenInfo.Expiration + expiresIn

	currTime := time.Now().Unix()

	return refreshTime < currTime
}
