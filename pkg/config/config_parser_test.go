package config

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleConnectionBlock(t *testing.T) {
	cases := []struct {
		input          *Connection
		expectedResult string
		expectedError  bool
	}{
		{
			&Connection{
				Username: `user`,
				Password: `pass`,
				Host:     `localhost`,
				Database: `postgres`,
			},
			"postgres://user:pass@localhost:5432/postgres",
			false,
		},
		{
			&Connection{
				Username: `user`,
				Type:     `postgres`,
				Host:     `localhost`,
				Port:     15432,
				Database: `postgres`,
			},
			"postgres://user@localhost:15432/postgres",
			false,
		},
		{
			&Connection{
				Username: `user`,
				Password: `pass`,
				Host:     `localhost`,
				Database: `postdb`,
				SSLMode:  `disable`,
				Extras:   []string{"a=b", "c=d", "e", "sslmode=enable"},
			},
			"postgres://user:pass@localhost:5432/postdb?a=b&c=d&e=&sslmode=disable",
			false,
		},
	}
	for i := range cases {
		tc := cases[i]
		t.Run("case #"+strconv.Itoa(i+1), func(t *testing.T) {
			err := handleConnectionBlock(tc.input)
			if tc.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedResult, tc.input.DSN)
			}
		})
	}
}

func TestParser_LoadConfigFromSourceConnectionOptionality(t *testing.T) {
	cases := []struct {
		cfg           string
		expectedDSN   string
		expectedError bool
	}{
		{
			`
cloudquery {
  connection {
    dsn =  "postgres://postgres:pass@localhost:5432/postgres"
  }
}
`,
			"postgres://postgres:pass@localhost:5432/postgres",
			false,
		},
		{
			`
cloudquery {
  connection {
    dsn =  "postgres://postgres:pass@localhost:5432/postgres"
    database = "cq"
  }
}
`,
			"",
			true,
		},
		{
			`
cloudquery {
  connection {
    username = "postgres"
    password = "pass"
    host = "localhost"
    port = 15432
    database = "cq"
    sslmode = "disable"
  }
}
`,
			"postgres://postgres:pass@localhost:15432/cq?sslmode=disable",
			false,
		},
		{
			`
cloudquery {
  connection {
    username = "postgres"
    password = "pass"
    type = "postgres"
    host = "localhost"
    port = 15432
    database = "cq"
    sslmode = "disable"
	extras = [ "search_path=myschema" ]
  }
}
`,
			"postgres://postgres:pass@localhost:15432/cq?search_path=myschema&sslmode=disable",
			false,
		},
	}
	for i := range cases {
		tc := cases[i]
		t.Run("case #"+strconv.Itoa(i+1), func(t *testing.T) {
			p := NewParser()
			parsedCfg, diags := p.LoadConfigFromSource("test.hcl", []byte(tc.cfg))
			if tc.expectedError {
				assert.True(t, diags.HasErrors())
			} else {
				assert.Len(t, diags.Errs(), 0)
				assert.Equal(t, tc.expectedDSN, parsedCfg.CloudQuery.Connection.DSN)
			}
		})
	}
}
