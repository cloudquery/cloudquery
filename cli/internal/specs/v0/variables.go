package specs

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/thoas/go-funk"
)

type PluginVariables struct {
	Connection string `json:"connection"`
}

type Variables struct {
	Plugins map[string]PluginVariables `json:"plugins"`
}

var reVariables = regexp.MustCompile(`@@(plugins\.[a-zA-Z0-9_\.-]+)`)

// ReplaceVariables replaces variables starting with @@ in the src string
// with the values from the values from variables by dot notation.
// Example: @@plugins.aws.connection will be replaced with the value of variables.Plugins["aws"].Connection
func ReplaceVariables(src string, variables Variables, registry Registry) (string, error) {
	var lastErr error
	bytes, err := json.Marshal(variables)
	if err != nil {
		return "", err
	}
	variablesMap := make(map[string]any)
	err = json.Unmarshal(bytes, &variablesMap)
	if err != nil {
		return "", err
	}
	result := reVariables.ReplaceAllStringFunc(src, func(s string) string {
		variablePath := s[2:]
		res := funk.Get(variablesMap, variablePath)
		if res == nil {
			lastErr = fmt.Errorf("variable %s not found", variablePath)
			return s
		}
		resString, ok := res.(string)
		if !ok {
			lastErr = fmt.Errorf("variable %s is not a string", variablePath)
			return s
		}
		// Edge case: if the plugin whose spec's variables are being replaced is a docker plugin,
		// it won't be able to connect to localhost, so we replace localhost with host.docker.internal
		if strings.Contains(variablePath, ".connection") && registry == RegistryDocker {
			for _, needle := range []string{"localhost", "0.0.0.0", "127.0.0.1"} {
				resString = strings.ReplaceAll(resString, needle, "host.docker.internal")
			}
		}
		// make safe for replacement into JSON string
		v, err := json.Marshal(resString)
		if err != nil {
			lastErr = err
			return s
		}
		resString = string(v[1 : len(v)-1])
		return resString
	})
	return result, lastErr
}
