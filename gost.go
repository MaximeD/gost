package main

import (
	"flag"
	"fmt"
	"github.com/MaximeD/gost/conf"
	"github.com/MaximeD/gost/gist"
	"os"
)

var baseUrl string = "https://api.github.com/"

// get command line arguments
var gistDescriptionFlag = flag.String("description", "", "Description of the gist")
var gistPrivateFlag = flag.Bool("private", false, "Set gist to private")
var listGistsFlag = flag.String("list", "", "List gists for a user")
var deleteGistFlag = flag.String("delete", "", "Delete a gist")
var downloadGistFlag = flag.String("get", "", "Get a single gist")

func init() {
	flag.StringVar(gistDescriptionFlag, "d", "", "Description of the gist")
	flag.BoolVar(gistPrivateFlag, "p", false, "Set gist to private")
	flag.StringVar(listGistsFlag, "l", "", "List gists for a user")
	flag.StringVar(deleteGistFlag, "D", "", "Delete a gist")
	flag.StringVar(downloadGistFlag, "g", "", "Get a single gist")
}

func main() {
	flag.Parse()
	isPublic := !*gistPrivateFlag

	// if nothing was given write message
	if (flag.NFlag() == 0) && (len(flag.Args()) == 0) {
		fmt.Println("No arguments or files given!")
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(2)
	}

	token := Configuration.GetToken()

	if *listGistsFlag != "" {
		username := *listGistsFlag
		url := baseUrl + "users/" + username + "/gists"
		Gist.List(url)
	} else if *deleteGistFlag != "" {
		Gist.Delete(baseUrl, token, *deleteGistFlag)
	} else if *downloadGistFlag != "" {
		Gist.Download(baseUrl, token, *downloadGistFlag)
	} else {
		filesName := flag.Args()
		if len(filesName) == 0 {
			fmt.Println("No files given!")
			os.Exit(2)
		}
		Gist.Post(baseUrl, token, isPublic, filesName, *gistDescriptionFlag)
	}
}
