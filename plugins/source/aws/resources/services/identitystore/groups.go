package identitystore

import (
	"github.com/aws/aws-sdk-go-v2/service/identitystore/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Groups() *schema.Table {
	tableName := "aws_identitystore_groups"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/singlesignon/latest/IdentityStoreAPIReference/API_Group.html`,
		Resolver:    fetchIdentitystoreGroups,
		Transform:   client.TransformWithStruct(&types.Group{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "identitystore"),

		Relations: []*schema.Table{
			GroupMemberships(),
		},
	}
}
