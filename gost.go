package main
import (
  "fmt"
  "encoding/json"
  "net/http"
  "io/ioutil"
  "os"
  "./lib"
)


func main() {
  /* common variables */
  base_url  := "https://api.github.com/"
  var username string

  fmt.Printf("Enter username\n")
  fmt.Scanf("%s", &username)

  url := base_url + "users/" + username + "/gists"

  listGists(url)
}

func listGists(url string) {
  res, err := http.Get(url)

  if err != nil {
    fmt.Printf("%s", err)
    os.Exit(1)
  } else {
    // close connexion
    defer res.Body.Close()
    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
      fmt.Printf("%s", err)
      os.Exit(1)
    } else {
      var json_res []GistTypes.Response
      err := json.Unmarshal(body, &json_res)
      if err != nil {
        fmt.Printf("%s", err)
        os.Exit(1)
      } else {
        for _,val := range json_res {
          fmt.Printf("%s\n", val.Html_url)
          fmt.Printf("\t%s\n", val.Description)
          fmt.Printf("\n")
        }
      }
    }
  }
}
