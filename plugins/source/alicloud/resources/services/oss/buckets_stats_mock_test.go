package oss

import (
	"testing"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/cloudquery/cloudquery/plugins/source/alicloud/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
)

func buildOssBucketStats(t *testing.T, mock *mocks.MockOssClient, bucketName string) {
	var b oss.GetBucketStatResult
	if err := faker.FakeObject(&b); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().GetBucketStat(bucketName).Return(b, nil)
}
