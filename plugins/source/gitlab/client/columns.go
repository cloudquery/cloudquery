package client

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

var BaseURLColumn = schema.Column{
	Name:        "base_url",
	Type:        arrow.BinaryTypes.String,
	Description: "GitLab instance base URL",
	NotNull:     true,
	PrimaryKey:  true,
	Resolver:    ResolveURL,
}
