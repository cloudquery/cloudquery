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
				DSN:  ptr(`host=localhost database=somedb port=5432 sslmode=disable`),
				Type: ptr("tsdb"),
			},
			"",
			true,
		},
		{
			&Connection{
				Username: ptr(`user`),
				Password: ptr(`pass`),
				Type:     ptr(`tsdb`),
				Host:     ptr(`localhost`),
				Database: ptr(`postgres`),
			},
			"tsdb://user:pass@localhost:5432/postgres",
			false,
		},
		{
			&Connection{
				Username: ptr(`user`),
				Type:     ptr(`tsdb`),
				Host:     ptr(`localhost`),
				Port:     func(i int) *int { return &i }(15432),
				Database: ptr(`postgres`),
			},
			"tsdb://user@localhost:15432/postgres",
			false,
		},
		{
			&Connection{
				Username: ptr(`user`),
				Password: ptr(`pass`),
				Host:     ptr(`localhost`),
				Database: ptr(`postdb`),
				SSLMode:  ptr(`disable`),
				Extras:   func(s []string) *[]string { return &s }([]string{"a=b", "c=d", "e", "sslmode=enable"}),
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
				assert.NotNil(t, tc.input.DSN)
				assert.Equal(t, tc.expectedResult, *tc.input.DSN)
			}
		})
	}
}
