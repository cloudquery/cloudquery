package ssm

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildAssociations(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockSsmClient(ctrl)

	var i types.Association
	if err := faker.FakeObject(&i); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListAssociations(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(
		&ssm.ListAssociationsOutput{Associations: []types.Association{i}},
		nil,
	)

	return client.Services{Ssm: mock}
}

func TestAssociations(t *testing.T) {
	client.AwsMockTestHelper(t, Associations(), buildAssociations, client.TestOptions{})
}
