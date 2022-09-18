// Auto generated code - DO NOT EDIT.

package cdn

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/cdn/mgmt/cdn"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/pkg/errors"
)

func routes() *schema.Table {
	return &schema.Table{
		Name:     "azure_cdn_routes",
		Resolver: fetchCDNRoutes,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "cq_id_parent",
				Type:     schema.TypeUUID,
				Resolver: schema.ParentIDResolver,
			},
			{
				Name:     "custom_domains",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CustomDomains"),
			},
			{
				Name:     "origin_group",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("OriginGroup"),
			},
			{
				Name:     "origin_path",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OriginPath"),
			},
			{
				Name:     "rule_sets",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("RuleSets"),
			},
			{
				Name:     "supported_protocols",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("SupportedProtocols"),
			},
			{
				Name:     "patterns_to_match",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("PatternsToMatch"),
			},
			{
				Name:     "query_string_caching_behavior",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("QueryStringCachingBehavior"),
			},
			{
				Name:     "forwarding_protocol",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ForwardingProtocol"),
			},
			{
				Name:     "link_to_default_domain",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LinkToDefaultDomain"),
			},
			{
				Name:     "https_redirect",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HTTPSRedirect"),
			},
			{
				Name:     "enabled_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EnabledState"),
			},
			{
				Name:     "provisioning_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ProvisioningState"),
			},
			{
				Name:     "deployment_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DeploymentStatus"),
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
				Name:     "system_data",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SystemData"),
			},
		},
	}
}

func fetchCDNRoutes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().CDN.Routes

	profile := parent.Parent.Item.(cdn.Profile)
	resource, err := client.ParseResourceID(*profile.ID)
	if err != nil {
		return errors.WithStack(err)
	}
	endpoint := parent.Item.(cdn.Endpoint)
	response, err := svc.ListByEndpoint(ctx, resource.ResourceGroup, *profile.Name, *endpoint.Name)

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
