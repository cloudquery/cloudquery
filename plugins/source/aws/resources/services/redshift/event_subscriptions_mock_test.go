package redshift

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildEventSubscriptionsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRedshiftClient(ctrl)

	var s types.EventSubscription
	require.NoError(t, faker.FakeObject(&s))

	m.EXPECT().DescribeEventSubscriptions(
		gomock.Any(),
		&redshift.DescribeEventSubscriptionsInput{MaxRecords: aws.Int32(100)},
		gomock.Any(),
	).Return(
		&redshift.DescribeEventSubscriptionsOutput{
			EventSubscriptionsList: []types.EventSubscription{s},
		},
		nil,
	)

	return client.Services{
		Redshift: m,
	}
}

func TestRedshiftEventSubscriptions(t *testing.T) {
	client.AwsMockTestHelper(t, EventSubscriptions(), buildEventSubscriptionsMock, client.TestOptions{})
}
