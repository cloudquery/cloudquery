package recordupdater

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/cloudquery/cloudquery/plugins/transformer/basic/client/schemaupdater"
	"github.com/cloudquery/cloudquery/plugins/transformer/basic/client/spec"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// RecordUpdater takes an `arrow.RecordBatch` and knows how to make simple subsequent changes to it.
// It doesn't know which table it belongs to or if the changes make sense.
type RecordUpdater struct {
	record        arrow.RecordBatch
	schemaUpdater *schemaupdater.SchemaUpdater
}

func New(record arrow.RecordBatch) *RecordUpdater {
	return &RecordUpdater{
		record:        record,
		schemaUpdater: schemaupdater.New(record.Schema()),
	}
}

const redactedByCQMessage = "Redacted by CloudQuery |"
const redactedByCQMessageNoSHA = "Redacted by CloudQuery"
const redactedByCQJSONName = "redacted_by_cloudquery"

// internalColumnPrefix identifies CloudQuery-managed columns (e.g. _cq_id, _cq_source_name)
// that must always pass through redaction untouched.
const internalColumnPrefix = "_cq_"

func isInternalColumn(name string) bool {
	return strings.HasPrefix(name, internalColumnPrefix)
}

func isPrimaryKeyField(f arrow.Field) bool {
	idx := f.Metadata.FindKey(schema.MetadataPrimaryKey)
	return idx >= 0 && f.Metadata.Values()[idx] == schema.MetadataTrue
}

// redactValue returns the redaction marker for a raw value. When includeSHA is true the
// SHA-256 hash of the value is appended so distinct values remain distinguishable; when
// false a bare marker is returned.
func redactValue(raw []byte, includeSHA bool) string {
	if includeSHA {
		return fmt.Sprintf("%s %x", redactedByCQMessage, sha256.Sum256(raw))
	}
	return redactedByCQMessageNoSHA
}

func (r *RecordUpdater) RemoveColumns(columnNames []string) (arrow.RecordBatch, error) {
	plainCols, jsonCols := r.splitJSONColumns(columnNames)

	if len(plainCols) > 0 {
		colIndices, err := r.colIndicesByNames(plainCols)
		if err != nil {
			return nil, err
		}
		if len(colIndices) == int(r.record.NumCols()) {
			return nil, errors.New("cannot remove all columns")
		}

		oldRecord := r.record.Columns()
		newColumns := make([]arrow.Array, 0, len(oldRecord)-len(colIndices))
		for i, column := range oldRecord {
			if _, ok := colIndices[i]; ok {
				continue
			}
			newColumns = append(newColumns, column)
		}

		r.record = array.NewRecordBatch(r.schemaUpdater.RemoveColumnIndices(colIndices), newColumns, r.record.NumRows())
	}

	if len(jsonCols) > 0 {
		for i, jcs := range r.jsonColIndicesByNames(jsonCols) {
			bld := types.NewJSONBuilder(memory.NewGoAllocator())
			for j := 0; j < r.record.Column(i).Len(); j++ {
				valStr := r.record.Column(i).ValueStr(j)
				if gjson.Valid(valStr) {
					for _, jc := range jcs {
						if val, err := sjson.Delete(valStr, jc.columnPath); err == nil {
							valStr = val
						}
					}
				}
				bld.AppendBytes([]byte(valStr))
			}

			rec, err := r.record.SetColumn(i, bld.NewJSONArray())
			if err != nil {
				return nil, err
			}
			r.record = rec
		}
	}
	return r.record, nil
}

