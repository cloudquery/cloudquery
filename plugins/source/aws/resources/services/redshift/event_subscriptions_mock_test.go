package redshift

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildRedshiftEventSubscriptionsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRedshiftClient(ctrl)

	var s types.EventSubscription
	if err := faker.FakeData(&s); err != nil {
		t.Fatal(err)
	}
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
	client.AwsMockTestHelper(t, EventSubscriptions(), buildRedshiftEventSubscriptionsMock, client.TestOptions{})
}
