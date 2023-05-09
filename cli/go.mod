module github.com/cloudquery/cloudquery/cli

go 1.19

require (
	github.com/avast/retry-go/v4 v4.3.3
	github.com/cloudquery/plugin-pb-go v1.0.5
	github.com/getsentry/sentry-go v0.20.0
	github.com/google/go-cmp v0.5.9
	github.com/google/uuid v1.3.0
	github.com/rs/zerolog v1.29.0
	github.com/schollz/progressbar/v3 v3.13.1
	github.com/spf13/cobra v1.6.1
	github.com/stretchr/testify v1.8.2
	github.com/thoas/go-funk v0.9.3
	golang.org/x/exp v0.0.0-20230425010034-47ecfdc1ba53
	google.golang.org/grpc v1.54.0
	google.golang.org/protobuf v1.30.0
)

// TODO: remove once the changes are merged to upstream
replace github.com/apache/arrow/go/v12 => github.com/cloudquery/arrow/go/v12 v12.0.0-20230411134648-a5b71cd453c8

require (
	github.com/cpuguy83/go-md2man/v2 v2.0.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/kr/pretty v0.3.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.18 // indirect
	github.com/mattn/go-runewidth v0.0.14 // indirect
	github.com/mitchellh/colorstring v0.0.0-20190213212951-d06e56a500db // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rivo/uniseg v0.4.4 // indirect
	github.com/rogpeppe/go-internal v1.9.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/net v0.9.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/term v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/genproto v0.0.0-20230331144136-dcfb400f0633 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
