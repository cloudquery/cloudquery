package sns

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildSnsSubscriptions(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockSnsClient(ctrl)
	sub := types.Subscription{}
	err := faker.FakeObject(&sub)
	if err != nil {
		t.Fatal(err)
	}

	subTemp := types.Subscription{}
	err = faker.FakeObject(&subTemp)
	if err != nil {
		t.Fatal(err)
	}
	emptySub := types.Subscription{
		SubscriptionArn: aws.String("PendingConfirmation"),
		Owner:           subTemp.Owner,
		Protocol:        subTemp.Protocol,
		TopicArn:        subTemp.TopicArn,
		Endpoint:        subTemp.Endpoint,
	}

	m.EXPECT().ListSubscriptions(
		gomock.Any(),
		&sns.ListSubscriptionsInput{},
	).Return(
		&sns.ListSubscriptionsOutput{
			Subscriptions: []types.Subscription{sub, emptySub},
		}, nil)

	m.EXPECT().GetSubscriptionAttributes(
		gomock.Any(),
		&sns.GetSubscriptionAttributesInput{SubscriptionArn: sub.SubscriptionArn},
	).Return(
		&sns.GetSubscriptionAttributesOutput{Attributes: map[string]string{
			"ConfirmationWasAuthenticated": "true",
			"DeliveryPolicy":               "{}",
			"EffectiveDeliveryPolicy":      "{}",
			"FilterPolicy":                 "{}",
			"PendingConfirmation":          "true",
			"RawMessageDelivery":           "true",
			"RedrivePolicy":                `{"deadLetterTargetArn": "test"}`,
			"SubscriptionRoleArn":          "some",
			"WeirdAndUnexpectedField":      "needs updating",
		}},
		nil,
	)

	return client.Services{
		Sns: m,
	}
}

func TestSnsSubscriptions(t *testing.T) {
	client.AwsMockTestHelper(t, Subscriptions(), buildSnsSubscriptions, client.TestOptions{})
}
