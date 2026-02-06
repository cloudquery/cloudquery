module github.com/{{.Org}}/cq-source-{{.Name}}

go 1.25.6

require (
	github.com/apache/arrow-go/v18 v18.5.1
	github.com/cloudquery/plugin-pb-go v1.27.6
	github.com/cloudquery/plugin-sdk/v4 v4.94.2
	github.com/rs/zerolog v1.34.0
)

require (
	google.golang.org/genproto/googleapis/rpc 1f4bbc51befe // indirect
)
