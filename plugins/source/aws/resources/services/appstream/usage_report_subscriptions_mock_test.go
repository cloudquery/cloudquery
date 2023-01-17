package appstream

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/appstream"
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildAppstreamUsageReportSubscriptionsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockAppstreamClient(ctrl)
	object := types.UsageReportSubscription{}
	err := faker.FakeObject(&object)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeUsageReportSubscriptions(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&appstream.DescribeUsageReportSubscriptionsOutput{
			UsageReportSubscriptions: []types.UsageReportSubscription{object},
		}, nil)

	tagsOutput := appstream.ListTagsForResourceOutput{}
	err = faker.FakeObject(&tagsOutput)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any()).Return(&tagsOutput, nil).AnyTimes()

	return client.Services{
		Appstream: m,
	}
}
func TestAppstreamUsageReportSubscriptions(t *testing.T) {
	client.AwsMockTestHelper(t, UsageReportSubscriptions(), buildAppstreamUsageReportSubscriptionsMock, client.TestOptions{})
}
