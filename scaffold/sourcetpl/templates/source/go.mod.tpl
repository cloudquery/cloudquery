module github.com/{{.Org}}/cq-source-{{.Name}}

go 1.22.4

require (
	github.com/apache/arrow-go/v18 v18.0.0
	github.com/cloudquery/plugin-pb-go v1.26.2
	github.com/cloudquery/plugin-sdk/v4 v4.72.4
	github.com/rs/zerolog v1.29.0
)

require (
	google.golang.org/genproto/googleapis/rpc 1f4bbc51befe // indirect
)
