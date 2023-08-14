package cdn

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func rule_sets() *schema.Table {
	return &schema.Table{
		Name:                 "azure_cdn_rule_sets",
		Resolver:             fetchRuleSets,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn@v1.0.0#RuleSet",
		Transform:            transformers.TransformWithStruct(&armcdn.RuleSet{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{},
	}
}
