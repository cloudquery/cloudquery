module github.com/cloudquery/cq-provider-aws

go 1.17

require (
	github.com/Masterminds/squirrel v1.5.0
	github.com/aws/aws-sdk-go-v2 v1.11.0
	github.com/aws/aws-sdk-go-v2/config v1.3.0
	github.com/aws/aws-sdk-go-v2/credentials v1.2.1
	github.com/aws/aws-sdk-go-v2/feature/s3/manager v1.2.1
	github.com/aws/aws-sdk-go-v2/internal/ini v1.0.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/accessanalyzer v1.4.1
	github.com/aws/aws-sdk-go-v2/service/apigateway v1.4.0
	github.com/aws/aws-sdk-go-v2/service/apigatewayv2 v1.3.1
	github.com/aws/aws-sdk-go-v2/service/autoscaling v1.2.0
	github.com/aws/aws-sdk-go-v2/service/cloudfront v1.3.0
	github.com/aws/aws-sdk-go-v2/service/cloudtrail v1.1.2
	github.com/aws/aws-sdk-go-v2/service/cloudwatch v1.1.2
	github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs v1.1.2
	github.com/aws/aws-sdk-go-v2/service/cognitoidentity v1.3.1
	github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider v1.3.3
	github.com/aws/aws-sdk-go-v2/service/configservice v1.5.1
	github.com/aws/aws-sdk-go-v2/service/directconnect v1.4.1
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.16.0
	github.com/aws/aws-sdk-go-v2/service/ecr v1.2.0
	github.com/aws/aws-sdk-go-v2/service/ecs v1.2.0
	github.com/aws/aws-sdk-go-v2/service/efs v1.2.0
	github.com/aws/aws-sdk-go-v2/service/eks v1.2.1
	github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk v1.2.0
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing v1.3.0
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 v1.2.0
	github.com/aws/aws-sdk-go-v2/service/elasticsearchservice v1.4.0
	github.com/aws/aws-sdk-go-v2/service/emr v1.2.0
	github.com/aws/aws-sdk-go-v2/service/fsx v1.2.0
	github.com/aws/aws-sdk-go-v2/service/iam v1.3.0
	github.com/aws/aws-sdk-go-v2/service/kms v1.2.1
	github.com/aws/aws-sdk-go-v2/service/lambda v1.3.0
	github.com/aws/aws-sdk-go-v2/service/mq v1.2.1
	github.com/aws/aws-sdk-go-v2/service/organizations v1.2.1
	github.com/aws/aws-sdk-go-v2/service/rds v1.2.1
	github.com/aws/aws-sdk-go-v2/service/redshift v1.3.0
	github.com/aws/aws-sdk-go-v2/service/route53 v1.4.0
	github.com/aws/aws-sdk-go-v2/service/s3 v1.8.0
	github.com/aws/aws-sdk-go-v2/service/sns v1.1.2
	github.com/aws/aws-sdk-go-v2/service/sqs v1.9.1
	github.com/aws/aws-sdk-go-v2/service/sts v1.4.1
	github.com/aws/aws-sdk-go-v2/service/waf v1.2.1
	github.com/aws/aws-sdk-go-v2/service/wafv2 v1.5.1
	github.com/aws/smithy-go v1.9.0
	github.com/cloudquery/cq-provider-sdk v0.5.1
	github.com/cloudquery/faker/v3 v3.7.5
	github.com/gocarina/gocsv v0.0.0-20210516172204-ca9e8a8ddea8
	github.com/golang/mock v1.6.0
	github.com/hashicorp/go-hclog v1.0.0
	github.com/jackc/pgx/v4 v4.13.0
	github.com/mitchellh/mapstructure v1.4.2
	github.com/spf13/cast v1.4.1
	github.com/stretchr/testify v1.7.0
)

require (
	github.com/agext/levenshtein v1.2.3 // indirect
	github.com/apparentlymart/go-textseg/v13 v13.0.0 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.1.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.1.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.3.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.3.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/route53domains v1.6.0
	github.com/aws/aws-sdk-go-v2/service/sso v1.2.1 // indirect
	github.com/creasty/defaults v1.5.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/doug-martin/goqu/v9 v9.17.0 // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/georgysavva/scany v0.2.9 // indirect
	github.com/go-test/deep v1.0.7 // indirect
	github.com/gofrs/uuid v4.0.0+incompatible // indirect
	github.com/golang-migrate/migrate/v4 v4.15.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-plugin v1.4.3 // indirect
	github.com/hashicorp/go-version v1.3.0 // indirect
	github.com/hashicorp/hcl/v2 v2.10.1 // indirect
	github.com/hashicorp/terraform-exec v0.15.0 // indirect
	github.com/hashicorp/terraform-json v0.13.0 // indirect
	github.com/hashicorp/yamux v0.0.0-20210826001029-26ff87cf9493 // indirect
	github.com/huandu/go-sqlbuilder v1.13.0 // indirect
	github.com/huandu/xstrings v1.3.2 // indirect
	github.com/iancoleman/strcase v0.2.0 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.10.0 // indirect
	github.com/jackc/pgerrcode v0.0.0-20201024163028-a0d42d470451 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.1.1 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.8.1 // indirect
	github.com/jackc/puddle v1.1.4 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/lann/builder v0.0.0-20180802200727-47ae307949d0 // indirect
	github.com/lann/ps v0.0.0-20150810152359-62de8c46ede0 // indirect
	github.com/lib/pq v1.10.3 // indirect
	github.com/mattn/go-colorable v0.1.11 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mitchellh/go-testing-interface v1.14.1 // indirect
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	github.com/mitchellh/hashstructure v1.1.0 // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/oklog/run v1.1.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/spf13/afero v1.6.0 // indirect
	github.com/stretchr/objx v0.3.0 // indirect
	github.com/thoas/go-funk v0.9.1 // indirect
	github.com/tmccombs/hcl2json v0.3.3 // indirect
	github.com/vmihailenco/msgpack/v5 v5.3.4 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	github.com/xo/dburl v0.8.4 // indirect
	github.com/zclconf/go-cty v1.9.1 // indirect
	go.uber.org/atomic v1.6.0 // indirect
	golang.org/x/crypto v0.0.0-20210711020723-a769d52b0f97 // indirect
	golang.org/x/net v0.0.0-20210520170846-37e1c6afe023 // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	golang.org/x/sys v0.0.0-20210927094055-39ccf1dd6fa6 // indirect
	golang.org/x/text v0.3.6 // indirect
	google.golang.org/genproto v0.0.0-20211005153810-c76a74d43a8e // indirect
	google.golang.org/grpc v1.41.0 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

require (
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.0 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.0.0 // indirect
)
