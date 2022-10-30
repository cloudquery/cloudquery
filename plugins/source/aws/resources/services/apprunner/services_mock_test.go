package apprunner

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/apprunner"
	"github.com/aws/aws-sdk-go-v2/service/apprunner/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildApprunnerGraphqlApisMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockAppRunnerClient(ctrl)
	s := types.Service{}
	err := faker.FakeObject(&s)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListServices(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apprunner.ListServicesOutput{
			ServiceSummaryList: []types.ServiceSummary{
				{ServiceArn: s.ServiceArn},
			},
		}, nil)

	m.EXPECT().DescribeService(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apprunner.DescribeServiceOutput{
			Service: &s,
		}, nil)

	return client.Services{
		Apprunner: m,
	}
}

func TestAppSyncGraphqlApis(t *testing.T) {
	client.AwsMockTestHelper(t, Services(), buildApprunnerGraphqlApisMock, client.TestOptions{})
}
