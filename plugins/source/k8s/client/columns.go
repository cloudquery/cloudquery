package client

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

var ContextColumn = schema.Column{Name: "context", Type: arrow.BinaryTypes.String, Resolver: ResolveContext}
