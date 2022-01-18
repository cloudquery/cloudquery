//go:build mock
// +build mock

package rds

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildRDSEventSubscriptions(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockRdsClient(ctrl)
	var s types.EventSubscription
	if err := faker.FakeData(&s); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().DescribeEventSubscriptions(gomock.Any(), &rds.DescribeEventSubscriptionsInput{}, gomock.Any()).Return(
		&rds.DescribeEventSubscriptionsOutput{EventSubscriptionsList: []types.EventSubscription{s}},
		nil,
	)

	mock.EXPECT().ListTagsForResource(
		gomock.Any(),
		&rds.ListTagsForResourceInput{ResourceName: s.EventSubscriptionArn},
		gomock.Any(),
	).Return(
		&rds.ListTagsForResourceOutput{
			TagList: []types.Tag{{Key: aws.String("key"), Value: aws.String("value")}},
		},
		nil,
	)
	return client.Services{RDS: mock}
}

func TestRDSEventSubscriptions(t *testing.T) {
	client.AwsMockTestHelper(t, RdsEventSubscriptions(), buildRDSEventSubscriptions, client.TestOptions{})
}
