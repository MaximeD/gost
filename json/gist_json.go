package GistJSON

type Response struct {
	Url         string                 `json:"url"`
	ForksUrl    string                 `json:"forks_url"`
	CommitsUrl  string                 `json:"commits_url"`
	Id          string                 `json:"id"`
	GitPullUrl  string                 `json:"git_pull_url"`
	GitPushUrl  string                 `json:"git_push_url"`
	HtmlUrl     string                 `json:"html_url"`
	Files       map[string]FileDetails `json:"files"`
	Public      bool                   `json:"public"`
	CreatedAt   string                 `json:"created_at"`
	UpdatedAt   string                 `json:"updated_at"`
	Description string                 `json:"description"`
	Comments    int                    `json:"comments"`
	User        User                   `json:"user"`
	CommentsUrl string                 `json:"comments_url"`
}

type User struct {
	Login             string `json:"login"`
	Id                int64  `json:"id"`
	AvatarUrl         string `json:"avatar_url"`
	GravatarId        string `json:"gravatar_id"`
	Url               string `json:"url"`
	HtmlUrl           string `json:"html_url"`
	FollowersUrl      string `json:"followers_url"`
	FollowingUrl      string `json:"followings_url"`
	GistsUrl          string `json:"gists_url"`
	StarredUrl        string `json:"starred_url"`
	SubscriptionsUrl  string `json:"subscriptions_url"`
	OrganizationsUrl  string `json:"organizations_url"`
	ReposUrl          string `json:"repos_url"`
	EventsUrl         string `json:"events_url"`
	ReceivedEventsUrl string `json:"received_events_url"`
	TypeUrl           string `json:"type_url"`
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
	FileName string `json:"filename"`
	Type     string `json:"type"`
	Language string `json:"language"`
	RawUrl   string `json:"raw_url"`
	Size     int    `json:"size"`
	Content  string `json:"content"`
}

type MessageResponse struct {
	Message          string `json:"message"`
	DocumentationUrl string `json:"documentation_url"`
}
