package pages

import "time"

type Pages struct {
	Object     string `json:"object"`
	Results    []Page `json:"results"`
	NextCursor string `json:"next_cursor"`
	HasMore    bool   `json:"has_more"`
}

type Page struct {
	Object         string                 `json:"object"`
	Id             string                 `json:"id"`
	CreatedTime    time.Time              `json:"created_time"`
	LastEditedTime time.Time              `json:"last_edited_time"`
	CreatedBy      map[string]interface{} `json:"created_by"`
	LastEditedBy   map[string]interface{} `json:"last_edited_by"`
	Cover          map[string]interface{} `json:"cover"`
	Icon           map[string]interface{} `json:"icon"`
	Parent         map[string]interface{} `json:"parent"`
	Archived       bool                   `json:"archived"`
	Properties     map[string]interface{} `json:"properties"`
	Url            string                 `json:"url"`
	PublicUrl      string                 `json:"public_url"`
}
