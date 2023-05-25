module github.com/cloudquery/plugins/source/gcp

go 1.19

require (
	cloud.google.com/go/aiplatform v1.37.0
	cloud.google.com/go/apigateway v1.5.0
	cloud.google.com/go/apikeys v0.6.0
	cloud.google.com/go/appengine v1.7.1
	cloud.google.com/go/artifactregistry v1.13.0
	cloud.google.com/go/baremetalsolution v0.5.0
	cloud.google.com/go/batch v0.7.0
	cloud.google.com/go/beyondcorp v0.5.0
	cloud.google.com/go/bigtable v1.18.1
	cloud.google.com/go/billing v1.13.0
	cloud.google.com/go/binaryauthorization v1.5.0
	cloud.google.com/go/certificatemanager v1.6.0
	cloud.google.com/go/compute v1.19.0
	cloud.google.com/go/container v1.15.0
	cloud.google.com/go/containeranalysis v0.9.0
	cloud.google.com/go/deploy v1.8.0
	cloud.google.com/go/domains v0.8.0
	cloud.google.com/go/errorreporting v0.3.0
	cloud.google.com/go/functions v1.13.0
	cloud.google.com/go/iam v0.13.0
	cloud.google.com/go/iot v1.6.0
	cloud.google.com/go/kms v1.10.1
	cloud.google.com/go/logging v1.7.0
	cloud.google.com/go/longrunning v0.4.1
	cloud.google.com/go/monitoring v1.13.0
	cloud.google.com/go/redis v1.11.0
	cloud.google.com/go/resourcemanager v1.7.0
	cloud.google.com/go/run v0.9.0
	cloud.google.com/go/scheduler v1.9.0
	cloud.google.com/go/secretmanager v1.10.0
	cloud.google.com/go/securitycenter v1.19.0
	cloud.google.com/go/serviceusage v1.6.0
	cloud.google.com/go/storage v1.28.1
	cloud.google.com/go/translate v1.7.0
	cloud.google.com/go/video v1.15.0
	cloud.google.com/go/vision/v2 v2.7.0
	cloud.google.com/go/vmmigration v1.6.0
	cloud.google.com/go/vpcaccess v1.6.0
	cloud.google.com/go/websecurityscanner v1.5.0
	cloud.google.com/go/workflows v1.10.0
	github.com/apache/arrow/go/v13 v13.0.0-20230509040948-de6c3cd2b604
	github.com/cloudquery/plugin-pb-go v1.0.8
	github.com/cloudquery/plugin-sdk/v3 v3.6.3
	github.com/golang/mock v1.6.0
	github.com/googleapis/gax-go/v2 v2.7.1
	github.com/grpc-ecosystem/go-grpc-middleware/providers/zerolog/v2 v2.0.0-rc.3
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.0.0-rc.3
	github.com/iancoleman/strcase v0.2.0
	github.com/julienschmidt/httprouter v1.3.0
	github.com/rs/zerolog v1.29.0
	github.com/spf13/cast v1.5.0
	github.com/stretchr/testify v1.8.2
	github.com/thoas/go-funk v0.9.3
	golang.org/x/exp v0.0.0-20230425010034-47ecfdc1ba53
	golang.org/x/sync v0.1.0
	google.golang.org/api v0.114.0
	google.golang.org/genproto v0.0.0-20230410155749-daa745c078e1
	google.golang.org/grpc v1.54.0
	google.golang.org/protobuf v1.30.0
)

// TODO: remove once all updates are merged
replace github.com/apache/arrow/go/v13 => github.com/cloudquery/arrow/go/v13 v13.0.0-20230509053643-898a79b1d3c8

require (
	cloud.google.com/go v0.110.0 // indirect
	cloud.google.com/go/compute/metadata v0.2.3 // indirect
	cloud.google.com/go/vision v1.2.0 // indirect
	github.com/andybalholm/brotli v1.0.5 // indirect
	github.com/apache/thrift v0.16.0 // indirect
	github.com/census-instrumentation/opencensus-proto v0.4.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/cloudquery/plugin-sdk/v2 v2.7.0 // indirect
	github.com/cncf/udpa/go v0.0.0-20220112060539-c52dc94e7fbe // indirect
	github.com/cncf/xds/go v0.0.0-20230428030218-4003588d1b74 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/envoyproxy/go-control-plane v0.10.3 // indirect
	github.com/envoyproxy/protoc-gen-validate v0.9.1 // indirect
	github.com/getsentry/sentry-go v0.20.0 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/goccy/go-json v0.9.11 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/flatbuffers v2.0.8+incompatible // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.2.3 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/klauspost/asmfmt v1.3.2 // indirect
	github.com/klauspost/compress v1.16.0 // indirect
	github.com/klauspost/cpuid/v2 v2.0.9 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.18 // indirect
	github.com/minio/asm2plan9s v0.0.0-20200509001527-cdd76441f9d8 // indirect
	github.com/minio/c2goasm v0.0.0-20190812172519-36a3d3bbc4f3 // indirect
	github.com/pierrec/lz4/v4 v4.1.15 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/spf13/cobra v1.6.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/zeebo/xxh3 v1.0.2 // indirect
	go.opencensus.io v0.24.0 // indirect
	golang.org/x/mod v0.8.0 // indirect
	golang.org/x/net v0.9.0 // indirect
	golang.org/x/oauth2 v0.6.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	golang.org/x/tools v0.6.0 // indirect
	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
