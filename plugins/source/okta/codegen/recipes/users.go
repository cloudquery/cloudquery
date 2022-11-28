package recipes

import (
	"github.com/okta/okta-sdk-golang/v2/okta"
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
