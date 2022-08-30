package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

var AthenaResources = combine(
	parentize(&Resource{
		DefaultColumns:       []codegen.ColumnDefinition{AccountIdColumn, RegionColumn},
		AWSStruct:            &types.DataCatalog{},
		AWSService:           "Athena",
		AWSSubService:        "DataCatalogs",
		Template:             "resource_list_and_detail",
		ListVerb:             "List",
		ListFieldName:        "DataCatalogsSummary",
		ResponseItemsName:    "CatalogName",
		ItemName:             "DataCatalog",
		DetailInputFieldName: "Name",
		ResponseItemsType:    "DataCatalogSummary",
		//CreateTableOptions: schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
		CustomErrorBlock: `
		// retrieving of default data catalog (AwsDataCatalog) returns "not found error" but it exists and its
		// relations can be fetched by its name
		if *itemSummary.CatalogName == "AwsDataCatalog" {
			resultsChan <- types.DataCatalog{Name: itemSummary.CatalogName, Type: itemSummary.Type}
			return
		}
`,
		CustomTagField: `aws.String(resolvers.CreateDataCatalogArn(cl, *item.CatalogName))`,
		ColumnOverrides: map[string]codegen.ColumnDefinition{
			"arn": {
				Type:     schema.TypeString,
				Resolver: "resolvers.ResolveDataCatalogArn",
			},
			"tags": {
				Type:        schema.TypeJSON,
				Description: "Tags associated with the Athena data catalog.",
				Resolver:    ResolverAuto,
			},
		},
	},
		parentize(&Resource{
			AWSStruct:         &types.Database{},
			AWSSubService:     "Databases",
			Template:          "resource_get",
			ChildFieldName:    "CatalogName",
			ParentFieldName:   "Name",
			Verb:              "List",
			ResponseItemsName: "DatabaseList",
		},
			&Resource{
				AWSStruct:            &types.TableMetadata{},
				AWSSubService:        "TableMetadata",
				CQSubserviceOverride: "tables",
				Template:             "resource_get",
				ChildFieldName:       "DatabaseName",
				ParentFieldName:      "Name",
				Verb:                 "List",
				ResponseItemsName:    "TableMetadataList",
			},
		)...,
	),
	parentize(&Resource{
		DefaultColumns:       []codegen.ColumnDefinition{AccountIdColumn, RegionColumn},
		AWSStruct:            &types.WorkGroup{},
		AWSService:           "Athena",
		AWSSubService:        "WorkGroups",
		Template:             "resource_list_and_detail",
		ListVerb:             "List",
		ListFieldName:        "WorkGroups",
		ResponseItemsName:    "Name",
		ItemName:             "WorkGroup",
		DetailInputFieldName: "WorkGroup",
		ResponseItemsType:    "WorkGroupSummary",
		//CreateTableOptions: schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
		CustomTagField: `aws.String(resolvers.CreateWorkGroupArn(cl, *item.Name))`,
		ColumnOverrides: map[string]codegen.ColumnDefinition{
			"arn": {
				Type:     schema.TypeString,
				Resolver: "resolvers.ResolveWorkGroupArn",
			},
			"tags": {
				Type:        schema.TypeJSON,
				Description: "Tags associated with the Athena work group.",
				Resolver:    ResolverAuto,
			},
		},
	},
		&Resource{
			AWSStruct:                &types.PreparedStatement{},
			AWSSubService:            "PreparedStatements",
			Template:                 "resource_list_describe",
			ItemName:                 "PreparedStatement",
			ListFieldName:            "StatementName",
			PaginatorListName:        "PreparedStatements",
			ChildFieldName:           "WorkGroup",
			ParentFieldName:          "Name",
			Verb:                     "Get",
			ResponseItemsName:        "PreparedStatement",
			MockRawPaginatorListType: "types.PreparedStatementSummary",
			MockRawListDetailType:    "types.PreparedStatement",
		},
		&Resource{
			AWSStruct:                &types.QueryExecution{},
			AWSSubService:            "QueryExecutions",
			Template:                 "resource_list_describe",
			ItemName:                 "QueryExecution",
			ListFieldName:            "QueryExecutionId",
			PaginatorListName:        "QueryExecutionIds",
			ChildFieldName:           "WorkGroup",
			ParentFieldName:          "Name",
			Verb:                     "Get",
			ResponseItemsName:        "QueryExecution",
			SkipDescribeParentInputs: true,
			RawDescribeFieldValue:    `&item`,
			MockRawPaginatorListType: "string",
			MockRawListDetailType:    "types.QueryExecution",
		},
		&Resource{
			AWSStruct:                &types.NamedQuery{},
			AWSSubService:            "NamedQueries",
			Template:                 "resource_list_describe",
			ItemName:                 "NamedQuery",
			ListFieldName:            "NamedQueryId",
			PaginatorListName:        "NamedQueryIds",
			ChildFieldName:           "WorkGroup",
			ParentFieldName:          "Name",
			Verb:                     "Get",
			ResponseItemsName:        "NamedQuery",
			SkipDescribeParentInputs: true,
			RawDescribeFieldValue:    `&item`,
			MockRawPaginatorListType: "string",
			MockRawListDetailType:    "types.NamedQuery",
		},
	),
)
