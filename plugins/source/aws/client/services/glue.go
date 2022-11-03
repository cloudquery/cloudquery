// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/glue"
)

//go:generate mockgen -package=mocks -destination=../mocks/glue.go . GlueClient
type GlueClient interface {
	BatchGetBlueprints(context.Context, *glue.BatchGetBlueprintsInput, ...func(*glue.Options)) (*glue.BatchGetBlueprintsOutput, error)
	BatchGetCrawlers(context.Context, *glue.BatchGetCrawlersInput, ...func(*glue.Options)) (*glue.BatchGetCrawlersOutput, error)
	BatchGetCustomEntityTypes(context.Context, *glue.BatchGetCustomEntityTypesInput, ...func(*glue.Options)) (*glue.BatchGetCustomEntityTypesOutput, error)
	BatchGetDevEndpoints(context.Context, *glue.BatchGetDevEndpointsInput, ...func(*glue.Options)) (*glue.BatchGetDevEndpointsOutput, error)
	BatchGetJobs(context.Context, *glue.BatchGetJobsInput, ...func(*glue.Options)) (*glue.BatchGetJobsOutput, error)
	BatchGetPartition(context.Context, *glue.BatchGetPartitionInput, ...func(*glue.Options)) (*glue.BatchGetPartitionOutput, error)
	BatchGetTriggers(context.Context, *glue.BatchGetTriggersInput, ...func(*glue.Options)) (*glue.BatchGetTriggersOutput, error)
	BatchGetWorkflows(context.Context, *glue.BatchGetWorkflowsInput, ...func(*glue.Options)) (*glue.BatchGetWorkflowsOutput, error)
	GetBlueprint(context.Context, *glue.GetBlueprintInput, ...func(*glue.Options)) (*glue.GetBlueprintOutput, error)
	GetBlueprintRun(context.Context, *glue.GetBlueprintRunInput, ...func(*glue.Options)) (*glue.GetBlueprintRunOutput, error)
	GetBlueprintRuns(context.Context, *glue.GetBlueprintRunsInput, ...func(*glue.Options)) (*glue.GetBlueprintRunsOutput, error)
	GetCatalogImportStatus(context.Context, *glue.GetCatalogImportStatusInput, ...func(*glue.Options)) (*glue.GetCatalogImportStatusOutput, error)
	GetClassifier(context.Context, *glue.GetClassifierInput, ...func(*glue.Options)) (*glue.GetClassifierOutput, error)
	GetClassifiers(context.Context, *glue.GetClassifiersInput, ...func(*glue.Options)) (*glue.GetClassifiersOutput, error)
	GetColumnStatisticsForPartition(context.Context, *glue.GetColumnStatisticsForPartitionInput, ...func(*glue.Options)) (*glue.GetColumnStatisticsForPartitionOutput, error)
	GetColumnStatisticsForTable(context.Context, *glue.GetColumnStatisticsForTableInput, ...func(*glue.Options)) (*glue.GetColumnStatisticsForTableOutput, error)
	GetConnection(context.Context, *glue.GetConnectionInput, ...func(*glue.Options)) (*glue.GetConnectionOutput, error)
	GetConnections(context.Context, *glue.GetConnectionsInput, ...func(*glue.Options)) (*glue.GetConnectionsOutput, error)
	GetCrawler(context.Context, *glue.GetCrawlerInput, ...func(*glue.Options)) (*glue.GetCrawlerOutput, error)
	GetCrawlerMetrics(context.Context, *glue.GetCrawlerMetricsInput, ...func(*glue.Options)) (*glue.GetCrawlerMetricsOutput, error)
	GetCrawlers(context.Context, *glue.GetCrawlersInput, ...func(*glue.Options)) (*glue.GetCrawlersOutput, error)
	GetCustomEntityType(context.Context, *glue.GetCustomEntityTypeInput, ...func(*glue.Options)) (*glue.GetCustomEntityTypeOutput, error)
	GetDataCatalogEncryptionSettings(context.Context, *glue.GetDataCatalogEncryptionSettingsInput, ...func(*glue.Options)) (*glue.GetDataCatalogEncryptionSettingsOutput, error)
	GetDatabase(context.Context, *glue.GetDatabaseInput, ...func(*glue.Options)) (*glue.GetDatabaseOutput, error)
	GetDatabases(context.Context, *glue.GetDatabasesInput, ...func(*glue.Options)) (*glue.GetDatabasesOutput, error)
	GetDataflowGraph(context.Context, *glue.GetDataflowGraphInput, ...func(*glue.Options)) (*glue.GetDataflowGraphOutput, error)
	GetDevEndpoint(context.Context, *glue.GetDevEndpointInput, ...func(*glue.Options)) (*glue.GetDevEndpointOutput, error)
	GetDevEndpoints(context.Context, *glue.GetDevEndpointsInput, ...func(*glue.Options)) (*glue.GetDevEndpointsOutput, error)
	GetJob(context.Context, *glue.GetJobInput, ...func(*glue.Options)) (*glue.GetJobOutput, error)
	GetJobBookmark(context.Context, *glue.GetJobBookmarkInput, ...func(*glue.Options)) (*glue.GetJobBookmarkOutput, error)
	GetJobRun(context.Context, *glue.GetJobRunInput, ...func(*glue.Options)) (*glue.GetJobRunOutput, error)
	GetJobRuns(context.Context, *glue.GetJobRunsInput, ...func(*glue.Options)) (*glue.GetJobRunsOutput, error)
	GetJobs(context.Context, *glue.GetJobsInput, ...func(*glue.Options)) (*glue.GetJobsOutput, error)
	GetMLTaskRun(context.Context, *glue.GetMLTaskRunInput, ...func(*glue.Options)) (*glue.GetMLTaskRunOutput, error)
	GetMLTaskRuns(context.Context, *glue.GetMLTaskRunsInput, ...func(*glue.Options)) (*glue.GetMLTaskRunsOutput, error)
	GetMLTransform(context.Context, *glue.GetMLTransformInput, ...func(*glue.Options)) (*glue.GetMLTransformOutput, error)
	GetMLTransforms(context.Context, *glue.GetMLTransformsInput, ...func(*glue.Options)) (*glue.GetMLTransformsOutput, error)
	GetMapping(context.Context, *glue.GetMappingInput, ...func(*glue.Options)) (*glue.GetMappingOutput, error)
	GetPartition(context.Context, *glue.GetPartitionInput, ...func(*glue.Options)) (*glue.GetPartitionOutput, error)
	GetPartitionIndexes(context.Context, *glue.GetPartitionIndexesInput, ...func(*glue.Options)) (*glue.GetPartitionIndexesOutput, error)
	GetPartitions(context.Context, *glue.GetPartitionsInput, ...func(*glue.Options)) (*glue.GetPartitionsOutput, error)
	GetPlan(context.Context, *glue.GetPlanInput, ...func(*glue.Options)) (*glue.GetPlanOutput, error)
	GetRegistry(context.Context, *glue.GetRegistryInput, ...func(*glue.Options)) (*glue.GetRegistryOutput, error)
	GetResourcePolicies(context.Context, *glue.GetResourcePoliciesInput, ...func(*glue.Options)) (*glue.GetResourcePoliciesOutput, error)
	GetResourcePolicy(context.Context, *glue.GetResourcePolicyInput, ...func(*glue.Options)) (*glue.GetResourcePolicyOutput, error)
	GetSchema(context.Context, *glue.GetSchemaInput, ...func(*glue.Options)) (*glue.GetSchemaOutput, error)
	GetSchemaByDefinition(context.Context, *glue.GetSchemaByDefinitionInput, ...func(*glue.Options)) (*glue.GetSchemaByDefinitionOutput, error)
	GetSchemaVersion(context.Context, *glue.GetSchemaVersionInput, ...func(*glue.Options)) (*glue.GetSchemaVersionOutput, error)
	GetSchemaVersionsDiff(context.Context, *glue.GetSchemaVersionsDiffInput, ...func(*glue.Options)) (*glue.GetSchemaVersionsDiffOutput, error)
	GetSecurityConfiguration(context.Context, *glue.GetSecurityConfigurationInput, ...func(*glue.Options)) (*glue.GetSecurityConfigurationOutput, error)
	GetSecurityConfigurations(context.Context, *glue.GetSecurityConfigurationsInput, ...func(*glue.Options)) (*glue.GetSecurityConfigurationsOutput, error)
	GetSession(context.Context, *glue.GetSessionInput, ...func(*glue.Options)) (*glue.GetSessionOutput, error)
	GetStatement(context.Context, *glue.GetStatementInput, ...func(*glue.Options)) (*glue.GetStatementOutput, error)
	GetTable(context.Context, *glue.GetTableInput, ...func(*glue.Options)) (*glue.GetTableOutput, error)
	GetTableVersion(context.Context, *glue.GetTableVersionInput, ...func(*glue.Options)) (*glue.GetTableVersionOutput, error)
	GetTableVersions(context.Context, *glue.GetTableVersionsInput, ...func(*glue.Options)) (*glue.GetTableVersionsOutput, error)
	GetTables(context.Context, *glue.GetTablesInput, ...func(*glue.Options)) (*glue.GetTablesOutput, error)
	GetTags(context.Context, *glue.GetTagsInput, ...func(*glue.Options)) (*glue.GetTagsOutput, error)
	GetTrigger(context.Context, *glue.GetTriggerInput, ...func(*glue.Options)) (*glue.GetTriggerOutput, error)
	GetTriggers(context.Context, *glue.GetTriggersInput, ...func(*glue.Options)) (*glue.GetTriggersOutput, error)
	GetUnfilteredPartitionMetadata(context.Context, *glue.GetUnfilteredPartitionMetadataInput, ...func(*glue.Options)) (*glue.GetUnfilteredPartitionMetadataOutput, error)
	GetUnfilteredPartitionsMetadata(context.Context, *glue.GetUnfilteredPartitionsMetadataInput, ...func(*glue.Options)) (*glue.GetUnfilteredPartitionsMetadataOutput, error)
	GetUnfilteredTableMetadata(context.Context, *glue.GetUnfilteredTableMetadataInput, ...func(*glue.Options)) (*glue.GetUnfilteredTableMetadataOutput, error)
	GetUserDefinedFunction(context.Context, *glue.GetUserDefinedFunctionInput, ...func(*glue.Options)) (*glue.GetUserDefinedFunctionOutput, error)
	GetUserDefinedFunctions(context.Context, *glue.GetUserDefinedFunctionsInput, ...func(*glue.Options)) (*glue.GetUserDefinedFunctionsOutput, error)
	GetWorkflow(context.Context, *glue.GetWorkflowInput, ...func(*glue.Options)) (*glue.GetWorkflowOutput, error)
	GetWorkflowRun(context.Context, *glue.GetWorkflowRunInput, ...func(*glue.Options)) (*glue.GetWorkflowRunOutput, error)
	GetWorkflowRunProperties(context.Context, *glue.GetWorkflowRunPropertiesInput, ...func(*glue.Options)) (*glue.GetWorkflowRunPropertiesOutput, error)
	GetWorkflowRuns(context.Context, *glue.GetWorkflowRunsInput, ...func(*glue.Options)) (*glue.GetWorkflowRunsOutput, error)
	ListBlueprints(context.Context, *glue.ListBlueprintsInput, ...func(*glue.Options)) (*glue.ListBlueprintsOutput, error)
	ListCrawlers(context.Context, *glue.ListCrawlersInput, ...func(*glue.Options)) (*glue.ListCrawlersOutput, error)
	ListCrawls(context.Context, *glue.ListCrawlsInput, ...func(*glue.Options)) (*glue.ListCrawlsOutput, error)
	ListCustomEntityTypes(context.Context, *glue.ListCustomEntityTypesInput, ...func(*glue.Options)) (*glue.ListCustomEntityTypesOutput, error)
	ListDevEndpoints(context.Context, *glue.ListDevEndpointsInput, ...func(*glue.Options)) (*glue.ListDevEndpointsOutput, error)
	ListJobs(context.Context, *glue.ListJobsInput, ...func(*glue.Options)) (*glue.ListJobsOutput, error)
	ListMLTransforms(context.Context, *glue.ListMLTransformsInput, ...func(*glue.Options)) (*glue.ListMLTransformsOutput, error)
	ListRegistries(context.Context, *glue.ListRegistriesInput, ...func(*glue.Options)) (*glue.ListRegistriesOutput, error)
	ListSchemaVersions(context.Context, *glue.ListSchemaVersionsInput, ...func(*glue.Options)) (*glue.ListSchemaVersionsOutput, error)
	ListSchemas(context.Context, *glue.ListSchemasInput, ...func(*glue.Options)) (*glue.ListSchemasOutput, error)
	ListSessions(context.Context, *glue.ListSessionsInput, ...func(*glue.Options)) (*glue.ListSessionsOutput, error)
	ListStatements(context.Context, *glue.ListStatementsInput, ...func(*glue.Options)) (*glue.ListStatementsOutput, error)
	ListTriggers(context.Context, *glue.ListTriggersInput, ...func(*glue.Options)) (*glue.ListTriggersOutput, error)
	ListWorkflows(context.Context, *glue.ListWorkflowsInput, ...func(*glue.Options)) (*glue.ListWorkflowsOutput, error)
	QuerySchemaVersionMetadata(context.Context, *glue.QuerySchemaVersionMetadataInput, ...func(*glue.Options)) (*glue.QuerySchemaVersionMetadataOutput, error)
	SearchTables(context.Context, *glue.SearchTablesInput, ...func(*glue.Options)) (*glue.SearchTablesOutput, error)
}
