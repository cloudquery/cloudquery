package resiliencehub

import (
	"github.com/aws/aws-sdk-go-v2/service/resiliencehub/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func appVersions() *schema.Table {
	tableName := "aws_resiliencehub_app_versions"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/resilience-hub/latest/APIReference/API_AppVersionSummary.html`,
		Resolver:    fetchAppVersions,
		Transform:   transformers.TransformWithStruct(&types.AppVersionSummary{}, transformers.WithPrimaryKeys("AppVersion")),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "resiliencehub"),
		Columns:     []schema.Column{client.DefaultAccountIDColumn(false), client.DefaultRegionColumn(false), appARNTop},
		Relations:   []*schema.Table{appVersionResources(), appVersionResourceMappings()},
	}
}
