module github.com/cloudquery/cloudquery/plugins/destination/azblob

go 1.19

require (
	github.com/Azure/azure-sdk-for-go/sdk/azidentity v1.2.2
	github.com/Azure/azure-sdk-for-go/sdk/storage/azblob v1.0.0
	github.com/apache/arrow/go/v12 v12.0.0-20230417014917-9888ac36c142
	github.com/cloudquery/filetypes/v2 v2.0.3
	github.com/cloudquery/plugin-sdk/v2 v2.5.0
	github.com/google/uuid v1.3.0
	github.com/rs/zerolog v1.29.0
)

replace github.com/apache/arrow/go/v12 => github.com/cloudquery/arrow/go/v12 v12.0.0-20230419074556-00ceafa3b033

require (
	github.com/JohnCGriffin/overflow v0.0.0-20211019200055-46fa312c352c // indirect
	github.com/andybalholm/brotli v1.0.5 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dnaeon/go-vcr v1.2.0 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/google/flatbuffers v23.3.3+incompatible // indirect
	github.com/klauspost/asmfmt v1.3.2 // indirect
	github.com/klauspost/cpuid/v2 v2.2.4 // indirect
	github.com/minio/asm2plan9s v0.0.0-20200509001527-cdd76441f9d8 // indirect
	github.com/minio/c2goasm v0.0.0-20190812172519-36a3d3bbc4f3 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/testify v1.8.2 // indirect
	github.com/zeebo/xxh3 v1.0.2 // indirect
	golang.org/x/exp v0.0.0-20230321023759-10a507213a29 // indirect
	golang.org/x/mod v0.9.0 // indirect
	golang.org/x/tools v0.7.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

require (
	github.com/apache/thrift v0.18.1 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/klauspost/compress v1.16.3 // indirect
	github.com/pierrec/lz4/v4 v4.1.17 // indirect
	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2 // indirect; indirect // indirect
)

require (
	github.com/Azure/azure-sdk-for-go/sdk/azcore v1.4.0 // indirect; indirect // indirect // indirect
	github.com/Azure/azure-sdk-for-go/sdk/internal v1.2.0 // indirect; indirect // indirect
	github.com/AzureAD/microsoft-authentication-library-for-go v0.9.0 // indirect
	github.com/getsentry/sentry-go v0.20.0 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/golang-jwt/jwt/v4 v4.5.0 // indirect; indirect // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware/providers/zerolog/v2 v2.0.0-rc.3 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.0.0-rc.3 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/kylelemons/godebug v1.1.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.18 // indirect
	github.com/pkg/browser v0.0.0-20210911075715-681adbf594b8 // indirect; indirect // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/cobra v1.6.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/thoas/go-funk v0.9.3 // indirect; indirect // indirect
	golang.org/x/crypto v0.7.0 // indirect
	golang.org/x/net v0.9.0 // indirect; indirect // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/genproto v0.0.0-20230331144136-dcfb400f0633 // indirect; indirect // indirect
	google.golang.org/grpc v1.54.0 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
