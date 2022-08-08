package internal

// FlattenRow flattens a row map and skips empty columns
func FlattenRow(m map[string]interface{}) map[string]interface{} {
	o := make(map[string]interface{})
	for k, v := range m {
		switch child := v.(type) {
		case map[string]interface{}:
			nm := FlattenRow(child)
			for nk, nv := range nm {
				o[nk] = nv
			}
		default:
			if v == nil {
				continue
			}
			o[k] = v
		}
	}
	return o
}
