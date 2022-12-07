package recipes

import (
	"time"
)

type Emoji struct {
	Name        string    `json:"name"`
	UploadedBy  string    `json:"uploaded_by"`
	DateCreated time.Time `json:"date_created"`
}

func EmojiResources() []*Resource {
	resources := []*Resource{
		{
			DataStruct: &Emoji{},
			PKColumns:  []string{"name"},
		},
	}
	for _, r := range resources {
		r.Service = "emojis"
	}
	return resources
}
