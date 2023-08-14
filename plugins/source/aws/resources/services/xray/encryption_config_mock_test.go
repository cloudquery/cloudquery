package xray

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/xray"
	"github.com/aws/aws-sdk-go-v2/service/xray/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildEncryptionConfig(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockXrayClient(ctrl)

	var config types.EncryptionConfig
	require.NoError(t, faker.FakeObject(&config))

	mock.EXPECT().GetEncryptionConfig(
		gomock.Any(),
		&xray.GetEncryptionConfigInput{},
		gomock.Any(),
	).Return(
		&xray.GetEncryptionConfigOutput{
			EncryptionConfig: &config,
		},
		nil,
	)

	return client.Services{Xray: mock}
}

func TestXrayEncryptionConfig(t *testing.T) {
	client.AwsMockTestHelper(t, EncryptionConfigs(), buildEncryptionConfig, client.TestOptions{})
}
