// Auto generated code - DO NOT EDIT.

package security

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/pkg/errors"
)

func JitNetworkAccessPolicies() *schema.Table {
	return &schema.Table{
		Name:      "azure_security_jit_network_access_policies",
		Resolver:  fetchSecurityJitNetworkAccessPolicies,
		Multiplex: client.SubscriptionMultiplex,
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
				Name:     "kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Kind"),
			},
			{
				Name:     "location",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Location"),
			},
			{
				Name:     "virtual_machines",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("VirtualMachines"),
			},
			{
				Name:     "requests",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Requests"),
			},
			{
				Name:     "provisioning_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ProvisioningState"),
			},
		},
	}
}

func fetchSecurityJitNetworkAccessPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Security.JitNetworkAccessPolicies

	response, err := svc.List(ctx)

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
