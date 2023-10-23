package users

type Users struct {
	Object     string `json:"object"`
	Results    []User `json:"results"`
	NextCursor string `json:"next_cursor"`
	HasMore    bool   `json:"has_more"`
}

type User struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	Object    string `json:"object"`
	AvatarUrl string `json:"avatar_url"`
}
