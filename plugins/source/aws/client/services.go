package client

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer"
	"github.com/aws/aws-sdk-go-v2/service/account"
	"github.com/aws/aws-sdk-go-v2/service/acm"
	"github.com/aws/aws-sdk-go-v2/service/acmpca"
	"github.com/aws/aws-sdk-go-v2/service/amp"
	"github.com/aws/aws-sdk-go-v2/service/amplify"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	"github.com/aws/aws-sdk-go-v2/service/appconfig"
	"github.com/aws/aws-sdk-go-v2/service/appflow"
	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling"
	"github.com/aws/aws-sdk-go-v2/service/appmesh"
	"github.com/aws/aws-sdk-go-v2/service/apprunner"
	"github.com/aws/aws-sdk-go-v2/service/appstream"
	"github.com/aws/aws-sdk-go-v2/service/appsync"
	"github.com/aws/aws-sdk-go-v2/service/athena"
	"github.com/aws/aws-sdk-go-v2/service/auditmanager"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/aws/aws-sdk-go-v2/service/autoscalingplans"
	"github.com/aws/aws-sdk-go-v2/service/backup"
	"github.com/aws/aws-sdk-go-v2/service/batch"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/aws/aws-sdk-go-v2/service/cloudhsmv2"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/codeartifact"
	"github.com/aws/aws-sdk-go-v2/service/codebuild"
	"github.com/aws/aws-sdk-go-v2/service/codecommit"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/computeoptimizer"
	"github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice"
	"github.com/aws/aws-sdk-go-v2/service/dax"
	"github.com/aws/aws-sdk-go-v2/service/detective"
	"github.com/aws/aws-sdk-go-v2/service/directconnect"
	"github.com/aws/aws-sdk-go-v2/service/docdb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodbstreams"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/ecrpublic"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/efs"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice"
	"github.com/aws/aws-sdk-go-v2/service/elastictranscoder"
	"github.com/aws/aws-sdk-go-v2/service/emr"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/aws/aws-sdk-go-v2/service/firehose"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector"
	"github.com/aws/aws-sdk-go-v2/service/fsx"
	"github.com/aws/aws-sdk-go-v2/service/glacier"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/guardduty"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/identitystore"
	"github.com/aws/aws-sdk-go-v2/service/inspector"
	"github.com/aws/aws-sdk-go-v2/service/inspector2"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/aws/aws-sdk-go-v2/service/kafka"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/aws/aws-sdk-go-v2/service/mq"
	"github.com/aws/aws-sdk-go-v2/service/mwaa"
	"github.com/aws/aws-sdk-go-v2/service/neptune"
	"github.com/aws/aws-sdk-go-v2/service/networkfirewall"
	"github.com/aws/aws-sdk-go-v2/service/networkmanager"
	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/aws/aws-sdk-go-v2/service/qldb"
	"github.com/aws/aws-sdk-go-v2/service/quicksight"
	"github.com/aws/aws-sdk-go-v2/service/ram"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/aws/aws-sdk-go-v2/service/resiliencehub"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroups"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/route53domains"
	"github.com/aws/aws-sdk-go-v2/service/route53recoverycontrolconfig"
	"github.com/aws/aws-sdk-go-v2/service/route53recoveryreadiness"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3control"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
	"github.com/aws/aws-sdk-go-v2/service/savingsplans"
	"github.com/aws/aws-sdk-go-v2/service/scheduler"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/securityhub"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalogappregistry"
	"github.com/aws/aws-sdk-go-v2/service/servicediscovery"
	"github.com/aws/aws-sdk-go-v2/service/servicequotas"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sfn"
	"github.com/aws/aws-sdk-go-v2/service/shield"
	"github.com/aws/aws-sdk-go-v2/service/signer"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin"
	"github.com/aws/aws-sdk-go-v2/service/support"
	"github.com/aws/aws-sdk-go-v2/service/timestreamwrite"
	"github.com/aws/aws-sdk-go-v2/service/transfer"
	"github.com/aws/aws-sdk-go-v2/service/waf"
	"github.com/aws/aws-sdk-go-v2/service/wafregional"
	"github.com/aws/aws-sdk-go-v2/service/wafv2"
	"github.com/aws/aws-sdk-go-v2/service/wellarchitected"
	"github.com/aws/aws-sdk-go-v2/service/workspaces"
	"github.com/aws/aws-sdk-go-v2/service/xray"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/services"
)

