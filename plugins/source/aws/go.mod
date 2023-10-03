module github.com/cloudquery/cloudquery/plugins/source/aws

go 1.21.1

require (
	github.com/apache/arrow/go/v14 v14.0.0-20230929201650-00efb06dc0de
	github.com/aws/aws-sdk-go-v2 v1.20.1
	github.com/aws/aws-sdk-go-v2/config v1.18.33
	github.com/aws/aws-sdk-go-v2/credentials v1.13.32
	github.com/aws/aws-sdk-go-v2/service/accessanalyzer v1.20.2
	github.com/aws/aws-sdk-go-v2/service/account v1.11.2
	github.com/aws/aws-sdk-go-v2/service/acm v1.18.2
	github.com/aws/aws-sdk-go-v2/service/acmpca v1.22.3
	github.com/aws/aws-sdk-go-v2/service/amp v1.17.2
	github.com/aws/aws-sdk-go-v2/service/amplify v1.14.2
	github.com/aws/aws-sdk-go-v2/service/apigateway v1.17.2
	github.com/aws/aws-sdk-go-v2/service/apigatewayv2 v1.14.2
	github.com/aws/aws-sdk-go-v2/service/appconfig v1.18.2
	github.com/aws/aws-sdk-go-v2/service/appflow v1.33.2
	github.com/aws/aws-sdk-go-v2/service/applicationautoscaling v1.22.2
	github.com/aws/aws-sdk-go-v2/service/appmesh v1.18.2
	github.com/aws/aws-sdk-go-v2/service/apprunner v1.18.2
	github.com/aws/aws-sdk-go-v2/service/appstream v1.22.2
	github.com/aws/aws-sdk-go-v2/service/appsync v1.22.2
	github.com/aws/aws-sdk-go-v2/service/athena v1.31.2
	github.com/aws/aws-sdk-go-v2/service/auditmanager v1.26.2
	github.com/aws/aws-sdk-go-v2/service/autoscaling v1.30.3
	github.com/aws/aws-sdk-go-v2/service/autoscalingplans v1.14.2
	github.com/aws/aws-sdk-go-v2/service/backup v1.24.0
	github.com/aws/aws-sdk-go-v2/service/batch v1.26.1
	github.com/aws/aws-sdk-go-v2/service/cloudformation v1.34.2
	github.com/aws/aws-sdk-go-v2/service/cloudfront v1.28.2
	github.com/aws/aws-sdk-go-v2/service/cloudhsmv2 v1.15.2
	github.com/aws/aws-sdk-go-v2/service/cloudtrail v1.28.3
	github.com/aws/aws-sdk-go-v2/service/cloudwatch v1.27.2
	github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs v1.23.2
	github.com/aws/aws-sdk-go-v2/service/codeartifact v1.19.2
	github.com/aws/aws-sdk-go-v2/service/codebuild v1.21.2
	github.com/aws/aws-sdk-go-v2/service/codecommit v1.15.2
	github.com/aws/aws-sdk-go-v2/service/codepipeline v1.16.2
	github.com/aws/aws-sdk-go-v2/service/cognitoidentity v1.16.2
	github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider v1.25.1
	github.com/aws/aws-sdk-go-v2/service/computeoptimizer v1.25.2
	github.com/aws/aws-sdk-go-v2/service/configservice v1.36.0
	github.com/aws/aws-sdk-go-v2/service/costexplorer v1.27.2
	github.com/aws/aws-sdk-go-v2/service/databasemigrationservice v1.30.1
	github.com/aws/aws-sdk-go-v2/service/dax v1.13.2
	github.com/aws/aws-sdk-go-v2/service/detective v1.21.0
	github.com/aws/aws-sdk-go-v2/service/directconnect v1.19.2
	github.com/aws/aws-sdk-go-v2/service/docdb v1.23.2
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.21.2
	github.com/aws/aws-sdk-go-v2/service/dynamodbstreams v1.15.2
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.111.0
	github.com/aws/aws-sdk-go-v2/service/ecr v1.19.2
	github.com/aws/aws-sdk-go-v2/service/ecrpublic v1.17.2
	github.com/aws/aws-sdk-go-v2/service/ecs v1.29.3
	github.com/aws/aws-sdk-go-v2/service/efs v1.21.3
	github.com/aws/aws-sdk-go-v2/service/eks v1.29.2
	github.com/aws/aws-sdk-go-v2/service/elasticache v1.29.0
	github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk v1.16.2
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing v1.16.2
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 v1.21.0
	github.com/aws/aws-sdk-go-v2/service/elasticsearchservice v1.20.2
	github.com/aws/aws-sdk-go-v2/service/elastictranscoder v1.15.2
	github.com/aws/aws-sdk-go-v2/service/emr v1.28.2
	github.com/aws/aws-sdk-go-v2/service/eventbridge v1.20.2
	github.com/aws/aws-sdk-go-v2/service/firehose v1.17.2
	github.com/aws/aws-sdk-go-v2/service/frauddetector v1.26.2
	github.com/aws/aws-sdk-go-v2/service/fsx v1.32.0
	github.com/aws/aws-sdk-go-v2/service/glacier v1.15.2
	github.com/aws/aws-sdk-go-v2/service/glue v1.60.1
	github.com/aws/aws-sdk-go-v2/service/guardduty v1.26.0
	github.com/aws/aws-sdk-go-v2/service/iam v1.22.2
	github.com/aws/aws-sdk-go-v2/service/identitystore v1.17.3
	github.com/aws/aws-sdk-go-v2/service/inspector v1.14.2
	github.com/aws/aws-sdk-go-v2/service/inspector2 v1.16.2
	github.com/aws/aws-sdk-go-v2/service/iot v1.39.2
	github.com/aws/aws-sdk-go-v2/service/kafka v1.22.2
	github.com/aws/aws-sdk-go-v2/service/kinesis v1.18.2
	github.com/aws/aws-sdk-go-v2/service/kms v1.24.2
	github.com/aws/aws-sdk-go-v2/service/lambda v1.39.2
	github.com/aws/aws-sdk-go-v2/service/lightsail v1.28.2
	github.com/aws/aws-sdk-go-v2/service/mq v1.16.2
	github.com/aws/aws-sdk-go-v2/service/mwaa v1.17.2
	github.com/aws/aws-sdk-go-v2/service/neptune v1.21.2
	github.com/aws/aws-sdk-go-v2/service/networkfirewall v1.29.2
	github.com/aws/aws-sdk-go-v2/service/networkmanager v1.18.2
	github.com/aws/aws-sdk-go-v2/service/organizations v1.20.2
	github.com/aws/aws-sdk-go-v2/service/qldb v1.16.2
	github.com/aws/aws-sdk-go-v2/service/quicksight v1.41.0
	github.com/aws/aws-sdk-go-v2/service/ram v1.20.2
	github.com/aws/aws-sdk-go-v2/service/rds v1.50.1
	github.com/aws/aws-sdk-go-v2/service/redshift v1.29.2
	github.com/aws/aws-sdk-go-v2/service/resiliencehub v1.13.1
	github.com/aws/aws-sdk-go-v2/service/resourcegroups v1.15.2
	github.com/aws/aws-sdk-go-v2/service/route53 v1.29.2
	github.com/aws/aws-sdk-go-v2/service/route53domains v1.16.2
	github.com/aws/aws-sdk-go-v2/service/route53recoverycontrolconfig v1.12.2
	github.com/aws/aws-sdk-go-v2/service/route53recoveryreadiness v1.10.2
	github.com/aws/aws-sdk-go-v2/service/route53resolver v1.19.3
	github.com/aws/aws-sdk-go-v2/service/s3 v1.38.2
	github.com/aws/aws-sdk-go-v2/service/s3control v1.32.2
	github.com/aws/aws-sdk-go-v2/service/sagemaker v1.101.0
	github.com/aws/aws-sdk-go-v2/service/savingsplans v1.13.2
	github.com/aws/aws-sdk-go-v2/service/scheduler v1.2.2
	github.com/aws/aws-sdk-go-v2/service/secretsmanager v1.21.0
	github.com/aws/aws-sdk-go-v2/service/securityhub v1.35.3
	github.com/aws/aws-sdk-go-v2/service/servicecatalog v1.21.0
	github.com/aws/aws-sdk-go-v2/service/servicecatalogappregistry v1.18.3
	github.com/aws/aws-sdk-go-v2/service/servicediscovery v1.22.2
	github.com/aws/aws-sdk-go-v2/service/servicequotas v1.15.2
	github.com/aws/aws-sdk-go-v2/service/ses v1.16.3
	github.com/aws/aws-sdk-go-v2/service/sesv2 v1.19.2
	github.com/aws/aws-sdk-go-v2/service/sfn v1.19.2
	github.com/aws/aws-sdk-go-v2/service/shield v1.19.2
	github.com/aws/aws-sdk-go-v2/service/signer v1.16.2
	github.com/aws/aws-sdk-go-v2/service/sns v1.21.2
	github.com/aws/aws-sdk-go-v2/service/sqs v1.24.2
	github.com/aws/aws-sdk-go-v2/service/ssm v1.37.2
	github.com/aws/aws-sdk-go-v2/service/ssoadmin v1.17.2
	github.com/aws/aws-sdk-go-v2/service/sts v1.21.2
	github.com/aws/aws-sdk-go-v2/service/support v1.16.2
	github.com/aws/aws-sdk-go-v2/service/timestreamwrite v1.18.2
	github.com/aws/aws-sdk-go-v2/service/transfer v1.33.3
	github.com/aws/aws-sdk-go-v2/service/waf v1.13.2
	github.com/aws/aws-sdk-go-v2/service/wafregional v1.14.2
	github.com/aws/aws-sdk-go-v2/service/wafv2 v1.37.2
	github.com/aws/aws-sdk-go-v2/service/wellarchitected v1.22.2
	github.com/aws/aws-sdk-go-v2/service/workspaces v1.29.2
	github.com/aws/aws-sdk-go-v2/service/xray v1.17.2
	github.com/aws/smithy-go v1.14.1
	github.com/basgys/goxml2json v1.1.0
	github.com/cloudquery/codegen v0.3.4
	github.com/cloudquery/plugin-sdk/v4 v4.12.0
	github.com/cockroachdb/cockroachdb-parser v0.0.0-20230705064001-302c9ad52e1a
	github.com/gertd/go-pluralize v0.2.1
	github.com/gocarina/gocsv v0.0.0-20230616125104-99d496ca653d
	github.com/golang/mock v1.6.0
	github.com/google/go-cmp v0.5.9
	github.com/mitchellh/hashstructure/v2 v2.0.2
	github.com/mitchellh/mapstructure v1.5.0
	github.com/mjibson/sqlfmt v0.5.0
	github.com/pkg/errors v0.9.1
	github.com/rs/zerolog v1.29.1
	github.com/stretchr/testify v1.8.4
	github.com/thoas/go-funk v0.9.3
	golang.org/x/exp v0.0.0-20230905200255-921286631fa9
	golang.org/x/sync v0.3.0
	google.golang.org/grpc v1.57.0
)

