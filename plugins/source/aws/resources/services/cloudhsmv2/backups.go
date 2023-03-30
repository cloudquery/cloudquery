package cloudhsmv2

import (
	"github.com/aws/aws-sdk-go-v2/service/cloudhsmv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Backups() *schema.Table {
	tableName := "aws_cloudhsmv2_backups"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/cloudhsm/latest/APIReference/API_Backup.html`,
		Resolver:    fetchCloudhsmv2Backups,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "cloudhsmv2"),
		Transform:   client.TransformWithStruct(&types.Backup{}, transformers.WithSkipFields("TagList")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
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
