package resources

import (
	"context"
	"encoding/json"
	"fmt"

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
func resolveLambdaFunctionEventSourceMappingAccessConfigurations(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(types.EventSourceMappingConfiguration)
	if !ok {
		return fmt.Errorf("wrong type assertion: got %T instead of EventSourceMappingConfiguration", p)
	}
	data, err := json.Marshal(p.SourceAccessConfigurations)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, data)
}

type RuntimeWrapper struct {
	Name string
}
