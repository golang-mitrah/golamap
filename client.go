package golamap

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type OLAMap struct {
	Token       string   // Ola map token
	RequestId   string   // Unique UUID for a request
	HttpService HttpServ // HTTP service interface
}

type HttpServ interface {
	SendOlaMapRequest(method, url, requestID, oauthToken string, responseObj interface{}) error
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

// Initialize the Olamap with X-RequestID
func Initialize(requestID string) *OLAMap {
	httpReq := &OlaRequest{}
	return &OLAMap{
		RequestId:   requestID,
		HttpService: httpReq,
	}
}

// Configure OLA access token
func (o *OLAMap) ConfigureAccessToken(clientID, clientSecret string) error {
	form := url.Values{}
	form.Set("grant_type", "client_credentials")
	form.Set("scope", "openid")
	form.Set("client_id", clientID)
	form.Set("client_secret", clientSecret)

	req, err := http.NewRequest("POST", TokenURL, strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("Failed to get token - statuscode %s", resp.StatusCode))
	}

	var tokenResponse TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		return err
	}

	o.Token = "Bearer " + tokenResponse.AccessToken

	return nil
}
