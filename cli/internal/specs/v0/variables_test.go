package specs

import "testing"

func TestReplaceVariables(t *testing.T) {
	vars := Variables{
		Plugins: map[string]PluginVariables{
			"aws": {
				Connection: "test",
			},
		},
	}
	cases := []struct {
		src       string
		variables Variables
		expect    string
		expectErr bool
	}{
		{
			src:       "test",
			variables: vars,
			expect:    "test",
		},
		{
			src:       "@@something",
			variables: vars,
			expect:    "@@something",
		},
		{
			src:       "@@plugins.aws",
			variables: vars,
			expectErr: true,
		},
		{
			src:       "@@plugins.aws.connection",
			variables: vars,
			expect:    "test",
		},
		{
			src:       "inside @@plugins.aws.connection string multiple times @@plugins.aws.connection",
			variables: vars,
			expect:    "inside test string multiple times test",
		},
	}
	for _, c := range cases {
		res, err := ReplaceVariables(c.src, c.variables, RegistryLocal)
		if err != nil && !c.expectErr {
			t.Fatalf("ReplaceVariables(%q) got unexpected error: %v", c.src, err)
		} else if err == nil && c.expectErr {
			t.Fatalf("ReplaceVariables(%q) expected error, got nil", c.src)
		}
		if c.expectErr {
			continue
		}
		if res != c.expect {
			t.Fatalf("ReplaceVariables(%q) = %q, want %q", c.src, res, c.expect)
		}
	}
}
