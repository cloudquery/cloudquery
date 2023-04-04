package elastictranscoder

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elastictranscoder"
	"github.com/aws/aws-sdk-go-v2/service/elastictranscoder/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Presets() *schema.Table {
	tableName := "aws_elastictranscoder_presets"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/elastictranscoder/latest/developerguide/list-presets.html`,
		Resolver:    fetchElastictranscoderPresets,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "elastictranscoder"),
		Transform:   transformers.TransformWithStruct(&types.Preset{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
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

func fetchElastictranscoderPresets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Elastictranscoder

	p := elastictranscoder.NewListPresetsPaginator(svc, nil)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx)
		if err != nil {
			return err
		}

		res <- response.Presets
	}

	return nil
}
