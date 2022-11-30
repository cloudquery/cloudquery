package recipes

import (
	"github.com/xanzy/go-gitlab"
)

func Releases() []*Resource {
	resources := []*Resource{
		{
			Service:    "releases",
			SubService: "releases",
			PKColumns:  []string{"id"},
			Struct:     &gitlab.Release{},
		},
	}

	return resources
}
