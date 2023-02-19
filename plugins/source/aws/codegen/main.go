package main

import (
	"path"
	"reflect"
	"runtime"

	"github.com/aws/aws-sdk-go-v2/service/resiliencehub"

	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer"
	"github.com/aws/aws-sdk-go-v2/service/account"
	"github.com/aws/aws-sdk-go-v2/service/acm"
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
	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/aws/aws-sdk-go-v2/service/qldb"
	"github.com/aws/aws-sdk-go-v2/service/quicksight"
	"github.com/aws/aws-sdk-go-v2/service/ram"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroups"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/route53domains"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3control"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
	"github.com/aws/aws-sdk-go-v2/service/savingsplans"
	"github.com/aws/aws-sdk-go-v2/service/scheduler"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/securityhub"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalogappregistry"
	"github.com/aws/aws-sdk-go-v2/service/servicequotas"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sfn"
	"github.com/aws/aws-sdk-go-v2/service/shield"
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
	"github.com/aws/aws-sdk-go-v2/service/workspaces"
	"github.com/aws/aws-sdk-go-v2/service/xray"
	"github.com/cloudquery/codegen/interfaces"
	"github.com/thoas/go-funk"
)

var clients = []any{
	&accessanalyzer.Client{},
	&account.Client{},
	&acm.Client{},
	&amp.Client{},
	&amplify.Client{},
	&apigateway.Client{},
	&apigatewayv2.Client{},
	&applicationautoscaling.Client{},
	&apprunner.Client{},
	&appstream.Client{},
	&appsync.Client{},
	&athena.Client{},
	&autoscaling.Client{},
	&backup.Client{},
	&cloudformation.Client{},
	&cloudfront.Client{},
	&cloudhsmv2.Client{},
	&cloudtrail.Client{},
	&cloudwatch.Client{},
	&cloudwatchlogs.Client{},
	&codebuild.Client{},
	&codepipeline.Client{},
	&cognitoidentity.Client{},
	&cognitoidentityprovider.Client{},
	&configservice.Client{},
	&databasemigrationservice.Client{},
	&dax.Client{},
	&directconnect.Client{},
	&docdb.Client{},
	&dynamodb.Client{},
	&ec2.Client{},
	&ecr.Client{},
	&ecrpublic.Client{},
	&ecs.Client{},
	&efs.Client{},
	&eks.Client{},
	&elasticache.Client{},
	&elasticbeanstalk.Client{},
	&elasticloadbalancing.Client{},
	&elasticloadbalancingv2.Client{},
	&elasticsearchservice.Client{},
	&elastictranscoder.Client{},
	&emr.Client{},
	&eventbridge.Client{},
	&firehose.Client{},
	&frauddetector.Client{},
	&fsx.Client{},
	&glacier.Client{},
	&glue.Client{},
	&guardduty.Client{},
	&iam.Client{},
	&identitystore.Client{},
	&inspector.Client{},
	&inspector2.Client{},
	&iot.Client{},
	&kafka.Client{},
	&kinesis.Client{},
	&kms.Client{},
	&lambda.Client{},
	&lightsail.Client{},
	&mq.Client{},
	&mwaa.Client{},
	&neptune.Client{},
	&organizations.Client{},
	&qldb.Client{},
	&quicksight.Client{},
	&ram.Client{},
	&rds.Client{},
	&redshift.Client{},
	&resourcegroups.Client{},
	&resiliencehub.Client{},
	&route53.Client{},
	&route53domains.Client{},
	&s3.Client{},
	&s3control.Client{},
	&sagemaker.Client{},
	&savingsplans.Client{},
	&scheduler.Client{},
	&secretsmanager.Client{},
	&securityhub.Client{},
	&servicecatalog.Client{},
	&servicecatalogappregistry.Client{},
	&servicequotas.Client{},
	&ses.Client{},
	&sesv2.Client{},
	&sfn.Client{},
	&shield.Client{},
	&sns.Client{},
	&sqs.Client{},
	&ssm.Client{},
	&ssoadmin.Client{},
	&support.Client{},
	&timestreamwrite.Client{},
	&transfer.Client{},
	&waf.Client{},
	&wafregional.Client{},
	&wafv2.Client{},
	&workspaces.Client{},
	&xray.Client{},
}

// Generate the service interfaces under in client/services for use with mockgen.
func main() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("failed to get caller information")
	}
	err := interfaces.Generate(
		clients,
		path.Join(path.Dir(filename), "../client/services"),
		interfaces.WithIncludeFunc(include),
		interfaces.WithExtraImports(extraImports),
	)
	if err != nil {
		panic(err)
	}
}

func include(m reflect.Method) bool {
	// these methods will be included despite not starting with an accepted prefix
	var exceptions = []string{
		"QuerySchemaVersionMetadata",
		"GenerateCredentialReport",
	}
	if funk.ContainsString(exceptions, m.Name) {
		return true
	}
	return interfaces.MethodHasAnyPrefix(m, []string{"List", "Get", "Describe", "Search", "BatchGet"})
}

func extraImports(_ reflect.Method) []string {
	return []string{"context"}
}
