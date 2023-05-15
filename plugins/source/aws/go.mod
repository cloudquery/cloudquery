module github.com/cloudquery/cloudquery/plugins/source/aws

go 1.19

require (
	github.com/aws/aws-sdk-go-v2 v1.18.0
	github.com/aws/aws-sdk-go-v2/config v1.18.25
	github.com/aws/aws-sdk-go-v2/credentials v1.13.24
	github.com/aws/aws-sdk-go-v2/service/accessanalyzer v1.19.12
	github.com/aws/aws-sdk-go-v2/service/account v1.8.0
	github.com/aws/aws-sdk-go-v2/service/acm v1.17.11
	github.com/aws/aws-sdk-go-v2/service/amp v1.16.11
	github.com/aws/aws-sdk-go-v2/service/amplify v1.13.10
	github.com/aws/aws-sdk-go-v2/service/apigateway v1.16.0
	github.com/aws/aws-sdk-go-v2/service/apigatewayv2 v1.13.0
	github.com/aws/aws-sdk-go-v2/service/applicationautoscaling v1.17.0
	github.com/aws/aws-sdk-go-v2/service/apprunner v1.16.0
	github.com/aws/aws-sdk-go-v2/service/appstream v1.19.0
	github.com/aws/aws-sdk-go-v2/service/appsync v1.18.0
	github.com/aws/aws-sdk-go-v2/service/athena v1.22.4
	github.com/aws/aws-sdk-go-v2/service/autoscaling v1.26.2
	github.com/aws/aws-sdk-go-v2/service/autoscalingplans v1.13.8
	github.com/aws/aws-sdk-go-v2/service/backup v1.19.2
	github.com/aws/aws-sdk-go-v2/service/batch v1.21.6
	github.com/aws/aws-sdk-go-v2/service/cloudformation v1.25.1
	github.com/aws/aws-sdk-go-v2/service/cloudfront v1.24.1
	github.com/aws/aws-sdk-go-v2/service/cloudhsmv2 v1.14.6
	github.com/aws/aws-sdk-go-v2/service/cloudtrail v1.22.1
	github.com/aws/aws-sdk-go-v2/service/cloudwatch v1.25.7
	github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs v1.20.7
	github.com/aws/aws-sdk-go-v2/service/codebuild v1.20.7
	github.com/aws/aws-sdk-go-v2/service/codepipeline v1.14.6
	github.com/aws/aws-sdk-go-v2/service/cognitoidentity v1.15.6
	github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider v1.22.6
	github.com/aws/aws-sdk-go-v2/service/computeoptimizer v1.21.5
	github.com/aws/aws-sdk-go-v2/service/configservice v1.29.6
	github.com/aws/aws-sdk-go-v2/service/costexplorer v1.25.8
	github.com/aws/aws-sdk-go-v2/service/databasemigrationservice v1.23.5
	github.com/aws/aws-sdk-go-v2/service/dax v1.12.6
	github.com/aws/aws-sdk-go-v2/service/directconnect v1.18.8
	github.com/aws/aws-sdk-go-v2/service/docdb v1.20.6
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.18.6
	github.com/aws/aws-sdk-go-v2/service/dynamodbstreams v1.14.7
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.80.0
	github.com/aws/aws-sdk-go-v2/service/ecr v1.18.7
	github.com/aws/aws-sdk-go-v2/service/ecrpublic v1.15.6
	github.com/aws/aws-sdk-go-v2/service/ecs v1.23.2
	github.com/aws/aws-sdk-go-v2/service/efs v1.19.9
	github.com/aws/aws-sdk-go-v2/service/eks v1.27.8
	github.com/aws/aws-sdk-go-v2/service/elasticache v1.26.6
	github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk v1.15.6
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing v1.15.6
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 v1.19.7
	github.com/aws/aws-sdk-go-v2/service/elasticsearchservice v1.18.7
	github.com/aws/aws-sdk-go-v2/service/elastictranscoder v1.14.6
	github.com/aws/aws-sdk-go-v2/service/emr v1.22.2
	github.com/aws/aws-sdk-go-v2/service/eventbridge v1.17.0
	github.com/aws/aws-sdk-go-v2/service/firehose v1.16.8
	github.com/aws/aws-sdk-go-v2/service/frauddetector v1.21.0
	github.com/aws/aws-sdk-go-v2/service/fsx v1.28.7
	github.com/aws/aws-sdk-go-v2/service/glacier v1.14.7
	github.com/aws/aws-sdk-go-v2/service/glue v1.40.0
	github.com/aws/aws-sdk-go-v2/service/guardduty v1.17.6
	github.com/aws/aws-sdk-go-v2/service/iam v1.19.8
	github.com/aws/aws-sdk-go-v2/service/identitystore v1.16.6
	github.com/aws/aws-sdk-go-v2/service/inspector v1.13.6
	github.com/aws/aws-sdk-go-v2/service/inspector2 v1.11.7
	github.com/aws/aws-sdk-go-v2/service/iot v1.33.0
	github.com/aws/aws-sdk-go-v2/service/kafka v1.19.7
	github.com/aws/aws-sdk-go-v2/service/kinesis v1.17.8
	github.com/aws/aws-sdk-go-v2/service/kms v1.20.8
	github.com/aws/aws-sdk-go-v2/service/lambda v1.29.0
	github.com/aws/aws-sdk-go-v2/service/lightsail v1.25.0
	github.com/aws/aws-sdk-go-v2/service/mq v1.14.0
	github.com/aws/aws-sdk-go-v2/service/mwaa v1.14.0
	github.com/aws/aws-sdk-go-v2/service/neptune v1.19.0
	github.com/aws/aws-sdk-go-v2/service/networkfirewall v1.26.3
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
	github.com/aws/aws-sdk-go-v2/service/s3 v1.32.0
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
	github.com/aws/aws-sdk-go-v2/service/shield v1.18.0
	github.com/aws/aws-sdk-go-v2/service/sns v1.19.0
	github.com/aws/aws-sdk-go-v2/service/sqs v1.20.0
	github.com/aws/aws-sdk-go-v2/service/ssm v1.35.0
	github.com/aws/aws-sdk-go-v2/service/ssoadmin v1.16.0
	github.com/aws/aws-sdk-go-v2/service/sts v1.19.0
	github.com/aws/aws-sdk-go-v2/service/support v1.14.1
	github.com/aws/aws-sdk-go-v2/service/timestreamwrite v1.15.0
	github.com/aws/aws-sdk-go-v2/service/transfer v1.28.0
	github.com/aws/aws-sdk-go-v2/service/waf v1.12.0
	github.com/aws/aws-sdk-go-v2/service/wafregional v1.13.1
	github.com/aws/aws-sdk-go-v2/service/wafv2 v1.24.2
	github.com/aws/aws-sdk-go-v2/service/workspaces v1.28.0
	github.com/aws/aws-sdk-go-v2/service/xray v1.16.0
	github.com/aws/smithy-go v1.13.5
	github.com/basgys/goxml2json v1.1.0
	github.com/cloudquery/codegen v0.2.1
	github.com/cloudquery/plugin-pb-go v1.0.5
	github.com/cloudquery/plugin-sdk/v2 v2.7.0
	github.com/gocarina/gocsv v0.0.0-20230325173030-9a18a846a479
	github.com/golang/mock v1.6.0
	github.com/google/go-cmp v0.5.9
	github.com/mitchellh/hashstructure/v2 v2.0.2
	github.com/mitchellh/mapstructure v1.5.0
	github.com/pkg/errors v0.9.1
	github.com/rs/zerolog v1.29.0
	github.com/stretchr/testify v1.8.2
	github.com/thoas/go-funk v0.9.3
	golang.org/x/exp v0.0.0-20230425010034-47ecfdc1ba53
	golang.org/x/sync v0.1.0
)

