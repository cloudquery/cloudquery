package elastictranscoder

import (
	"github.com/aws/aws-sdk-go-v2/service/elastictranscoder/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Presets() *schema.Table {
	return &schema.Table{
		Name:        "aws_elastictranscoder_presets",
		Description: `https://docs.aws.amazon.com/elastictranscoder/latest/developerguide/list-presets.html`,
		Resolver:    fetchElastictranscoderPresets,
		Multiplex:   client.ServiceAccountRegionMultiplexer("elastictranscoder"),
		Transform:   transformers.TransformWithStruct(&types.Preset{}),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
