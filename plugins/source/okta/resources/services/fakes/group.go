package fakes

import (
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/okta/okta-sdk-golang/v4/okta"
)

func Group() okta.Group {
	var g okta.Group
	if err := faker.FakeObject(&g); err != nil {
		panic(err)
	}
	g.Links = &okta.GroupLinks{
		Self: &okta.HrefObjectSelfLink{Href: "#"},
	}
	g.AdditionalProperties = map[string]any{"key": "value"}
	g.Embedded = map[string]map[string]any{"top-key": {"key": "value"}}

	return g
}
