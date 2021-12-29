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

func buildSnsSubscriptions(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockSnsClient(ctrl)
	sub := types.Subscription{}
	err := faker.FakeData(&sub)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListSubscriptions(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&sns.ListSubscriptionsOutput{
			Subscriptions: []types.Subscription{sub},
		}, nil)
	return client.Services{
		SNS: m,
	}
}

func TestSnsSubscriptions(t *testing.T) {
	client.AwsMockTestHelper(t, SnsSubscriptions(), buildSnsSubscriptions, client.TestOptions{})
}
