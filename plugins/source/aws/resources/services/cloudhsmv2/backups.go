package cloudhsmv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/cloudhsmv2"
	"github.com/aws/aws-sdk-go-v2/service/cloudhsmv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func Backups() *schema.Table {
	tableName := "aws_cloudhsmv2_backups"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/cloudhsm/latest/APIReference/API_Backup.html`,
		Resolver:    fetchCloudhsmv2Backups,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "cloudhsmv2"),
		Transform:   transformers.TransformWithStruct(&types.Backup{}, transformers.WithSkipFields("TagList")),
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

func fetchCloudhsmv2Backups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Cloudhsmv2
	var input cloudhsmv2.DescribeBackupsInput
	paginator := cloudhsmv2.NewDescribeBackupsPaginator(svc, &input)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- output.Backups
	}
	return nil
}

func resolveBackupArn(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	item := resource.Item.(types.Backup)
	a := arn.ARN{
		Partition: cl.Partition,
		Service:   "hsm",
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  "backup/" + aws.ToString(item.BackupId),
	}
	return resource.Set(c.Name, a.String())
}
