package recipes

import (
	"github.com/xanzy/go-gitlab"
)

func Projects() []*Resource {
	resources := []*Resource{
		{
			Service:    "projects",
			SubService: "projects",
			PKColumns:  []string{"id"},
			Struct:     &gitlab.Project{},
		},
	}

	return resources
}
