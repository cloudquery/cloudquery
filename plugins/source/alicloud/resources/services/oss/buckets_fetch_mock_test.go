package oss

import (
	"testing"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/cloudquery/cloudquery/plugins/source/alicloud/client"
	"github.com/cloudquery/cloudquery/plugins/source/alicloud/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
)

func buildOssBuckets(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockOssClient(ctrl)

	var b oss.ListBucketsResult
	if err := faker.FakeObject(&b); err != nil {
		t.Fatal(err)
	}
	b.Buckets[0].Location = "cn-hangzhou"
	mock.EXPECT().ListBuckets().Return(b, nil)

	buildOssBucketStats(t, mock, b.Buckets[0].Name)

	return client.Services{OSS: mock}
}

func TestOssBuckets(t *testing.T) {
	client.MockTestHelper(t, Buckets(), buildOssBuckets, client.TestOptions{})
}
