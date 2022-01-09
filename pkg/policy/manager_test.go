package policy

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/hashicorp/go-hclog"

	"github.com/stretchr/testify/assert"
)

func TestManager_Load(t *testing.T) {
	// Skip test for now since github is annoying
	cases := []struct {
		Name           string
		Policy         *Policy
		ExpectedPolicy *Policy
		ErrorOutput    string
	}{
		{
			Name: "load github policy",
			Policy: &Policy{
				Name:   "test",
				Source: "github.com/cloudquery-policies/test_policy",
			},
			ExpectedPolicy: &Policy{
				Name:  "test",
				Title: "this is a test policy",
				Doc:   "MAIN README",
				Policies: Policies{
					{
						Name:  "sub-policy",
						Title: "sub policy description",
						Doc:   "README FOR SUBPOLICY",
						Checks: []*Check{
							{
								Name:         "check",
								Title:        "test check",
								Doc:          "some doc md",
								ExpectOutput: true,
								Type:         AutomaticQuery,
								Query:        "SELECT 1;",
							},
						},
					},
				},
				Source: "github.com/cloudquery-policies/test_policy",
				meta: &Meta{
					Type:      "github",
					Version:   "",
					subPolicy: "",
					Directory: "cq/policies/manager/github.com/cloudquery-policies/test_policy",
				},
			},
		},

		{
			Name: "load github policy",
			Policy: &Policy{
				Name:   "test",
				Source: "tests/empty",
			},
			ExpectedPolicy: &Policy{
				Name:  "test",
				Title: "this is a test policy",
				Doc:   "MAIN README",
				Policies: Policies{
					{
						Name:  "sub-policy",
						Title: "sub policy description",
						Doc:   "README FOR SUBPOLICY",
						Checks: []*Check{
							{
								Name:         "check",
								Title:        "test check",
								Doc:          "some doc md",
								ExpectOutput: true,
								Type:         AutomaticQuery,
								Query:        "SELECT 1;",
							},
						},
					},
				},
				Source: "github.com/cloudquery-policies/test_policy",
				meta: &Meta{
					Type:      "github",
					Version:   "",
					subPolicy: "",
					Directory: "cq/policies/manager/github.com/cloudquery-policies/test_policy",
				},
			},
		},
	}

	_ = os.RemoveAll("./test")
	m := NewManager("./test", nil, hclog.Default())
	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			p, err := m.Load(context.Background(), tc.Policy)
			assert.NoError(t, err)
			assert.NotNil(t, p)
			if p.meta != nil {
				p.meta.Directory = filepath.ToSlash(p.meta.Directory)
			}
			assert.Equal(t, tc.ExpectedPolicy, p)
		})
	}

}
