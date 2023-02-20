package support

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/support"
	"github.com/aws/aws-sdk-go-v2/service/support/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildCases(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockSupportClient(ctrl)
	details := []types.CaseDetails{}
	err := faker.FakeObject(&details)
	if err != nil {
		t.Fatal(err)
	}

	input := support.DescribeCasesInput{MaxResults: aws.Int32(100), IncludeResolvedCases: true}
	m.EXPECT().DescribeCases(gomock.Any(), &input).Return(&support.DescribeCasesOutput{Cases: details}, nil)

	err = mockCommunications(details[0], m)
	if err != nil {
		t.Fatal(err)
	}

	return client.Services{
		Support: m,
	}
}

func TestCases(t *testing.T) {
	client.AwsMockTestHelper(t, Cases(), buildCases, client.TestOptions{})
}
