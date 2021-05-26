package client

// define interfaces so we can easily mock AWS calls
//go:generate mockgen -destination=./mocks/services.go -package=mocks . AutoscalingClient,ApigatewayClient,Apigatewayv2Client,CloudfrontClient,CloudtrailClient,CloudwatchClient,CloudwatchLogsClient,DirectconnectClient,Ec2Client,EcrClient,EfsClient,ElasticbeanstalkClient,ElbV1Client,ElbV2Client,EmrClient,FsxClient,IamClient,KmsClient,OrganizationsClient,RdsClient,S3Client,SnsClient,EksClient,RedshiftClient,Route53Client,EcsClient,S3ManagerClient,LambdaClient,ConfigServiceClient,AnalyzerClient

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/configservice"

	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/directconnect"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/efs"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
	elbv1 "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
	elbv2 "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/aws/aws-sdk-go-v2/service/emr"
	"github.com/aws/aws-sdk-go-v2/service/fsx"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

type AutoscalingClient interface {
	DescribeLaunchConfigurations(context.Context, *autoscaling.DescribeLaunchConfigurationsInput, ...func(*autoscaling.Options)) (*autoscaling.DescribeLaunchConfigurationsOutput, error)
}

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
}

type CloudfrontClient interface {
	ListDistributions(ctx context.Context, params *cloudfront.ListDistributionsInput, optFns ...func(*cloudfront.Options)) (*cloudfront.ListDistributionsOutput, error)
	ListCachePolicies(ctx context.Context, params *cloudfront.ListCachePoliciesInput, optFns ...func(*cloudfront.Options)) (*cloudfront.ListCachePoliciesOutput, error)
}

type CloudtrailClient interface {
	GetEventSelectors(ctx context.Context, params *cloudtrail.GetEventSelectorsInput, optFns ...func(*cloudtrail.Options)) (*cloudtrail.GetEventSelectorsOutput, error)
	DescribeTrails(ctx context.Context, params *cloudtrail.DescribeTrailsInput, optFns ...func(*cloudtrail.Options)) (*cloudtrail.DescribeTrailsOutput, error)
	GetTrailStatus(ctx context.Context, params *cloudtrail.GetTrailStatusInput, optFns ...func(*cloudtrail.Options)) (*cloudtrail.GetTrailStatusOutput, error)
}

type CloudwatchClient interface {
	DescribeAlarms(ctx context.Context, params *cloudwatch.DescribeAlarmsInput, optFns ...func(*cloudwatch.Options)) (*cloudwatch.DescribeAlarmsOutput, error)
}

type CloudwatchLogsClient interface {
	DescribeMetricFilters(ctx context.Context, params *cloudwatchlogs.DescribeMetricFiltersInput, optFns ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.DescribeMetricFiltersOutput, error)
}

type ConfigServiceClient interface {
	DescribeConfigurationRecorders(ctx context.Context, params *configservice.DescribeConfigurationRecordersInput, optFns ...func(*configservice.Options)) (*configservice.DescribeConfigurationRecordersOutput, error)
	configservice.DescribeConformancePacksAPIClient
}

