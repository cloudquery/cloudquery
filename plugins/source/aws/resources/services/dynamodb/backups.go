package dynamodb

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
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
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("BackupDetails.BackupArn"),
				PrimaryKey: true,
			},
		},
	}
}

func listBackups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Dynamodb

	config := dynamodb.ListBackupsInput{}
	// No paginator available
	for {
		output, err := svc.ListBackups(ctx, &config, func(options *dynamodb.Options) {
			options.Region = cl.Region
		})
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
	cl := meta.(*client.Client)
	svc := cl.Services().Dynamodb

	backupSummary := resource.Item.(types.BackupSummary)

	response, err := svc.DescribeBackup(ctx, &dynamodb.DescribeBackupInput{BackupArn: backupSummary.BackupArn}, func(options *dynamodb.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}

	resource.Item = response.BackupDescription
	return nil
}
