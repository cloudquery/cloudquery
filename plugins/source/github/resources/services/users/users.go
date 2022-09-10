package users

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

// Users is a utility table to be used as a relation
func Users(name, description string) *schema.Table {
	return &schema.Table{
		Name:                 name,
		Description:          description,
		Columns:              nil,
		Relations:            nil,
		Resolver:             nil,
		Multiplex:            nil,
		PostResourceResolver: nil,
		PreResourceResolver:  nil,
		IgnoreInTests:        false,
		Parent:               nil,
		Serial:               "",
	}
}
