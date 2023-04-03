package backup

import (
	"context"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/backup"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func vaultRecoveryPoints() *schema.Table {
	tableName := "aws_backup_vault_recovery_points"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/aws-backup/latest/devguide/API_RecoveryPointByBackupVault.html`,
		Resolver:    fetchBackupVaultRecoveryPoints,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "backup"),
		Transform:   transformers.TransformWithStruct(&types.RecoveryPointByBackupVault{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "vault_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RecoveryPointArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveRecoveryPointTags,
			},
		},
	}
}

func fetchBackupVaultRecoveryPoints(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Backup
	vault := parent.Item.(types.BackupVaultListMember)
	params := backup.ListRecoveryPointsByBackupVaultInput{BackupVaultName: vault.BackupVaultName, MaxResults: aws.Int32(100)}
	for {
		result, err := svc.ListRecoveryPointsByBackupVault(ctx, &params)
		if err != nil {
			return err
		}
		res <- result.RecoveryPoints
		if aws.ToString(result.NextToken) == "" {
			break
		}
		params.NextToken = result.NextToken
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

	// decide if the backed up resource supports tags
	switch client.AWSService(resourceARN.Service) {
	case client.S3Service, client.EFSService, client.DynamoDBService:
		// these services are ok
	case client.EC2Service:
		if !strings.HasPrefix(resourceARN.Resource, "instance/") {
			return nil
		}
	default:
		// full backup management not supported, so no tags
		return nil
	}

	cl := meta.(*client.Client)
	svc := cl.Services().Backup
	params := backup.ListTagsInput{ResourceArn: rp.RecoveryPointArn}
	tags := make(map[string]string)
	for {
		result, err := svc.ListTags(ctx, &params, func(o *backup.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			if client.IsAWSError(err, "ERROR_2603") {
				// ignoring "ERROR_2603: Cannot find recovery point."
				return nil
			}
			if resourceARN.Service == string(client.DynamoDBService) && client.IsAWSError(err, "ERROR_3930") {
				// advanced backup features are not enabled for dynamodb
				return nil
			}
			return err
		}

		if result == nil {
			break
		}

		for k, v := range result.Tags {
			tags[k] = v
		}

		if aws.ToString(result.NextToken) == "" {
			break
		}
		params.NextToken = result.NextToken
	}
	return resource.Set(c.Name, tags)
}
