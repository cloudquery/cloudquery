package resources

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func LambdaRuntimes() *schema.Table {
	return &schema.Table{
		Name:        "aws_lambda_runtimes",
		Description: "All known values for Runtime",
		Resolver:    fetchLambdaRuntimes,
		Options:     schema.TableCreationOptions{PrimaryKeys: []string{"name"}},
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
func fetchLambdaRuntimes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	for _, runtime := range types.RuntimeProvidedal2.Values() {
		res <- &RuntimeWrapper{
			Name: string(runtime),
		}
	}

	return nil
}

type RuntimeWrapper struct {
	Name string
}
