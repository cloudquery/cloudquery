package cdn

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func rule_sets() *schema.Table {
	return &schema.Table{
		Name:        "azure_cdn_rule_sets",
		Resolver:    fetchRuleSets,
		Description: "https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn@v1.0.0#RuleSet",
		Transform:   transformers.TransformWithStruct(&armcdn.RuleSet{}, transformers.WithPrimaryKeys("ID")),
		Columns:     schema.ColumnList{},
	}
}
