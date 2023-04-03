package identitystore

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/identitystore"
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

func fetchIdentitystoreUsers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	instance, err := getIamInstance(ctx, meta)
	if err != nil {
		return err
	}
	svc := meta.(*client.Client).Services().Identitystore
	config := identitystore.ListUsersInput{}
	config.IdentityStoreId = instance.IdentityStoreId
	for {
		response, err := svc.ListUsers(ctx, &config)
		if err != nil {
			return err
		}
		res <- response.Users

		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}
