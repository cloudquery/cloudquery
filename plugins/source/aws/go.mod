module github.com/cloudquery/cloudquery/plugins/source/aws

go 1.21.1

require (
	github.com/apache/arrow/go/v14 v14.0.0-20231031200323-c49e24273160
	github.com/aws/aws-sdk-go-v2 v1.23.4
	github.com/aws/aws-sdk-go-v2/config v1.25.5
	github.com/aws/aws-sdk-go-v2/credentials v1.16.4
	github.com/aws/aws-sdk-go-v2/service/accessanalyzer v1.23.3
	github.com/aws/aws-sdk-go-v2/service/account v1.13.3
	github.com/aws/aws-sdk-go-v2/service/acm v1.21.3
	github.com/aws/aws-sdk-go-v2/service/acmpca v1.24.3
	github.com/aws/aws-sdk-go-v2/service/amp v1.19.3
	github.com/aws/aws-sdk-go-v2/service/amplify v1.17.3
	github.com/aws/aws-sdk-go-v2/service/apigateway v1.20.3
	github.com/aws/aws-sdk-go-v2/service/apigatewayv2 v1.17.3
	github.com/aws/aws-sdk-go-v2/service/appconfig v1.25.3
	github.com/aws/aws-sdk-go-v2/service/appflow v1.38.3
	github.com/aws/aws-sdk-go-v2/service/applicationautoscaling v1.24.3
	github.com/aws/aws-sdk-go-v2/service/appmesh v1.22.1
	github.com/aws/aws-sdk-go-v2/service/apprunner v1.24.4
	github.com/aws/aws-sdk-go-v2/service/appstream v1.28.3
	github.com/aws/aws-sdk-go-v2/service/appsync v1.24.3
	github.com/aws/aws-sdk-go-v2/service/athena v1.35.1
	github.com/aws/aws-sdk-go-v2/service/auditmanager v1.29.3
	github.com/aws/aws-sdk-go-v2/service/autoscaling v1.35.2
	github.com/aws/aws-sdk-go-v2/service/autoscalingplans v1.17.3
	github.com/aws/aws-sdk-go-v2/service/backup v1.28.2
	github.com/aws/aws-sdk-go-v2/service/batch v1.29.3
	github.com/aws/aws-sdk-go-v2/service/cloudformation v1.40.1
	github.com/aws/aws-sdk-go-v2/service/cloudfront v1.31.0
	github.com/aws/aws-sdk-go-v2/service/cloudhsmv2 v1.18.3
	github.com/aws/aws-sdk-go-v2/service/cloudtrail v1.33.1
	github.com/aws/aws-sdk-go-v2/service/cloudwatch v1.30.3
	github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs v1.27.2
	github.com/aws/aws-sdk-go-v2/service/codeartifact v1.22.3
	github.com/aws/aws-sdk-go-v2/service/codebuild v1.25.3
	github.com/aws/aws-sdk-go-v2/service/codecommit v1.18.3
	github.com/aws/aws-sdk-go-v2/service/codepipeline v1.21.1
	github.com/aws/aws-sdk-go-v2/service/cognitoidentity v1.20.3
	github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider v1.30.3
	github.com/aws/aws-sdk-go-v2/service/computeoptimizer v1.29.4
	github.com/aws/aws-sdk-go-v2/service/configservice v1.41.3
	github.com/aws/aws-sdk-go-v2/service/costexplorer v1.32.3
	github.com/aws/aws-sdk-go-v2/service/databasemigrationservice v1.34.2
	github.com/aws/aws-sdk-go-v2/service/dax v1.16.3
	github.com/aws/aws-sdk-go-v2/service/detective v1.23.3
	github.com/aws/aws-sdk-go-v2/service/directconnect v1.21.3
	github.com/aws/aws-sdk-go-v2/service/docdb v1.28.0
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.26.2
	github.com/aws/aws-sdk-go-v2/service/dynamodbstreams v1.17.3
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.137.1
	github.com/aws/aws-sdk-go-v2/service/ecr v1.23.1
	github.com/aws/aws-sdk-go-v2/service/ecrpublic v1.20.3
	github.com/aws/aws-sdk-go-v2/service/ecs v1.33.2
	github.com/aws/aws-sdk-go-v2/service/efs v1.23.3
	github.com/aws/aws-sdk-go-v2/service/eks v1.33.2
	github.com/aws/aws-sdk-go-v2/service/elasticache v1.32.3
	github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk v1.19.3
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing v1.20.3
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 v1.24.3
	github.com/aws/aws-sdk-go-v2/service/elasticsearchservice v1.23.3
	github.com/aws/aws-sdk-go-v2/service/elastictranscoder v1.18.3
	github.com/aws/aws-sdk-go-v2/service/emr v1.34.1
	github.com/aws/aws-sdk-go-v2/service/eventbridge v1.25.1
	github.com/aws/aws-sdk-go-v2/service/firehose v1.21.3
	github.com/aws/aws-sdk-go-v2/service/frauddetector v1.28.3
	github.com/aws/aws-sdk-go-v2/service/fsx v1.37.1
	github.com/aws/aws-sdk-go-v2/service/glacier v1.18.3
	github.com/aws/aws-sdk-go-v2/service/glue v1.72.0
	github.com/aws/aws-sdk-go-v2/service/guardduty v1.33.3
	github.com/aws/aws-sdk-go-v2/service/iam v1.27.3
	github.com/aws/aws-sdk-go-v2/service/identitystore v1.20.3
	github.com/aws/aws-sdk-go-v2/service/inspector v1.18.3
	github.com/aws/aws-sdk-go-v2/service/inspector2 v1.19.3
	github.com/aws/aws-sdk-go-v2/service/iot v1.45.1
	github.com/aws/aws-sdk-go-v2/service/kafka v1.27.1
	github.com/aws/aws-sdk-go-v2/service/kinesis v1.23.0
	github.com/aws/aws-sdk-go-v2/service/kms v1.26.3
	github.com/aws/aws-sdk-go-v2/service/lambda v1.48.1
	github.com/aws/aws-sdk-go-v2/service/lightsail v1.31.3
	github.com/aws/aws-sdk-go-v2/service/mq v1.19.3
	github.com/aws/aws-sdk-go-v2/service/mwaa v1.21.1
	github.com/aws/aws-sdk-go-v2/service/neptune v1.26.3
	github.com/aws/aws-sdk-go-v2/service/networkfirewall v1.35.3
	github.com/aws/aws-sdk-go-v2/service/networkmanager v1.22.3
	github.com/aws/aws-sdk-go-v2/service/organizations v1.22.3
	github.com/aws/aws-sdk-go-v2/service/qldb v1.18.3
	github.com/aws/aws-sdk-go-v2/service/quicksight v1.51.1
	github.com/aws/aws-sdk-go-v2/service/ram v1.22.3
	github.com/aws/aws-sdk-go-v2/service/rds v1.63.1
	github.com/aws/aws-sdk-go-v2/service/redshift v1.37.1
	github.com/aws/aws-sdk-go-v2/service/resiliencehub v1.17.3
	github.com/aws/aws-sdk-go-v2/service/resourcegroups v1.18.3
	github.com/aws/aws-sdk-go-v2/service/route53 v1.34.3
	github.com/aws/aws-sdk-go-v2/service/route53domains v1.19.3
	github.com/aws/aws-sdk-go-v2/service/route53recoverycontrolconfig v1.17.3
	github.com/aws/aws-sdk-go-v2/service/route53recoveryreadiness v1.14.3
	github.com/aws/aws-sdk-go-v2/service/route53resolver v1.22.3
	github.com/aws/aws-sdk-go-v2/service/s3 v1.44.0
	github.com/aws/aws-sdk-go-v2/service/s3control v1.38.0
	github.com/aws/aws-sdk-go-v2/service/sagemaker v1.120.0
	github.com/aws/aws-sdk-go-v2/service/savingsplans v1.15.3
	github.com/aws/aws-sdk-go-v2/service/scheduler v1.5.3
	github.com/aws/aws-sdk-go-v2/service/secretsmanager v1.23.3
	github.com/aws/aws-sdk-go-v2/service/securityhub v1.40.3
	github.com/aws/aws-sdk-go-v2/service/servicecatalog v1.24.3
	github.com/aws/aws-sdk-go-v2/service/servicecatalogappregistry v1.23.3
	github.com/aws/aws-sdk-go-v2/service/servicediscovery v1.26.3
	github.com/aws/aws-sdk-go-v2/service/servicequotas v1.18.3
	github.com/aws/aws-sdk-go-v2/service/ses v1.18.3
	github.com/aws/aws-sdk-go-v2/service/sesv2 v1.23.3
	github.com/aws/aws-sdk-go-v2/service/sfn v1.22.2
	github.com/aws/aws-sdk-go-v2/service/shield v1.22.3
	github.com/aws/aws-sdk-go-v2/service/signer v1.18.4
	github.com/aws/aws-sdk-go-v2/service/sns v1.25.3
	github.com/aws/aws-sdk-go-v2/service/sqs v1.28.2
	github.com/aws/aws-sdk-go-v2/service/ssm v1.43.1
	github.com/aws/aws-sdk-go-v2/service/ssoadmin v1.22.1
	github.com/aws/aws-sdk-go-v2/service/sts v1.25.4
	github.com/aws/aws-sdk-go-v2/service/support v1.18.3
	github.com/aws/aws-sdk-go-v2/service/timestreamwrite v1.23.2
	github.com/aws/aws-sdk-go-v2/service/transfer v1.38.1
	github.com/aws/aws-sdk-go-v2/service/waf v1.17.3
	github.com/aws/aws-sdk-go-v2/service/wafregional v1.18.3
	github.com/aws/aws-sdk-go-v2/service/wafv2 v1.42.3
	github.com/aws/aws-sdk-go-v2/service/wellarchitected v1.26.3
	github.com/aws/aws-sdk-go-v2/service/workspaces v1.33.4
	github.com/aws/aws-sdk-go-v2/service/xray v1.22.3
	github.com/aws/smithy-go v1.18.1
	github.com/basgys/goxml2json v1.1.0
	github.com/cloudquery/codegen v0.3.12
	github.com/cloudquery/plugin-sdk/v4 v4.19.0
	github.com/gertd/go-pluralize v0.2.1
	github.com/ghodss/yaml v1.0.0
	github.com/gocarina/gocsv v0.0.0-20231116093920-b87c2d0e983a
	github.com/golang/mock v1.6.0
	github.com/google/go-cmp v0.6.0
	github.com/google/uuid v1.4.0
	github.com/invopop/jsonschema v0.11.0
	github.com/mitchellh/hashstructure/v2 v2.0.2
	github.com/mitchellh/mapstructure v1.5.0
	github.com/pkg/errors v0.9.1
	github.com/rs/zerolog v1.29.1
	github.com/stretchr/testify v1.8.4
	github.com/thoas/go-funk v0.9.3
	github.com/wk8/go-ordered-map/v2 v2.1.8
	golang.org/x/sync v0.4.0
	google.golang.org/grpc v1.59.0
)

