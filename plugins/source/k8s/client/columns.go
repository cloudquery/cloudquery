package client

import (
	"github.com/cloudquery/plugin-sdk/v2/schema"
)

var ContextColumn = schema.Column{Name: "context", Type: schema.TypeString, Resolver: ResolveContext}
