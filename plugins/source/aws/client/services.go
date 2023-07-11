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

func initServices(c aws.Config, regions []string) Services {
	awsCfg := c.Copy()
	return Services{
		Regions:                   regions,
		Accessanalyzer:            accessanalyzer.NewFromConfig(awsCfg),
		Account:                   account.NewFromConfig(awsCfg),
		Acm:                       acm.NewFromConfig(awsCfg),
		Acmpca:                    acmpca.NewFromConfig(awsCfg),
		Amp:                       amp.NewFromConfig(awsCfg),
		Amplify:                   amplify.NewFromConfig(awsCfg),
		Apigateway:                apigateway.NewFromConfig(awsCfg),
		Apigatewayv2:              apigatewayv2.NewFromConfig(awsCfg),
		Applicationautoscaling:    applicationautoscaling.NewFromConfig(awsCfg),
		Apprunner:                 apprunner.NewFromConfig(awsCfg),
		Appstream:                 appstream.NewFromConfig(awsCfg),
		Appsync:                   appsync.NewFromConfig(awsCfg),
		Athena:                    athena.NewFromConfig(awsCfg),
		Autoscaling:               autoscaling.NewFromConfig(awsCfg),
		Autoscalingplans:          autoscalingplans.NewFromConfig(awsCfg),
		Batch:                     batch.NewFromConfig(awsCfg),
		Backup:                    backup.NewFromConfig(awsCfg),
		Cloudformation:            cloudformation.NewFromConfig(awsCfg),
		Cloudfront:                cloudfront.NewFromConfig(awsCfg),
		Cloudhsmv2:                cloudhsmv2.NewFromConfig(awsCfg),
		Cloudtrail:                cloudtrail.NewFromConfig(awsCfg),
		Cloudwatch:                cloudwatch.NewFromConfig(awsCfg),
		Cloudwatchlogs:            cloudwatchlogs.NewFromConfig(awsCfg),
		Codeartifact:              codeartifact.NewFromConfig(awsCfg),
		Codebuild:                 codebuild.NewFromConfig(awsCfg),
		Codecommit:                codecommit.NewFromConfig(awsCfg),
		Codepipeline:              codepipeline.NewFromConfig(awsCfg),
		Cognitoidentity:           cognitoidentity.NewFromConfig(awsCfg),
		Cognitoidentityprovider:   cognitoidentityprovider.NewFromConfig(awsCfg),
		Computeoptimizer:          computeoptimizer.NewFromConfig(awsCfg),
		Configservice:             configservice.NewFromConfig(awsCfg),
		Costexplorer:              costexplorer.NewFromConfig(awsCfg),
		Databasemigrationservice:  databasemigrationservice.NewFromConfig(awsCfg),
		Dax:                       dax.NewFromConfig(awsCfg),
		Directconnect:             directconnect.NewFromConfig(awsCfg),
		Detective:                 detective.NewFromConfig(awsCfg),
		Docdb:                     docdb.NewFromConfig(awsCfg),
		Dynamodb:                  dynamodb.NewFromConfig(awsCfg),
		Dynamodbstreams:           dynamodbstreams.NewFromConfig(awsCfg),
		Ec2:                       ec2.NewFromConfig(awsCfg),
		Ecr:                       ecr.NewFromConfig(awsCfg),
		Ecrpublic:                 ecrpublic.NewFromConfig(awsCfg),
		Ecs:                       ecs.NewFromConfig(awsCfg),
		Efs:                       efs.NewFromConfig(awsCfg),
		Eks:                       eks.NewFromConfig(awsCfg),
		Elasticache:               elasticache.NewFromConfig(awsCfg),
		Elasticbeanstalk:          elasticbeanstalk.NewFromConfig(awsCfg),
		Elasticloadbalancing:      elasticloadbalancing.NewFromConfig(awsCfg),
		Elasticloadbalancingv2:    elasticloadbalancingv2.NewFromConfig(awsCfg),
		Elasticsearchservice:      elasticsearchservice.NewFromConfig(awsCfg),
		Elastictranscoder:         elastictranscoder.NewFromConfig(awsCfg),
		Emr:                       emr.NewFromConfig(awsCfg),
		Eventbridge:               eventbridge.NewFromConfig(awsCfg),
		Firehose:                  firehose.NewFromConfig(awsCfg),
		Frauddetector:             frauddetector.NewFromConfig(awsCfg),
		Fsx:                       fsx.NewFromConfig(awsCfg),
		Glacier:                   glacier.NewFromConfig(awsCfg),
		Glue:                      glue.NewFromConfig(awsCfg),
		Guardduty:                 guardduty.NewFromConfig(awsCfg),
		Iam:                       iam.NewFromConfig(awsCfg),
		Identitystore:             identitystore.NewFromConfig(awsCfg),
		Inspector:                 inspector.NewFromConfig(awsCfg),
		Inspector2:                inspector2.NewFromConfig(awsCfg),
		Iot:                       iot.NewFromConfig(awsCfg),
		Kafka:                     kafka.NewFromConfig(awsCfg),
		Kinesis:                   kinesis.NewFromConfig(awsCfg),
		Kms:                       kms.NewFromConfig(awsCfg),
		Lambda:                    lambda.NewFromConfig(awsCfg),
		Lightsail:                 lightsail.NewFromConfig(awsCfg),
		Mq:                        mq.NewFromConfig(awsCfg),
		Mwaa:                      mwaa.NewFromConfig(awsCfg),
		Neptune:                   neptune.NewFromConfig(awsCfg),
		Networkfirewall:           networkfirewall.NewFromConfig(awsCfg),
		Networkmanager:            networkmanager.NewFromConfig(awsCfg),
		Organizations:             organizations.NewFromConfig(awsCfg),
		Qldb:                      qldb.NewFromConfig(awsCfg),
		Quicksight:                quicksight.NewFromConfig(awsCfg),
		Ram:                       ram.NewFromConfig(awsCfg),
		Rds:                       rds.NewFromConfig(awsCfg),
		Redshift:                  redshift.NewFromConfig(awsCfg),
		Resiliencehub:             resiliencehub.NewFromConfig(awsCfg),
		Resourcegroups:            resourcegroups.NewFromConfig(awsCfg),
		Route53:                   route53.NewFromConfig(awsCfg),
		Route53domains:            route53domains.NewFromConfig(awsCfg),
		Route53resolver:           route53resolver.NewFromConfig(awsCfg),
		S3:                        s3.NewFromConfig(awsCfg),
		S3control:                 s3control.NewFromConfig(awsCfg),
		Sagemaker:                 sagemaker.NewFromConfig(awsCfg),
		Savingsplans:              savingsplans.NewFromConfig(awsCfg),
		Scheduler:                 scheduler.NewFromConfig(awsCfg),
		Secretsmanager:            secretsmanager.NewFromConfig(awsCfg),
		Securityhub:               securityhub.NewFromConfig(awsCfg),
		Servicecatalog:            servicecatalog.NewFromConfig(awsCfg),
		Servicecatalogappregistry: servicecatalogappregistry.NewFromConfig(awsCfg),
		Servicediscovery:          servicediscovery.NewFromConfig(awsCfg),
		Servicequotas:             servicequotas.NewFromConfig(awsCfg),
		Ses:                       ses.NewFromConfig(awsCfg),
		Sesv2:                     sesv2.NewFromConfig(awsCfg),
		Sfn:                       sfn.NewFromConfig(awsCfg),
		Shield:                    shield.NewFromConfig(awsCfg),
		Signer:                    signer.NewFromConfig(awsCfg),
		Sns:                       sns.NewFromConfig(awsCfg),
		Sqs:                       sqs.NewFromConfig(awsCfg),
		Ssm:                       ssm.NewFromConfig(awsCfg),
		Ssoadmin:                  ssoadmin.NewFromConfig(awsCfg),
		Support:                   support.NewFromConfig(awsCfg),
		Timestreamwrite:           timestreamwrite.NewFromConfig(awsCfg),
		Transfer:                  transfer.NewFromConfig(awsCfg),
		Waf:                       waf.NewFromConfig(awsCfg),
		Wafregional:               wafregional.NewFromConfig(awsCfg),
		Wafv2:                     wafv2.NewFromConfig(awsCfg),
		Wellarchitected:           wellarchitected.NewFromConfig(awsCfg),
		Workspaces:                workspaces.NewFromConfig(awsCfg),
		Xray:                      xray.NewFromConfig(awsCfg),
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
	Networkmanager            services.NetworkmanagerClient
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