func (r *RecordUpdater) AddLiteralStringColumn(columnName, columnValue string, position int) (arrow.RecordBatch, error) {
	if position == -1 {
		position = int(r.record.NumCols())
	}
	if position < 0 || position > int(r.record.NumCols()) {
		return nil, fmt.Errorf("invalid position %v", position)
	}

	newColumns := make([]arrow.Array, 0, int(r.record.NumCols())+1)
	for i := 0; i < int(r.record.NumCols()); i++ {
		if i == position {
			newColumns = append(newColumns, r.buildStringColumn(columnValue, int(r.record.NumRows())))
		}
		newColumns = append(newColumns, r.record.Column(i))
	}
	if position == int(r.record.NumCols()) {
		newColumns = append(newColumns, r.buildStringColumn(columnValue, int(r.record.NumRows())))
	}
	newSchema, err := r.schemaUpdater.AddStringColumnAtPos(columnName, position, false)
	if err != nil {
		return nil, err
	}
	r.record = array.NewRecordBatch(newSchema, newColumns, r.record.NumRows())
	return r.record, nil
}

func (r *RecordUpdater) AddTimestampColumn(columnName string, position int) (arrow.RecordBatch, error) {
	if position == -1 {
		position = int(r.record.NumCols())
	}
	if position < 0 || position > int(r.record.NumCols()) {
		return nil, fmt.Errorf("invalid position %v", position)
	}
	timeVal := time.Now()

	newColumns := make([]arrow.Array, 0, int(r.record.NumCols())+1)
	for i := 0; i < int(r.record.NumCols()); i++ {
		if i == position {
			newColumns = append(newColumns, r.buildCurrentTimestampColumn(timeVal, int(r.record.NumRows())))
		}
		newColumns = append(newColumns, r.record.Column(i))
	}
	if position == int(r.record.NumCols()) {
		newColumns = append(newColumns, r.buildCurrentTimestampColumn(timeVal, int(r.record.NumRows())))
	}
	newSchema, err := r.schemaUpdater.AddTimestampColumnAtPos(columnName, position, true)
	if err != nil {
		return nil, err
	}
	r.record = array.NewRecordBatch(newSchema, newColumns, r.record.NumRows())
	return r.record, nil
}

func (r *RecordUpdater) ObfuscateSensitiveColumns(includeSHA bool) (arrow.RecordBatch, error) {
	if r.record.Schema() == nil {
		return nil, errors.New("record schema is nil")
	}
	s, ok := r.record.Schema().Metadata().GetValue(schema.MetadataTableSensitiveColumns)
	if !ok {
		return r.record, nil
	}
	var sensitiveColumnsArr []string
	err := json.Unmarshal([]byte(s), &sensitiveColumnsArr)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal sensitive columns: %w", err)
	}
	if len(sensitiveColumnsArr) == 0 {
		return r.record, nil
	}
	return r.ObfuscateColumns(sensitiveColumnsArr, includeSHA)
}

func (r *RecordUpdater) DropRows(columnNames []string, value *string) (arrow.RecordBatch, error) {
	cols := r.record.Columns()

	rowsToDrop := make(map[int]bool)
	for j, column := range cols {
		if !slices.Contains(columnNames, r.record.ColumnName(j)) {
			continue
		}
		for i := range column.Len() {
			// check if i in map already, if so, keep going
			if rowsToDrop[i] {
				continue
			}
			// If Value specified by the user is nil, and Column is null, we drop the row.
			// Or if Value specified by the user is not nil, and Column is valid and equal to the Value, we drop the row.
			if column.IsNull(i) && value == nil || value != nil && column.IsValid(i) && column.ValueStr(i) == *value {
				rowsToDrop[i] = true
			}
		}
	}
	if len(rowsToDrop) == 0 {
		return r.record, nil
	}
	newRowLen := int(r.record.NumRows()) - len(rowsToDrop)
	rowSlices := make([]arrow.RecordBatch, 0, newRowLen)

	// This section builds slices of rows that are not to be dropped.
	currentSliceStart := -1
	for row := range r.record.NumRows() {
		if !rowsToDrop[int(row)] {
			if currentSliceStart == -1 {
				currentSliceStart = int(row)
			}
			// This handles the edge case of checking the last row
			if row == r.record.NumRows()-1 && currentSliceStart != -1 {
				rowSlices = append(rowSlices, r.record.NewSlice(int64(currentSliceStart), row+1))
			}
			continue
		}
		// if we reach here, it means that the current row is supposed to be dropped, so we create a NewSlice and reset currentSliceStart
		if currentSliceStart != -1 {
			rowSlices = append(rowSlices, r.record.NewSlice(int64(currentSliceStart), row))
			currentSliceStart = -1
		}
	}
	concatenatedCols := make([]arrow.Array, int(r.record.NumCols()))
	for i := range r.record.NumCols() {
		var colChunks []arrow.Array
		for _, slice := range rowSlices {
			colChunks = append(colChunks, slice.Column(int(i)))
		}

		if len(rowSlices) > 0 {
			concat, err := array.Concatenate(colChunks, memory.DefaultAllocator)
			if err != nil {
				return nil, fmt.Errorf("failed to concatenate arrays: %w", err)
			}
			concatenatedCols[i] = concat
		} else {
			builder := array.NewBuilder(memory.DefaultAllocator, r.record.Column(int(i)).DataType())
			concatenatedCols[i] = builder.NewArray()
		}
	}

	r.record = array.NewRecordBatch(r.record.Schema(), concatenatedCols, int64(newRowLen))
	return r.record, nil
}

