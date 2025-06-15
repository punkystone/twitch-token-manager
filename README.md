
# Twitch Token Manager

## Methods

```go 
    func NewTokenManager(clientID, clientSecret string) *TokenManager
    func (manager *TokenManager) IsValid(accessToken string) (bool, error)
    func (manager *TokenManager) RefreshToken(refreshToken string) (bool, string, string, error)
    func (manager *TokenManager) ValidateAndRefreshToken(accessToken string, refreshToken string) (bool, string, string, error)
```

## Usage
```go
package main

import (
	manager "github.com/punkystone/twitch-token-manager"
)

func main() {
	accessToken := "accessToken"
	refreshToken := "refreshToken"
	tokenManager := manager.NewTokenManager("clientId", "clientSecret")

	isValid, err := tokenManager.IsValid(accessToken)

	if err != nil {
		println(err)
	}

	if isValid {
		println("Token is valid")
	} else {
		println("Token is invalid")
	}

	success, newAccessToken, newRefreshToken, err := tokenManager.RefreshToken(refreshToken)

	if err != nil {
		println(err)
		return
	}

	if success {
		println("New Access Token:", newAccessToken)
		println("New Refresh Token:", newRefreshToken)
	} else {
		println("Failed to refresh token")
	}

	success, newAccessToken, newRefreshToken, err = tokenManager.ValidateAndRefreshToken(accessToken, refreshToken)

	if err != nil {
		println(err)
		return
	}

	if success {
		if newAccessToken == accessToken && newRefreshToken == refreshToken {
			println("Access Token was valid and not refreshed")
		} else {
			println("New Access Token:", newAccessToken)
			println("New Refresh Token:", newRefreshToken)
		}
	} else {
		println("Failed to validate or refresh token")
	}
}
```
