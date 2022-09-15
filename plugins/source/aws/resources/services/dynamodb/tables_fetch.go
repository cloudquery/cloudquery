package dynamodb

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchDynamodbTables(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().DynamoDB

	config := dynamodb.ListTablesInput{}
	for {
		output, err := svc.ListTables(ctx, &config)
		if err != nil {
			return err
		}

		for i := range output.TableNames {
			response, err := svc.DescribeTable(ctx, &dynamodb.DescribeTableInput{TableName: &output.TableNames[i]})
			if err != nil {
				if c.IsNotFoundError(err) {
					continue
				}
				return err
			}
			res <- response.Table
		}

		if aws.ToString(output.LastEvaluatedTableName) == "" {
			break
		}
		config.ExclusiveStartTableName = output.LastEvaluatedTableName
	}

	return nil
}
func resolveDynamodbTableTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	table := resource.Item.(*types.TableDescription)

	cl := meta.(*client.Client)
	svc := cl.Services().DynamoDB
	response, err := svc.ListTagsOfResource(ctx, &dynamodb.ListTagsOfResourceInput{
		ResourceArn: table.TableArn,
	})
	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	return resource.Set(c.Name, client.TagsToMap(response.Tags))
}
func resolveDynamodbTableArchivalSummary(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*types.TableDescription)
	if r.ArchivalSummary == nil {
		return nil
	}
	return resource.Set(c.Name, map[string]interface{}{
		"date_time":  r.ArchivalSummary.ArchivalDateTime,
		"backup_arn": r.ArchivalSummary.ArchivalBackupArn,
		"reason":     r.ArchivalSummary.ArchivalReason,
	})
}
func resolveDynamodbTableAttributeDefinitions(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*types.TableDescription)
	val := make(map[string]interface{}, len(r.AttributeDefinitions))
	for i := range r.AttributeDefinitions {
		val[aws.ToString(r.AttributeDefinitions[i].AttributeName)] = map[string]interface{}{
			"type": r.AttributeDefinitions[i].AttributeType,
			"name": r.AttributeDefinitions[i].AttributeName,
		}
	}
	return resource.Set(c.Name, val)
}
func resolveDynamodbTableBillingModeSummary(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*types.TableDescription)
	if r.BillingModeSummary == nil {
		return nil
	}
	return resource.Set(c.Name, map[string]interface{}{
		"billing_mode": r.BillingModeSummary.BillingMode,
		"last_update_to_pay_per_request_date_time": r.BillingModeSummary.LastUpdateToPayPerRequestDateTime,
	})
}
func resolveDynamodbTableKeySchema(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*types.TableDescription)
	return resource.Set(c.Name, marshalKeySchema(r.KeySchema))
}
func resolveDynamodbTableRestoreSummary(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*types.TableDescription)
	if r.RestoreSummary == nil {
		return nil
	}
	return resource.Set(c.Name, map[string]interface{}{
		"date_time":         r.RestoreSummary.RestoreDateTime,
		"in_progress":       r.RestoreSummary.RestoreInProgress,
		"source_table_arn":  r.RestoreSummary.SourceTableArn,
		"source_backup_arn": r.RestoreSummary.SourceBackupArn,
	})
}
func resolveDynamodbTableStreamSpecification(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*types.TableDescription)
	if r.StreamSpecification == nil {
		return nil
	}
	return resource.Set(c.Name, map[string]interface{}{
		"enabled":   r.StreamSpecification.StreamEnabled,
		"view_type": r.StreamSpecification.StreamViewType,
	})
}
func fetchDynamodbTableReplicaAutoScalings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	par := parent.Item.(*types.TableDescription)

	if aws.ToString(par.GlobalTableVersion) == "" {
		// "This operation only applies to Version 2019.11.21 of global tables"
		return nil
	}

	c := meta.(*client.Client)
	svc := c.Services().DynamoDB

	output, err := svc.DescribeTableReplicaAutoScaling(ctx, &dynamodb.DescribeTableReplicaAutoScalingInput{
		TableName: par.TableName,
	})
	if err != nil {
		if c.IsNotFoundError(err) {
			return nil
		}
		return err
	}

	for i := range output.TableAutoScalingDescription.Replicas {
		res <- output.TableAutoScalingDescription.Replicas[i]
	}
	return nil
}
func fetchDynamodbTableContinuousBackups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	par := parent.Item.(*types.TableDescription)

	c := meta.(*client.Client)
	svc := c.Services().DynamoDB

	output, err := svc.DescribeContinuousBackups(ctx, &dynamodb.DescribeContinuousBackupsInput{
		TableName: par.TableName,
	})
	if err != nil {
		if c.IsNotFoundError(err) {
			return nil
		}
		return err
	}

	res <- output.ContinuousBackupsDescription
	return nil
}
func marshalKeySchema(k []types.KeySchemaElement) []byte {
	if len(k) == 0 {
		return nil
	}
	val := make([]map[string]interface{}, len(k))
	for i := range k {
		val[i] = map[string]interface{}{
			"type": k[i].KeyType,
			"name": k[i].AttributeName,
		}
	}
	b, _ := json.Marshal(val)
	return b
}