func initServices(config aws.Config, regions []string) Services {
	return Services{
		AWSConfig: config,
		Regions:   regions,
	}
}

type Services struct {
	AWSConfig                    aws.Config
	Regions                      []string
	Accessanalyzer               services.AccessanalyzerClient
	Account                      services.AccountClient
	Acm                          services.AcmClient
	Acmpca                       services.AcmpcaClient
	Amp                          services.AmpClient
	Amplify                      services.AmplifyClient
	Apigateway                   services.ApigatewayClient
	Apigatewayv2                 services.Apigatewayv2Client
	Appconfig                    services.AppconfigClient
	Appflow                      services.AppflowClient
	Applicationautoscaling       services.ApplicationautoscalingClient
	Appmesh                      services.AppmeshClient
	Apprunner                    services.ApprunnerClient
	Appstream                    services.AppstreamClient
	Appsync                      services.AppsyncClient
	Athena                       services.AthenaClient
	Auditmanager                 services.AuditmanagerClient
	Autoscaling                  services.AutoscalingClient
	Autoscalingplans             services.AutoscalingplansClient
	Backup                       services.BackupClient
	Batch                        services.BatchClient
	Cloudformation               services.CloudformationClient
	Cloudfront                   services.CloudfrontClient
	Cloudhsmv2                   services.Cloudhsmv2Client
	Cloudtrail                   services.CloudtrailClient
	Cloudwatch                   services.CloudwatchClient
	Cloudwatchlogs               services.CloudwatchlogsClient
	Codeartifact                 services.CodeartifactClient
	Codebuild                    services.CodebuildClient
	Codecommit                   services.CodecommitClient
	Codepipeline                 services.CodepipelineClient
	Cognitoidentity              services.CognitoidentityClient
	Cognitoidentityprovider      services.CognitoidentityproviderClient
	Computeoptimizer             services.ComputeoptimizerClient
	Configservice                services.ConfigserviceClient
	Costexplorer                 services.CostexplorerClient
	Databasemigrationservice     services.DatabasemigrationserviceClient
	Dax                          services.DaxClient
	Detective                    services.DetectiveClient
	Directconnect                services.DirectconnectClient
	Docdb                        services.DocdbClient
	Dynamodb                     services.DynamodbClient
	Dynamodbstreams              services.DynamodbstreamsClient
	Ec2                          services.Ec2Client
	Ecr                          services.EcrClient
	Ecrpublic                    services.EcrpublicClient
	Ecs                          services.EcsClient
	Efs                          services.EfsClient
	Eks                          services.EksClient
	Elasticache                  services.ElasticacheClient
	Elasticbeanstalk             services.ElasticbeanstalkClient
	Elasticloadbalancing         services.ElasticloadbalancingClient
	Elasticloadbalancingv2       services.Elasticloadbalancingv2Client
	Elasticsearchservice         services.ElasticsearchserviceClient
	Elastictranscoder            services.ElastictranscoderClient
	Emr                          services.EmrClient
	Eventbridge                  services.EventbridgeClient
	Firehose                     services.FirehoseClient
	Frauddetector                services.FrauddetectorClient
	Fsx                          services.FsxClient
	Glacier                      services.GlacierClient
	Glue                         services.GlueClient
	Guardduty                    services.GuarddutyClient
	Iam                          services.IamClient
	Identitystore                services.IdentitystoreClient
	Inspector                    services.InspectorClient
	Inspector2                   services.Inspector2Client
	Iot                          services.IotClient
	Kafka                        services.KafkaClient
	Kinesis                      services.KinesisClient
	Kms                          services.KmsClient
	Lambda                       services.LambdaClient
	Lightsail                    services.LightsailClient
	Mq                           services.MqClient
	Mwaa                         services.MwaaClient
	Neptune                      services.NeptuneClient
	Networkfirewall              services.NetworkfirewallClient
	Networkmanager               services.NetworkmanagerClient
	Organizations                services.OrganizationsClient
	Qldb                         services.QldbClient
	Quicksight                   services.QuicksightClient
	Ram                          services.RamClient
	Rds                          services.RdsClient
	Redshift                     services.RedshiftClient
	Resourcegroups               services.ResourcegroupsClient
	Resiliencehub                services.ResiliencehubClient
	Route53                      services.Route53Client
	Route53domains               services.Route53domainsClient
	Route53recoverycontrolconfig services.Route53recoverycontrolconfigClient
	Route53recoveryreadiness     services.Route53recoveryreadinessClient
	Route53resolver              services.Route53resolverClient
	S3                           services.S3Client
	S3control                    services.S3controlClient
	Sagemaker                    services.SagemakerClient
	Savingsplans                 services.SavingsplansClient
	Scheduler                    services.SchedulerClient
	Secretsmanager               services.SecretsmanagerClient
	Securityhub                  services.SecurityhubClient
	Servicecatalog               services.ServicecatalogClient
	Servicecatalogappregistry    services.ServicecatalogappregistryClient
	Servicediscovery             services.ServicediscoveryClient
	Servicequotas                services.ServicequotasClient
	Ses                          services.SesClient
	Sesv2                        services.Sesv2Client
	Sfn                          services.SfnClient
	Shield                       services.ShieldClient
	Signer                       services.SignerClient
	Sns                          services.SnsClient
	Sqs                          services.SqsClient
	Ssm                          services.SsmClient
	Ssoadmin                     services.SsoadminClient
	Support                      services.SupportClient
	Timestreamwrite              services.TimestreamwriteClient
	Transfer                     services.TransferClient
	Waf                          services.WafClient
	Wafregional                  services.WafregionalClient
	Wafv2                        services.Wafv2Client
	Wellarchitected              services.WellarchitectedClient
	Workspaces                   services.WorkspacesClient
	Xray                         services.XrayClient
}

