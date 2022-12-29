package recipes

import (
	"github.com/fastly/go-fastly/v7/fastly"
)

func AuthResources() []*Resource {
	resources := []*Resource{
		{
			TableName:   "auth_tokens",
			DataStruct:  &fastly.Token{},
			Description: "https://developer.fastly.com/reference/api/auth-tokens/",
			PKColumns:   []string{"id"},
		},
	}
	for _, r := range resources {
		r.Service = "auth"
	}
	return resources
}
