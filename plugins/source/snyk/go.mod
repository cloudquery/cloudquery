module github.com/cloudquery/cloudquery/plugins/source/snyk

go 1.19

require (
	github.com/cloudquery/plugin-sdk v1.44.0
	github.com/google/uuid v1.3.0
	github.com/julienschmidt/httprouter v1.3.0
	github.com/pavel-snyk/snyk-sdk-go v0.4.1
	github.com/rs/zerolog v1.29.0
	golang.org/x/exp v0.0.0-20230224173230-c95f2b4c22f2
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/testify v1.8.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/pavel-snyk/snyk-sdk-go => github.com/cloudquery/snyk-sdk-go v0.1.0

require (
	github.com/getsentry/sentry-go v0.18.0 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware/providers/zerolog/v2 v2.0.0-rc.3 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.0.0-rc.3 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/cobra v1.6.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/thoas/go-funk v0.9.3 // indirect
	golang.org/x/net v0.7.0 // indirect; indirect // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	google.golang.org/genproto v0.0.0-20230227214838-9b19f0bdc514 // indirect; indirect // indirect
	google.golang.org/grpc v1.53.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