func (s *Services) InitService(service AWSServiceName) {
	c := s.AWSConfig.Copy()
	switch service {
	case AWSServiceAccessanalyzer:
		s.Accessanalyzer = accessanalyzer.NewFromConfig(c)
	case AWSServiceAccount:
		s.Account = account.NewFromConfig(c)
	case AWSServiceAcm:
		s.Acm = acm.NewFromConfig(c)
	case AWSServiceAcmpca:
		s.Acmpca = acmpca.NewFromConfig(c)
	case AWSServiceAmp:
		s.Amp = amp.NewFromConfig(c)
	case AWSServiceAmplify:
		s.Amplify = amplify.NewFromConfig(c)
	case AWSServiceApigateway:
		s.Apigateway = apigateway.NewFromConfig(c)
	case AWSServiceApigatewayv2:
		s.Apigatewayv2 = apigatewayv2.NewFromConfig(c)
	case AWSServiceAppconfig:
		s.Appconfig = appconfig.NewFromConfig(c)
	case AWSServiceAppflow:
		s.Appflow = appflow.NewFromConfig(c)
	case AWSServiceApplicationautoscaling:
		s.Applicationautoscaling = applicationautoscaling.NewFromConfig(c)
	case AWSServiceAppmesh:
		s.Appmesh = appmesh.NewFromConfig(c)
	case AWSServiceApprunner:
		s.Apprunner = apprunner.NewFromConfig(c)
	case AWSServiceAppstream:
		s.Appstream = appstream.NewFromConfig(c)
	case AWSServiceAppsync:
		s.Appsync = appsync.NewFromConfig(c)
	case AWSServiceAthena:
		s.Athena = athena.NewFromConfig(c)
	case AWSServiceAuditmanager:
		s.Auditmanager = auditmanager.NewFromConfig(c)
	case AWSServiceAutoscaling:
		s.Autoscaling = autoscaling.NewFromConfig(c)
	case AWSServiceAutoscalingplans:
		s.Autoscalingplans = autoscalingplans.NewFromConfig(c)
	case AWSServiceBackup:
		s.Backup = backup.NewFromConfig(c)
	case AWSServiceBatch:
		s.Batch = batch.NewFromConfig(c)
	case AWSServiceCloudformation:
		s.Cloudformation = cloudformation.NewFromConfig(c)
	case AWSServiceCloudfront:
		s.Cloudfront = cloudfront.NewFromConfig(c)
	case AWSServiceCloudhsmv2:
		s.Cloudhsmv2 = cloudhsmv2.NewFromConfig(c)
	case AWSServiceCloudtrail:
		s.Cloudtrail = cloudtrail.NewFromConfig(c)
	case AWSServiceCloudwatch:
		s.Cloudwatch = cloudwatch.NewFromConfig(c)
	case AWSServiceCloudwatchlogs:
		s.Cloudwatchlogs = cloudwatchlogs.NewFromConfig(c)
	case AWSServiceCodeartifact:
		s.Codeartifact = codeartifact.NewFromConfig(c)
	case AWSServiceCodebuild:
		s.Codebuild = codebuild.NewFromConfig(c)
	case AWSServiceCodecommit:
		s.Codecommit = codecommit.NewFromConfig(c)
	case AWSServiceCodepipeline:
		s.Codepipeline = codepipeline.NewFromConfig(c)
	case AWSServiceCognitoidentity:
		s.Cognitoidentity = cognitoidentity.NewFromConfig(c)
	case AWSServiceCognitoidentityprovider:
		s.Cognitoidentityprovider = cognitoidentityprovider.NewFromConfig(c)
	case AWSServiceComputeoptimizer:
		s.Computeoptimizer = computeoptimizer.NewFromConfig(c)
	case AWSServiceConfigservice:
		s.Configservice = configservice.NewFromConfig(c)
	case AWSServiceCostexplorer:
		s.Costexplorer = costexplorer.NewFromConfig(c)
	case AWSServiceDatabasemigrationservice:
		s.Databasemigrationservice = databasemigrationservice.NewFromConfig(c)
	case AWSServiceDax:
		s.Dax = dax.NewFromConfig(c)
	case AWSServiceDetective:
		s.Detective = detective.NewFromConfig(c)
	case AWSServiceDirectconnect:
		s.Directconnect = directconnect.NewFromConfig(c)
	case AWSServiceDocdb:
		s.Docdb = docdb.NewFromConfig(c)
	case AWSServiceDynamodb:
		s.Dynamodb = dynamodb.NewFromConfig(c)
	case AWSServiceDynamodbstreams:
		s.Dynamodbstreams = dynamodbstreams.NewFromConfig(c)
	case AWSServiceEc2:
		s.Ec2 = ec2.NewFromConfig(c)
	case AWSServiceEcr:
		s.Ecr = ecr.NewFromConfig(c)
	case AWSServiceEcrpublic:
		s.Ecrpublic = ecrpublic.NewFromConfig(c)
	case AWSServiceEcs:
		s.Ecs = ecs.NewFromConfig(c)
	case AWSServiceEfs:
		s.Efs = efs.NewFromConfig(c)
	case AWSServiceEks:
		s.Eks = eks.NewFromConfig(c)
	case AWSServiceElasticache:
		s.Elasticache = elasticache.NewFromConfig(c)
	case AWSServiceElasticbeanstalk:
		s.Elasticbeanstalk = elasticbeanstalk.NewFromConfig(c)
	case AWSServiceElasticloadbalancing:
		s.Elasticloadbalancing = elasticloadbalancing.NewFromConfig(c)
	case AWSServiceElasticloadbalancingv2:
		s.Elasticloadbalancingv2 = elasticloadbalancingv2.NewFromConfig(c)
	case AWSServiceElasticsearchservice:
		s.Elasticsearchservice = elasticsearchservice.NewFromConfig(c)
	case AWSServiceElastictranscoder:
		s.Elastictranscoder = elastictranscoder.NewFromConfig(c)
	case AWSServiceEmr:
		s.Emr = emr.NewFromConfig(c)
	case AWSServiceEventbridge:
		s.Eventbridge = eventbridge.NewFromConfig(c)
	case AWSServiceFirehose:
		s.Firehose = firehose.NewFromConfig(c)
	case AWSServiceFrauddetector:
		s.Frauddetector = frauddetector.NewFromConfig(c)
	case AWSServiceFsx:
		s.Fsx = fsx.NewFromConfig(c)
	case AWSServiceGlacier:
		s.Glacier = glacier.NewFromConfig(c)
	case AWSServiceGlue:
		s.Glue = glue.NewFromConfig(c)
	case AWSServiceGuardduty:
		s.Guardduty = guardduty.NewFromConfig(c)
	case AWSServiceIam:
		s.Iam = iam.NewFromConfig(c)
	case AWSServiceIdentitystore:
		s.Identitystore = identitystore.NewFromConfig(c)
	case AWSServiceInspector:
		s.Inspector = inspector.NewFromConfig(c)
	case AWSServiceInspector2:
		s.Inspector2 = inspector2.NewFromConfig(c)
	case AWSServiceIot:
		s.Iot = iot.NewFromConfig(c)
	case AWSServiceKafka:
		s.Kafka = kafka.NewFromConfig(c)
	case AWSServiceKinesis:
		s.Kinesis = kinesis.NewFromConfig(c)
	case AWSServiceKms:
		s.Kms = kms.NewFromConfig(c)
	case AWSServiceLambda:
		s.Lambda = lambda.NewFromConfig(c)
	case AWSServiceLightsail:
		s.Lightsail = lightsail.NewFromConfig(c)
	case AWSServiceMq:
		s.Mq = mq.NewFromConfig(c)
	case AWSServiceMwaa:
		s.Mwaa = mwaa.NewFromConfig(c)
	case AWSServiceNeptune:
		s.Neptune = neptune.NewFromConfig(c)
	case AWSServiceNetworkfirewall:
		s.Networkfirewall = networkfirewall.NewFromConfig(c)
	case AWSServiceNetworkmanager:
		s.Networkmanager = networkmanager.NewFromConfig(c)
	case AWSServiceOrganizations:
		s.Organizations = organizations.NewFromConfig(c)
	case AWSServiceQldb:
		s.Qldb = qldb.NewFromConfig(c)
	case AWSServiceQuicksight:
		s.Quicksight = quicksight.NewFromConfig(c)
	case AWSServiceRam:
		s.Ram = ram.NewFromConfig(c)
	case AWSServiceRds:
		s.Rds = rds.NewFromConfig(c)
	case AWSServiceRedshift:
		s.Redshift = redshift.NewFromConfig(c)
	case AWSServiceResiliencehub:
		s.Resiliencehub = resiliencehub.NewFromConfig(c)
	case AWSServiceResourcegroups:
		s.Resourcegroups = resourcegroups.NewFromConfig(c)
	case AWSServiceRoute53:
		s.Route53 = route53.NewFromConfig(c)
	case AWSServiceRoute53domains:
		s.Route53domains = route53domains.NewFromConfig(c)
	case AWSServiceRoute53recoverycontrolconfig:
		s.Route53recoverycontrolconfig = route53recoverycontrolconfig.NewFromConfig(c)
	case AWSServiceRoute53recoveryreadiness:
		s.Route53recoveryreadiness = route53recoveryreadiness.NewFromConfig(c)
	case AWSServiceRoute53resolver:
		s.Route53resolver = route53resolver.NewFromConfig(c)
	case AWSServiceS3:
		s.S3 = s3.NewFromConfig(c)
	case AWSServiceS3control:
		s.S3control = s3control.NewFromConfig(c)
	case AWSServiceSagemaker:
		s.Sagemaker = sagemaker.NewFromConfig(c)
	case AWSServiceSavingsplans:
		s.Savingsplans = savingsplans.NewFromConfig(c)
	case AWSServiceScheduler:
		s.Scheduler = scheduler.NewFromConfig(c)
	case AWSServiceSecretsmanager:
		s.Secretsmanager = secretsmanager.NewFromConfig(c)
	case AWSServiceSecurityhub:
		s.Securityhub = securityhub.NewFromConfig(c)
	case AWSServiceServicecatalog:
		s.Servicecatalog = servicecatalog.NewFromConfig(c)
	case AWSServiceServicecatalogappregistry:
		s.Servicecatalogappregistry = servicecatalogappregistry.NewFromConfig(c)
	case AWSServiceServicediscovery:
		s.Servicediscovery = servicediscovery.NewFromConfig(c)
	case AWSServiceServicequotas:
		s.Servicequotas = servicequotas.NewFromConfig(c)
	case AWSServiceSes:
		s.Ses = ses.NewFromConfig(c)
	case AWSServiceSesv2:
		s.Sesv2 = sesv2.NewFromConfig(c)
	case AWSServiceSfn:
		s.Sfn = sfn.NewFromConfig(c)
	case AWSServiceShield:
		s.Shield = shield.NewFromConfig(c)
	case AWSServiceSigner:
		s.Signer = signer.NewFromConfig(c)
	case AWSServiceSns:
		s.Sns = sns.NewFromConfig(c)
	case AWSServiceSqs:
		s.Sqs = sqs.NewFromConfig(c)
	case AWSServiceSsm:
		s.Ssm = ssm.NewFromConfig(c)
	case AWSServiceSsoadmin:
		s.Ssoadmin = ssoadmin.NewFromConfig(c)
	case AWSServiceSupport:
		s.Support = support.NewFromConfig(c)
	case AWSServiceTimestreamwrite:
		s.Timestreamwrite = timestreamwrite.NewFromConfig(c)
	case AWSServiceTransfer:
		s.Transfer = transfer.NewFromConfig(c)
	case AWSServiceWaf:
		s.Waf = waf.NewFromConfig(c)
	case AWSServiceWafv2:
		s.Wafv2 = wafv2.NewFromConfig(c)
	case AWSServiceWafregional:
		s.Wafregional = wafregional.NewFromConfig(c)
	case AWSServiceWellarchitected:
		s.Wellarchitected = wellarchitected.NewFromConfig(c)
	case AWSServiceWorkspaces:
		s.Workspaces = workspaces.NewFromConfig(c)
	case AWSServiceXray:
		s.Xray = xray.NewFromConfig(c)
	default:
		panic("unknown service: " + service.String())
	}
}

