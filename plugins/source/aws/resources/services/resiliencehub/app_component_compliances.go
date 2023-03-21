package resiliencehub

import (
	"github.com/aws/aws-sdk-go-v2/service/resiliencehub/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func appComponentCompliances() *schema.Table {
	tableName := "aws_resiliencehub_app_component_compliances"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/resilience-hub/latest/APIReference/API_AppComponentCompliance.html`,
		Resolver:    fetchAppComponentCompliances,
		Transform:   transformers.TransformWithStruct(&types.AppComponentCompliance{}, transformers.WithPrimaryKeys("AppComponentName")),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "resiliencehub"),
		Columns:     []schema.Column{client.DefaultAccountIDColumn(false), client.DefaultRegionColumn(false), appARN, assessmentARN},
	}
}
