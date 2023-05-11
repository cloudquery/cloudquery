package client

import (
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

var ContextColumn = schema.Column{Name: "context", Type: schema.TypeString, Resolver: ResolveContext}
