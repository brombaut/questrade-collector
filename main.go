package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/brombaut/questrade-collector/model"
	"github.com/spf13/viper"
)

type SSMap map[string]string

func main() {
	setup()
	refreshToken()
	accounts := getAccounts()
	for _, a := range accounts {
		getAccountBalances(a.Number)  // TODO: Write this to csv (with date and account details)
		getAccountPositions(a.Number) // TODO: Write this to csv (with date and account details)
	}
}

func setup() {
	viper.SetConfigFile(".env")
}

func refreshToken() {
	log.Print("refreshToken()")
	url := "https://login.questrade.com/oauth2/token"
	query := make(SSMap)
	headers := make(SSMap)
	query["grant_type"] = "refresh_token"
	query["refresh_token"] = readEnvVariable("QT_REFRESH_TOKEN")
	oa2Resp := model.OAuth2Response{}
	err := makeUnauthRequest(url, query, headers, &oa2Resp)
	if err != nil {
		log.Fatalln(err)
	}
	setEnvVariable("QT_REFRESH_TOKEN", oa2Resp.RefreshToken)
	setEnvVariable("QT_ACCESS_TOKEN", oa2Resp.AccessToken)
	setEnvVariable("QT_API_SERVER", oa2Resp.ApiServer)
}

func getAccounts() []model.Account {
	url := readEnvVariable("QT_API_SERVER") + "v1/accounts"
	query := make(SSMap)
	headers := make(SSMap)
	accountsResponse := model.AccountsResponse{}
	err := makeAuthRequest(url, query, headers, &accountsResponse)
	if err != nil {
		log.Fatalln(err)
	}
	log.Print(accountsResponse.TextOutput())
	return accountsResponse.Accounts
}

func getAccountBalances(accountNumber string) model.BalancesResponse {
	url := readEnvVariable("QT_API_SERVER") + "v1/accounts/" + accountNumber + "/balances"
	query := make(SSMap)
	headers := make(SSMap)
	log.Print(url)
	balancesResponse := model.BalancesResponse{}
	err := makeAuthRequest(url, query, headers, &balancesResponse)
	if err != nil {
		log.Fatalln(err)
	}
	log.Print(balancesResponse.TextOutput())
	return balancesResponse
}

func getAccountPositions(accountNumber string) []model.Position {
	url := readEnvVariable("QT_API_SERVER") + "v1/accounts/" + accountNumber + "/positions"
	query := make(SSMap)
	headers := make(SSMap)
	log.Print(url)
	positionsResponse := model.PositionsResponse{}
	err := makeAuthRequest(url, query, headers, &positionsResponse)
	if err != nil {
		log.Fatalln(err)
	}
	log.Print(positionsResponse.TextOutput())
	return positionsResponse.Positions
}

func makeAuthRequest(url string, query SSMap, headers SSMap, target interface{}) error {
	headers["Authorization"] = "Bearer " + readEnvVariable("QT_ACCESS_TOKEN")
	return makeUnauthRequest(url, query, headers, target)
}

func makeUnauthRequest(url string, query SSMap, headers SSMap, target interface{}) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	q := req.URL.Query()
	for key, value := range query {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(target)
}

func readEnvVariable(key string) string {
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}
	return value
}

func setEnvVariable(key string, value string) {
	viper.Set(key, value)
	viper.WriteConfig()
}
