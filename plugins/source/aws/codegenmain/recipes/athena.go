package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/athena"
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func init() {
	add(combine(
		parentize(&Resource{
			DefaultColumns:     []codegen.ColumnDefinition{AccountIdColumn, RegionColumn},
			AWSStruct:          &types.DataCatalog{},
			AWSService:         "Athena",
			Template:           "resource_list_and_detail",
			PaginatorStruct:    &athena.ListDataCatalogsOutput{},
			PaginatorGetStruct: &athena.GetDataCatalogInput{},
			ItemsStruct:        &athena.GetDataCatalogOutput{},
			//CreateTableOptions: schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
			CustomErrorBlock: `
		// retrieving of default data catalog (AwsDataCatalog) returns "not found error" but it exists and its
		// relations can be fetched by its name
		if *item.CatalogName == "AwsDataCatalog" {
			resultsChan <- types.DataCatalog{Name: item.CatalogName, Type: item.Type}
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
				AWSStruct:       &types.Database{},
				Template:        "resource_get",
				ChildFieldName:  "CatalogName",
				ParentFieldName: "Name",
				ItemsStruct:     &athena.ListDatabasesOutput{},
			},
				&Resource{
					AWSStruct:            &types.TableMetadata{},
					CQSubserviceOverride: "tables",
					Template:             "resource_get",
					ChildFieldName:       "DatabaseName",
					ParentFieldName:      "Name",
					ItemsStruct:          &athena.ListTableMetadataOutput{},
				},
			)...,
		),
		parentize(&Resource{
			DefaultColumns:     []codegen.ColumnDefinition{AccountIdColumn, RegionColumn},
			AWSStruct:          &types.WorkGroup{},
			AWSService:         "Athena",
			Template:           "resource_list_and_detail",
			PaginatorStruct:    &athena.ListWorkGroupsOutput{},
			PaginatorGetStruct: &athena.GetWorkGroupInput{},
			ItemsStruct:        &athena.GetWorkGroupOutput{},
			ItemName:           "WorkGroup",
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
				AWSStruct:          &types.PreparedStatement{},
				Template:           "resource_list_describe",
				ItemName:           "PreparedStatement",
				PaginatorStruct:    &athena.ListPreparedStatementsOutput{},
				PaginatorGetStruct: &athena.GetPreparedStatementInput{},
				ItemsStruct:        &athena.GetPreparedStatementOutput{},
				ChildFieldName:     "WorkGroup",
				ParentFieldName:    "Name",
			},
			&Resource{
				AWSStruct:                &types.QueryExecution{},
				Template:                 "resource_list_describe",
				ItemName:                 "QueryExecution",
				PaginatorStruct:          &athena.ListQueryExecutionsOutput{},
				PaginatorGetStruct:       &athena.GetQueryExecutionInput{},
				ItemsStruct:              &athena.GetQueryExecutionOutput{},
				ChildFieldName:           "WorkGroup",
				ParentFieldName:          "Name",
				SkipDescribeParentInputs: true,
			},
			&Resource{
				AWSStruct:                &types.NamedQuery{},
				Template:                 "resource_list_describe",
				ItemName:                 "NamedQuery",
				PaginatorStruct:          &athena.ListNamedQueriesOutput{},
				PaginatorGetStruct:       &athena.GetNamedQueryInput{},
				ItemsStruct:              &athena.GetNamedQueryOutput{},
				ChildFieldName:           "WorkGroup",
				ParentFieldName:          "Name",
				SkipDescribeParentInputs: true,
			},
		),
	)...)
}
