package provider

import (
	"embed"

	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/accessanalyzer"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/acm"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/apigateway"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/apigatewayv2"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/applicationautoscaling"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/appsync"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/athena"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/autoscaling"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/backup"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/cloudformation"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/cloudfront"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/cloudtrail"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/cloudwatch"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/cloudwatchlogs"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/codebuild"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/codepipeline"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/cognito"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/config"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/dax"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/directconnect"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/dms"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/dynamodb"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/ec2"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/ecr"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/ecs"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/efs"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/eks"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/elasticache"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/elasticbeanstalk"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/elasticsearch"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/elbv1"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/elbv2"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/emr"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/eventbridge"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/firehose"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/fsx"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/glue"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/guardduty"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/iam"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/inspector"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/inspector2"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/iot"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/kinesis"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/kms"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/lambda"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/lightsail"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/mq"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/organizations"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/qldb"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/rds"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/redshift"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/resourcegroups"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/route53"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/s3"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/sagemaker"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/secretsmanager"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/ses"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/shield"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/sns"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/sqs"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/ssm"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/transfer"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/waf"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/wafregional"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/wafv2"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/workspaces"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/xray"
	"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/module"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

var (
	//go:embed moduledata/*
	moduleData embed.FS

	Version = "Development"
)

func Provider() *provider.Provider {
	return &provider.Provider{
		Name:             "aws",
		Version:          Version,
		Configure:        client.Configure,
		ErrorClassifier:  client.ErrorClassifier,
		ModuleInfoReader: module.EmbeddedReader(moduleData, "moduledata"),
		ResourceMap: map[string]*schema.Table{
			"accessanalyzer.analyzers":                   accessanalyzer.Analyzers(),
			"acm.certificates":                           acm.Certificates(),
			"apigateway.api_keys":                        apigateway.APIKeys(),
			"apigateway.client_certificates":             apigateway.ClientCertificates(),
			"apigateway.domain_names":                    apigateway.DomainNames(),
			"apigateway.rest_apis":                       apigateway.RestApis(),
			"apigateway.usage_plans":                     apigateway.UsagePlans(),
			"apigateway.vpc_links":                       apigateway.VpcLinks(),
			"apigatewayv2.apis":                          apigatewayv2.Apigatewayv2Apis(),
			"apigatewayv2.domain_names":                  apigatewayv2.Apigatewayv2DomainNames(),
			"apigatewayv2.vpc_links":                     apigatewayv2.Apigatewayv2VpcLinks(),
			"applicationautoscaling.policies":            applicationautoscaling.ApplicationautoscalingPolicies(),
			"appsync.graphql_apis":                       appsync.GraphqlApis(),
			"athena.data_catalogs":                       athena.DataCatalogs(),
			"athena.work_groups":                         athena.WorkGroups(),
			"autoscaling.groups":                         autoscaling.AutoscalingGroups(),
			"autoscaling.launch_configurations":          autoscaling.AutoscalingLaunchConfigurations(),
			"autoscaling.scheduled_actions":              autoscaling.AutoscalingScheduledActions(),
			"aws.regions":                                ec2.AwsRegions(),
			"backup.global_settings":                     backup.GlobalSettings(),
			"backup.plans":                               backup.Plans(),
			"backup.region_settings":                     backup.RegionSettings(),
			"backup.vaults":                              backup.Vaults(),
			"cloudformation.stacks":                      cloudformation.Stacks(),
			"cloudfront.cache_policies":                  cloudfront.CloudfrontCachePolicies(),
			"cloudfront.distributions":                   cloudfront.CloudfrontDistributions(),
			"cloudtrail.trails":                          cloudtrail.CloudtrailTrails(),
			"cloudwatch.alarms":                          cloudwatch.CloudwatchAlarms(),
			"cloudwatchlogs.filters":                     cloudwatchlogs.CloudwatchlogsFilters(),
			"cloudwatchlogs.log_groups":                  cloudwatchlogs.LogGroups(),
			"codebuild.projects":                         codebuild.CodebuildProjects(),
			"codepipeline.pipelines":                     codepipeline.Pipelines(),
			"codepipeline.webhooks":                      codepipeline.Webhooks(),
			"cognito.identity_pools":                     cognito.CognitoIdentityPools(),
			"cognito.user_pools":                         cognito.CognitoUserPools(),
			"config.configuration_recorders":             config.ConfigConfigurationRecorders(),
			"config.conformance_packs":                   config.ConfigConformancePack(),
			"dax.clusters":                               dax.DaxClusters(),
			"directconnect.connections":                  directconnect.DirectconnectConnections(),
			"directconnect.gateways":                     directconnect.DirectconnectGateways(),
			"directconnect.lags":                         directconnect.DirectconnectLags(),
			"directconnect.virtual_gateways":             directconnect.DirectconnectVirtualGateways(),
			"directconnect.virtual_interfaces":           directconnect.DirectconnectVirtualInterfaces(),
			"dms.replication_instances":                  dms.DmsReplicationInstances(),
			"dynamodb.tables":                            dynamodb.DynamodbTables(),
			"ec2.byoip_cidrs":                            ec2.Ec2ByoipCidrs(),
			"ec2.customer_gateways":                      ec2.Ec2CustomerGateways(),
			"ec2.ebs_snapshots":                          ec2.Ec2EbsSnapshots(),
			"ec2.ebs_volumes":                            ec2.Ec2EbsVolumes(),
			"ec2.egress_only_internet_gateways":          ec2.EgressOnlyInternetGateways(),
			"ec2.eips":                                   ec2.Ec2Eips(),
			"ec2.flow_logs":                              ec2.Ec2FlowLogs(),
			"ec2.hosts":                                  ec2.Hosts(),
			"ec2.images":                                 ec2.Ec2Images(),
			"ec2.instance_statuses":                      ec2.Ec2InstanceStatuses(),
			"ec2.instance_types":                         ec2.InstanceTypes(),
			"ec2.instances":                              ec2.Ec2Instances(),
			"ec2.internet_gateways":                      ec2.Ec2InternetGateways(),
			"ec2.key_pairs":                              ec2.KeyPairs(),
			"ec2.nat_gateways":                           ec2.Ec2NatGateways(),
			"ec2.network_acls":                           ec2.Ec2NetworkAcls(),
			"ec2.network_interfaces":                     ec2.NetworkInterfaces(),
			"ec2.regional_config":                        ec2.Ec2RegionalConfig(),
			"ec2.route_tables":                           ec2.Ec2RouteTables(),
			"ec2.security_groups":                        ec2.Ec2SecurityGroups(),
			"ec2.subnets":                                ec2.Ec2Subnets(),
			"ec2.transit_gateways":                       ec2.Ec2TransitGateways(),
			"ec2.vpc_endpoint_service_configurations":    ec2.Ec2VpcEndpointServiceConfigurations(),
			"ec2.vpc_endpoint_services":                  ec2.Ec2VpcEndpointServices(),
			"ec2.vpc_endpoints":                          ec2.Ec2VpcEndpoints(),
			"ec2.vpc_peering_connections":                ec2.Ec2VpcPeeringConnections(),
			"ec2.vpcs":                                   ec2.Ec2Vpcs(),
			"ec2.vpn_gateways":                           ec2.Ec2VpnGateways(),
			"ecr.repositories":                           ecr.Repositories(),
			"ecs.clusters":                               ecs.Clusters(),
			"ecs.task_definitions":                       ecs.EcsTaskDefinitions(),
			"efs.filesystems":                            efs.EfsFilesystems(),
			"eks.clusters":                               eks.EksClusters(),
			"elasticache.clusters":                       elasticache.Clusters(),
			"elasticache.engine_versions":                elasticache.EngineVersions(),
			"elasticache.global_replication_groups":      elasticache.GlobalReplicationGroups(),
			"elasticache.parameter_groups":               elasticache.ParameterGroups(),
			"elasticache.replication_groups":             elasticache.ReplicationGroups(),
			"elasticache.reserved_cache_nodes":           elasticache.ReservedCacheNodes(),
			"elasticache.reserved_cache_nodes_offerings": elasticache.ReservedCacheNodesOfferings(),
			"elasticache.service_updates":                elasticache.ServiceUpdates(),
			"elasticache.snapshots":                      elasticache.Snapshots(),
			"elasticache.subnet_groups":                  elasticache.SubnetGroups(),
			"elasticache.user_groups":                    elasticache.UserGroups(),
			"elasticache.users":                          elasticache.Users(),
			"elasticbeanstalk.application_versions":      elasticbeanstalk.ApplicationVersions(),
			"elasticbeanstalk.applications":              elasticbeanstalk.ElasticbeanstalkApplications(),
			"elasticbeanstalk.environments":              elasticbeanstalk.ElasticbeanstalkEnvironments(),
			"elasticsearch.domains":                      elasticsearch.ElasticsearchDomains(),
			"elbv1.load_balancers":                       elbv1.Elbv1LoadBalancers(),
			"elbv2.load_balancers":                       elbv2.Elbv2LoadBalancers(),
			"elbv2.target_groups":                        elbv2.Elbv2TargetGroups(),
			"emr.block_public_access_configs":            emr.EmrBlockPublicAccessConfigs(),
			"emr.clusters":                               emr.EmrClusters(),
			"eventbridge.event_buses":                    eventbridge.EventBuses(),
			"fsx.backups":                                fsx.Backups(),
			"fsx.data_repo_associations":                 fsx.DataRepoAssociations(),
			"fsx.data_repo_tasks":                        fsx.DataRepoTasks(),
			"fsx.filesystems":                            fsx.Filesystems(),
			"fsx.snapshots":                              fsx.Snapshots(),
			"fsx.storage_vms":                            fsx.StorageVms(),
			"fsx.volumes":                                fsx.Volumes(),
			"glue.classifiers":                           glue.Classifiers(),
			"glue.connections":                           glue.Connections(),
			"glue.crawlers":                              glue.Crawlers(),
			"glue.databases":                             glue.Databases(),
			"glue.datacatalog_encryption_settings":       glue.DatacatalogEncryptionSettings(),
			"glue.dev_endpoints":                         glue.DevEndpoints(),
			"glue.jobs":                                  glue.Jobs(),
			"glue.ml_transforms":                         glue.MlTransforms(),
			"glue.registries":                            glue.Registries(),
			"glue.triggers":                              glue.Triggers(),
			"glue.security_configurations":               glue.SecurityConfigurations(),
			"glue.workflows":                             glue.Workflows(),
			"guardduty.detectors":                        guardduty.GuarddutyDetectors(),
			"iam.accounts":                               iam.IamAccounts(),
			"iam.groups":                                 iam.IamGroups(),
			"iam.openid_connect_identity_providers":      iam.IamOpenidConnectIdentityProviders(),
			"iam.password_policies":                      iam.IamPasswordPolicies(),
			"iam.policies":                               iam.IamPolicies(),
			"iam.roles":                                  iam.Roles(),
			"iam.saml_identity_providers":                iam.IamSamlIdentityProviders(),
			"iam.server_certificates":                    iam.IamServerCertificates(),
			"iam.users":                                  iam.IamUsers(),
			"iam.virtual_mfa_devices":                    iam.IamVirtualMfaDevices(),
			"inspector.findings":                         inspector.Findings(),
			"inspector2.findings":                        inspector2.Findings(),
			"iot.billing_groups":                         iot.IotBillingGroups(),
			"iot.ca_certificates":                        iot.IotCaCertificates(),
			"iot.certificates":                           iot.IotCertificates(),
			"iot.policies":                               iot.IotPolicies(),
			"iot.streams":                                iot.IotStreams(),
			"iot.thing_groups":                           iot.IotThingGroups(),
			"iot.thing_types":                            iot.IotThingTypes(),
			"iot.things":                                 iot.IotThings(),
			"iot.topic_rules":                            iot.IotTopicRules(),
			"kinesis.data_streams":                       kinesis.Streams(),
			"firehose.delivery_streams":                  firehose.DeliveryStreams(),
			"kms.keys":                                   kms.Keys(),
			"lambda.functions":                           lambda.Functions(),
			"lambda.layers":                              lambda.LambdaLayers(),
			"lambda.runtimes":                            lambda.LambdaRuntimes(),
			"lightsail.alarms":                           lightsail.Alarms(),
			"lightsail.buckets":                          lightsail.Buckets(),
			"lightsail.certificates":                     lightsail.Certificates(),
			"lightsail.container_services":               lightsail.ContainerServices(),
			"lightsail.database_snapshots":               lightsail.DatabaseSnapshots(),
			"lightsail.databases":                        lightsail.Databases(),
			"lightsail.disks":                            lightsail.Disks(),
			"lightsail.distributions":                    lightsail.Distributions(),
			"lightsail.instance_snapshots":               lightsail.InstanceSnapshots(),
			"lightsail.instances":                        lightsail.Instances(),
			"lightsail.load_balancers":                   lightsail.LoadBalancers(),
			"lightsail.static_ips":                       lightsail.StaticIps(),
			"mq.brokers":                                 mq.Brokers(),
			"organizations.accounts":                     organizations.Accounts(),
			"qldb.ledgers":                               qldb.Ledgers(),
			"rds.certificates":                           rds.RdsCertificates(),
			"rds.cluster_parameter_groups":               rds.RdsClusterParameterGroups(),
			"rds.cluster_snapshots":                      rds.RdsClusterSnapshots(),
			"rds.clusters":                               rds.RdsClusters(),
			"rds.db_parameter_groups":                    rds.RdsDbParameterGroups(),
			"rds.db_security_groups":                     rds.RdsDbSecurityGroups(),
			"rds.db_snapshots":                           rds.RdsDbSnapshots(),
			"rds.db_subnet_groups":                       rds.RdsSubnetGroups(),
			"rds.event_subscriptions":                    rds.RdsEventSubscriptions(),
			"rds.instances":                              rds.RdsInstances(),
			"redshift.clusters":                          redshift.RedshiftClusters(),
			"redshift.event_subscriptions":               redshift.EventSubscriptions(),
			"redshift.subnet_groups":                     redshift.RedshiftSubnetGroups(),
			"resourcegroups.resource_groups":             resourcegroups.ResourceGroups(),
			"route53.domains":                            route53.Route53Domains(),
			"route53.health_checks":                      route53.Route53HealthChecks(),
			"route53.hosted_zones":                       route53.Route53HostedZones(),
			"route53.reusable_delegation_sets":           route53.Route53ReusableDelegationSets(),
			"route53.traffic_policies":                   route53.Route53TrafficPolicies(),
			"s3.accounts":                                s3.Accounts(),
			"s3.buckets":                                 s3.Buckets(),
			"sagemaker.endpoint_configurations":          sagemaker.SagemakerEndpointConfigurations(),
			"sagemaker.models":                           sagemaker.SagemakerModels(),
			"sagemaker.notebook_instances":               sagemaker.SagemakerNotebookInstances(),
			"sagemaker.training_jobs":                    sagemaker.SagemakerTrainingJobs(),
			"secretsmanager.secrets":                     secretsmanager.SecretsmanagerSecrets(),
			"ses.templates":                              ses.Templates(),
			"shield.attacks":                             shield.Attacks(),
			"shield.protections_groups":                  shield.ProtectionGroups(),
			"shield.protections":                         shield.Protections(),
			"shield.subscriptions":                       shield.Subscriptions(),
			"sns.subscriptions":                          sns.Subscriptions(),
			"sns.topics":                                 sns.Topics(),
			"sqs.queues":                                 sqs.Queues(),
			"ssm.documents":                              ssm.SsmDocuments(),
			"ssm.instances":                              ssm.SsmInstances(),
			"ssm.parameters":                             ssm.Parameters(),
			"transfer.servers":                           transfer.Servers(),
			"waf.rule_groups":                            waf.WafRuleGroups(),
			"waf.rules":                                  waf.WafRules(),
			"waf.subscribed_rule_groups":                 waf.WafSubscribedRuleGroups(),
			"waf.web_acls":                               waf.WafWebAcls(),
			"wafregional.rate_based_rules":               wafregional.RateBasedRules(),
			"wafregional.rule_groups":                    wafregional.RuleGroups(),
			"wafregional.rules":                          wafregional.Rules(),
			"wafregional.web_acls":                       wafregional.WebAcls(),
			"wafv2.ipsets":                               wafv2.Ipsets(),
			"wafv2.managed_rule_groups":                  wafv2.Wafv2ManagedRuleGroups(),
			"wafv2.regex_pattern_sets":                   wafv2.RegexPatternSets(),
			"wafv2.rule_groups":                          wafv2.Wafv2RuleGroups(),
			"wafv2.web_acls":                             wafv2.Wafv2WebAcls(),
			"workspaces.directories":                     workspaces.Directories(),
			"workspaces.workspaces":                      workspaces.Workspaces(),
			"xray.encryption_config":                     xray.EncryptionConfigs(),
			"xray.groups":                                xray.Groups(),
			"xray.sampling_rules":                        xray.SamplingRules(),
			//"iot.security_profiles": 				 iot.IotSecurityProfiles(), //TODO disabled because of api error NotFoundException: No method found matching route security-profiles for http method GET.
		},
		Config: func() provider.Config {
			return &client.Config{}
		},
	}
}
