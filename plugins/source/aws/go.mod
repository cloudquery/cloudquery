module github.com/cloudquery/cloudquery/plugins/source/aws

go 1.19

require (
	github.com/aws/aws-sdk-go-v2 v1.17.5
	github.com/aws/aws-sdk-go-v2/config v1.18.15
	github.com/aws/aws-sdk-go-v2/credentials v1.13.15
	github.com/aws/aws-sdk-go-v2/feature/s3/manager v1.11.55
	github.com/aws/aws-sdk-go-v2/service/accessanalyzer v1.19.5
	github.com/aws/aws-sdk-go-v2/service/account v1.8.1
	github.com/aws/aws-sdk-go-v2/service/acm v1.17.5
	github.com/aws/aws-sdk-go-v2/service/amp v1.16.4
	github.com/aws/aws-sdk-go-v2/service/apigateway v1.16.5
	github.com/aws/aws-sdk-go-v2/service/apigatewayv2 v1.13.5
	github.com/aws/aws-sdk-go-v2/service/applicationautoscaling v1.17.5
	github.com/aws/aws-sdk-go-v2/service/apprunner v1.16.1
	github.com/aws/aws-sdk-go-v2/service/appstream v1.19.0
	github.com/aws/aws-sdk-go-v2/service/appsync v1.18.0
	github.com/aws/aws-sdk-go-v2/service/autoscaling v1.26.2
	github.com/aws/aws-sdk-go-v2/service/cloudformation v1.25.1
	github.com/aws/aws-sdk-go-v2/service/cloudfront v1.24.1
	github.com/aws/aws-sdk-go-v2/service/cloudhsmv2 v1.14.4
	github.com/aws/aws-sdk-go-v2/service/cloudtrail v1.22.1
	github.com/aws/aws-sdk-go-v2/service/cloudwatch v1.25.4
	github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs v1.20.5
	github.com/aws/aws-sdk-go-v2/service/codebuild v1.20.5
	github.com/aws/aws-sdk-go-v2/service/codepipeline v1.14.4
	github.com/aws/aws-sdk-go-v2/service/cognitoidentity v1.15.4
	github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider v1.22.4
	github.com/aws/aws-sdk-go-v2/service/configservice v1.29.5
	github.com/aws/aws-sdk-go-v2/service/databasemigrationservice v1.23.3
	github.com/aws/aws-sdk-go-v2/service/dax v1.12.2
	github.com/aws/aws-sdk-go-v2/service/directconnect v1.18.3
	github.com/aws/aws-sdk-go-v2/service/docdb v1.20.2
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.18.3
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.80.0
	github.com/aws/aws-sdk-go-v2/service/ecr v1.18.3
	github.com/aws/aws-sdk-go-v2/service/ecrpublic v1.15.1
	github.com/aws/aws-sdk-go-v2/service/ecs v1.23.2
	github.com/aws/aws-sdk-go-v2/service/efs v1.19.3
	github.com/aws/aws-sdk-go-v2/service/eks v1.27.2
	github.com/aws/aws-sdk-go-v2/service/elasticache v1.26.2
	github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk v1.15.1
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing v1.15.2
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 v1.19.3
	github.com/aws/aws-sdk-go-v2/service/elasticsearchservice v1.18.2
	github.com/aws/aws-sdk-go-v2/service/elastictranscoder v1.14.1
	github.com/aws/aws-sdk-go-v2/service/emr v1.22.2
	github.com/aws/aws-sdk-go-v2/service/eventbridge v1.17.0
	github.com/aws/aws-sdk-go-v2/service/firehose v1.16.2
	github.com/aws/aws-sdk-go-v2/service/frauddetector v1.21.0
	github.com/aws/aws-sdk-go-v2/service/fsx v1.28.2
	github.com/aws/aws-sdk-go-v2/service/glacier v1.14.2
	github.com/aws/aws-sdk-go-v2/service/guardduty v1.17.2
	github.com/aws/aws-sdk-go-v2/service/iam v1.19.0
	github.com/aws/aws-sdk-go-v2/service/identitystore v1.16.0
	github.com/aws/aws-sdk-go-v2/service/inspector v1.13.0
	github.com/aws/aws-sdk-go-v2/service/inspector2 v1.11.0
	github.com/aws/aws-sdk-go-v2/service/iot v1.33.0
	github.com/aws/aws-sdk-go-v2/service/kafka v1.19.0
	github.com/aws/aws-sdk-go-v2/service/kinesis v1.17.0
	github.com/aws/aws-sdk-go-v2/service/kms v1.20.0
	github.com/aws/aws-sdk-go-v2/service/lambda v1.29.0
	github.com/aws/aws-sdk-go-v2/service/lightsail v1.25.0
	github.com/aws/aws-sdk-go-v2/service/mq v1.14.0
	github.com/aws/aws-sdk-go-v2/service/mwaa v1.14.0
	github.com/aws/aws-sdk-go-v2/service/neptune v1.19.0
	github.com/aws/aws-sdk-go-v2/service/organizations v1.18.0
	github.com/aws/aws-sdk-go-v2/service/qldb v1.15.0
	github.com/aws/aws-sdk-go-v2/service/quicksight v1.30.0
	github.com/aws/aws-sdk-go-v2/service/ram v1.17.0
	github.com/aws/aws-sdk-go-v2/service/rds v1.40.0
	github.com/aws/aws-sdk-go-v2/service/redshift v1.27.0
	github.com/aws/aws-sdk-go-v2/service/resiliencehub v1.8.0
	github.com/aws/aws-sdk-go-v2/service/resourcegroups v1.14.0
	github.com/aws/aws-sdk-go-v2/service/route53 v1.26.0
	github.com/aws/aws-sdk-go-v2/service/route53domains v1.14.0
	github.com/aws/aws-sdk-go-v2/service/s3 v1.30.5
	github.com/aws/aws-sdk-go-v2/service/s3control v1.29.0
	github.com/aws/aws-sdk-go-v2/service/sagemaker v1.63.0
	github.com/aws/aws-sdk-go-v2/service/savingsplans v1.12.0
	github.com/aws/aws-sdk-go-v2/service/scheduler v1.1.0
	github.com/aws/aws-sdk-go-v2/service/secretsmanager v1.18.1
	github.com/aws/aws-sdk-go-v2/service/securityhub v1.27.1
	github.com/aws/aws-sdk-go-v2/service/servicecatalog v1.16.0
	github.com/aws/aws-sdk-go-v2/service/servicecatalogappregistry v1.16.1
	github.com/aws/aws-sdk-go-v2/service/servicequotas v1.14.0
	github.com/aws/aws-sdk-go-v2/service/ses v1.15.0
	github.com/aws/aws-sdk-go-v2/service/sesv2 v1.16.0
	github.com/aws/aws-sdk-go-v2/service/sfn v1.17.0
	github.com/aws/aws-sdk-go-v2/service/sns v1.19.0
	github.com/aws/aws-sdk-go-v2/service/sqs v1.20.0
	github.com/aws/aws-sdk-go-v2/service/ssm v1.35.0
	github.com/aws/aws-sdk-go-v2/service/ssoadmin v1.16.0
	github.com/aws/aws-sdk-go-v2/service/sts v1.18.5
	github.com/aws/aws-sdk-go-v2/service/timestreamwrite v1.15.0
	github.com/aws/aws-sdk-go-v2/service/waf v1.12.0
	github.com/aws/aws-sdk-go-v2/service/wafv2 v1.24.2
	github.com/aws/aws-sdk-go-v2/service/workspaces v1.28.0
	github.com/aws/aws-sdk-go-v2/service/xray v1.16.0
	github.com/aws/smithy-go v1.13.5
	github.com/basgys/goxml2json v1.1.0
	github.com/cloudquery/codegen v0.2.1
	github.com/cloudquery/plugin-sdk v1.40.0
	github.com/gocarina/gocsv v0.0.0-20230226133904-70c27cb2918a
	github.com/golang/mock v1.6.0
	github.com/google/go-cmp v0.5.9
	github.com/mitchellh/hashstructure/v2 v2.0.2
	github.com/mitchellh/mapstructure v1.5.0
	github.com/rs/zerolog v1.29.0
	github.com/stretchr/testify v1.8.2
	github.com/thoas/go-funk v0.9.3
	golang.org/x/sync v0.1.0
)