require (
	github.com/BurntSushi/toml v1.3.2 // indirect
	github.com/CloudyKit/fastprinter v0.0.0-20200109182630-33d98a066a53 // indirect
	github.com/CloudyKit/jet/v6 v6.2.0 // indirect
	github.com/Joker/jade v1.1.3 // indirect
	github.com/Shopify/goreferrer v0.0.0-20220729165902-8cddb4f5de06 // indirect
	github.com/andybalholm/brotli v1.0.6 // indirect
	github.com/apache/arrow/go/v13 v13.0.0-20230731205701-112f94971882 // indirect
	github.com/apapsch/go-jsonmerge/v2 v2.0.0 // indirect
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.5.1 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.14.5 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.2.7 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.5.7 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.7.1 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.2.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.10.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.2.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.8.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.10.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.16.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.17.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.20.1 // indirect
	github.com/aymerick/douceur v0.2.0 // indirect
	github.com/bahlo/generic-list-go v0.2.0 // indirect
	github.com/bitly/go-simplejson v0.5.0 // indirect
	github.com/buger/jsonparser v1.1.1 // indirect
	github.com/bytedance/sonic v1.10.2 // indirect
	github.com/cenkalti/backoff/v4 v4.2.1 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20230717121745-296ad89f973d // indirect
	github.com/chenzhuoyu/iasm v0.9.0 // indirect
	github.com/cloudquery/cloudquery-api-go v1.6.0 // indirect
	github.com/cloudquery/plugin-pb-go v1.14.2 // indirect
	github.com/cloudquery/plugin-sdk/v2 v2.7.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/deepmap/oapi-codegen v1.15.0 // indirect
	github.com/fatih/structs v1.1.0 // indirect
	github.com/flosch/pongo2/v4 v4.0.2 // indirect
	github.com/gabriel-vasile/mimetype v1.4.3 // indirect
	github.com/getsentry/sentry-go v0.24.1 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/gin-gonic/gin v1.9.1 // indirect
	github.com/go-logr/logr v1.3.0 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.14.0 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/gomarkdown/markdown v0.0.0-20231115200524-a660076da3fd // indirect
	github.com/google/flatbuffers v23.5.26+incompatible // indirect
	github.com/gorilla/css v1.0.0 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.0.1 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.16.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/iris-contrib/schema v0.0.6 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/jpillora/longestcommon v0.0.0-20161227235612-adb9d91ee629 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/kataras/blocks v0.0.8 // indirect
	github.com/kataras/golog v0.1.9 // indirect
	github.com/kataras/iris/v12 v12.2.7 // indirect
	github.com/kataras/pio v0.0.12 // indirect
	github.com/kataras/sitemap v0.0.6 // indirect
	github.com/kataras/tunnel v0.0.4 // indirect
	github.com/klauspost/compress v1.17.0 // indirect
	github.com/klauspost/cpuid/v2 v2.2.5 // indirect
	github.com/labstack/echo/v4 v4.11.1 // indirect
	github.com/labstack/gommon v0.4.0 // indirect
	github.com/leodido/go-urn v1.2.4 // indirect
	github.com/mailgun/raymond/v2 v2.0.48 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/microcosm-cc/bluemonday v1.0.26 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.0.8 // indirect
	github.com/pierrec/lz4/v4 v4.1.18 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/santhosh-tekuri/jsonschema/v5 v5.3.1 // indirect
	github.com/schollz/closestmatch v2.1.0+incompatible // indirect
	github.com/sergi/go-diff v1.2.0 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/spf13/cobra v1.6.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/tdewolff/minify/v2 v2.12.9 // indirect
	github.com/tdewolff/parse/v2 v2.6.8 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.11 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	github.com/vmihailenco/msgpack/v5 v5.3.5 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	github.com/xeipuuv/gojsonpointer v0.0.0-20190905194746-02993c407bfb // indirect
	github.com/yosssi/ace v0.0.5 // indirect
	github.com/zeebo/xxh3 v1.0.2 // indirect
	go.opentelemetry.io/otel v1.20.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.20.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.20.0 // indirect
	go.opentelemetry.io/otel/metric v1.20.0 // indirect
	go.opentelemetry.io/otel/sdk v1.20.0 // indirect
	go.opentelemetry.io/otel/trace v1.20.0 // indirect
	go.opentelemetry.io/proto/otlp v1.0.0 // indirect
	golang.org/x/arch v0.3.0 // indirect
	golang.org/x/crypto v0.14.0 // indirect
	golang.org/x/exp v0.0.0-20231006140011-7918f672742d // indirect
	golang.org/x/mod v0.13.0 // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/sys v0.14.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	golang.org/x/time v0.3.0 // indirect
	golang.org/x/tools v0.14.0 // indirect
	golang.org/x/xerrors v0.0.0-20231012003039-104605ab7028 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20231030173426-d783a09b4405 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231030173426-d783a09b4405 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// TODO: remove once all updates are merged
replace github.com/apache/arrow/go/v14 => github.com/cloudquery/arrow/go/v14 v14.0.0-20231023001216-f46436fa3561

// github.com/cloudquery/jsonschema @ cqmain
replace github.com/invopop/jsonschema => github.com/cloudquery/jsonschema v0.0.0-20231018073309-6c617a23d42f
