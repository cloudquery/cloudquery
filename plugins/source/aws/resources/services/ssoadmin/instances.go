package ssoadmin

import (
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Instances() *schema.Table {
	return &schema.Table{
		Name:        "aws_ssoadmin_instances",
		Description: `https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_InstanceMetadata.html`,
		Resolver:    fetchSsoadminInstances,
		Transform:   transformers.TransformWithStruct(&types.InstanceMetadata{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("identitystore"),

		Relations: []*schema.Table{
			PermissionSets(),
		},
	}
}
