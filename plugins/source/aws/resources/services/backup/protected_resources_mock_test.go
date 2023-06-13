package backup

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/backup"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v3/faker"
	"github.com/golang/mock/gomock"
)

func buildProtectedResourcesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockBackupClient(ctrl)

	var pr types.ProtectedResource
	if err := faker.FakeObject(&pr); err != nil {
		t.Fatal(err)
	}

	var dpo backup.DescribeProtectedResourceOutput
	if err := faker.FakeObject(&dpo); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListProtectedResources(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(
		&backup.ListProtectedResourcesOutput{
			Results: []types.ProtectedResource{pr},
		},
		nil,
	)

	m.EXPECT().DescribeProtectedResource(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(
		&dpo,
		nil,
	)

	return client.Services{
		Backup: m,
	}
}

func TestProtectedResources(t *testing.T) {
	client.AwsMockTestHelper(t, ProtectedResources(), buildProtectedResourcesMock, client.TestOptions{})
}
