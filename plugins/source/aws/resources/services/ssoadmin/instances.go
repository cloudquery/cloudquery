package ssoadmin

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ssoadmin"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Instances() *schema.Table {
	tableName := "aws_ssoadmin_instances"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_InstanceMetadata.html`,
		Resolver:    fetchSsoadminInstances,
		Transform:   transformers.TransformWithStruct(&types.InstanceMetadata{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "identitystore"),

		Relations: []*schema.Table{
			permissionSets(),
		},
	}
}

func fetchSsoadminInstances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	svc := meta.(*client.Client).Services().Ssoadmin
	config := ssoadmin.ListInstancesInput{}
	response, err := svc.ListInstances(ctx, &config)
	if err != nil {
		return err
	}
	// TODO: replace with paginator
	for _, i := range response.Instances {
		res <- i
	}
	return nil
}
