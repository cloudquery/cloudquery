package backup

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/backup"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildBackupVaultsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockBackupClient(ctrl)

	var vault types.BackupVaultListMember
	if err := faker.FakeData(&vault); err != nil {
		t.Fatal(err)
	}
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
	if err := faker.FakeData(&rp); err != nil {
		t.Fatal(err)
	}
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