func (r *RecordUpdater) ObfuscateColumns(columnNames []string, includeSHA bool) (arrow.RecordBatch, error) {
	plainCols, jsonCols := r.splitJSONColumns(columnNames)

	plainColIndex, err := r.colIndicesByNames(plainCols)
	if err != nil {
		return nil, err
	}
	jsonColIndex := r.jsonColIndicesByNames(jsonCols)

	oldRecord := r.record.Columns()
	newColumns := make([]arrow.Array, 0, len(oldRecord))
	for i, column := range oldRecord {
		if _, ok := plainColIndex[i]; ok {
			if column.DataType().ID() == arrow.STRING {
				newColumns = append(newColumns, r.obfuscateColumn(column, includeSHA))
				continue
			}
			if _, ok := column.DataType().(*types.JSONType); ok {
				newColumns = append(newColumns, r.obfuscateEntireJSONColumn(column, includeSHA))
				continue
			}
			if column.DataType().ID() == arrow.BINARY {
				newColumns = append(newColumns, r.obfuscateBinaryColumn(column, includeSHA))
				continue
			}
			return nil, fmt.Errorf("column %v is not a string, binary or JSON column", r.record.ColumnName(i))
		}

		jcs, ok := jsonColIndex[i]
		if !ok {
			newColumns = append(newColumns, column)
			continue
		}

		if _, ok := column.DataType().(*types.JSONType); !ok {
			return nil, fmt.Errorf("column %v is not a JSON column", r.record.ColumnName(i))
		}

		newColumns = append(newColumns, r.obfuscateJSONColumns(column, jcs, includeSHA))
	}

	r.record = array.NewRecordBatch(r.record.Schema(), newColumns, r.record.NumRows())

	return r.record, nil
}

