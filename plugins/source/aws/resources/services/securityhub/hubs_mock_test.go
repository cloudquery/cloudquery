package securityhub

import (
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/securityhub"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildHubs(t *testing.T, ctrl *gomock.Controller) client.Services {
	shMock := mocks.NewMockSecurityhubClient(ctrl)
	hub := securityhub.DescribeHubOutput{}
	require.NoError(t, faker.FakeObject(&hub))
	hub.SubscribedAt = aws.String(time.Now().Format(time.RFC3339))

	shMock.EXPECT().DescribeHub(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(&hub, nil)

	tags := make(map[string]string)
	require.NoError(t, faker.FakeObject(&tags))
	shMock.EXPECT().ListTagsForResource(
		gomock.Any(),
		&securityhub.ListTagsForResourceInput{ResourceArn: hub.HubArn},
		gomock.Any(),
	).Return(&securityhub.ListTagsForResourceOutput{Tags: tags}, nil)

	return client.Services{Securityhub: shMock}
}

func TestHubs(t *testing.T) {
	client.AwsMockTestHelper(t, Hubs(), buildHubs, client.TestOptions{})
}
