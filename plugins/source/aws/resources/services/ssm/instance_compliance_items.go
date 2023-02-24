package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func InstanceComplianceItems() *schema.Table {
	return &schema.Table{
		Name:        "aws_ssm_instance_compliance_items",
		Description: `https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_ComplianceItem.html`,
		Resolver:    fetchSsmInstanceComplianceItems,
		Transform:   transformers.TransformWithStruct(&types.ComplianceItem{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("ssm"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Id"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "instance_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchSsmInstanceComplianceItems(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	instance := parent.Item.(types.InstanceInformation)
	cl := meta.(*client.Client)
	svc := cl.Services().Ssm

	input := ssm.ListComplianceItemsInput{
		ResourceIds: []string{*instance.InstanceId},
	}
	for {
		output, err := svc.ListComplianceItems(ctx, &input)
		if err != nil {
			return err
		}
		res <- output.ComplianceItems
		if aws.ToString(output.NextToken) == "" {
			break
		}
		input.NextToken = output.NextToken
	}
	return nil
}
