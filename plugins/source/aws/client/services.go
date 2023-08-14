package client

import (
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer"
	"github.com/aws/aws-sdk-go-v2/service/account"
	"github.com/aws/aws-sdk-go-v2/service/acm"
	"github.com/aws/aws-sdk-go-v2/service/acmpca"
	"github.com/aws/aws-sdk-go-v2/service/amp"
	"github.com/aws/aws-sdk-go-v2/service/amplify"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling"
	"github.com/aws/aws-sdk-go-v2/service/apprunner"
	"github.com/aws/aws-sdk-go-v2/service/appstream"
	"github.com/aws/aws-sdk-go-v2/service/appsync"
	"github.com/aws/aws-sdk-go-v2/service/athena"
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

func initServices(_ aws.Config, regions []string) Services {
	return Services{
		Regions: regions,
	}
}

type Services struct {
	Regions                   []string
	Accessanalyzer            services.AccessanalyzerClient
	Account                   services.AccountClient
	Acm                       services.AcmClient
	Acmpca                    services.AcmpcaClient
	Amp                       services.AmpClient
	Amplify                   services.AmplifyClient
	Apigateway                services.ApigatewayClient
	Apigatewayv2              services.Apigatewayv2Client
	Applicationautoscaling    services.ApplicationautoscalingClient
	Apprunner                 services.ApprunnerClient
	Appstream                 services.AppstreamClient
	Appsync                   services.AppsyncClient
	Athena                    services.AthenaClient
	Autoscaling               services.AutoscalingClient
	Autoscalingplans          services.AutoscalingplansClient
	Backup                    services.BackupClient
	Batch                     services.BatchClient
	Cloudformation            services.CloudformationClient
	Cloudfront                services.CloudfrontClient
	Cloudhsmv2                services.Cloudhsmv2Client
	Cloudtrail                services.CloudtrailClient
	Cloudwatch                services.CloudwatchClient
	Cloudwatchlogs            services.CloudwatchlogsClient
	Codeartifact              services.CodeartifactClient
	Codebuild                 services.CodebuildClient
	Codecommit                services.CodecommitClient
	Codepipeline              services.CodepipelineClient
	Cognitoidentity           services.CognitoidentityClient
	Cognitoidentityprovider   services.CognitoidentityproviderClient
	Computeoptimizer          services.ComputeoptimizerClient
	Configservice             services.ConfigserviceClient
	Costexplorer              services.CostexplorerClient
	Databasemigrationservice  services.DatabasemigrationserviceClient
	Dax                       services.DaxClient
	Detective                 services.DetectiveClient
	Directconnect             services.DirectconnectClient
	Docdb                     services.DocdbClient
	Dynamodb                  services.DynamodbClient
	Dynamodbstreams           services.DynamodbstreamsClient
	Ec2                       services.Ec2Client
	Ecr                       services.EcrClient
	Ecrpublic                 services.EcrpublicClient
	Ecs                       services.EcsClient
	Efs                       services.EfsClient
	Eks                       services.EksClient
	Elasticache               services.ElasticacheClient
	Elasticbeanstalk          services.ElasticbeanstalkClient
	Elasticloadbalancing      services.ElasticloadbalancingClient
	Elasticloadbalancingv2    services.Elasticloadbalancingv2Client
	Elasticsearchservice      services.ElasticsearchserviceClient
	Elastictranscoder         services.ElastictranscoderClient
	Emr                       services.EmrClient
	Eventbridge               services.EventbridgeClient
	Firehose                  services.FirehoseClient
	Frauddetector             services.FrauddetectorClient
	Fsx                       services.FsxClient
	Glacier                   services.GlacierClient
	Glue                      services.GlueClient
	Guardduty                 services.GuarddutyClient
	Iam                       services.IamClient
	Identitystore             services.IdentitystoreClient
	Inspector                 services.InspectorClient
	Inspector2                services.Inspector2Client
	Iot                       services.IotClient
	Kafka                     services.KafkaClient
	Kinesis                   services.KinesisClient
	Kms                       services.KmsClient
	Lambda                    services.LambdaClient
	Lightsail                 services.LightsailClient
	Mq                        services.MqClient
	Mwaa                      services.MwaaClient
	Neptune                   services.NeptuneClient
	Networkfirewall           services.NetworkfirewallClient
	Organizations             services.OrganizationsClient
	Qldb                      services.QldbClient
	Quicksight                services.QuicksightClient
	Ram                       services.RamClient
	Rds                       services.RdsClient
	Redshift                  services.RedshiftClient
	Resourcegroups            services.ResourcegroupsClient
	Resiliencehub             services.ResiliencehubClient
	Route53                   services.Route53Client
	Route53domains            services.Route53domainsClient
	Route53resolver           services.Route53resolverClient
	S3                        services.S3Client
	S3control                 services.S3controlClient
	Sagemaker                 services.SagemakerClient
	Savingsplans              services.SavingsplansClient
	Scheduler                 services.SchedulerClient
	Secretsmanager            services.SecretsmanagerClient
	Securityhub               services.SecurityhubClient
	Servicecatalog            services.ServicecatalogClient
	Servicecatalogappregistry services.ServicecatalogappregistryClient
	Servicediscovery          services.ServicediscoveryClient
	Servicequotas             services.ServicequotasClient
	Ses                       services.SesClient
	Sesv2                     services.Sesv2Client
	Sfn                       services.SfnClient
	Shield                    services.ShieldClient
	Signer                    services.SignerClient
	Sns                       services.SnsClient
	Sqs                       services.SqsClient
	Ssm                       services.SsmClient
	Ssoadmin                  services.SsoadminClient
	Support                   services.SupportClient
	Timestreamwrite           services.TimestreamwriteClient
	Transfer                  services.TransferClient
	Waf                       services.WafClient
	Wafregional               services.WafregionalClient
	Wafv2                     services.Wafv2Client
	Wellarchitected           services.WellarchitectedClient
	Workspaces                services.WorkspacesClient
	Xray                      services.XrayClient
}

func (s *Services) InitService(awsConfig *aws.Config, service string) {
	c := awsConfig.Copy()
	switch strings.ToLower(service) {
	case "accessanalyzer":
		s.Accessanalyzer = accessanalyzer.NewFromConfig(c)
	case "account":
		s.Account = account.NewFromConfig(c)
	case "acm":
		s.Acm = acm.NewFromConfig(c)
	case "acmpca":
		s.Acmpca = acmpca.NewFromConfig(c)
	case "amp":
		s.Amp = amp.NewFromConfig(c)
	case "amplify":
		s.Amplify = amplify.NewFromConfig(c)
	case "apigateway":
		s.Apigateway = apigateway.NewFromConfig(c)
	case "apigatewayv2":
		s.Apigatewayv2 = apigatewayv2.NewFromConfig(c)
	case "applicationautoscaling":
		s.Applicationautoscaling = applicationautoscaling.NewFromConfig(c)
	case "apprunner":
		s.Apprunner = apprunner.NewFromConfig(c)
	case "appstream":
		s.Appstream = appstream.NewFromConfig(c)
	case "appsync":
		s.Appsync = appsync.NewFromConfig(c)
	case "athena":
		s.Athena = athena.NewFromConfig(c)
	case "autoscaling":
		s.Autoscaling = autoscaling.NewFromConfig(c)
	case "autoscalingplans":
		s.Autoscalingplans = autoscalingplans.NewFromConfig(c)
	case "backup":
		s.Backup = backup.NewFromConfig(c)
	case "batch":
		s.Batch = batch.NewFromConfig(c)
	case "cloudformation":
		s.Cloudformation = cloudformation.NewFromConfig(c)
	case "cloudfront":
		s.Cloudfront = cloudfront.NewFromConfig(c)
	case "cloudhsmv2":
		s.Cloudhsmv2 = cloudhsmv2.NewFromConfig(c)
	case "cloudtrail":
		s.Cloudtrail = cloudtrail.NewFromConfig(c)
	case "cloudwatch":
		s.Cloudwatch = cloudwatch.NewFromConfig(c)
	case "cloudwatchlogs":
		s.Cloudwatchlogs = cloudwatchlogs.NewFromConfig(c)
	case "codeartifact":
		s.Codeartifact = codeartifact.NewFromConfig(c)
	case "codebuild":
		s.Codebuild = codebuild.NewFromConfig(c)
	case "codecommit":
		s.Codecommit = codecommit.NewFromConfig(c)
	case "codepipeline":
		s.Codepipeline = codepipeline.NewFromConfig(c)
	case "cognitoidentity":
		s.Cognitoidentity = cognitoidentity.NewFromConfig(c)
	case "cognitoidentityprovider":
		s.Cognitoidentityprovider = cognitoidentityprovider.NewFromConfig(c)
	case "computeoptimizer":
		s.Computeoptimizer = computeoptimizer.NewFromConfig(c)
	case "configservice":
		s.Configservice = configservice.NewFromConfig(c)
	case "costexplorer":
		s.Costexplorer = costexplorer.NewFromConfig(c)
	case "databasemigrationservice":
		s.Databasemigrationservice = databasemigrationservice.NewFromConfig(c)
	case "dax":
		s.Dax = dax.NewFromConfig(c)
	case "detective":
		s.Detective = detective.NewFromConfig(c)
	case "directconnect":
		s.Directconnect = directconnect.NewFromConfig(c)
	case "docdb":
		s.Docdb = docdb.NewFromConfig(c)
	case "dynamodb":
		s.Dynamodb = dynamodb.NewFromConfig(c)
	case "dynamodbstreams":
		s.Dynamodbstreams = dynamodbstreams.NewFromConfig(c)
	case "ec2":
		s.Ec2 = ec2.NewFromConfig(c)
	case "ecr":
		s.Ecr = ecr.NewFromConfig(c)
	case "ecrpublic":
		s.Ecrpublic = ecrpublic.NewFromConfig(c)
	case "ecs":
		s.Ecs = ecs.NewFromConfig(c)
	case "efs":
		s.Efs = efs.NewFromConfig(c)
	case "eks":
		s.Eks = eks.NewFromConfig(c)
	case "elasticache":
		s.Elasticache = elasticache.NewFromConfig(c)
	case "elasticbeanstalk":
		s.Elasticbeanstalk = elasticbeanstalk.NewFromConfig(c)
	case "elasticloadbalancing":
		s.Elasticloadbalancing = elasticloadbalancing.NewFromConfig(c)
	case "elasticloadbalancingv2":
		s.Elasticloadbalancingv2 = elasticloadbalancingv2.NewFromConfig(c)
	case "elasticsearchservice":
		s.Elasticsearchservice = elasticsearchservice.NewFromConfig(c)
	case "elastictranscoder":
		s.Elastictranscoder = elastictranscoder.NewFromConfig(c)
	case "emr":
		s.Emr = emr.NewFromConfig(c)
	case "eventbridge":
		s.Eventbridge = eventbridge.NewFromConfig(c)
	case "firehose":
		s.Firehose = firehose.NewFromConfig(c)
	case "frauddetector":
		s.Frauddetector = frauddetector.NewFromConfig(c)
	case "fsx":
		s.Fsx = fsx.NewFromConfig(c)
	case "glacier":
		s.Glacier = glacier.NewFromConfig(c)
	case "glue":
		s.Glue = glue.NewFromConfig(c)
	case "guardduty":
		s.Guardduty = guardduty.NewFromConfig(c)
	case "iam":
		s.Iam = iam.NewFromConfig(c)
	case "identitystore":
		s.Identitystore = identitystore.NewFromConfig(c)
	case "inspector":
		s.Inspector = inspector.NewFromConfig(c)
	case "inspector2":
		s.Inspector2 = inspector2.NewFromConfig(c)
	case "iot":
		s.Iot = iot.NewFromConfig(c)
	case "kafka":
		s.Kafka = kafka.NewFromConfig(c)
	case "kinesis":
		s.Kinesis = kinesis.NewFromConfig(c)
	case "kms":
		s.Kms = kms.NewFromConfig(c)
	case "lambda":
		s.Lambda = lambda.NewFromConfig(c)
	case "lightsail":
		s.Lightsail = lightsail.NewFromConfig(c)
	case "mq":
		s.Mq = mq.NewFromConfig(c)
	case "mwaa":
		s.Mwaa = mwaa.NewFromConfig(c)
	case "neptune":
		s.Neptune = neptune.NewFromConfig(c)
	case "networkfirewall":
		s.Networkfirewall = networkfirewall.NewFromConfig(c)
	case "organizations":
		s.Organizations = organizations.NewFromConfig(c)
	case "qldb":
		s.Qldb = qldb.NewFromConfig(c)
	case "quicksight":
		s.Quicksight = quicksight.NewFromConfig(c)
	case "ram":
		s.Ram = ram.NewFromConfig(c)
	case "rds":
		s.Rds = rds.NewFromConfig(c)
	case "redshift":
		s.Redshift = redshift.NewFromConfig(c)
	case "resiliencehub":
		s.Resiliencehub = resiliencehub.NewFromConfig(c)
	case "resourcegroups":
		s.Resourcegroups = resourcegroups.NewFromConfig(c)
	case "route53":
		s.Route53 = route53.NewFromConfig(c)
	case "route53domains":
		s.Route53domains = route53domains.NewFromConfig(c)
	case "route53resolver":
		s.Route53resolver = route53resolver.NewFromConfig(c)
	case "s3":
		s.S3 = s3.NewFromConfig(c)
	case "s3control":
		s.S3control = s3control.NewFromConfig(c)
	case "sagemaker":
		s.Sagemaker = sagemaker.NewFromConfig(c)
	case "savingsplans":
		s.Savingsplans = savingsplans.NewFromConfig(c)
	case "scheduler":
		s.Scheduler = scheduler.NewFromConfig(c)
	case "secretsmanager":
		s.Secretsmanager = secretsmanager.NewFromConfig(c)
	case "securityhub":
		s.Securityhub = securityhub.NewFromConfig(c)
	case "servicecatalog":
		s.Servicecatalog = servicecatalog.NewFromConfig(c)
	case "servicecatalogappregistry":
		s.Servicecatalogappregistry = servicecatalogappregistry.NewFromConfig(c)
	case "servicediscovery":
		s.Servicediscovery = servicediscovery.NewFromConfig(c)
	case "servicequotas":
		s.Servicequotas = servicequotas.NewFromConfig(c)
	case "ses":
		s.Ses = ses.NewFromConfig(c)
	case "sesv2":
		s.Sesv2 = sesv2.NewFromConfig(c)
	case "sfn":
		s.Sfn = sfn.NewFromConfig(c)
	case "shield":
		s.Shield = shield.NewFromConfig(c)
	case "signer":
		s.Signer = signer.NewFromConfig(c)
	case "sns":
		s.Sns = sns.NewFromConfig(c)
	case "sqs":
		s.Sqs = sqs.NewFromConfig(c)
	case "ssm":
		s.Ssm = ssm.NewFromConfig(c)
	case "ssoadmin":
		s.Ssoadmin = ssoadmin.NewFromConfig(c)
	case "support":
		s.Support = support.NewFromConfig(c)
	case "timestreamwrite":
		s.Timestreamwrite = timestreamwrite.NewFromConfig(c)
	case "transfer":
		s.Transfer = transfer.NewFromConfig(c)
	case "waf":
		s.Waf = waf.NewFromConfig(c)
	case "wafv2":
		s.Wafv2 = wafv2.NewFromConfig(c)
	case "wafregional":
		s.Wafregional = wafregional.NewFromConfig(c)
	case "wellarchitected":
		s.Wellarchitected = wellarchitected.NewFromConfig(c)
	case "workspaces":
		s.Workspaces = workspaces.NewFromConfig(c)
	case "xray":
		s.Xray = xray.NewFromConfig(c)
	}
}
