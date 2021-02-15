package model

import "fmt"

type OAuth2Response struct {
	AccessToken  string `json:"access_token"`
	ApiServer    string `json:"api_server"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
}

func (oa2 OAuth2Response) TextOutput() string {
	result := fmt.Sprintf(
		"AccessToken: %s\nApiServer : %s\nExpiresIn: %d\nRefreshToken: %s\nTokenType: %s\n",
		oa2.AccessToken, oa2.ApiServer, oa2.ExpiresIn, oa2.RefreshToken, oa2.TokenType)
	return result
}
