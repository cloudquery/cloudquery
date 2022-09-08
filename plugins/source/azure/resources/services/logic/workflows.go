// Auto generated code - DO NOT EDIT.

package logic

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/pkg/errors"
)

func Workflows() *schema.Table {
	return &schema.Table{
		Name:      "azure_logic_workflows",
		Resolver:  fetchLogicWorkflows,
		Multiplex: client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "provisioning_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ProvisioningState"),
			},
			{
				Name:     "created_time",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CreatedTime"),
			},
			{
				Name:     "changed_time",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ChangedTime"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:     "version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Version"),
			},
			{
				Name:     "access_endpoint",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccessEndpoint"),
			},
			{
				Name:     "endpoints_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("EndpointsConfiguration"),
			},
			{
				Name:     "access_control",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AccessControl"),
			},
			{
				Name:     "sku",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Sku"),
			},
			{
				Name:     "integration_account",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("IntegrationAccount"),
			},
			{
				Name:     "integration_service_environment",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("IntegrationServiceEnvironment"),
			},
			{
				Name:     "parameters",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Parameters"),
			},
			{
				Name:     "identity",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Identity"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
			{
				Name:     "location",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Location"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
		},

		Relations: []*schema.Table{
			diagnosticSettings(),
		},
	}
}

func fetchLogicWorkflows(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Logic.Workflows

	var top int32 = 100
	response, err := svc.ListBySubscription(ctx, &top, "")

	if err != nil {
		return errors.WithStack(err)
	}

	for response.NotDone() {
		res <- response.Values()
		if err := response.NextWithContext(ctx); err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}
