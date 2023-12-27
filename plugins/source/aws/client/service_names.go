package client

const (
	AWSServiceAccessanalyzer AWSServiceName = iota
	AWSServiceAccount
	AWSServiceAcm
	AWSServiceAcmpca
	AWSServiceAmp
	AWSServiceAmplify
	AWSServiceApigateway
	AWSServiceApigatewayv2
	AWSServiceAppconfig
	AWSServiceAppflow
	AWSServiceApplicationautoscaling
	AWSServiceAppmesh
	AWSServiceApprunner
	AWSServiceAppstream
	AWSServiceAppsync
	AWSServiceAthena
	AWSServiceAuditmanager
	AWSServiceAutoscaling
	AWSServiceAutoscalingplans
	AWSServiceBackup
	AWSServiceBatch
	AWSServiceCloudformation
	AWSServiceCloudfront
	AWSServiceCloudhsmv2
	AWSServiceCloudtrail
	AWSServiceCloudwatch
	AWSServiceCloudwatchlogs
	AWSServiceCodeartifact
	AWSServiceCodebuild
	AWSServiceCodecommit
	AWSServiceCodepipeline
	AWSServiceCognitoidentity
	AWSServiceCognitoidentityprovider
	AWSServiceComputeoptimizer
	AWSServiceConfigservice
	AWSServiceCostexplorer
	AWSServiceDatabasemigrationservice
	AWSServiceDax
	AWSServiceDetective
	AWSServiceDirectconnect
	AWSServiceDocdb
	AWSServiceDynamodb
	AWSServiceDynamodbstreams
	AWSServiceEc2
	AWSServiceEcr
	AWSServiceEcrpublic
	AWSServiceEcs
	AWSServiceEfs
	AWSServiceEks
	AWSServiceElasticache
	AWSServiceElasticbeanstalk
	AWSServiceElasticloadbalancing
	AWSServiceElasticloadbalancingv2
	AWSServiceElasticsearchservice
	AWSServiceElastictranscoder
	AWSServiceEmr
	AWSServiceEventbridge
	AWSServiceFirehose
	AWSServiceFrauddetector
	AWSServiceFsx
	AWSServiceGlacier
	AWSServiceGlue
	AWSServiceGuardduty
	AWSServiceIam
	AWSServiceIdentitystore
	AWSServiceInspector
	AWSServiceInspector2
	AWSServiceIot
	AWSServiceKafka
	AWSServiceKinesis
	AWSServiceKms
	AWSServiceLambda
	AWSServiceLightsail
	AWSServiceMq
	AWSServiceMwaa
	AWSServiceNeptune
	AWSServiceNetworkfirewall
	AWSServiceNetworkmanager
	AWSServiceOrganizations
	AWSServiceQldb
	AWSServiceQuicksight
	AWSServiceRam
	AWSServiceRds
	AWSServiceRedshift
	AWSServiceResiliencehub
	AWSServiceResourcegroups
	AWSServiceRoute53
	AWSServiceRoute53domains
	AWSServiceRoute53recoverycontrolconfig
	AWSServiceRoute53recoveryreadiness
	AWSServiceRoute53resolver
	AWSServiceS3
	AWSServiceS3control
	AWSServiceSagemaker
	AWSServiceSavingsplans
	AWSServiceScheduler
	AWSServiceSecretsmanager
	AWSServiceSecurityhub
	AWSServiceServicecatalog
	AWSServiceServicecatalogappregistry
	AWSServiceServicediscovery
	AWSServiceServicequotas
	AWSServiceSes
	AWSServiceSesv2
	AWSServiceSfn
	AWSServiceShield
	AWSServiceSigner
	AWSServiceSns
	AWSServiceSqs
	AWSServiceSsm
	AWSServiceSsoadmin
	AWSServiceSupport
	AWSServiceTimestreamwrite
	AWSServiceTransfer
	AWSServiceWaf
	AWSServiceWafregional
	AWSServiceWafv2
	AWSServiceWellarchitected
	AWSServiceWorkspaces
	AWSServiceXray
)

type AWSServiceName int

func (s AWSServiceName) String() string {
	return AllAWSServiceNames[s]
}

