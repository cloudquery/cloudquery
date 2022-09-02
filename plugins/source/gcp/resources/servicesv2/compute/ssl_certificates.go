// Code generated by codegen; DO NOT EDIT.

package compute

import (
	"context"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"

	"google.golang.org/api/compute/v1"
)

func SslCertificates() *schema.Table {
	return &schema.Table{
		Name:      "gcp_compute_ssl_certificates",
		Resolver:  fetchSslCertificates,
		Multiplex: client.ProjectMultiplex,
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name:     "certificate",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Certificate"),
			},
			{
				Name:     "creation_timestamp",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CreationTimestamp"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "expire_time",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ExpireTime"),
			},
			{
				Name:     "id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Id"),
			},
			{
				Name:     "kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Kind"),
			},
			{
				Name:     "managed",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Managed"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "private_key",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PrivateKey"),
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Region"),
			},
			{
				Name: "self_link",
				Type: schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "self_managed",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SelfManaged"),
			},
			{
				Name:     "subject_alternative_names",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("SubjectAlternativeNames"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
		},
	}
}

func fetchSslCertificates(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		output, err := c.Services.Compute.SslCertificates.AggregatedList(c.ProjectId).PageToken(nextPageToken).Do()
		if err != nil {
			return errors.WithStack(err)
		}

		var allItems []*compute.SslCertificate
		for _, items := range output.Items {
			allItems = append(allItems, items.SslCertificates...)
		}
		res <- allItems

		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
