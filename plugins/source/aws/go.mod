module github.com/cloudquery/cloudquery/plugins/source/aws

go 1.19

require (
	github.com/aws/aws-sdk-go v1.44.95
	github.com/aws/aws-sdk-go-v2 v1.16.11
	github.com/aws/aws-sdk-go-v2/config v1.15.14
	github.com/aws/aws-sdk-go-v2/credentials v1.12.9
	github.com/aws/aws-sdk-go-v2/feature/s3/manager v1.11.20
	github.com/aws/aws-sdk-go-v2/service/accessanalyzer v1.15.8
	github.com/aws/aws-sdk-go-v2/service/acm v1.14.12
	github.com/aws/aws-sdk-go-v2/service/apigateway v1.15.14
	github.com/aws/aws-sdk-go-v2/service/apigatewayv2 v1.12.8
	github.com/aws/aws-sdk-go-v2/service/applicationautoscaling v1.15.8
	github.com/aws/aws-sdk-go-v2/service/appsync v1.15.1
	github.com/aws/aws-sdk-go-v2/service/autoscaling v1.23.5
	github.com/aws/aws-sdk-go-v2/service/cloudformation v1.21.2
	github.com/aws/aws-sdk-go-v2/service/cloudfront v1.18.4
	github.com/aws/aws-sdk-go-v2/service/cloudtrail v1.16.4
	github.com/aws/aws-sdk-go-v2/service/cloudwatch v1.18.6
	github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs v1.15.10
	github.com/aws/aws-sdk-go-v2/service/codebuild v1.19.8
	github.com/aws/aws-sdk-go-v2/service/codepipeline v1.13.8
	github.com/aws/aws-sdk-go-v2/service/cognitoidentity v1.13.8
	github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider v1.17.3
	github.com/aws/aws-sdk-go-v2/service/configservice v1.21.4
	github.com/aws/aws-sdk-go-v2/service/databasemigrationservice v1.20.0
	github.com/aws/aws-sdk-go-v2/service/dax v1.11.8
	github.com/aws/aws-sdk-go-v2/service/directconnect v1.17.8
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.15.9
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.47.1
	github.com/aws/aws-sdk-go-v2/service/ecr v1.17.8
	github.com/aws/aws-sdk-go-v2/service/ecs v1.18.11
	github.com/aws/aws-sdk-go-v2/service/efs v1.17.6
	github.com/aws/aws-sdk-go-v2/service/eks v1.21.4
	github.com/aws/aws-sdk-go-v2/service/elasticache v1.22.1
	github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk v1.14.8
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing v1.14.7
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 v1.18.7
	github.com/aws/aws-sdk-go-v2/service/elasticsearchservice v1.15.7
	github.com/aws/aws-sdk-go-v2/service/emr v1.20.0
	github.com/aws/aws-sdk-go-v2/service/eventbridge v1.16.8
	github.com/aws/aws-sdk-go-v2/service/firehose v1.14.10
	github.com/aws/aws-sdk-go-v2/service/fsx v1.24.2
	github.com/aws/aws-sdk-go-v2/service/guardduty v1.14.1
	github.com/aws/aws-sdk-go-v2/service/iam v1.18.8
	github.com/aws/aws-sdk-go-v2/service/inspector v1.12.11
	github.com/aws/aws-sdk-go-v2/service/inspector2 v1.7.3
	github.com/aws/aws-sdk-go-v2/service/iot v1.25.4
	github.com/aws/aws-sdk-go-v2/service/kinesis v1.15.9
	github.com/aws/aws-sdk-go-v2/service/kms v1.17.4
	github.com/aws/aws-sdk-go-v2/service/lambda v1.23.3
	github.com/aws/aws-sdk-go-v2/service/lightsail v1.22.2
	github.com/aws/aws-sdk-go-v2/service/mq v1.13.3
	github.com/aws/aws-sdk-go-v2/service/organizations v1.16.3
	github.com/aws/aws-sdk-go-v2/service/qldb v1.14.8
	github.com/aws/aws-sdk-go-v2/service/rds v1.21.5
	github.com/aws/aws-sdk-go-v2/service/redshift v1.25.1
	github.com/aws/aws-sdk-go-v2/service/resourcegroups v1.12.9
	github.com/aws/aws-sdk-go-v2/service/route53 v1.21.2
	github.com/aws/aws-sdk-go-v2/service/route53domains v1.12.7
	github.com/aws/aws-sdk-go-v2/service/s3 v1.27.1
	github.com/aws/aws-sdk-go-v2/service/s3control v1.21.8
	github.com/aws/aws-sdk-go-v2/service/sagemaker v1.34.0
	github.com/aws/aws-sdk-go-v2/service/secretsmanager v1.15.12
	github.com/aws/aws-sdk-go-v2/service/sesv2 v1.13.8
	github.com/aws/aws-sdk-go-v2/service/sns v1.17.13
	github.com/aws/aws-sdk-go-v2/service/sqs v1.19.4
	github.com/aws/aws-sdk-go-v2/service/ssm v1.27.3
	github.com/aws/aws-sdk-go-v2/service/sts v1.16.9
	github.com/aws/aws-sdk-go-v2/service/waf v1.11.7
	github.com/aws/aws-sdk-go-v2/service/wafv2 v1.20.4
	github.com/aws/aws-sdk-go-v2/service/workspaces v1.19.1
	github.com/aws/aws-sdk-go-v2/service/xray v1.13.8
	github.com/aws/smithy-go v1.12.1
	github.com/basgys/goxml2json v1.1.0
	github.com/bxcodec/faker v2.0.1+incompatible
	github.com/cloudquery/faker/v3 v3.7.7
	github.com/cloudquery/plugin-sdk v0.7.12
	github.com/ettle/strcase v0.1.1
	github.com/gocarina/gocsv v0.0.0-20220823132111-71f3a5cb2654
	github.com/golang/mock v1.6.0
	github.com/google/go-cmp v0.5.9
	github.com/mitchellh/mapstructure v1.5.0
	github.com/rs/zerolog v1.28.0
	github.com/stretchr/testify v1.8.0
	github.com/thoas/go-funk v0.9.2
	golang.org/x/sync v0.0.0-20220907140024-f12130a52804
)

