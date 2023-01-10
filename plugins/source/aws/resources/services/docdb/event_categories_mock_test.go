package docdb

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/docdb"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildEventCategoriesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockDocdbClient(ctrl)
	services := client.Services{
		Docdb: m,
	}
	var ec docdb.DescribeEventCategoriesOutput
	if err := faker.FakeObject(&ec); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeEventCategories(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec,
		nil,
	)

	return services
}

func TestEventCategories(t *testing.T) {
	client.AwsMockTestHelper(t, EventCategories(), buildEventCategoriesMock, client.TestOptions{})
}
