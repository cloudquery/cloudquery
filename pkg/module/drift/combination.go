package drift

// MatrixProduct combines each item in `multiplier` with each element in `base`
//
// [(A),(B),(C)] x (1,2) = [(A,1)], [(B,1)], [(C,1)], [(A,2)], [(B,2)], [(C,2)]
func MatrixProduct(base [][]string, multiplier []string) [][]string {
	if len(multiplier) == 0 {
		return base
	}

	if len(base) == 0 {
		ret := make([][]string, len(multiplier))
		for i, v := range multiplier {
			ret[i] = []string{v}
		}
		return ret
	}

	ret := make([][]string, 0, len(base)*len(multiplier))
	for _, v := range multiplier {
		for _, l := range base {
			ll := make([]string, len(l))
			copy(ll, l)

			ret = append(ret, append(ll, v))
		}
	}

	return ret
}