require (
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.0.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.1.9 // indirect
	github.com/bitly/go-simplejson v0.5.0 // indirect
	github.com/getsentry/sentry-go v0.13.0 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware/providers/zerolog/v2 v2.0.0-rc.2 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.0.0-rc.2 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/cobra v1.5.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/xeipuuv/gojsonpointer v0.0.0-20180127040702-4e3ac2762d5f // indirect
	github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415 // indirect
	github.com/xeipuuv/gojsonschema v1.2.0 // indirect
)

require (
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.4.3 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.12.8 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.18 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.12 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.3.15 // indirect
	github.com/aws/aws-sdk-go-v2/service/athena v1.16.0
	github.com/aws/aws-sdk-go-v2/service/backup v1.16.3
	github.com/aws/aws-sdk-go-v2/service/glue v1.28.1
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.9.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.7.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.9.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.13.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/shield v1.16.7
	github.com/aws/aws-sdk-go-v2/service/sso v1.11.12 // indirect
	github.com/aws/aws-sdk-go-v2/service/transfer v1.21.4
	github.com/aws/aws-sdk-go-v2/service/wafregional v1.12.8
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/gofrs/uuid v4.3.0+incompatible // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/iancoleman/strcase v0.2.0
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pkg/errors v0.9.1
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/mod v0.6.0-dev.0.20220419223038-86c51ed26bb4 // indirect
	golang.org/x/net v0.0.0-20220728211354-c7608f3a8462 // indirect
	golang.org/x/sys v0.0.0-20220731174439-a90be440212d // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/tools v0.1.12 // indirect
	google.golang.org/genproto v0.0.0-20220801145646-83ce21fca29f // indirect
	google.golang.org/grpc v1.49.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
