package manager

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

type TokenManager struct {
	clientID     string
	clientSecret string
}

func NewTokenManager(clientID, clientSecret string) *TokenManager {
	return &TokenManager{
		clientID,
		clientSecret,
	}
}

func (manager *TokenManager) IsValid(accessToken string) (bool, error) {
	request, err := http.NewRequest(http.MethodGet, "https://id.twitch.tv/oauth2/validate", nil)
	if err != nil {
		return false, err
	}
	request.Header.Add("Authorization", "Bearer "+accessToken)
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return false, err
	}
	defer response.Body.Close()
	return response.StatusCode == http.StatusOK, nil
}

func (manager *TokenManager) RefreshToken(refreshToken string) (bool, string, string, error) {
	data := url.Values{}
	data.Add("client_id", manager.clientID)
	data.Add("client_secret", manager.clientSecret)
	data.Add("grant_type", "refresh_token")
	data.Add("refresh_token", refreshToken)
	response, err := http.Post("https://id.twitch.tv/oauth2/token", "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
	if err != nil {
		return false, "", "", err
	}
	if response.StatusCode == http.StatusBadRequest {
		return false, "", "", nil
	}
	defer response.Body.Close()
	tokenResponse := &TokenResponse{}
	err = json.NewDecoder(response.Body).Decode(tokenResponse)
	if err != nil {
		return false, "", "", err
	}
	return true, tokenResponse.AccessToken, tokenResponse.RefreshToken, nil
}

func (manager *TokenManager) ValidateAndRefreshToken(accessToken string, refreshToken string) (bool, string, string, error) {
	valid, err := manager.IsValid(accessToken)
	if err != nil {
		return false, accessToken, refreshToken, err
	}
	if valid {
		return true, accessToken, refreshToken, nil
	}
	return manager.RefreshToken(refreshToken)
}
