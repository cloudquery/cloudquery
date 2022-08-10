package lightsail

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildBucketsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockLightsailClient(ctrl)

	b := lightsail.GetBucketsOutput{}
	err := faker.FakeData(&b)
	if err != nil {
		t.Fatal(err)
	}
	b.NextPageToken = nil
	m.EXPECT().GetBuckets(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&b, nil)

	ac := lightsail.GetBucketAccessKeysOutput{}
	err = faker.FakeData(&ac)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().GetBucketAccessKeys(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ac, nil)

	return client.Services{
		Lightsail: m,
	}
}

func TestBuckets(t *testing.T) {
	client.AwsMockTestHelper(t, Buckets(), buildBucketsMock, client.TestOptions{})
}
