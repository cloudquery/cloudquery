package ssoadmin

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchSsoadminInstances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	fmt.Print("*******FETCH********\n")
	svc := meta.(*client.Client).Services().SSOAdmin
	config := ssoadmin.ListInstancesInput{}
	fmt.Print(config)
	for {
		response, err := svc.ListInstances(ctx, &config)
		fmt.Print("*******RESPONSE********\n")
		fmt.Printf("%#v\n", response)
		if err != nil {
			return err
		}
		for _, i := range response.Instances {
			fmt.Print(i)
			res <- i
		}
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}
