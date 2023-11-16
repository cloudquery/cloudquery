package neptune

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/neptune"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildNeptuneEventSubscriptions(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockNeptuneClient(ctrl)
	var s types.EventSubscription
	require.NoError(t, faker.FakeObject(&s))

	mock.EXPECT().DescribeEventSubscriptions(gomock.Any(), &neptune.DescribeEventSubscriptionsInput{
		Filters: []types.Filter{{Name: aws.String("engine"), Values: []string{"neptune"}}},
	}, gomock.Any()).Return(
		&neptune.DescribeEventSubscriptionsOutput{EventSubscriptionsList: []types.EventSubscription{s}},
		nil,
	)

	mock.EXPECT().ListTagsForResource(
		gomock.Any(),
		&neptune.ListTagsForResourceInput{ResourceName: s.EventSubscriptionArn},
		gomock.Any(),
	).Return(
		&neptune.ListTagsForResourceOutput{
			TagList: []types.Tag{{Key: aws.String("key"), Value: aws.String("value")}},
		},
		nil,
	)
	return client.Services{Neptune: mock}
}

func TestNeptuneEventSubscriptions(t *testing.T) {
	client.AwsMockTestHelper(t, EventSubscriptions(), buildNeptuneEventSubscriptions, client.TestOptions{})
}
