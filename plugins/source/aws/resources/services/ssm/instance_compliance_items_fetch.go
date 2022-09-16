package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func resolveInstanceComplianceItemInstanceARN(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	instance := resource.Parent.Item.(types.InstanceInformation)
	cl := meta.(*client.Client)
	return resource.Set(c.Name, cl.ARN("ssm", "managed-instance", *instance.InstanceId))
}
