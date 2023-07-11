package support

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/support"
	"github.com/aws/aws-sdk-go-v2/service/support/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildCases(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockSupportClient(ctrl)
	details := []types.CaseDetails{}
	require.NoError(t, faker.FakeObject(&details))

	input := support.DescribeCasesInput{MaxResults: aws.Int32(100), IncludeResolvedCases: true}
	m.EXPECT().DescribeCases(gomock.Any(), &input, gomock.Any()).Return(&support.DescribeCasesOutput{Cases: details}, nil)

	require.NoError(t, mockCommunications(details[0], m))

	return client.Services{
		Support: m,
	}
}

func TestCases(t *testing.T) {
	client.AwsMockTestHelper(t, Cases(), buildCases, client.TestOptions{})
}
