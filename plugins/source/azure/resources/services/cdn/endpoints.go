// Auto generated code - DO NOT EDIT.

package cdn

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/cdn/mgmt/cdn"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func endpoints() *schema.Table {
	return &schema.Table{
		Name:        "azure_cdn_endpoints",
		Description: `https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/cdn/mgmt/2020-09-01/cdn#Endpoint`,
		Resolver:    fetchCDNEndpoints,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "cdn_profile_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
			},
			{
				Name:     "host_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HostName"),
			},
			{
				Name:     "origins",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Origins"),
			},
			{
				Name:     "origin_groups",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("OriginGroups"),
			},
			{
				Name:     "resource_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceState"),
			},
			{
				Name:     "provisioning_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ProvisioningState"),
			},
			{
				Name:     "origin_path",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OriginPath"),
			},
			{
				Name:     "content_types_to_compress",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("ContentTypesToCompress"),
			},
			{
				Name:     "origin_host_header",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OriginHostHeader"),
			},
			{
				Name:     "is_compression_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsCompressionEnabled"),
			},
			{
				Name:     "is_http_allowed",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsHTTPAllowed"),
			},
			{
				Name:     "is_https_allowed",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsHTTPSAllowed"),
			},
			{
				Name:     "query_string_caching_behavior",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("QueryStringCachingBehavior"),
			},
			{
				Name:     "optimization_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OptimizationType"),
			},
			{
				Name:     "probe_path",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ProbePath"),
			},
			{
				Name:     "geo_filters",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("GeoFilters"),
			},
			{
				Name:     "default_origin_group",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DefaultOriginGroup"),
			},
			{
				Name:     "url_signing_keys",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("URLSigningKeys"),
			},
			{
				Name:     "delivery_policy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DeliveryPolicy"),
			},
			{
				Name:     "web_application_firewall_policy_link",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("WebApplicationFirewallPolicyLink"),
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

		Relations: []*schema.Table{
			customDomains(),
			routes(),
		},
	}
}

func fetchCDNEndpoints(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().CDN.Endpoints

	profile := parent.Item.(cdn.Profile)
	resource, err := client.ParseResourceID(*profile.ID)
	if err != nil {
		return err
	}
	response, err := svc.ListByProfile(ctx, resource.ResourceGroup, *profile.Name)

	if err != nil {
		return err
	}

	for response.NotDone() {
		res <- response.Values()
		if err := response.NextWithContext(ctx); err != nil {
			return err
		}
	}

	return nil
}
