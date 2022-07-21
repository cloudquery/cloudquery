module github.com/cloudquery/cq-provider-aws

go 1.17

require (
	github.com/aws/aws-sdk-go-v2 v1.16.7
	github.com/aws/aws-sdk-go-v2/config v1.15.14
	github.com/aws/aws-sdk-go-v2/credentials v1.12.9
	github.com/aws/aws-sdk-go-v2/feature/s3/manager v1.11.20
	github.com/aws/aws-sdk-go-v2/service/accessanalyzer v1.15.8
	github.com/aws/aws-sdk-go-v2/service/acm v1.14.8
	github.com/aws/aws-sdk-go-v2/service/apigateway v1.15.10
	github.com/aws/aws-sdk-go-v2/service/apigatewayv2 v1.12.8
	github.com/aws/aws-sdk-go-v2/service/applicationautoscaling v1.15.8
	github.com/aws/aws-sdk-go-v2/service/autoscaling v1.23.5
	github.com/aws/aws-sdk-go-v2/service/cloudformation v1.21.2
	github.com/aws/aws-sdk-go-v2/service/cloudfront v1.18.4
	github.com/aws/aws-sdk-go-v2/service/cloudtrail v1.16.4
	github.com/aws/aws-sdk-go-v2/service/cloudwatch v1.18.6
	github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs v1.15.10
	github.com/aws/aws-sdk-go-v2/service/codebuild v1.19.7
	github.com/aws/aws-sdk-go-v2/service/codepipeline v1.13.7
	github.com/aws/aws-sdk-go-v2/service/cognitoidentity v1.13.7
	github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider v1.17.2
	github.com/aws/aws-sdk-go-v2/service/configservice v1.21.4
	github.com/aws/aws-sdk-go-v2/service/databasemigrationservice v1.20.0
	github.com/aws/aws-sdk-go-v2/service/dax v1.11.7
	github.com/aws/aws-sdk-go-v2/service/directconnect v1.17.7
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.15.8
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.47.1
	github.com/aws/aws-sdk-go-v2/service/ecr v1.17.7
	github.com/aws/aws-sdk-go-v2/service/ecs v1.18.10
	github.com/aws/aws-sdk-go-v2/service/efs v1.17.5
	github.com/aws/aws-sdk-go-v2/service/eks v1.21.3
	github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk v1.14.8
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing v1.14.7
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 v1.18.7
	github.com/aws/aws-sdk-go-v2/service/elasticsearchservice v1.15.7
	github.com/aws/aws-sdk-go-v2/service/emr v1.20.0
	github.com/aws/aws-sdk-go-v2/service/fsx v1.24.2
	github.com/aws/aws-sdk-go-v2/service/guardduty v1.14.1
	github.com/aws/aws-sdk-go-v2/service/iam v1.18.8
	github.com/aws/aws-sdk-go-v2/service/iot v1.25.4
	github.com/aws/aws-sdk-go-v2/service/kms v1.17.4
	github.com/aws/aws-sdk-go-v2/service/lambda v1.23.3
	github.com/aws/aws-sdk-go-v2/service/lightsail v1.22.2
	github.com/aws/aws-sdk-go-v2/service/mq v1.13.3
	github.com/aws/aws-sdk-go-v2/service/organizations v1.16.3
	github.com/aws/aws-sdk-go-v2/service/qldb v1.14.8
	github.com/aws/aws-sdk-go-v2/service/rds v1.21.5
	github.com/aws/aws-sdk-go-v2/service/redshift v1.25.1
	github.com/aws/aws-sdk-go-v2/service/route53 v1.21.2
	github.com/aws/aws-sdk-go-v2/service/route53domains v1.12.7
	github.com/aws/aws-sdk-go-v2/service/s3 v1.27.1
	github.com/aws/aws-sdk-go-v2/service/s3control v1.21.8
	github.com/aws/aws-sdk-go-v2/service/sagemaker v1.34.0
	github.com/aws/aws-sdk-go-v2/service/secretsmanager v1.15.12
	github.com/aws/aws-sdk-go-v2/service/sesv2 v1.13.8
	github.com/aws/aws-sdk-go-v2/service/sns v1.17.8
	github.com/aws/aws-sdk-go-v2/service/sqs v1.18.7
	github.com/aws/aws-sdk-go-v2/service/ssm v1.27.3
	github.com/aws/aws-sdk-go-v2/service/sts v1.16.9
	github.com/aws/aws-sdk-go-v2/service/waf v1.11.7
	github.com/aws/aws-sdk-go-v2/service/wafv2 v1.20.4
	github.com/aws/aws-sdk-go-v2/service/workspaces v1.19.1
	github.com/aws/aws-sdk-go-v2/service/xray v1.13.8
	github.com/aws/smithy-go v1.12.0
	github.com/basgys/goxml2json v1.1.0
	github.com/bxcodec/faker v2.0.1+incompatible
	github.com/cloudquery/cq-gen v0.0.6
	github.com/cloudquery/cq-provider-sdk v0.14.2
	github.com/cloudquery/faker/v3 v3.7.7
	github.com/gocarina/gocsv v0.0.0-20220712153207-8b2118da4570
	github.com/golang/mock v1.6.0
	github.com/google/go-cmp v0.5.8
	github.com/hashicorp/go-hclog v1.2.1
	github.com/mitchellh/mapstructure v1.5.0
	github.com/spf13/cast v1.5.0
	github.com/stretchr/testify v1.8.0
	github.com/thoas/go-funk v0.9.2
	golang.org/x/sync v0.0.0-20220601150217-0de741cfad7f
)

