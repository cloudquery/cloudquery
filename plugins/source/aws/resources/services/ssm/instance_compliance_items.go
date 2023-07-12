package ssm

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func instanceComplianceItems() *schema.Table {
	tableName := "aws_ssm_instance_compliance_items"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_ComplianceItem.html`,
		Resolver:    fetchSsmInstanceComplianceItems,
		Transform:   transformers.TransformWithStruct(&types.ComplianceItem{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "ssm"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("Id"),
				PrimaryKey: true,
			},
			{
				Name:       "instance_arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("arn"),
				PrimaryKey: true,
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
	paginator := ssm.NewListComplianceItemsPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(o *ssm.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.ComplianceItems
	}
	return nil
}
