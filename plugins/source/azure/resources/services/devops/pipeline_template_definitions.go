package devops

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devops/armdevops"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func PipelineTemplateDefinitions() *schema.Table {
	return &schema.Table{
		Name:        "azure_devops_pipeline_template_definitions",
		Resolver:    fetchPipelineTemplateDefinitions,
		Description: "https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devops/armdevops@v0.5.0#PipelineTemplateDefinition",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_devops_pipeline_template_definitions", client.Namespacemicrosoft_devops),
		Transform:   transformers.TransformWithStruct(&armdevops.PipelineTemplateDefinition{}),
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchPipelineTemplateDefinitions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armdevops.NewPipelineTemplateDefinitionsClient(cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListPager(nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
