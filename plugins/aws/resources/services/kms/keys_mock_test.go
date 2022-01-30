package kms

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildKmsKeys(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockKmsClient(ctrl)

	keys := kms.ListKeysOutput{}
	err := faker.FakeData(&keys)
	if err != nil {
		t.Fatal(err)
	}
	keys.NextMarker = nil
	m.EXPECT().ListKeys(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&keys, nil)

	tags := kms.ListResourceTagsOutput{}
	err = faker.FakeData(&tags)
	if err != nil {
		t.Fatal(err)
	}
	tags.NextMarker = nil
	m.EXPECT().ListResourceTags(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&tags, nil)

	key := kms.DescribeKeyOutput{}
	err = faker.FakeData(&key)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeKey(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&key, nil)

	rotation := kms.GetKeyRotationStatusOutput{}
	err = faker.FakeData(&rotation)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetKeyRotationStatus(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&rotation, nil)

	return client.Services{
		KMS: m,
	}
}

func TestKmsKeys(t *testing.T) {
	client.AwsMockTestHelper(t, KmsKeys(), buildKmsKeys, client.TestOptions{})
}
