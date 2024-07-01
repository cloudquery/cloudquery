module github.com/{{.Org}}/cq-source-{{.Name}}

go 1.22.3

require (
	github.com/apache/arrow/go/v16 v16.1.0
	github.com/cloudquery/plugin-pb-go v1.20.2
	github.com/cloudquery/plugin-sdk/v4 v4.49.0
	github.com/rs/zerolog v1.29.0
)

require (
	google.golang.org/genproto/googleapis/rpc 1f4bbc51befe // indirect
)
