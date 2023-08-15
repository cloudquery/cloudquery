package bss

import (
	"testing"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/bssopenapi"
	"github.com/cloudquery/cloudquery/plugins/source/alicloud/client"
	"github.com/cloudquery/cloudquery/plugins/source/alicloud/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
)

func buildBillDetails(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockBssopenapiClient(ctrl)

	var b *bssopenapi.QueryBillResponse
	if err := faker.FakeObject(&b); err != nil {
		t.Fatal(err)
	}
	b.Success = true
	b.Data.TotalCount = 2
	mock.EXPECT().QueryBill(gomock.Any()).AnyTimes().Return(b, nil)
	return client.Services{BSS: mock}
}

func TestBillDetails(t *testing.T) {
	client.MockTestHelper(t, Bill(), buildBillDetails, client.TestOptions{})
}
