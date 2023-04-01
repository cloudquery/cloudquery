package glue

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
)

func devEndpointARN(cl *client.Client, name string) string {
	return arn.ARN{
		Partition: cl.Partition,
		Service:   string(client.GlueService),
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  fmt.Sprintf("devEndpoint/%s", name),
	}.String()
}
