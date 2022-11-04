package identitystore

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ssoadmin"
	types "github.com/aws/aws-sdk-go-v2/service/ssoadmin/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func getIamInstance(ctx context.Context, meta schema.ClientMeta) (types.InstanceMetadata, error) {
	svc := meta.(*client.Client).Services().Ssoadmin
	config := ssoadmin.ListInstancesInput{}
	response, err := svc.ListInstances(ctx, &config)
	if err == nil {
		for _, i := range response.Instances {
			return i, err
		}
	}
	return types.InstanceMetadata{}, err
}
