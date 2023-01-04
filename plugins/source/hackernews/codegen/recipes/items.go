package recipes

import (
	"github.com/hermanschaaf/hackernews"
)

func ItemResources() []*Resource {
	return []*Resource{
		{
			Service:       "items", // this will be the directory name under resources/services
			TableName:     "items", // will become hackernews_items
			DataStruct:    &hackernews.Item{},
			Description:   "https://github.com/HackerNews/API#items",
			PKColumns:     []string{"id"},
			IsIncremental: true,
		},
	}
}
