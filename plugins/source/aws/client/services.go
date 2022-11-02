package client

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer"
	"github.com/aws/aws-sdk-go-v2/service/acm"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling"
	"github.com/aws/aws-sdk-go-v2/service/apprunner"
	"github.com/aws/aws-sdk-go-v2/service/appsync"
	"github.com/aws/aws-sdk-go-v2/service/athena"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/aws/aws-sdk-go-v2/service/backup"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/aws/aws-sdk-go-v2/service/cloudhsmv2"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/codebuild"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice"
	"github.com/aws/aws-sdk-go-v2/service/dax"
	"github.com/aws/aws-sdk-go-v2/service/directconnect"
	"github.com/aws/aws-sdk-go-v2/service/docdb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/ecrpublic"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/efs"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
	elbv1 "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
	elbv2 "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice"
	"github.com/aws/aws-sdk-go-v2/service/emr"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/aws/aws-sdk-go-v2/service/firehose"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector"
	"github.com/aws/aws-sdk-go-v2/service/fsx"
	"github.com/aws/aws-sdk-go-v2/service/glacier"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/guardduty"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/inspector"
	"github.com/aws/aws-sdk-go-v2/service/inspector2"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/aws/aws-sdk-go-v2/service/mq"
	"github.com/aws/aws-sdk-go-v2/service/neptune"
	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/aws/aws-sdk-go-v2/service/qldb"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroups"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/route53domains"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3control"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalogappregistry"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/shield"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/transfer"
	"github.com/aws/aws-sdk-go-v2/service/waf"
	"github.com/aws/aws-sdk-go-v2/service/wafregional"
	"github.com/aws/aws-sdk-go-v2/service/wafv2"
	"github.com/aws/aws-sdk-go-v2/service/workspaces"
	"github.com/aws/aws-sdk-go-v2/service/xray"
)

