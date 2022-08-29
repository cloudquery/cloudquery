// Auto generated code - DO NOT EDIT.

package {{.AzureService | ToLower}}

import (
	"context"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/pkg/errors"
    {{template "imports.go.tpl" .}}
)

func {{.AzureSubService}}() *schema.Table {
    return &schema.Table{{template "table.go.tpl" .Table}}
}