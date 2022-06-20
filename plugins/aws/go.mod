module github.com/cloudquery/cq-provider-aws

go 1.17

require (
	github.com/aws/aws-sdk-go-v2 v1.16.3
	github.com/aws/aws-sdk-go-v2/config v1.3.0
	github.com/aws/aws-sdk-go-v2/credentials v1.2.1
	github.com/aws/aws-sdk-go-v2/feature/s3/manager v1.2.1
	github.com/aws/aws-sdk-go-v2/service/accessanalyzer v1.15.0
	github.com/aws/aws-sdk-go-v2/service/acm v1.9.1
	github.com/aws/aws-sdk-go-v2/service/apigateway v1.4.0
	github.com/aws/aws-sdk-go-v2/service/apigatewayv2 v1.3.1
	github.com/aws/aws-sdk-go-v2/service/applicationautoscaling v1.10.2
	github.com/aws/aws-sdk-go-v2/service/autoscaling v1.2.0
	github.com/aws/aws-sdk-go-v2/service/cloudformation v1.20.0
	github.com/aws/aws-sdk-go-v2/service/cloudfront v1.3.0
	github.com/aws/aws-sdk-go-v2/service/cloudtrail v1.1.2
	github.com/aws/aws-sdk-go-v2/service/cloudwatch v1.1.2
	github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs v1.1.2
	github.com/aws/aws-sdk-go-v2/service/codebuild v1.14.1
	github.com/aws/aws-sdk-go-v2/service/codepipeline v1.13.3
	github.com/aws/aws-sdk-go-v2/service/cognitoidentity v1.3.1
	github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider v1.3.3
	github.com/aws/aws-sdk-go-v2/service/configservice v1.5.1
	github.com/aws/aws-sdk-go-v2/service/databasemigrationservice v1.12.0
	github.com/aws/aws-sdk-go-v2/service/dax v1.7.2
	github.com/aws/aws-sdk-go-v2/service/directconnect v1.4.1
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.10.0
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.32.1
	github.com/aws/aws-sdk-go-v2/service/ecr v1.2.0
	github.com/aws/aws-sdk-go-v2/service/ecs v1.18.1
	github.com/aws/aws-sdk-go-v2/service/efs v1.2.0
	github.com/aws/aws-sdk-go-v2/service/eks v1.2.1
	github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk v1.14.3
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing v1.3.0
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 v1.2.0
	github.com/aws/aws-sdk-go-v2/service/elasticsearchservice v1.4.0
	github.com/aws/aws-sdk-go-v2/service/emr v1.2.0
	github.com/aws/aws-sdk-go-v2/service/fsx v1.2.0
	github.com/aws/aws-sdk-go-v2/service/guardduty v1.7.1
	github.com/aws/aws-sdk-go-v2/service/iam v1.3.0
	github.com/aws/aws-sdk-go-v2/service/iot v1.20.0
	github.com/aws/aws-sdk-go-v2/service/kms v1.17.0
	github.com/aws/aws-sdk-go-v2/service/lambda v1.22.0
	github.com/aws/aws-sdk-go-v2/service/mq v1.2.1
	github.com/aws/aws-sdk-go-v2/service/organizations v1.2.1
	github.com/aws/aws-sdk-go-v2/service/qldb v1.14.3
	github.com/aws/aws-sdk-go-v2/service/rds v1.2.1
	github.com/aws/aws-sdk-go-v2/service/redshift v1.3.0
	github.com/aws/aws-sdk-go-v2/service/route53 v1.4.0
	github.com/aws/aws-sdk-go-v2/service/route53domains v1.6.0
	github.com/aws/aws-sdk-go-v2/service/s3 v1.19.1
	github.com/aws/aws-sdk-go-v2/service/s3control v1.14.1
	github.com/aws/aws-sdk-go-v2/service/sagemaker v1.19.1
	github.com/aws/aws-sdk-go-v2/service/secretsmanager v1.10.2
	github.com/aws/aws-sdk-go-v2/service/sns v1.1.2
	github.com/aws/aws-sdk-go-v2/service/sqs v1.9.1
	github.com/aws/aws-sdk-go-v2/service/ssm v1.16.0
	github.com/aws/aws-sdk-go-v2/service/sts v1.4.1
	github.com/aws/aws-sdk-go-v2/service/waf v1.2.1
	github.com/aws/aws-sdk-go-v2/service/wafv2 v1.5.1
	github.com/aws/aws-sdk-go-v2/service/workspaces v1.16.0
	github.com/aws/aws-sdk-go-v2/service/xray v1.13.4
	github.com/aws/smithy-go v1.11.3
	github.com/basgys/goxml2json v1.1.0
	github.com/bxcodec/faker v2.0.1+incompatible
	github.com/cloudquery/cq-gen v0.0.5
	github.com/cloudquery/cq-provider-sdk v0.11.4
	github.com/cloudquery/faker/v3 v3.7.5
	github.com/gocarina/gocsv v0.0.0-20210516172204-ca9e8a8ddea8
	github.com/golang/mock v1.6.0
	github.com/google/go-cmp v0.5.8
	github.com/hashicorp/go-hclog v1.2.1
	github.com/mitchellh/mapstructure v1.4.3
	github.com/spf13/cast v1.4.1
	github.com/stretchr/testify v1.7.2
)

