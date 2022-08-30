package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/acm/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

var ACMResources = []*Resource{
	{
		DefaultColumns:           []codegen.ColumnDefinition{AccountIdColumn, RegionColumn},
		AWSStruct:                &types.CertificateDetail{},
		AWSService:               "ACM",
		AWSSubService:            "Certificates",
		ItemName:                 "Certificate",
		ListFieldName:            "CertificateArn",
		PaginatorListName:        "CertificateSummaryList",
		MockRawPaginatorListType: "types.CertificateSummary",
		MockRawListDetailType:    "types.CertificateDetail",
		Template:                 "resource_list_describe",
		CreateTableOptions:       schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
		ColumnOverrides: map[string]codegen.ColumnDefinition{
			"certificate_arn": {
				Name: "arn",
			},
			"key_usages": {
				Resolver: `schema.PathResolver("KeyUsages.Name")`,
			},
			"tags": {
				Type:        schema.TypeJSON,
				Description: "The tags that have been applied to the ACM certificate",
			},
		},
	},
}
