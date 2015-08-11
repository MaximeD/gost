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
var deleteGistFlag = flag.String("delete", "", "Delete a gist")
var downloadGistFlag = flag.String("get", "", "Get a single gist")
var gistDescriptionFlag = flag.String("description", "", "Description of the gist")
var gistPrivateFlag = flag.Bool("private", false, "Set gist to private")
var listGistsFlag = flag.String("list", "", "List gists for a user")
var openBrowserFlag = flag.Bool("open", false, "Open result in browser")
var updateGistFlag = flag.String("update", "", "Update an existing gist")

func init() {
	flag.BoolVar(gistPrivateFlag, "p", false, "Set gist to private")
	flag.BoolVar(openBrowserFlag, "o", false, "Open result in browser")
	flag.StringVar(deleteGistFlag, "D", "", "Delete a gist")
	flag.StringVar(downloadGistFlag, "g", "", "Get a single gist")
	flag.StringVar(gistDescriptionFlag, "d", "", "Description of the gist")
	flag.StringVar(listGistsFlag, "l", "", "List gists for a user")
	flag.StringVar(updateGistFlag, "u", "", "Update an existing gist")
}

func main() {
	flag.Parse()
	isPublic := !*gistPrivateFlag

	// if nothing was given, display help
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
		Gist.List(url, token)
	} else if *deleteGistFlag != "" {
		Gist.Delete(baseUrl, token, *deleteGistFlag)
	} else if *downloadGistFlag != "" {
		Gist.Download(baseUrl, token, *downloadGistFlag)
	} else {
		filesName := flag.Args()
		if len(filesName) == 0 && *updateGistFlag == "" {
			fmt.Println("No files given!")
			os.Exit(2)
		}
		if *updateGistFlag != "" {
			Gist.Update(baseUrl, token, filesName, *updateGistFlag, *gistDescriptionFlag, *openBrowserFlag)
		} else {
			Gist.Post(baseUrl, token, isPublic, filesName, *gistDescriptionFlag, *openBrowserFlag)
		}
	}
}
