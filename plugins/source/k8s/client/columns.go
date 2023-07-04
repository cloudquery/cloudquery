package client

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

var ContextColumn = schema.Column{Name: "context", Type: arrow.BinaryTypes.String, Resolver: ResolveContext}
