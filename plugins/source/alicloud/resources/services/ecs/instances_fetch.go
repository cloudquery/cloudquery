package ecs

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/cloudquery/cloudquery/plugins/source/alicloud/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

var (
	maxLimit = 100
)

func fetchEcsInstances(_ context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	pageNum := 1
	total := 0
	req := ecs.CreateDescribeInstancesRequest()
	req.PageNumber = requests.NewInteger(pageNum)
	req.PageSize = requests.NewInteger(maxLimit)
	for {
		resp, err := c.Services().ECS.DescribeInstances(req)
		if err != nil {
			return err
		}
		if !resp.IsSuccess() {
			code := resp.GetHttpStatus()
			return fmt.Errorf("got response status code %d (%v)", code, http.StatusText(code))
		}
		for _, instance := range resp.Instances.Instance {
			res <- instance
		}
		total += len(resp.Instances.Instance)
		if len(resp.Instances.Instance) == 0 || total >= resp.TotalCount {
			break
		}
		pageNum++
		req.PageNumber = requests.NewInteger(pageNum)
	}
	return nil
}
