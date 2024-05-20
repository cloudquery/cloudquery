module github.com/{{.Org}}/cq-source-{{.Name}}

go 1.21.4

require (
	github.com/apache/arrow/go/v16 v16.1.0
	github.com/cloudquery/plugin-pb-go v1.19.15
	github.com/cloudquery/plugin-sdk/v4 v4.43.1
	github.com/rs/zerolog v1.29.0
)

require (
	google.golang.org/genproto/googleapis/rpc 1f4bbc51befe // indirect
)
