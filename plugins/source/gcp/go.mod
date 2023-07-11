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
	github.com/apache/arrow/go/v13 v13.0.0-20230630125530-5a06b2ec2a8e
	github.com/cloudquery/plugin-sdk/v4 v4.8.1-rc1
	github.com/cockroachdb/cockroachdb-parser v0.0.0-20230515042840-c9c144eab71a
	github.com/golang/mock v1.6.0
	github.com/googleapis/gax-go/v2 v2.7.1
	github.com/grpc-ecosystem/go-grpc-middleware/providers/zerolog/v2 v2.0.0-rc.3
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.0.0-rc.3
	github.com/iancoleman/strcase v0.2.0
	github.com/julienschmidt/httprouter v1.3.0
	github.com/mjibson/sqlfmt v0.5.0
	github.com/rs/zerolog v1.29.1
	github.com/spf13/cast v1.5.0
	github.com/stretchr/testify v1.8.4
	github.com/thoas/go-funk v0.9.3
	golang.org/x/exp v0.0.0-20230626212559-97b1e661b5df
	golang.org/x/sync v0.1.0
	google.golang.org/api v0.114.0
	google.golang.org/genproto v0.0.0-20230530153820-e85fd2cbaebc
	google.golang.org/grpc v1.55.0
	google.golang.org/protobuf v1.30.0
)

// TODO: remove once all updates are merged
replace github.com/apache/arrow/go/v13 => github.com/cloudquery/arrow/go/v13 v13.0.0-20230710001530-a2a76ebbb85f

require (
	cloud.google.com/go v0.110.0 // indirect
	cloud.google.com/go/compute/metadata v0.2.3 // indirect
	cloud.google.com/go/vision v1.2.0 // indirect
	github.com/biogo/store v0.0.0-20201120204734-aad293a2328f // indirect
	github.com/blevesearch/snowballstem v0.9.0 // indirect
	github.com/census-instrumentation/opencensus-proto v0.4.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/cloudquery/plugin-pb-go v1.6.0 // indirect
	github.com/cloudquery/plugin-sdk/v2 v2.7.0 // indirect
	github.com/cncf/udpa/go v0.0.0-20220112060539-c52dc94e7fbe // indirect
	github.com/cncf/xds/go v0.0.0-20230607035331-e9ce68804cb4 // indirect
	github.com/cockroachdb/apd/v3 v3.1.0 // indirect
	github.com/cockroachdb/errors v1.9.0 // indirect
	github.com/cockroachdb/logtags v0.0.0-20230118201751-21c54148d20b // indirect
	github.com/cockroachdb/redact v1.1.3 // indirect
	github.com/dave/dst v0.27.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/envoyproxy/go-control-plane v0.11.0 // indirect
	github.com/envoyproxy/protoc-gen-validate v0.10.0 // indirect
	github.com/getsentry/sentry-go v0.20.0 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/geo v0.0.0-20230421003525-6adc56603217 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/flatbuffers v23.1.21+incompatible // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.2.3 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/klauspost/compress v1.16.6 // indirect
	github.com/klauspost/cpuid/v2 v2.2.5 // indirect
	github.com/kr/pretty v0.3.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/lib/pq v1.10.6 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/petermattis/goid v0.0.0-20211229010228-4d14c490ee36 // indirect
	github.com/pierrec/lz4/v4 v4.1.17 // indirect
	github.com/pierrre/geohash v1.0.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rogpeppe/go-internal v1.9.0 // indirect
	github.com/sasha-s/go-deadlock v0.3.1 // indirect
	github.com/spf13/cobra v1.6.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/twpayne/go-geom v1.4.2 // indirect
	github.com/twpayne/go-kml v1.5.2 // indirect
	github.com/zeebo/xxh3 v1.0.2 // indirect
	go.opencensus.io v0.24.0 // indirect
	golang.org/x/mod v0.11.0 // indirect
	golang.org/x/net v0.9.0 // indirect
	golang.org/x/oauth2 v0.6.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	golang.org/x/tools v0.7.0 // indirect
	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20230526203410-71b5a4ffd15e // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230629202037-9506855d4529 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
