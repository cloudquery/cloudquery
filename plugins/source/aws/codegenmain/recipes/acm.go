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
		ItemName:           "Certificate",
		DescribeFieldName:  "CertificateArn",
		Template:           "resource_describe",
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
			"tags": {
				Type:        schema.TypeJSON,
				Description: "The tags that have been applied to the ACM certificate",
			},
		},
	},
}
