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
