package config

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleConnectionBlock(t *testing.T) {
	s := func(s string) *string {
		return &s
	}

	cases := []struct {
		input          *Connection
		expectedResult string
		expectedError  bool
	}{
		{
			&Connection{
				DSN:  `host=localhost database=somedb port=5432 sslmode=disable`,
				Type: s(`tsdb`),
			},
			"tsdb://localhost:5432/somedb?sslmode=disable",
			false,
		},
		{
			&Connection{
				DSN:      `host=localhost database=somedb port=5432 sslmode=disable`,
				Username: s(`user`),
			},
			"",
			true,
		},
		{
			&Connection{
				DSN:      `host=localhost database=somedb port=5432 sslmode=disable`,
				Password: s(`pass`),
			},
			"",
			true,
		},
		{
			&Connection{
				DSN:      `host=localhost database=somedb port=5432 sslmode=disable`,
				Username: s(`user`),
				Password: s(`pass`),
				Type:     s(`tsdb`),
			},
			"",
			true,
		},
		{
			&Connection{
				DSN:  `host=localhost user=postgres password=pass database=somedb port=5432 sslmode=disable`,
				Type: s(`tsdb`),
			},
			"tsdb://postgres:pass@localhost:5432/somedb?sslmode=disable",
			false,
		},
		{
			&Connection{
				DSN:      `postgres://user:pass@localhost:5432/somedb?sslmode=disable`,
				Username: s(`us3r`),
				Password: s(`p4ss`),
			},
			"",
			true,
		},
		{
			&Connection{
				DSN:      `postgres://localhost:5432/somedb?sslmode=disable`,
				Username: s(`us3r`),
				Password: s(`p4ss`),
			},
			"postgres://us3r:p4ss@localhost:5432/somedb?sslmode=disable",
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
