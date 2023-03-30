package identitystore

import (
	"github.com/aws/aws-sdk-go-v2/service/identitystore/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Users() *schema.Table {
	tableName := "aws_identitystore_users"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/singlesignon/latest/IdentityStoreAPIReference/API_User.html`,
		Resolver:    fetchIdentitystoreUsers,
		Transform:   transformers.TransformWithStruct(&types.User{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "identitystore"),
	}
}
