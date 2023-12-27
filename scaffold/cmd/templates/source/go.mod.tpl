module github.com/{{.Org}}/cq-source-{{.Name}}

go 1.21

require (
	github.com/apache/arrow/go/v15 v15.0.0-20231219235838-1c48d69844cb
	github.com/cloudquery/plugin-pb-go v1.14.4
	github.com/cloudquery/plugin-sdk/v4 v4.21.3
	github.com/rs/zerolog v1.29.0
)

require (
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230530153820-e85fd2cbaebc // indirect
)
