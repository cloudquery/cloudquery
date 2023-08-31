package backup

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/backup"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildReportPlansMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockBackupClient(ctrl)

	var reportPlan types.ReportPlan
	require.NoError(t, faker.FakeObject(&reportPlan))

	m.EXPECT().ListReportPlans(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(
		&backup.ListReportPlansOutput{
			ReportPlans: []types.ReportPlan{reportPlan},
		},
		nil,
	)

	m.EXPECT().ListTags(
		gomock.Any(),
		&backup.ListTagsInput{ResourceArn: reportPlan.ReportPlanArn},
		gomock.Any(),
	).Return(
		&backup.ListTagsOutput{
			Tags: map[string]string{"plan1": "value1"},
		},
		nil,
	)

	return client.Services{
		Backup: m,
	}
}

func TestReportPlans(t *testing.T) {
	client.AwsMockTestHelper(t, ReportPlans(), buildReportPlansMock, client.TestOptions{})
}
