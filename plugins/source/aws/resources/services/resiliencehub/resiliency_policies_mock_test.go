package resiliencehub

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/resiliencehub"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildResiliencyPolicies(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockResiliencehubClient(ctrl)
	var l resiliencehub.ListResiliencyPoliciesOutput
	require.NoError(t, faker.FakeObject(&l))

	l.NextToken = nil
	mock.EXPECT().ListResiliencyPolicies(
		gomock.Any(),
		&resiliencehub.ListResiliencyPoliciesInput{},
		gomock.Any(),
	).Return(&l, nil)

	return client.Services{Resiliencehub: mock}
}

func TestResiilencehubResiliencyPolicies(t *testing.T) {
	client.AwsMockTestHelper(t, ResiliencyPolicies(), buildResiliencyPolicies, client.TestOptions{})
}
