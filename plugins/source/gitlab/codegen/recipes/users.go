package recipes

import (
	"github.com/xanzy/go-gitlab"
)

func Users() []*Resource {
	resources := []*Resource{
		{
			Service:    "users",
			SubService: "users",
			Struct:     &gitlab.User{},
			PKColumns:  []string{"id"},
		},
		{
			Service:    "users",
			SubService: "groups",
			PKColumns:  []string{"id", "name"},
			Struct:     &gitlab.Group{},
			Relations:  []string{"GroupMembers()"},
		},
		{
			Service:    "users",
			SubService: "group_members",
			PKColumns:  []string{"id"},
			Struct:     &gitlab.GroupMember{},
			Relations:  []string{"Users()"},
		},
	}

	return resources
}
