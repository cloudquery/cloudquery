package xray

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/xray"
	"github.com/aws/aws-sdk-go-v2/service/xray/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildResourcePolicies(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockXrayClient(ctrl)

	var pols types.ResourcePolicy
	require.NoError(t, faker.FakeObject(&pols))

	mock.EXPECT().ListResourcePolicies(
		gomock.Any(),
		&xray.ListResourcePoliciesInput{},
		gomock.Any(),
	).Return(
		&xray.ListResourcePoliciesOutput{
			ResourcePolicies: []types.ResourcePolicy{
				pols,
			},
		},
		nil,
	)

	return client.Services{Xray: mock}
}

func TestResourcePolicies(t *testing.T) {
	client.AwsMockTestHelper(t, ResourcePolicies(), buildResourcePolicies, client.TestOptions{})
}