// ObfuscateColumnsExcept redacts every column EXCEPT those on the keepColumns allowlist.
// It is the opt-in inverse of ObfuscateColumns:
//   - CloudQuery internal columns (_cq_*) always pass through untouched.
//   - A bare column name on the allowlist keeps that column verbatim.
//   - A dotted allowlist entry (e.g. "spec_containers.#.image") keeps that JSON sub-path
//     and redacts every other leaf within that JSON column in place.
//   - A non-allowlisted column is obfuscated when it is a string/JSON/binary column, and
//     dropped when its type cannot be hashed into itself (numbers, timestamps, lists, ...).
//
// Allowlist paths use gjson syntax; "#" matches any array index.
func (r *RecordUpdater) ObfuscateColumnsExcept(keepColumns []string, includeSHA bool) (arrow.RecordBatch, error) {
	keepWhole := make(map[string]struct{})
	keepSub := make(map[string][][]string)
	for _, c := range keepColumns {
		if idx := strings.Index(c, "."); idx > -1 {
			col := c[:idx]
			keepSub[col] = append(keepSub[col], splitJSONPath(c[idx+1:]))
			continue
		}
		keepWhole[c] = struct{}{}
	}

	oldColumns := r.record.Columns()
	fields := r.record.Schema().Fields()

	// Fail closed on misconfiguration: an allowlist entry that names a column not present
	// in the table silently over-redacts the column the user meant to keep. Surface it.
	if err := r.validateKeepColumnsExist(keepWhole, keepSub); err != nil {
		return nil, err
	}
	// Guard row identity: a non-allowlisted primary-key column would be redacted, which
	// breaks upserts at the destination — collapsing to one key with include_sha=false, or
	// vanishing entirely if its type can't be hashed.
	if err := r.validatePrimaryKeysSurvive(keepWhole, keepSub, includeSHA); err != nil {
		return nil, err
	}

	newFields := make([]arrow.Field, 0, len(oldColumns))
	newColumns := make([]arrow.Array, 0, len(oldColumns))

	for i, column := range oldColumns {
		name := r.record.ColumnName(i)
		_, isJSON := column.DataType().(*types.JSONType)
		_, keptWhole := keepWhole[name]

		switch {
		case isInternalColumn(name):
			// CloudQuery internal columns always pass through.
			newFields = append(newFields, fields[i])
			newColumns = append(newColumns, column)
		case keptWhole:
			// Fully allowlisted column: keep verbatim.
			newFields = append(newFields, fields[i])
			newColumns = append(newColumns, column)
		case len(keepSub[name]) > 0:
			// JSON column with allowlisted sub-paths: keep them, redact the rest.
			if !isJSON {
				return nil, fmt.Errorf("column %q is referenced with a nested path in the allowlist but is not a JSON column", name)
			}
			newFields = append(newFields, fields[i])
			newColumns = append(newColumns, r.redactJSONColumnExcept(column, keepSub[name], includeSHA))
		case column.DataType().ID() == arrow.STRING:
			newFields = append(newFields, fields[i])
			newColumns = append(newColumns, r.obfuscateColumn(column, includeSHA))
		case isJSON:
			newFields = append(newFields, fields[i])
			newColumns = append(newColumns, r.obfuscateEntireJSONColumn(column, includeSHA))
		case column.DataType().ID() == arrow.BINARY:
			newFields = append(newFields, fields[i])
			newColumns = append(newColumns, r.obfuscateBinaryColumn(column, includeSHA))
		default:
			// Un-hashable type (number, timestamp, list, struct, ...): drop it.
		}
	}

	if len(newColumns) == 0 {
		return nil, errors.New("obfuscate_columns_except would drop every column; allowlist at least one string/JSON/binary column")
	}

	md := r.record.Schema().Metadata()
	newSchema := arrow.NewSchema(newFields, &md)
	r.record = array.NewRecordBatch(newSchema, newColumns, r.record.NumRows())
	return r.record, nil
}

// validateKeepColumnsExist errors if any allowlist entry names a top-level column that is
// not present in the record, which would otherwise silently over-redact.
func (r *RecordUpdater) validateKeepColumnsExist(keepWhole map[string]struct{}, keepSub map[string][][]string) error {
	present := make(map[string]struct{}, r.record.NumCols())
	for i := 0; i < int(r.record.NumCols()); i++ {
		present[r.record.ColumnName(i)] = struct{}{}
	}
	var missing []string
	for name := range keepWhole {
		if _, ok := present[name]; !ok {
			missing = append(missing, name)
		}
	}
	for name := range keepSub {
		if _, ok := present[name]; !ok {
			missing = append(missing, name)
		}
	}
	if len(missing) > 0 {
		slices.Sort(missing)
		return fmt.Errorf("obfuscate_columns_except allowlist references unknown column(s): %s", strings.Join(missing, ", "))
	}
	return nil
}

