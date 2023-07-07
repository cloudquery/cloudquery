package devops

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devops/armdevops"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func PipelineTemplateDefinitions() *schema.Table {
	return &schema.Table{
		Name:                 "azure_devops_pipeline_template_definitions",
		Resolver:             fetchPipelineTemplateDefinitions,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devops/armdevops@v0.5.0#PipelineTemplateDefinition",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_devops_pipeline_template_definitions", client.Namespacemicrosoft_devops),
		Transform:            transformers.TransformWithStruct(&armdevops.PipelineTemplateDefinition{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
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
