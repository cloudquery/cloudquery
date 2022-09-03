module github.com/cloudquery/plugins/source/gcp

go 1.18

require (
	github.com/bxcodec/faker/v4 v4.0.0-beta
	github.com/cloudquery/faker/v3 v3.7.7
	github.com/cloudquery/plugin-sdk v0.0.11
	github.com/golang/mock v1.6.0
	github.com/googleapis/gax-go/v2 v2.4.0
	github.com/iancoleman/strcase v0.2.0
	github.com/julienschmidt/httprouter v1.3.0
	github.com/pkg/errors v0.9.1
	github.com/rs/zerolog v1.28.0
	github.com/spf13/cast v1.5.0
	github.com/stretchr/testify v1.8.0
	github.com/thoas/go-funk v0.9.2
	google.golang.org/api v0.92.0
	google.golang.org/grpc v1.49.0
)

replace github.com/bxcodec/faker/v4 => github.com/cloudquery/faker/v4 v4.0.1-beta

replace github.com/cloudquery/plugin-sdk => ../../../../plugin-sdk

require (
	cloud.google.com/go v0.102.1 // indirect
	cloud.google.com/go/asset v1.4.0 // indirect
	cloud.google.com/go/compute v1.7.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/gofrs/uuid v4.2.0+incompatible // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.8 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.1.0 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware/providers/zerolog/v2 v2.0.0-rc.2 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.0.0-rc.2 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/spf13/cobra v1.5.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/xeipuuv/gojsonpointer v0.0.0-20180127040702-4e3ac2762d5f // indirect
	github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415 // indirect
	github.com/xeipuuv/gojsonschema v1.2.0 // indirect
	go.opencensus.io v0.23.0 // indirect
	golang.org/x/mod v0.6.0-dev.0.20220419223038-86c51ed26bb4 // indirect
	golang.org/x/net v0.0.0-20220728211354-c7608f3a8462 // indirect
	golang.org/x/oauth2 v0.0.0-20220622183110-fd043fe589d2 // indirect
	golang.org/x/sync v0.0.0-20220819030929-7fc1605a5dde // indirect
	golang.org/x/sys v0.0.0-20220731174439-a90be440212d // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/tools v0.1.12 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20220812140447-cec7f5303424 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
