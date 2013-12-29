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
	scopes := []string{"gist"}
	authorization := OAuthJSON.GetSingleAuth{Scopes: scopes, Note: "gost", NoteUrl: "https://github.com/MaximeD"}
	buf, err := json.Marshal(authorization)

	// encode json
	if err != nil {
		fmt.Printf("%s", err)
	}
	body := bytes.NewBuffer(buf)

	// create client to handle basic auth
	client := &http.Client{}
	req, err := http.NewRequest("POST", authorizationUrl, body)
	username, password := getCredentials()
	req.SetBasicAuth(username, password)

	// post json
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		// close connexion
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		} else {
			var jsonRes OAuthJSON.GetSingleAuthResponse
			err := json.Unmarshal(body, &jsonRes)
			if err != nil {
				fmt.Printf("%s", err)
				os.Exit(1)
			} else {
				token = jsonRes.Token
			}
		}
	}
	return token
}

func getCredentials() (username string, password string) {
	fmt.Println("GitHub username:")
	fmt.Scanln(&username)
	fmt.Println("GitHub password:")
	fmt.Scanln(&password)

	return username, password
}
