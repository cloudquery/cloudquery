package docdb

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/docdb"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildClusterParameterGroupsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockDocdbClient(ctrl)
	services := client.Services{
		Docdb: m,
	}
	var parameterGroups docdb.DescribeDBClusterParameterGroupsOutput
	if err := faker.FakeObject(&parameterGroups); err != nil {
		t.Fatal(err)
	}
	parameterGroups.Marker = nil
	m.EXPECT().DescribeDBClusterParameterGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&parameterGroups,
		nil,
	)

	var parameters docdb.DescribeDBClusterParametersOutput
	if err := faker.FakeObject(&parameters); err != nil {
		t.Fatal(err)
	}
	parameters.Marker = nil
	m.EXPECT().DescribeDBClusterParameters(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&parameters,
		nil,
	)

	var tags docdb.ListTagsForResourceOutput
	if err := faker.FakeObject(&tags); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&tags,
		nil,
	)

	return services
}

func TestClusterParameterGroups(t *testing.T) {
	client.AwsMockTestHelper(t, ClusterParameterGroups(), buildClusterParameterGroupsMock, client.TestOptions{})
}
