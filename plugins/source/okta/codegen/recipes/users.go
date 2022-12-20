package recipes

import (
	"github.com/okta/okta-sdk-golang/v3/okta"
)

func Users() []*Resource {
	return []*Resource{
		{
			DataStruct: &okta.User{},
			PKColumns:  []string{"id"},
			Service:    "users",
		},
	}
}
