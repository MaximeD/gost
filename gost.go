package main
import (
  "flag"
  "github.com/MaximeD/gost/conf"
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
    token := Configuration.GetToken()
    filesName := flag.Args()
    Gist.Post(baseUrl, token, filesName, *gistDescriptionFlag)
  }
}