func (s *Services) GetService(service AWSServiceName) any {
	switch service {
	case AWSServiceAccessanalyzer:
		return s.Accessanalyzer
	case AWSServiceAccount:
		return s.Account
	case AWSServiceAcm:
		return s.Acm
	case AWSServiceAcmpca:
		return s.Acmpca
	case AWSServiceAmp:
		return s.Amp
	case AWSServiceAmplify:
		return s.Amplify
	case AWSServiceApigateway:
		return s.Apigateway
	case AWSServiceApigatewayv2:
		return s.Apigatewayv2
	case AWSServiceAppconfig:
		return s.Appconfig
	case AWSServiceAppflow:
		return s.Appflow
	case AWSServiceApplicationautoscaling:
		return s.Applicationautoscaling
	case AWSServiceAppmesh:
		return s.Appmesh
	case AWSServiceApprunner:
		return s.Apprunner
	case AWSServiceAppstream:
		return s.Appstream
	case AWSServiceAppsync:
		return s.Appsync
	case AWSServiceAthena:
		return s.Athena
	case AWSServiceAuditmanager:
		return s.Auditmanager
	case AWSServiceAutoscaling:
		return s.Autoscaling
	case AWSServiceAutoscalingplans:
		return s.Autoscalingplans
	case AWSServiceBackup:
		return s.Backup
	case AWSServiceBatch:
		return s.Batch
	case AWSServiceCloudformation:
		return s.Cloudformation
	case AWSServiceCloudfront:
		return s.Cloudfront
	case AWSServiceCloudhsmv2:
		return s.Cloudhsmv2
	case AWSServiceCloudtrail:
		return s.Cloudtrail
	case AWSServiceCloudwatch:
		return s.Cloudwatch
	case AWSServiceCloudwatchlogs:
		return s.Cloudwatchlogs
	case AWSServiceCodeartifact:
		return s.Codeartifact
	case AWSServiceCodebuild:
		return s.Codebuild
	case AWSServiceCodecommit:
		return s.Codecommit
	case AWSServiceCodepipeline:
		return s.Codepipeline
	case AWSServiceCognitoidentity:
		return s.Cognitoidentity
	case AWSServiceCognitoidentityprovider:
		return s.Cognitoidentityprovider
	case AWSServiceComputeoptimizer:
		return s.Computeoptimizer
	case AWSServiceConfigservice:
		return s.Configservice
	case AWSServiceCostexplorer:
		return s.Costexplorer
	case AWSServiceDatabasemigrationservice:
		return s.Databasemigrationservice
	case AWSServiceDax:
		return s.Dax
	case AWSServiceDetective:
		return s.Detective
	case AWSServiceDirectconnect:
		return s.Directconnect
	case AWSServiceDocdb:
		return s.Docdb
	case AWSServiceDynamodb:
		return s.Dynamodb
	case AWSServiceDynamodbstreams:
		return s.Dynamodbstreams
	case AWSServiceEc2:
		return s.Ec2
	case AWSServiceEcr:
		return s.Ecr
	case AWSServiceEcrpublic:
		return s.Ecrpublic
	case AWSServiceEcs:
		return s.Ecs
	case AWSServiceEfs:
		return s.Efs
	case AWSServiceEks:
		return s.Eks
	case AWSServiceElasticache:
		return s.Elasticache
	case AWSServiceElasticbeanstalk:
		return s.Elasticbeanstalk
	case AWSServiceElasticloadbalancing:
		return s.Elasticloadbalancing
	case AWSServiceElasticloadbalancingv2:
		return s.Elasticloadbalancingv2
	case AWSServiceElasticsearchservice:
		return s.Elasticsearchservice
	case AWSServiceElastictranscoder:
		return s.Elastictranscoder
	case AWSServiceEmr:
		return s.Emr
	case AWSServiceEventbridge:
		return s.Eventbridge
	case AWSServiceFirehose:
		return s.Firehose
	case AWSServiceFrauddetector:
		return s.Frauddetector
	case AWSServiceFsx:
		return s.Fsx
	case AWSServiceGlacier:
		return s.Glacier
	case AWSServiceGlue:
		return s.Glue
	case AWSServiceGuardduty:
		return s.Guardduty
	case AWSServiceIam:
		return s.Iam
	case AWSServiceIdentitystore:
		return s.Identitystore
	case AWSServiceInspector:
		return s.Inspector
	case AWSServiceInspector2:
		return s.Inspector2
	case AWSServiceIot:
		return s.Iot
	case AWSServiceKafka:
		return s.Kafka
	case AWSServiceKinesis:
		return s.Kinesis
	case AWSServiceKms:
		return s.Kms
	case AWSServiceLambda:
		return s.Lambda
	case AWSServiceLightsail:
		return s.Lightsail
	case AWSServiceMq:
		return s.Mq
	case AWSServiceMwaa:
		return s.Mwaa
	case AWSServiceNeptune:
		return s.Neptune
	case AWSServiceNetworkfirewall:
		return s.Networkfirewall
	case AWSServiceNetworkmanager:
		return s.Networkmanager
	case AWSServiceOrganizations:
		return s.Organizations
	case AWSServiceQldb:
		return s.Qldb
	case AWSServiceQuicksight:
		return s.Quicksight
	case AWSServiceRam:
		return s.Ram
	case AWSServiceRds:
		return s.Rds
	case AWSServiceRedshift:
		return s.Redshift
	case AWSServiceResiliencehub:
		return s.Resiliencehub
	case AWSServiceResourcegroups:
		return s.Resourcegroups
	case AWSServiceRoute53:
		return s.Route53
	case AWSServiceRoute53domains:
		return s.Route53domains
	case AWSServiceRoute53recoverycontrolconfig:
		return s.Route53recoverycontrolconfig
	case AWSServiceRoute53recoveryreadiness:
		return s.Route53recoveryreadiness
	case AWSServiceRoute53resolver:
		return s.Route53resolver
	case AWSServiceS3:
		return s.S3
	case AWSServiceS3control:
		return s.S3control
	case AWSServiceSagemaker:
		return s.Sagemaker
	case AWSServiceSavingsplans:
		return s.Savingsplans
	case AWSServiceScheduler:
		return s.Scheduler
	case AWSServiceSecretsmanager:
		return s.Secretsmanager
	case AWSServiceSecurityhub:
		return s.Securityhub
	case AWSServiceServicecatalog:
		return s.Servicecatalog
	case AWSServiceServicecatalogappregistry:
		return s.Servicecatalogappregistry
	case AWSServiceServicediscovery:
		return s.Servicediscovery
	case AWSServiceServicequotas:
		return s.Servicequotas
	case AWSServiceSes:
		return s.Ses
	case AWSServiceSesv2:
		return s.Sesv2
	case AWSServiceSfn:
		return s.Sfn
	case AWSServiceShield:
		return s.Shield
	case AWSServiceSigner:
		return s.Signer
	case AWSServiceSns:
		return s.Sns
	case AWSServiceSqs:
		return s.Sqs
	case AWSServiceSsm:
		return s.Ssm
	case AWSServiceSsoadmin:
		return s.Ssoadmin
	case AWSServiceSupport:
		return s.Support
	case AWSServiceTimestreamwrite:
		return s.Timestreamwrite
	case AWSServiceTransfer:
		return s.Transfer
	case AWSServiceWaf:
		return s.Waf
	case AWSServiceWafv2:
		return s.Wafv2
	case AWSServiceWafregional:
		return s.Wafregional
	case AWSServiceWellarchitected:
		return s.Wellarchitected
	case AWSServiceWorkspaces:
		return s.Workspaces
	case AWSServiceXray:
		return s.Xray
	default:
		panic("unknown service: " + service.String())
	}
}
