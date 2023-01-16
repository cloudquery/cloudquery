package ecs

import (
	"testing"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/cloudquery/cloudquery/plugins/source/alicloud/client"
	"github.com/cloudquery/cloudquery/plugins/source/alicloud/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildEcsInstances(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockEcsClient(ctrl)

	var b ecs.DescribeInstancesResponse
	if err := faker.FakeObject(&b); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().DescribeInstances(gomock.Any()).Return(&b, nil)
	return client.Services{ECS: mock}
}

func TestEcsInstances(t *testing.T) {
	client.MockTestHelper(t, Instances(), buildEcsInstances, client.TestOptions{})
}
