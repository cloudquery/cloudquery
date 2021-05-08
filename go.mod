module github.com/cloudquery/cloudquery

go 1.15

require (
	github.com/VividCortex/ewma v1.2.0 // indirect
	github.com/briandowns/spinner v1.12.0 // indirect
	github.com/cenkalti/backoff/v3 v3.2.2 // indirect
	github.com/cloudquery/cq-provider-sdk v0.1.7
	github.com/containerd/continuity v0.0.0-20201208142359-180525291bb7 // indirect
	github.com/fatih/color v1.10.0
	github.com/google/go-github/v35 v35.1.0
	github.com/hashicorp/go-hclog v0.16.0
	github.com/hashicorp/go-plugin v1.4.1
	github.com/hashicorp/go-version v1.3.0
	github.com/hashicorp/hcl v1.0.0
	github.com/hashicorp/hcl/v2 v2.10.0
	github.com/jackc/pgx/v4 v4.10.1
	github.com/leaanthony/spinner v0.5.3 // indirect
	github.com/mitchellh/go-glint v0.0.0-20201119015200-53f6eb3bf4d2 // indirect
	github.com/moby/term v0.0.0-20201216013528-df9cb8a40635 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/rs/zerolog v1.20.0
	github.com/shopspring/decimal v1.2.0 // indirect
	github.com/sirupsen/logrus v1.7.0 // indirect
	github.com/spf13/afero v1.1.2
	github.com/spf13/cobra v1.1.1
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.7.0
	github.com/vbauerster/mpb/v6 v6.0.3
	github.com/zclconf/go-cty v1.8.0
	golang.org/x/crypto v0.0.0-20210322153248-0c34fe9e7dc2
	golang.org/x/sync v0.0.0-20200625203802-6e8e738ad208
	google.golang.org/genproto v0.0.0-20210202153253-cf70463f6119 // indirect
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.1.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

replace github.com/cloudquery/cq-provider-sdk v0.1.7 => ../cq-provider-sdk
