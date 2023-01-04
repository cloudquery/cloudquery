package cloudhsmv2

import (
	"github.com/aws/aws-sdk-go-v2/service/cloudhsmv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Backups() *schema.Table {
	return &schema.Table{
		Name:        "aws_cloudhsmv2_backups",
		Description: `https://docs.aws.amazon.com/cloudhsm/latest/APIReference/API_Backup.html`,
		Resolver:    fetchCloudhsmv2Backups,
		Multiplex:   client.ServiceAccountRegionMultiplexer("cloudhsmv2"),
		Transform: transformers.TransformWithStruct(&types.Backup{}),
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
				Resolver: resolveBackupArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTagField("TagList"),
			},
		},
	}
}
