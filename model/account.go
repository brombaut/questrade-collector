package model

import "fmt"

type Account struct {
	Type              string `json:"type"`
	Number            string `json:"number"`
	Status            string `json:"status"`
	IsPrimary         bool   `json:"isPrimary"`
	IsBilling         bool   `json:"isBilling"`
	ClientAccountType string `json:"clientAccountType"`
}

func (a Account) TextOutput() string {
	result := fmt.Sprintf("%+v", a)
	return result
}

type AccountsResponse struct {
	Accounts []Account `json:"accounts"`
	UserId   int       `json:"userId"`
}

func (ar AccountsResponse) TextOutput() string {
	result := fmt.Sprintf("%+v", ar)
	return result
}