require (
	github.com/aws/aws-sdk-go-v2/service/amplify v1.13.4
	github.com/aws/aws-sdk-go-v2/service/support v1.14.1
)

require (
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.0.21 // indirect; indirect // indirect // indirect // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.1.24 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.14.4 // indirect
	github.com/bitly/go-simplejson v0.5.0 // indirect
	github.com/getsentry/sentry-go v0.18.0 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware/providers/zerolog/v2 v2.0.0-rc.3 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.0.0-rc.3 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/cobra v1.6.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/exp v0.0.0-20230224173230-c95f2b4c22f2 // indirect; indirect // indirect // indirect
)

require (
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.4.10 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.12.23 // indirect; indirect // indirect // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.29 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.23 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.3.30 // indirect; indirect // indirect // indirect
	github.com/aws/aws-sdk-go-v2/service/athena v1.22.4
	github.com/aws/aws-sdk-go-v2/service/backup v1.19.2
	github.com/aws/aws-sdk-go-v2/service/glue v1.40.0
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.9.11 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.7.22 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.9.23 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.13.23 // indirect
	github.com/aws/aws-sdk-go-v2/service/shield v1.18.0
	github.com/aws/aws-sdk-go-v2/service/sso v1.12.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/transfer v1.28.0
	github.com/aws/aws-sdk-go-v2/service/wafregional v1.13.1
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/pkg/errors v0.9.1
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/mod v0.6.0 // indirect
	golang.org/x/net v0.7.0 // indirect; indirect // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	golang.org/x/tools v0.2.0 // indirect
	google.golang.org/genproto v0.0.0-20230209215440-0dfe4f8abfcc // indirect; indirect // indirect
	google.golang.org/grpc v1.53.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
