package identitystore

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ssoadmin"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func getIamInstances(ctx context.Context, meta schema.ClientMeta) ([]types.InstanceMetadata, error) {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceSsoadmin).Ssoadmin
	config := ssoadmin.ListInstancesInput{}
	paginator := ssoadmin.NewListInstancesPaginator(svc, &config)
	instances := make([]types.InstanceMetadata, 0)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *ssoadmin.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return nil, err
		}
		instances = append(instances, page.Instances...)
	}

	return instances, nil
}
