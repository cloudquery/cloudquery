package ssm

import (
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func PatchBaselines() *schema.Table {
	return &schema.Table{
		Name:        "aws_ssm_patch_baselines",
		Description: `https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_PatchBaselineIdentity.html`,
		Resolver:    fetchSsmPatchBaselines,
		Transform:   transformers.TransformWithStruct(&types.PatchBaselineIdentity{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("ssm"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:     "baseline_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("BaselineId"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
