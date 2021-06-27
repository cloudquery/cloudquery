module github.com/cloudquery/cloudquery

go 1.16

require (
	github.com/VividCortex/ewma v1.2.0 // indirect
	github.com/aws/aws-lambda-go v1.23.0
	github.com/cloudquery/cq-provider-sdk v0.2.7
	github.com/fatih/color v1.10.0
	github.com/fsnotify/fsnotify v1.4.9
	github.com/go-git/go-git/v5 v5.4.2
	github.com/google/go-github/v35 v35.1.0
	github.com/hashicorp/go-hclog v0.16.1
	github.com/hashicorp/go-plugin v1.4.1
	github.com/hashicorp/go-version v1.3.0
	github.com/hashicorp/hcl/v2 v2.10.0
	github.com/jackc/pgx/v4 v4.11.0
	github.com/mattn/go-isatty v0.0.12
	github.com/mitchellh/mapstructure v1.4.1 // indirect
	github.com/rs/zerolog v1.20.0
	github.com/satori/go.uuid v1.2.0
	github.com/spf13/afero v1.1.2
	github.com/spf13/cobra v1.1.3
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.7.0
	github.com/vbauerster/mpb/v6 v6.0.3
	github.com/zclconf/go-cty v1.8.3
	golang.org/x/crypto v0.0.0-20210506145944-38f3c27a63bf
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	golang.org/x/term v0.0.0-20201126162022-7de9c90e9dd1
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

// TODO: remove after approval
replace github.com/cloudquery/cq-provider-sdk v0.2.7 => ../forks/cq-provider-sdk
