package backup

import (
	"context"
	"strings"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/backup"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"
)

func vaultRecoveryPoints() *schema.Table {
	tableName := "aws_backup_vault_recovery_points"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/aws-backup/latest/devguide/API_RecoveryPointByBackupVault.html`,
		Resolver:    fetchBackupVaultRecoveryPoints,
		Transform:   transformers.TransformWithStruct(&types.RecoveryPointByBackupVault{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "vault_arn",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("RecoveryPointArn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveRecoveryPointTags,
			},
		},
	}
}

func fetchBackupVaultRecoveryPoints(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceBackup).Backup
	vault := parent.Item.(types.BackupVaultListMember)
	params := backup.ListRecoveryPointsByBackupVaultInput{BackupVaultName: vault.BackupVaultName, MaxResults: aws.Int32(100)}
	paginator := backup.NewListRecoveryPointsByBackupVaultPaginator(svc, &params)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *backup.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.RecoveryPoints
	}
	return nil
}

func resolveRecoveryPointTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	rp := resource.Item.(types.RecoveryPointByBackupVault)
	if rp.ResourceArn == nil || rp.RecoveryPointArn == nil {
		return nil
	}
	resourceARN, err := arn.Parse(*rp.ResourceArn)
	if err != nil {
		return err
	}
	recoveryPointArn, err := arn.Parse(*rp.RecoveryPointArn)
	if err != nil {
		return err
	}

	// decide if the backed up resource supports tags
	switch client.AWSService(resourceARN.Service) {
	case client.S3Service, client.EFSService, client.RDSService:

		// these services are ok
	case client.DynamoDBService:
		// DynamoDB backups in accounts without "Advanced DynamoDB Backups" do not have Full Backup Management and do not support tagging.
		// DynamoDB backups in such accounts are in the "dynamodb" service namespace, instead of "awsbackup".
		// https://docs.aws.amazon.com/aws-backup/latest/devguide/advanced-ddb-backup.html#advanced-ddb-backup-other-benefits
		if recoveryPointArn.Service == "dynamodb" {
			return nil
		}
	case client.EC2Service:
		if !strings.HasPrefix(resourceARN.Resource, "instance/") {
			return nil
		}
	default:
		// full backup management not supported, so no tags
		return nil
	}

	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceBackup).Backup
	params := backup.ListTagsInput{ResourceArn: rp.RecoveryPointArn}
	tags := make(map[string]string)
	paginator := backup.NewListTagsPaginator(svc, &params)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *backup.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			if client.IsAWSError(err, "ERROR_2603") {
				// ignoring "ERROR_2603: Cannot find recovery point."
				return nil
			}
			return err
		}
		for k, v := range page.Tags {
			tags[k] = v
		}
	}
	return resource.Set(c.Name, tags)
}
