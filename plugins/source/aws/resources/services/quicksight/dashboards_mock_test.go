package quicksight

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/quicksight"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildDashboardsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockQuicksightClient(ctrl)

	var ld quicksight.ListDashboardsOutput
	if err := faker.FakeObject(&ld); err != nil {
		t.Fatal(err)
	}
	ld.NextToken = nil
	m.EXPECT().ListDashboards(gomock.Any(), gomock.Any(), gomock.Any()).Return(&ld, nil)

	var to quicksight.ListTagsForResourceOutput
	if err := faker.FakeObject(&to); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(&to, nil)

	return client.Services{
		Quicksight: m,
	}
}
func TestQuicksightDashboards(t *testing.T) {
	client.AwsMockTestHelper(t, Dashboards(), buildDashboardsMock, client.TestOptions{})
}