// validatePrimaryKeysSurvive errors if a non-allowlisted primary-key column would lose its
// row-identifying value (dropped for un-hashable types, or collapsed to one value under
// include_sha=false), which would corrupt upserts at the destination.
func (r *RecordUpdater) validatePrimaryKeysSurvive(keepWhole map[string]struct{}, keepSub map[string][][]string, includeSHA bool) error {
	fields := r.record.Schema().Fields()
	for i := 0; i < int(r.record.NumCols()); i++ {
		name := r.record.ColumnName(i)
		if isInternalColumn(name) || !isPrimaryKeyField(fields[i]) {
			continue
		}
		if _, ok := keepWhole[name]; ok {
			continue
		}
		col := r.record.Column(i)
		_, isJSON := col.DataType().(*types.JSONType)
		if _, ok := keepSub[name]; ok && isJSON {
			continue
		}
		redactable := col.DataType().ID() == arrow.STRING || col.DataType().ID() == arrow.BINARY || isJSON
		switch {
		case !redactable:
			return fmt.Errorf("primary-key column %q has type %s that cannot be redacted in place, so it would be dropped and break row identity at the destination; add %q to the allowlist", name, col.DataType(), name)
		case !includeSHA:
			return fmt.Errorf("primary-key column %q is not on the allowlist; with include_sha=false its redacted value is identical for every row and would break upserts at the destination; add %q to the allowlist or use include_sha=true", name, name)
		}
	}
	return nil
}

// redactJSONColumnExcept redacts every leaf of each JSON value except those covered by the
// keep patterns.
func (*RecordUpdater) redactJSONColumnExcept(column arrow.Array, patterns [][]string, includeSHA bool) arrow.Array {
	bld := types.NewJSONBuilder(memory.NewGoAllocator())
	for i := 0; i < column.Len(); i++ {
		if !column.IsValid(i) {
			bld.AppendNull()
			continue
		}
		s := column.ValueStr(i)
		out, err := redactJSONExcept(s, patterns, includeSHA)
		if err != nil {
			// Not parseable as JSON: fall back to redacting the whole value so nothing leaks.
			fallback, ok := entireJSONRedaction([]byte(s), includeSHA)
			if !ok {
				bld.AppendNull()
				continue
			}
			out = string(fallback)
		}
		bld.AppendBytes([]byte(out))
	}
	return bld.NewJSONArray()
}

// redactJSONExcept parses jsonStr, redacts every leaf not covered by a keep pattern, and
// returns the re-serialized JSON. Number fidelity is preserved via json.Number; HTML
// escaping is disabled so values are not mangled.
func redactJSONExcept(jsonStr string, patterns [][]string, includeSHA bool) (string, error) {
	dec := json.NewDecoder(strings.NewReader(jsonStr))
	dec.UseNumber()
	var v any
	if err := dec.Decode(&v); err != nil {
		return "", err
	}

	v = walkRedactExcept(v, nil, patterns, includeSHA)

	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	if err := enc.Encode(v); err != nil {
		return "", err
	}
	return strings.TrimRight(buf.String(), "\n"), nil
}

func walkRedactExcept(node any, path []string, patterns [][]string, includeSHA bool) any {
	switch n := node.(type) {
	case map[string]any:
		for k, val := range n {
			n[k] = walkRedactExcept(val, appendPath(path, k), patterns, includeSHA)
		}
		return n
	case []any:
		for i, val := range n {
			n[i] = walkRedactExcept(val, appendPath(path, strconv.Itoa(i)), patterns, includeSHA)
		}
		return n
	case nil:
		// Preserve nulls: they carry no value to redact.
		return nil
	default:
		if pathKept(path, patterns) {
			return node
		}
		raw, _ := json.Marshal(node)
		return redactValue(raw, includeSHA)
	}
}

// appendPath returns a fresh slice so recursive siblings never alias the same backing array.
func appendPath(path []string, seg string) []string {
	out := make([]string, len(path)+1)
	copy(out, path)
	out[len(path)] = seg
	return out
}

// pathKept reports whether the concrete leaf path is at or under any keep pattern.
func pathKept(path []string, patterns [][]string) bool {
	for _, pat := range patterns {
		if patternIsPrefix(pat, path) {
			return true
		}
	}
	return false
}

