package client

import (
	"context"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
)

type Client struct {
	logger         *zerolog.Logger
	Spec           Spec
	OSSClient      *oss.Client
	ossClientCache map[string]*oss.Client
	Client         *sdk.Client
}

func (c *Client) Logger() *zerolog.Logger {
	return c.logger
}

func (*Client) ID() string {
	return "alicloud-client"
}

func (c *Client) GetOSSClient(location string) (*oss.Client, error) {
	client, ok := c.ossClientCache[location]
	if ok {
		return client, nil
	}
	endpoint := fmt.Sprintf("%s.aliyuncs.com", location)
	ossCli, err := oss.New(endpoint, c.Spec.AccessKey, c.Spec.SecretKey)
	if err != nil {
		return nil, err
	}
	c.ossClientCache[location] = ossCli
	return ossCli, nil
}

func New(_ context.Context, logger zerolog.Logger, s specs.Source) (schema.ClientMeta, error) {
	var spec Spec
	err := s.UnmarshalSpec(&spec)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal GitHub spec: %w", err)
	}
	// validate plugin config
	if spec.RegionID == "" {
		return nil, fmt.Errorf("missing alicloud region id in configuration")
	}
	if spec.AccessKey == "" {
		return nil, fmt.Errorf("missing alicloud access key in configuration")
	}
	if spec.SecretKey == "" {
		return nil, fmt.Errorf("missing alicloud secret key in configuration")
	}
	endpoint := fmt.Sprintf("oss-%s.aliyuncs.com", spec.RegionID)
	ossCli, err := oss.New(endpoint, spec.AccessKey, spec.SecretKey, oss.Timeout(15, 30))
	if err != nil {
		return nil, err
	}
	client, err := sdk.NewClientWithAccessKey(spec.RegionID, spec.AccessKey, spec.SecretKey)
	if err != nil {
		return nil, err
	}
	ossClientCache := make(map[string]*oss.Client)
	ossClientCache["oss-"+spec.RegionID] = ossCli

	return &Client{logger: &logger, Spec: spec, OSSClient: ossCli, ossClientCache: ossClientCache, Client: client}, nil
}
