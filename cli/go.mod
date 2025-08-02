module github.com/cloudquery/cloudquery/cli/v6

go 1.24.5

require (
	github.com/apache/arrow-go/v18 v18.4.0
	github.com/bradleyjkemp/cupaloy/v2 v2.8.0
	github.com/cenkalti/backoff/v4 v4.3.0
	github.com/cloudquery/cloudquery-api-go v1.14.1
	github.com/cloudquery/codegen v0.3.31
	github.com/cloudquery/plugin-pb-go v1.26.18
	github.com/cloudquery/plugin-sdk/v4 v4.88.1
	github.com/distribution/reference v0.6.0
	github.com/docker/distribution v2.8.3+incompatible
	github.com/docker/docker v28.3.3+incompatible
	github.com/fatih/color v1.18.0
	github.com/getsentry/sentry-go v0.30.0
	github.com/ghodss/yaml v1.0.0
	github.com/google/go-cmp v0.7.0
	github.com/google/uuid v1.6.0
	github.com/invopop/jsonschema v0.13.0
	github.com/jedib0t/go-pretty/v6 v6.6.5
	github.com/manifoldco/promptui v0.9.0
	github.com/opencontainers/go-digest v1.0.0
	github.com/opencontainers/image-spec v1.1.1
	github.com/pkg/browser v0.0.0-20240102092130-5ac0b6a4141c
	github.com/rs/zerolog v1.34.0
	github.com/rudderlabs/analytics-go/v4 v4.2.1
	github.com/samber/lo v1.49.1
	github.com/santhosh-tekuri/jsonschema/v5 v5.3.1
	github.com/schollz/progressbar/v3 v3.14.6
	github.com/spf13/cobra v1.9.0
	github.com/stretchr/testify v1.10.0
	github.com/thoas/go-funk v0.9.3
	github.com/vnteamopen/godebouncer v1.1.1-0.20230626172639-4b59d27e1b8c
	github.com/wk8/go-ordered-map/v2 v2.1.8
	github.com/yuin/goldmark v1.7.8
	go.opentelemetry.io/collector/component v1.36.1
	go.opentelemetry.io/collector/config/configgrpc v0.130.1
	go.opentelemetry.io/collector/config/confighttp v0.130.1
	go.opentelemetry.io/collector/config/configoptional v0.130.1
	go.opentelemetry.io/collector/config/configtls v1.36.1
	go.opentelemetry.io/collector/consumer v1.36.1
	go.opentelemetry.io/collector/pdata v1.36.1
	go.opentelemetry.io/collector/receiver v1.36.1
	go.opentelemetry.io/collector/receiver/otlpreceiver v0.130.1
	go.opentelemetry.io/otel v1.37.0
	go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp v0.13.0
	go.opentelemetry.io/otel/log v0.13.0
	go.opentelemetry.io/otel/metric v1.37.0
	go.opentelemetry.io/otel/sdk v1.37.0
	go.opentelemetry.io/otel/sdk/log v0.13.0
	go.opentelemetry.io/otel/trace v1.37.0
	go.uber.org/zap v1.27.0
	golang.org/x/exp v0.0.0-20250718183923-645b1fa84792
	golang.org/x/net v0.42.0
	golang.org/x/sync v0.16.0
	golang.org/x/term v0.33.0
	google.golang.org/grpc v1.74.1
	google.golang.org/protobuf v1.36.6
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/Azure/go-ansiterm v0.0.0-20230124172434-306776ec8161 // indirect
	github.com/Masterminds/semver v1.5.0 // indirect
	github.com/Microsoft/go-winio v0.6.1 // indirect
	github.com/adrg/xdg v0.5.3 // indirect
	github.com/apapsch/go-jsonmerge/v2 v2.0.0 // indirect
	github.com/avast/retry-go/v4 v4.6.1 // indirect
	github.com/bahlo/generic-list-go v0.2.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/buger/jsonparser v1.1.1 // indirect
	github.com/cenkalti/backoff/v5 v5.0.2 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/chzyer/readline v0.0.0-20180603132655-2972be24d48e // indirect
	github.com/containerd/errdefs v1.0.0 // indirect
	github.com/containerd/errdefs/pkg v0.3.0 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.6 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/docker/go-connections v0.5.0 // indirect
	github.com/docker/go-metrics v0.0.1 // indirect
	github.com/docker/go-units v0.5.0 // indirect
	github.com/docker/libtrust v0.0.0-20160708172513-aabc10ec26b7 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/foxboron/go-tpm-keyfiles v0.0.0-20250323135004-b31fac66206e // indirect
	github.com/fsnotify/fsnotify v1.9.0 // indirect
	github.com/go-logr/logr v1.4.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-viper/mapstructure/v2 v2.3.0 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/goccy/go-json v0.10.5 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/snappy v1.0.0 // indirect
	github.com/google/flatbuffers v25.2.10+incompatible // indirect
	github.com/google/go-tpm v0.9.5 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.27.1 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.8 // indirect
	github.com/hashicorp/go-version v1.7.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.18.0 // indirect
	github.com/klauspost/cpuid/v2 v2.2.11 // indirect
	github.com/knadh/koanf/maps v0.1.2 // indirect
	github.com/knadh/koanf/providers/confmap v1.0.0 // indirect
	github.com/knadh/koanf/v2 v2.2.1 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.16 // indirect
	github.com/mitchellh/colorstring v0.0.0-20190213212951-d06e56a500db // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/moby/docker-image-spec v1.3.1 // indirect
	github.com/moby/sys/sequential v0.6.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/mostynb/go-grpc-compression v1.2.3 // indirect
	github.com/oapi-codegen/runtime v1.1.2 // indirect
	github.com/pierrec/lz4/v4 v4.1.22 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/prometheus/client_golang v1.19.1 // indirect
	github.com/prometheus/client_model v0.6.1 // indirect
	github.com/prometheus/common v0.54.0 // indirect
	github.com/prometheus/procfs v0.15.0 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/rs/cors v1.11.1 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/santhosh-tekuri/jsonschema/v6 v6.0.2 // indirect
	github.com/segmentio/backo-go v1.1.0 // indirect
	github.com/spf13/pflag v1.0.6 // indirect
	github.com/tidwall/gjson v1.14.4 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.1 // indirect
	github.com/zeebo/xxh3 v1.0.2 // indirect
	go.opentelemetry.io/auto/sdk v1.1.0 // indirect
	go.opentelemetry.io/collector v0.130.1 // indirect
	go.opentelemetry.io/collector/client v1.36.1 // indirect
	go.opentelemetry.io/collector/component/componentstatus v0.130.1 // indirect
	go.opentelemetry.io/collector/config/configauth v0.130.1 // indirect
	go.opentelemetry.io/collector/config/configcompression v1.36.1 // indirect
	go.opentelemetry.io/collector/config/configmiddleware v0.130.1 // indirect
	go.opentelemetry.io/collector/config/confignet v1.36.1 // indirect
	go.opentelemetry.io/collector/config/configopaque v1.36.1 // indirect
	go.opentelemetry.io/collector/confmap v1.36.1 // indirect
	go.opentelemetry.io/collector/consumer/consumererror v0.130.1 // indirect
	go.opentelemetry.io/collector/consumer/xconsumer v0.130.1 // indirect
	go.opentelemetry.io/collector/extension/extensionauth v1.36.1 // indirect
	go.opentelemetry.io/collector/extension/extensionmiddleware v0.130.1 // indirect
	go.opentelemetry.io/collector/featuregate v1.36.1 // indirect
	go.opentelemetry.io/collector/internal/sharedcomponent v0.130.1 // indirect
	go.opentelemetry.io/collector/internal/telemetry v0.130.1 // indirect
	go.opentelemetry.io/collector/pdata/pprofile v0.130.1 // indirect
	go.opentelemetry.io/collector/pipeline v0.130.1 // indirect
	go.opentelemetry.io/collector/receiver/receiverhelper v0.130.1 // indirect
	go.opentelemetry.io/collector/receiver/xreceiver v0.130.1 // indirect
	go.opentelemetry.io/contrib/bridges/otelzap v0.12.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.62.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.62.0 // indirect
	go.opentelemetry.io/proto/otlp v1.7.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.yaml.in/yaml/v3 v3.0.4 // indirect
	golang.org/x/crypto v0.40.0 // indirect
	golang.org/x/mod v0.26.0 // indirect
	golang.org/x/sys v0.34.0 // indirect
	golang.org/x/text v0.27.0 // indirect
	golang.org/x/tools v0.35.0 // indirect
	golang.org/x/xerrors v0.0.0-20240903120638-7835f813f4da // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20250603155806-513f23925822 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250603155806-513f23925822 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

// github.com/cloudquery/jsonschema @ cqmain
replace github.com/invopop/jsonschema => github.com/cloudquery/jsonschema v0.0.0-20240220124159-92878faa2a66

// github.com/cloudquery/godebouncer @ fix-race
replace github.com/vnteamopen/godebouncer => github.com/cloudquery/godebouncer v0.0.0-20230626172639-4b59d27e1b8c
