package identitystore

import (
	"context"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
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
		Resolver:    fetchGroups,
		Transform:   transformers.TransformWithStruct(&types.Group{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "identitystore"),
		Columns: schema.ColumnList{
			client.RequestAccountIDColumn(true),
			client.RequestRegionColumn(true),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            resolveGroupARN,
				PrimaryKeyComponent: true,
			},
		},
		Relations: []*schema.Table{
			groupMemberships(),
		},
	}
}

func fetchGroups(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	instances, err := getIamInstances(ctx, meta)
	if err != nil {
		return err
	}
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceIdentitystore).Identitystore
	for _, instance := range instances {
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
	}
	return nil
}

func resolveGroupARN(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	return resource.Set(c.Name, groupARN(cl, aws.ToString(resource.Item.(types.Group).GroupId)))
}

func groupARN(cl *client.Client, groupID string) string {
	// https: //docs.aws.amazon.com/service-authorization/latest/reference/list_awsidentitystore.html#awsidentitystore-resources-for-iam-policies
	return arn.ARN{
		Partition: cl.Partition,
		Service:   string(client.IdentitystoreService),
		Resource:  "group/" + groupID,
	}.String()
}
