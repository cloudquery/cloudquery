package fakes

import (
	"github.com/cloudquery/plugin-sdk/v2/faker"
	"github.com/okta/okta-sdk-golang/v3/okta"
)

func Group() okta.Group {
	var g okta.Group
	if err := faker.FakeObject(&g); err != nil {
		panic(err)
	}
	g.Type = &okta.AllowedGroupTypeEnumValues[0]
	g.Links = &okta.GroupLinks{
		Self: &okta.HrefObject{Href: "#"},
	}

	return g
}
