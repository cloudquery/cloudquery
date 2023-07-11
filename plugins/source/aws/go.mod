module github.com/cloudquery/cloudquery/plugins/source/aws

go 1.19

require (
	github.com/apache/arrow/go/v13 v13.0.0-20230630125530-5a06b2ec2a8e
	github.com/aws/aws-sdk-go-v2 v1.18.1
	github.com/aws/aws-sdk-go-v2/config v1.18.27
	github.com/aws/aws-sdk-go-v2/credentials v1.13.26
	github.com/aws/aws-sdk-go-v2/service/accessanalyzer v1.19.14
	github.com/aws/aws-sdk-go-v2/service/account v1.10.8
	github.com/aws/aws-sdk-go-v2/service/acm v1.17.13
	github.com/aws/aws-sdk-go-v2/service/acmpca v1.21.15
	github.com/aws/aws-sdk-go-v2/service/amp v1.16.13
	github.com/aws/aws-sdk-go-v2/service/amplify v1.13.12
	github.com/aws/aws-sdk-go-v2/service/apigateway v1.16.13
	github.com/aws/aws-sdk-go-v2/service/apigatewayv2 v1.13.13
	github.com/aws/aws-sdk-go-v2/service/applicationautoscaling v1.21.2
	github.com/aws/aws-sdk-go-v2/service/apprunner v1.17.11
	github.com/aws/aws-sdk-go-v2/service/appstream v1.20.13
	github.com/aws/aws-sdk-go-v2/service/appsync v1.21.2
	github.com/aws/aws-sdk-go-v2/service/athena v1.30.2
	github.com/aws/aws-sdk-go-v2/service/autoscaling v1.28.9
	github.com/aws/aws-sdk-go-v2/service/autoscalingplans v1.13.12
	github.com/aws/aws-sdk-go-v2/service/backup v1.22.2
	github.com/aws/aws-sdk-go-v2/service/batch v1.23.8
	github.com/aws/aws-sdk-go-v2/service/cloudformation v1.30.0
	github.com/aws/aws-sdk-go-v2/service/cloudfront v1.26.8
	github.com/aws/aws-sdk-go-v2/service/cloudhsmv2 v1.14.12
	github.com/aws/aws-sdk-go-v2/service/cloudtrail v1.27.1
	github.com/aws/aws-sdk-go-v2/service/cloudwatch v1.26.2
	github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs v1.21.2
	github.com/aws/aws-sdk-go-v2/service/codeartifact v1.18.6
	github.com/aws/aws-sdk-go-v2/service/codebuild v1.20.15
	github.com/aws/aws-sdk-go-v2/service/codecommit v1.14.14
	github.com/aws/aws-sdk-go-v2/service/codepipeline v1.15.2
	github.com/aws/aws-sdk-go-v2/service/cognitoidentity v1.15.13
	github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider v1.22.12
	github.com/aws/aws-sdk-go-v2/service/computeoptimizer v1.24.2
	github.com/aws/aws-sdk-go-v2/service/configservice v1.34.0
	github.com/aws/aws-sdk-go-v2/service/costexplorer v1.25.12
	github.com/aws/aws-sdk-go-v2/service/databasemigrationservice v1.25.7
	github.com/aws/aws-sdk-go-v2/service/dax v1.12.12
	github.com/aws/aws-sdk-go-v2/service/detective v1.19.2
	github.com/aws/aws-sdk-go-v2/service/directconnect v1.18.17
	github.com/aws/aws-sdk-go-v2/service/docdb v1.21.5
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.19.11
	github.com/aws/aws-sdk-go-v2/service/dynamodbstreams v1.14.14
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.102.0
	github.com/aws/aws-sdk-go-v2/service/ecr v1.18.13
	github.com/aws/aws-sdk-go-v2/service/ecrpublic v1.16.4
	github.com/aws/aws-sdk-go-v2/service/ecs v1.27.4
	github.com/aws/aws-sdk-go-v2/service/efs v1.20.3
	github.com/aws/aws-sdk-go-v2/service/eks v1.27.14
	github.com/aws/aws-sdk-go-v2/service/elasticache v1.27.2
	github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk v1.15.12
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing v1.15.12
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 v1.19.13
	github.com/aws/aws-sdk-go-v2/service/elasticsearchservice v1.19.2
	github.com/aws/aws-sdk-go-v2/service/elastictranscoder v1.14.12
	github.com/aws/aws-sdk-go-v2/service/emr v1.27.0
	github.com/aws/aws-sdk-go-v2/service/eventbridge v1.19.4
	github.com/aws/aws-sdk-go-v2/service/firehose v1.16.14
	github.com/aws/aws-sdk-go-v2/service/frauddetector v1.25.2
	github.com/aws/aws-sdk-go-v2/service/fsx v1.29.3
	github.com/aws/aws-sdk-go-v2/service/glacier v1.14.13
	github.com/aws/aws-sdk-go-v2/service/glue v1.52.0
	github.com/aws/aws-sdk-go-v2/service/guardduty v1.24.0
	github.com/aws/aws-sdk-go-v2/service/iam v1.21.0
	github.com/aws/aws-sdk-go-v2/service/identitystore v1.16.13
	github.com/aws/aws-sdk-go-v2/service/inspector v1.13.12
	github.com/aws/aws-sdk-go-v2/service/inspector2 v1.15.0
	github.com/aws/aws-sdk-go-v2/service/iot v1.38.2
	github.com/aws/aws-sdk-go-v2/service/kafka v1.20.5
	github.com/aws/aws-sdk-go-v2/service/kinesis v1.17.14
	github.com/aws/aws-sdk-go-v2/service/kms v1.22.2
	github.com/aws/aws-sdk-go-v2/service/lambda v1.36.0
	github.com/aws/aws-sdk-go-v2/service/lightsail v1.27.1
	github.com/aws/aws-sdk-go-v2/service/mq v1.15.0
	github.com/aws/aws-sdk-go-v2/service/mwaa v1.16.2
	github.com/aws/aws-sdk-go-v2/service/neptune v1.20.7
	github.com/aws/aws-sdk-go-v2/service/networkfirewall v1.28.3
	github.com/aws/aws-sdk-go-v2/service/networkmanager v1.17.14
	github.com/aws/aws-sdk-go-v2/service/organizations v1.19.8
	github.com/aws/aws-sdk-go-v2/service/qldb v1.15.13
	github.com/aws/aws-sdk-go-v2/service/quicksight v1.37.2
	github.com/aws/aws-sdk-go-v2/service/ram v1.18.4
	github.com/aws/aws-sdk-go-v2/service/rds v1.45.3
	github.com/aws/aws-sdk-go-v2/service/redshift v1.28.0
	github.com/aws/aws-sdk-go-v2/service/resiliencehub v1.11.3
	github.com/aws/aws-sdk-go-v2/service/resourcegroups v1.14.13
	github.com/aws/aws-sdk-go-v2/service/route53 v1.28.3
	github.com/aws/aws-sdk-go-v2/service/route53domains v1.15.0
	github.com/aws/aws-sdk-go-v2/service/route53resolver v1.17.13
	github.com/aws/aws-sdk-go-v2/service/s3 v1.35.0
	github.com/aws/aws-sdk-go-v2/service/s3control v1.31.7
	github.com/aws/aws-sdk-go-v2/service/sagemaker v1.86.0
	github.com/aws/aws-sdk-go-v2/service/savingsplans v1.12.12
	github.com/aws/aws-sdk-go-v2/service/scheduler v1.1.13
	github.com/aws/aws-sdk-go-v2/service/secretsmanager v1.19.10
	github.com/aws/aws-sdk-go-v2/service/securityhub v1.33.2
	github.com/aws/aws-sdk-go-v2/service/servicecatalog v1.19.2
	github.com/aws/aws-sdk-go-v2/service/servicecatalogappregistry v1.17.6
	github.com/aws/aws-sdk-go-v2/service/servicediscovery v1.21.7
	github.com/aws/aws-sdk-go-v2/service/servicequotas v1.14.14
	github.com/aws/aws-sdk-go-v2/service/ses v1.15.11
	github.com/aws/aws-sdk-go-v2/service/sesv2 v1.18.2
	github.com/aws/aws-sdk-go-v2/service/sfn v1.18.0
	github.com/aws/aws-sdk-go-v2/service/shield v1.18.12
	github.com/aws/aws-sdk-go-v2/service/signer v1.15.2
	github.com/aws/aws-sdk-go-v2/service/sns v1.20.13
	github.com/aws/aws-sdk-go-v2/service/sqs v1.23.2
	github.com/aws/aws-sdk-go-v2/service/ssm v1.36.6
	github.com/aws/aws-sdk-go-v2/service/ssoadmin v1.16.12
	github.com/aws/aws-sdk-go-v2/service/sts v1.19.2
	github.com/aws/aws-sdk-go-v2/service/support v1.15.2
	github.com/aws/aws-sdk-go-v2/service/timestreamwrite v1.17.2
	github.com/aws/aws-sdk-go-v2/service/transfer v1.30.0
	github.com/aws/aws-sdk-go-v2/service/waf v1.12.12
	github.com/aws/aws-sdk-go-v2/service/wafregional v1.13.15
	github.com/aws/aws-sdk-go-v2/service/wafv2 v1.35.1
	github.com/aws/aws-sdk-go-v2/service/wellarchitected v1.21.1
	github.com/aws/aws-sdk-go-v2/service/workspaces v1.28.15
	github.com/aws/aws-sdk-go-v2/service/xray v1.16.13
	github.com/aws/smithy-go v1.13.5
	github.com/basgys/goxml2json v1.1.0
	github.com/cloudquery/codegen v0.2.1
	github.com/cloudquery/plugin-sdk/v4 v4.8.1-rc1
	github.com/cockroachdb/cockroachdb-parser v0.0.0-20221207165326-ea0ac1a4778b
	github.com/gocarina/gocsv v0.0.0-20230513223533-9ddd7fd60602
	github.com/golang/mock v1.6.0
	github.com/google/go-cmp v0.5.9
	github.com/mitchellh/hashstructure/v2 v2.0.2
	github.com/mitchellh/mapstructure v1.5.0
	github.com/mjibson/sqlfmt v0.5.0
	github.com/pkg/errors v0.9.1
	github.com/rs/zerolog v1.29.1
	github.com/stretchr/testify v1.8.4
	github.com/thoas/go-funk v0.9.3
	golang.org/x/exp v0.0.0-20230626212559-97b1e661b5df
	golang.org/x/sync v0.1.0
	google.golang.org/grpc v1.55.0
)

