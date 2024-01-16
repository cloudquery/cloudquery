package kms

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/aws/smithy-go"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildKmsKeys(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockKmsClient(ctrl)

	keyListEntry := types.KeyListEntry{}
	require.NoError(t, faker.FakeObject(&keyListEntry))

	m.EXPECT().ListKeys(gomock.Any(), gomock.Any(), gomock.Any()).Return(&kms.ListKeysOutput{Keys: []types.KeyListEntry{keyListEntry, keyListEntry}}, nil)

	tags := kms.ListResourceTagsOutput{}
	require.NoError(t, faker.FakeObject(&tags))

	tags.NextMarker = nil
	m.EXPECT().ListResourceTags(gomock.Any(), gomock.Any(), gomock.Any()).Return(&tags, nil).Times(2)

	key := kms.DescribeKeyOutput{}
	require.NoError(t, faker.FakeObject(&key))

	err := smithy.GenericAPIError{Code: "AccessDenied", Message: "This is an error message"}

	// There are 2 calls to DescribeKey, one succeeds and the other fails
	gomock.InOrder(
		m.EXPECT().DescribeKey(gomock.Any(), &kms.DescribeKeyInput{KeyId: keyListEntry.KeyId}, gomock.Any()).Return(&key, nil),

		m.EXPECT().DescribeKey(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(nil, &err),
	)

	rotation := kms.GetKeyRotationStatusOutput{}
	require.NoError(t, faker.FakeObject(&rotation))

	m.EXPECT().GetKeyRotationStatus(gomock.Any(), gomock.Any(), gomock.Any()).Return(&rotation, nil).Times(2)

	g := kms.ListGrantsOutput{}
	require.NoError(t, faker.FakeObject(&g))

	g.NextMarker = nil
	m.EXPECT().ListGrants(gomock.Any(), gomock.Any(), gomock.Any()).Return(&g, nil).Times(2)

	pj := `{"data":["data"]}`
	m.EXPECT().GetKeyPolicy(gomock.Any(), &kms.GetKeyPolicyInput{
		KeyId:      keyListEntry.KeyId,
		PolicyName: aws.String("default"),
	}, gomock.Any()).Return(&kms.GetKeyPolicyOutput{Policy: &pj}, nil).Times(2)

	return client.Services{
		Kms: m,
	}
}

func TestKmsKeys(t *testing.T) {
	client.AwsMockTestHelper(t, Keys(), buildKmsKeys, client.TestOptions{})
}
