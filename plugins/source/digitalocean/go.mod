module github.com/cloudquery/cloudquery/plugins/source/digitalocean

go 1.18

require (
	github.com/avast/retry-go v3.0.0+incompatible
	github.com/aws/aws-sdk-go-v2 v1.16.7
	github.com/aws/aws-sdk-go-v2/config v1.8.0
	github.com/aws/aws-sdk-go-v2/service/s3 v1.27.1
	github.com/aws/smithy-go v1.12.0
	github.com/cloudquery/cq-provider-sdk v0.14.7
	github.com/cloudquery/faker/v3 v3.7.7
	github.com/cloudquery/plugin-sdk v0.1.1
	github.com/digitalocean/godo v1.81.0
	github.com/golang/mock v1.6.0
	github.com/iancoleman/strcase v0.2.0
	github.com/julienschmidt/httprouter v1.3.0
	github.com/pkg/errors v0.9.1
	github.com/rs/zerolog v1.28.0
	github.com/spf13/cast v1.5.0
	github.com/thoas/go-funk v0.9.2
)

replace github.com/cloudquery/plugin-sdk => ../../../../plugin-sdk

require (
	github.com/agext/levenshtein v1.2.3 // indirect
	github.com/apparentlymart/go-textseg/v13 v13.0.0 // indirect
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.4.3 // indirect
	github.com/aws/aws-sdk-go-v2/credentials v1.4.0 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.5.0 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.14 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.8 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.2.2 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.0.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.9.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.1.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.9.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.13.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.4.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.7.0 // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/gofrs/uuid v4.2.0+incompatible // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware/providers/zerolog/v2 v2.0.0-rc.2 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.0.0-rc.2 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-hclog v1.2.2 // indirect
	github.com/hashicorp/hcl/v2 v2.13.0 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	github.com/mitchellh/hashstructure/v2 v2.0.2 // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/spf13/cobra v1.5.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/xeipuuv/gojsonpointer v0.0.0-20180127040702-4e3ac2762d5f // indirect
	github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415 // indirect
	github.com/xeipuuv/gojsonschema v1.2.0 // indirect
	github.com/zclconf/go-cty v1.10.0 // indirect
	golang.org/x/mod v0.6.0-dev.0.20220419223038-86c51ed26bb4 // indirect
	golang.org/x/net v0.0.0-20220728211354-c7608f3a8462 // indirect
	golang.org/x/oauth2 v0.0.0-20220411215720-9780585627b5 // indirect
	golang.org/x/sync v0.0.0-20220819030929-7fc1605a5dde // indirect
	golang.org/x/sys v0.0.0-20220731174439-a90be440212d // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/tools v0.1.12 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20220801145646-83ce21fca29f // indirect
	google.golang.org/grpc v1.49.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
