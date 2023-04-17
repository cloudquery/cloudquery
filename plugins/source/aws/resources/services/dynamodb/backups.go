package dynamodb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func Backups() *schema.Table {
	tableName := "aws_dynamodb_backups"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_BackupDescription.html`,
		Resolver:            listBackups,
		PreResourceResolver: getBackup,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "dynamodb"),
		Transform:           transformers.TransformWithStruct(&types.BackupDescription{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("BackupDetails.BackupArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func listBackups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Dynamodb

	config := dynamodb.ListBackupsInput{}
	// No paginator available
	for {
		output, err := svc.ListBackups(ctx, &config)
		if err != nil {
			return err
		}
		res <- output.BackupSummaries

		if aws.ToString(output.LastEvaluatedBackupArn) == "" {
			break
		}
		config.ExclusiveStartBackupArn = output.LastEvaluatedBackupArn
	}

	return nil
}

func getBackup(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	svc := meta.(*client.Client).Services().Dynamodb

	backupSummary := resource.Item.(types.BackupSummary)

	response, err := svc.DescribeBackup(ctx, &dynamodb.DescribeBackupInput{BackupArn: backupSummary.BackupArn})
	if err != nil {
		return err
	}

	resource.Item = response.BackupDescription
	return nil
}
