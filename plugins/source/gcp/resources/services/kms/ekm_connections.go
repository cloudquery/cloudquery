package kms

import (
	pb "cloud.google.com/go/kms/apiv1/kmspb"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func EkmConnections() *schema.Table {
	return &schema.Table{
		Name:        "gcp_kms_ekm_connections",
		Description: `https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.ekmConnections#EkmConnection`,
		Resolver:    fetchEkmConnections,
		Multiplex:   client.ProjectMultiplexEnabledServices("cloudkms.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.EkmConnection{}, transformers.WithPrimaryKeys("Name")),
		Columns: []schema.Column{
			{
				Name:       "project_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveProject,
				PrimaryKey: true,
			},
		},
	}
}
