package sns

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/stretchr/testify/require"

	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
)

func buildSnsSubscriptions(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockSnsClient(ctrl)
	sub := types.Subscription{}
	require.NoError(t, faker.FakeObject(&sub))

	subTemp := types.Subscription{}
	require.NoError(t, faker.FakeObject(&subTemp))
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
		gomock.Any(),
	).Return(
		&sns.ListSubscriptionsOutput{
			Subscriptions: []types.Subscription{sub, emptySub},
		}, nil)

	m.EXPECT().GetSubscriptionAttributes(
		gomock.Any(),
		&sns.GetSubscriptionAttributesInput{SubscriptionArn: sub.SubscriptionArn},
		gomock.Any(),
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