require (
	github.com/BurntSushi/toml v1.1.0 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.0.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.1.9 // indirect
	github.com/elliotchance/orderedmap v1.4.0 // indirect
	github.com/lorenzosaino/go-sysctl v0.3.1 // indirect
	github.com/mitchellh/hashstructure/v2 v2.0.2 // indirect
	github.com/segmentio/stats/v4 v4.6.3 // indirect
	golang.org/x/exp/typeparams v0.0.0-20220613132600-b0d781184e0d // indirect
	golang.org/x/lint v0.0.0-20210508222113-6edffad5e616 // indirect
	honnef.co/go/tools v0.3.2 // indirect
)

require (
	github.com/Masterminds/squirrel v1.5.3 // indirect
	github.com/agext/levenshtein v1.2.3 // indirect
	github.com/agnivade/levenshtein v1.0.1 // indirect
	github.com/apparentlymart/go-textseg/v13 v13.0.0 // indirect
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.4.3 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.12.8 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.14 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.8 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.3.15 // indirect
	github.com/aws/aws-sdk-go-v2/service/athena v1.16.0
	github.com/aws/aws-sdk-go-v2/service/backup v1.16.3
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.9.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.7.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.9.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.13.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/shield v1.16.7
	github.com/aws/aws-sdk-go-v2/service/sso v1.11.12 // indirect
	github.com/aws/aws-sdk-go-v2/service/wafregional v1.12.8
	github.com/creasty/defaults v1.6.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/doug-martin/goqu/v9 v9.18.0 // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/georgysavva/scany v1.0.0 // indirect
	github.com/getkin/kin-openapi v0.83.0 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/swag v0.19.14 // indirect
	github.com/gofrs/uuid v4.2.0+incompatible // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-plugin v1.4.4 // indirect
	github.com/hashicorp/hcl/v2 v2.13.0 // indirect
	github.com/hashicorp/yamux v0.0.0-20211028200310-0bc27b27de87 // indirect
	github.com/iancoleman/strcase v0.2.0 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.12.1 // indirect
	github.com/jackc/pgerrcode v0.0.0-20220416144525-469b46aa5efa // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.11.0 // indirect
	github.com/jackc/pgx/v4 v4.16.1 // indirect
	github.com/jackc/puddle v1.2.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/lann/builder v0.0.0-20180802200727-47ae307949d0 // indirect
	github.com/lann/ps v0.0.0-20150810152359-62de8c46ede0 // indirect
	github.com/mailru/easyjson v0.7.6 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mitchellh/go-testing-interface v1.14.1 // indirect
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/oklog/run v1.1.0 // indirect
	github.com/pbnjay/memory v0.0.0-20210728143218-7b4eea64cf58 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/spf13/afero v1.8.2 // indirect
	github.com/vektah/gqlparser/v2 v2.2.0 // indirect
	github.com/vmihailenco/msgpack/v5 v5.3.5 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	github.com/xo/dburl v0.11.0 // indirect
	github.com/zclconf/go-cty v1.10.0 // indirect
	golang.org/x/crypto v0.0.0-20220622213112-05595931fe9d // indirect
	golang.org/x/mod v0.6.0-dev.0.20220419223038-86c51ed26bb4 // indirect
	golang.org/x/net v0.0.0-20220630215102-69896b714898 // indirect
	golang.org/x/sys v0.0.0-20220704084225-05e143d24a9e // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/tools v0.1.11 // indirect
	google.golang.org/genproto v0.0.0-20220630174209-ad1d48641aa7 // indirect
	google.golang.org/grpc v1.48.0 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
