package client

import (
	"github.com/cloudquery/plugin-sdk/v2/schema"
)

var (
	BaseURLColumn = schema.Column{
		Name:            "base_url",
		Type:            schema.TypeString,
		Description:     "GitLab instance base URL",
		Resolver:        ResolveURL,
		CreationOptions: schema.ColumnCreationOptions{NotNull: true, PrimaryKey: true},
	}
)
