package client

import (
	"fmt"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func initServices(account AccountSpec, regionID string) (Services, error) {
	ossCli, err := oss.New(fmt.Sprintf("%s.aliyuncs.com", regionID), account.AccessKey, account.SecretKey)
	if err != nil {
		return Services{}, fmt.Errorf("failed to initialize oss client: %w", err)
	}
	return Services{
		OSS: ossCli,
	}
}

type Services struct {
}
