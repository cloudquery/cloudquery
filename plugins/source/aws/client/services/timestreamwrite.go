// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/timestreamwrite"
	
)

//go:generate mockgen -package=mocks -destination=../mocks/timestreamwrite.go -source=timestreamwrite.go TimestreamwriteClient
type TimestreamwriteClient interface {
	DescribeDatabase(context.Context, *timestreamwrite.DescribeDatabaseInput, ...func(*timestreamwrite.Options)) (*timestreamwrite.DescribeDatabaseOutput, error)
	DescribeEndpoints(context.Context, *timestreamwrite.DescribeEndpointsInput, ...func(*timestreamwrite.Options)) (*timestreamwrite.DescribeEndpointsOutput, error)
	DescribeTable(context.Context, *timestreamwrite.DescribeTableInput, ...func(*timestreamwrite.Options)) (*timestreamwrite.DescribeTableOutput, error)
	ListDatabases(context.Context, *timestreamwrite.ListDatabasesInput, ...func(*timestreamwrite.Options)) (*timestreamwrite.ListDatabasesOutput, error)
	ListTables(context.Context, *timestreamwrite.ListTablesInput, ...func(*timestreamwrite.Options)) (*timestreamwrite.ListTablesOutput, error)
	ListTagsForResource(context.Context, *timestreamwrite.ListTagsForResourceInput, ...func(*timestreamwrite.Options)) (*timestreamwrite.ListTagsForResourceOutput, error)
}
