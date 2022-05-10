package client

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer"
	"github.com/aws/aws-sdk-go-v2/service/acm"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/aws/aws-sdk-go-v2/service/backup"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
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
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/efs"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
	elbv1 "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
	elbv2 "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice"
	"github.com/aws/aws-sdk-go-v2/service/emr"
	"github.com/aws/aws-sdk-go-v2/service/fsx"
	"github.com/aws/aws-sdk-go-v2/service/guardduty"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/mq"
	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/aws/aws-sdk-go-v2/service/qldb"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/route53domains"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3control"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/shield"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
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

//go:generate mockgen -package=mocks -destination=./mocks/mock_application_autoscaling.go . ApplicationAutoscalingClient
type ApplicationAutoscalingClient interface {
	DescribeScalingPolicies(ctx context.Context, params *applicationautoscaling.DescribeScalingPoliciesInput, optFns ...func(*applicationautoscaling.Options)) (*applicationautoscaling.DescribeScalingPoliciesOutput, error)
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

//go:generate mockgen -package=mocks -destination=./mocks/mock_apigateway.go . ApigatewayClient
type ApigatewayClient interface {
	GetRestApis(ctx context.Context, params *apigateway.GetRestApisInput, optFns ...func(*apigateway.Options)) (*apigateway.GetRestApisOutput, error)
	GetAuthorizers(ctx context.Context, params *apigateway.GetAuthorizersInput, optFns ...func(*apigateway.Options)) (*apigateway.GetAuthorizersOutput, error)
	GetDeployments(ctx context.Context, params *apigateway.GetDeploymentsInput, optFns ...func(*apigateway.Options)) (*apigateway.GetDeploymentsOutput, error)
	GetDocumentationParts(ctx context.Context, params *apigateway.GetDocumentationPartsInput, optFns ...func(*apigateway.Options)) (*apigateway.GetDocumentationPartsOutput, error)
	GetDocumentationVersions(ctx context.Context, params *apigateway.GetDocumentationVersionsInput, optFns ...func(*apigateway.Options)) (*apigateway.GetDocumentationVersionsOutput, error)
	GetGatewayResponses(ctx context.Context, params *apigateway.GetGatewayResponsesInput, optFns ...func(*apigateway.Options)) (*apigateway.GetGatewayResponsesOutput, error)
	GetModels(ctx context.Context, params *apigateway.GetModelsInput, optFns ...func(*apigateway.Options)) (*apigateway.GetModelsOutput, error)
	GetModelTemplate(ctx context.Context, params *apigateway.GetModelTemplateInput, optFns ...func(*apigateway.Options)) (*apigateway.GetModelTemplateOutput, error)
	GetRequestValidators(ctx context.Context, params *apigateway.GetRequestValidatorsInput, optFns ...func(*apigateway.Options)) (*apigateway.GetRequestValidatorsOutput, error)
	GetResources(ctx context.Context, params *apigateway.GetResourcesInput, optFns ...func(*apigateway.Options)) (*apigateway.GetResourcesOutput, error)
	GetStages(ctx context.Context, params *apigateway.GetStagesInput, optFns ...func(*apigateway.Options)) (*apigateway.GetStagesOutput, error)
	GetUsagePlans(ctx context.Context, params *apigateway.GetUsagePlansInput, optFns ...func(*apigateway.Options)) (*apigateway.GetUsagePlansOutput, error)
	GetUsagePlanKeys(ctx context.Context, params *apigateway.GetUsagePlanKeysInput, optFns ...func(*apigateway.Options)) (*apigateway.GetUsagePlanKeysOutput, error)
	GetDomainNames(ctx context.Context, params *apigateway.GetDomainNamesInput, optFns ...func(*apigateway.Options)) (*apigateway.GetDomainNamesOutput, error)
	GetBasePathMappings(ctx context.Context, params *apigateway.GetBasePathMappingsInput, optFns ...func(*apigateway.Options)) (*apigateway.GetBasePathMappingsOutput, error)
	GetClientCertificates(ctx context.Context, params *apigateway.GetClientCertificatesInput, optFns ...func(*apigateway.Options)) (*apigateway.GetClientCertificatesOutput, error)
	GetApiKeys(ctx context.Context, params *apigateway.GetApiKeysInput, optFns ...func(*apigateway.Options)) (*apigateway.GetApiKeysOutput, error)
	GetVpcLinks(ctx context.Context, params *apigateway.GetVpcLinksInput, optFns ...func(*apigateway.Options)) (*apigateway.GetVpcLinksOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_apigatewayv2.go . Apigatewayv2Client
type Apigatewayv2Client interface {
	GetApis(ctx context.Context, params *apigatewayv2.GetApisInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.GetApisOutput, error)
	GetApiMappings(ctx context.Context, params *apigatewayv2.GetApiMappingsInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.GetApiMappingsOutput, error)
	GetAuthorizers(ctx context.Context, params *apigatewayv2.GetAuthorizersInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.GetAuthorizersOutput, error)
	GetDeployments(ctx context.Context, params *apigatewayv2.GetDeploymentsInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.GetDeploymentsOutput, error)
	GetIntegrations(ctx context.Context, params *apigatewayv2.GetIntegrationsInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.GetIntegrationsOutput, error)
	GetIntegrationResponses(ctx context.Context, params *apigatewayv2.GetIntegrationResponsesInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.GetIntegrationResponsesOutput, error)
	GetModels(ctx context.Context, params *apigatewayv2.GetModelsInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.GetModelsOutput, error)
	GetModelTemplate(ctx context.Context, params *apigatewayv2.GetModelTemplateInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.GetModelTemplateOutput, error)
	GetRoutes(ctx context.Context, params *apigatewayv2.GetRoutesInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.GetRoutesOutput, error)
	GetRouteResponses(ctx context.Context, params *apigatewayv2.GetRouteResponsesInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.GetRouteResponsesOutput, error)
	GetStages(ctx context.Context, params *apigatewayv2.GetStagesInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.GetStagesOutput, error)
	GetVpcLinks(ctx context.Context, params *apigatewayv2.GetVpcLinksInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.GetVpcLinksOutput, error)
	GetDomainNames(ctx context.Context, params *apigatewayv2.GetDomainNamesInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.GetDomainNamesOutput, error)
	GetTags(ctx context.Context, params *apigatewayv2.GetTagsInput, optFns ...func(*apigatewayv2.Options)) (*apigatewayv2.GetTagsOutput, error)
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
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_cloudwatchlogs.go . CloudwatchLogsClient
type CloudwatchLogsClient interface {
	DescribeMetricFilters(ctx context.Context, params *cloudwatchlogs.DescribeMetricFiltersInput, optFns ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.DescribeMetricFiltersOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_cloudformation.go . CloudFormationClient
type CloudFormationClient interface {
	cloudformation.DescribeStacksAPIClient
	cloudformation.ListStackResourcesAPIClient
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

//go:generate mockgen -package=mocks -destination=./mocks/mock_dax.go . DAXClient
type DAXClient interface {
	DescribeClusters(ctx context.Context, params *dax.DescribeClustersInput, optFns ...func(*dax.Options)) (*dax.DescribeClustersOutput, error)
	ListTags(ctx context.Context, params *dax.ListTagsInput, optFns ...func(*dax.Options)) (*dax.ListTagsOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_directconnect.go . DirectconnectClient
type DirectconnectClient interface {
	DescribeConnections(ctx context.Context, params *directconnect.DescribeConnectionsInput, optFns ...func(*directconnect.Options)) (*directconnect.DescribeConnectionsOutput, error)
	DescribeDirectConnectGateways(ctx context.Context, params *directconnect.DescribeDirectConnectGatewaysInput, optFns ...func(*directconnect.Options)) (*directconnect.DescribeDirectConnectGatewaysOutput, error)
	DescribeDirectConnectGatewayAssociations(ctx context.Context, params *directconnect.DescribeDirectConnectGatewayAssociationsInput, optFns ...func(*directconnect.Options)) (*directconnect.DescribeDirectConnectGatewayAssociationsOutput, error)
	DescribeDirectConnectGatewayAttachments(ctx context.Context, params *directconnect.DescribeDirectConnectGatewayAttachmentsInput, optFns ...func(*directconnect.Options)) (*directconnect.DescribeDirectConnectGatewayAttachmentsOutput, error)
	DescribeLags(ctx context.Context, params *directconnect.DescribeLagsInput, optFns ...func(*directconnect.Options)) (*directconnect.DescribeLagsOutput, error)
	DescribeVirtualGateways(ctx context.Context, params *directconnect.DescribeVirtualGatewaysInput, optFns ...func(*directconnect.Options)) (*directconnect.DescribeVirtualGatewaysOutput, error)
	DescribeVirtualInterfaces(ctx context.Context, params *directconnect.DescribeVirtualInterfacesInput, optFns ...func(*directconnect.Options)) (*directconnect.DescribeVirtualInterfacesOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_dynamodb.go . DynamoDBClient
type DynamoDBClient interface {
	ListTables(ctx context.Context, params *dynamodb.ListTablesInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ListTablesOutput, error)
	DescribeTable(ctx context.Context, params *dynamodb.DescribeTableInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DescribeTableOutput, error)
	DescribeTableReplicaAutoScaling(ctx context.Context, params *dynamodb.DescribeTableReplicaAutoScalingInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DescribeTableReplicaAutoScalingOutput, error)
	DescribeContinuousBackups(ctx context.Context, params *dynamodb.DescribeContinuousBackupsInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DescribeContinuousBackupsOutput, error)
	ListTagsOfResource(ctx context.Context, params *dynamodb.ListTagsOfResourceInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ListTagsOfResourceOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_databasemigrationservice.go . DatabasemigrationserviceClient
type DatabasemigrationserviceClient interface {
	DescribeReplicationInstances(ctx context.Context, params *databasemigrationservice.DescribeReplicationInstancesInput, optFns ...func(*databasemigrationservice.Options)) (*databasemigrationservice.DescribeReplicationInstancesOutput, error)
	ListTagsForResource(ctx context.Context, params *databasemigrationservice.ListTagsForResourceInput, optFns ...func(*databasemigrationservice.Options)) (*databasemigrationservice.ListTagsForResourceOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_ec2.go . Ec2Client
type Ec2Client interface {
	DescribeRegions(ctx context.Context, params *ec2.DescribeRegionsInput, optFns ...func(*ec2.Options)) (*ec2.DescribeRegionsOutput, error)
	DescribeByoipCidrs(ctx context.Context, params *ec2.DescribeByoipCidrsInput, optFns ...func(*ec2.Options)) (*ec2.DescribeByoipCidrsOutput, error)
	DescribeCustomerGateways(ctx context.Context, params *ec2.DescribeCustomerGatewaysInput, optFns ...func(*ec2.Options)) (*ec2.DescribeCustomerGatewaysOutput, error)
	DescribeEgressOnlyInternetGateways(ctx context.Context, params *ec2.DescribeEgressOnlyInternetGatewaysInput, optFns ...func(*ec2.Options)) (*ec2.DescribeEgressOnlyInternetGatewaysOutput, error)
	DescribeFlowLogs(ctx context.Context, params *ec2.DescribeFlowLogsInput, optFns ...func(*ec2.Options)) (*ec2.DescribeFlowLogsOutput, error)
	DescribeHosts(ctx context.Context, params *ec2.DescribeHostsInput, optFns ...func(*ec2.Options)) (*ec2.DescribeHostsOutput, error)
	DescribeImageAttribute(ctx context.Context, params *ec2.DescribeImageAttributeInput, optFns ...func(*ec2.Options)) (*ec2.DescribeImageAttributeOutput, error)
	DescribeImages(ctx context.Context, params *ec2.DescribeImagesInput, optFns ...func(*ec2.Options)) (*ec2.DescribeImagesOutput, error)
	DescribeInstances(ctx context.Context, params *ec2.DescribeInstancesInput, optFns ...func(*ec2.Options)) (*ec2.DescribeInstancesOutput, error)
	DescribeInstanceStatus(ctx context.Context, params *ec2.DescribeInstanceStatusInput, optFns ...func(*ec2.Options)) (*ec2.DescribeInstanceStatusOutput, error)
	DescribeInternetGateways(ctx context.Context, params *ec2.DescribeInternetGatewaysInput, optFns ...func(*ec2.Options)) (*ec2.DescribeInternetGatewaysOutput, error)
	DescribeNatGateways(ctx context.Context, params *ec2.DescribeNatGatewaysInput, optFns ...func(*ec2.Options)) (*ec2.DescribeNatGatewaysOutput, error)
	DescribeNetworkInterfaces(ctx context.Context, params *ec2.DescribeNetworkInterfacesInput, optFns ...func(*ec2.Options)) (*ec2.DescribeNetworkInterfacesOutput, error)
	DescribeNetworkAcls(ctx context.Context, params *ec2.DescribeNetworkAclsInput, optFns ...func(*ec2.Options)) (*ec2.DescribeNetworkAclsOutput, error)
	DescribeRouteTables(ctx context.Context, params *ec2.DescribeRouteTablesInput, optFns ...func(*ec2.Options)) (*ec2.DescribeRouteTablesOutput, error)
	DescribeSecurityGroups(ctx context.Context, params *ec2.DescribeSecurityGroupsInput, optFns ...func(*ec2.Options)) (*ec2.DescribeSecurityGroupsOutput, error)
	DescribeSubnets(ctx context.Context, params *ec2.DescribeSubnetsInput, optFns ...func(*ec2.Options)) (*ec2.DescribeSubnetsOutput, error)
	DescribeTransitGatewayAttachments(ctx context.Context, params *ec2.DescribeTransitGatewayAttachmentsInput, optFns ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayAttachmentsOutput, error)
	DescribeTransitGatewayMulticastDomains(ctx context.Context, params *ec2.DescribeTransitGatewayMulticastDomainsInput, optFns ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayMulticastDomainsOutput, error)
	DescribeTransitGatewayRouteTables(ctx context.Context, params *ec2.DescribeTransitGatewayRouteTablesInput, optFns ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayRouteTablesOutput, error)
	DescribeTransitGatewayPeeringAttachments(ctx context.Context, params *ec2.DescribeTransitGatewayPeeringAttachmentsInput, optFns ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayPeeringAttachmentsOutput, error)
	DescribeTransitGateways(ctx context.Context, params *ec2.DescribeTransitGatewaysInput, optFns ...func(*ec2.Options)) (*ec2.DescribeTransitGatewaysOutput, error)
	DescribeTransitGatewayVpcAttachments(ctx context.Context, params *ec2.DescribeTransitGatewayVpcAttachmentsInput, optFns ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayVpcAttachmentsOutput, error)
	DescribeVpcPeeringConnections(ctx context.Context, params *ec2.DescribeVpcPeeringConnectionsInput, optFns ...func(*ec2.Options)) (*ec2.DescribeVpcPeeringConnectionsOutput, error)
	DescribeVolumes(ctx context.Context, params *ec2.DescribeVolumesInput, optFns ...func(*ec2.Options)) (*ec2.DescribeVolumesOutput, error)
	DescribeSnapshots(ctx context.Context, params *ec2.DescribeSnapshotsInput, optFns ...func(*ec2.Options)) (*ec2.DescribeSnapshotsOutput, error)
	DescribeSnapshotAttribute(ctx context.Context, params *ec2.DescribeSnapshotAttributeInput, optFns ...func(*ec2.Options)) (*ec2.DescribeSnapshotAttributeOutput, error)
	DescribeVpcs(ctx context.Context, params *ec2.DescribeVpcsInput, optFns ...func(*ec2.Options)) (*ec2.DescribeVpcsOutput, error)
	DescribeVpcEndpoints(ctx context.Context, params *ec2.DescribeVpcEndpointsInput, optFns ...func(*ec2.Options)) (*ec2.DescribeVpcEndpointsOutput, error)
	DescribeVpnGateways(ctx context.Context, params *ec2.DescribeVpnGatewaysInput, optFns ...func(*ec2.Options)) (*ec2.DescribeVpnGatewaysOutput, error)
	DescribeAddresses(ctx context.Context, params *ec2.DescribeAddressesInput, optFns ...func(*ec2.Options)) (*ec2.DescribeAddressesOutput, error)
	GetEbsEncryptionByDefault(ctx context.Context, params *ec2.GetEbsEncryptionByDefaultInput, optFns ...func(*ec2.Options)) (*ec2.GetEbsEncryptionByDefaultOutput, error)
	GetEbsDefaultKmsKeyId(ctx context.Context, params *ec2.GetEbsDefaultKmsKeyIdInput, optFns ...func(*ec2.Options)) (*ec2.GetEbsDefaultKmsKeyIdOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_ecr.go . EcrClient
type EcrClient interface {
	DescribeRepositories(ctx context.Context, params *ecr.DescribeRepositoriesInput, optFns ...func(*ecr.Options)) (*ecr.DescribeRepositoriesOutput, error)
	DescribeImages(ctx context.Context, params *ecr.DescribeImagesInput, optFns ...func(*ecr.Options)) (*ecr.DescribeImagesOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_efs.go . EfsClient
type EfsClient interface {
	DescribeFileSystems(ctx context.Context, params *efs.DescribeFileSystemsInput, optFns ...func(*efs.Options)) (*efs.DescribeFileSystemsOutput, error)
	DescribeBackupPolicy(ctx context.Context, params *efs.DescribeBackupPolicyInput, optFns ...func(*efs.Options)) (*efs.DescribeBackupPolicyOutput, error)
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

//go:generate mockgen -package=mocks -destination=./mocks/mock_elbv2.go . ElbV2Client
type ElbV2Client interface {
	DescribeListenerCertificates(ctx context.Context, params *elbv2.DescribeListenerCertificatesInput, optFns ...func(*elbv2.Options)) (*elbv2.DescribeListenerCertificatesOutput, error)
	DescribeListeners(ctx context.Context, params *elbv2.DescribeListenersInput, optFns ...func(*elbv2.Options)) (*elbv2.DescribeListenersOutput, error)
	DescribeLoadBalancers(ctx context.Context, params *elbv2.DescribeLoadBalancersInput, optFns ...func(*elbv2.Options)) (*elbv2.DescribeLoadBalancersOutput, error)
	DescribeLoadBalancerAttributes(ctx context.Context, params *elbv2.DescribeLoadBalancerAttributesInput, optFns ...func(*elbv2.Options)) (*elbv2.DescribeLoadBalancerAttributesOutput, error)
	DescribeTargetGroups(ctx context.Context, params *elbv2.DescribeTargetGroupsInput, optFns ...func(*elbv2.Options)) (*elbv2.DescribeTargetGroupsOutput, error)
	DescribeTags(ctx context.Context, params *elbv2.DescribeTagsInput, optFns ...func(*elbv2.Options)) (*elbv2.DescribeTagsOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_elbv1.go . ElbV1Client
type ElbV1Client interface {
	DescribeLoadBalancers(ctx context.Context, params *elbv1.DescribeLoadBalancersInput, optFns ...func(*elbv1.Options)) (*elbv1.DescribeLoadBalancersOutput, error)
	DescribeLoadBalancerPolicies(ctx context.Context, params *elbv1.DescribeLoadBalancerPoliciesInput, optFns ...func(*elbv1.Options)) (*elbv1.DescribeLoadBalancerPoliciesOutput, error)
	DescribeTags(ctx context.Context, params *elbv1.DescribeTagsInput, optFns ...func(*elbv1.Options)) (*elbv1.DescribeTagsOutput, error)
	DescribeLoadBalancerAttributes(ctx context.Context, params *elbv1.DescribeLoadBalancerAttributesInput, optFns ...func(*elbv1.Options)) (*elbv1.DescribeLoadBalancerAttributesOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_emr.go . EmrClient
type EmrClient interface {
	DescribeCluster(ctx context.Context, params *emr.DescribeClusterInput, optFns ...func(*emr.Options)) (*emr.DescribeClusterOutput, error)
	GetBlockPublicAccessConfiguration(ctx context.Context, params *emr.GetBlockPublicAccessConfigurationInput, optFns ...func(*emr.Options)) (*emr.GetBlockPublicAccessConfigurationOutput, error)
	ListClusters(ctx context.Context, params *emr.ListClustersInput, optFns ...func(*emr.Options)) (*emr.ListClustersOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_fsx.go . FsxClient
type FsxClient interface {
	DescribeBackups(ctx context.Context, params *fsx.DescribeBackupsInput, optFns ...func(*fsx.Options)) (*fsx.DescribeBackupsOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_iam.go . IamClient
type IamClient interface {
	GetAccountAuthorizationDetails(context.Context, *iam.GetAccountAuthorizationDetailsInput, ...func(*iam.Options)) (*iam.GetAccountAuthorizationDetailsOutput, error)
	ListGroups(ctx context.Context, params *iam.ListGroupsInput, optFns ...func(*iam.Options)) (*iam.ListGroupsOutput, error)
	ListAttachedGroupPolicies(ctx context.Context, params *iam.ListAttachedGroupPoliciesInput, optFns ...func(*iam.Options)) (*iam.ListAttachedGroupPoliciesOutput, error)
	GetAccountPasswordPolicy(ctx context.Context, params *iam.GetAccountPasswordPolicyInput, optFns ...func(*iam.Options)) (*iam.GetAccountPasswordPolicyOutput, error)
	GetCredentialReport(ctx context.Context, params *iam.GetCredentialReportInput, optFns ...func(*iam.Options)) (*iam.GetCredentialReportOutput, error)
	GenerateCredentialReport(ctx context.Context, params *iam.GenerateCredentialReportInput, optFns ...func(*iam.Options)) (*iam.GenerateCredentialReportOutput, error)
	ListUsers(ctx context.Context, params *iam.ListUsersInput, optFns ...func(*iam.Options)) (*iam.ListUsersOutput, error)
	GetAccessKeyLastUsed(ctx context.Context, params *iam.GetAccessKeyLastUsedInput, optFns ...func(*iam.Options)) (*iam.GetAccessKeyLastUsedOutput, error)
	ListAttachedUserPolicies(ctx context.Context, params *iam.ListAttachedUserPoliciesInput, optFns ...func(*iam.Options)) (*iam.ListAttachedUserPoliciesOutput, error)
	GetUser(ctx context.Context, params *iam.GetUserInput, optFns ...func(*iam.Options)) (*iam.GetUserOutput, error)
	ListAccessKeys(ctx context.Context, params *iam.ListAccessKeysInput, optFns ...func(*iam.Options)) (*iam.ListAccessKeysOutput, error)
	ListRoles(ctx context.Context, params *iam.ListRolesInput, optFns ...func(*iam.Options)) (*iam.ListRolesOutput, error)
	ListAttachedRolePolicies(ctx context.Context, params *iam.ListAttachedRolePoliciesInput, optFns ...func(*iam.Options)) (*iam.ListAttachedRolePoliciesOutput, error)
	ListVirtualMFADevices(ctx context.Context, params *iam.ListVirtualMFADevicesInput, optFns ...func(*iam.Options)) (*iam.ListVirtualMFADevicesOutput, error)
	ListGroupsForUser(ctx context.Context, params *iam.ListGroupsForUserInput, optFns ...func(*iam.Options)) (*iam.ListGroupsForUserOutput, error)
	ListUserTags(ctx context.Context, params *iam.ListUserTagsInput, optFns ...func(*iam.Options)) (*iam.ListUserTagsOutput, error)
	ListRolePolicies(ctx context.Context, params *iam.ListRolePoliciesInput, optFns ...func(*iam.Options)) (*iam.ListRolePoliciesOutput, error)
	ListUserPolicies(ctx context.Context, params *iam.ListUserPoliciesInput, optFns ...func(*iam.Options)) (*iam.ListUserPoliciesOutput, error)
	ListGroupPolicies(ctx context.Context, params *iam.ListGroupPoliciesInput, optFns ...func(*iam.Options)) (*iam.ListGroupPoliciesOutput, error)
	GetRolePolicy(ctx context.Context, params *iam.GetRolePolicyInput, optFns ...func(*iam.Options)) (*iam.GetRolePolicyOutput, error)
	GetGroupPolicy(ctx context.Context, params *iam.GetGroupPolicyInput, optFns ...func(*iam.Options)) (*iam.GetGroupPolicyOutput, error)
	GetUserPolicy(ctx context.Context, params *iam.GetUserPolicyInput, optFns ...func(*iam.Options)) (*iam.GetUserPolicyOutput, error)
	ListOpenIDConnectProviders(ctx context.Context, params *iam.ListOpenIDConnectProvidersInput, optFns ...func(*iam.Options)) (*iam.ListOpenIDConnectProvidersOutput, error)
	GetOpenIDConnectProvider(ctx context.Context, params *iam.GetOpenIDConnectProviderInput, optFns ...func(*iam.Options)) (*iam.GetOpenIDConnectProviderOutput, error)
	ListSAMLProviders(ctx context.Context, params *iam.ListSAMLProvidersInput, optFns ...func(*iam.Options)) (*iam.ListSAMLProvidersOutput, error)
	GetSAMLProvider(ctx context.Context, params *iam.GetSAMLProviderInput, optFns ...func(*iam.Options)) (*iam.GetSAMLProviderOutput, error)
	ListRoleTags(ctx context.Context, params *iam.ListRoleTagsInput, optFns ...func(*iam.Options)) (*iam.ListRoleTagsOutput, error)

	iam.ListServerCertificatesAPIClient
	iam.ListAccountAliasesAPIClient
	GetAccountSummary(ctx context.Context, params *iam.GetAccountSummaryInput, optFns ...func(*iam.Options)) (*iam.GetAccountSummaryOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_kms.go . KmsClient
type KmsClient interface {
	ListKeys(ctx context.Context, params *kms.ListKeysInput, optFns ...func(*kms.Options)) (*kms.ListKeysOutput, error)
	DescribeKey(ctx context.Context, params *kms.DescribeKeyInput, optFns ...func(*kms.Options)) (*kms.DescribeKeyOutput, error)
	GetKeyRotationStatus(ctx context.Context, params *kms.GetKeyRotationStatusInput, optFns ...func(*kms.Options)) (*kms.GetKeyRotationStatusOutput, error)
	ListResourceTags(ctx context.Context, params *kms.ListResourceTagsInput, optFns ...func(*kms.Options)) (*kms.ListResourceTagsOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_mq.go . MQClient
type MQClient interface {
	DescribeBroker(ctx context.Context, params *mq.DescribeBrokerInput, optFns ...func(*mq.Options)) (*mq.DescribeBrokerOutput, error)
	DescribeConfiguration(ctx context.Context, params *mq.DescribeConfigurationInput, optFns ...func(*mq.Options)) (*mq.DescribeConfigurationOutput, error)
	ListConfigurationRevisions(ctx context.Context, params *mq.ListConfigurationRevisionsInput, optFns ...func(*mq.Options)) (*mq.ListConfigurationRevisionsOutput, error)
	DescribeConfigurationRevision(ctx context.Context, params *mq.DescribeConfigurationRevisionInput, optFns ...func(*mq.Options)) (*mq.DescribeConfigurationRevisionOutput, error)
	DescribeUser(ctx context.Context, params *mq.DescribeUserInput, optFns ...func(*mq.Options)) (*mq.DescribeUserOutput, error)
	ListBrokers(ctx context.Context, params *mq.ListBrokersInput, optFns ...func(*mq.Options)) (*mq.ListBrokersOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_organizations.go . OrganizationsClient
type OrganizationsClient interface {
	ListAccountsForParent(ctx context.Context, params *organizations.ListAccountsForParentInput, optFns ...func(*organizations.Options)) (*organizations.ListAccountsForParentOutput, error)
	ListAccounts(ctx context.Context, params *organizations.ListAccountsInput, optFns ...func(*organizations.Options)) (*organizations.ListAccountsOutput, error)
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

//go:generate mockgen -package=mocks -destination=./mocks/mock_s3Control.go . S3ControlClient
type S3ControlClient interface {
	GetPublicAccessBlock(ctx context.Context, params *s3control.GetPublicAccessBlockInput, optFns ...func(*s3control.Options)) (*s3control.GetPublicAccessBlockOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_s3.go . S3Client
type S3Client interface {
	ListBuckets(ctx context.Context, params *s3.ListBucketsInput, optFns ...func(*s3.Options)) (*s3.ListBucketsOutput, error)
	GetBucketLocation(ctx context.Context, params *s3.GetBucketLocationInput, optFns ...func(*s3.Options)) (*s3.GetBucketLocationOutput, error)
	GetBucketLogging(ctx context.Context, params *s3.GetBucketLoggingInput, optFns ...func(*s3.Options)) (*s3.GetBucketLoggingOutput, error)
	GetBucketAcl(ctx context.Context, params *s3.GetBucketAclInput, optFns ...func(*s3.Options)) (*s3.GetBucketAclOutput, error)
	GetBucketCors(ctx context.Context, params *s3.GetBucketCorsInput, optFns ...func(*s3.Options)) (*s3.GetBucketCorsOutput, error)
	GetBucketPolicy(ctx context.Context, params *s3.GetBucketPolicyInput, optFns ...func(*s3.Options)) (*s3.GetBucketPolicyOutput, error)
	GetBucketVersioning(ctx context.Context, params *s3.GetBucketVersioningInput, optFns ...func(*s3.Options)) (*s3.GetBucketVersioningOutput, error)
	GetBucketEncryption(ctx context.Context, params *s3.GetBucketEncryptionInput, optFns ...func(*s3.Options)) (*s3.GetBucketEncryptionOutput, error)
	GetPublicAccessBlock(ctx context.Context, params *s3.GetPublicAccessBlockInput, optFns ...func(*s3.Options)) (*s3.GetPublicAccessBlockOutput, error)
	GetBucketReplication(ctx context.Context, params *s3.GetBucketReplicationInput, optFns ...func(*s3.Options)) (*s3.GetBucketReplicationOutput, error)
	GetBucketLifecycleConfiguration(ctx context.Context, params *s3.GetBucketLifecycleConfigurationInput, optFns ...func(*s3.Options)) (*s3.GetBucketLifecycleConfigurationOutput, error)
	GetBucketTagging(ctx context.Context, params *s3.GetBucketTaggingInput, optFns ...func(*s3.Options)) (*s3.GetBucketTaggingOutput, error)
	GetBucketOwnershipControls(ctx context.Context, params *s3.GetBucketOwnershipControlsInput, optFns ...func(*s3.Options)) (*s3.GetBucketOwnershipControlsOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_sns.go . SnsClient
type SnsClient interface {
	ListTopics(ctx context.Context, params *sns.ListTopicsInput, optFns ...func(*sns.Options)) (*sns.ListTopicsOutput, error)
	ListSubscriptions(ctx context.Context, params *sns.ListSubscriptionsInput, optFns ...func(*sns.Options)) (*sns.ListSubscriptionsOutput, error)
	GetTopicAttributes(ctx context.Context, params *sns.GetTopicAttributesInput, optFns ...func(*sns.Options)) (*sns.GetTopicAttributesOutput, error)
	ListTagsForResource(ctx context.Context, params *sns.ListTagsForResourceInput, optFns ...func(*sns.Options)) (*sns.ListTagsForResourceOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_ecs.go . EcsClient
type EcsClient interface {
	DescribeClusters(ctx context.Context, params *ecs.DescribeClustersInput, optFns ...func(*ecs.Options)) (*ecs.DescribeClustersOutput, error)
	ListClusters(ctx context.Context, params *ecs.ListClustersInput, optFns ...func(*ecs.Options)) (*ecs.ListClustersOutput, error)
	ListTagsForResource(ctx context.Context, params *ecs.ListTagsForResourceInput, optFns ...func(*ecs.Options)) (*ecs.ListTagsForResourceOutput, error)
	DescribeServices(ctx context.Context, params *ecs.DescribeServicesInput, optFns ...func(*ecs.Options)) (*ecs.DescribeServicesOutput, error)
	DescribeContainerInstances(ctx context.Context, params *ecs.DescribeContainerInstancesInput, optFns ...func(*ecs.Options)) (*ecs.DescribeContainerInstancesOutput, error)
	ListServices(ctx context.Context, params *ecs.ListServicesInput, optFns ...func(*ecs.Options)) (*ecs.ListServicesOutput, error)
	ListContainerInstances(ctx context.Context, params *ecs.ListContainerInstancesInput, optFns ...func(*ecs.Options)) (*ecs.ListContainerInstancesOutput, error)
	ListTaskDefinitions(ctx context.Context, params *ecs.ListTaskDefinitionsInput, optFns ...func(*ecs.Options)) (*ecs.ListTaskDefinitionsOutput, error)
	DescribeTaskDefinition(ctx context.Context, params *ecs.DescribeTaskDefinitionInput, optFns ...func(*ecs.Options)) (*ecs.DescribeTaskDefinitionOutput, error)
	ListTasks(ctx context.Context, params *ecs.ListTasksInput, optFns ...func(*ecs.Options)) (*ecs.ListTasksOutput, error)
	DescribeTasks(ctx context.Context, params *ecs.DescribeTasksInput, optFns ...func(*ecs.Options)) (*ecs.DescribeTasksOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_elasticsearch.go . ElasticSearch
type ElasticSearch interface {
	ListDomainNames(ctx context.Context, params *elasticsearchservice.ListDomainNamesInput, optFns ...func(*elasticsearchservice.Options)) (*elasticsearchservice.ListDomainNamesOutput, error)
	DescribeElasticsearchDomain(ctx context.Context, params *elasticsearchservice.DescribeElasticsearchDomainInput, optFns ...func(*elasticsearchservice.Options)) (*elasticsearchservice.DescribeElasticsearchDomainOutput, error)
	ListTags(ctx context.Context, params *elasticsearchservice.ListTagsInput, optFns ...func(*elasticsearchservice.Options)) (*elasticsearchservice.ListTagsOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_eks.go . EksClient
type EksClient interface {
	ListClusters(ctx context.Context, params *eks.ListClustersInput, optFns ...func(*eks.Options)) (*eks.ListClustersOutput, error)
	DescribeCluster(ctx context.Context, params *eks.DescribeClusterInput, optFns ...func(*eks.Options)) (*eks.DescribeClusterOutput, error)
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

//go:generate mockgen -package=mocks -destination=./mocks/mock_route53.go . Route53Client
type Route53Client interface {
	ListHostedZones(ctx context.Context, params *route53.ListHostedZonesInput, optFns ...func(*route53.Options)) (*route53.ListHostedZonesOutput, error)
	ListTagsForResource(ctx context.Context, params *route53.ListTagsForResourceInput, optFns ...func(*route53.Options)) (*route53.ListTagsForResourceOutput, error)
	ListTagsForResources(ctx context.Context, params *route53.ListTagsForResourcesInput, optFns ...func(*route53.Options)) (*route53.ListTagsForResourcesOutput, error)
	ListQueryLoggingConfigs(ctx context.Context, params *route53.ListQueryLoggingConfigsInput, optFns ...func(*route53.Options)) (*route53.ListQueryLoggingConfigsOutput, error)
	ListResourceRecordSets(ctx context.Context, params *route53.ListResourceRecordSetsInput, optFns ...func(*route53.Options)) (*route53.ListResourceRecordSetsOutput, error)
	ListTrafficPolicies(ctx context.Context, params *route53.ListTrafficPoliciesInput, optFns ...func(*route53.Options)) (*route53.ListTrafficPoliciesOutput, error)
	ListTrafficPolicyInstancesByHostedZone(ctx context.Context, params *route53.ListTrafficPolicyInstancesByHostedZoneInput, optFns ...func(*route53.Options)) (*route53.ListTrafficPolicyInstancesByHostedZoneOutput, error)
	GetTrafficPolicy(ctx context.Context, params *route53.GetTrafficPolicyInput, optFns ...func(*route53.Options)) (*route53.GetTrafficPolicyOutput, error)
	ListHealthChecks(ctx context.Context, params *route53.ListHealthChecksInput, optFns ...func(*route53.Options)) (*route53.ListHealthChecksOutput, error)
	ListVPCAssociationAuthorizations(ctx context.Context, params *route53.ListVPCAssociationAuthorizationsInput, optFns ...func(*route53.Options)) (*route53.ListVPCAssociationAuthorizationsOutput, error)
	ListTrafficPolicyVersions(ctx context.Context, params *route53.ListTrafficPolicyVersionsInput, optFns ...func(*route53.Options)) (*route53.ListTrafficPolicyVersionsOutput, error)
	GetHostedZone(ctx context.Context, params *route53.GetHostedZoneInput, optFns ...func(*route53.Options)) (*route53.GetHostedZoneOutput, error)
	ListReusableDelegationSets(ctx context.Context, params *route53.ListReusableDelegationSetsInput, optFns ...func(*route53.Options)) (*route53.ListReusableDelegationSetsOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_route53_domains.go . Route53DomainsClient
type Route53DomainsClient interface {
	GetDomainDetail(ctx context.Context, params *route53domains.GetDomainDetailInput, optFns ...func(*route53domains.Options)) (*route53domains.GetDomainDetailOutput, error)
	ListDomains(ctx context.Context, params *route53domains.ListDomainsInput, optFns ...func(*route53domains.Options)) (*route53domains.ListDomainsOutput, error)
	ListTagsForDomain(ctx context.Context, params *route53domains.ListTagsForDomainInput, optFns ...func(*route53domains.Options)) (*route53domains.ListTagsForDomainOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_s3manager.go . S3ManagerClient
type S3ManagerClient interface {
	GetBucketRegion(ctx context.Context, bucket string, optFns ...func(*s3.Options)) (string, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_lambda.go . LambdaClient
type LambdaClient interface {
	ListFunctions(ctx context.Context, params *lambda.ListFunctionsInput, optFns ...func(*lambda.Options)) (*lambda.ListFunctionsOutput, error)
	GetFunction(ctx context.Context, params *lambda.GetFunctionInput, optFns ...func(*lambda.Options)) (*lambda.GetFunctionOutput, error)
	ListAliases(ctx context.Context, params *lambda.ListAliasesInput, optFns ...func(*lambda.Options)) (*lambda.ListAliasesOutput, error)
	ListFunctionEventInvokeConfigs(ctx context.Context, params *lambda.ListFunctionEventInvokeConfigsInput, optFns ...func(*lambda.Options)) (*lambda.ListFunctionEventInvokeConfigsOutput, error)
	ListProvisionedConcurrencyConfigs(ctx context.Context, params *lambda.ListProvisionedConcurrencyConfigsInput, optFns ...func(*lambda.Options)) (*lambda.ListProvisionedConcurrencyConfigsOutput, error)
	ListVersionsByFunction(ctx context.Context, params *lambda.ListVersionsByFunctionInput, optFns ...func(*lambda.Options)) (*lambda.ListVersionsByFunctionOutput, error)
	ListLayers(ctx context.Context, params *lambda.ListLayersInput, optFns ...func(*lambda.Options)) (*lambda.ListLayersOutput, error)
	ListLayerVersions(ctx context.Context, params *lambda.ListLayerVersionsInput, optFns ...func(*lambda.Options)) (*lambda.ListLayerVersionsOutput, error)
	ListEventSourceMappings(ctx context.Context, params *lambda.ListEventSourceMappingsInput, optFns ...func(*lambda.Options)) (*lambda.ListEventSourceMappingsOutput, error)
	GetPolicy(ctx context.Context, params *lambda.GetPolicyInput, optFns ...func(*lambda.Options)) (*lambda.GetPolicyOutput, error)
	GetFunctionCodeSigningConfig(ctx context.Context, params *lambda.GetFunctionCodeSigningConfigInput, optFns ...func(*lambda.Options)) (*lambda.GetFunctionCodeSigningConfigOutput, error)
	GetCodeSigningConfig(ctx context.Context, params *lambda.GetCodeSigningConfigInput, optFns ...func(*lambda.Options)) (*lambda.GetCodeSigningConfigOutput, error)
	GetLayerVersionPolicy(ctx context.Context, params *lambda.GetLayerVersionPolicyInput, optFns ...func(*lambda.Options)) (*lambda.GetLayerVersionPolicyOutput, error)
	GetFunctionUrlConfig(ctx context.Context, params *lambda.GetFunctionUrlConfigInput, optFns ...func(*lambda.Options)) (*lambda.GetFunctionUrlConfigOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_analyzer.go . AnalyzerClient
type AnalyzerClient interface {
	accessanalyzer.ListAnalyzersAPIClient
	accessanalyzer.ListFindingsAPIClient
	accessanalyzer.ListArchiveRulesAPIClient
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_waf.go . WafClient
type WafClient interface {
	ListWebACLs(ctx context.Context, params *waf.ListWebACLsInput, optFns ...func(*waf.Options)) (*waf.ListWebACLsOutput, error)
	GetWebACL(ctx context.Context, params *waf.GetWebACLInput, optFns ...func(*waf.Options)) (*waf.GetWebACLOutput, error)
	ListRuleGroups(ctx context.Context, params *waf.ListRuleGroupsInput, optFns ...func(*waf.Options)) (*waf.ListRuleGroupsOutput, error)
	GetRuleGroup(ctx context.Context, params *waf.GetRuleGroupInput, optFns ...func(*waf.Options)) (*waf.GetRuleGroupOutput, error)
	ListActivatedRulesInRuleGroup(ctx context.Context, params *waf.ListActivatedRulesInRuleGroupInput, optFns ...func(*waf.Options)) (*waf.ListActivatedRulesInRuleGroupOutput, error)
	ListSubscribedRuleGroups(ctx context.Context, params *waf.ListSubscribedRuleGroupsInput, optFns ...func(*waf.Options)) (*waf.ListSubscribedRuleGroupsOutput, error)
	ListRules(ctx context.Context, params *waf.ListRulesInput, optFns ...func(*waf.Options)) (*waf.ListRulesOutput, error)
	GetRule(ctx context.Context, params *waf.GetRuleInput, optFns ...func(*waf.Options)) (*waf.GetRuleOutput, error)
	ListTagsForResource(ctx context.Context, params *waf.ListTagsForResourceInput, optFns ...func(*waf.Options)) (*waf.ListTagsForResourceOutput, error)
	GetLoggingConfiguration(ctx context.Context, params *waf.GetLoggingConfigurationInput, optFns ...func(*waf.Options)) (*waf.GetLoggingConfigurationOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_wafv2.go . WafV2Client
type WafV2Client interface {
	ListWebACLs(ctx context.Context, params *wafv2.ListWebACLsInput, optFns ...func(*wafv2.Options)) (*wafv2.ListWebACLsOutput, error)
	GetWebACL(ctx context.Context, params *wafv2.GetWebACLInput, optFns ...func(*wafv2.Options)) (*wafv2.GetWebACLOutput, error)
	ListRuleGroups(ctx context.Context, params *wafv2.ListRuleGroupsInput, optFns ...func(*wafv2.Options)) (*wafv2.ListRuleGroupsOutput, error)
	GetRuleGroup(ctx context.Context, params *wafv2.GetRuleGroupInput, optFns ...func(*wafv2.Options)) (*wafv2.GetRuleGroupOutput, error)
	ListAvailableManagedRuleGroups(ctx context.Context, params *wafv2.ListAvailableManagedRuleGroupsInput, optFns ...func(*wafv2.Options)) (*wafv2.ListAvailableManagedRuleGroupsOutput, error)
	DescribeManagedRuleGroup(ctx context.Context, params *wafv2.DescribeManagedRuleGroupInput, optFns ...func(*wafv2.Options)) (*wafv2.DescribeManagedRuleGroupOutput, error)
	ListResourcesForWebACL(ctx context.Context, params *wafv2.ListResourcesForWebACLInput, optFns ...func(*wafv2.Options)) (*wafv2.ListResourcesForWebACLOutput, error)
	ListTagsForResource(ctx context.Context, params *wafv2.ListTagsForResourceInput, optFns ...func(*wafv2.Options)) (*wafv2.ListTagsForResourceOutput, error)
	GetPermissionPolicy(ctx context.Context, params *wafv2.GetPermissionPolicyInput, optFns ...func(*wafv2.Options)) (*wafv2.GetPermissionPolicyOutput, error)
	GetWebACLForResource(ctx context.Context, params *wafv2.GetWebACLForResourceInput, optFns ...func(*wafv2.Options)) (*wafv2.GetWebACLForResourceOutput, error)
	GetLoggingConfiguration(ctx context.Context, params *wafv2.GetLoggingConfigurationInput, optFns ...func(*wafv2.Options)) (*wafv2.GetLoggingConfigurationOutput, error)
	GetIPSet(ctx context.Context, params *wafv2.GetIPSetInput, optFns ...func(*wafv2.Options)) (*wafv2.GetIPSetOutput, error)
	GetRegexPatternSet(ctx context.Context, params *wafv2.GetRegexPatternSetInput, optFns ...func(*wafv2.Options)) (*wafv2.GetRegexPatternSetOutput, error)
	ListIPSets(ctx context.Context, params *wafv2.ListIPSetsInput, optFns ...func(*wafv2.Options)) (*wafv2.ListIPSetsOutput, error)
	ListRegexPatternSets(ctx context.Context, params *wafv2.ListRegexPatternSetsInput, optFns ...func(*wafv2.Options)) (*wafv2.ListRegexPatternSetsOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_sqs.go . SQSClient
type SQSClient interface {
	GetQueueAttributes(ctx context.Context, params *sqs.GetQueueAttributesInput, optFns ...func(*sqs.Options)) (*sqs.GetQueueAttributesOutput, error)
	ListQueues(ctx context.Context, params *sqs.ListQueuesInput, optFns ...func(*sqs.Options)) (*sqs.ListQueuesOutput, error)
	ListQueueTags(ctx context.Context, params *sqs.ListQueueTagsInput, optFns ...func(*sqs.Options)) (*sqs.ListQueueTagsOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_guardduty.go . GuardDutyClient
type GuardDutyClient interface {
	guardduty.ListDetectorsAPIClient
	guardduty.ListMembersAPIClient
	GetDetector(ctx context.Context, params *guardduty.GetDetectorInput, optFns ...func(*guardduty.Options)) (*guardduty.GetDetectorOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_ssm.go . SSMClient
type SSMClient interface {
	DescribeDocument(ctx context.Context, params *ssm.DescribeDocumentInput, optFns ...func(*ssm.Options)) (*ssm.DescribeDocumentOutput, error)
	DescribeDocumentPermission(ctx context.Context, params *ssm.DescribeDocumentPermissionInput, optFns ...func(*ssm.Options)) (*ssm.DescribeDocumentPermissionOutput, error)
	DescribeInstanceInformation(ctx context.Context, params *ssm.DescribeInstanceInformationInput, optFns ...func(*ssm.Options)) (*ssm.DescribeInstanceInformationOutput, error)
	ListComplianceItems(ctx context.Context, params *ssm.ListComplianceItemsInput, optFns ...func(*ssm.Options)) (*ssm.ListComplianceItemsOutput, error)
	ListDocuments(ctx context.Context, params *ssm.ListDocumentsInput, optFns ...func(*ssm.Options)) (*ssm.ListDocumentsOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_sagemaker.go . SageMakerClient
type SageMakerClient interface {
	ListNotebookInstances(ctx context.Context, params *sagemaker.ListNotebookInstancesInput, optFns ...func(*sagemaker.Options)) (*sagemaker.ListNotebookInstancesOutput, error)
	DescribeNotebookInstance(ctx context.Context, params *sagemaker.DescribeNotebookInstanceInput, optFns ...func(*sagemaker.Options)) (*sagemaker.DescribeNotebookInstanceOutput, error)
	ListTags(ctx context.Context, params *sagemaker.ListTagsInput, optFns ...func(*sagemaker.Options)) (*sagemaker.ListTagsOutput, error)
	ListModels(ctx context.Context, params *sagemaker.ListModelsInput, optFns ...func(*sagemaker.Options)) (*sagemaker.ListModelsOutput, error)
	DescribeModel(ctx context.Context, params *sagemaker.DescribeModelInput, optFns ...func(*sagemaker.Options)) (*sagemaker.DescribeModelOutput, error)
	ListEndpointConfigs(ctx context.Context, params *sagemaker.ListEndpointConfigsInput, optFns ...func(*sagemaker.Options)) (*sagemaker.ListEndpointConfigsOutput, error)
	DescribeEndpointConfig(ctx context.Context, params *sagemaker.DescribeEndpointConfigInput, optFns ...func(*sagemaker.Options)) (*sagemaker.DescribeEndpointConfigOutput, error)
	ListTrainingJobs(ctx context.Context, params *sagemaker.ListTrainingJobsInput, optFns ...func(*sagemaker.Options)) (*sagemaker.ListTrainingJobsOutput, error)
	DescribeTrainingJob(ctx context.Context, params *sagemaker.DescribeTrainingJobInput, optFns ...func(*sagemaker.Options)) (*sagemaker.DescribeTrainingJobOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/codebuild.go . CodebuildClient
type CodebuildClient interface {
	BatchGetProjects(ctx context.Context, params *codebuild.BatchGetProjectsInput, optFns ...func(*codebuild.Options)) (*codebuild.BatchGetProjectsOutput, error)
	ListProjects(ctx context.Context, params *codebuild.ListProjectsInput, optFns ...func(*codebuild.Options)) (*codebuild.ListProjectsOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_secrets_manager.go . SecretsManagerClient
type SecretsManagerClient interface {
	ListSecrets(ctx context.Context, params *secretsmanager.ListSecretsInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.ListSecretsOutput, error)
	DescribeSecret(ctx context.Context, params *secretsmanager.DescribeSecretInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.DescribeSecretOutput, error)
	GetResourcePolicy(ctx context.Context, params *secretsmanager.GetResourcePolicyInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.GetResourcePolicyOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_iot.go . IOTClient
type IOTClient interface {
	ListThings(ctx context.Context, params *iot.ListThingsInput, optFns ...func(*iot.Options)) (*iot.ListThingsOutput, error)
	ListThingTypes(ctx context.Context, params *iot.ListThingTypesInput, optFns ...func(*iot.Options)) (*iot.ListThingTypesOutput, error)
	ListThingGroups(ctx context.Context, params *iot.ListThingGroupsInput, optFns ...func(*iot.Options)) (*iot.ListThingGroupsOutput, error)
	DescribeThingGroup(ctx context.Context, params *iot.DescribeThingGroupInput, optFns ...func(*iot.Options)) (*iot.DescribeThingGroupOutput, error)
	ListThingsInThingGroup(ctx context.Context, params *iot.ListThingsInThingGroupInput, optFns ...func(*iot.Options)) (*iot.ListThingsInThingGroupOutput, error)
	ListBillingGroups(ctx context.Context, params *iot.ListBillingGroupsInput, optFns ...func(*iot.Options)) (*iot.ListBillingGroupsOutput, error)
	DescribeBillingGroup(ctx context.Context, params *iot.DescribeBillingGroupInput, optFns ...func(*iot.Options)) (*iot.DescribeBillingGroupOutput, error)
	ListTagsForResource(ctx context.Context, params *iot.ListTagsForResourceInput, optFns ...func(*iot.Options)) (*iot.ListTagsForResourceOutput, error)
	ListThingsInBillingGroup(ctx context.Context, params *iot.ListThingsInBillingGroupInput, optFns ...func(*iot.Options)) (*iot.ListThingsInBillingGroupOutput, error)
	ListThingPrincipals(ctx context.Context, params *iot.ListThingPrincipalsInput, optFns ...func(*iot.Options)) (*iot.ListThingPrincipalsOutput, error)
	ListStreams(ctx context.Context, params *iot.ListStreamsInput, optFns ...func(*iot.Options)) (*iot.ListStreamsOutput, error)
	DescribeStream(ctx context.Context, params *iot.DescribeStreamInput, optFns ...func(*iot.Options)) (*iot.DescribeStreamOutput, error)
	ListSecurityProfiles(ctx context.Context, params *iot.ListSecurityProfilesInput, optFns ...func(*iot.Options)) (*iot.ListSecurityProfilesOutput, error)
	DescribeSecurityProfile(ctx context.Context, params *iot.DescribeSecurityProfileInput, optFns ...func(*iot.Options)) (*iot.DescribeSecurityProfileOutput, error)
	ListCACertificates(ctx context.Context, params *iot.ListCACertificatesInput, optFns ...func(*iot.Options)) (*iot.ListCACertificatesOutput, error)
	DescribeCACertificate(ctx context.Context, params *iot.DescribeCACertificateInput, optFns ...func(*iot.Options)) (*iot.DescribeCACertificateOutput, error)
	ListCertificatesByCA(ctx context.Context, params *iot.ListCertificatesByCAInput, optFns ...func(*iot.Options)) (*iot.ListCertificatesByCAOutput, error)
	DescribeCertificate(ctx context.Context, params *iot.DescribeCertificateInput, optFns ...func(*iot.Options)) (*iot.DescribeCertificateOutput, error)
	ListCertificates(ctx context.Context, params *iot.ListCertificatesInput, optFns ...func(*iot.Options)) (*iot.ListCertificatesOutput, error)
	ListTargetsForSecurityProfile(ctx context.Context, params *iot.ListTargetsForSecurityProfileInput, optFns ...func(*iot.Options)) (*iot.ListTargetsForSecurityProfileOutput, error)
	ListAttachedPolicies(ctx context.Context, params *iot.ListAttachedPoliciesInput, optFns ...func(*iot.Options)) (*iot.ListAttachedPoliciesOutput, error)
	GetPolicy(ctx context.Context, params *iot.GetPolicyInput, optFns ...func(*iot.Options)) (*iot.GetPolicyOutput, error)
	ListPolicies(ctx context.Context, params *iot.ListPoliciesInput, optFns ...func(*iot.Options)) (*iot.ListPoliciesOutput, error)
	ListJobs(ctx context.Context, params *iot.ListJobsInput, optFns ...func(*iot.Options)) (*iot.ListJobsOutput, error)
	DescribeJob(ctx context.Context, params *iot.DescribeJobInput, optFns ...func(*iot.Options)) (*iot.DescribeJobOutput, error)
	GetTopicRule(ctx context.Context, params *iot.GetTopicRuleInput, optFns ...func(*iot.Options)) (*iot.GetTopicRuleOutput, error)
	ListTopicRules(ctx context.Context, params *iot.ListTopicRulesInput, optFns ...func(*iot.Options)) (*iot.ListTopicRulesOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_workspaces.go . WorkspacesClient
type WorkspacesClient interface {
	DescribeWorkspaces(ctx context.Context, params *workspaces.DescribeWorkspacesInput, optFns ...func(*workspaces.Options)) (*workspaces.DescribeWorkspacesOutput, error)
	DescribeWorkspaceDirectories(ctx context.Context, params *workspaces.DescribeWorkspaceDirectoriesInput, optFns ...func(*workspaces.Options)) (*workspaces.DescribeWorkspaceDirectoriesOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_codepipeline.go . CodePipelineClient
type CodePipelineClient interface {
	ListPipelines(ctx context.Context, params *codepipeline.ListPipelinesInput, optFns ...func(*codepipeline.Options)) (*codepipeline.ListPipelinesOutput, error)
	GetPipeline(ctx context.Context, params *codepipeline.GetPipelineInput, optFns ...func(*codepipeline.Options)) (*codepipeline.GetPipelineOutput, error)
	ListTagsForResource(ctx context.Context, params *codepipeline.ListTagsForResourceInput, optFns ...func(*codepipeline.Options)) (*codepipeline.ListTagsForResourceOutput, error)
	ListWebhooks(ctx context.Context, params *codepipeline.ListWebhooksInput, optFns ...func(*codepipeline.Options)) (*codepipeline.ListWebhooksOutput, error)
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

//go:generate mockgen -package=mocks -destination=./mocks/mock_qldb.go . QLDBClient
type QLDBClient interface {
	qldb.ListLedgersAPIClient
	qldb.ListJournalKinesisStreamsForLedgerAPIClient
	qldb.ListJournalS3ExportsForLedgerAPIClient
	DescribeLedger(ctx context.Context, params *qldb.DescribeLedgerInput, optFns ...func(*qldb.Options)) (*qldb.DescribeLedgerOutput, error)
	ListTagsForResource(ctx context.Context, params *qldb.ListTagsForResourceInput, optFns ...func(*qldb.Options)) (*qldb.ListTagsForResourceOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/wafregional.go . WafRegionalClient
type WafRegionalClient interface {
	GetRateBasedRule(ctx context.Context, params *wafregional.GetRateBasedRuleInput, optFns ...func(*wafregional.Options)) (*wafregional.GetRateBasedRuleOutput, error)
	GetRule(ctx context.Context, params *wafregional.GetRuleInput, optFns ...func(*wafregional.Options)) (*wafregional.GetRuleOutput, error)
	GetRuleGroup(ctx context.Context, params *wafregional.GetRuleGroupInput, optFns ...func(*wafregional.Options)) (*wafregional.GetRuleGroupOutput, error)
	GetWebACL(ctx context.Context, params *wafregional.GetWebACLInput, optFns ...func(*wafregional.Options)) (*wafregional.GetWebACLOutput, error)
	ListRateBasedRules(ctx context.Context, params *wafregional.ListRateBasedRulesInput, optFns ...func(*wafregional.Options)) (*wafregional.ListRateBasedRulesOutput, error)
	ListRuleGroups(ctx context.Context, params *wafregional.ListRuleGroupsInput, optFns ...func(*wafregional.Options)) (*wafregional.ListRuleGroupsOutput, error)
	ListRules(ctx context.Context, params *wafregional.ListRulesInput, optFns ...func(*wafregional.Options)) (*wafregional.ListRulesOutput, error)
	ListTagsForResource(ctx context.Context, params *wafregional.ListTagsForResourceInput, optFns ...func(*wafregional.Options)) (*wafregional.ListTagsForResourceOutput, error)
	ListWebACLs(ctx context.Context, params *wafregional.ListWebACLsInput, optFns ...func(*wafregional.Options)) (*wafregional.ListWebACLsOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/xray.go . XrayClient
type XrayClient interface {
	GetEncryptionConfig(ctx context.Context, params *xray.GetEncryptionConfigInput, optFns ...func(*xray.Options)) (*xray.GetEncryptionConfigOutput, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/shield.go . ShieldClient
type ShieldClient interface {
	ListProtections(ctx context.Context, params *shield.ListProtectionsInput, optFns ...func(*shield.Options)) (*shield.ListProtectionsOutput, error)
	ListTagsForResource(ctx context.Context, params *shield.ListTagsForResourceInput, optFns ...func(*shield.Options)) (*shield.ListTagsForResourceOutput, error)
	DescribeSubscription(ctx context.Context, params *shield.DescribeSubscriptionInput, optFns ...func(*shield.Options)) (*shield.DescribeSubscriptionOutput, error)
	ListAttacks(ctx context.Context, params *shield.ListAttacksInput, optFns ...func(*shield.Options)) (*shield.ListAttacksOutput, error)
	DescribeAttack(ctx context.Context, params *shield.DescribeAttackInput, optFns ...func(*shield.Options)) (*shield.DescribeAttackOutput, error)
	ListProtectionGroups(ctx context.Context, params *shield.ListProtectionGroupsInput, optFns ...func(*shield.Options)) (*shield.ListProtectionGroupsOutput, error)
}
