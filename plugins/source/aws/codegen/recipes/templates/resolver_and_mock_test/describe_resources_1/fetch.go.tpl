// Code generated by codegen; DO NOT EDIT.

package {{.Service}}

import (
	"context"

    "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/{{.Service}}"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func {{.Table.Resolver}}(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var input {{.Service}}.{{.DescribeMethod.Method.Name}}Input = {{ if .CustomDescribeInput }}{{.CustomDescribeInput}}{{ else }}{{.Service}}.{{.DescribeMethod.Method.Name}}Input{}{{ end }}
	c := meta.(*client.Client)
	svc := c.Services().{{.CloudQueryServiceName}}
	for {
        response, err := svc.{{.DescribeMethod.Method.Name}}(ctx, &input)
        if err != nil {
            return err
        }
        {{- if .DescribeMethod.OutputFieldName }}
        res <- response.{{.DescribeMethod.OutputFieldName}}
        {{- else }}
        res <- response
        {{- end }}

        if aws.ToString(response.NextToken) == "" {
            break
        }
        input.NextToken = response.NextToken
    }
    return nil
}
