package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ManagedPrefixLists() *schema.Table {
	return &schema.Table{
		Name: "aws_ec2_managed_prefix_lists",
		Description: `https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ManagedPrefixList.html. 
The 'request_account_id' and 'request_region' columns are added to show the account_id and region of where the request was made from.`,
		Resolver:  fetchEc2ManagedPrefixLists,
		Multiplex: client.ServiceAccountRegionMultiplexer("ec2"),
		Transform: transformers.TransformWithStruct(&types.ManagedPrefixList{}),
		Columns: []schema.Column{
			{
				Name:            "request_account_id",
				Type:            schema.TypeString,
				Resolver:        client.ResolveAWSAccount,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:            "request_region",
				Type:            schema.TypeString,
				Resolver:        client.ResolveAWSRegion,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:            "arn",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("PrefixListArn"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
		},
	}
}
func fetchEc2ManagedPrefixLists(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Ec2
	paginator := ec2.NewDescribeManagedPrefixListsPaginator(svc, &ec2.DescribeManagedPrefixListsInput{})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.PrefixLists
	}
	return nil
}
