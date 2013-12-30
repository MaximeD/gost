package GistJSON

type Response struct {
	Url          string
	Forks_url    string
	Commits_url  string
	Id           string
	Git_pull_url string
	Git_push_url string
	Html_url     string
	Files        map[string]FileDetails
	Public       bool
	Created_at   string
	Updated_at   string
	Description  string
	Comments     int
	User         User
	Comments_url string
}

type User struct {
	Login             string
	Id                int64
	AvatarUrl         string
	GravatarId        string
	Url               string
	HtmlUrl           string
	FollowersUrl      string
	FollowingUrl      string
	GistsUrl          string
	StarredUrl        string
	SubscriptionsUrl  string
	OrganizationsUrl  string
	ReposUrl          string
	EventsUrl         string
	ReceivedEventsUrl string
	TypeUrl           string
}

type Post struct {
	Desc   string          `json:"description"`
	Public bool            `json:"public"`
	Files  map[string]File `json:"files"`
}

type File struct {
	Content string `json:"content"`
}

type FileDetails struct {
	FileName string
	Type     string
	Language string
	RawUrl   string
	Size     int
  Content  string
}

type MessageResponse struct {
  Message string
  DocumentationUrl string
}
