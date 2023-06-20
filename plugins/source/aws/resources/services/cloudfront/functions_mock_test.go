package cloudfront

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	cloudfrontTypes "github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v3/faker"
	"github.com/golang/mock/gomock"
)

func buildCloudfronFunctionsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockCloudfrontClient(ctrl)
	services := client.Services{
		Cloudfront: m,
	}
	fs := cloudfrontTypes.FunctionSummary{}
	if err := faker.FakeObject(&fs); err != nil {
		t.Fatal(err)
	}
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
	if err := faker.FakeObject(&function); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeFunction(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		function,
		nil,
	)

	return services
}

func TestCloudfrontFunctions(t *testing.T) {
	client.AwsMockTestHelper(t, Functions(), buildCloudfronFunctionsMock, client.TestOptions{})
}
