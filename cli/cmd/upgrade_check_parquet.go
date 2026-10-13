package cmd

import (
	"fmt"
	"slices"
	"strings"
)

const upgradeMaxStructFields = 3

type upgradeParquetNode struct {
	repetition string
	physical   string
	logical    string
	fields     []upgradeParquetField
}

type upgradeParquetField struct {
	name string
	node upgradeParquetNode
}

func upgradeParquetTypeName(parquetType string) (string, bool) {
	rest, node, ok := parseUpgradeParquetNode(parquetType)
	if !ok || rest != "" {
		return "", false
	}
	name, ok := node.shortName()
	if !ok {
		return "", false
	}
	if node.repetition == "required" {
		name = "required " + name
	}
	return name, true
}

func parseUpgradeParquetNode(s string) (string, upgradeParquetNode, bool) {
	var node upgradeParquetNode
	var ok bool
	if node.repetition, s, ok = strings.Cut(s, " "); !ok || !slices.Contains([]string{"required", "optional", "repeated"}, node.repetition) {
		return "", node, false
	}
	end := strings.IndexAny(s, " ;}")
	if end < 0 {
		end = len(s)
	}
	node.physical, s = s[:end], s[end:]
	if strings.HasPrefix(s, " (") {
		closing := upgradeClosingParen(s[1:])
		if closing < 0 {
			return "", node, false
		}
		node.logical, s = s[2:closing+1], s[closing+2:]
	}
	if node.physical != "group" {
		return s, node, true
	}
	if s, ok = strings.CutPrefix(s, " {"); !ok {
		return "", node, false
	}
	for {
		var field upgradeParquetField
		if field.name, s, ok = strings.Cut(s, ": "); !ok {
			return "", node, false
		}
		if s, field.node, ok = parseUpgradeParquetNode(s); !ok {
			return "", node, false
		}
		node.fields = append(node.fields, field)
		if rest, ok := strings.CutPrefix(s, "}"); ok {
			return rest, node, true
		}
		if s, ok = strings.CutPrefix(s, "; "); !ok {
			return "", node, false
		}
	}
}

func upgradeClosingParen(s string) int {
	depth := 0
	for i, r := range s {
		switch r {
		case '(':
			depth++
		case ')':
			if depth--; depth == 0 {
				return i
			}
		}
	}
	return -1
}

func (n upgradeParquetNode) shortName() (string, bool) {
	if n.physical == "group" {
		return n.groupName()
	}
	name, params, _ := strings.Cut(strings.TrimSuffix(n.logical, ")"), "(")
	values := upgradeParquetParams(params)
	switch name {
	case "", "None":
		return n.physicalName(), true
	case "String", "Enum":
		return "string", true
	case "JSON":
		return "json", true
	case "BSON":
		return "bson", true
	case "UUID":
		return "uuid", true
	case "Date":
		return "date", true
	case "Float16":
		return "float16", true
	case "Interval":
		return "interval", true
	case "Int":
		prefix := "int"
		if values["isSigned"] == "false" {
			prefix = "uint"
		}
		return prefix + values["bitWidth"], values["bitWidth"] != ""
	case "Decimal":
		return fmt.Sprintf("decimal(%s,%s)", values["precision"], values["scale"]), values["precision"] != ""
	case "Timestamp", "Time":
		return upgradeParquetTimeName(strings.ToLower(name), values)
	}
	return "", false
}

func (n upgradeParquetNode) physicalName() string {
	switch {
	case n.physical == "boolean":
		return "bool"
	case n.physical == "byte_array":
		return "binary"
	case strings.HasPrefix(n.physical, "fixed_len_byte_array"):
		return "binary" + strings.TrimPrefix(n.physical, "fixed_len_byte_array")
	}
	return n.physical
}

func upgradeParquetParams(params string) map[string]string {
	values := make(map[string]string)
	for _, param := range strings.Split(params, ", ") {
		if key, value, ok := strings.Cut(param, "="); ok {
			values[key] = value
		}
	}
	return values
}

func upgradeParquetTimeName(name string, values map[string]string) (string, bool) {
	unit, ok := map[string]string{"milliseconds": "ms", "microseconds": "us", "nanoseconds": "ns"}[values["timeUnit"]]
	if !ok {
		return "", false
	}
	if values["isAdjustedToUTC"] == "true" {
		return fmt.Sprintf("%s (%s, UTC)", name, unit), true
	}
	return fmt.Sprintf("%s (%s)", name, unit), true
}

func (n upgradeParquetNode) groupName() (string, bool) {
	switch n.logical {
	case "List":
		element, ok := n.repeatedChild(1)
		if !ok {
			return "", false
		}
		elementName, ok := element[0].shortName()
		return "list<" + elementName + ">", ok
	case "Map":
		keyValue, ok := n.repeatedChild(2)
		if !ok {
			return "", false
		}
		keyName, keyOK := keyValue[0].shortName()
		valueName, valueOK := keyValue[1].shortName()
		return "map<" + keyName + "," + valueName + ">", keyOK && valueOK
	case "":
		return n.structName()
	}
	return "", false
}

func (n upgradeParquetNode) repeatedChild(fieldCount int) ([]upgradeParquetNode, bool) {
	if len(n.fields) != 1 || n.fields[0].node.repetition != "repeated" {
		return nil, false
	}
	child := n.fields[0].node
	if child.physical != "group" {
		child.repetition = "optional"
		return []upgradeParquetNode{child}, fieldCount == 1
	}
	if len(child.fields) != fieldCount {
		return nil, false
	}
	nodes := make([]upgradeParquetNode, fieldCount)
	for i, field := range child.fields {
		nodes[i] = field.node
	}
	return nodes, true
}

func (n upgradeParquetNode) structName() (string, bool) {
	var fields []string
	for _, field := range n.fields[:min(len(n.fields), upgradeMaxStructFields)] {
		fieldName, ok := field.node.shortName()
		if !ok {
			return "", false
		}
		fields = append(fields, field.name+": "+fieldName)
	}
	if len(n.fields) > upgradeMaxStructFields {
		fields = append(fields, "…")
	}
	return "struct<" + strings.Join(fields, ", ") + ">", true
}
