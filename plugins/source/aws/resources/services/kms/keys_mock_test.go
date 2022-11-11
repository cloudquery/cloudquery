package kms

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildKmsKeys(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockKmsClient(ctrl)

	keyListEntry := types.KeyListEntry{}
	if err := faker.FakeObject(&keyListEntry); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListKeys(gomock.Any(), gomock.Any(), gomock.Any()).Return(&kms.ListKeysOutput{Keys: []types.KeyListEntry{keyListEntry}}, nil)

	tags := kms.ListResourceTagsOutput{}
	if err := faker.FakeObject(&tags); err != nil {
		t.Fatal(err)
	}
	tags.NextMarker = nil
	m.EXPECT().ListResourceTags(gomock.Any(), gomock.Any(), gomock.Any()).Return(&tags, nil)

	key := kms.DescribeKeyOutput{}
	if err := faker.FakeObject(&key); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeKey(gomock.Any(), &kms.DescribeKeyInput{KeyId: keyListEntry.KeyId}, gomock.Any()).Return(&key, nil)

	rotation := kms.GetKeyRotationStatusOutput{}
	if err := faker.FakeObject(&rotation); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetKeyRotationStatus(gomock.Any(), gomock.Any(), gomock.Any()).Return(&rotation, nil)

	g := kms.ListGrantsOutput{}
	if err := faker.FakeObject(&g); err != nil {
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