//go:generate mockgen -package=mocks -destination=./mocks/mock_acm.go . ACMClient
type ACMClient interface {
	DescribeCertificate(ctx context.Context, params *acm.DescribeCertificateInput, optFns ...func(*acm.Options)) (*acm.DescribeCertificateOutput, error)
	ListCertificates(ctx context.Context, params *acm.ListCertificatesInput, optFns ...func(*acm.Options)) (*acm.ListCertificatesOutput, error)
	ListTagsForCertificate(ctx context.Context, params *acm.ListTagsForCertificateInput, optFns ...func(*acm.Options)) (*acm.ListTagsForCertificateOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_analyzer.go . AnalyzerClient
type AnalyzerClient interface {
	accessanalyzer.ListAnalyzersAPIClient
	accessanalyzer.ListFindingsAPIClient
	accessanalyzer.ListArchiveRulesAPIClient
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_application_autoscaling.go . ApplicationAutoscalingClient
type ApplicationAutoscalingClient interface {
	DescribeScalingPolicies(ctx context.Context, params *applicationautoscaling.DescribeScalingPoliciesInput, optFns ...func(*applicationautoscaling.Options)) (*applicationautoscaling.DescribeScalingPoliciesOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/apigateway.go . ApigatewayClient
type ApigatewayClient interface {
	GetApiKeys(ctx context.Context, params *apigateway.GetApiKeysInput, optFns ...func(*apigateway.Options)) (*apigateway.GetApiKeysOutput, error)
	GetAuthorizers(ctx context.Context, params *apigateway.GetAuthorizersInput, optFns ...func(*apigateway.Options)) (*apigateway.GetAuthorizersOutput, error)
	GetBasePathMappings(ctx context.Context, params *apigateway.GetBasePathMappingsInput, optFns ...func(*apigateway.Options)) (*apigateway.GetBasePathMappingsOutput, error)
	GetClientCertificates(ctx context.Context, params *apigateway.GetClientCertificatesInput, optFns ...func(*apigateway.Options)) (*apigateway.GetClientCertificatesOutput, error)
	GetDeployments(ctx context.Context, params *apigateway.GetDeploymentsInput, optFns ...func(*apigateway.Options)) (*apigateway.GetDeploymentsOutput, error)
	GetDocumentationParts(ctx context.Context, params *apigateway.GetDocumentationPartsInput, optFns ...func(*apigateway.Options)) (*apigateway.GetDocumentationPartsOutput, error)
	GetDocumentationVersions(ctx context.Context, params *apigateway.GetDocumentationVersionsInput, optFns ...func(*apigateway.Options)) (*apigateway.GetDocumentationVersionsOutput, error)
	GetDomainNames(ctx context.Context, params *apigateway.GetDomainNamesInput, optFns ...func(*apigateway.Options)) (*apigateway.GetDomainNamesOutput, error)
	GetGatewayResponses(ctx context.Context, params *apigateway.GetGatewayResponsesInput, optFns ...func(*apigateway.Options)) (*apigateway.GetGatewayResponsesOutput, error)
	GetModels(ctx context.Context, params *apigateway.GetModelsInput, optFns ...func(*apigateway.Options)) (*apigateway.GetModelsOutput, error)
	GetModelTemplate(ctx context.Context, params *apigateway.GetModelTemplateInput, optFns ...func(*apigateway.Options)) (*apigateway.GetModelTemplateOutput, error)
	GetRequestValidators(ctx context.Context, params *apigateway.GetRequestValidatorsInput, optFns ...func(*apigateway.Options)) (*apigateway.GetRequestValidatorsOutput, error)
	GetResources(ctx context.Context, params *apigateway.GetResourcesInput, optFns ...func(*apigateway.Options)) (*apigateway.GetResourcesOutput, error)
	GetRestApis(ctx context.Context, params *apigateway.GetRestApisInput, optFns ...func(*apigateway.Options)) (*apigateway.GetRestApisOutput, error)
	GetStages(ctx context.Context, params *apigateway.GetStagesInput, optFns ...func(*apigateway.Options)) (*apigateway.GetStagesOutput, error)
	GetUsagePlanKeys(ctx context.Context, params *apigateway.GetUsagePlanKeysInput, optFns ...func(*apigateway.Options)) (*apigateway.GetUsagePlanKeysOutput, error)
	GetUsagePlans(ctx context.Context, params *apigateway.GetUsagePlansInput, optFns ...func(*apigateway.Options)) (*apigateway.GetUsagePlansOutput, error)
	GetVpcLinks(ctx context.Context, params *apigateway.GetVpcLinksInput, optFns ...func(*apigateway.Options)) (*apigateway.GetVpcLinksOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_apigatewayv2.go . Apigatewayv2Client
type Apigatewayv2Client interface {
	GetApiMappings(ctx context.Context, params *apigatewayv2.GetApiMappingsInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.GetApiMappingsOutput, error)
	GetApis(ctx context.Context, params *apigatewayv2.GetApisInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.GetApisOutput, error)
	GetAuthorizers(ctx context.Context, params *apigatewayv2.GetAuthorizersInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.GetAuthorizersOutput, error)
	GetDeployments(ctx context.Context, params *apigatewayv2.GetDeploymentsInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.GetDeploymentsOutput, error)
	GetDomainNames(ctx context.Context, params *apigatewayv2.GetDomainNamesInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.GetDomainNamesOutput, error)
	GetIntegrationResponses(ctx context.Context, params *apigatewayv2.GetIntegrationResponsesInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.GetIntegrationResponsesOutput, error)
	GetIntegrations(ctx context.Context, params *apigatewayv2.GetIntegrationsInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.GetIntegrationsOutput, error)
	GetModels(ctx context.Context, params *apigatewayv2.GetModelsInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.GetModelsOutput, error)
	GetModelTemplate(ctx context.Context, params *apigatewayv2.GetModelTemplateInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.GetModelTemplateOutput, error)
	GetRouteResponses(ctx context.Context, params *apigatewayv2.GetRouteResponsesInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.GetRouteResponsesOutput, error)
	GetRoutes(ctx context.Context, params *apigatewayv2.GetRoutesInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.GetRoutesOutput, error)
	GetStages(ctx context.Context, params *apigatewayv2.GetStagesInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.GetStagesOutput, error)
	GetTags(ctx context.Context, params *apigatewayv2.GetTagsInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.GetTagsOutput, error)
	GetVpcLinks(ctx context.Context, params *apigatewayv2.GetVpcLinksInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.GetVpcLinksOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_apprunner.go . AppRunnerClient
type AppRunnerClient interface {
	DescribeAutoScalingConfiguration(ctx context.Context, params *apprunner.DescribeAutoScalingConfigurationInput, optFns ...func(*apprunner.Options)) (*apprunner.DescribeAutoScalingConfigurationOutput, error)
	DescribeCustomDomains(ctx context.Context, params *apprunner.DescribeCustomDomainsInput, optFns ...func(*apprunner.Options)) (*apprunner.DescribeCustomDomainsOutput, error)
	DescribeObservabilityConfiguration(ctx context.Context, params *apprunner.DescribeObservabilityConfigurationInput, optFns ...func(*apprunner.Options)) (*apprunner.DescribeObservabilityConfigurationOutput, error)
	DescribeService(ctx context.Context, params *apprunner.DescribeServiceInput, optFns ...func(*apprunner.Options)) (*apprunner.DescribeServiceOutput, error)
	DescribeVpcConnector(ctx context.Context, params *apprunner.DescribeVpcConnectorInput, optFns ...func(*apprunner.Options)) (*apprunner.DescribeVpcConnectorOutput, error)
	ListAutoScalingConfigurations(ctx context.Context, params *apprunner.ListAutoScalingConfigurationsInput, optFns ...func(*apprunner.Options)) (*apprunner.ListAutoScalingConfigurationsOutput, error)
	ListConnections(ctx context.Context, params *apprunner.ListConnectionsInput, optFns ...func(*apprunner.Options)) (*apprunner.ListConnectionsOutput, error)
	ListObservabilityConfigurations(ctx context.Context, params *apprunner.ListObservabilityConfigurationsInput, optFns ...func(*apprunner.Options)) (*apprunner.ListObservabilityConfigurationsOutput, error)
	ListOperations(ctx context.Context, params *apprunner.ListOperationsInput, optFns ...func(*apprunner.Options)) (*apprunner.ListOperationsOutput, error)
	ListServices(ctx context.Context, params *apprunner.ListServicesInput, optFns ...func(*apprunner.Options)) (*apprunner.ListServicesOutput, error)
	ListTagsForResource(ctx context.Context, params *apprunner.ListTagsForResourceInput, optFns ...func(*apprunner.Options)) (*apprunner.ListTagsForResourceOutput, error)
	ListVpcConnectors(ctx context.Context, params *apprunner.ListVpcConnectorsInput, optFns ...func(*apprunner.Options)) (*apprunner.ListVpcConnectorsOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_appsync.go . AppSyncClient
type AppSyncClient interface {
	ListGraphqlApis(ctx context.Context, params *appsync.ListGraphqlApisInput, optFns ...func(*appsync.Options)) (*appsync.ListGraphqlApisOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_autoscaling.go . AutoscalingClient
type AutoscalingClient interface {
	DescribeLaunchConfigurations(context.Context, *autoscaling.DescribeLaunchConfigurationsInput, ...func(*autoscaling.Options)) (*autoscaling.DescribeLaunchConfigurationsOutput, error)
	DescribeAutoScalingGroups(ctx context.Context, params *autoscaling.DescribeAutoScalingGroupsInput, optFns ...func(*autoscaling.Options)) (*autoscaling.DescribeAutoScalingGroupsOutput, error)
	DescribePolicies(ctx context.Context, params *autoscaling.DescribePoliciesInput, optFns ...func(*autoscaling.Options)) (*autoscaling.DescribePoliciesOutput, error)
	DescribeTags(ctx context.Context, params *autoscaling.DescribeTagsInput, optFns ...func(*autoscaling.Options)) (*autoscaling.DescribeTagsOutput, error)
	DescribeNotificationConfigurations(ctx context.Context, params *autoscaling.DescribeNotificationConfigurationsInput, optFns ...func(*autoscaling.Options)) (*autoscaling.DescribeNotificationConfigurationsOutput, error)
	DescribeLoadBalancers(ctx context.Context, params *autoscaling.DescribeLoadBalancersInput, optFns ...func(*autoscaling.Options)) (*autoscaling.DescribeLoadBalancersOutput, error)
	DescribeLoadBalancerTargetGroups(ctx context.Context, params *autoscaling.DescribeLoadBalancerTargetGroupsInput, optFns ...func(*autoscaling.Options)) (*autoscaling.DescribeLoadBalancerTargetGroupsOutput, error)
	DescribeLifecycleHooks(ctx context.Context, params *autoscaling.DescribeLifecycleHooksInput, optFns ...func(*autoscaling.Options)) (*autoscaling.DescribeLifecycleHooksOutput, error)
	DescribeScheduledActions(ctx context.Context, params *autoscaling.DescribeScheduledActionsInput, optFns ...func(options *autoscaling.Options)) (*autoscaling.DescribeScheduledActionsOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_athena.go . AthenaClient
type AthenaClient interface {
	ListDataCatalogs(ctx context.Context, params *athena.ListDataCatalogsInput, optFns ...func(*athena.Options)) (*athena.ListDataCatalogsOutput, error)
	GetDataCatalog(ctx context.Context, params *athena.GetDataCatalogInput, optFns ...func(*athena.Options)) (*athena.GetDataCatalogOutput, error)
	ListDatabases(ctx context.Context, params *athena.ListDatabasesInput, optFns ...func(*athena.Options)) (*athena.ListDatabasesOutput, error)
	ListTableMetadata(ctx context.Context, params *athena.ListTableMetadataInput, optFns ...func(*athena.Options)) (*athena.ListTableMetadataOutput, error)
	ListTagsForResource(ctx context.Context, params *athena.ListTagsForResourceInput, optFns ...func(*athena.Options)) (*athena.ListTagsForResourceOutput, error)
	ListWorkGroups(ctx context.Context, params *athena.ListWorkGroupsInput, optFns ...func(*athena.Options)) (*athena.ListWorkGroupsOutput, error)
	GetWorkGroup(ctx context.Context, params *athena.GetWorkGroupInput, optFns ...func(*athena.Options)) (*athena.GetWorkGroupOutput, error)
	ListPreparedStatements(ctx context.Context, params *athena.ListPreparedStatementsInput, optFns ...func(*athena.Options)) (*athena.ListPreparedStatementsOutput, error)
	GetPreparedStatement(ctx context.Context, params *athena.GetPreparedStatementInput, optFns ...func(*athena.Options)) (*athena.GetPreparedStatementOutput, error)
	ListQueryExecutions(ctx context.Context, params *athena.ListQueryExecutionsInput, optFns ...func(*athena.Options)) (*athena.ListQueryExecutionsOutput, error)
	ListNamedQueries(ctx context.Context, params *athena.ListNamedQueriesInput, optFns ...func(*athena.Options)) (*athena.ListNamedQueriesOutput, error)
	GetNamedQuery(ctx context.Context, params *athena.GetNamedQueryInput, optFns ...func(*athena.Options)) (*athena.GetNamedQueryOutput, error)
	GetQueryExecution(ctx context.Context, params *athena.GetQueryExecutionInput, optFns ...func(*athena.Options)) (*athena.GetQueryExecutionOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/backup.go . BackupClient
type BackupClient interface {
	GetBackupPlan(ctx context.Context, params *backup.GetBackupPlanInput, optFns ...func(*backup.Options)) (*backup.GetBackupPlanOutput, error)
	GetBackupSelection(ctx context.Context, params *backup.GetBackupSelectionInput, optFns ...func(*backup.Options)) (*backup.GetBackupSelectionOutput, error)
	GetBackupVaultAccessPolicy(ctx context.Context, params *backup.GetBackupVaultAccessPolicyInput, optFns ...func(*backup.Options)) (*backup.GetBackupVaultAccessPolicyOutput, error)
	GetBackupVaultNotifications(ctx context.Context, params *backup.GetBackupVaultNotificationsInput, optFns ...func(*backup.Options)) (*backup.GetBackupVaultNotificationsOutput, error)
	ListBackupPlans(ctx context.Context, params *backup.ListBackupPlansInput, optFns ...func(*backup.Options)) (*backup.ListBackupPlansOutput, error)
	ListBackupSelections(ctx context.Context, params *backup.ListBackupSelectionsInput, optFns ...func(*backup.Options)) (*backup.ListBackupSelectionsOutput, error)
	ListBackupVaults(ctx context.Context, params *backup.ListBackupVaultsInput, optFns ...func(*backup.Options)) (*backup.ListBackupVaultsOutput, error)
	ListRecoveryPointsByBackupVault(ctx context.Context, params *backup.ListRecoveryPointsByBackupVaultInput, optFns ...func(*backup.Options)) (*backup.ListRecoveryPointsByBackupVaultOutput, error)
	ListTags(ctx context.Context, params *backup.ListTagsInput, optFns ...func(*backup.Options)) (*backup.ListTagsOutput, error)
	DescribeGlobalSettings(ctx context.Context, params *backup.DescribeGlobalSettingsInput, optFns ...func(*backup.Options)) (*backup.DescribeGlobalSettingsOutput, error)
	DescribeRegionSettings(ctx context.Context, params *backup.DescribeRegionSettingsInput, optFns ...func(*backup.Options)) (*backup.DescribeRegionSettingsOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_cloudhsmv2.go . CloudHSMV2Client
type CloudHSMV2Client interface {
	DescribeBackups(ctx context.Context, params *cloudhsmv2.DescribeBackupsInput, optFns ...func(*cloudhsmv2.Options)) (*cloudhsmv2.DescribeBackupsOutput, error)
	DescribeClusters(ctx context.Context, params *cloudhsmv2.DescribeClustersInput, optFns ...func(*cloudhsmv2.Options)) (*cloudhsmv2.DescribeClustersOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_cloudformation.go . CloudFormationClient
type CloudFormationClient interface {
	cloudformation.DescribeStacksAPIClient
	cloudformation.ListStackResourcesAPIClient
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_cloudfront.go . CloudfrontClient
type CloudfrontClient interface {
	ListDistributions(ctx context.Context, params *cloudfront.ListDistributionsInput, optFns ...func(*cloudfront.Options)) (*cloudfront.ListDistributionsOutput, error)
	ListDistributionsByWebACLId(ctx context.Context, params *cloudfront.ListDistributionsByWebACLIdInput, optFns ...func(options *cloudfront.Options)) (*cloudfront.ListDistributionsByWebACLIdOutput, error)
	ListCachePolicies(ctx context.Context, params *cloudfront.ListCachePoliciesInput, optFns ...func(*cloudfront.Options)) (*cloudfront.ListCachePoliciesOutput, error)
	ListTagsForResource(ctx context.Context, params *cloudfront.ListTagsForResourceInput, optFns ...func(*cloudfront.Options)) (*cloudfront.ListTagsForResourceOutput, error)
	GetDistribution(ctx context.Context, params *cloudfront.GetDistributionInput, optFns ...func(*cloudfront.Options)) (*cloudfront.GetDistributionOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_cloudtrail.go . CloudtrailClient
type CloudtrailClient interface {
	GetEventSelectors(ctx context.Context, params *cloudtrail.GetEventSelectorsInput, optFns ...func(*cloudtrail.Options)) (*cloudtrail.GetEventSelectorsOutput, error)
	DescribeTrails(ctx context.Context, params *cloudtrail.DescribeTrailsInput, optFns ...func(*cloudtrail.Options)) (*cloudtrail.DescribeTrailsOutput, error)
	GetTrailStatus(ctx context.Context, params *cloudtrail.GetTrailStatusInput, optFns ...func(*cloudtrail.Options)) (*cloudtrail.GetTrailStatusOutput, error)
	ListTags(ctx context.Context, params *cloudtrail.ListTagsInput, optFns ...func(*cloudtrail.Options)) (*cloudtrail.ListTagsOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_cloudwatch.go . CloudwatchClient
type CloudwatchClient interface {
	DescribeAlarms(ctx context.Context, params *cloudwatch.DescribeAlarmsInput, optFns ...func(*cloudwatch.Options)) (*cloudwatch.DescribeAlarmsOutput, error)
	ListTagsForResource(ctx context.Context, params *cloudwatch.ListTagsForResourceInput, optFns ...func(*cloudwatch.Options)) (*cloudwatch.ListTagsForResourceOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_cloudwatchlogs.go . CloudwatchLogsClient
type CloudwatchLogsClient interface {
	DescribeMetricFilters(ctx context.Context, params *cloudwatchlogs.DescribeMetricFiltersInput, optFns ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.DescribeMetricFiltersOutput, error)
	DescribeLogGroups(ctx context.Context, params *cloudwatchlogs.DescribeLogGroupsInput, optFns ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.DescribeLogGroupsOutput, error)
	ListTagsLogGroup(ctx context.Context, params *cloudwatchlogs.ListTagsLogGroupInput, optFns ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.ListTagsLogGroupOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/codebuild.go . CodebuildClient
type CodebuildClient interface {
	BatchGetProjects(ctx context.Context, params *codebuild.BatchGetProjectsInput, optFns ...func(*codebuild.Options)) (*codebuild.BatchGetProjectsOutput, error)
	ListProjects(ctx context.Context, params *codebuild.ListProjectsInput, optFns ...func(*codebuild.Options)) (*codebuild.ListProjectsOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_codepipeline.go . CodePipelineClient
type CodePipelineClient interface {
	ListPipelines(ctx context.Context, params *codepipeline.ListPipelinesInput, optFns ...func(*codepipeline.Options)) (*codepipeline.ListPipelinesOutput, error)
	GetPipeline(ctx context.Context, params *codepipeline.GetPipelineInput, optFns ...func(*codepipeline.Options)) (*codepipeline.GetPipelineOutput, error)
	ListTagsForResource(ctx context.Context, params *codepipeline.ListTagsForResourceInput, optFns ...func(*codepipeline.Options)) (*codepipeline.ListTagsForResourceOutput, error)
	ListWebhooks(ctx context.Context, params *codepipeline.ListWebhooksInput, optFns ...func(*codepipeline.Options)) (*codepipeline.ListWebhooksOutput, error)
}

//go:generate mockgen -destination=./mocks/mock_cognitoidentitypools.go -package=mocks . CognitoIdentityPoolsClient
type CognitoIdentityPoolsClient interface {
	DescribeIdentityPool(ctx context.Context, params *cognitoidentity.DescribeIdentityPoolInput, optFns ...func(*cognitoidentity.Options)) (*cognitoidentity.DescribeIdentityPoolOutput, error)
	ListIdentityPools(ctx context.Context, params *cognitoidentity.ListIdentityPoolsInput, optFns ...func(*cognitoidentity.Options)) (*cognitoidentity.ListIdentityPoolsOutput, error)
}

//go:generate mockgen -destination=./mocks/mock_cognitouserpools.go -package=mocks . CognitoUserPoolsClient
type CognitoUserPoolsClient interface {
	DescribeIdentityProvider(ctx context.Context, params *cognitoidentityprovider.DescribeIdentityProviderInput, optFns ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DescribeIdentityProviderOutput, error)
	DescribeUserPool(ctx context.Context, params *cognitoidentityprovider.DescribeUserPoolInput, optFns ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DescribeUserPoolOutput, error)
	ListIdentityProviders(ctx context.Context, params *cognitoidentityprovider.ListIdentityProvidersInput, optFns ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListIdentityProvidersOutput, error)
	ListUserPools(ctx context.Context, params *cognitoidentityprovider.ListUserPoolsInput, optFns ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListUserPoolsOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_configservice.go . ConfigServiceClient
type ConfigServiceClient interface {
	DescribeConfigurationRecorders(ctx context.Context, params *configservice.DescribeConfigurationRecordersInput, optFns ...func(*configservice.Options)) (*configservice.DescribeConfigurationRecordersOutput, error)
	DescribeConfigurationRecorderStatus(ctx context.Context, params *configservice.DescribeConfigurationRecorderStatusInput, optFns ...func(*configservice.Options)) (*configservice.DescribeConfigurationRecorderStatusOutput, error)
	GetConformancePackComplianceDetails(ctx context.Context, params *configservice.GetConformancePackComplianceDetailsInput, optFns ...func(*configservice.Options)) (*configservice.GetConformancePackComplianceDetailsOutput, error)
	configservice.DescribeConformancePackComplianceAPIClient
	configservice.DescribeConformancePacksAPIClient
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_databasemigrationservice.go . DatabasemigrationserviceClient
type DatabasemigrationserviceClient interface {
	DescribeReplicationInstances(ctx context.Context, params *databasemigrationservice.DescribeReplicationInstancesInput, optFns ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeReplicationInstancesOutput, error)
	ListTagsForResource(ctx context.Context, params *databasemigrationservice.ListTagsForResourceInput, optFns ...func(*databasemigrationservice.Options)) (*databasemigrationservice.ListTagsForResourceOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_dax.go . DAXClient
type DAXClient interface {
	DescribeClusters(ctx context.Context, params *dax.DescribeClustersInput, optFns ...func(*dax.Options)) (*dax.DescribeClustersOutput, error)
	ListTags(ctx context.Context, params *dax.ListTagsInput, optFns ...func(*dax.Options)) (*dax.ListTagsOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_directconnect.go . DirectconnectClient
type DirectconnectClient interface {
	DescribeConnections(ctx context.Context, params *directconnect.DescribeConnectionsInput, optFns ...func(*directconnect.Options)) (*directconnect.DescribeConnectionsOutput, error)
	DescribeDirectConnectGatewayAssociations(ctx context.Context, params *directconnect.DescribeDirectConnectGatewayAssociationsInput, optFns ...func(*directconnect.Options)) (*directconnect.DescribeDirectConnectGatewayAssociationsOutput, error)
	DescribeDirectConnectGatewayAttachments(ctx context.Context, params *directconnect.DescribeDirectConnectGatewayAttachmentsInput, optFns ...func(*directconnect.Options)) (*directconnect.DescribeDirectConnectGatewayAttachmentsOutput, error)
	DescribeDirectConnectGateways(ctx context.Context, params *directconnect.DescribeDirectConnectGatewaysInput, optFns ...func(*directconnect.Options)) (*directconnect.DescribeDirectConnectGatewaysOutput, error)
	DescribeLags(ctx context.Context, params *directconnect.DescribeLagsInput, optFns ...func(*directconnect.Options)) (*directconnect.DescribeLagsOutput, error)
	DescribeVirtualGateways(ctx context.Context, params *directconnect.DescribeVirtualGatewaysInput, optFns ...func(*directconnect.Options)) (*directconnect.DescribeVirtualGatewaysOutput, error)
	DescribeVirtualInterfaces(ctx context.Context, params *directconnect.DescribeVirtualInterfacesInput, optFns ...func(*directconnect.Options)) (*directconnect.DescribeVirtualInterfacesOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_docdb.go . DocDBClient
type DocDBClient interface {
	DescribeDBClusters(ctx context.Context, params *docdb.DescribeDBClustersInput, optFns ...func(*docdb.Options)) (*docdb.DescribeDBClustersOutput, error)
	ListTagsForResource(ctx context.Context, params *docdb.ListTagsForResourceInput, optFns ...func(*docdb.Options)) (*docdb.ListTagsForResourceOutput, error)
	DescribeDBClusterSnapshots(ctx context.Context, params *docdb.DescribeDBClusterSnapshotsInput, optFns ...func(*docdb.Options)) (*docdb.DescribeDBClusterSnapshotsOutput, error)
	DescribeDBClusterSnapshotAttributes(ctx context.Context, params *docdb.DescribeDBClusterSnapshotAttributesInput, optFns ...func(*docdb.Options)) (*docdb.DescribeDBClusterSnapshotAttributesOutput, error)
	DescribeDBClusterParameters(ctx context.Context, params *docdb.DescribeDBClusterParametersInput, optFns ...func(*docdb.Options)) (*docdb.DescribeDBClusterParametersOutput, error)
	DescribeDBClusterParameterGroups(ctx context.Context, params *docdb.DescribeDBClusterParameterGroupsInput, optFns ...func(*docdb.Options)) (*docdb.DescribeDBClusterParameterGroupsOutput, error)
	DescribeCertificates(ctx context.Context, params *docdb.DescribeCertificatesInput, optFns ...func(*docdb.Options)) (*docdb.DescribeCertificatesOutput, error)
	DescribeDBEngineVersions(ctx context.Context, params *docdb.DescribeDBEngineVersionsInput, optFns ...func(*docdb.Options)) (*docdb.DescribeDBEngineVersionsOutput, error)
	DescribeDBInstances(ctx context.Context, params *docdb.DescribeDBInstancesInput, optFns ...func(*docdb.Options)) (*docdb.DescribeDBInstancesOutput, error)
	DescribeDBSubnetGroups(ctx context.Context, params *docdb.DescribeDBSubnetGroupsInput, optFns ...func(*docdb.Options)) (*docdb.DescribeDBSubnetGroupsOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_dynamodb.go . DynamoDBClient
type DynamoDBClient interface {
	ListTables(ctx context.Context, params *dynamodb.ListTablesInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ListTablesOutput, error)
	DescribeTable(ctx context.Context, params *dynamodb.DescribeTableInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DescribeTableOutput, error)
	DescribeTableReplicaAutoScaling(ctx context.Context, params *dynamodb.DescribeTableReplicaAutoScalingInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DescribeTableReplicaAutoScalingOutput, error)
	DescribeContinuousBackups(ctx context.Context, params *dynamodb.DescribeContinuousBackupsInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DescribeContinuousBackupsOutput, error)
	ListTagsOfResource(ctx context.Context, params *dynamodb.ListTagsOfResourceInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ListTagsOfResourceOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_ec2.go . Ec2Client
type Ec2Client interface {
	DescribeAddresses(ctx context.Context, params *ec2.DescribeAddressesInput, optFns ...func(*ec2.Options)) (*ec2.DescribeAddressesOutput, error)
	DescribeByoipCidrs(ctx context.Context, params *ec2.DescribeByoipCidrsInput, optFns ...func(*ec2.Options)) (*ec2.DescribeByoipCidrsOutput, error)
	DescribeCustomerGateways(ctx context.Context, params *ec2.DescribeCustomerGatewaysInput, optFns ...func(*ec2.Options)) (*ec2.DescribeCustomerGatewaysOutput, error)
	DescribeEgressOnlyInternetGateways(ctx context.Context, params *ec2.DescribeEgressOnlyInternetGatewaysInput, optFns ...func(*ec2.Options)) (*ec2.DescribeEgressOnlyInternetGatewaysOutput, error)
	DescribeFlowLogs(ctx context.Context, params *ec2.DescribeFlowLogsInput, optFns ...func(*ec2.Options)) (*ec2.DescribeFlowLogsOutput, error)
	DescribeHosts(ctx context.Context, params *ec2.DescribeHostsInput, optFns ...func(*ec2.Options)) (*ec2.DescribeHostsOutput, error)
	DescribeImageAttribute(ctx context.Context, params *ec2.DescribeImageAttributeInput, optFns ...func(*ec2.Options)) (*ec2.DescribeImageAttributeOutput, error)
	DescribeImages(ctx context.Context, params *ec2.DescribeImagesInput, optFns ...func(*ec2.Options)) (*ec2.DescribeImagesOutput, error)
	DescribeInstances(ctx context.Context, params *ec2.DescribeInstancesInput, optFns ...func(*ec2.Options)) (*ec2.DescribeInstancesOutput, error)
	DescribeInstanceStatus(ctx context.Context, params *ec2.DescribeInstanceStatusInput, optFns ...func(*ec2.Options)) (*ec2.DescribeInstanceStatusOutput, error)
	DescribeInstanceTypes(ctx context.Context, params *ec2.DescribeInstanceTypesInput, optFns ...func(*ec2.Options)) (*ec2.DescribeInstanceTypesOutput, error)
	DescribeInternetGateways(ctx context.Context, params *ec2.DescribeInternetGatewaysInput, optFns ...func(*ec2.Options)) (*ec2.DescribeInternetGatewaysOutput, error)
	DescribeKeyPairs(ctx context.Context, params *ec2.DescribeKeyPairsInput, optFns ...func(*ec2.Options)) (*ec2.DescribeKeyPairsOutput, error)
	DescribeNatGateways(ctx context.Context, params *ec2.DescribeNatGatewaysInput, optFns ...func(*ec2.Options)) (*ec2.DescribeNatGatewaysOutput, error)
	DescribeNetworkAcls(ctx context.Context, params *ec2.DescribeNetworkAclsInput, optFns ...func(*ec2.Options)) (*ec2.DescribeNetworkAclsOutput, error)
	DescribeNetworkInterfaces(ctx context.Context, params *ec2.DescribeNetworkInterfacesInput, optFns ...func(*ec2.Options)) (*ec2.DescribeNetworkInterfacesOutput, error)
	DescribeRegions(ctx context.Context, params *ec2.DescribeRegionsInput, optFns ...func(*ec2.Options)) (*ec2.DescribeRegionsOutput, error)
	DescribeReservedInstances(ctx context.Context, params *ec2.DescribeReservedInstancesInput, optFns ...func(*ec2.Options)) (*ec2.DescribeReservedInstancesOutput, error)
	DescribeRouteTables(ctx context.Context, params *ec2.DescribeRouteTablesInput, optFns ...func(*ec2.Options)) (*ec2.DescribeRouteTablesOutput, error)
	DescribeSecurityGroups(ctx context.Context, params *ec2.DescribeSecurityGroupsInput, optFns ...func(*ec2.Options)) (*ec2.DescribeSecurityGroupsOutput, error)
	DescribeSnapshotAttribute(ctx context.Context, params *ec2.DescribeSnapshotAttributeInput, optFns ...func(*ec2.Options)) (*ec2.DescribeSnapshotAttributeOutput, error)
	DescribeSnapshots(ctx context.Context, params *ec2.DescribeSnapshotsInput, optFns ...func(*ec2.Options)) (*ec2.DescribeSnapshotsOutput, error)
	DescribeSubnets(ctx context.Context, params *ec2.DescribeSubnetsInput, optFns ...func(*ec2.Options)) (*ec2.DescribeSubnetsOutput, error)
	DescribeTransitGatewayAttachments(ctx context.Context, params *ec2.DescribeTransitGatewayAttachmentsInput, optFns ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayAttachmentsOutput, error)
	DescribeTransitGatewayMulticastDomains(ctx context.Context, params *ec2.DescribeTransitGatewayMulticastDomainsInput, optFns ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayMulticastDomainsOutput, error)
	DescribeTransitGatewayPeeringAttachments(ctx context.Context, params *ec2.DescribeTransitGatewayPeeringAttachmentsInput, optFns ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayPeeringAttachmentsOutput, error)
	DescribeTransitGatewayRouteTables(ctx context.Context, params *ec2.DescribeTransitGatewayRouteTablesInput, optFns ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayRouteTablesOutput, error)
	DescribeTransitGateways(ctx context.Context, params *ec2.DescribeTransitGatewaysInput, optFns ...func(*ec2.Options)) (*ec2.DescribeTransitGatewaysOutput, error)
	DescribeTransitGatewayVpcAttachments(ctx context.Context, params *ec2.DescribeTransitGatewayVpcAttachmentsInput, optFns ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayVpcAttachmentsOutput, error)
	DescribeVolumes(ctx context.Context, params *ec2.DescribeVolumesInput, optFns ...func(*ec2.Options)) (*ec2.DescribeVolumesOutput, error)
	DescribeVpcEndpoints(ctx context.Context, params *ec2.DescribeVpcEndpointsInput, optFns ...func(*ec2.Options)) (*ec2.DescribeVpcEndpointsOutput, error)
	DescribeVpcEndpointServiceConfigurations(ctx context.Context, params *ec2.DescribeVpcEndpointServiceConfigurationsInput, optFns ...func(*ec2.Options)) (*ec2.DescribeVpcEndpointServiceConfigurationsOutput, error)
	DescribeVpcEndpointServices(ctx context.Context, params *ec2.DescribeVpcEndpointServicesInput, optFns ...func(*ec2.Options)) (*ec2.DescribeVpcEndpointServicesOutput, error)
	DescribeVpcPeeringConnections(ctx context.Context, params *ec2.DescribeVpcPeeringConnectionsInput, optFns ...func(*ec2.Options)) (*ec2.DescribeVpcPeeringConnectionsOutput, error)
	DescribeVpcs(ctx context.Context, params *ec2.DescribeVpcsInput, optFns ...func(*ec2.Options)) (*ec2.DescribeVpcsOutput, error)
	DescribeVpnGateways(ctx context.Context, params *ec2.DescribeVpnGatewaysInput, optFns ...func(*ec2.Options)) (*ec2.DescribeVpnGatewaysOutput, error)
	GetEbsDefaultKmsKeyId(ctx context.Context, params *ec2.GetEbsDefaultKmsKeyIdInput, optFns ...func(*ec2.Options)) (*ec2.GetEbsDefaultKmsKeyIdOutput, error)
	GetEbsEncryptionByDefault(ctx context.Context, params *ec2.GetEbsEncryptionByDefaultInput, optFns ...func(*ec2.Options)) (*ec2.GetEbsEncryptionByDefaultOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_ecr.go . EcrClient
type EcrClient interface {
	DescribeRegistry(ctx context.Context, params *ecr.DescribeRegistryInput, optFns ...func(*ecr.Options)) (*ecr.DescribeRegistryOutput, error)
	DescribeRepositories(ctx context.Context, params *ecr.DescribeRepositoriesInput, optFns ...func(*ecr.Options)) (*ecr.DescribeRepositoriesOutput, error)
	DescribeImages(ctx context.Context, params *ecr.DescribeImagesInput, optFns ...func(*ecr.Options)) (*ecr.DescribeImagesOutput, error)
	GetRegistryPolicy(ctx context.Context, params *ecr.GetRegistryPolicyInput, optFns ...func(*ecr.Options)) (*ecr.GetRegistryPolicyOutput, error)
	ListTagsForResource(ctx context.Context, params *ecr.ListTagsForResourceInput, optFns ...func(*ecr.Options)) (*ecr.ListTagsForResourceOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_ecrpublic.go . EcrPublicClient
type EcrPublicClient interface {
	DescribeImageTags(ctx context.Context, params *ecrpublic.DescribeImageTagsInput, optFns ...func(*ecrpublic.Options)) (*ecrpublic.DescribeImageTagsOutput, error)
	DescribeImages(ctx context.Context, params *ecrpublic.DescribeImagesInput, optFns ...func(*ecrpublic.Options)) (*ecrpublic.DescribeImagesOutput, error)
	DescribeRegistries(ctx context.Context, params *ecrpublic.DescribeRegistriesInput, optFns ...func(*ecrpublic.Options)) (*ecrpublic.DescribeRegistriesOutput, error)
	DescribeRepositories(ctx context.Context, params *ecrpublic.DescribeRepositoriesInput, optFns ...func(*ecrpublic.Options)) (*ecrpublic.DescribeRepositoriesOutput, error)
	GetRepositoryPolicy(ctx context.Context, params *ecrpublic.GetRepositoryPolicyInput, optFns ...func(*ecrpublic.Options)) (*ecrpublic.GetRepositoryPolicyOutput, error)
	ListTagsForResource(ctx context.Context, params *ecrpublic.ListTagsForResourceInput, optFns ...func(*ecrpublic.Options)) (*ecrpublic.ListTagsForResourceOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_ecs.go . EcsClient
type EcsClient interface {
	DescribeClusters(ctx context.Context, params *ecs.DescribeClustersInput, optFns ...func(*ecs.Options)) (*ecs.DescribeClustersOutput, error)
	ListClusters(ctx context.Context, params *ecs.ListClustersInput, optFns ...func(*ecs.Options)) (*ecs.ListClustersOutput, error)
	DescribeServices(ctx context.Context, params *ecs.DescribeServicesInput, optFns ...func(*ecs.Options)) (*ecs.DescribeServicesOutput, error)
	DescribeContainerInstances(ctx context.Context, params *ecs.DescribeContainerInstancesInput, optFns ...func(*ecs.Options)) (*ecs.DescribeContainerInstancesOutput, error)
	ListServices(ctx context.Context, params *ecs.ListServicesInput, optFns ...func(*ecs.Options)) (*ecs.ListServicesOutput, error)
	ListContainerInstances(ctx context.Context, params *ecs.ListContainerInstancesInput, optFns ...func(*ecs.Options)) (*ecs.ListContainerInstancesOutput, error)
	ListTaskDefinitions(ctx context.Context, params *ecs.ListTaskDefinitionsInput, optFns ...func(*ecs.Options)) (*ecs.ListTaskDefinitionsOutput, error)
	DescribeTaskDefinition(ctx context.Context, params *ecs.DescribeTaskDefinitionInput, optFns ...func(*ecs.Options)) (*ecs.DescribeTaskDefinitionOutput, error)
	ListTasks(ctx context.Context, params *ecs.ListTasksInput, optFns ...func(*ecs.Options)) (*ecs.ListTasksOutput, error)
	DescribeTasks(ctx context.Context, params *ecs.DescribeTasksInput, optFns ...func(*ecs.Options)) (*ecs.DescribeTasksOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_efs.go . EfsClient
type EfsClient interface {
	DescribeFileSystems(ctx context.Context, params *efs.DescribeFileSystemsInput, optFns ...func(*efs.Options)) (*efs.DescribeFileSystemsOutput, error)
	DescribeBackupPolicy(ctx context.Context, params *efs.DescribeBackupPolicyInput, optFns ...func(*efs.Options)) (*efs.DescribeBackupPolicyOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_eks.go . EksClient
type EksClient interface {
	ListClusters(ctx context.Context, params *eks.ListClustersInput, optFns ...func(*eks.Options)) (*eks.ListClustersOutput, error)
	DescribeCluster(ctx context.Context, params *eks.DescribeClusterInput, optFns ...func(*eks.Options)) (*eks.DescribeClusterOutput, error)
}

// go:generate mockgen -package=mocks -destination=./mocks/mock_elasticache.go . ElastiCache
type ElastiCache interface {
	DescribeCacheClusters(ctx context.Context, params *elasticache.DescribeCacheClustersInput, optFns ...func(*elasticache.Options)) (*elasticache.DescribeCacheClustersOutput, error)
	DescribeCacheEngineVersions(ctx context.Context, params *elasticache.DescribeCacheEngineVersionsInput, optFns ...func(*elasticache.Options)) (*elasticache.DescribeCacheEngineVersionsOutput, error)
	DescribeCacheParameterGroups(ctx context.Context, params *elasticache.DescribeCacheParameterGroupsInput, optFns ...func(*elasticache.Options)) (*elasticache.DescribeCacheParameterGroupsOutput, error)
	DescribeCacheParameters(ctx context.Context, params *elasticache.DescribeCacheParametersInput, optFns ...func(*elasticache.Options)) (*elasticache.DescribeCacheParametersOutput, error)
	DescribeCacheSubnetGroups(ctx context.Context, params *elasticache.DescribeCacheSubnetGroupsInput, optFns ...func(*elasticache.Options)) (*elasticache.DescribeCacheSubnetGroupsOutput, error)
	DescribeGlobalReplicationGroups(ctx context.Context, params *elasticache.DescribeGlobalReplicationGroupsInput, optFns ...func(*elasticache.Options)) (*elasticache.DescribeGlobalReplicationGroupsOutput, error)
	DescribeReplicationGroups(ctx context.Context, params *elasticache.DescribeReplicationGroupsInput, optFns ...func(*elasticache.Options)) (*elasticache.DescribeReplicationGroupsOutput, error)
	DescribeReservedCacheNodes(ctx context.Context, params *elasticache.DescribeReservedCacheNodesInput, optFns ...func(*elasticache.Options)) (*elasticache.DescribeReservedCacheNodesOutput, error)
	DescribeReservedCacheNodesOfferings(ctx context.Context, params *elasticache.DescribeReservedCacheNodesOfferingsInput, optFns ...func(*elasticache.Options)) (*elasticache.DescribeReservedCacheNodesOfferingsOutput, error)
	DescribeServiceUpdates(ctx context.Context, params *elasticache.DescribeServiceUpdatesInput, optFns ...func(*elasticache.Options)) (*elasticache.DescribeServiceUpdatesOutput, error)
	DescribeSnapshots(ctx context.Context, params *elasticache.DescribeSnapshotsInput, optFns ...func(*elasticache.Options)) (*elasticache.DescribeSnapshotsOutput, error)
	DescribeUserGroups(ctx context.Context, params *elasticache.DescribeUserGroupsInput, optFns ...func(*elasticache.Options)) (*elasticache.DescribeUserGroupsOutput, error)
	DescribeUsers(ctx context.Context, params *elasticache.DescribeUsersInput, optFns ...func(*elasticache.Options)) (*elasticache.DescribeUsersOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_elasticbeanstalk.go . ElasticbeanstalkClient
type ElasticbeanstalkClient interface {
	DescribeConfigurationOptions(ctx context.Context, params *elasticbeanstalk.DescribeConfigurationOptionsInput, optFns ...func(*elasticbeanstalk.Options)) (*elasticbeanstalk.DescribeConfigurationOptionsOutput, error)
	DescribeConfigurationSettings(ctx context.Context, params *elasticbeanstalk.DescribeConfigurationSettingsInput, optFns ...func(*elasticbeanstalk.Options)) (*elasticbeanstalk.DescribeConfigurationSettingsOutput, error)
	DescribeApplications(ctx context.Context, params *elasticbeanstalk.DescribeApplicationsInput, optFns ...func(*elasticbeanstalk.Options)) (*elasticbeanstalk.DescribeApplicationsOutput, error)
	DescribeApplicationVersions(ctx context.Context, params *elasticbeanstalk.DescribeApplicationVersionsInput, optFns ...func(*elasticbeanstalk.Options)) (*elasticbeanstalk.DescribeApplicationVersionsOutput, error)
	DescribeEnvironments(ctx context.Context, params *elasticbeanstalk.DescribeEnvironmentsInput, optFns ...func(*elasticbeanstalk.Options)) (*elasticbeanstalk.DescribeEnvironmentsOutput, error)
	ListTagsForResource(ctx context.Context, params *elasticbeanstalk.ListTagsForResourceInput, optFns ...func(*elasticbeanstalk.Options)) (*elasticbeanstalk.ListTagsForResourceOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_elasticsearch.go . ElasticSearch
type ElasticSearch interface {
	ListDomainNames(ctx context.Context, params *elasticsearchservice.ListDomainNamesInput, optFns ...func(*elasticsearchservice.Options)) (*elasticsearchservice.ListDomainNamesOutput, error)
	DescribeElasticsearchDomain(ctx context.Context, params *elasticsearchservice.DescribeElasticsearchDomainInput, optFns ...func(*elasticsearchservice.Options)) (*elasticsearchservice.DescribeElasticsearchDomainOutput, error)
	ListTags(ctx context.Context, params *elasticsearchservice.ListTagsInput, optFns ...func(*elasticsearchservice.Options)) (*elasticsearchservice.ListTagsOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_elbv1.go . ElbV1Client
type ElbV1Client interface {
	DescribeLoadBalancers(ctx context.Context, params *elbv1.DescribeLoadBalancersInput, optFns ...func(*elbv1.Options)) (*elbv1.DescribeLoadBalancersOutput, error)
	DescribeLoadBalancerPolicies(ctx context.Context, params *elbv1.DescribeLoadBalancerPoliciesInput, optFns ...func(*elbv1.Options)) (*elbv1.DescribeLoadBalancerPoliciesOutput, error)
	DescribeTags(ctx context.Context, params *elbv1.DescribeTagsInput, optFns ...func(*elbv1.Options)) (*elbv1.DescribeTagsOutput, error)
	DescribeLoadBalancerAttributes(ctx context.Context, params *elbv1.DescribeLoadBalancerAttributesInput, optFns ...func(*elbv1.Options)) (*elbv1.DescribeLoadBalancerAttributesOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_elbv2.go . ElbV2Client
type ElbV2Client interface {
	DescribeListenerCertificates(ctx context.Context, params *elbv2.DescribeListenerCertificatesInput, optFns ...func(*elbv2.Options)) (*elbv2.DescribeListenerCertificatesOutput, error)
	DescribeListeners(ctx context.Context, params *elbv2.DescribeListenersInput, optFns ...func(*elbv2.Options)) (*elbv2.DescribeListenersOutput, error)
	DescribeLoadBalancers(ctx context.Context, params *elbv2.DescribeLoadBalancersInput, optFns ...func(*elbv2.Options)) (*elbv2.DescribeLoadBalancersOutput, error)
	DescribeLoadBalancerAttributes(ctx context.Context, params *elbv2.DescribeLoadBalancerAttributesInput, optFns ...func(*elbv2.Options)) (*elbv2.DescribeLoadBalancerAttributesOutput, error)
	DescribeTargetGroups(ctx context.Context, params *elbv2.DescribeTargetGroupsInput, optFns ...func(*elbv2.Options)) (*elbv2.DescribeTargetGroupsOutput, error)
	DescribeTargetHealth(ctx context.Context, params *elbv2.DescribeTargetHealthInput, optFns ...func(*elbv2.Options)) (*elbv2.DescribeTargetHealthOutput, error)
	DescribeTags(ctx context.Context, params *elbv2.DescribeTagsInput, optFns ...func(*elbv2.Options)) (*elbv2.DescribeTagsOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_emr.go . EmrClient
type EmrClient interface {
	DescribeCluster(ctx context.Context, params *emr.DescribeClusterInput, optFns ...func(*emr.Options)) (*emr.DescribeClusterOutput, error)
	GetBlockPublicAccessConfiguration(ctx context.Context, params *emr.GetBlockPublicAccessConfigurationInput, optFns ...func(*emr.Options)) (*emr.GetBlockPublicAccessConfigurationOutput, error)
	ListClusters(ctx context.Context, params *emr.ListClustersInput, optFns ...func(*emr.Options)) (*emr.ListClustersOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_eventbridge.go . EventBridgeClient
type EventBridgeClient interface {
	ListApiDestinations(ctx context.Context, params *eventbridge.ListApiDestinationsInput, optFns ...func(*eventbridge.Options)) (*eventbridge.ListApiDestinationsOutput, error)
	ListArchives(ctx context.Context, params *eventbridge.ListArchivesInput, optFns ...func(*eventbridge.Options)) (*eventbridge.ListArchivesOutput, error)
	ListConnections(ctx context.Context, params *eventbridge.ListConnectionsInput, optFns ...func(*eventbridge.Options)) (*eventbridge.ListConnectionsOutput, error)
	ListEndpoints(ctx context.Context, params *eventbridge.ListEndpointsInput, optFns ...func(*eventbridge.Options)) (*eventbridge.ListEndpointsOutput, error)
	ListEventBuses(ctx context.Context, params *eventbridge.ListEventBusesInput, optFns ...func(*eventbridge.Options)) (*eventbridge.ListEventBusesOutput, error)
	ListEventSources(ctx context.Context, params *eventbridge.ListEventSourcesInput, optFns ...func(*eventbridge.Options)) (*eventbridge.ListEventSourcesOutput, error)
	ListReplays(ctx context.Context, params *eventbridge.ListReplaysInput, optFns ...func(*eventbridge.Options)) (*eventbridge.ListReplaysOutput, error)
	ListRules(ctx context.Context, params *eventbridge.ListRulesInput, optFns ...func(*eventbridge.Options)) (*eventbridge.ListRulesOutput, error)
	ListTagsForResource(ctx context.Context, params *eventbridge.ListTagsForResourceInput, optFns ...func(*eventbridge.Options)) (*eventbridge.ListTagsForResourceOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/firehose.go . FirehoseClient
type FirehoseClient interface {
	DescribeDeliveryStream(ctx context.Context, params *firehose.DescribeDeliveryStreamInput, optFns ...func(*firehose.Options)) (*firehose.DescribeDeliveryStreamOutput, error)
	ListDeliveryStreams(ctx context.Context, params *firehose.ListDeliveryStreamsInput, optFns ...func(*firehose.Options)) (*firehose.ListDeliveryStreamsOutput, error)
	ListTagsForDeliveryStream(ctx context.Context, params *firehose.ListTagsForDeliveryStreamInput, optFns ...func(*firehose.Options)) (*firehose.ListTagsForDeliveryStreamOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/frauddetector.go . FraudDetectorClient
type FraudDetectorClient interface {
	frauddetector.DescribeModelVersionsAPIClient
	frauddetector.GetBatchImportJobsAPIClient
	frauddetector.GetBatchPredictionJobsAPIClient
	frauddetector.GetDetectorsAPIClient
	frauddetector.GetEntityTypesAPIClient
	frauddetector.GetEventTypesAPIClient
	frauddetector.GetExternalModelsAPIClient
	frauddetector.GetLabelsAPIClient
	frauddetector.GetModelsAPIClient
	frauddetector.GetOutcomesAPIClient
	frauddetector.GetRulesAPIClient
	frauddetector.GetVariablesAPIClient
	frauddetector.ListTagsForResourceAPIClient
}

//go:generate mockgen -package=mocks -destination=./mocks/fsx.go . FsxClient
type FsxClient interface {
	DescribeBackups(ctx context.Context, params *fsx.DescribeBackupsInput, optFns ...func(*fsx.Options)) (*fsx.DescribeBackupsOutput, error)
	DescribeDataRepositoryAssociations(ctx context.Context, params *fsx.DescribeDataRepositoryAssociationsInput, optFns ...func(*fsx.Options)) (*fsx.DescribeDataRepositoryAssociationsOutput, error)
	DescribeDataRepositoryTasks(ctx context.Context, params *fsx.DescribeDataRepositoryTasksInput, optFns ...func(*fsx.Options)) (*fsx.DescribeDataRepositoryTasksOutput, error)
	DescribeFileSystems(ctx context.Context, params *fsx.DescribeFileSystemsInput, optFns ...func(*fsx.Options)) (*fsx.DescribeFileSystemsOutput, error)
	DescribeSnapshots(ctx context.Context, params *fsx.DescribeSnapshotsInput, optFns ...func(*fsx.Options)) (*fsx.DescribeSnapshotsOutput, error)
	DescribeStorageVirtualMachines(ctx context.Context, params *fsx.DescribeStorageVirtualMachinesInput, optFns ...func(*fsx.Options)) (*fsx.DescribeStorageVirtualMachinesOutput, error)
	DescribeVolumes(ctx context.Context, params *fsx.DescribeVolumesInput, optFns ...func(*fsx.Options)) (*fsx.DescribeVolumesOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/glacier.go . GlacierClient
type GlacierClient interface {
	ListVaults(ctx context.Context, params *glacier.ListVaultsInput, optFns ...func(*glacier.Options)) (*glacier.ListVaultsOutput, error)
	ListTagsForVault(ctx context.Context, params *glacier.ListTagsForVaultInput, optFns ...func(*glacier.Options)) (*glacier.ListTagsForVaultOutput, error)
	GetVaultAccessPolicy(ctx context.Context, params *glacier.GetVaultAccessPolicyInput, optFns ...func(*glacier.Options)) (*glacier.GetVaultAccessPolicyOutput, error)
	GetVaultLock(ctx context.Context, params *glacier.GetVaultLockInput, optFns ...func(*glacier.Options)) (*glacier.GetVaultLockOutput, error)
	GetVaultNotifications(ctx context.Context, params *glacier.GetVaultNotificationsInput, optFns ...func(*glacier.Options)) (*glacier.GetVaultNotificationsOutput, error)
	GetDataRetrievalPolicy(ctx context.Context, params *glacier.GetDataRetrievalPolicyInput, optFns ...func(*glacier.Options)) (*glacier.GetDataRetrievalPolicyOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/glue.go . GlueClient
type GlueClient interface {
	GetClassifiers(ctx context.Context, params *glue.GetClassifiersInput, optFns ...func(*glue.Options)) (*glue.GetClassifiersOutput, error)
	GetConnections(ctx context.Context, params *glue.GetConnectionsInput, optFns ...func(*glue.Options)) (*glue.GetConnectionsOutput, error)
	GetCrawlers(ctx context.Context, params *glue.GetCrawlersInput, optFns ...func(*glue.Options)) (*glue.GetCrawlersOutput, error)
	GetDatabases(ctx context.Context, params *glue.GetDatabasesInput, optFns ...func(*glue.Options)) (*glue.GetDatabasesOutput, error)
	GetDataCatalogEncryptionSettings(ctx context.Context, params *glue.GetDataCatalogEncryptionSettingsInput, optFns ...func(*glue.Options)) (*glue.GetDataCatalogEncryptionSettingsOutput, error)
	GetDevEndpoints(ctx context.Context, params *glue.GetDevEndpointsInput, optFns ...func(*glue.Options)) (*glue.GetDevEndpointsOutput, error)
	GetJobRuns(ctx context.Context, params *glue.GetJobRunsInput, optFns ...func(*glue.Options)) (*glue.GetJobRunsOutput, error)
	GetJobs(ctx context.Context, params *glue.GetJobsInput, optFns ...func(*glue.Options)) (*glue.GetJobsOutput, error)
	GetMLTaskRuns(ctx context.Context, params *glue.GetMLTaskRunsInput, optFns ...func(*glue.Options)) (*glue.GetMLTaskRunsOutput, error)
	GetMLTransforms(ctx context.Context, params *glue.GetMLTransformsInput, optFns ...func(*glue.Options)) (*glue.GetMLTransformsOutput, error)
	GetPartitionIndexes(ctx context.Context, params *glue.GetPartitionIndexesInput, optFns ...func(*glue.Options)) (*glue.GetPartitionIndexesOutput, error)
	GetSchema(ctx context.Context, params *glue.GetSchemaInput, optFns ...func(*glue.Options)) (*glue.GetSchemaOutput, error)
	GetSchemaVersion(ctx context.Context, params *glue.GetSchemaVersionInput, optFns ...func(*glue.Options)) (*glue.GetSchemaVersionOutput, error)
	GetSecurityConfigurations(ctx context.Context, params *glue.GetSecurityConfigurationsInput, optFns ...func(*glue.Options)) (*glue.GetSecurityConfigurationsOutput, error)
	GetTables(ctx context.Context, params *glue.GetTablesInput, optFns ...func(*glue.Options)) (*glue.GetTablesOutput, error)
	GetTags(ctx context.Context, params *glue.GetTagsInput, optFns ...func(*glue.Options)) (*glue.GetTagsOutput, error)
	GetTrigger(ctx context.Context, params *glue.GetTriggerInput, optFns ...func(*glue.Options)) (*glue.GetTriggerOutput, error)
	GetWorkflow(ctx context.Context, params *glue.GetWorkflowInput, optFns ...func(*glue.Options)) (*glue.GetWorkflowOutput, error)
	ListRegistries(ctx context.Context, params *glue.ListRegistriesInput, optFns ...func(*glue.Options)) (*glue.ListRegistriesOutput, error)
	ListSchemas(ctx context.Context, params *glue.ListSchemasInput, optFns ...func(*glue.Options)) (*glue.ListSchemasOutput, error)
	ListSchemaVersions(ctx context.Context, params *glue.ListSchemaVersionsInput, optFns ...func(*glue.Options)) (*glue.ListSchemaVersionsOutput, error)
	ListTriggers(ctx context.Context, params *glue.ListTriggersInput, optFns ...func(*glue.Options)) (*glue.ListTriggersOutput, error)
	ListWorkflows(ctx context.Context, params *glue.ListWorkflowsInput, optFns ...func(*glue.Options)) (*glue.ListWorkflowsOutput, error)
	QuerySchemaVersionMetadata(ctx context.Context, params *glue.QuerySchemaVersionMetadataInput, optFns ...func(*glue.Options)) (*glue.QuerySchemaVersionMetadataOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_guardduty.go . GuardDutyClient
type GuardDutyClient interface {
	guardduty.ListDetectorsAPIClient
	guardduty.ListMembersAPIClient
	GetDetector(ctx context.Context, params *guardduty.GetDetectorInput, optFns ...func(*guardduty.Options)) (*guardduty.GetDetectorOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_iam.go . IamClient
type IamClient interface {
	GenerateCredentialReport(ctx context.Context, params *iam.GenerateCredentialReportInput, optFns ...func(*iam.Options)) (*iam.GenerateCredentialReportOutput, error)
	GetAccessKeyLastUsed(ctx context.Context, params *iam.GetAccessKeyLastUsedInput, optFns ...func(*iam.Options)) (*iam.GetAccessKeyLastUsedOutput, error)
	GetAccountAuthorizationDetails(context.Context, *iam.GetAccountAuthorizationDetailsInput, ...func(*iam.Options)) (*iam.GetAccountAuthorizationDetailsOutput, error)
	GetAccountPasswordPolicy(ctx context.Context, params *iam.GetAccountPasswordPolicyInput, optFns ...func(*iam.Options)) (*iam.GetAccountPasswordPolicyOutput, error)
	GetCredentialReport(ctx context.Context, params *iam.GetCredentialReportInput, optFns ...func(*iam.Options)) (*iam.GetCredentialReportOutput, error)
	GetGroupPolicy(ctx context.Context, params *iam.GetGroupPolicyInput, optFns ...func(*iam.Options)) (*iam.GetGroupPolicyOutput, error)
	GetOpenIDConnectProvider(ctx context.Context, params *iam.GetOpenIDConnectProviderInput, optFns ...func(*iam.Options)) (*iam.GetOpenIDConnectProviderOutput, error)
	GetRole(ctx context.Context, params *iam.GetRoleInput, optFns ...func(*iam.Options)) (*iam.GetRoleOutput, error)
	GetRolePolicy(ctx context.Context, params *iam.GetRolePolicyInput, optFns ...func(*iam.Options)) (*iam.GetRolePolicyOutput, error)
	GetSAMLProvider(ctx context.Context, params *iam.GetSAMLProviderInput, optFns ...func(*iam.Options)) (*iam.GetSAMLProviderOutput, error)
	GetUser(ctx context.Context, params *iam.GetUserInput, optFns ...func(*iam.Options)) (*iam.GetUserOutput, error)
	GetUserPolicy(ctx context.Context, params *iam.GetUserPolicyInput, optFns ...func(*iam.Options)) (*iam.GetUserPolicyOutput, error)
	ListAccessKeys(ctx context.Context, params *iam.ListAccessKeysInput, optFns ...func(*iam.Options)) (*iam.ListAccessKeysOutput, error)
	ListAttachedGroupPolicies(ctx context.Context, params *iam.ListAttachedGroupPoliciesInput, optFns ...func(*iam.Options)) (*iam.ListAttachedGroupPoliciesOutput, error)
	ListAttachedRolePolicies(ctx context.Context, params *iam.ListAttachedRolePoliciesInput, optFns ...func(*iam.Options)) (*iam.ListAttachedRolePoliciesOutput, error)
	ListAttachedUserPolicies(ctx context.Context, params *iam.ListAttachedUserPoliciesInput, optFns ...func(*iam.Options)) (*iam.ListAttachedUserPoliciesOutput, error)
	ListGroupPolicies(ctx context.Context, params *iam.ListGroupPoliciesInput, optFns ...func(*iam.Options)) (*iam.ListGroupPoliciesOutput, error)
	ListGroups(ctx context.Context, params *iam.ListGroupsInput, optFns ...func(*iam.Options)) (*iam.ListGroupsOutput, error)
	ListGroupsForUser(ctx context.Context, params *iam.ListGroupsForUserInput, optFns ...func(*iam.Options)) (*iam.ListGroupsForUserOutput, error)
	ListOpenIDConnectProviders(ctx context.Context, params *iam.ListOpenIDConnectProvidersInput, optFns ...func(*iam.Options)) (*iam.ListOpenIDConnectProvidersOutput, error)
	ListPolicyTags(ctx context.Context, params *iam.ListPolicyTagsInput, optFns ...func(*iam.Options)) (*iam.ListPolicyTagsOutput, error)
	ListRolePolicies(ctx context.Context, params *iam.ListRolePoliciesInput, optFns ...func(*iam.Options)) (*iam.ListRolePoliciesOutput, error)
	ListRoles(ctx context.Context, params *iam.ListRolesInput, optFns ...func(*iam.Options)) (*iam.ListRolesOutput, error)
	ListSAMLProviders(ctx context.Context, params *iam.ListSAMLProvidersInput, optFns ...func(*iam.Options)) (*iam.ListSAMLProvidersOutput, error)
	ListUserPolicies(ctx context.Context, params *iam.ListUserPoliciesInput, optFns ...func(*iam.Options)) (*iam.ListUserPoliciesOutput, error)
	ListUsers(context.Context, *iam.ListUsersInput, ...func(*iam.Options)) (*iam.ListUsersOutput, error)
	ListVirtualMFADevices(ctx context.Context, params *iam.ListVirtualMFADevicesInput, optFns ...func(*iam.Options)) (*iam.ListVirtualMFADevicesOutput, error)
	iam.ListServerCertificatesAPIClient
	iam.ListAccountAliasesAPIClient
	GetAccountSummary(ctx context.Context, params *iam.GetAccountSummaryInput, optFns ...func(*iam.Options)) (*iam.GetAccountSummaryOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/inspector.go . InspectorClient
type InspectorClient interface {
	ListFindings(ctx context.Context, params *inspector.ListFindingsInput, optFns ...func(*inspector.Options)) (*inspector.ListFindingsOutput, error)
	DescribeFindings(ctx context.Context, params *inspector.DescribeFindingsInput, optFns ...func(*inspector.Options)) (*inspector.DescribeFindingsOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/inspector_v2.go . InspectorV2Client
type InspectorV2Client interface {
	ListFindings(ctx context.Context, params *inspector2.ListFindingsInput, optFns ...func(*inspector2.Options)) (*inspector2.ListFindingsOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_iot.go . IOTClient
type IOTClient interface {
	DescribeBillingGroup(ctx context.Context, params *iot.DescribeBillingGroupInput, optFns ...func(*iot.Options)) (*iot.DescribeBillingGroupOutput, error)
	DescribeCACertificate(ctx context.Context, params *iot.DescribeCACertificateInput, optFns ...func(*iot.Options)) (*iot.DescribeCACertificateOutput, error)
	DescribeCertificate(ctx context.Context, params *iot.DescribeCertificateInput, optFns ...func(*iot.Options)) (*iot.DescribeCertificateOutput, error)
	DescribeJob(ctx context.Context, params *iot.DescribeJobInput, optFns ...func(*iot.Options)) (*iot.DescribeJobOutput, error)
	DescribeSecurityProfile(ctx context.Context, params *iot.DescribeSecurityProfileInput, optFns ...func(*iot.Options)) (*iot.DescribeSecurityProfileOutput, error)
	DescribeStream(ctx context.Context, params *iot.DescribeStreamInput, optFns ...func(*iot.Options)) (*iot.DescribeStreamOutput, error)
	DescribeThingGroup(ctx context.Context, params *iot.DescribeThingGroupInput, optFns ...func(*iot.Options)) (*iot.DescribeThingGroupOutput, error)
	GetPolicy(ctx context.Context, params *iot.GetPolicyInput, optFns ...func(*iot.Options)) (*iot.GetPolicyOutput, error)
	GetTopicRule(ctx context.Context, params *iot.GetTopicRuleInput, optFns ...func(*iot.Options)) (*iot.GetTopicRuleOutput, error)
	ListAttachedPolicies(ctx context.Context, params *iot.ListAttachedPoliciesInput, optFns ...func(*iot.Options)) (*iot.ListAttachedPoliciesOutput, error)
	ListBillingGroups(ctx context.Context, params *iot.ListBillingGroupsInput, optFns ...func(*iot.Options)) (*iot.ListBillingGroupsOutput, error)
	ListCACertificates(ctx context.Context, params *iot.ListCACertificatesInput, optFns ...func(*iot.Options)) (*iot.ListCACertificatesOutput, error)
	ListCertificates(ctx context.Context, params *iot.ListCertificatesInput, optFns ...func(*iot.Options)) (*iot.ListCertificatesOutput, error)
	ListCertificatesByCA(ctx context.Context, params *iot.ListCertificatesByCAInput, optFns ...func(*iot.Options)) (*iot.ListCertificatesByCAOutput, error)
	ListJobs(ctx context.Context, params *iot.ListJobsInput, optFns ...func(*iot.Options)) (*iot.ListJobsOutput, error)
	ListPolicies(ctx context.Context, params *iot.ListPoliciesInput, optFns ...func(*iot.Options)) (*iot.ListPoliciesOutput, error)
	ListSecurityProfiles(ctx context.Context, params *iot.ListSecurityProfilesInput, optFns ...func(*iot.Options)) (*iot.ListSecurityProfilesOutput, error)
	ListStreams(ctx context.Context, params *iot.ListStreamsInput, optFns ...func(*iot.Options)) (*iot.ListStreamsOutput, error)
	ListTagsForResource(ctx context.Context, params *iot.ListTagsForResourceInput, optFns ...func(*iot.Options)) (*iot.ListTagsForResourceOutput, error)
	ListTargetsForSecurityProfile(ctx context.Context, params *iot.ListTargetsForSecurityProfileInput, optFns ...func(*iot.Options)) (*iot.ListTargetsForSecurityProfileOutput, error)
	ListThingGroups(ctx context.Context, params *iot.ListThingGroupsInput, optFns ...func(*iot.Options)) (*iot.ListThingGroupsOutput, error)
	ListThingPrincipals(ctx context.Context, params *iot.ListThingPrincipalsInput, optFns ...func(*iot.Options)) (*iot.ListThingPrincipalsOutput, error)
	ListThings(ctx context.Context, params *iot.ListThingsInput, optFns ...func(*iot.Options)) (*iot.ListThingsOutput, error)
	ListThingsInBillingGroup(ctx context.Context, params *iot.ListThingsInBillingGroupInput, optFns ...func(*iot.Options)) (*iot.ListThingsInBillingGroupOutput, error)
	ListThingsInThingGroup(ctx context.Context, params *iot.ListThingsInThingGroupInput, optFns ...func(*iot.Options)) (*iot.ListThingsInThingGroupOutput, error)
	ListThingTypes(ctx context.Context, params *iot.ListThingTypesInput, optFns ...func(*iot.Options)) (*iot.ListThingTypesOutput, error)
	ListTopicRules(ctx context.Context, params *iot.ListTopicRulesInput, optFns ...func(*iot.Options)) (*iot.ListTopicRulesOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/kinesis.go . KinesisClient
type KinesisClient interface {
	DescribeStreamSummary(ctx context.Context, params *kinesis.DescribeStreamSummaryInput, optFns ...func(*kinesis.Options)) (*kinesis.DescribeStreamSummaryOutput, error)
	ListStreams(ctx context.Context, params *kinesis.ListStreamsInput, optFns ...func(*kinesis.Options)) (*kinesis.ListStreamsOutput, error)
	ListTagsForStream(ctx context.Context, params *kinesis.ListTagsForStreamInput, optFns ...func(*kinesis.Options)) (*kinesis.ListTagsForStreamOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_kms.go . KmsClient
type KmsClient interface {
	DescribeKey(ctx context.Context, params *kms.DescribeKeyInput, optFns ...func(*kms.Options)) (*kms.DescribeKeyOutput, error)
	GetKeyRotationStatus(ctx context.Context, params *kms.GetKeyRotationStatusInput, optFns ...func(*kms.Options)) (*kms.GetKeyRotationStatusOutput, error)
	ListKeys(ctx context.Context, params *kms.ListKeysInput, optFns ...func(*kms.Options)) (*kms.ListKeysOutput, error)
	ListResourceTags(ctx context.Context, params *kms.ListResourceTagsInput, optFns ...func(*kms.Options)) (*kms.ListResourceTagsOutput, error)
	ListAliases(ctx context.Context, params *kms.ListAliasesInput, optFns ...func(*kms.Options)) (*kms.ListAliasesOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_lambda.go . LambdaClient
type LambdaClient interface {
	GetCodeSigningConfig(ctx context.Context, params *lambda.GetCodeSigningConfigInput, optFns ...func(*lambda.Options)) (*lambda.GetCodeSigningConfigOutput, error)
	GetFunction(ctx context.Context, params *lambda.GetFunctionInput, optFns ...func(*lambda.Options)) (*lambda.GetFunctionOutput, error)
	GetFunctionCodeSigningConfig(ctx context.Context, params *lambda.GetFunctionCodeSigningConfigInput, optFns ...func(*lambda.Options)) (*lambda.GetFunctionCodeSigningConfigOutput, error)
	GetFunctionUrlConfig(ctx context.Context, params *lambda.GetFunctionUrlConfigInput, optFns ...func(*lambda.Options)) (*lambda.GetFunctionUrlConfigOutput, error)
	GetLayerVersionPolicy(ctx context.Context, params *lambda.GetLayerVersionPolicyInput, optFns ...func(*lambda.Options)) (*lambda.GetLayerVersionPolicyOutput, error)
	GetPolicy(ctx context.Context, params *lambda.GetPolicyInput, optFns ...func(*lambda.Options)) (*lambda.GetPolicyOutput, error)
	ListAliases(ctx context.Context, params *lambda.ListAliasesInput, optFns ...func(*lambda.Options)) (*lambda.ListAliasesOutput, error)
	ListEventSourceMappings(ctx context.Context, params *lambda.ListEventSourceMappingsInput, optFns ...func(*lambda.Options)) (*lambda.ListEventSourceMappingsOutput, error)
	ListFunctionEventInvokeConfigs(ctx context.Context, params *lambda.ListFunctionEventInvokeConfigsInput, optFns ...func(*lambda.Options)) (*lambda.ListFunctionEventInvokeConfigsOutput, error)
	ListFunctions(ctx context.Context, params *lambda.ListFunctionsInput, optFns ...func(*lambda.Options)) (*lambda.ListFunctionsOutput, error)
	ListLayers(ctx context.Context, params *lambda.ListLayersInput, optFns ...func(*lambda.Options)) (*lambda.ListLayersOutput, error)
	ListLayerVersions(ctx context.Context, params *lambda.ListLayerVersionsInput, optFns ...func(*lambda.Options)) (*lambda.ListLayerVersionsOutput, error)
	ListProvisionedConcurrencyConfigs(ctx context.Context, params *lambda.ListProvisionedConcurrencyConfigsInput, optFns ...func(*lambda.Options)) (*lambda.ListProvisionedConcurrencyConfigsOutput, error)
	ListVersionsByFunction(ctx context.Context, params *lambda.ListVersionsByFunctionInput, optFns ...func(*lambda.Options)) (*lambda.ListVersionsByFunctionOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_lightsail.go . LightsailClient
type LightsailClient interface {
	GetAlarms(ctx context.Context, params *lightsail.GetAlarmsInput, optFns ...func(*lightsail.Options)) (*lightsail.GetAlarmsOutput, error)
	GetBucketAccessKeys(ctx context.Context, params *lightsail.GetBucketAccessKeysInput, optFns ...func(*lightsail.Options)) (*lightsail.GetBucketAccessKeysOutput, error)
	GetBuckets(ctx context.Context, params *lightsail.GetBucketsInput, optFns ...func(*lightsail.Options)) (*lightsail.GetBucketsOutput, error)
	GetCertificates(ctx context.Context, params *lightsail.GetCertificatesInput, optFns ...func(*lightsail.Options)) (*lightsail.GetCertificatesOutput, error)
	GetContainerImages(ctx context.Context, params *lightsail.GetContainerImagesInput, optFns ...func(*lightsail.Options)) (*lightsail.GetContainerImagesOutput, error)
	GetContainerServiceDeployments(ctx context.Context, params *lightsail.GetContainerServiceDeploymentsInput, optFns ...func(*lightsail.Options)) (*lightsail.GetContainerServiceDeploymentsOutput, error)
	GetContainerServices(ctx context.Context, params *lightsail.GetContainerServicesInput, optFns ...func(*lightsail.Options)) (*lightsail.GetContainerServicesOutput, error)
	GetDisks(ctx context.Context, params *lightsail.GetDisksInput, optFns ...func(*lightsail.Options)) (*lightsail.GetDisksOutput, error)
	GetDiskSnapshots(ctx context.Context, params *lightsail.GetDiskSnapshotsInput, optFns ...func(*lightsail.Options)) (*lightsail.GetDiskSnapshotsOutput, error)
	GetDistributionLatestCacheReset(ctx context.Context, params *lightsail.GetDistributionLatestCacheResetInput, optFns ...func(*lightsail.Options)) (*lightsail.GetDistributionLatestCacheResetOutput, error)
	GetDistributions(ctx context.Context, params *lightsail.GetDistributionsInput, optFns ...func(*lightsail.Options)) (*lightsail.GetDistributionsOutput, error)
	GetInstanceAccessDetails(ctx context.Context, params *lightsail.GetInstanceAccessDetailsInput, optFns ...func(*lightsail.Options)) (*lightsail.GetInstanceAccessDetailsOutput, error)
	GetInstancePortStates(ctx context.Context, params *lightsail.GetInstancePortStatesInput, optFns ...func(*lightsail.Options)) (*lightsail.GetInstancePortStatesOutput, error)
	GetInstances(ctx context.Context, params *lightsail.GetInstancesInput, optFns ...func(*lightsail.Options)) (*lightsail.GetInstancesOutput, error)
	GetInstanceSnapshots(ctx context.Context, params *lightsail.GetInstanceSnapshotsInput, optFns ...func(*lightsail.Options)) (*lightsail.GetInstanceSnapshotsOutput, error)
	GetLoadBalancers(ctx context.Context, params *lightsail.GetLoadBalancersInput, optFns ...func(*lightsail.Options)) (*lightsail.GetLoadBalancersOutput, error)
	GetLoadBalancerTlsCertificates(ctx context.Context, params *lightsail.GetLoadBalancerTlsCertificatesInput, optFns ...func(*lightsail.Options)) (*lightsail.GetLoadBalancerTlsCertificatesOutput, error)
	GetRelationalDatabaseEvents(ctx context.Context, params *lightsail.GetRelationalDatabaseEventsInput, optFns ...func(*lightsail.Options)) (*lightsail.GetRelationalDatabaseEventsOutput, error)
	GetRelationalDatabaseLogEvents(ctx context.Context, params *lightsail.GetRelationalDatabaseLogEventsInput, optFns ...func(*lightsail.Options)) (*lightsail.GetRelationalDatabaseLogEventsOutput, error)
	GetRelationalDatabaseLogStreams(ctx context.Context, params *lightsail.GetRelationalDatabaseLogStreamsInput, optFns ...func(*lightsail.Options)) (*lightsail.GetRelationalDatabaseLogStreamsOutput, error)
	GetRelationalDatabaseParameters(ctx context.Context, params *lightsail.GetRelationalDatabaseParametersInput, optFns ...func(*lightsail.Options)) (*lightsail.GetRelationalDatabaseParametersOutput, error)
	GetRelationalDatabases(ctx context.Context, params *lightsail.GetRelationalDatabasesInput, optFns ...func(*lightsail.Options)) (*lightsail.GetRelationalDatabasesOutput, error)
	GetRelationalDatabaseSnapshots(ctx context.Context, params *lightsail.GetRelationalDatabaseSnapshotsInput, optFns ...func(*lightsail.Options)) (*lightsail.GetRelationalDatabaseSnapshotsOutput, error)
	GetStaticIps(ctx context.Context, params *lightsail.GetStaticIpsInput, optFns ...func(*lightsail.Options)) (*lightsail.GetStaticIpsOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_mq.go . MQClient
type MQClient interface {
	DescribeBroker(ctx context.Context, params *mq.DescribeBrokerInput, optFns ...func(*mq.Options)) (*mq.DescribeBrokerOutput, error)
	DescribeConfiguration(ctx context.Context, params *mq.DescribeConfigurationInput, optFns ...func(*mq.Options)) (*mq.DescribeConfigurationOutput, error)
	DescribeConfigurationRevision(ctx context.Context, params *mq.DescribeConfigurationRevisionInput, optFns ...func(*mq.Options)) (*mq.DescribeConfigurationRevisionOutput, error)
	DescribeUser(ctx context.Context, params *mq.DescribeUserInput, optFns ...func(*mq.Options)) (*mq.DescribeUserOutput, error)
	ListBrokers(ctx context.Context, params *mq.ListBrokersInput, optFns ...func(*mq.Options)) (*mq.ListBrokersOutput, error)
	ListConfigurationRevisions(ctx context.Context, params *mq.ListConfigurationRevisionsInput, optFns ...func(*mq.Options)) (*mq.ListConfigurationRevisionsOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_neptune.go . NeptuneClient
type NeptuneClient interface {
	DescribeDBClusterEndpoints(ctx context.Context, params *neptune.DescribeDBClusterEndpointsInput, optFns ...func(*neptune.Options)) (*neptune.DescribeDBClusterEndpointsOutput, error)
	DescribeDBClusterParameterGroups(ctx context.Context, params *neptune.DescribeDBClusterParameterGroupsInput, optFns ...func(*neptune.Options)) (*neptune.DescribeDBClusterParameterGroupsOutput, error)
	DescribeDBClusterParameters(ctx context.Context, params *neptune.DescribeDBClusterParametersInput, optFns ...func(*neptune.Options)) (*neptune.DescribeDBClusterParametersOutput, error)
	DescribeDBClusterSnapshotAttributes(ctx context.Context, params *neptune.DescribeDBClusterSnapshotAttributesInput, optFns ...func(*neptune.Options)) (*neptune.DescribeDBClusterSnapshotAttributesOutput, error)
	DescribeDBClusterSnapshots(ctx context.Context, params *neptune.DescribeDBClusterSnapshotsInput, optFns ...func(*neptune.Options)) (*neptune.DescribeDBClusterSnapshotsOutput, error)
	DescribeDBClusters(ctx context.Context, params *neptune.DescribeDBClustersInput, optFns ...func(*neptune.Options)) (*neptune.DescribeDBClustersOutput, error)
	DescribeDBInstances(ctx context.Context, params *neptune.DescribeDBInstancesInput, optFns ...func(*neptune.Options)) (*neptune.DescribeDBInstancesOutput, error)
	DescribeDBParameterGroups(ctx context.Context, params *neptune.DescribeDBParameterGroupsInput, optFns ...func(*neptune.Options)) (*neptune.DescribeDBParameterGroupsOutput, error)
	DescribeDBParameters(ctx context.Context, params *neptune.DescribeDBParametersInput, optFns ...func(*neptune.Options)) (*neptune.DescribeDBParametersOutput, error)
	DescribeDBSubnetGroups(ctx context.Context, params *neptune.DescribeDBSubnetGroupsInput, optFns ...func(*neptune.Options)) (*neptune.DescribeDBSubnetGroupsOutput, error)
	DescribeEventSubscriptions(ctx context.Context, params *neptune.DescribeEventSubscriptionsInput, optFns ...func(*neptune.Options)) (*neptune.DescribeEventSubscriptionsOutput, error)
	DescribeGlobalClusters(ctx context.Context, params *neptune.DescribeGlobalClustersInput, optFns ...func(*neptune.Options)) (*neptune.DescribeGlobalClustersOutput, error)
	ListTagsForResource(ctx context.Context, params *neptune.ListTagsForResourceInput, optFns ...func(*neptune.Options)) (*neptune.ListTagsForResourceOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_organizations.go . OrganizationsClient
type OrganizationsClient interface {
	ListAccounts(ctx context.Context, params *organizations.ListAccountsInput, optFns ...func(*organizations.Options)) (*organizations.ListAccountsOutput, error)
	ListAccountsForParent(ctx context.Context, params *organizations.ListAccountsForParentInput, optFns ...func(*organizations.Options)) (*organizations.ListAccountsForParentOutput, error)
	organizations.ListTagsForResourceAPIClient
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_qldb.go . QLDBClient
type QLDBClient interface {
	qldb.ListLedgersAPIClient
	qldb.ListJournalKinesisStreamsForLedgerAPIClient
	qldb.ListJournalS3ExportsForLedgerAPIClient
	DescribeLedger(ctx context.Context, params *qldb.DescribeLedgerInput, optFns ...func(*qldb.Options)) (*qldb.DescribeLedgerOutput, error)
	ListTagsForResource(ctx context.Context, params *qldb.ListTagsForResourceInput, optFns ...func(*qldb.Options)) (*qldb.ListTagsForResourceOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_rds.go . RdsClient
type RdsClient interface {
	DescribeCertificates(ctx context.Context, params *rds.DescribeCertificatesInput, optFns ...func(*rds.Options)) (*rds.DescribeCertificatesOutput, error)
	DescribeDBClusterParameterGroups(ctx context.Context, params *rds.DescribeDBClusterParameterGroupsInput, optFns ...func(*rds.Options)) (*rds.DescribeDBClusterParameterGroupsOutput, error)
	DescribeDBClusterParameters(ctx context.Context, params *rds.DescribeDBClusterParametersInput, optFns ...func(*rds.Options)) (*rds.DescribeDBClusterParametersOutput, error)
	DescribeDBClusters(ctx context.Context, params *rds.DescribeDBClustersInput, optFns ...func(*rds.Options)) (*rds.DescribeDBClustersOutput, error)
	DescribeDBClusterSnapshotAttributes(ctx context.Context, params *rds.DescribeDBClusterSnapshotAttributesInput, optFns ...func(*rds.Options)) (*rds.DescribeDBClusterSnapshotAttributesOutput, error)
	DescribeDBClusterSnapshots(ctx context.Context, params *rds.DescribeDBClusterSnapshotsInput, optFns ...func(*rds.Options)) (*rds.DescribeDBClusterSnapshotsOutput, error)
	DescribeDBInstances(ctx context.Context, params *rds.DescribeDBInstancesInput, optFns ...func(*rds.Options)) (*rds.DescribeDBInstancesOutput, error)
	DescribeDBParameterGroups(ctx context.Context, params *rds.DescribeDBParameterGroupsInput, optFns ...func(*rds.Options)) (*rds.DescribeDBParameterGroupsOutput, error)
	DescribeDBParameters(ctx context.Context, params *rds.DescribeDBParametersInput, optFns ...func(*rds.Options)) (*rds.DescribeDBParametersOutput, error)
	DescribeDBSecurityGroups(ctx context.Context, params *rds.DescribeDBSecurityGroupsInput, optFns ...func(*rds.Options)) (*rds.DescribeDBSecurityGroupsOutput, error)
	DescribeDBSnapshotAttributes(ctx context.Context, params *rds.DescribeDBSnapshotAttributesInput, optFns ...func(*rds.Options)) (*rds.DescribeDBSnapshotAttributesOutput, error)
	DescribeDBSnapshots(ctx context.Context, params *rds.DescribeDBSnapshotsInput, optFns ...func(*rds.Options)) (*rds.DescribeDBSnapshotsOutput, error)
	DescribeDBSubnetGroups(ctx context.Context, params *rds.DescribeDBSubnetGroupsInput, optFns ...func(*rds.Options)) (*rds.DescribeDBSubnetGroupsOutput, error)
	DescribeEventSubscriptions(ctx context.Context, params *rds.DescribeEventSubscriptionsInput, optFns ...func(*rds.Options)) (*rds.DescribeEventSubscriptionsOutput, error)
	ListTagsForResource(ctx context.Context, params *rds.ListTagsForResourceInput, optFns ...func(*rds.Options)) (*rds.ListTagsForResourceOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_redshift.go . RedshiftClient
type RedshiftClient interface {
	DescribeClusterParameters(ctx context.Context, params *redshift.DescribeClusterParametersInput, optFns ...func(*redshift.Options)) (*redshift.DescribeClusterParametersOutput, error)
	DescribeClusters(ctx context.Context, params *redshift.DescribeClustersInput, optFns ...func(*redshift.Options)) (*redshift.DescribeClustersOutput, error)
	DescribeClusterSnapshots(ctx context.Context, params *redshift.DescribeClusterSnapshotsInput, optFns ...func(*redshift.Options)) (*redshift.DescribeClusterSnapshotsOutput, error)
	DescribeClusterSubnetGroups(ctx context.Context, params *redshift.DescribeClusterSubnetGroupsInput, optFns ...func(*redshift.Options)) (*redshift.DescribeClusterSubnetGroupsOutput, error)
	DescribeEventSubscriptions(ctx context.Context, params *redshift.DescribeEventSubscriptionsInput, optFns ...func(*redshift.Options)) (*redshift.DescribeEventSubscriptionsOutput, error)
	DescribeLoggingStatus(ctx context.Context, params *redshift.DescribeLoggingStatusInput, optFns ...func(*redshift.Options)) (*redshift.DescribeLoggingStatusOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_resourcegroups.go . ResourceGroupsClient
type ResourceGroupsClient interface {
	GetGroup(ctx context.Context, params *resourcegroups.GetGroupInput, optFns ...func(*resourcegroups.Options)) (*resourcegroups.GetGroupOutput, error)
	GetGroupQuery(ctx context.Context, params *resourcegroups.GetGroupQueryInput, optFns ...func(*resourcegroups.Options)) (*resourcegroups.GetGroupQueryOutput, error)
	GetTags(ctx context.Context, params *resourcegroups.GetTagsInput, optFns ...func(*resourcegroups.Options)) (*resourcegroups.GetTagsOutput, error)
	ListGroups(ctx context.Context, params *resourcegroups.ListGroupsInput, optFns ...func(*resourcegroups.Options)) (*resourcegroups.ListGroupsOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_route53.go . Route53Client
type Route53Client interface {
	GetHostedZone(ctx context.Context, params *route53.GetHostedZoneInput, optFns ...func(*route53.Options)) (*route53.GetHostedZoneOutput, error)
	GetTrafficPolicy(ctx context.Context, params *route53.GetTrafficPolicyInput, optFns ...func(*route53.Options)) (*route53.GetTrafficPolicyOutput, error)
	ListHealthChecks(ctx context.Context, params *route53.ListHealthChecksInput, optFns ...func(*route53.Options)) (*route53.ListHealthChecksOutput, error)
	ListHostedZones(ctx context.Context, params *route53.ListHostedZonesInput, optFns ...func(*route53.Options)) (*route53.ListHostedZonesOutput, error)
	ListQueryLoggingConfigs(ctx context.Context, params *route53.ListQueryLoggingConfigsInput, optFns ...func(*route53.Options)) (*route53.ListQueryLoggingConfigsOutput, error)
	ListResourceRecordSets(ctx context.Context, params *route53.ListResourceRecordSetsInput, optFns ...func(*route53.Options)) (*route53.ListResourceRecordSetsOutput, error)
	ListReusableDelegationSets(ctx context.Context, params *route53.ListReusableDelegationSetsInput, optFns ...func(*route53.Options)) (*route53.ListReusableDelegationSetsOutput, error)
	ListTagsForResource(ctx context.Context, params *route53.ListTagsForResourceInput, optFns ...func(*route53.Options)) (*route53.ListTagsForResourceOutput, error)
	ListTagsForResources(ctx context.Context, params *route53.ListTagsForResourcesInput, optFns ...func(*route53.Options)) (*route53.ListTagsForResourcesOutput, error)
	ListTrafficPolicies(ctx context.Context, params *route53.ListTrafficPoliciesInput, optFns ...func(*route53.Options)) (*route53.ListTrafficPoliciesOutput, error)
	ListTrafficPolicyInstancesByHostedZone(ctx context.Context, params *route53.ListTrafficPolicyInstancesByHostedZoneInput, optFns ...func(*route53.Options)) (*route53.ListTrafficPolicyInstancesByHostedZoneOutput, error)
	ListTrafficPolicyVersions(ctx context.Context, params *route53.ListTrafficPolicyVersionsInput, optFns ...func(*route53.Options)) (*route53.ListTrafficPolicyVersionsOutput, error)
	ListVPCAssociationAuthorizations(ctx context.Context, params *route53.ListVPCAssociationAuthorizationsInput, optFns ...func(*route53.Options)) (*route53.ListVPCAssociationAuthorizationsOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_route53_domains.go . Route53DomainsClient
type Route53DomainsClient interface {
	GetDomainDetail(ctx context.Context, params *route53domains.GetDomainDetailInput, optFns ...func(*route53domains.Options)) (*route53domains.GetDomainDetailOutput, error)
	ListDomains(ctx context.Context, params *route53domains.ListDomainsInput, optFns ...func(*route53domains.Options)) (*route53domains.ListDomainsOutput, error)
	ListTagsForDomain(ctx context.Context, params *route53domains.ListTagsForDomainInput, optFns ...func(*route53domains.Options)) (*route53domains.ListTagsForDomainOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_s3.go . S3Client
type S3Client interface {
	GetBucketAcl(ctx context.Context, params *s3.GetBucketAclInput, optFns ...func(*s3.Options)) (*s3.GetBucketAclOutput, error)
	GetBucketCors(ctx context.Context, params *s3.GetBucketCorsInput, optFns ...func(*s3.Options)) (*s3.GetBucketCorsOutput, error)
	GetBucketEncryption(ctx context.Context, params *s3.GetBucketEncryptionInput, optFns ...func(*s3.Options)) (*s3.GetBucketEncryptionOutput, error)
	GetBucketLifecycleConfiguration(ctx context.Context, params *s3.GetBucketLifecycleConfigurationInput, optFns ...func(*s3.Options)) (*s3.GetBucketLifecycleConfigurationOutput, error)
	GetBucketLocation(ctx context.Context, params *s3.GetBucketLocationInput, optFns ...func(*s3.Options)) (*s3.GetBucketLocationOutput, error)
	GetBucketLogging(ctx context.Context, params *s3.GetBucketLoggingInput, optFns ...func(*s3.Options)) (*s3.GetBucketLoggingOutput, error)
	GetBucketOwnershipControls(ctx context.Context, params *s3.GetBucketOwnershipControlsInput, optFns ...func(*s3.Options)) (*s3.GetBucketOwnershipControlsOutput, error)
	GetBucketPolicy(ctx context.Context, params *s3.GetBucketPolicyInput, optFns ...func(*s3.Options)) (*s3.GetBucketPolicyOutput, error)
	GetBucketReplication(ctx context.Context, params *s3.GetBucketReplicationInput, optFns ...func(*s3.Options)) (*s3.GetBucketReplicationOutput, error)
	GetBucketTagging(ctx context.Context, params *s3.GetBucketTaggingInput, optFns ...func(*s3.Options)) (*s3.GetBucketTaggingOutput, error)
	GetBucketVersioning(ctx context.Context, params *s3.GetBucketVersioningInput, optFns ...func(*s3.Options)) (*s3.GetBucketVersioningOutput, error)
	GetPublicAccessBlock(ctx context.Context, params *s3.GetPublicAccessBlockInput, optFns ...func(*s3.Options)) (*s3.GetPublicAccessBlockOutput, error)
	ListBuckets(ctx context.Context, params *s3.ListBucketsInput, optFns ...func(*s3.Options)) (*s3.ListBucketsOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_s3Control.go . S3ControlClient
type S3ControlClient interface {
	GetPublicAccessBlock(ctx context.Context, params *s3control.GetPublicAccessBlockInput, optFns ...func(*s3control.Options)) (*s3control.GetPublicAccessBlockOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_s3manager.go . S3ManagerClient
type S3ManagerClient interface {
	GetBucketRegion(ctx context.Context, bucket string, optFns ...func(*s3.Options)) (string, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_sagemaker.go . SageMakerClient
type SageMakerClient interface {
	DescribeEndpointConfig(ctx context.Context, params *sagemaker.DescribeEndpointConfigInput, optFns ...func(*sagemaker.Options)) (*sagemaker.DescribeEndpointConfigOutput, error)
	DescribeModel(ctx context.Context, params *sagemaker.DescribeModelInput, optFns ...func(*sagemaker.Options)) (*sagemaker.DescribeModelOutput, error)
	DescribeNotebookInstance(ctx context.Context, params *sagemaker.DescribeNotebookInstanceInput, optFns ...func(*sagemaker.Options)) (*sagemaker.DescribeNotebookInstanceOutput, error)
	DescribeTrainingJob(ctx context.Context, params *sagemaker.DescribeTrainingJobInput, optFns ...func(*sagemaker.Options)) (*sagemaker.DescribeTrainingJobOutput, error)
	ListEndpointConfigs(ctx context.Context, params *sagemaker.ListEndpointConfigsInput, optFns ...func(*sagemaker.Options)) (*sagemaker.ListEndpointConfigsOutput, error)
	ListModels(ctx context.Context, params *sagemaker.ListModelsInput, optFns ...func(*sagemaker.Options)) (*sagemaker.ListModelsOutput, error)
	ListNotebookInstances(ctx context.Context, params *sagemaker.ListNotebookInstancesInput, optFns ...func(*sagemaker.Options)) (*sagemaker.ListNotebookInstancesOutput, error)
	ListTags(ctx context.Context, params *sagemaker.ListTagsInput, optFns ...func(*sagemaker.Options)) (*sagemaker.ListTagsOutput, error)
	ListTrainingJobs(ctx context.Context, params *sagemaker.ListTrainingJobsInput, optFns ...func(*sagemaker.Options)) (*sagemaker.ListTrainingJobsOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_secrets_manager.go . SecretsManagerClient
type SecretsManagerClient interface {
	DescribeSecret(ctx context.Context, params *secretsmanager.DescribeSecretInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.DescribeSecretOutput, error)
	GetResourcePolicy(ctx context.Context, params *secretsmanager.GetResourcePolicyInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.GetResourcePolicyOutput, error)
	ListSecrets(ctx context.Context, params *secretsmanager.ListSecretsInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.ListSecretsOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_service_catalog.go . ServiceCatalogClient
type ServiceCatalogClient interface {
	ListPortfolios(ctx context.Context, input *servicecatalog.ListPortfoliosInput, optFns ...func(*servicecatalog.Options)) (*servicecatalog.ListPortfoliosOutput, error)
	SearchProductsAsAdmin(ctx context.Context, params *servicecatalog.SearchProductsAsAdminInput, optFns ...func(*servicecatalog.Options)) (*servicecatalog.SearchProductsAsAdminOutput, error)
	SearchProvisionedProducts(ctx context.Context, params *servicecatalog.SearchProvisionedProductsInput, optFns ...func(*servicecatalog.Options)) (*servicecatalog.SearchProvisionedProductsOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_service_catalog_app_registry.go . ServiceCatalogAppRegistryClient
type ServiceCatalogAppRegistryClient interface {
	ListTagsForResource(ctx context.Context, params *servicecatalogappregistry.ListTagsForResourceInput, optFns ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.ListTagsForResourceOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/ses.go . SESClient
type SESClient interface {
	GetEmailTemplate(ctx context.Context, params *sesv2.GetEmailTemplateInput, optFns ...func(*sesv2.Options)) (*sesv2.GetEmailTemplateOutput, error)
	ListEmailTemplates(ctx context.Context, params *sesv2.ListEmailTemplatesInput, optFns ...func(*sesv2.Options)) (*sesv2.ListEmailTemplatesOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/shield.go . ShieldClient
type ShieldClient interface {
	DescribeAttack(ctx context.Context, params *shield.DescribeAttackInput, optFns ...func(*shield.Options)) (*shield.DescribeAttackOutput, error)
	DescribeSubscription(ctx context.Context, params *shield.DescribeSubscriptionInput, optFns ...func(*shield.Options)) (*shield.DescribeSubscriptionOutput, error)
	ListAttacks(ctx context.Context, params *shield.ListAttacksInput, optFns ...func(*shield.Options)) (*shield.ListAttacksOutput, error)
	ListProtectionGroups(ctx context.Context, params *shield.ListProtectionGroupsInput, optFns ...func(*shield.Options)) (*shield.ListProtectionGroupsOutput, error)
	ListProtections(ctx context.Context, params *shield.ListProtectionsInput, optFns ...func(*shield.Options)) (*shield.ListProtectionsOutput, error)
	ListTagsForResource(ctx context.Context, params *shield.ListTagsForResourceInput, optFns ...func(*shield.Options)) (*shield.ListTagsForResourceOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/sns.go . SnsClient
type SnsClient interface {
	GetSubscriptionAttributes(ctx context.Context, params *sns.GetSubscriptionAttributesInput, optFns ...func(*sns.Options)) (*sns.GetSubscriptionAttributesOutput, error)
	GetTopicAttributes(ctx context.Context, params *sns.GetTopicAttributesInput, optFns ...func(*sns.Options)) (*sns.GetTopicAttributesOutput, error)
	ListSubscriptions(ctx context.Context, params *sns.ListSubscriptionsInput, optFns ...func(*sns.Options)) (*sns.ListSubscriptionsOutput, error)
	ListTagsForResource(ctx context.Context, params *sns.ListTagsForResourceInput, optFns ...func(*sns.Options)) (*sns.ListTagsForResourceOutput, error)
	ListTopics(ctx context.Context, params *sns.ListTopicsInput, optFns ...func(*sns.Options)) (*sns.ListTopicsOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_sqs.go . SQSClient
type SQSClient interface {
	GetQueueAttributes(ctx context.Context, params *sqs.GetQueueAttributesInput, optFns ...func(*sqs.Options)) (*sqs.GetQueueAttributesOutput, error)
	ListQueues(ctx context.Context, params *sqs.ListQueuesInput, optFns ...func(*sqs.Options)) (*sqs.ListQueuesOutput, error)
	ListQueueTags(ctx context.Context, params *sqs.ListQueueTagsInput, optFns ...func(*sqs.Options)) (*sqs.ListQueueTagsOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_ssm.go . SSMClient
type SSMClient interface {
	DescribeDocument(ctx context.Context, params *ssm.DescribeDocumentInput, optFns ...func(*ssm.Options)) (*ssm.DescribeDocumentOutput, error)
	DescribeDocumentPermission(ctx context.Context, params *ssm.DescribeDocumentPermissionInput, optFns ...func(*ssm.Options)) (*ssm.DescribeDocumentPermissionOutput, error)
	DescribeInstanceInformation(ctx context.Context, params *ssm.DescribeInstanceInformationInput, optFns ...func(*ssm.Options)) (*ssm.DescribeInstanceInformationOutput, error)
	DescribeParameters(ctx context.Context, params *ssm.DescribeParametersInput, optFns ...func(*ssm.Options)) (*ssm.DescribeParametersOutput, error)
	ListComplianceItems(ctx context.Context, params *ssm.ListComplianceItemsInput, optFns ...func(*ssm.Options)) (*ssm.ListComplianceItemsOutput, error)
	ListDocuments(ctx context.Context, params *ssm.ListDocumentsInput, optFns ...func(*ssm.Options)) (*ssm.ListDocumentsOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_waf.go . WafClient
type WafClient interface {
	GetLoggingConfiguration(ctx context.Context, params *waf.GetLoggingConfigurationInput, optFns ...func(*waf.Options)) (*waf.GetLoggingConfigurationOutput, error)
	GetRule(ctx context.Context, params *waf.GetRuleInput, optFns ...func(*waf.Options)) (*waf.GetRuleOutput, error)
	GetRuleGroup(ctx context.Context, params *waf.GetRuleGroupInput, optFns ...func(*waf.Options)) (*waf.GetRuleGroupOutput, error)
	GetWebACL(ctx context.Context, params *waf.GetWebACLInput, optFns ...func(*waf.Options)) (*waf.GetWebACLOutput, error)
	ListActivatedRulesInRuleGroup(ctx context.Context, params *waf.ListActivatedRulesInRuleGroupInput, optFns ...func(*waf.Options)) (*waf.ListActivatedRulesInRuleGroupOutput, error)
	ListRuleGroups(ctx context.Context, params *waf.ListRuleGroupsInput, optFns ...func(*waf.Options)) (*waf.ListRuleGroupsOutput, error)
	ListRules(ctx context.Context, params *waf.ListRulesInput, optFns ...func(*waf.Options)) (*waf.ListRulesOutput, error)
	ListSubscribedRuleGroups(ctx context.Context, params *waf.ListSubscribedRuleGroupsInput, optFns ...func(*waf.Options)) (*waf.ListSubscribedRuleGroupsOutput, error)
	ListTagsForResource(ctx context.Context, params *waf.ListTagsForResourceInput, optFns ...func(*waf.Options)) (*waf.ListTagsForResourceOutput, error)
	ListWebACLs(ctx context.Context, params *waf.ListWebACLsInput, optFns ...func(*waf.Options)) (*waf.ListWebACLsOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/wafregional.go . WafRegionalClient
type WafRegionalClient interface {
	GetRateBasedRule(ctx context.Context, params *wafregional.GetRateBasedRuleInput, optFns ...func(*wafregional.Options)) (*wafregional.GetRateBasedRuleOutput, error)
	GetRule(ctx context.Context, params *wafregional.GetRuleInput, optFns ...func(*wafregional.Options)) (*wafregional.GetRuleOutput, error)
	GetRuleGroup(ctx context.Context, params *wafregional.GetRuleGroupInput, optFns ...func(*wafregional.Options)) (*wafregional.GetRuleGroupOutput, error)
	GetWebACL(ctx context.Context, params *wafregional.GetWebACLInput, optFns ...func(*wafregional.Options)) (*wafregional.GetWebACLOutput, error)
	ListRateBasedRules(ctx context.Context, params *wafregional.ListRateBasedRulesInput, optFns ...func(*wafregional.Options)) (*wafregional.ListRateBasedRulesOutput, error)
	ListResourcesForWebACL(ctx context.Context, params *wafregional.ListResourcesForWebACLInput, optFns ...func(*wafregional.Options)) (*wafregional.ListResourcesForWebACLOutput, error)
	ListRuleGroups(ctx context.Context, params *wafregional.ListRuleGroupsInput, optFns ...func(*wafregional.Options)) (*wafregional.ListRuleGroupsOutput, error)
	ListRules(ctx context.Context, params *wafregional.ListRulesInput, optFns ...func(*wafregional.Options)) (*wafregional.ListRulesOutput, error)
	ListTagsForResource(ctx context.Context, params *wafregional.ListTagsForResourceInput, optFns ...func(*wafregional.Options)) (*wafregional.ListTagsForResourceOutput, error)
	ListWebACLs(ctx context.Context, params *wafregional.ListWebACLsInput, optFns ...func(*wafregional.Options)) (*wafregional.ListWebACLsOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_wafv2.go . WafV2Client
type WafV2Client interface {
	DescribeManagedRuleGroup(ctx context.Context, params *wafv2.DescribeManagedRuleGroupInput, optFns ...func(*wafv2.Options)) (*wafv2.DescribeManagedRuleGroupOutput, error)
	GetIPSet(ctx context.Context, params *wafv2.GetIPSetInput, optFns ...func(*wafv2.Options)) (*wafv2.GetIPSetOutput, error)
	GetLoggingConfiguration(ctx context.Context, params *wafv2.GetLoggingConfigurationInput, optFns ...func(*wafv2.Options)) (*wafv2.GetLoggingConfigurationOutput, error)
	GetPermissionPolicy(ctx context.Context, params *wafv2.GetPermissionPolicyInput, optFns ...func(*wafv2.Options)) (*wafv2.GetPermissionPolicyOutput, error)
	GetRegexPatternSet(ctx context.Context, params *wafv2.GetRegexPatternSetInput, optFns ...func(*wafv2.Options)) (*wafv2.GetRegexPatternSetOutput, error)
	GetRuleGroup(ctx context.Context, params *wafv2.GetRuleGroupInput, optFns ...func(*wafv2.Options)) (*wafv2.GetRuleGroupOutput, error)
	GetWebACL(ctx context.Context, params *wafv2.GetWebACLInput, optFns ...func(*wafv2.Options)) (*wafv2.GetWebACLOutput, error)
	GetWebACLForResource(ctx context.Context, params *wafv2.GetWebACLForResourceInput, optFns ...func(*wafv2.Options)) (*wafv2.GetWebACLForResourceOutput, error)
	ListAvailableManagedRuleGroups(ctx context.Context, params *wafv2.ListAvailableManagedRuleGroupsInput, optFns ...func(*wafv2.Options)) (*wafv2.ListAvailableManagedRuleGroupsOutput, error)
	ListIPSets(ctx context.Context, params *wafv2.ListIPSetsInput, optFns ...func(*wafv2.Options)) (*wafv2.ListIPSetsOutput, error)
	ListRegexPatternSets(ctx context.Context, params *wafv2.ListRegexPatternSetsInput, optFns ...func(*wafv2.Options)) (*wafv2.ListRegexPatternSetsOutput, error)
	ListResourcesForWebACL(ctx context.Context, params *wafv2.ListResourcesForWebACLInput, optFns ...func(*wafv2.Options)) (*wafv2.ListResourcesForWebACLOutput, error)
	ListRuleGroups(ctx context.Context, params *wafv2.ListRuleGroupsInput, optFns ...func(*wafv2.Options)) (*wafv2.ListRuleGroupsOutput, error)
	ListTagsForResource(ctx context.Context, params *wafv2.ListTagsForResourceInput, optFns ...func(*wafv2.Options)) (*wafv2.ListTagsForResourceOutput, error)
	ListWebACLs(ctx context.Context, params *wafv2.ListWebACLsInput, optFns ...func(*wafv2.Options)) (*wafv2.ListWebACLsOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_workspaces.go . WorkspacesClient
type WorkspacesClient interface {
	DescribeWorkspaces(ctx context.Context, params *workspaces.DescribeWorkspacesInput, optFns ...func(*workspaces.Options)) (*workspaces.DescribeWorkspacesOutput, error)
	DescribeWorkspaceDirectories(ctx context.Context, params *workspaces.DescribeWorkspaceDirectoriesInput, optFns ...func(*workspaces.Options)) (*workspaces.DescribeWorkspaceDirectoriesOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/xray.go . XrayClient
type XrayClient interface {
	GetEncryptionConfig(ctx context.Context, params *xray.GetEncryptionConfigInput, optFns ...func(*xray.Options)) (*xray.GetEncryptionConfigOutput, error)
	GetSamplingRules(ctx context.Context, params *xray.GetSamplingRulesInput, optFns ...func(*xray.Options)) (*xray.GetSamplingRulesOutput, error)
	GetGroups(ctx context.Context, params *xray.GetGroupsInput, optFns ...func(*xray.Options)) (*xray.GetGroupsOutput, error)
	ListTagsForResource(ctx context.Context, params *xray.ListTagsForResourceInput, optFns ...func(*xray.Options)) (*xray.ListTagsForResourceOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/transfer.go . TransferClient
type TransferClient interface {
	DescribeServer(ctx context.Context, params *transfer.DescribeServerInput, optFns ...func(*transfer.Options)) (*transfer.DescribeServerOutput, error)
	ListServers(ctx context.Context, params *transfer.ListServersInput, optFns ...func(*transfer.Options)) (*transfer.ListServersOutput, error)
	ListTagsForResource(ctx context.Context, params *transfer.ListTagsForResourceInput, optFns ...func(*transfer.Options)) (*transfer.ListTagsForResourceOutput, error)
}
