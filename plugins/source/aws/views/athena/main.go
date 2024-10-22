package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/athena"
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
)

type UpdateResourcesViewEvent struct {
	Catalog      string   `json:"catalog"`
	Database     string   `json:"database"`
	Output       string   `json:"output"`
	View         string   `json:"view"`
	Region       string   `json:"region"`
	ExtraColumns []string `json:"extra_columns"`
}

func HandleRequest(ctx context.Context, event UpdateResourcesViewEvent) (string, error) {
	log.Println("Setting up...")

	awsCfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(event.Region))
	if err != nil {
		return "Error loading config", err
	}

	// Create an Athena client
	svc := athena.NewFromConfig(awsCfg)

	// Set the query string
	queryString := `WITH tables AS (
SELECT table_name
    FROM information_schema.columns
    WHERE table_name LIKE 'aws_%s' and COLUMN_NAME = 'account_id' and table_catalog = '${CATALOG}' and table_schema = '${DATABASE}'
    INTERSECT
    SELECT table_name
    FROM information_schema.columns
    WHERE table_name LIKE 'aws_%s' and COLUMN_NAME = 'arn' and table_catalog = '${CATALOG}' and table_schema = '${DATABASE}'
	AND table_name NOT IN (
	SELECT table_name
	FROM information_schema.views
	WHERE table_catalog = '${CATALOG}' and table_schema = '${DATABASE}')
) 
SELECT 
    t.table_name, 
    (CASE WHEN EXISTS (SELECT 1 FROM information_schema.columns cols WHERE column_name='region' AND cols.table_name=t.table_name and cols.table_catalog = '${CATALOG}' and cols.table_schema = '${DATABASE}') THEN true ELSE false END) AS has_region,
    (CASE WHEN EXISTS (SELECT 1 FROM information_schema.columns cols WHERE column_name='tags' AND cols.table_name=t.table_name and cols.table_catalog = '${CATALOG}' and cols.table_schema = '${DATABASE}') THEN true ELSE false END) AS has_tags,
	(CASE WHEN EXISTS (SELECT 1 FROM information_schema.columns cols WHERE column_name='tags' AND cols.table_name=t.table_name and cols.table_catalog = '${CATALOG}' and cols.table_schema = '${DATABASE}') THEN (SELECT data_type FROM information_schema.columns cols WHERE column_name='tags' AND cols.table_name=t.table_name and cols.table_catalog = '${CATALOG}' and cols.table_schema = '${DATABASE}') ELSE '' END) AS tags_data_type
FROM tables t;`

	// Replace the catalog and database placeholders with the actual values
	queryString = strings.ReplaceAll(queryString, "${CATALOG}", event.Catalog)
	queryString = strings.ReplaceAll(queryString, "${DATABASE}", event.Database)

	// Set up the query input parameters
	input := &athena.StartQueryExecutionInput{
		QueryString: aws.String(queryString),
		QueryExecutionContext: &types.QueryExecutionContext{
			Database: aws.String(event.Database),
			Catalog:  aws.String(event.Catalog),
		},
		ResultConfiguration: &types.ResultConfiguration{
			OutputLocation: aws.String(event.Output),
		},
	}

	log.Println("Starting query to get tables...")
	// Start the query execution
	result, err := svc.StartQueryExecution(ctx, input)
	if err != nil {
		fmt.Println("Error starting query execution:", err)
		return "Error starting query execution", err
	}

	// Get the query execution ID
	queryExecutionID := *result.QueryExecutionId

	waitForResults(ctx, svc, queryExecutionID)

	log.Println("Reading query results...")
	// Get the query results
	getQueryResultsOutput, err := svc.GetQueryResults(ctx, &athena.GetQueryResultsInput{
		QueryExecutionId: aws.String(queryExecutionID),
	})
	if err != nil {
		fmt.Println("Error getting query results:", err)
		return "Error getting query results", err
	}

	type table struct {
		name         string
		hasRegion    bool
		hasTags      bool
		tagsDataType string
	}

	// Create a slice to store the results
	var tables []table

	// Loop through the rows of the query results
	for i, row := range getQueryResultsOutput.ResultSet.Rows {
		if i == 0 {
			// skip the header
			continue
		}
		// Get the first column value from the row
		tables = append(tables, table{
			name:         *row.Data[0].VarCharValue,
			hasRegion:    *row.Data[1].VarCharValue == "true",
			hasTags:      *row.Data[2].VarCharValue == "true",
			tagsDataType: *row.Data[3].VarCharValue,
		})
	}
	tableNames := make([]string, len(tables))
	for i, t := range tables {
		tableNames[i] = t.name
	}
	log.Println("Found", len(tables), "matching tables:", tableNames)
	if len(tables) == 0 {
		return "No tables found", errors.New("no matching tables found")
	}
	log.Println("Query results read, creating or updating view...")

	// Create the view
	var sb strings.Builder
	sb.WriteString(`CREATE OR REPLACE VIEW aws_resources AS (`)
	sb.WriteString("\n")
	for i, t := range tables {
		if i > 0 {
			sb.WriteString("  UNION ALL\n")
		}
		region := "region"
		if !t.hasRegion {
			region = "''"
		}
		tags := `tags`
		if !t.hasTags || t.tagsDataType != "varchar" {
			// we can't convert the tags to a string if they aren't varchar already,
			// so just use an empty map
			tags = "'{}'"
		}
		q := `    SELECT _cq_id, _cq_source_name, cast (_cq_sync_time as timestamp) as _cq_sync_time, '%s' as _cq_table, account_id, %s as region, arn, %s as tags`
		if len(event.ExtraColumns) > 0 {
			for _, c := range event.ExtraColumns {
				q += ", "
				q += c
			}
		}
		sb.WriteString(fmt.Sprintf(q, t.name, region, tags))
		sb.WriteString(` FROM ` + t.name + "\n")
		sb.WriteString("\n")
	}
	sb.WriteString(`)`)
	fmt.Println(sb.String())

	// Set up the query input parameters
	input = &athena.StartQueryExecutionInput{
		QueryString: aws.String(sb.String()),
		QueryExecutionContext: &types.QueryExecutionContext{
			Database: aws.String(event.Database),
			Catalog:  aws.String(event.Catalog),
		},
		ResultConfiguration: &types.ResultConfiguration{
			OutputLocation: aws.String(event.Output),
		},
	}

	result, err = svc.StartQueryExecution(ctx, input)
	if err != nil {
		fmt.Println("Error starting query execution:", err)
		return "Error starting query execution", err
	}
	err = waitForResults(ctx, svc, *result.QueryExecutionId)
	if err != nil {
		return "Error while creating or updating view", err
	}

	log.Println("Success!")
	return "", nil
}