// TODO: remove once all updates are merged
replace github.com/apache/arrow/go/v14 => github.com/cloudquery/arrow/go/v14 v14.0.0-20231002001222-7ded38b478cd

require (
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.4.12 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.13.8 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.38 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.32 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.3.39 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.1.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.9.13 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.1.33 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.7.32 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.9.32 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.15.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.13.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.15.2 // indirect
	github.com/bitly/go-simplejson v0.5.0 // indirect
	github.com/cloudquery/plugin-sdk v1.45.0
	github.com/cloudquery/plugin-sdk/v2 v2.7.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/getsentry/sentry-go v0.24.1 // indirect
	github.com/ghodss/yaml v1.0.0
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/flatbuffers v23.5.26+incompatible // indirect
	github.com/google/uuid v1.3.1 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.0.1 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/klauspost/compress v1.16.7 // indirect
	github.com/klauspost/cpuid/v2 v2.2.5 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/pierrec/lz4/v4 v4.1.18 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/spf13/cobra v1.6.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/zeebo/xxh3 v1.0.2 // indirect
	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2 // indirect
)

require (
	github.com/apache/arrow/go/v13 v13.0.0-20230731205701-112f94971882 // indirect
	github.com/bahlo/generic-list-go v0.2.0 // indirect
	github.com/biogo/store v0.0.0-20201120204734-aad293a2328f // indirect
	github.com/blevesearch/snowballstem v0.9.0 // indirect
	github.com/buger/jsonparser v1.1.1 // indirect
	github.com/cenkalti/backoff/v4 v4.2.1 // indirect
	github.com/cloudquery/plugin-pb-go v1.11.0 // indirect
	github.com/cockroachdb/apd/v3 v3.1.0 // indirect
	github.com/cockroachdb/errors v1.9.0 // indirect
	github.com/cockroachdb/logtags v0.0.0-20230118201751-21c54148d20b // indirect
	github.com/cockroachdb/redact v1.1.3 // indirect
	github.com/dave/dst v0.27.2 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/geo v0.0.0-20230421003525-6adc56603217 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.7.0 // indirect
	github.com/invopop/jsonschema v0.11.0 // indirect
	github.com/jpillora/longestcommon v0.0.0-20161227235612-adb9d91ee629 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/lib/pq v1.10.6 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/petermattis/goid v0.0.0-20230904192822-1876fd5063bc // indirect
	github.com/pierrre/geohash v1.0.0 // indirect
	github.com/rogpeppe/go-internal v1.10.0 // indirect
	github.com/santhosh-tekuri/jsonschema/v5 v5.3.1 // indirect
	github.com/sasha-s/go-deadlock v0.3.1 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/twpayne/go-geom v1.4.2 // indirect
	github.com/twpayne/go-kml v1.5.2 // indirect
	github.com/wk8/go-ordered-map/v2 v2.1.8 // indirect
	go.opentelemetry.io/otel v1.16.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/internal/retry v1.16.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.16.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.16.0 // indirect
	go.opentelemetry.io/otel/metric v1.16.0 // indirect
	go.opentelemetry.io/otel/sdk v1.16.0 // indirect
	go.opentelemetry.io/otel/trace v1.16.0 // indirect
	go.opentelemetry.io/proto/otlp v0.19.0 // indirect
	golang.org/x/mod v0.12.0 // indirect
	golang.org/x/net v0.15.0 // indirect
	golang.org/x/sys v0.12.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	golang.org/x/tools v0.13.0 // indirect
	google.golang.org/genproto v0.0.0-20230920204549-e6e6cdab5c13 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20230920204549-e6e6cdab5c13 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230920204549-e6e6cdab5c13 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