var AllAWSServiceNames = [...]string{
	AWSServiceAccessanalyzer:               "accessanalyzer",
	AWSServiceAccount:                      "account",
	AWSServiceAcm:                          "acm",
	AWSServiceAcmpca:                       "acmpca",
	AWSServiceAmp:                          "amp",
	AWSServiceAmplify:                      "amplify",
	AWSServiceApigateway:                   "apigateway",
	AWSServiceApigatewayv2:                 "apigatewayv2",
	AWSServiceAppconfig:                    "appconfig",
	AWSServiceAppflow:                      "appflow",
	AWSServiceApplicationautoscaling:       "applicationautoscaling",
	AWSServiceAppmesh:                      "appmesh",
	AWSServiceApprunner:                    "apprunner",
	AWSServiceAppstream:                    "appstream",
	AWSServiceAppsync:                      "appsync",
	AWSServiceAthena:                       "athena",
	AWSServiceAuditmanager:                 "auditmanager",
	AWSServiceAutoscaling:                  "autoscaling",
	AWSServiceAutoscalingplans:             "autoscalingplans",
	AWSServiceBackup:                       "backup",
	AWSServiceBatch:                        "batch",
	AWSServiceCloudformation:               "cloudformation",
	AWSServiceCloudfront:                   "cloudfront",
	AWSServiceCloudhsmv2:                   "cloudhsmv2",
	AWSServiceCloudtrail:                   "cloudtrail",
	AWSServiceCloudwatch:                   "cloudwatch",
	AWSServiceCloudwatchlogs:               "cloudwatchlogs",
	AWSServiceCodeartifact:                 "codeartifact",
	AWSServiceCodebuild:                    "codebuild",
	AWSServiceCodecommit:                   "codecommit",
	AWSServiceCodepipeline:                 "codepipeline",
	AWSServiceCognitoidentity:              "cognitoidentity",
	AWSServiceCognitoidentityprovider:      "cognitoidentityprovider",
	AWSServiceComputeoptimizer:             "computeoptimizer",
	AWSServiceConfigservice:                "configservice",
	AWSServiceCostexplorer:                 "costexplorer",
	AWSServiceDatabasemigrationservice:     "databasemigrationservice",
	AWSServiceDax:                          "dax",
	AWSServiceDetective:                    "detective",
	AWSServiceDirectconnect:                "directconnect",
	AWSServiceDocdb:                        "docdb",
	AWSServiceDynamodb:                     "dynamodb",
	AWSServiceDynamodbstreams:              "dynamodbstreams",
	AWSServiceEc2:                          "ec2",
	AWSServiceEcr:                          "ecr",
	AWSServiceEcrpublic:                    "ecrpublic",
	AWSServiceEcs:                          "ecs",
	AWSServiceEfs:                          "efs",
	AWSServiceEks:                          "eks",
	AWSServiceElasticache:                  "elasticache",
	AWSServiceElasticbeanstalk:             "elasticbeanstalk",
	AWSServiceElasticloadbalancing:         "elasticloadbalancing",
	AWSServiceElasticloadbalancingv2:       "elasticloadbalancingv2",
	AWSServiceElasticsearchservice:         "elasticsearchservice",
	AWSServiceElastictranscoder:            "elastictranscoder",
	AWSServiceEmr:                          "emr",
	AWSServiceEventbridge:                  "eventbridge",
	AWSServiceFirehose:                     "firehose",
	AWSServiceFrauddetector:                "frauddetector",
	AWSServiceFsx:                          "fsx",
	AWSServiceGlacier:                      "glacier",
	AWSServiceGlue:                         "glue",
	AWSServiceGuardduty:                    "guardduty",
	AWSServiceIam:                          "iam",
	AWSServiceIdentitystore:                "identitystore",
	AWSServiceInspector:                    "inspector",
	AWSServiceInspector2:                   "inspector2",
	AWSServiceIot:                          "iot",
	AWSServiceKafka:                        "kafka",
	AWSServiceKinesis:                      "kinesis",
	AWSServiceKms:                          "kms",
	AWSServiceLambda:                       "lambda",
	AWSServiceLightsail:                    "lightsail",
	AWSServiceMq:                           "mq",
	AWSServiceMwaa:                         "mwaa",
	AWSServiceNeptune:                      "neptune",
	AWSServiceNetworkfirewall:              "networkfirewall",
	AWSServiceNetworkmanager:               "networkmanager",
	AWSServiceOrganizations:                "organizations",
	AWSServiceQldb:                         "qldb",
	AWSServiceQuicksight:                   "quicksight",
	AWSServiceRam:                          "ram",
	AWSServiceRds:                          "rds",
	AWSServiceRedshift:                     "redshift",
	AWSServiceResiliencehub:                "resiliencehub",
	AWSServiceResourcegroups:               "resourcegroups",
	AWSServiceRoute53:                      "route53",
	AWSServiceRoute53domains:               "route53domains",
	AWSServiceRoute53recoverycontrolconfig: "route53recoverycontrolconfig",
	AWSServiceRoute53recoveryreadiness:     "route53recoveryreadiness",
	AWSServiceRoute53resolver:              "route53resolver",
	AWSServiceS3:                           "s3",
	AWSServiceS3control:                    "s3control",
	AWSServiceSagemaker:                    "sagemaker",
	AWSServiceSavingsplans:                 "savingsplans",
	AWSServiceScheduler:                    "scheduler",
	AWSServiceSecretsmanager:               "secretsmanager",
	AWSServiceSecurityhub:                  "securityhub",
	AWSServiceServicecatalog:               "servicecatalog",
	AWSServiceServicecatalogappregistry:    "servicecatalogappregistry",
	AWSServiceServicediscovery:             "servicediscovery",
	AWSServiceServicequotas:                "servicequotas",
	AWSServiceSes:                          "ses",
	AWSServiceSesv2:                        "sesv2",
	AWSServiceSfn:                          "sfn",
	AWSServiceShield:                       "shield",
	AWSServiceSigner:                       "signer",
	AWSServiceSns:                          "sns",
	AWSServiceSqs:                          "sqs",
	AWSServiceSsm:                          "ssm",
	AWSServiceSsoadmin:                     "ssoadmin",
	AWSServiceSupport:                      "support",
	AWSServiceTimestreamwrite:              "timestreamwrite",
	AWSServiceTransfer:                     "transfer",
	AWSServiceWaf:                          "waf",
	AWSServiceWafregional:                  "wafregional",
	AWSServiceWafv2:                        "wafv2",
	AWSServiceWellarchitected:              "wellarchitected",
	AWSServiceWorkspaces:                   "workspaces",
	AWSServiceXray:                         "xray",
}
