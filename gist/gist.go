package Gist

import (
  "fmt"
  "time"
  "bytes"
  "encoding/json"
  "net/http"
  "io/ioutil"
  "github.com/MaximeD/gost/json"
  "os"
)

func List(url string) {
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
      var json_res []JSONStruct.Response
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

func Post(baseUrl string, filesName []string, description string) {
  files := make(map[string]JSONStruct.File)

  for i:=0; i < len(filesName); i++ {
    content, err := ioutil.ReadFile(filesName[i])
    if err != nil {
      fmt.Printf("%s", err)
      os.Exit(1)
    }
    files[filesName[i]] = JSONStruct.File{Content: string(content)}
  }

  gist := JSONStruct.Post{Desc: description, Public: true, Files: files}

  // encode json
  buf, err := json.Marshal(gist)
  if err != nil {
    fmt.Printf("%s", err)
  }
  body := bytes.NewBuffer(buf)

  // post json
  resp, err := http.Post(baseUrl + "gists", "text/json", body)
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
      var json_res JSONStruct.Response
      err := json.Unmarshal(body, &json_res)
      if err != nil {
        fmt.Printf("%s", err)
        os.Exit(1)
      } else {
        fmt.Printf("%s\n", json_res.Html_url)
      }
    }
  }
}

func shortDate(date time.Time) string {
  return date.Format("2006-01-02")
}
