package OAuthJSON

type GetSingleAuth struct {
	Scopes  []string `json:"scopes"`
	Note    string   `json:"note"`
	NoteUrl string   `json:"note_url"`
}

type GetSingleAuthResponse struct {
	Id        int
	Url       string
	Scopes    []string
	Token     string
	App       map[string]string
	Note      string
	NoteUrl   string
	CreatedAt string
	UpdatedAt string
}
