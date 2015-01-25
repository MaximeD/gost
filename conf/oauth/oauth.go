package OAuth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/MaximeD/gost/json/oauth"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var baseUrl string = "https://api.github.com/"
var authorizationUrl string = baseUrl + "authorizations"

// GetToken will query github api to create an authentication token for gost.
// This allows user to create gists on his behalf rather than anonymously.
// It returns the created token.
func GetToken() string {
	client := &http.Client{}
	var authorizationResponseBody []byte
	var twoFA string

	username, password := askUserForCredentials()
	authenticationJson := getAuthenticationJson(false)

	resp := makeBasicAuthRequest(client, username, password, authenticationJson, twoFA)
	defer resp.Body.Close()

	// check for 2fa
	if resp.StatusCode == 401 && resp.Header.Get("X-Github-Otp") != "" {
		twoFA = askUserForTwoFA()
		resp = makeBasicAuthRequest(client, username, password, authenticationJson, twoFA)
		defer resp.Body.Close()
	}

	// user might already have registered a token for this application
	if resp.StatusCode == 422 {
		authenticationJson = getAuthenticationJson(true)
		resp = makeBasicAuthRequest(client, username, password, authenticationJson, twoFA)
		defer resp.Body.Close()
	}

	if resp.StatusCode != 201 {
		fmt.Println("Sorry but we could not authenticate you.")
		fmt.Println("No gist were created...")
		os.Exit(1)
	}

	authorizationResponseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	var jsonRes OAuthJSON.GetSingleAuthResponse
	err = json.Unmarshal(authorizationResponseBody, &jsonRes)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	return jsonRes.Token
}

// askUserForCredentials will prompt the user to enter his username and password.
// It returns user's username and user's password.
func askUserForCredentials() (username string, password string) {
	fmt.Println("GitHub username:")
	fmt.Scanln(&username)
	fmt.Println("GitHub password:")
	fmt.Scanln(&password)

	return username, password
}

// makeBasicAuthRequest will create a POST request to create a new authorization
// with basic auth enabled.
// See: https://developer.github.com/v3/oauth_authorizations/#create-a-new-authorization
// It takes user's username, user's password, marshalled json and 2fa code if user has enabled it.
// It returns a new Request.
func makeBasicAuthRequest(client *http.Client, username string, password string, marshalledJson []byte, twoFA string) *http.Response {
	jsonBody := bytes.NewBuffer(marshalledJson)

	req, err := http.NewRequest("POST", authorizationUrl, jsonBody)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
	req.SetBasicAuth(username, password)
	if twoFA != "" {
		req.Header.Add("X-GitHub-OTP", twoFA)
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	return resp
}

// getAuthenticationJson creates a marshalled json to be used to create a new authorization.
// See: https://developer.github.com/v3/oauth_authorizations/#create-a-new-authorization
// It takes a boolean parameter telling if note should include a timestamp.
// You will need this in case user already has a token for this application and needs to create a new one.
// Otherwise github will return an empty token.
// It returns the marshalled json.
func getAuthenticationJson(withTimestamp bool) []byte {
	description := "gost"

	if withTimestamp {
		fmt.Println("You already have a personal access token for gost, though I cannot find it on your computer...")
		fmt.Println("I will create another access token for you.")
		fmt.Println("(You can see them here https://github.com/settings/applications)")
		description = fmt.Sprintf("%s (%s)", description, time.Now())
	}

	authorization := OAuthJSON.GetSingleAuth{Scopes: []string{"gist"}, Note: description, NoteUrl: "https://github.com/MaximeD/gost"}
	marshalledJson, err := json.Marshal(authorization)
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	return marshalledJson
}

// askUserForTwoFA will prompt user to enter his 2f-a code if GitHub requests one
// It returns the 2f-a code.
func askUserForTwoFA() string {
	var twoFA string

	fmt.Println("You have enabled two-factor authentication, please enter the code GitHub sent you:")
	fmt.Scanln(&twoFA)

	return twoFA
}
