package rds

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildEngineVersionsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRdsClient(ctrl)
	services := client.Services{
		Rds: m,
	}
	var ev rds.DescribeDBEngineVersionsOutput
	if err := faker.FakeObject(&ev); err != nil {
		t.Fatal(err)
	}
	ev.Marker = nil
	m.EXPECT().DescribeDBEngineVersions(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ev,
		nil,
	)

	var parameters rds.DescribeEngineDefaultClusterParametersOutput
	if err := faker.FakeObject(&parameters); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeEngineDefaultClusterParameters(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&parameters,
		nil,
	)

	return services
}

func TestEngineVersions(t *testing.T) {
	client.AwsMockTestHelper(t, EngineVersions(), buildEngineVersionsMock, client.TestOptions{})
}
