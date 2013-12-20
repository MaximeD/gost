package JSONStruct

import "time"

/* define types of json response */
type Response struct {
  Url string
  Forks_url string
  Commits_url string
  Id string
  Git_pull_url string
  Git_push_url string
  Html_url string
  Files Files
  Public bool
  Created_at time.Time
  Updated_at time.Time
  Description string
  Comments int
  User User
  Comments_url string
}

type User struct {
  Login string
  Id int64
  AvatarUrl string
  GravatarId string
  Url string
  HtmlUrl string
  FollowersUrl string
  FollowingUrl string
  GistsUrl string
  StarredUrl string
  SubscriptionsUrl string
  OrganizationsUrl string
  ReposUrl string
  EventsUrl string
  ReceivedEventsUrl string
  TypeUrl string
}

type Post struct {
  Desc string `json:"description"`
  Public bool `json:"public"`
  Files map[string]File `json:"files"`
}

type File struct {
  Content string `json:"content"`
}


type Files struct {
  //"http_response.coffee": {
  //  "filename": "http_response.coffee",
  //  "type": "text/coffescript",
  //  "language": "CoffeeScript",
  //  "raw_url": "https://gist.github.com/raw/6163410/a3a11aa559da8bc0e40828742cf482f25a4a54c2/http_response.coffee",
  //  "size": 1113
}
