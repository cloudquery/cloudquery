package client

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

var ContextColumn = schema.Column{Name: "context", Type: schema.TypeString, Resolver: ResolveContext}
