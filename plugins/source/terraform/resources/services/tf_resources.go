package services

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/terraform/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func tfResources() *schema.Table {
	return &schema.Table{
		Name:        "tf_resources",
		Description: "Terraform resources",
		Resolver:    resolveTerraformResources,
		Columns: []schema.Column{
			{
				Name:     "data_backend_name",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("backend_name"),
			},
			{
				Name:        "module",
				Description: "Resource module if exists",
				Type:        arrow.BinaryTypes.String,
			},
			{
				Name:        "mode",
				Description: "Resource mode, for example: data, managed, etc",
				Type:        arrow.BinaryTypes.String,
			},
			{
				Name:        "type",
				Description: "Resource type",
				Type:        arrow.BinaryTypes.String,
			},
			{
				Name:        "name",
				Description: "Resource name",
				Type:        arrow.BinaryTypes.String,
			},
			{
				Name:        "provider_path",
				Description: "Resource provider full path, for example: provider[\"registry.terraform.io/hashicorp/aws\"]",
				Type:        arrow.BinaryTypes.String,
				Resolver:    schema.PathResolver("ProviderConfig"),
			},
			{
				Name:        "provider",
				Description: "Resource provider name, for example: aws, gcp, etc",
				Type:        arrow.BinaryTypes.String,
				Resolver:    resolveProviderName,
			},
		},
		Relations: []*schema.Table{tfResourceInstances()},
	}
}

func resolveTerraformResources(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	state := parent.Item.(client.State)
	for _, resource := range state.Resources {
		res <- resource
	}
	return nil
}

func resolveProviderName(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	res := resource.Item.(client.Resource)
	matches := providerNameRegex.FindStringSubmatch(res.ProviderConfig)
	typeIndex := providerNameRegex.SubexpIndex("Type")
	if len(matches) >= 3 {
		return resource.Set(c.Name, matches[typeIndex])
	}
	return nil
}
