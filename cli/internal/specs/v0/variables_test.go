package specs

import "testing"

func TestReplaceVariables(t *testing.T) {
	cases := []struct {
		src                    string
		connection             string
		shouldReplaceLocalhost bool
		expect                 string
		expectErr              bool
	}{
		{
			src:        "test",
			connection: "localhost:7777",
			expect:     "test",
		},
		{
			src:        "@@something",
			connection: "localhost:7777",
			expect:     "@@something",
		},
		{
			src:        "@@plugins.aws",
			connection: "localhost:7777",
			expectErr:  true,
		},
		{
			src:        "@@plugins.aws.connection",
			connection: "localhost:7777",
			expect:     "localhost:7777",
		},
		{
			src:        "inside @@plugins.aws.connection string multiple times @@plugins.aws.connection",
			connection: "localhost:7777",
			expect:     "inside localhost:7777 string multiple times localhost:7777",
		},
		{
			src:                    "nothing to replace",
			connection:             "localhost:7777",
			shouldReplaceLocalhost: true,
			expect:                 "nothing to replace",
		},
		{
			src:                    "@@plugins.aws.connection",
			connection:             "localhost:7777",
			shouldReplaceLocalhost: true,
			expect:                 "host.docker.internal:7777",
		},
		{
			src:                    "@@plugins.aws.connection",
			connection:             "127.0.0.1:7777",
			shouldReplaceLocalhost: true,
			expect:                 "host.docker.internal:7777",
		},
		{
			src:                    "@@plugins.aws.connection",
			connection:             "0.0.0.0:7777",
			shouldReplaceLocalhost: true,
			expect:                 "host.docker.internal:7777",
		},
	}
	for _, c := range cases {
		vars := Variables{
			Plugins: map[string]PluginVariables{
				"aws": {
					Connection: c.connection,
				},
			},
		}
		res, err := ReplaceVariables(c.src, vars, c.shouldReplaceLocalhost)
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