type DirectconnectClient interface {
	DescribeDirectConnectGateways(ctx context.Context, params *directconnect.DescribeDirectConnectGatewaysInput, optFns ...func(*directconnect.Options)) (*directconnect.DescribeDirectConnectGatewaysOutput, error)
	DescribeVirtualGateways(ctx context.Context, params *directconnect.DescribeVirtualGatewaysInput, optFns ...func(*directconnect.Options)) (*directconnect.DescribeVirtualGatewaysOutput, error)
	DescribeVirtualInterfaces(ctx context.Context, params *directconnect.DescribeVirtualInterfacesInput, optFns ...func(*directconnect.Options)) (*directconnect.DescribeVirtualInterfacesOutput, error)
}
type Ec2Client interface {
	DescribeByoipCidrs(ctx context.Context, params *ec2.DescribeByoipCidrsInput, optFns ...func(*ec2.Options)) (*ec2.DescribeByoipCidrsOutput, error)
	DescribeCustomerGateways(ctx context.Context, params *ec2.DescribeCustomerGatewaysInput, optFns ...func(*ec2.Options)) (*ec2.DescribeCustomerGatewaysOutput, error)
	DescribeFlowLogs(ctx context.Context, params *ec2.DescribeFlowLogsInput, optFns ...func(*ec2.Options)) (*ec2.DescribeFlowLogsOutput, error)
	DescribeImages(ctx context.Context, params *ec2.DescribeImagesInput, optFns ...func(*ec2.Options)) (*ec2.DescribeImagesOutput, error)
	DescribeInstances(ctx context.Context, params *ec2.DescribeInstancesInput, optFns ...func(*ec2.Options)) (*ec2.DescribeInstancesOutput, error)
	DescribeInternetGateways(ctx context.Context, params *ec2.DescribeInternetGatewaysInput, optFns ...func(*ec2.Options)) (*ec2.DescribeInternetGatewaysOutput, error)
	DescribeNatGateways(ctx context.Context, params *ec2.DescribeNatGatewaysInput, optFns ...func(*ec2.Options)) (*ec2.DescribeNatGatewaysOutput, error)
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
	DescribeVpcs(ctx context.Context, params *ec2.DescribeVpcsInput, optFns ...func(*ec2.Options)) (*ec2.DescribeVpcsOutput, error)
	DescribeVpcEndpoints(ctx context.Context, params *ec2.DescribeVpcEndpointsInput, optFns ...func(*ec2.Options)) (*ec2.DescribeVpcEndpointsOutput, error)
}

type EcrClient interface {
	DescribeRepositories(ctx context.Context, params *ecr.DescribeRepositoriesInput, optFns ...func(*ecr.Options)) (*ecr.DescribeRepositoriesOutput, error)
	DescribeImages(ctx context.Context, params *ecr.DescribeImagesInput, optFns ...func(*ecr.Options)) (*ecr.DescribeImagesOutput, error)
}

type EfsClient interface {
	DescribeFileSystems(ctx context.Context, params *efs.DescribeFileSystemsInput, optFns ...func(*efs.Options)) (*efs.DescribeFileSystemsOutput, error)
}

type ElasticbeanstalkClient interface {
	DescribeEnvironments(ctx context.Context, params *elasticbeanstalk.DescribeEnvironmentsInput, optFns ...func(*elasticbeanstalk.Options)) (*elasticbeanstalk.DescribeEnvironmentsOutput, error)
}

type ElbV2Client interface {
	DescribeLoadBalancers(ctx context.Context, params *elbv2.DescribeLoadBalancersInput, optFns ...func(*elbv2.Options)) (*elbv2.DescribeLoadBalancersOutput, error)
	DescribeTargetGroups(ctx context.Context, params *elbv2.DescribeTargetGroupsInput, optFns ...func(*elbv2.Options)) (*elbv2.DescribeTargetGroupsOutput, error)
}

type ElbV1Client interface {
	DescribeLoadBalancers(ctx context.Context, params *elbv1.DescribeLoadBalancersInput, optFns ...func(*elbv1.Options)) (*elbv1.DescribeLoadBalancersOutput, error)
	DescribeLoadBalancerPolicies(ctx context.Context, params *elbv1.DescribeLoadBalancerPoliciesInput, optFns ...func(*elbv1.Options)) (*elbv1.DescribeLoadBalancerPoliciesOutput, error)
	DescribeTags(ctx context.Context, params *elbv1.DescribeTagsInput, optFns ...func(*elbv1.Options)) (*elbv1.DescribeTagsOutput, error)
	DescribeLoadBalancerAttributes(ctx context.Context, params *elbv1.DescribeLoadBalancerAttributesInput, optFns ...func(*elbv1.Options)) (*elbv1.DescribeLoadBalancerAttributesOutput, error)
}

type EmrClient interface {
	ListClusters(ctx context.Context, params *emr.ListClustersInput, optFns ...func(*emr.Options)) (*emr.ListClustersOutput, error)
}

