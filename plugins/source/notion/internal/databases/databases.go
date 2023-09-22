package databases

import "time"

type Databases struct {
	Object     string     `json:"object"`
	Results    []Database `json:"results"`
	NextCursor string     `json:"next_cursor"`
	HasMore    bool       `json:"has_more"`
}

type Database struct {
	Object         string           `json:"object"`
	Id             string           `json:"id"`
	CreatedTime    time.Time        `json:"created_time"`
	LastEditedTime time.Time        `json:"last_edited_time"`
	CreatedBy      map[string]any   `json:"created_by"`
	LastEditedBy   map[string]any   `json:"last_edited_by"`
	Title          []map[string]any `json:"title"`
	Description    []map[string]any `json:"description"`
	IsInline       bool             `json:"is_inline"`
	Cover          map[string]any   `json:"cover"`
	Icon           map[string]any   `json:"icon"`
	Parent         map[string]any   `json:"parent"`
	Archived       bool             `json:"archived"`
	Properties     map[string]any   `json:"properties"`
	Url            string           `json:"url"`
	PublicUrl      string           `json:"public_url"`
}
