module github.com/cloudquery/cloudquery/cli

go 1.19

require (
	github.com/cloudquery/plugin-sdk/v2 v2.5.3
	github.com/getsentry/sentry-go v0.20.0
	github.com/google/uuid v1.3.0
	github.com/rs/zerolog v1.29.0
	github.com/schollz/progressbar/v3 v3.13.1
	github.com/spf13/cobra v1.6.1
	github.com/stretchr/testify v1.8.2
	github.com/thoas/go-funk v0.9.3
	golang.org/x/exp v0.0.0-20230425010034-47ecfdc1ba53
	golang.org/x/sync v0.1.0
	google.golang.org/grpc v1.54.0
	google.golang.org/protobuf v1.30.0
)

// TODO: remove once the changes are merged to upstream
replace github.com/apache/arrow/go/v12 => github.com/cloudquery/arrow/go/v12 v12.0.0-20230411134648-a5b71cd453c8

require (
	github.com/andybalholm/brotli v1.0.5 // indirect
	github.com/apache/arrow/go/v12 v12.0.0-20230430004532-0ea1a103dfc2 // indirect
	github.com/apache/thrift v0.16.0 // indirect
	github.com/avast/retry-go/v4 v4.3.4 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/goccy/go-json v0.9.11 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/flatbuffers v2.0.8+incompatible // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/klauspost/asmfmt v1.3.2 // indirect
	github.com/klauspost/compress v1.16.0 // indirect
	github.com/klauspost/cpuid/v2 v2.0.9 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.18 // indirect
	github.com/mattn/go-runewidth v0.0.14 // indirect
	github.com/minio/asm2plan9s v0.0.0-20200509001527-cdd76441f9d8 // indirect
	github.com/minio/c2goasm v0.0.0-20190812172519-36a3d3bbc4f3 // indirect
	github.com/mitchellh/colorstring v0.0.0-20190213212951-d06e56a500db // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rivo/uniseg v0.4.4 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/zeebo/xxh3 v1.0.2 // indirect
	golang.org/x/mod v0.8.0 // indirect
	golang.org/x/net v0.9.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/term v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	golang.org/x/tools v0.6.0 // indirect
	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2 // indirect
	google.golang.org/genproto v0.0.0-20230410155749-daa745c078e1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
