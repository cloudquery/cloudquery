package client

import (
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/bssopenapi"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/cloudquery/cloudquery/plugins/source/alicloud/client/services"
)

type Services struct {
	OSS services.OssClient
	BSS services.BssopenapiClient
}

func initServices(account AccountSpec, region string) (*Services, error) {
	ossCli, err := oss.New(fmt.Sprintf("oss-%s.aliyuncs.com", region), account.AccessKey, account.SecretKey, oss.Timeout(15, 30))
	if err != nil {
		return nil, fmt.Errorf("failed to initialize oss client: %w", err)
	}
	bssCli, err := bssopenapi.NewClientWithAccessKey(region, account.AccessKey, account.SecretKey)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize bssopenapi client: %w", err)
	}
	return &Services{
		OSS: ossCli,
		BSS: bssCli,
	}, nil
}
