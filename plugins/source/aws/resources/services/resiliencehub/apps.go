package resiliencehub

import (
	"github.com/aws/aws-sdk-go-v2/service/resiliencehub/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Apps() *schema.Table {
	return &schema.Table{
		Name:                "aws_resiliencehub_apps",
		Description:         `https://docs.aws.amazon.com/resilience-hub/latest/APIReference/API_App.html`,
		Resolver:            fetchApps,
		PreResourceResolver: describeApp,
		Transform:           transformers.TransformWithStruct(&types.App{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer("resiliencehub"),
		Columns:             []schema.Column{client.DefaultAccountIDColumn(false), client.DefaultRegionColumn(false), arnColumn("AppArn")},
		Relations:           []*schema.Table{appAssesments(), appVersions()},
	}
}
