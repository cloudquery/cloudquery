package recipes

import (
	"github.com/xanzy/go-gitlab"
)

func Groups() []*Resource {
	resources := []*Resource{
		{
			Service:    "groups",
			SubService: "groups",
			PKColumns:  []string{"id", "name"},
			Struct:     &gitlab.Group{},
			Relations:  []string{"GroupMembers()"},
		},
		{
			Service:    "groups",
			SubService: "group_members",
			PKColumns:  []string{"id"},
			Struct:     &gitlab.GroupMember{},
		},
	}

	return resources
}
