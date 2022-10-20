module github.com/cloudquery/cloudquery/plugins/source/aws

go 1.19

require (
	github.com/aws/aws-sdk-go-v2 v1.16.16
	github.com/aws/aws-sdk-go-v2/config v1.17.8
	github.com/aws/aws-sdk-go-v2/credentials v1.12.21
	github.com/aws/aws-sdk-go-v2/feature/s3/manager v1.11.34
	github.com/aws/aws-sdk-go-v2/service/accessanalyzer v1.16.0
	github.com/aws/aws-sdk-go-v2/service/acm v1.15.0
	github.com/aws/aws-sdk-go-v2/service/apigateway v1.15.20
	github.com/aws/aws-sdk-go-v2/service/apigatewayv2 v1.12.18
	github.com/aws/aws-sdk-go-v2/service/applicationautoscaling v1.15.18
	github.com/aws/aws-sdk-go-v2/service/apprunner v1.13.0
	github.com/aws/aws-sdk-go-v2/service/appsync v1.15.10
	github.com/aws/aws-sdk-go-v2/service/autoscaling v1.23.16
	github.com/aws/aws-sdk-go-v2/service/cloudformation v1.22.10
	github.com/aws/aws-sdk-go-v2/service/cloudfront v1.20.5
	github.com/aws/aws-sdk-go-v2/service/cloudhsmv2 v1.13.19
	github.com/aws/aws-sdk-go-v2/service/cloudtrail v1.18.2
	github.com/aws/aws-sdk-go-v2/service/cloudwatch v1.21.6
	github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs v1.15.20
	github.com/aws/aws-sdk-go-v2/service/codebuild v1.19.17
	github.com/aws/aws-sdk-go-v2/service/codepipeline v1.13.17
	github.com/aws/aws-sdk-go-v2/service/cognitoidentity v1.14.1
	github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider v1.20.1
	github.com/aws/aws-sdk-go-v2/service/configservice v1.26.1
	github.com/aws/aws-sdk-go-v2/service/databasemigrationservice v1.21.12
	github.com/aws/aws-sdk-go-v2/service/dax v1.11.17
	github.com/aws/aws-sdk-go-v2/service/directconnect v1.17.18
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.17.1
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.63.1
	github.com/aws/aws-sdk-go-v2/service/ecr v1.17.18
	github.com/aws/aws-sdk-go-v2/service/ecrpublic v1.13.17
	github.com/aws/aws-sdk-go-v2/service/ecs v1.18.24
	github.com/aws/aws-sdk-go-v2/service/efs v1.17.15
	github.com/aws/aws-sdk-go-v2/service/eks v1.22.1
	github.com/aws/aws-sdk-go-v2/service/elasticache v1.22.10
	github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk v1.14.18
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing v1.14.18
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 v1.18.20
	github.com/aws/aws-sdk-go-v2/service/elasticsearchservice v1.16.10
	github.com/aws/aws-sdk-go-v2/service/emr v1.20.11
	github.com/aws/aws-sdk-go-v2/service/eventbridge v1.16.15
	github.com/aws/aws-sdk-go-v2/service/firehose v1.14.19
	github.com/aws/aws-sdk-go-v2/service/fsx v1.25.0
	github.com/aws/aws-sdk-go-v2/service/glacier v1.13.17
	github.com/aws/aws-sdk-go-v2/service/guardduty v1.16.0
	github.com/aws/aws-sdk-go-v2/service/iam v1.18.20
	github.com/aws/aws-sdk-go-v2/service/inspector v1.12.17
	github.com/aws/aws-sdk-go-v2/service/inspector2 v1.8.1
	github.com/aws/aws-sdk-go-v2/service/iot v1.29.1
	github.com/aws/aws-sdk-go-v2/service/kinesis v1.15.19
	github.com/aws/aws-sdk-go-v2/service/kms v1.18.12
	github.com/aws/aws-sdk-go-v2/service/lambda v1.24.6
	github.com/aws/aws-sdk-go-v2/service/lightsail v1.23.0
	github.com/aws/aws-sdk-go-v2/service/mq v1.13.13
	github.com/aws/aws-sdk-go-v2/service/neptune v1.17.12
	github.com/aws/aws-sdk-go-v2/service/organizations v1.16.13
	github.com/aws/aws-sdk-go-v2/service/qldb v1.14.18
	github.com/aws/aws-sdk-go-v2/service/rds v1.26.1
	github.com/aws/aws-sdk-go-v2/service/redshift v1.26.10
	github.com/aws/aws-sdk-go-v2/service/resourcegroups v1.12.18
	github.com/aws/aws-sdk-go-v2/service/route53 v1.22.2
	github.com/aws/aws-sdk-go-v2/service/route53domains v1.12.17
	github.com/aws/aws-sdk-go-v2/service/s3 v1.27.11
	github.com/aws/aws-sdk-go-v2/service/s3control v1.24.0
	github.com/aws/aws-sdk-go-v2/service/sagemaker v1.47.0
	github.com/aws/aws-sdk-go-v2/service/secretsmanager v1.16.2
	github.com/aws/aws-sdk-go-v2/service/sesv2 v1.13.18
	github.com/aws/aws-sdk-go-v2/service/sns v1.18.1
	github.com/aws/aws-sdk-go-v2/service/sqs v1.19.10
	github.com/aws/aws-sdk-go-v2/service/ssm v1.31.0
	github.com/aws/aws-sdk-go-v2/service/sts v1.16.19
	github.com/aws/aws-sdk-go-v2/service/waf v1.11.17
	github.com/aws/aws-sdk-go-v2/service/wafv2 v1.22.9
	github.com/aws/aws-sdk-go-v2/service/workspaces v1.23.0
	github.com/aws/aws-sdk-go-v2/service/xray v1.13.19
	github.com/aws/smithy-go v1.13.3
	github.com/basgys/goxml2json v1.1.0
	github.com/cloudquery/plugin-sdk v0.13.12
	github.com/ettle/strcase v0.1.1
	github.com/gocarina/gocsv v0.0.0-20220927221512-ad3251f9fa25
	github.com/golang/mock v1.6.0
	github.com/google/go-cmp v0.5.9
	github.com/mitchellh/mapstructure v1.5.0
	github.com/rs/zerolog v1.28.0
	github.com/stretchr/testify v1.8.0
	github.com/thoas/go-funk v0.9.2
	golang.org/x/sync v0.1.0
)

