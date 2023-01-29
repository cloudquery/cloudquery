package cdn

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func security_policies() *schema.Table {
	return &schema.Table{
		Name:        "azure_cdn_security_policies",
		Resolver:    fetchSecurityPolicies,
		Description: "https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn@v1.0.0#SecurityPolicy",
		Transform:   transformers.TransformWithStruct(&armcdn.SecurityPolicy{}, transformers.WithPrimaryKeys("ID")),
		Columns:     schema.ColumnList{},
	}
}
