package lightsail

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildBucketsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockLightsailClient(ctrl)

	b := lightsail.GetBucketsOutput{}
	require.NoError(t, faker.FakeObject(&b))
	b.NextPageToken = nil
	m.EXPECT().GetBuckets(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&b, nil)

	ac := lightsail.GetBucketAccessKeysOutput{}
	require.NoError(t, faker.FakeObject(&ac))

	m.EXPECT().GetBucketAccessKeys(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ac, nil)

	return client.Services{
		Lightsail: m,
	}
}

func TestBuckets(t *testing.T) {
	client.AwsMockTestHelper(t, Buckets(), buildBucketsMock, client.TestOptions{})
}
