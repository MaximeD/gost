package main
import (
  "flag"
  "github.com/MaximeD/gost/conf"
  "github.com/MaximeD/gost/gist"
)

var baseUrl string = "https://api.github.com/"

// get command line arguments
var gistDescriptionFlag = flag.String("description", "", "Description of the gist")
var gistPrivateFlag = flag.Bool("private", false, "Tells if the gist is private")
var listGistsFlag = flag.String("list", "", "List gists for a user")
func init() {
  flag.StringVar(gistDescriptionFlag, "d", "", "description")
  flag.BoolVar(gistPrivateFlag, "p", false, "private")
  flag.StringVar(listGistsFlag, "l", "", "list")
}


func main() {
  flag.Parse()
  isPublic := !*gistPrivateFlag

  if *listGistsFlag != "" {
    username := *listGistsFlag
    url := baseUrl + "users/" + username + "/gists"
    Gist.List(url)
  } else {
    token := Configuration.GetToken()
    filesName := flag.Args()
    Gist.Post(baseUrl, token, isPublic, filesName, *gistDescriptionFlag)
  }
}
