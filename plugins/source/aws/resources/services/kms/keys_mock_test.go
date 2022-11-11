package kms

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildKmsKeys(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockKmsClient(ctrl)

	keys := kms.ListKeysOutput{}
	err := faker.FakeObject(&keys)
	if err != nil {
		t.Fatal(err)
	}
	keys.NextMarker = nil
	m.EXPECT().ListKeys(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&keys, nil)

	tags := kms.ListResourceTagsOutput{}
	err = faker.FakeObject(&tags)
	if err != nil {
		t.Fatal(err)
	}
	tags.NextMarker = nil
	m.EXPECT().ListResourceTags(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&tags, nil)

	key := kms.DescribeKeyOutput{}
	err = faker.FakeObject(&key)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeKey(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&key, nil)

	rotation := kms.GetKeyRotationStatusOutput{}
	err = faker.FakeObject(&rotation)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetKeyRotationStatus(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&rotation, nil)

	g := kms.ListGrantsOutput{}
	err = faker.FakeObject(&g)
	if err != nil {
		t.Fatal(err)
	}
	g.NextMarker = nil
	m.EXPECT().ListGrants(gomock.Any(), gomock.Any(), gomock.Any()).Return(&g, nil)

	return client.Services{
		Kms: m,
	}
}

func TestKmsKeys(t *testing.T) {
	client.AwsMockTestHelper(t, Keys(), buildKmsKeys, client.TestOptions{})
}
