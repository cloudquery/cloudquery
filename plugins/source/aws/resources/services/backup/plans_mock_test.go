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

func buildBackupPlansMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockBackupClient(ctrl)

	var plan backup.GetBackupPlanOutput
	require.NoError(t, faker.FakeObject(&plan))

	m.EXPECT().ListBackupPlans(
		gomock.Any(),
		&backup.ListBackupPlansInput{MaxResults: aws.Int32(1000)},
		gomock.Any(),
	).Return(
		&backup.ListBackupPlansOutput{BackupPlansList: []types.BackupPlansListMember{
			{
				BackupPlanId: plan.BackupPlanId,
				VersionId:    plan.VersionId,
			},
		}},
		nil,
	)

	m.EXPECT().GetBackupPlan(
		gomock.Any(),
		&backup.GetBackupPlanInput{BackupPlanId: plan.BackupPlanId, VersionId: plan.VersionId},
		gomock.Any(),
	).Return(
		&plan,
		nil,
	)

	m.EXPECT().ListTags(
		gomock.Any(),
		&backup.ListTagsInput{ResourceArn: plan.BackupPlanArn},
		gomock.Any(),
	).Return(
		&backup.ListTagsOutput{
			Tags: map[string]string{"plan1": "value1"},
		},
		nil,
	)

	var selection backup.GetBackupSelectionOutput
	require.NoError(t, faker.FakeObject(&selection))

	selection.BackupPlanId = plan.BackupPlanId
	m.EXPECT().ListBackupSelections(
		gomock.Any(),
		&backup.ListBackupSelectionsInput{
			BackupPlanId: plan.BackupPlanId,
			MaxResults:   aws.Int32(1000),
		},
		gomock.Any(),
	).Return(
		&backup.ListBackupSelectionsOutput{
			BackupSelectionsList: []types.BackupSelectionsListMember{{SelectionId: selection.SelectionId}},
		},
		nil,
	)

	m.EXPECT().GetBackupSelection(
		gomock.Any(),
		&backup.GetBackupSelectionInput{
			BackupPlanId: plan.BackupPlanId,
			SelectionId:  selection.SelectionId,
		},
		gomock.Any(),
	).Return(&selection, nil)

	return client.Services{
		Backup: m,
	}
}

func TestPlans(t *testing.T) {
	client.AwsMockTestHelper(t, Plans(), buildBackupPlansMock, client.TestOptions{})
}
