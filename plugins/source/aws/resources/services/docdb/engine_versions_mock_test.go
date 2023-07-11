package docdb

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/docdb"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildEngineVersionsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockDocdbClient(ctrl)
	services := client.Services{
		Docdb: m,
	}
	var ev docdb.DescribeDBEngineVersionsOutput
	require.NoError(t, faker.FakeObject(&ev))

	ev.Marker = nil
	m.EXPECT().DescribeDBEngineVersions(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ev,
		nil,
	)

	var parameters docdb.DescribeEngineDefaultClusterParametersOutput
	require.NoError(t, faker.FakeObject(&parameters))

	m.EXPECT().DescribeEngineDefaultClusterParameters(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&parameters,
		nil,
	)

	var instanceOptions docdb.DescribeOrderableDBInstanceOptionsOutput
	require.NoError(t, faker.FakeObject(&instanceOptions))

	instanceOptions.Marker = nil
	m.EXPECT().DescribeOrderableDBInstanceOptions(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&instanceOptions,
		nil,
	)

	return services
}

func TestEngineVersions(t *testing.T) {
	client.AwsMockTestHelper(t, EngineVersions(), buildEngineVersionsMock, client.TestOptions{})
}
