package compute

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
	"google.golang.org/api/compute/v1"
)

func ComputeSslPolicies() *schema.Table {
	return &schema.Table{
		Name:        "gcp_compute_ssl_policies",
		Description: "Represents an SSL Policy resource",
		Resolver:    fetchComputeSslPolicies,
		Multiplex:   client.ProjectMultiplex,

		Options: schema.TableCreationOptions{PrimaryKeys: []string{"project_id", "id"}},
		Columns: []schema.Column{
			{
				Name:        "project_id",
				Description: "GCP Project Id of the resource",
				Type:        schema.TypeString,
				Resolver:    client.ResolveProject,
			},
			{
				Name:        "creation_timestamp",
				Description: "Creation timestamp in RFC3339 text format",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.DateResolver("CreationTimestamp"),
			},
			{
				Name:          "custom_features",
				Description:   "A list of features enabled when the selected profile is CUSTOM The - method returns the set of features that can be specified in this list This field must be empty if the profile is not CUSTOM",
				Type:          schema.TypeStringArray,
				IgnoreInTests: true,
			},
			{
				Name:        "description",
				Description: "An optional description of this resource",
				Type:        schema.TypeString,
			},
			{
				Name:          "enabled_features",
				Description:   "The list of features enabled in the SSL policy",
				Type:          schema.TypeStringArray,
				IgnoreInTests: true,
			},
			{
				Name:        "fingerprint",
				Description: "Fingerprint of this resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "id",
				Description: "The unique identifier for the resource This identifier is defined by the server",
				Type:        schema.TypeString,
				Resolver:    client.ResolveResourceId,
			},
			{
				Name:        "kind",
				Description: "Type of the resource Always compute#sslPolicy for SSL policies",
				Type:        schema.TypeString,
			},
			{
				Name:        "min_tls_version",
				Description: "The minimum version of SSL protocol that can be used by the clients to establish a connection with the load balancer This can be one of TLS_1_0, TLS_1_1, TLS_1_2",
				Type:        schema.TypeString,
			},
			{
				Name:        "name",
				Description: "Name of the resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "profile",
				Description: "Profile specifies the set of SSL features that can be used by the load balancer when negotiating SSL with clients This can be one of COMPATIBLE, MODERN, RESTRICTED, or CUSTOM If using CUSTOM, the set of SSL features to enable must be specified in the customFeatures field",
				Type:        schema.TypeString,
			},
			{
				Name:        "self_link",
				Description: "Server-defined URL for the resource",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:          "gcp_compute_ssl_policy_warnings",
				Resolver:      fetchComputeSslPolicyWarnings,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "ssl_policy_cq_id",
						Description: "Unique ID of gcp_compute_ssl_policies table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "policy_id",
						Type:     schema.TypeString,
						Resolver: schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "code",
						Description: "A warning code, if applicable For example, Compute Engine returns NO_RESULTS_ON_PAGE if there are no results in the response",
						Type:        schema.TypeString,
					},
					{
						Name:        "data",
						Description: "Metadata about this warning in key: value format",
						Type:        schema.TypeJSON,
						Resolver:    resolveComputeSslPolicyWarningData,
					},
					{
						Name:        "message",
						Description: "A human-readable description of the warning code",
						Type:        schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchComputeSslPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		output, err := c.Services.Compute.SslPolicies.List(c.ProjectId).PageToken(nextPageToken).Do()
		if err != nil {
			return errors.WithStack(err)
		}

		res <- output.Items

		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}

func fetchComputeSslPolicyWarnings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Item.(*compute.SslPolicy)
	res <- p.Warnings
	return nil
}
func resolveComputeSslPolicyWarningData(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(*compute.SslPolicyWarnings)
	data := make(map[string]string)
	for _, v := range p.Data {
		data[v.Key] = v.Value
	}
	return errors.WithStack(resource.Set(c.Name, data))
}
