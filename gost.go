package main
import (
  "fmt"
  "bytes"
  "flag"
  "encoding/json"
  "net/http"
  "io/ioutil"
  "os"
  "./lib"
  "time"
)

// get command line arguments
var gistDescriptionFlag = flag.String("description", "", "Description of the gist")
func init() {
  flag.StringVar(gistDescriptionFlag, "d", "", "Description")
}

func main() {
  flag.Parse()

  /* common variables */
  baseUrl  := "https://api.github.com/"
  username := "MaximeD"
  url := baseUrl + "users/" + username + "/gists"

  listGists(url)
  postGist(baseUrl, *gistDescriptionFlag)
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
          fmt.Printf("(%s)\t%s\n", shortDate(val.Created_at), val.Description)
          fmt.Printf("\n")
        }
      }
    }
  }
}

type Gist struct {
  Desc string `json:"description"`
  Public bool `json:"public"`
  Files map[string]GistFile `json:"files"`
}

type GistFile struct {
  Content string `json:"content"`
}

func postGist(baseUrl string, description string) {
  // add some content for tests
  fileContent := "a test"
  files := make(map[string]GistFile)
  files["test.txt"] = GistFile{Content: fileContent}
  gist := Gist{Desc: description, Public: true, Files: files}

  // encode json
  buf, err := json.Marshal(gist)
  if err != nil {
    fmt.Printf("%s", err)
  }

  // print our shiny new json
  fmt.Println(string(buf))

  body := bytes.NewBuffer(buf)
  resp, err := http.Post(baseUrl + "gists", "text/json", body)
  fmt.Println(resp)
}

func shortDate(date time.Time) string {
  return date.Format("2006-01-02")
}
