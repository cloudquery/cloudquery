package drift

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/stretchr/testify/assert"
)

func TestReadBuiltinConfig(t *testing.T) {
	t.Parallel()

	d := &Drift{
		logger: hclog.NewNullLogger(),
	}
	val, err := d.readBuiltinConfig()
	assert.NoError(t, err)
	assert.Nil(t, val.Terraform)
	assert.NotNil(t, val.WildProvider)
	assert.Equal(t, 1, len(val.Providers))
	assert.NotNil(t, val.Providers[0].WildResource)
	assert.NotZero(t, len(val.Providers[0].Resources))
}

func TestEmptyProfileConfig(t *testing.T) {
	t.Parallel()

	d := &Drift{
		logger: hclog.NewNullLogger(),
	}
	base, err := d.readBuiltinConfig()
	assert.NoError(t, err)
	assert.NotNil(t, base)

	bc, err := d.readProfileConfig(base, nil)
	assert.NoError(t, err)
	assert.NotNil(t, bc)
}

func TestProfileConfig(t *testing.T) {
	t.Parallel()

	d := &Drift{
		logger: hclog.NewNullLogger(),
	}
	base, err := d.readBuiltinConfig()
	assert.NoError(t, err)
	assert.NotNil(t, base)

	configRaw, diags := hclparse.NewParser().ParseHCL([]byte(`
provider "aws" {
	resource "ec2.instances" {
		identifiers = [ "test" ]
		deep = true
	}
}
`), "")
	assert.False(t, diags.HasErrors(), "%s", diags.Error())

	cfg, err := d.readProfileConfig(base, configRaw.Body)
	assert.NoError(t, err)
	assert.NotNil(t, cfg)

	a := cfg.FindProvider("aws")
	assert.NotNil(t, a)

	r := a.Resources["ec2.instances"]
	assert.NotNil(t, r)

	assert.Equal(t, []string{"test"}, r.Identifiers)
	assert.EqualValues(t, aws.Bool(true), r.Deep)
}

func TestHandleIdentifiers(t *testing.T) {
	t.Parallel()

	table := []struct {
		Name          string
		Identifiers   []string
		ExpectedExp   exp.Expression
		ExpectedError bool
	}{
		{
			Name:        "Single",
			Identifiers: []string{"id"},
			ExpectedExp: goqu.L(`c."id" AS id`),
		},
		{
			Name:        "Multiple",
			Identifiers: []string{"id1", "id2"},
			ExpectedExp: goqu.L(fmt.Sprintf(`CONCAT(c."id1",'%s',c."id2") AS id`, idSeparator)),
		},
		{
			Name:        "Multiple with refs",
			Identifiers: []string{"parent.id", "c.id"},
			ExpectedExp: goqu.L(fmt.Sprintf(`CONCAT(parent.id,'%s',c.id) AS id`, idSeparator)),
		},
		{
			Name:        "Single with SQL",
			Identifiers: []string{`${sql:LEFT(id, 5)}`},
			ExpectedExp: goqu.L(`LEFT(id, 5) AS id`),
		},
		{
			Name:        "Multiple with SQL",
			Identifiers: []string{"id1", `${sql:LEFT(id2, 5)}`},
			ExpectedExp: goqu.L(fmt.Sprintf(`CONCAT(c."id1",'%s',LEFT(id2, 5)) AS id`, idSeparator)),
		},
		{
			Name:        "Multiple with refs and SQL",
			Identifiers: []string{"parent.id", `${sql:LEFT(id2, 5)}`},
			ExpectedExp: goqu.L(fmt.Sprintf(`CONCAT(parent.id,'%s',LEFT(id2, 5)) AS id`, idSeparator)),
		},
	}

	for i := range table {
		t.Run(table[i].Name, func(t *testing.T) {
			out, err := handleIdentifiers(table[i].Identifiers)
			if table[i].ExpectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.EqualValues(t, table[i].ExpectedExp, out)
		})
	}
}
