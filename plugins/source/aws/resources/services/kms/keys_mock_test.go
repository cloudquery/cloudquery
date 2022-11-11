package kms

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/aws/smithy-go"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildKmsKeys(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockKmsClient(ctrl)

	keyListEntry1, keyListEntry2 := types.KeyListEntry{}, types.KeyListEntry{}
	if err := faker.FakeObject(&keyListEntry1); err != nil {
		t.Fatal(err)
	}
	if err := faker.FakeObject(&keyListEntry2); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListKeys(gomock.Any(), gomock.Any(), gomock.Any()).Return(&kms.ListKeysOutput{Keys: []types.KeyListEntry{keyListEntry1, keyListEntry2}}, nil)

	tags := kms.ListResourceTagsOutput{}
	if err := faker.FakeObject(&tags); err != nil {
		t.Fatal(err)
	}
	tags.NextMarker = nil
	m.EXPECT().ListResourceTags(gomock.Any(), gomock.Any(), gomock.Any()).Return(&tags, nil).MaxTimes(2)

	key1 := kms.DescribeKeyOutput{}
	if err := faker.FakeObject(&key1); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeKey(gomock.Any(), &kms.DescribeKeyInput{KeyId: keyListEntry1.KeyId}, gomock.Any()).Return(&key1, nil)
	m.EXPECT().DescribeKey(gomock.Any(), &kms.DescribeKeyInput{KeyId: keyListEntry2.KeyId}, gomock.Any()).Return(nil, &smithy.GenericAPIError{Code: "NotFound"})

	rotation := kms.GetKeyRotationStatusOutput{}
	if err := faker.FakeObject(&rotation); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetKeyRotationStatus(gomock.Any(), gomock.Any(), gomock.Any()).Return(&rotation, nil).MaxTimes(1)

	g := kms.ListGrantsOutput{}
	if err := faker.FakeObject(&g); err != nil {
		t.Fatal(err)
	}
	g.NextMarker = nil
	m.EXPECT().ListGrants(gomock.Any(), gomock.Any(), gomock.Any()).Return(&g, nil).MaxTimes(1)

	return client.Services{
		Kms: m,
	}
}

func TestKmsKeys(t *testing.T) {
	client.AwsMockTestHelper(t, Keys(), buildKmsKeys, client.TestOptions{})
}
