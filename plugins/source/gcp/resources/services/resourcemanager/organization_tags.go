package resourcemanager

import (
	pb "cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	"github.com/apache/arrow/go/v14/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func OrganizationTagKeys() *schema.Table {
	return &schema.Table{
		Name:        "gcp_resourcemanager_organization_tag_keys",
		Description: `https://cloud.google.com/resource-manager/reference/rest/v3/tagKeys/list`,
		Resolver:    fetchOrganizationTagKeys,
		Multiplex:   client.OrgMultiplex,
		Transform:   client.TransformWithStruct(&pb.TagKey{}, transformers.WithPrimaryKeys("Name")),
		Columns: []schema.Column{
			{
				Name:       "organization_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveOrganization,
				PrimaryKey: true,
			},
		},
		Relations: []*schema.Table{
			organizationTagValues(),
		},
	}
}

func organizationTagValues() *schema.Table {
	return &schema.Table{
		Name:        "gcp_resourcemanager_organization_tag_values",
		Description: `https://cloud.google.com/resource-manager/reference/rest/v3/tagValues/list`,
		Resolver:    fetchOgranizationTagValues,
		Multiplex:   client.OrgMultiplex,
		Transform:   client.TransformWithStruct(&pb.TagValue{}, transformers.WithPrimaryKeys("Name")),
		Columns: []schema.Column{
			{
				Name:       "organization_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveOrganization,
				PrimaryKey: true,
			},
		},
	}
}
