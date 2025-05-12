module github.com/{{.Org}}/cq-source-{{.Name}}

go 1.24.2

require (
	github.com/apache/arrow-go/v18 v18.2.0
	github.com/cloudquery/plugin-pb-go v1.26.11
	github.com/cloudquery/plugin-sdk/v4 v4.80.1
	github.com/rs/zerolog v1.33.0
)

require (
	google.golang.org/genproto/googleapis/rpc 1f4bbc51befe // indirect
)
