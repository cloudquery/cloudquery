package specs

import "testing"

func TestReplaceVariables(t *testing.T) {
	variables := Variables{
		Plugins: map[string]PluginVariables{
			"aws": {
				Connection: "test",
			},
		},
	}
	res, err := ReplaceVariables("test", variables)
	if err != nil {
		t.Fatal(err)
	}
	if res != "test" {
		t.Fatalf("expected test, got %s", res)
	}

	_, err = ReplaceVariables("@@plugins", variables)
	if err == nil {
		t.Fatal("expected error")
	}

	_, err = ReplaceVariables("@@plugins.aws", variables)
	if err == nil {
		t.Fatal("expected error")
	}

	_, err = ReplaceVariables("@@plugins.aws.test", variables)
	if err == nil {
		t.Fatal("expected error")
	}

	res, err = ReplaceVariables("@@plugins.aws.connection", variables)
	if err != nil {
		t.Fatal(err)
	}
	if res != "test" {
		t.Fatalf("expected %q, got %q", "test", res)
	}

	res, err = ReplaceVariables("inside @@plugins.aws.connection string", variables)
	if err != nil {
		t.Fatal(err)
	}
	if res != "inside test string" {
		t.Fatalf("expected %q, got %q", "inside test string", res)
	}
}
