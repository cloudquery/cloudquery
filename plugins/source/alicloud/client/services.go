package client

import (
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/bssopenapi"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/cloudquery/cloudquery/plugins/source/alicloud/client/services"
)

func initServices(account AccountSpec, regionID string) (*Services, error) {
	ossCli, err := oss.New(fmt.Sprintf("%s.aliyuncs.com", regionID), account.AccessKey, account.SecretKey, oss.Timeout(15, 30))
	if err != nil {
		return nil, fmt.Errorf("failed to initialize oss client: %w", err)
	}
	bssCli, err := bssopenapi.NewClientWithAccessKey(regionID, account.AccessKey, account.SecretKey)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize bssopenapi client: %w", err)
	}
	return &Services{
		OSS: ossCli,
		BSS: bssCli,
	}, nil
}

type Services struct {
	OSS services.OssClient
	BSS services.BssopenapiClient
}
