package ssm

import (
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func InstancePatches() *schema.Table {
	tableName := "aws_ssm_instance_patches"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_PatchComplianceData.html`,
		Resolver:    fetchSsmInstancePatches,
		Transform:   client.TransformWithStruct(&types.PatchComplianceData{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "ssm"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:     "kb_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("KBId"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
