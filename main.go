package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/brombaut/questrade-collector/model"
	"github.com/spf13/viper"
)

func main() {
	setup()
	refreshToken()
}

func setup() {
	viper.SetConfigFile(".env")
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
	writeCurrEnvVariables()
}

func writeCurrEnvVariables() {
	viper.WriteConfig()
}

func refreshToken() {
	url := "https://login.questrade.com/oauth2/token"
	query := make(map[string]string)
	headers := make(map[string]string)
	query["grant_type"] = "refresh_token"
	query["refresh_token"] = readEnvVariable("QT_REFRESH_TOKEN")
	oa2Resp := model.OAuth2Response{}
	err := makeUnauthRequest(url, query, headers, &oa2Resp)
	if err != nil {
		log.Fatalln(err)
	}
	log.Print(oa2Resp)
	setEnvVariable("QT_REFRESH_TOKEN", oa2Resp.RefreshToken)
	setEnvVariable("QT_ACCESS_TOKEN", oa2Resp.AccessToken)
	setEnvVariable("AT_API_SERVER", oa2Resp.ApiServer)
}

func makeAuthRequest() {
	headers := make(map[string]string)
	headers["Authorization"] = "Bearer " + readEnvVariable("QT_ACCESS_TOKEN")
}

func makeUnauthRequest(url string, query map[string]string, headers map[string]string, target interface{}) error {
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
