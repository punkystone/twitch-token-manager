package manager

import "net/http"

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
