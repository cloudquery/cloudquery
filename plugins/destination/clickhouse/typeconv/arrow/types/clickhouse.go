package types

// This file contains some code extracted from github.com/ClickHouse/clickhouse-go
//
// Copyright 2016-2023 ClickHouse, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
}

func tupleFieldSpec(spec string) *namedCol {
	spec = strings.TrimSpace(spec)
	if len(spec) == 0 {
		return nil
	}

	if parts := strings.SplitN(spec, " ", 2); len(parts) == 2 {
		if !strings.Contains(parts[0], "(") {
			return &namedCol{
				name:    strings.TrimSpace(parts[0]),
				colType: column.Type(strings.TrimSpace(parts[1])),
			}
		}
	}

	return &namedCol{colType: column.Type(strings.TrimSpace(spec))}
}

func parseTupleType(t column.Type, tz *time.Location) ([]column.Interface, error) {
	var (
		elements []namedCol
		brackets int
	)

	p := params(t)
	spec := make([]rune, 0, len(p))

	for _, r := range p {
		switch r {
		case '(':
			brackets++
		case ')':
			brackets--
		case ',':
			if brackets == 0 {
				col := tupleFieldSpec(string(spec))
				if col != nil {
					elements = append(elements, *col)
				}
				spec = spec[:0] // cleanup
				continue
			}
		}
		spec = append(spec, r)
	}
	col := tupleFieldSpec(string(spec))
	if col != nil {
		elements = append(elements, *col)
	}

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
