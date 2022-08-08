package xray

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/xray"
	"github.com/aws/aws-sdk-go-v2/service/xray/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildEncryptionConfig(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockXrayClient(ctrl)

	var config types.EncryptionConfig
	if err := faker.FakeData(&config); err != nil {
		t.Fatal(err)
	}

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