require (
	github.com/BurntSushi/toml v1.0.0 // indirect
	github.com/elliotchance/orderedmap v1.4.0 // indirect
	github.com/lorenzosaino/go-sysctl v0.3.0 // indirect
	github.com/mitchellh/hashstructure/v2 v2.0.2 // indirect
	github.com/segmentio/stats/v4 v4.6.3 // indirect
	golang.org/x/exp/typeparams v0.0.0-20220317015231-48e79f11773a // indirect
	golang.org/x/lint v0.0.0-20210508222113-6edffad5e616 // indirect
	honnef.co/go/tools v0.3.0-0.dev.0.20220306074811-23e1086441d2 // indirect
)

require (
	github.com/Masterminds/squirrel v1.5.3 // indirect
	github.com/agext/levenshtein v1.2.3 // indirect
	github.com/agnivade/levenshtein v1.0.1 // indirect
	github.com/apparentlymart/go-textseg/v13 v13.0.0 // indirect
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.0.0 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.1.1 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.10 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.4 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.0.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/athena v1.15.1
	github.com/aws/aws-sdk-go-v2/service/backup v1.15.4
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.5.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.3.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.9.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.9.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/shield v1.16.4
	github.com/aws/aws-sdk-go-v2/service/sso v1.2.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/wafregional v1.12.3
	github.com/creasty/defaults v1.6.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/doug-martin/goqu/v9 v9.18.0 // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/georgysavva/scany v0.2.9 // indirect
	github.com/getkin/kin-openapi v0.83.0 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/swag v0.19.5 // indirect
	github.com/gofrs/uuid v4.2.0+incompatible // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-plugin v1.4.4 // indirect
	github.com/hashicorp/hcl/v2 v2.11.1 // indirect
	github.com/hashicorp/yamux v0.0.0-20210826001029-26ff87cf9493 // indirect
	github.com/iancoleman/strcase v0.2.0 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.10.0 // indirect
	github.com/jackc/pgerrcode v0.0.0-20220416144525-469b46aa5efa // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.1.1 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.8.1 // indirect
	github.com/jackc/pgx/v4 v4.13.0 // indirect
	github.com/jackc/puddle v1.1.4 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/lann/builder v0.0.0-20180802200727-47ae307949d0 // indirect
	github.com/lann/ps v0.0.0-20150810152359-62de8c46ede0 // indirect
	github.com/mailru/easyjson v0.7.0 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mitchellh/go-testing-interface v1.14.1 // indirect
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/oklog/run v1.1.0 // indirect
	github.com/pbnjay/memory v0.0.0-20210728143218-7b4eea64cf58 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/spf13/afero v1.8.1 // indirect
	github.com/stretchr/objx v0.3.0 // indirect
	github.com/thoas/go-funk v0.9.2
	github.com/vektah/gqlparser/v2 v2.2.0 // indirect
	github.com/vmihailenco/msgpack/v5 v5.3.5 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	github.com/xo/dburl v0.8.4 // indirect
	github.com/zclconf/go-cty v1.10.0 // indirect
	golang.org/x/crypto v0.0.0-20211108221036-ceb1ce70b4fa // indirect
	golang.org/x/mod v0.6.0-dev.0.20220106191415-9b9b3d81d5e3 // indirect
	golang.org/x/net v0.0.0-20220225172249-27dd8689420f // indirect
	golang.org/x/sync v0.0.0-20220601150217-0de741cfad7f
	golang.org/x/sys v0.0.0-20220503163025-988cb79eb6c6 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/tools v0.1.10 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/genproto v0.0.0-20220314164441-57ef72a4c106 // indirect
	google.golang.org/grpc v1.45.0 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
