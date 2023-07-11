package ssoadmin

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ssoadmin"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Instances() *schema.Table {
	tableName := "aws_ssoadmin_instances"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_InstanceMetadata.html`,
		Resolver:    fetchSsoadminInstances,
		Transform:   transformers.TransformWithStruct(&types.InstanceMetadata{}, transformers.WithPrimaryKeys("InstanceArn")),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "identitystore"),

		Relations: []*schema.Table{
			permissionSets(),
		},
	}
}

func fetchSsoadminInstances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Ssoadmin
	config := ssoadmin.ListInstancesInput{}
	paginator := ssoadmin.NewListInstancesPaginator(svc, &config)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx, func(o *ssoadmin.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- output.Instances
	}
	return nil
}
