package recipes

import (
	"github.com/xanzy/go-gitlab"
)

func Projects() []*Resource {
	resources := []*Resource{
		{
			Service:    "projects",
			SubService: "projects",
			PKColumns:  []string{"base_url", "id"},
			Struct:     &gitlab.Project{},
			Relations:  []string{"ProjectsReleases()"},
		},
		{
			Service:    "projects",
			SubService: "projects_releases",
			Struct:     &gitlab.Release{},
		},
	}

	return resources
}
