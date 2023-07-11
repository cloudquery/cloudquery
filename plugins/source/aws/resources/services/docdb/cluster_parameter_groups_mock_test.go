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

func buildClusterParameterGroupsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockDocdbClient(ctrl)
	services := client.Services{
		Docdb: m,
	}
	var parameterGroups docdb.DescribeDBClusterParameterGroupsOutput
	require.NoError(t, faker.FakeObject(&parameterGroups))

	parameterGroups.Marker = nil
	m.EXPECT().DescribeDBClusterParameterGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&parameterGroups,
		nil,
	)

	var parameters docdb.DescribeDBClusterParametersOutput
	require.NoError(t, faker.FakeObject(&parameters))

	parameters.Marker = nil
	m.EXPECT().DescribeDBClusterParameters(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&parameters,
		nil,
	)

	var tags docdb.ListTagsForResourceOutput
	require.NoError(t, faker.FakeObject(&tags))

	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&tags,
		nil,
	)

	return services
}

func TestClusterParameterGroups(t *testing.T) {
	client.AwsMockTestHelper(t, ClusterParameterGroups(), buildClusterParameterGroupsMock, client.TestOptions{})
}
