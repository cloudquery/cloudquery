package fsx

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/fsx"
	"github.com/aws/aws-sdk-go-v2/service/fsx/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildVolumesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockFsxClient(ctrl)

	var v types.Volume
	if err := faker.FakeObject(&v); err != nil {
		t.Fatal(err)
	}
	v.Lifecycle = types.VolumeLifecycleAvailable
	v.VolumeType = types.VolumeTypeOntap
	m.EXPECT().DescribeVolumes(
		gomock.Any(),
		&fsx.DescribeVolumesInput{MaxResults: aws.Int32(1000)},
	).Return(
		&fsx.DescribeVolumesOutput{Volumes: []types.Volume{v}},
		nil,
	)

	return client.Services{
		Fsx: m,
	}
}

func TestVolumes(t *testing.T) {
	client.AwsMockTestHelper(t, Volumes(), buildVolumesMock, client.TestOptions{})
}
