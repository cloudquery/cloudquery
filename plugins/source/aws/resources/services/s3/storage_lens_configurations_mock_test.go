package s3

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/s3control"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildStorageLensConfigurations(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockS3controlClient(ctrl)
	lo := s3control.ListStorageLensConfigurationsOutput{}
	if err := faker.FakeObject(&lo); err != nil {
		t.Fatal(err)
	}
	lo.NextToken = nil

	out := s3control.GetStorageLensConfigurationOutput{}
	if err := faker.FakeObject(&out); err != nil {
		t.Fatal(err)
	}

	to := s3control.GetStorageLensConfigurationTaggingOutput{}
	if err := faker.FakeObject(&to); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListStorageLensConfigurations(gomock.Any(), gomock.Any(), gomock.Any()).Return(&lo, nil)
	m.EXPECT().GetStorageLensConfiguration(gomock.Any(), gomock.Any(), gomock.Any()).Return(&out, nil)
	m.EXPECT().GetStorageLensConfigurationTagging(gomock.Any(), gomock.Any(), gomock.Any()).Return(&to, nil)

	return client.Services{
		S3control: m,
	}
}

func TestStorageLensConfigurations(t *testing.T) {
	client.AwsMockTestHelper(t, StorageLensConfigurations(), buildStorageLensConfigurations, client.TestOptions{})
}
