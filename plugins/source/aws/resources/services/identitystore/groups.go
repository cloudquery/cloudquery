package identitystore

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/identitystore"
	"github.com/aws/aws-sdk-go-v2/service/identitystore/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Groups() *schema.Table {
	tableName := "aws_identitystore_groups"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/singlesignon/latest/IdentityStoreAPIReference/API_Group.html`,
		Resolver:    fetchIdentitystoreGroups,
		Transform:   transformers.TransformWithStruct(&types.Group{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "identitystore"),

		Relations: []*schema.Table{
			groupMemberships(),
		},
	}
}

func fetchIdentitystoreGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	instance, err := getIamInstance(ctx, meta)
	if err != nil {
		return err
	}
	cl := meta.(*client.Client)
	svc := cl.Services().Identitystore
	config := identitystore.ListGroupsInput{
		IdentityStoreId: instance.IdentityStoreId,
	}
	paginator := identitystore.NewListGroupsPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *identitystore.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Groups
	}
	return nil
}
