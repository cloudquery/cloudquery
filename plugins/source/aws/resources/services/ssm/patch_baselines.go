package ssm

import (
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func PatchBaselines() *schema.Table {
	tableName := "aws_ssm_patch_baselines"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_PatchBaselineIdentity.html`,
		Resolver:    fetchSsmPatchBaselines,
		Transform:   client.TransformWithStruct(&types.PatchBaselineIdentity{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "ssm"),
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
