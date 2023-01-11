package resourcemanager

import (
	pb "cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func Folders() *schema.Table {
	return &schema.Table{
		Name:        "gcp_resourcemanager_folders",
		Description: `https://cloud.google.com/resource-manager/reference/rest/v3/folders#Folder`,
		Resolver:    fetchFolders,
		Multiplex:   client.OrgMultiplex,
		Transform:   transformers.TransformWithStruct(&pb.Folder{}, client.Options()...),
		Columns: []schema.Column{
			{
				Name:     "organization_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveOrganization,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
