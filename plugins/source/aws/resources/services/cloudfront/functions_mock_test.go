package cloudfront

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	cloudfrontTypes "github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildCloudfronFunctionsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockCloudfrontClient(ctrl)
	services := client.Services{
		Cloudfront: m,
	}
	fs := cloudfrontTypes.FunctionSummary{}
	require.NoError(t, faker.FakeObject(&fs))

	cloudfrontOutput := &cloudfront.ListFunctionsOutput{
		FunctionList: &cloudfrontTypes.FunctionList{
			Items: []cloudfrontTypes.FunctionSummary{fs},
		},
	}
	m.EXPECT().ListFunctions(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		cloudfrontOutput,
		nil,
	)

	function := &cloudfront.DescribeFunctionOutput{}
	require.NoError(t, faker.FakeObject(&function))

	m.EXPECT().DescribeFunction(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		function,
		nil,
	)

	return services
}

func TestCloudfrontFunctions(t *testing.T) {
	client.AwsMockTestHelper(t, Functions(), buildCloudfronFunctionsMock, client.TestOptions{})
}
