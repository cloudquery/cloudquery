package backup

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/backup"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildBackupVaultsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockBackupClient(ctrl)

	var vault types.BackupVaultListMember
	require.NoError(t, faker.FakeObject(&vault))

	m.EXPECT().ListBackupVaults(
		gomock.Any(),
		&backup.ListBackupVaultsInput{MaxResults: aws.Int32(1000)},
		gomock.Any(),
	).Return(
		&backup.ListBackupVaultsOutput{BackupVaultList: []types.BackupVaultListMember{vault}},
		nil,
	)

	// list tags for backup vault
	m.EXPECT().ListTags(
		gomock.Any(),
		&backup.ListTagsInput{ResourceArn: vault.BackupVaultArn},
		gomock.Any(),
	).Return(
		&backup.ListTagsOutput{
			Tags: map[string]string{"tag1": "value1"},
		},
		nil,
	)

	m.EXPECT().GetBackupVaultAccessPolicy(
		gomock.Any(),
		&backup.GetBackupVaultAccessPolicyInput{BackupVaultName: vault.BackupVaultName},
		gomock.Any(),
	).Return(
		&backup.GetBackupVaultAccessPolicyOutput{
			Policy: aws.String(`{"key":"value"}`),
		},
		nil,
	)

	m.EXPECT().GetBackupVaultNotifications(
		gomock.Any(),
		&backup.GetBackupVaultNotificationsInput{BackupVaultName: vault.BackupVaultName},
		gomock.Any(),
	).Return(
		&backup.GetBackupVaultNotificationsOutput{
			BackupVaultEvents: []types.BackupVaultEvent{types.BackupVaultEventBackupJobFailed},
			SNSTopicArn:       aws.String("not really an ARN"),
		},
		nil,
	)

	var rp types.RecoveryPointByBackupVault
	require.NoError(t, faker.FakeObject(&rp))

	rp.ResourceArn = aws.String("arn:aws:s3:eu-central-1:testAccount:resource/id")

	m.EXPECT().ListRecoveryPointsByBackupVault(
		gomock.Any(),
		&backup.ListRecoveryPointsByBackupVaultInput{BackupVaultName: vault.BackupVaultName, MaxResults: aws.Int32(100)},
		gomock.Any(),
	).Return(
		&backup.ListRecoveryPointsByBackupVaultOutput{RecoveryPoints: []types.RecoveryPointByBackupVault{rp}},
		nil,
	)

	// list tags for recovery point
	m.EXPECT().ListTags(
		gomock.Any(),
		&backup.ListTagsInput{ResourceArn: rp.RecoveryPointArn},
		gomock.Any(),
	).Return(
		&backup.ListTagsOutput{
			Tags: map[string]string{"tag1": "value1"},
		},
		nil,
	)

	return client.Services{
		Backup: m,
	}
}

func TestVaults(t *testing.T) {
	client.AwsMockTestHelper(t, Vaults(), buildBackupVaultsMock, client.TestOptions{})
}
