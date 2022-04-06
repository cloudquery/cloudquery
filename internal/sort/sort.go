package sort

import stdsort "sort"

// Unique sorts input alphabetically and returns only non-duplicate entries
func Unique(input []string) []string {
	if input == nil {
		return nil
	}

	stdsort.Strings(input)
	ret := make([]string, 0, len(input))
	for i := range input {
		if i == len(input)-1 { // always append last element
			ret = append(ret, input[i])
			continue
		}

		if input[i] == input[i+1] { // skip if it's the same as the next one
			continue
		}

		ret = append(ret, input[i])
	}

	return ret
}
