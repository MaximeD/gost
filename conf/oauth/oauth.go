package OAuth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/MaximeD/gost/json/oauth"
	"io/ioutil"
	"net/http"
	"os"
)

var baseUrl string = "https://api.github.com/"
var authorizationUrl string = baseUrl + "authorizations"

func GetToken() (token string) {
	client := &http.Client{}
	var resp *http.Response
	var authorizationResponseBody []byte

	// json structure
	scopes := []string{"gist"}
	authorization := OAuthJSON.GetSingleAuth{Scopes: scopes, Note: "gost", NoteUrl: "https://github.com/MaximeD"}
	encodedJson, err := json.Marshal(authorization)
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	// ask for user credentials
	username, password := getCredentials()

	// make request
	req := makeRequest(username, password, encodedJson)

	// post json
	resp, err = client.Do(req)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
	authorizationResponseBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	// check github response
	if resp.StatusCode == 401 {
		header := resp.Header

		// user may have enable 2fa
		if header.Get("X-Github-Otp") != "" {
			fmt.Println("2-factor authentication code:")
			var twoFA string
			fmt.Scanln(&twoFA)

			// make request
			req := makeRequest(username, password, encodedJson)
			req.Header.Add("X-Github-Otp", twoFA)

			resp, err = client.Do(req)
			if err != nil {
				fmt.Printf("%s\n", err)
				os.Exit(1)
			}

			authorizationResponseBody, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Printf("%s\n", err)
				os.Exit(1)
			}

			// has authorization been created?
			if resp.StatusCode != 201 {
				fmt.Println("Sorry but we could not authenticate you")
				os.Exit(1)
			}
		} else {
			fmt.Println("Sorry but we could not authenticate you")
			os.Exit(1)
		}
	}

	var jsonRes OAuthJSON.GetSingleAuthResponse
	err = json.Unmarshal(authorizationResponseBody, &jsonRes)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	token = jsonRes.Token
	return token
}

func getCredentials() (username string, password string) {
	fmt.Println("GitHub username:")
	fmt.Scanln(&username)
	fmt.Println("GitHub password:")
	fmt.Scanln(&password)

	return username, password
}

func makeRequest(username string, password string, encodedJson []byte) *http.Request {
	jsonBody := bytes.NewBuffer(encodedJson)
	req, err := http.NewRequest("POST", authorizationUrl, jsonBody)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
	req.SetBasicAuth(username, password)

	return req
}
