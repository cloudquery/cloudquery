package route53

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/route53domains"
	"github.com/aws/aws-sdk-go-v2/service/route53domains/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildRoute53Operations(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockRoute53domainsClient(ctrl)

	var os types.OperationSummary
	if err := faker.FakeObject(&os); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListOperations(gomock.Any(), &route53domains.ListOperationsInput{}, gomock.Any()).Return(
		&route53domains.ListOperationsOutput{Operations: []types.OperationSummary{os}},
		nil,
	)

	var detail route53domains.GetOperationDetailOutput
	if err := faker.FakeObject(&detail); err != nil {
		t.Fatal(err)
	}
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
