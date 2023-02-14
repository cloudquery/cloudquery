package client

import (
	"strconv"
	"strings"
)

func unescape(str string) string {
	out := ""
	for len(str) > 0 {
		if str[0] == '\\' {
			if len(str) > 5 && str[1] == 'u' {
				u, err := strconv.ParseUint(str[2:6], 16, 64)
				if err == nil {
					out += string(byte(u))
					str = str[6:]
					continue
				}
			}
		}
		out += str[:1]
		str = str[1:]
	}

	// check for unquote, too
	if !strings.Contains(out, `\"`) && !strings.Contains(out, `\n`) {
		return out
	}

	if unquoted, err := strconv.Unquote(`"` + out + `"`); err == nil {
		return unquoted
	}

	return out
}
