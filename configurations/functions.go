package configurations

import (
	"bytes"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// function to return the base path for endpoints

func baseEndpoint(env string) string {
	var endpoint string
	if env == "SANDBOX" {
		endpoint = "https://sandbox.safaricom.co.ke/mpesa/"
	} else if env == "LIVE" {
		endpoint = "https://sandbox.safaricom.co.ke/api/"
	}
	return endpoint
}

func basicAuth(consumerKey, appSecret string) string {
	auth := consumerKey + ":" + appSecret
	return b64.StdEncoding.EncodeToString([]byte(auth))
}

func generateToken(env string) string {
	var auth_endpoint, consumerKey, appSecret string
	if env == "SANDBOX" {
		auth_endpoint = endpoints["SANDBOX_ENDPOINTS"]["AUTHENTICATION"]
		consumerKey = daraja["SANDBOX_CONFIGS"]["APP_CONSUMER_KEY"]
		appSecret = daraja["SANDBOX_CONFIGS"]["APP_SECRET_KEY"]
	} else if env == "LIVE" {
		auth_endpoint = endpoints["LIVE_ENDPOINTS"]["AUTHENTICATION"]
		consumerKey = daraja["LIVE_CONFIGS"]["APP_CONSUMER_KEY"]
		appSecret = daraja["LIVE_CONFIGS"]["APP_SECRET_KEY"]
	}

	request, err := http.NewRequest("GET", auth_endpoint, nil)

	if err != nil {
		panic(err)
	}

	request.Header.Add("Authorization", "Basic "+basicAuth(consumerKey, appSecret))

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	type responseMap map[string]string

	var data responseMap

	err = json.Unmarshal([]byte(body), &data)
	if err != nil {
		log.Fatal(err)
	}

	accessToken := data["access_token"]

	defer response.Body.Close()

	return accessToken
}

func RegisterURLs(env, responseType string) string {
	var register_urls_endpoint, confirmation_url, validation_url, shortcode string
	if env == "SANDBOX" {
		shortcode = daraja["SANDBOX_CONFIGS"]["C2B_SHORTCODE"]
		register_urls_endpoint = endpoints["SANDBOX_ENDPOINTS"]["C2B_REGISTER_URLS"]
		confirmation_url = callbacks["CONFIRMATION_URL"]
		validation_url = callbacks["VALIDATION_URL"]
	} else if env == "LIVE" {
		shortcode = daraja["LIVE_CONFIGS"]["C2B_SHORTCODE"]
		register_urls_endpoint = endpoints["LIVE_ENDPOINTS"]["C2B_REGISTER_URLS"]
		confirmation_url = callbacks["CONFIRMATION_URL"]
		validation_url = callbacks["VALIDATION_URL"]
	}

	postData := RegisterUrl{
		ShortCode:       shortcode,
		ResponseType:    responseType,
		ConfirmationURL: confirmation_url,
		ValidationURL:   validation_url,
	}

	jsonData, err := json.Marshal(postData)

	if err != nil {
		fmt.Println("error:", err)
	}

	request, err := http.NewRequest("POST", register_urls_endpoint, bytes.NewBuffer(jsonData))

	if err != nil {
		log.Fatal(err)
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+generateToken(enviroment))

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	return string(body)
}

func SimulateC2B(command, phoneNumber, accountNo string, amount int) string {
	endPoint := endpoints["SANDBOX_ENDPOINTS"]["SIMULATE_C2B"]
	shortcode := daraja["SANDBOX_CONFIGS"]["B2B_B2C_SHORTCODE"]

	postData := simulateC2BStruct{
		ShortCode:   shortcode,
		CommandId:   command,
		PhoneNumber: phoneNumber,
		Amount:      amount,
		AccountNo:   accountNo,
	}

	jsonData, err := json.Marshal(postData)

	if err != nil {
		fmt.Println("error:", err)
	}

	request, err := http.NewRequest("POST", endPoint, bytes.NewBuffer(jsonData))

	if err != nil {
		log.Fatal(err)
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+generateToken(enviroment))

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	return string(body)
}
