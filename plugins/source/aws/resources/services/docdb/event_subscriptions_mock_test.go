package docdb

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/docdb"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildEventSubscriptionsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockDocdbClient(ctrl)
	services := client.Services{
		Docdb: m,
	}
	var ec docdb.DescribeEventSubscriptionsOutput
	if err := faker.FakeObject(&ec); err != nil {
		t.Fatal(err)
	}
	ec.Marker = nil
	m.EXPECT().DescribeEventSubscriptions(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec,
		nil,
	)

	return services
}

func TestEventSubscriptions(t *testing.T) {
	client.AwsMockTestHelper(t, EventSubscriptions(), buildEventSubscriptionsMock, client.TestOptions{})
}
