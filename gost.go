package main
import (
  "flag"
  "github.com/MaximeD/gost/gist"
)

var baseUrl string = "https://api.github.com/"

// get command line arguments
var gistDescriptionFlag = flag.String("description", "", "Description of the gist")
var listGistsFlag = flag.String("list", "", "List gists for a user")
func init() {
  flag.StringVar(gistDescriptionFlag, "d", "", "Description")
  flag.StringVar(listGistsFlag, "l", "", "list")
}


func main() {
  flag.Parse()

  if *listGistsFlag != "" {
    username := *listGistsFlag
    url := baseUrl + "users/" + username + "/gists"
    Gist.List(url)
  } else {
    filesName := flag.Args()
    /* common variables */
    Gist.Post(baseUrl, filesName, *gistDescriptionFlag)
  }
}