// patternIsPrefix reports whether pat matches a prefix of path. A "#" segment in pat matches
// any array index (a numeric segment); every other segment must match exactly.
func patternIsPrefix(pat, path []string) bool {
	if len(pat) > len(path) {
		return false
	}
	for i := range pat {
		if pat[i] == "#" {
			if !isArrayIndex(path[i]) {
				return false
			}
			continue
		}
		if pat[i] != path[i] {
			return false
		}
	}
	return true
}

func isArrayIndex(seg string) bool {
	if seg == "" {
		return false
	}
	for _, r := range seg {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

// splitJSONPath splits a gjson-style path on unescaped dots, unescaping "\." into a literal dot.
func splitJSONPath(p string) []string {
	var segs []string
	var cur strings.Builder
	for i := 0; i < len(p); i++ {
		if p[i] == '\\' && i+1 < len(p) && p[i+1] == '.' {
			cur.WriteByte('.')
			i++
			continue
		}
		if p[i] == '.' {
			segs = append(segs, cur.String())
			cur.Reset()
			continue
		}
		cur.WriteByte(p[i])
	}
	segs = append(segs, cur.String())
	return segs
}

func (r *RecordUpdater) AddPrimaryKeys(columnNames []string) (arrow.RecordBatch, error) {
	newSchema, err := r.schemaUpdater.AddPrimaryKeys(columnNames)
	if err != nil {
		return nil, err
	}
	r.record = array.NewRecordBatch(newSchema, r.record.Columns(), r.record.NumRows())
	return r.record, nil
}

func (r *RecordUpdater) ChangeTableName(newTableNamePattern string) (arrow.RecordBatch, error) {
	newSchema, err := r.schemaUpdater.ChangeTableName(newTableNamePattern)
	if err != nil {
		return nil, err
	}
	r.record = array.NewRecordBatch(newSchema, r.record.Columns(), r.record.NumRows())
	return r.record, nil
}

func (r *RecordUpdater) RenameColumn(oldName, newName string) (arrow.RecordBatch, error) {
	newSchema, err := r.schemaUpdater.RenameColumn(oldName, newName)
	if err != nil {
		return nil, err
	}
	r.record = array.NewRecordBatch(newSchema, r.record.Columns(), r.record.NumRows())
	return r.record, nil
}

func (r *RecordUpdater) ChangeCase(caseType string, columnNames []string) (arrow.RecordBatch, error) {
	plainCols, jsonCols := r.splitJSONColumns(columnNames)

	plainColIndex, err := r.colIndicesByNames(plainCols)
	if err != nil {
		return nil, err
	}
	jsonColIndex := r.jsonColIndicesByNames(jsonCols)

	caser := strings.ToLower
	if caseType == spec.KindUppercase {
		caser = strings.ToUpper
	}

	oldRecord := r.record.Columns()
	newColumns := make([]arrow.Array, 0, len(oldRecord))
	for i, column := range oldRecord {
		if _, ok := plainColIndex[i]; ok {
			if column.DataType().ID() == arrow.STRING {
				newColumns = append(newColumns, r.changeColumnCase(column, caser))
				continue
			}
			if _, ok := column.DataType().(*types.JSONType); ok {
				newColumns = append(newColumns, r.changeCaseEntireJSONColumn(column, caser))
				continue
			}
			return nil, fmt.Errorf("column %v is not a string or JSON column", r.record.ColumnName(i))
		}

		jcs, ok := jsonColIndex[i]
		if !ok {
			newColumns = append(newColumns, column)
			continue
		}

		if _, ok := column.DataType().(*types.JSONType); !ok {
			return nil, fmt.Errorf("column %v is not a JSON column", r.record.ColumnName(i))
		}

		newColumns = append(newColumns, r.chanceCaseJSONColumns(column, jcs, caser))
	}

	r.record = array.NewRecordBatch(r.record.Schema(), newColumns, r.record.NumRows())

	return r.record, nil
}

func (r *RecordUpdater) colIndicesByNames(columnNames []string) (map[int]struct{}, error) {
	colNameMap := make(map[string]struct{})
	for _, columnName := range columnNames {
		colNameMap[columnName] = struct{}{}
	}

	colIndexes := make(map[int]struct{})
	for i := 0; i < int(r.record.NumCols()); i++ {
		colName := r.record.ColumnName(i)
		if _, ok := colNameMap[colName]; ok {
			colIndexes[i] = struct{}{}
			delete(colNameMap, colName)
		}
	}
	if len(colNameMap) > 0 {
		missingColumns := make([]string, 0, len(colNameMap))
		for colName := range colNameMap {
			missingColumns = append(missingColumns, colName)
		}
		return nil, fmt.Errorf("columns %v not found", strings.Join(missingColumns, ", "))
	}

	return colIndexes, nil
}

type jsonColumn struct {
	columnName string
	columnPath string
}

func (r *RecordUpdater) jsonColIndicesByNames(columns map[string]jsonColumn) map[int][]jsonColumn {
	colNameMap := make(map[int][]jsonColumn)
	for i := 0; i < int(r.record.NumCols()); i++ {
		for _, jc := range columns {
			if jc.columnName == r.record.ColumnName(i) {
				if _, ok := colNameMap[i]; !ok {
					colNameMap[i] = []jsonColumn{jc}
				} else {
					colNameMap[i] = append(colNameMap[i], jc)
				}
			}
		}
	}
	return colNameMap
}

func (*RecordUpdater) buildStringColumn(literalValue string, numRows int) arrow.Array {
	bld := array.NewStringBuilder(memory.DefaultAllocator)
	for i := 0; i < numRows; i++ {
		bld.AppendString(literalValue)
	}
	return bld.NewStringArray()
}

func (*RecordUpdater) buildCurrentTimestampColumn(t time.Time, numRows int) arrow.Array {
	ts, _ := arrow.TimestampFromTime(t, arrow.Microsecond)
	syncTimeBldr := array.NewTimestampBuilder(memory.DefaultAllocator, &arrow.TimestampType{Unit: arrow.Microsecond, TimeZone: "UTC"})
	for i := 0; i < numRows; i++ {
		syncTimeBldr.Append(ts)
	}
	return syncTimeBldr.NewArray()
}

func (*RecordUpdater) obfuscateColumn(column arrow.Array, includeSHA bool) arrow.Array {
	bld := array.NewStringBuilder(memory.DefaultAllocator)
	for i := 0; i < column.Len(); i++ {
		if !column.IsValid(i) {
			bld.AppendNull()
			continue
		}
		bld.AppendString(redactValue([]byte(column.ValueStr(i)), includeSHA))
	}
	return bld.NewStringArray()
}

func (*RecordUpdater) obfuscateJSONColumns(column arrow.Array, jcs []jsonColumn, includeSHA bool) arrow.Array {
	bld := types.NewJSONBuilder(memory.NewGoAllocator())
	for i := 0; i < column.Len(); i++ {
		if !column.IsValid(i) {
			bld.AppendNull()
			continue
		}

		str := column.ValueStr(i)
		for _, jc := range jcs {
			val := gjson.Get(column.ValueStr(i), jc.columnPath)
			// todo: Currently nested types will create a single SHA hash for all matched array elements. Consider changing this to hash for each element separately.
			if val.Exists() {
				if modified, err := sjson.Set(str, jc.columnPath, redactValue([]byte(val.Raw), includeSHA)); err == nil {
					str = modified
					continue
				}
			}
		}
		bld.AppendBytes([]byte(str))
	}
	return bld.NewJSONArray()
}

func (*RecordUpdater) obfuscateBinaryColumn(column arrow.Array, includeSHA bool) arrow.Array {
	bld := array.NewBinaryBuilder(memory.DefaultAllocator, &arrow.BinaryType{})
	for i := 0; i < column.Len(); i++ {
		if !column.IsValid(i) {
			bld.AppendNull()
			continue
		}
		bc, ok := column.(*array.Binary)
		if !ok {
			bld.AppendNull()
			continue
		}
		bld.Append([]byte(redactValue(bc.Value(i), includeSHA)))
	}
	return bld.NewBinaryArray()
}

// entireJSONRedaction returns the {"redacted_by_cloudquery": <marker>} document used when a
// whole JSON value is redacted. The marker is the SHA-256 of the value, or a static string
// when include_sha is false. The bool reports whether serialization succeeded.
func entireJSONRedaction(raw []byte, includeSHA bool) ([]byte, bool) {
	marker := redactedByCQMessageNoSHA
	if includeSHA {
		marker = fmt.Sprintf("%x", sha256.Sum256(raw))
	}
	out, err := sjson.Set("{}", redactedByCQJSONName, marker)
	if err != nil {
		return nil, false
	}
	return []byte(out), true
}

func (*RecordUpdater) obfuscateEntireJSONColumn(column arrow.Array, includeSHA bool) arrow.Array {
	bld := types.NewJSONBuilder(memory.NewGoAllocator())
	for i := 0; i < column.Len(); i++ {
		if !column.IsValid(i) {
			bld.AppendNull()
			continue
		}
		if out, ok := entireJSONRedaction([]byte(column.ValueStr(i)), includeSHA); ok {
			bld.AppendBytes(out)
		} else {
			bld.AppendNull()
		}
	}
	return bld.NewJSONArray()
}

func (r *RecordUpdater) splitJSONColumns(columnNames []string) (plainCols []string, jsonCols map[string]jsonColumn) {
	plainColMap := make(map[string]struct{})
	jsonColMap := make(map[string]jsonColumn)
	for _, columnName := range columnNames {
		if idx := strings.Index(columnName, "."); idx > -1 {
			jsonColMap[columnName] = jsonColumn{columnName: columnName[:idx], columnPath: columnName[idx+1:]}
		} else {
			plainColMap[columnName] = struct{}{}
		}
	}

	for i := 0; i < int(r.record.NumCols()); i++ {
		if _, ok := plainColMap[r.record.ColumnName(i)]; ok {
			plainCols = append(plainCols, r.record.ColumnName(i))
			continue
		}
	}

	for k, jc := range jsonColMap {
		if slices.Contains(plainCols, jc.columnName) {
			delete(jsonColMap, k)
		}
	}
	jsonCols = make(map[string]jsonColumn)
	for i := 0; i < int(r.record.NumCols()); i++ {
		for _, jc := range jsonColMap {
			if jc.columnName == r.record.ColumnName(i) {
				jsonCols[jc.columnName+"."+jc.columnPath] = jc
			}
		}
	}

	return plainCols, jsonCols
}

func (*RecordUpdater) changeColumnCase(column arrow.Array, caser func(string) string) arrow.Array {
	bld := array.NewStringBuilder(memory.DefaultAllocator)
	for i := 0; i < column.Len(); i++ {
		if !column.IsValid(i) {
			bld.AppendNull()
			continue
		}
		bld.AppendString(caser(column.ValueStr(i)))
	}
	return bld.NewStringArray()
}

func (*RecordUpdater) chanceCaseJSONColumns(column arrow.Array, jcs []jsonColumn, caser func(string) string) arrow.Array {
	bld := types.NewJSONBuilder(memory.NewGoAllocator())
	for i := 0; i < column.Len(); i++ {
		if !column.IsValid(i) {
			bld.AppendNull()
			continue
		}

		str := column.ValueStr(i)
		for _, jc := range jcs {
			val := gjson.Get(column.ValueStr(i), jc.columnPath)
			if val.Exists() && val.Type == gjson.String {
				if modified, err := sjson.Set(str, jc.columnPath, caser(val.Str)); err == nil {
					str = modified
					continue
				}
			}
		}
		bld.AppendBytes([]byte(str))
	}
	return bld.NewJSONArray()
}

func (*RecordUpdater) changeCaseEntireJSONColumn(column arrow.Array, caser func(string) string) arrow.Array {
	bld := types.NewJSONBuilder(memory.NewGoAllocator())
	for i := 0; i < column.Len(); i++ {
		if !column.IsValid(i) {
			bld.AppendNull()
			continue
		}

		str := column.ValueStr(i)
		bld.AppendBytes([]byte(caser(str)))
	}
	return bld.NewJSONArray()
}
