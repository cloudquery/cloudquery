package efs

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/efs"
	"github.com/aws/aws-sdk-go-v2/service/efs/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v3/faker"
	"github.com/golang/mock/gomock"
)

func buildEfsAccessPointsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEfsClient(ctrl)
	l := types.AccessPointDescription{}
	err := faker.FakeObject(&l)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeAccessPoints(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&efs.DescribeAccessPointsOutput{
			AccessPoints: []types.AccessPointDescription{l},
		}, nil)

	return client.Services{
		Efs: m,
	}
}

func TestEfsAccesspoints(t *testing.T) {
	client.AwsMockTestHelper(t, AccessPoints(), buildEfsAccessPointsMock, client.TestOptions{})
}
