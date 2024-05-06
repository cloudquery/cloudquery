module github.com/{{.Org}}/cq-source-{{.Name}}

go 1.21.4

require (
	github.com/apache/arrow/go/v15 bcaeaa8c2d97
	github.com/cloudquery/plugin-pb-go v1.19.11
	github.com/cloudquery/plugin-sdk/v4 v4.40.1
	github.com/rs/zerolog v1.29.0
)

require (
	google.golang.org/genproto/googleapis/rpc 1f4bbc51befe // indirect
)