type FsxClient interface {
	DescribeBackups(ctx context.Context, params *fsx.DescribeBackupsInput, optFns ...func(*fsx.Options)) (*fsx.DescribeBackupsOutput, error)
}

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

	iam.ListServerCertificatesAPIClient
	iam.ListAccountAliasesAPIClient
	GetAccountSummary(ctx context.Context, params *iam.GetAccountSummaryInput, optFns ...func(*iam.Options)) (*iam.GetAccountSummaryOutput, error)
}

type KmsClient interface {
	ListKeys(ctx context.Context, params *kms.ListKeysInput, optFns ...func(*kms.Options)) (*kms.ListKeysOutput, error)
	DescribeKey(ctx context.Context, params *kms.DescribeKeyInput, optFns ...func(*kms.Options)) (*kms.DescribeKeyOutput, error)
	GetKeyRotationStatus(ctx context.Context, params *kms.GetKeyRotationStatusInput, optFns ...func(*kms.Options)) (*kms.GetKeyRotationStatusOutput, error)
}

type OrganizationsClient interface {
	ListAccounts(ctx context.Context, params *organizations.ListAccountsInput, optFns ...func(*organizations.Options)) (*organizations.ListAccountsOutput, error)
}

type RdsClient interface {
	DescribeDBInstances(ctx context.Context, params *rds.DescribeDBInstancesInput, optFns ...func(*rds.Options)) (*rds.DescribeDBInstancesOutput, error)
	DescribeDBClusters(ctx context.Context, params *rds.DescribeDBClustersInput, optFns ...func(*rds.Options)) (*rds.DescribeDBClustersOutput, error)
	DescribeDBSubnetGroups(ctx context.Context, params *rds.DescribeDBSubnetGroupsInput, optFns ...func(*rds.Options)) (*rds.DescribeDBSubnetGroupsOutput, error)
	DescribeCertificates(ctx context.Context, params *rds.DescribeCertificatesInput, optFns ...func(*rds.Options)) (*rds.DescribeCertificatesOutput, error)
}

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
}

type SnsClient interface {
	ListTopics(ctx context.Context, params *sns.ListTopicsInput, optFns ...func(*sns.Options)) (*sns.ListTopicsOutput, error)
	ListSubscriptions(ctx context.Context, params *sns.ListSubscriptionsInput, optFns ...func(*sns.Options)) (*sns.ListSubscriptionsOutput, error)
	GetTopicAttributes(ctx context.Context, params *sns.GetTopicAttributesInput, optFns ...func(*sns.Options)) (*sns.GetTopicAttributesOutput, error)
}

type EcsClient interface {
	DescribeClusters(ctx context.Context, params *ecs.DescribeClustersInput, optFns ...func(*ecs.Options)) (*ecs.DescribeClustersOutput, error)
	ListClusters(ctx context.Context, params *ecs.ListClustersInput, optFns ...func(*ecs.Options)) (*ecs.ListClustersOutput, error)
}

type EksClient interface {
	ListClusters(ctx context.Context, params *eks.ListClustersInput, optFns ...func(*eks.Options)) (*eks.ListClustersOutput, error)
	DescribeCluster(ctx context.Context, params *eks.DescribeClusterInput, optFns ...func(*eks.Options)) (*eks.DescribeClusterOutput, error)
}

type RedshiftClient interface {
	DescribeClusters(ctx context.Context, params *redshift.DescribeClustersInput, optFns ...func(*redshift.Options)) (*redshift.DescribeClustersOutput, error)
	DescribeClusterSubnetGroups(ctx context.Context, params *redshift.DescribeClusterSubnetGroupsInput, optFns ...func(*redshift.Options)) (*redshift.DescribeClusterSubnetGroupsOutput, error)
}
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

type S3ManagerClient interface {
	GetBucketRegion(ctx context.Context, bucket string, optFns ...func(*s3.Options)) (string, error)
}

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
}

type AnalyzerClient interface {
	accessanalyzer.ListAnalyzersAPIClient
	accessanalyzer.ListFindingsAPIClient
}
