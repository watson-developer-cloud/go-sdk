package core

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	assert "github.com/stretchr/testify/assert"
)

func TestRequestToken(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{
			"access_token": "oAeisG8yqPY7sFR_x66Z15",
			"token_type": "Bearer",
			"expires_in": 3600,
			"expiration": 1524167011,
			"refresh_token": "jy4gl91BQ"
		}`)
	}))
	defer server.Close()

	tokenManager := NewTokenManager("", server.URL, "")
	tokenInfo := tokenManager.requestToken()
	assert.Equal(t, tokenInfo.AccessToken, "oAeisG8yqPY7sFR_x66Z15")
}

func TestRefreshToken(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{
			"access_token": "oAeisG8yqPY7sFR_x66Z15",
			"token_type": "Bearer",
			"expires_in": 3600,
			"expiration": 1524167011,
			"refresh_token": "jy4gl91BQ"
		}`)
	}))
	defer server.Close()

	tokenManager := NewTokenManager("", server.URL, "")
	tokenInfo := tokenManager.refreshToken()
	assert.Equal(t, tokenInfo.AccessToken, "oAeisG8yqPY7sFR_x66Z15")
}

func TestIsTokenExpired(t *testing.T) {
	tokenManager := NewTokenManager("iamApiKey", "", "")
	tokenManager.tokenInfo = &TokenInfo{
		AccessToken:  "oAeisG8yqPY7sFR_x66Z15",
		TokenType:    "Bearer",
		ExpiresIn:    3600,
		Expiration:   time.Now().Unix() + 6000,
		RefreshToken: "jy4gl91BQ",
	}

	assert.Equal(t, tokenManager.isTokenExpired(), false)
	tokenManager.tokenInfo.Expiration = time.Now().Unix() - 3600
	assert.Equal(t, tokenManager.isTokenExpired(), true)
}

func TestIsRefreshTokenExpired(t *testing.T) {
	tokenManager := NewTokenManager("iamApiKey", "", "")
	tokenManager.tokenInfo = &TokenInfo{
		AccessToken:  "oAeisG8yqPY7sFR_x66Z15",
		TokenType:    "Bearer",
		ExpiresIn:    3600,
		Expiration:   time.Now().Unix(),
		RefreshToken: "jy4gl91BQ",
	}

	assert.Equal(t, tokenManager.isRefreshTokenExpired(), false)
	tokenManager.tokenInfo.Expiration = time.Now().Unix() - (8 * 24 * 3600)
	assert.Equal(t, tokenManager.isRefreshTokenExpired(), true)
}

func TestGetToken(t *testing.T) {
	// # Case 1:
	tokenManager := NewTokenManager("iamApiKey", "", "user access token")
	token := tokenManager.GetToken()
	assert.Equal(t, token, "user access token")

	// Case 2:
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{
			"access_token": "hellohello",
			"token_type": "Bearer",
			"expires_in": 3600,
			"expiration": 1524167011,
			"refresh_token": "jy4gl91BQ"
		}`)
	}))
	defer server.Close()
	tokenManager = NewTokenManager("iamApiKey", server.URL, "")
	token = tokenManager.GetToken()
	assert.Equal(t, token, "hellohello")

	// Case 3:
	server2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{
			"access_token": "hellohello",
			"token_type": "Bearer",
			"expires_in": 3600,
			"expiration": 1524167011,
			"refresh_token": "jy4gl91BQ"
		}`)
		body, _ := ioutil.ReadAll(r.Body)
		assert.Contains(t, string(body), "grant_type=urn")
	}))
	defer server2.Close()
	tokenManager = NewTokenManager("iamApiKey", server2.URL, "")
	tokenManager.tokenInfo.Expiration = time.Now().Unix() - (20 * 24 * 3600)
	tokenManager.GetToken()
}
