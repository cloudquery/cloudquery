package identitystore

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ssoadmin"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func getIamInstance(ctx context.Context, meta schema.ClientMeta) (types.InstanceMetadata, error) {
	cl := meta.(*client.Client)
	svc := cl.Services().Ssoadmin
	config := ssoadmin.ListInstancesInput{}
	response, err := svc.ListInstances(ctx, &config, func(options *ssoadmin.Options) {
		options.Region = cl.Region
	})
	if err == nil {
		for _, i := range response.Instances {
			return i, err
		}
	}
	return types.InstanceMetadata{}, err
}
