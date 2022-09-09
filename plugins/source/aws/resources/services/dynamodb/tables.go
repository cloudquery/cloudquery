package dynamodb

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/plugin-sdk/schema"
)

type simplifiedGlobalSecondaryIndex struct {
	Name          *string                `json:"name"`
	Status        types.IndexStatus      `json:"status"`
	ReadCapacity  map[string]interface{} `json:"read_capacity"`
	WriteCapacity map[string]interface{} `json:"write_capacity"`
}

func DynamodbTables() *schema.Table {
	return &schema.Table{
		Name:          "aws_dynamodb_tables",
		Description:   "Information about a DynamoDB table.",
		Resolver:      fetchDynamodbTables,
		Multiplex:     client.ServiceAccountRegionMultiplexer("dynamodb"),
		IgnoreInTests: true,
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:        "tags",
				Description: "The tags associated with the table.",
				Type:        schema.TypeJSON,
				Resolver:    resolveDynamodbTableTags,
			},
			{
				Name:        "archival_summary",
				Description: "Contains information about the table archive.",
				Type:        schema.TypeJSON,
				Resolver:    resolveDynamodbTableArchivalSummary,
			},
			{
				Name:        "attribute_definitions",
				Description: "An array of AttributeDefinition objects",
				Type:        schema.TypeJSON,
				Resolver:    resolveDynamodbTableAttributeDefinitions,
			},
			{
				Name:        "billing_mode_summary",
				Description: "Contains the details for the read/write capacity mode.",
				Type:        schema.TypeJSON,
				Resolver:    resolveDynamodbTableBillingModeSummary,
			},
			{
				Name:        "creation_date_time",
				Description: "The date and time when the table was created, in UNIX epoch time (http://www.epochconverter.com/) format.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "global_table_version",
				Description: "Represents the version of global tables (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/GlobalTables.html) in use, if the table is replicated across Amazon Web Services Regions.",
				Type:        schema.TypeString,
			},
			{
				Name:        "item_count",
				Description: "The number of items in the specified table",
				Type:        schema.TypeInt,
			},
			{
				Name:        "key_schema",
				Description: "The primary key structure for the table",
				Type:        schema.TypeJSON,
				Resolver:    resolveDynamodbTableKeySchema,
			},
			{
				Name:        "latest_stream_arn",
				Description: "The Amazon Resource Name (ARN) that uniquely identifies the latest stream for this table.",
				Type:        schema.TypeString,
			},
			{
				Name:        "latest_stream_label",
				Description: "A timestamp, in ISO 8601 format, for this stream",
				Type:        schema.TypeString,
			},
			{
				Name:     "provisioned_throughput",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ProvisionedThroughput"),
			},
			{
				Name:        "restore_summary",
				Description: "Contains details for the restore.",
				Type:        schema.TypeJSON,
				Resolver:    resolveDynamodbTableRestoreSummary,
			},
			{
				Name:        "inaccessible_encryption_date_time",
				Description: "Indicates the time, in UNIX epoch date format, when DynamoDB detected that the table's KMS key was inaccessible",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("SSEDescription.InaccessibleEncryptionDateTime"),
			},
			{
				Name:     "sse_description",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SSEDescription"),
			},
			{
				Name:        "stream_specification",
				Description: "The current DynamoDB Streams configuration for the table.",
				Type:        schema.TypeJSON,
				Resolver:    resolveDynamodbTableStreamSpecification,
			},
			{
				Name:            "arn",
				Description:     "The Amazon Resource Name (ARN) that uniquely identifies the table.",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("TableArn"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:     "table_class_summary",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("TableClassSummary"),
			},
			{
				Name:        "id",
				Description: "Unique identifier for the table for which the backup was created.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("TableId"),
			},
			{
				Name:        "name",
				Description: "The name of the table.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("TableName"),
			},
			{
				Name:        "size_bytes",
				Description: "The total size of the specified table, in bytes",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("TableSizeBytes"),
			},
			{
				Name:        "status",
				Description: "The current state of the table:  * CREATING - The table is being created.  * UPDATING - The table is being updated.  * DELETING - The table is being deleted.  * ACTIVE - The table is ready for use.  * INACCESSIBLE_ENCRYPTION_CREDENTIALS - The KMS key used to encrypt the table in inaccessible",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("TableStatus"),
			},
			{
				Name:        "global_secondary_indexes",
				Description: "Represents the properties of a global secondary index.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("GlobalSecondaryIndexes"),
			},
			{
				Name:        "local_secondary_indexes",
				Description: "Represents the properties of a local secondary index.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("LocalSecondaryIndexes"),
			},
			{
				Name:        "replicas",
				Description: "Contains the details of replicas.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Replicas"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_dynamodb_table_replica_auto_scalings",
				Description: "Represents the auto scaling settings of the replica.",
				Resolver:    fetchDynamodbTableReplicaAutoScalings,
				Columns: []schema.Column{
					{
						Name:        "table_cq_id",
						Description: "Unique CloudQuery ID of aws_dynamodb_tables table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "global_secondary_indexes",
						Description: "Replica-specific global secondary index auto scaling settings.",
						Type:        schema.TypeJSON,
						Resolver:    resolveDynamodbTableReplicaAutoScalingGlobalSecondaryIndexes,
					},
					{
						Name:        "region_name",
						Description: "The Region where the replica exists.",
						Type:        schema.TypeString,
					},
					{
						Name:        "read_capacity",
						Description: "Represents the auto scaling settings for a global table or global secondary index.",
						Type:        schema.TypeJSON,
						Resolver:    resolveDynamodbTableReplicaAutoScalingReadCapacity,
					},
					{
						Name:        "write_capacity",
						Description: "Represents the auto scaling settings for a global table or global secondary index.",
						Type:        schema.TypeJSON,
						Resolver:    resolveDynamodbTableReplicaAutoScalingWriteCapacity,
					},
					{
						Name:        "replica_status",
						Description: "The current state of the replica:  * CREATING - The replica is being created.  * UPDATING - The replica is being updated.  * DELETING - The replica is being deleted.  * ACTIVE - The replica is ready for use.",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "aws_dynamodb_table_continuous_backups",
				Description: "Represents the continuous backups and point in time recovery settings on the table.",
				Resolver:    fetchDynamodbTableContinuousBackups,
				Columns: []schema.Column{
					{
						Name:        "table_cq_id",
						Description: "Unique CloudQuery ID of aws_dynamodb_tables table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "continuous_backups_status",
						Description: "ContinuousBackupsStatus can be one of the following states: ENABLED, DISABLED ",
						Type:        schema.TypeString,
					},
					{
						Name:     "point_in_time_recovery_description",
						Type:     schema.TypeJSON,
						Resolver: schema.PathResolver("PointInTimeRecoveryDescription"),
					},
					{
						Name:        "point_in_time_recovery_status",
						Description: "The current state of point in time recovery:  * ENABLING - Point in time recovery is being enabled.  * ENABLED - Point in time recovery is enabled.  * DISABLED - Point in time recovery is disabled.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("PointInTimeRecoveryDescription.PointInTimeRecoveryStatus"),
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//
//	Table Resolver Functions
//
// ====================================================================================================================
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
	return diag.WrapError(resource.Set(c.Name, map[string]interface{}{
		"date_time":  r.ArchivalSummary.ArchivalDateTime,
		"backup_arn": r.ArchivalSummary.ArchivalBackupArn,
		"reason":     r.ArchivalSummary.ArchivalReason,
	}))
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
	return diag.WrapError(resource.Set(c.Name, map[string]interface{}{
		"billing_mode": r.BillingModeSummary.BillingMode,
		"last_update_to_pay_per_request_date_time": r.BillingModeSummary.LastUpdateToPayPerRequestDateTime,
	}))
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
	return diag.WrapError(resource.Set(c.Name, map[string]interface{}{
		"date_time":         r.RestoreSummary.RestoreDateTime,
		"in_progress":       r.RestoreSummary.RestoreInProgress,
		"source_table_arn":  r.RestoreSummary.SourceTableArn,
		"source_backup_arn": r.RestoreSummary.SourceBackupArn,
	}))
}
func resolveDynamodbTableStreamSpecification(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*types.TableDescription)
	if r.StreamSpecification == nil {
		return nil
	}
	return diag.WrapError(resource.Set(c.Name, map[string]interface{}{
		"enabled":   r.StreamSpecification.StreamEnabled,
		"view_type": r.StreamSpecification.StreamViewType,
	}))
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
func resolveDynamodbTableReplicaAutoScalingGlobalSecondaryIndexes(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.ReplicaAutoScalingDescription)
	if len(r.GlobalSecondaryIndexes) == 0 {
		return nil
	}

	val := make(map[string]simplifiedGlobalSecondaryIndex, len(r.GlobalSecondaryIndexes))
	for i := range r.GlobalSecondaryIndexes {
		d := r.GlobalSecondaryIndexes[i]
		val[aws.ToString(d.IndexName)] = simplifiedGlobalSecondaryIndex{
			Name:          d.IndexName,
			Status:        d.IndexStatus,
			ReadCapacity:  marshalAutoScalingSettingsDescription(d.ProvisionedReadCapacityAutoScalingSettings),
			WriteCapacity: marshalAutoScalingSettingsDescription(d.ProvisionedWriteCapacityAutoScalingSettings),
		}
	}

	return resource.Set(c.Name, val)
}
func resolveDynamodbTableReplicaAutoScalingReadCapacity(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.ReplicaAutoScalingDescription)
	if r.ReplicaProvisionedReadCapacityAutoScalingSettings == nil {
		return nil
	}
	return resource.Set(c.Name, marshalAutoScalingSettingsDescription(r.ReplicaProvisionedReadCapacityAutoScalingSettings))
}
func resolveDynamodbTableReplicaAutoScalingWriteCapacity(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.ReplicaAutoScalingDescription)
	if r.ReplicaProvisionedWriteCapacityAutoScalingSettings == nil {
		return nil
	}
	return resource.Set(c.Name, marshalAutoScalingSettingsDescription(r.ReplicaProvisionedWriteCapacityAutoScalingSettings))
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

// ====================================================================================================================
//                                                  User Defined Helpers
// ====================================================================================================================

func marshalAutoScalingSettingsDescription(d *types.AutoScalingSettingsDescription) map[string]interface{} {
	if d == nil {
		return nil
	}
	return map[string]interface{}{
		"auto_scaling_disabled": d.AutoScalingDisabled,
		"role_arn":              d.AutoScalingRoleArn,
		"scaling_policies":      d.ScalingPolicies,
		"minimum_units":         d.MinimumUnits,
		"maximum_units":         d.MaximumUnits,
	}
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
