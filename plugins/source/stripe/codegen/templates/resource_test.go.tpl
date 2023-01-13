package {{.Service}}_test

import (
  "testing"

  "github.com/cloudquery/cloudquery/plugins/source/stripe/client"
  "github.com/cloudquery/cloudquery/plugins/source/stripe/resources/services/{{.Service}}"
)

func Test{{.TableName | ToPascal}}(t *testing.T) {
  client.MockTestHelper(t, {{.Service}}.{{.TableName | ToPascal}}(), client.TestOptions{})
}
