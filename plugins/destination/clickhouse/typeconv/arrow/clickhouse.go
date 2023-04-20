package arrow

import (
	"fmt"
	"strings"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2/lib/column"
)

// From clickhouse-go
func params(t column.Type) string {
	switch start, end := strings.Index(string(t), "("), strings.LastIndex(string(t), ")"); {
	case len(t) == 0, start <= 0, end <= 0, end < start:
		return ""
	default:
		return string(t[start+1 : end])
	}
}

type namedCol struct {
	name    string
	colType column.Type
	iface   column.Interface
}

func parseTupleType(t column.Type, tz *time.Location) ([]column.Interface, error) {
	var (
		element       []rune
		elements      []namedCol
		brackets      int
		appendElement = func() {
			if len(element) != 0 {
				cType := strings.TrimSpace(string(element))
				name := ""
				if parts := strings.SplitN(cType, " ", 2); len(parts) == 2 {
					if !strings.Contains(parts[0], "(") {
						name = parts[0]
						cType = parts[1]
					}
				}
				elements = append(elements, namedCol{
					name:    name,
					colType: column.Type(strings.TrimSpace(cType)),
				})
			}
		}
	)
	for _, r := range params(t) {
		switch r {
		case '(':
			brackets++
		case ')':
			brackets--
		case ',':
			if brackets == 0 {
				appendElement()
				element = element[:0]
				continue
			}
		}
		element = append(element, r)
	}
	appendElement()
	columns := make([]column.Interface, len(elements))
	for i, ct := range elements {
		if len(ct.name) == 0 {
			return nil, fmt.Errorf("unsupported ClickHouse Tuple type (Aache Arrow requires named fields): %s", t)
		}
		col, err := ct.colType.Column(ct.name, tz)
		if err != nil {
			return nil, err
		}
		columns[i] = col
	}

	return columns, nil
}
