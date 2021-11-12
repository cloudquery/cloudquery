package drift

import (
	"os"
	"testing"

	"github.com/cloudquery/cloudquery/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestEnvVars(t *testing.T) {
	const (
		varName = "TEST_VARIABLE"
		val     = "VALUE"
	)
	_ = os.Setenv(config.EnvVarPrefix+varName, val)

	p := NewParser("")

	varVal, ok := p.HCLContext.Variables[varName]
	assert.True(t, ok)
	assert.Equal(t, val, varVal.AsString())
}
