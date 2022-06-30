package lambda

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

type RuntimeWrapper struct {
	Name string
}

func LambdaRuntimes() *schema.Table {
	return &schema.Table{
		Name:         "aws_lambda_runtimes",
		Description:  "All known values for Runtime",
		Resolver:     fetchLambdaRuntimes,
		DeleteFilter: client.DeleteAllFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"name"}},
		Columns: []schema.Column{
			{
				Name:        "name",
				Description: "Runtime name",
				Type:        schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchLambdaRuntimes(_ context.Context, _ schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	runtimes := make([]RuntimeWrapper, len(types.RuntimeProvidedal2.Values()))
	for i, runtime := range types.RuntimeProvidedal2.Values() {
		runtimes[i] = RuntimeWrapper{
			Name: string(runtime),
		}
	}
	res <- runtimes
	return nil
}
