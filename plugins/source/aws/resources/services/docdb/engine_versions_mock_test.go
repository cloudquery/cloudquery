package docdb

import (
	"github.com/aws/aws-sdk-go-v2/service/docdb"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"testing"
)

func buildEngineVersionsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockDocDBClient(ctrl)
	services := client.Services{
		DocDB: m,
	}
	var ev docdb.DescribeDBEngineVersionsOutput
	if err := faker.FakeObject(&ev); err != nil {
		t.Fatal(err)
	}
	ev.Marker = nil
	m.EXPECT().DescribeDBEngineVersions(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ev,
		nil,
	)

	return services
}

func TestEngineVersions(t *testing.T) {
	client.AwsMockTestHelper(t, EngineVersions(), buildEngineVersionsMock, client.TestOptions{})
}
