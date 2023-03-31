package athena

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
)

func createDataCatalogArn(cl *client.Client, catalogName string) string {
	return arn.ARN{
		Partition: cl.Partition,
		Service:   string(client.Athena),
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  fmt.Sprintf("datacatalog/%s", catalogName),
	}.String()
}
