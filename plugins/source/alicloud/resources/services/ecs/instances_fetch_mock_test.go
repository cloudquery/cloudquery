package ecs

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/cloudquery/cloudquery/plugins/source/alicloud/client"
	"github.com/cloudquery/cloudquery/plugins/source/alicloud/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildEcsInstances(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockEcsClient(ctrl)

	b := ecs.CreateDescribeInstancesResponse()
	if err := faker.FakeObject(b); err != nil {
		t.Fatal(err)
	}
	b.BaseResponse = fakeSuccessRequest(t)
	b.TotalCount = 2
	mock.EXPECT().DescribeInstances(gomock.Any()).Times(2).Return(b, nil)
	return client.Services{ECS: mock}
}

func fakeSuccessRequest(t *testing.T) *responses.BaseResponse {
	baseResponse := &responses.BaseResponse{}
	resp := &http.Response{
		Status:     http.StatusText(http.StatusOK),
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(strings.NewReader("{}")),
	}
	err := responses.Unmarshal(baseResponse, resp, "JSON")
	if err != nil {
		t.Fatal(err)
	}
	return baseResponse
}

func TestEcsInstances(t *testing.T) {
	client.MockTestHelper(t, Instances(), buildEcsInstances, client.TestOptions{})
}
