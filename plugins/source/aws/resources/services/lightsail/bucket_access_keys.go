package lightsail

import (
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func BucketAccessKeys() *schema.Table {
	tableName := "aws_lightsail_bucket_access_keys"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_AccessKey.html`,
		Resolver:    fetchLightsailBucketAccessKeys,
		Transform:   transformers.TransformWithStruct(&types.AccessKey{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "lightsail"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "bucket_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},
	}
}