// TODO: remove once all updates are merged
replace github.com/apache/arrow/go/v13 => github.com/cloudquery/arrow/go/v13 v13.0.0-20230509053643-898a79b1d3c8

require (
	github.com/andybalholm/brotli v1.0.5 // indirect
	github.com/apache/arrow/go/v13 v13.0.0-20230509040948-de6c3cd2b604 // indirect
	github.com/apache/thrift v0.16.0 // indirect
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.4.10 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.13.3 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.33 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.27 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.3.34 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.0.25 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.9.11 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.1.27 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.7.23 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.9.27 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.14.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.12.10 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.14.10 // indirect
	github.com/bitly/go-simplejson v0.5.0 // indirect
	github.com/cloudquery/plugin-sdk v1.24.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/getsentry/sentry-go v0.20.0 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/goccy/go-json v0.9.11 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/flatbuffers v2.0.8+incompatible // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware/providers/zerolog/v2 v2.0.0-rc.3 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.0.0-rc.3 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/klauspost/asmfmt v1.3.2 // indirect
	github.com/klauspost/compress v1.16.0 // indirect
	github.com/klauspost/cpuid/v2 v2.0.9 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.18 // indirect
	github.com/minio/asm2plan9s v0.0.0-20200509001527-cdd76441f9d8 // indirect
	github.com/minio/c2goasm v0.0.0-20190812172519-36a3d3bbc4f3 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/cobra v1.6.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/zeebo/xxh3 v1.0.2 // indirect
	golang.org/x/mod v0.8.0 // indirect
	golang.org/x/net v0.9.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	golang.org/x/tools v0.6.0 // indirect
	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2 // indirect
	google.golang.org/genproto v0.0.0-20230410155749-daa745c078e1 // indirect
	google.golang.org/grpc v1.54.0 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