func waitForResults(ctx context.Context, svc *athena.Client, queryExecutionID string) error {
	log.Println("Waiting for query results...")

	// Check the query execution status until it's complete
	for {
		queryExecutionOutput, err := svc.GetQueryExecution(ctx, &athena.GetQueryExecutionInput{
			QueryExecutionId: aws.String(queryExecutionID),
		})

		if err != nil {
			fmt.Println("Error getting query execution:", err)
			return err
		}

		log.Println("Still waiting for query results...")
		queryExecution := queryExecutionOutput.QueryExecution
		switch queryExecution.Status.State {
		case types.QueryExecutionStateSucceeded:
			return nil
		case types.QueryExecutionStateFailed:
			return errors.New("query failed. Check the query results for more information")
		case types.QueryExecutionStateCancelled:
			return errors.New("query cancelled")
		}
		time.Sleep(3 * time.Second)
	}
}

func main() {
	if len(os.Args) == 1 {
		log.Println("Running as Lambda function (see --help for local usage)")
		lambda.Start(HandleRequest)
		return
	}

	e := UpdateResourcesViewEvent{}
	var extraCols string
	flag.StringVar(&e.Database, "database", "", "Database name")
	flag.StringVar(&e.Output, "output", "", "Query output path (e.g. s3://bucket/path)")
	flag.StringVar(&e.Catalog, "catalog", "awsdatacatalog", "Catalog name")
	flag.StringVar(&e.View, "view", "aws_resources", "View name (default: aws_resources)")
	flag.StringVar(&e.Region, "region", "us-east-1", "View name (default: us-east-1)")
	flag.StringVar(&extraCols, "extra-columns", "", "Extra columns to add to the view, separated by commas (default: none)")
	flag.Parse()

	if extraCols != "" {
		e.ExtraColumns = strings.Split(extraCols, ",")
	}

	if e.Database == "" {
		log.Fatal("database name is required")
	}
	if e.Output == "" {
		log.Fatal("S3 output path is required")
	}

	_, err := HandleRequest(context.Background(), e)
	if err != nil {
		log.Fatal(err)
	}
}
