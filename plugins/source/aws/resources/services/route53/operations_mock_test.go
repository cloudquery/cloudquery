package route53

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/route53domains"
	"github.com/aws/aws-sdk-go-v2/service/route53domains/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildRoute53Operations(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockRoute53domainsClient(ctrl)

	var os types.OperationSummary
	require.NoError(t, faker.FakeObject(&os))

	mock.EXPECT().ListOperations(gomock.Any(), &route53domains.ListOperationsInput{}, gomock.Any()).Return(
		&route53domains.ListOperationsOutput{Operations: []types.OperationSummary{os}},
		nil,
	)

	var detail route53domains.GetOperationDetailOutput
	require.NoError(t, faker.FakeObject(&detail))

	detail.OperationId = os.OperationId
	mock.EXPECT().GetOperationDetail(gomock.Any(), &route53domains.GetOperationDetailInput{OperationId: os.OperationId}, gomock.Any()).Return(
		&detail, nil,
	)

	return client.Services{
		Route53domains: mock,
	}
}

func TestRoute53Operations(t *testing.T) {
	client.AwsMockTestHelper(t, Operations(), buildRoute53Operations, client.TestOptions{})
}
