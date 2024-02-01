package ssm

import (
	"context"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func instancePatches() *schema.Table {
	tableName := "aws_ssm_instance_patches"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_PatchComplianceData.html`,
		Resolver:    fetchSsmInstancePatches,
		Transform:   transformers.TransformWithStruct(&types.PatchComplianceData{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "instance_arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("arn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:                "kb_id",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("KBId"),
				PrimaryKeyComponent: true,
			},
		},
	}
}

func fetchSsmInstancePatches(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceSsm).Ssm
	item := parent.Item.(types.InstanceInformation)

	paginator := ssm.NewDescribeInstancePatchesPaginator(svc, &ssm.DescribeInstancePatchesInput{
		InstanceId: item.InstanceId,
	})
	for paginator.HasMorePages() {
		v, err := paginator.NextPage(ctx, func(o *ssm.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- v.Patches
	}
	return nil
}
