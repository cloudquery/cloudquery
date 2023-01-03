package recipes

import (
	"github.com/hermanschaaf/hackernews"
)

func ItemResources() []*Resource {
	resources := []*Resource{
		{
			TableName:   "items",
			DataStruct:  &hackernews.Item{},
			Description: "https://github.com/HackerNews/API#items",
			PKColumns:   []string{"id"},
		},
	}
	for _, r := range resources {
		r.Service = "items"
	}
	return resources
}
