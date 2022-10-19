package docdb

import (
	"github.com/aws/aws-sdk-go-v2/service/docdb"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"testing"
)

func buildInstancesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockDocDBClient(ctrl)
	services := client.Services{
		DocDB: m,
	}
	var ev docdb.DescribeDBInstancesOutput
	if err := faker.FakeObject(&ev); err != nil {
		t.Fatal(err)
	}
	ev.Marker = nil
	m.EXPECT().DescribeDBInstances(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ev,
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

func TestInstances(t *testing.T) {
	client.AwsMockTestHelper(t, Instances(), buildInstancesMock, client.TestOptions{})
}
