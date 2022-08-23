package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/acm/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

var ACMResources = []Resource{
	{
		DefaultColumns: []codegen.ColumnDefinition{AccountIdColumn, RegionColumn},
		//Table:              nil, // will be "generated" at "runtime"
		AWSStruct:          &types.CertificateDetail{},
		AWSService:         "ACM",
		AWSSubService:      "certificates",
		Template:           "resource_describe",
		ListFieldPrefix:    "CertificateSummary",
		ItemName:           "Certificate",
		DescribeFieldName:  "CertificateArn",
		DescribeResultName: "Certificate",
		Imports:            nil,
		MockImports:        nil,
		MockListStruct:     "",
		SkipFields:         nil,
		CreateTableOptions: schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
		ColumnOverrides: map[string]codegen.ColumnDefinition{
			"certificate_arn": {
				Name: "arn",
			},
			"key_usages": {
				Resolver: `schema.PathResolver("KeyUsages.Name")`,
			},
			"renewal_summary_renewal_status": {
				Name: "renewal_summary_status",
			},
			"renewal_summary_renewal_status_reason": {
				Name: "renewal_summary_failure_reason",
			},
			"tags": {
				Type: schema.TypeJSON,
			},
		},
	},
}
