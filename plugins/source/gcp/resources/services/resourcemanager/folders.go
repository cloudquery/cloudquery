package resourcemanager

import (
	pb "cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	"github.com/apache/arrow/go/v14/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Folders() *schema.Table {
	return &schema.Table{
		Name:        "gcp_resourcemanager_folders",
		Description: `https://cloud.google.com/resource-manager/reference/rest/v3/folders#Folder`,
		Resolver:    fetchFolders,
		Multiplex:   client.OrgMultiplex,
		Transform:   client.TransformWithStruct(&pb.Folder{}, transformers.WithPrimaryKeys("Name")),
		Columns: []schema.Column{
			{
				Name:       "organization_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveOrganization,
				PrimaryKey: true,
			},
		},
		Relations: schema.Tables{
			subFolders(),
		},
	}
}

func subFolders() *schema.Table {
	return &schema.Table{
		Name: "gcp_resourcemanager_subfolders",
		Description: `https://cloud.google.com/resource-manager/reference/rest/v3/folders#Folder

Lower level folders (where the parent is another folder) are stored in this table.`,
		Resolver:  fetchSubFolders,
		Transform: client.TransformWithStruct(&pb.Folder{}, transformers.WithPrimaryKeys("Name")),
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