// TODO: remove once all updates are merged
replace github.com/apache/arrow/go/v13 => github.com/cloudquery/arrow/go/v13 v13.0.0-20230710001530-a2a76ebbb85f

require (
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.4.10 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.13.4 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.34 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.28 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.3.35 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.0.26 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.9.11 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.1.29 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.7.28 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.9.28 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.14.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.12.12 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.14.12 // indirect
	github.com/bitly/go-simplejson v0.5.0 // indirect
	github.com/cloudquery/plugin-sdk v1.45.0
	github.com/cloudquery/plugin-sdk/v2 v2.7.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/getsentry/sentry-go v0.20.0 // indirect
	github.com/ghodss/yaml v1.0.0
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/flatbuffers v23.1.21+incompatible // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware/providers/zerolog/v2 v2.0.0-rc.3 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.0.0-rc.3 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/klauspost/compress v1.16.6 // indirect
	github.com/klauspost/cpuid/v2 v2.2.5 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/pierrec/lz4/v4 v4.1.17 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/spf13/cobra v1.6.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/zeebo/xxh3 v1.0.2 // indirect
	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2 // indirect
)

require (
	github.com/biogo/store v0.0.0-20201120204734-aad293a2328f // indirect
	github.com/cloudquery/plugin-pb-go v1.6.0 // indirect
	github.com/cockroachdb/apd/v3 v3.1.0 // indirect
	github.com/cockroachdb/errors v1.9.0 // indirect
	github.com/cockroachdb/logtags v0.0.0-20230118201751-21c54148d20b // indirect
	github.com/cockroachdb/redact v1.1.3 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/geo v0.0.0-20230421003525-6adc56603217 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/kr/pretty v0.3.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/lib/pq v1.10.6 // indirect
	github.com/petermattis/goid v0.0.0-20211229010228-4d14c490ee36 // indirect
	github.com/pierrre/geohash v1.0.0 // indirect
	github.com/rogpeppe/go-internal v1.9.0 // indirect
	github.com/sasha-s/go-deadlock v0.3.1 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/twpayne/go-geom v1.4.2 // indirect
	github.com/twpayne/go-kml v1.5.2 // indirect
	golang.org/x/mod v0.11.0 // indirect
	golang.org/x/net v0.9.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	golang.org/x/tools v0.6.0 // indirect
	google.golang.org/genproto v0.0.0-20230525234025-438c736192d0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20230525234020-1aefcd67740a // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230629202037-9506855d4529 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
