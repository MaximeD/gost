package Configuration

import (
	"fmt"
	"github.com/MaximeD/gost/conf/oauth"
	"io/ioutil"
	"os"
)

var homeDir string = os.Getenv("HOME")
var configurationFilePath string = homeDir + "/.gost"

func GetToken() (token string) {
	return readConf()
}

func readConf() (token string) {
	file, err := ioutil.ReadFile(configurationFilePath)

	if err != nil {
		// file does not exist
		return createConfigurationFile()
	}
	return string(file)
}

func createConfigurationFile() string {
	var OAuthToken string

	fmt.Println("You don't have any configuration file")
	fmt.Println("Do you want to create one? [Y/n]")
	var answer string
	fmt.Scanln(&answer)

	if answer == "y" || answer == "Y" || answer == "" {
		OAuthToken = OAuth.GetToken()
		ioutil.WriteFile(configurationFilePath, []byte(OAuthToken), 0660)
		fmt.Printf("[configuration file written to '%s']\n", configurationFilePath)
	} else {
		fmt.Println("Posting an anonymous gist...")
		fmt.Println("I'll ask you again next time!")
	}

	return OAuthToken
}
