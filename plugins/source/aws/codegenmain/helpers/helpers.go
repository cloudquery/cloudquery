package helpers

func StringSwitch(b bool, ifTrue, ifFalse string) string {
	if b {
		return ifTrue
	}
	return ifFalse
}

func Coalesce(input, defValue string) string {
	if input == "" {
		return defValue
	}
	return input
}

func ReverseStringSlice(input []string) []string {
	ret := make([]string, 0, len(input))
	for i := len(input) - 1; i >= 0; i-- {
		ret = append(ret, input[i])
	}
	return ret
}
