package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/acm/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ACMResources() []*Resource {
	resources := []*Resource{
		{
			SubService:          "certificates",
			Struct:              &types.CertificateDetail{},
			Description:         "https://docs.aws.amazon.com/acm/latest/APIReference/API_CertificateDetail.html",
			SkipFields:          []string{"CertificateArn"},
			PreResourceResolver: "getCertificate",
			Multiplex:           `client.ServiceAccountRegionMultiplexer("acm")`,
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("CertificateArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveCertificateTags`,
					},
				}...),
		},
	}

	for _, r := range resources {
		r.Service = "acm"
	}
	return resources
}