require (
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.0.14 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.1.18 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.13.6 // indirect
	github.com/bitly/go-simplejson v0.5.0 // indirect
	github.com/getsentry/sentry-go v0.14.0 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware/providers/zerolog/v2 v2.0.0-rc.2 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.0.0-rc.2 // indirect
	github.com/inconshreveable/mousetrap v1.0.1 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/cobra v1.5.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/xeipuuv/gojsonpointer v0.0.0-20190905194746-02993c407bfb // indirect
	github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415 // indirect
	github.com/xeipuuv/gojsonschema v1.2.0 // indirect
	golang.org/x/exp v0.0.0-20221012211006-4de253d81b95 // indirect
)

require (
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.4.8 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.12.17 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.23 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.17 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.3.24 // indirect
	github.com/aws/aws-sdk-go-v2/service/athena v1.18.10
	github.com/aws/aws-sdk-go-v2/service/backup v1.17.9
	github.com/aws/aws-sdk-go-v2/service/glue v1.33.0
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.9.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.7.17 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.9.17 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.13.17 // indirect
	github.com/aws/aws-sdk-go-v2/service/shield v1.17.9
	github.com/aws/aws-sdk-go-v2/service/sso v1.11.23 // indirect
	github.com/aws/aws-sdk-go-v2/service/transfer v1.23.0
	github.com/aws/aws-sdk-go-v2/service/wafregional v1.12.18
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/gofrs/uuid v4.3.0+incompatible // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/iancoleman/strcase v0.2.0
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pkg/errors v0.9.1
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/mod v0.6.0-dev.0.20220419223038-86c51ed26bb4 // indirect
	golang.org/x/net v0.0.0-20221004154528-8021a29435af // indirect
	golang.org/x/sys v0.1.0 // indirect
	golang.org/x/text v0.4.0 // indirect
	golang.org/x/tools v0.1.12 // indirect
	google.golang.org/genproto v0.0.0-20220930163606-c98284e70a91 // indirect
	google.golang.org/grpc v1.49.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
