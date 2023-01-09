package identitystore

import (
	"github.com/aws/aws-sdk-go-v2/service/identitystore/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Users() *schema.Table {
	return &schema.Table{
		Name:      "aws_identitystore_users",
		Resolver:  fetchIdentitystoreUsers,
		Transform: transformers.TransformWithStruct(&types.User{}),
		Multiplex: client.ServiceAccountRegionMultiplexer("identitystore"),
	}
}
