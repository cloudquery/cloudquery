// +build mock

package sns

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildSnsTopics(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockSnsClient(ctrl)
	topic := types.Topic{}
	err := faker.FakeData(&topic)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListTopics(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&sns.ListTopicsOutput{
			Topics: []types.Topic{topic},
		}, nil)
	m.EXPECT().GetTopicAttributes(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&sns.GetTopicAttributesOutput{
			Attributes: map[string]string{
				"SubscriptionsConfirmed":    "5",
				"SubscriptionsDeleted":      "3",
				"SubscriptionsPending":      "0",
				"FifoTopic":                 "false",
				"ContentBasedDeduplication": "true",
				"DisplayName":               "cloudquery",
				"KmsMasterKeyId":            "test/key",
				"Owner":                     "owner",
				"Policy":                    `{"stuff": 3}`,
				"DeliveryPolicy":            `{"stuff": 3}`,
				"EffectiveDeliveryPolicy":   `{"stuff": 3}`,
			},
		}, nil)
	return client.Services{
		SNS: m,
	}
}

func TestSnsTopics(t *testing.T) {
	client.AwsMockTestHelper(t, SnsTopics(), buildSnsTopics, client.TestOptions{})
}
