package apprunner

import (
	"github.com/aws/aws-sdk-go-v2/service/apprunner/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func CustomDomains() *schema.Table {
	return &schema.Table{
		Name:        "aws_apprunner_custom_domains",
		Description: `https://docs.aws.amazon.com/apprunner/latest/api/API_CustomDomain.html`,
		Resolver:    fetchApprunnerCustomDomains,
		Transform:   transformers.TransformWithStruct(&types.CustomDomain{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "enable_www_subdomain",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("EnableWWWSubdomain"),
			},
		},
	}
}
