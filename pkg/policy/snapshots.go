package policy

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/jeremywohl/flatten"
)

func (ce *Executor) StoreSnapshot(ctx context.Context, path string, tables []string) error {
	if len(tables) == 0 {
		return errors.New("no tables to snapshot")
	}

	for _, table := range tables {
		ef, err := os.OpenFile(fmt.Sprintf("%s/table_%s.csv", path, table), os.O_CREATE|os.O_WRONLY, 0777)
		if err != nil {
			ce.log.Error("error opening file.", "error", err)
			return err
		}
		defer ef.Close()
		query := fmt.Sprintf("COPY (select * from %s) TO STDOUT DELIMITER '|' CSV HEADER", table)
		ce.log.Debug("exporting", "table", table, "query", query)

		err = ce.storeOutput(ctx, ef, query)
		if err != nil {
			ce.log.Error("error exporting file.", "error", err)
			return fmt.Errorf("error exporting file: %+v", err)
		}

	}

	return nil
}

func (ce *Executor) RestoreSnapshot(ctx context.Context, source string) error {

	ef, err := os.OpenFile(source, os.O_RDONLY, 0777)
	if err != nil {
		ce.log.Error("error opening file.", "error", err)
		return err
	}
	defer ef.Close()
	fileName := path.Base(source)
	if !strings.HasPrefix(fileName, "table_") || !strings.HasSuffix(fileName, ".csv") {
		ce.log.Error("truncating", "source", source, "file", fileName)
		return errors.New("invalid filename")
	}
	source_name := strings.TrimPrefix(strings.TrimSuffix(fileName, ".csv"), "table_")
	truncQuery := fmt.Sprintf("TRUNCATE %s CASCADE", source_name)
	ce.log.Debug("truncating", "table", source_name, "query", truncQuery)
	err = ce.conn.Exec(ctx, truncQuery)
	if err != nil {
		ce.log.Error("error importing data from file.", "error", err, "source", source)
		return fmt.Errorf("error importing file: %+v", err)
	}

	query := fmt.Sprintf("copy \"%s\" from stdin DELIMITER '|' CSV HEADER;", source_name)
	ce.log.Debug("importing", "table", source_name, "query", query)

	err = ce.conn.RawCopyFrom(ctx, ef, query)
	if err != nil {
		return fmt.Errorf("error importing file: %+v", err)
	}

	return nil
}

func cleanQuery(query string) string {
	// This removes comments from the query
	// single line queries are not supported in pg_dump
	var re = regexp.MustCompile(`(?s)\/\*.*?\*\/|--.*?\n`)
	query = re.ReplaceAllString(query, "")
	query = strings.TrimSuffix(query, ";")
	query = strings.TrimSpace(query)
	query = strings.TrimSuffix(query, ";")

	return strings.TrimSpace(query)
}

func (ce *Executor) ExtractTableNames(ctx context.Context, query string) (tableNames []string, err error) {
	cleanedQuery := cleanQuery(query)

	explainQuery := fmt.Sprintf("EXPLAIN (FORMAT JSON) %s", cleanedQuery)

	ce.log.Debug("extracting Table names-cleaned", "explainQuery", explainQuery, "raw query", query)

	rows, err := ce.conn.Query(ctx, explainQuery)
	if err != nil {
		return tableNames, err
	}
	var s string
	for rows.Next() {
		if err := rows.Scan(&s); err != nil {
			ce.log.Error("error scanning into variable", "error", err)
			return nil, err
		}
	}

	var arrayJsonMap []map[string](interface{})
	err = json.Unmarshal([]byte(s), &arrayJsonMap)
	if err != nil {
		ce.log.Error("failed to unmarshal json", "err", err)
		return tableNames, err
	}

	flat, err := flatten.Flatten(arrayJsonMap[0], "", flatten.DotStyle)
	if err != nil {
		ce.log.Error("failed to flatten json", "err", err)
		return tableNames, err
	}
	for key, val := range flat {
		if strings.HasSuffix(key, "Relation Name") {
			ce.log.Debug("Found relation Name", "Table", val)
			tableNames = append(tableNames, val.(string))
		}
		if strings.HasSuffix(key, "Alias") { // Aliases could be query aliases or table aliases (views)

			// Check if query alias, if it is the table will not exist
			viewQuery, err := ce.checkTableExistence(ctx, val.(string))
			if err != nil && strings.Contains(err.Error(), "does not exist (SQLSTATE 42P01)") {
				ce.log.Error("failed to grab query", "err", err, "resp", viewQuery)
				continue
			} else if err != nil {
				return tableNames, err
			}
			// If the alias is a view name, then recursively extract the table names from the subquery
			if viewQuery != "" {
				viewTables, err := ce.ExtractTableNames(ctx, viewQuery)
				if err != nil {
					return []string{}, err
				}
				ce.log.Error("Added tables from view", "resp", viewTables)
				tableNames = append(tableNames, viewTables...)
			}
		}
	}

	if err := rows.Err(); err != nil {
		ce.log.Error("Error fetching rows", "query", query, "error", err)
		return tableNames, err
	}
	// It is possible that tables are used multiple times so need to dedupe prior to returning
	return removeDuplicateValues(tableNames), err
}

func removeDuplicateValues(stringSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func (ce *Executor) StoreOutput(ctx context.Context, pol *Policy, destination string) (err error) {
	if !pol.HasChecks() {
		return errors.New("no checks")
	}
	err = ce.createViews(ctx, pol)
	if err != nil {
		ce.log.Error("error creating views:", err)
		return err
	}

	ef, err := os.OpenFile(fmt.Sprintf("%s/data.csv", destination), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		ce.log.Error("error opening file:", err)
		return err
	}
	defer ef.Close()
	query := fmt.Sprintf("COPY (%s)  TO STDOUT DELIMITER '|' CSV HEADER", cleanQuery(pol.Checks[0].Query))
	err = ce.storeOutput(ctx, ef, query)
	if err != nil {
		ce.log.Error("Storing output error", "query", query, "error", err)
		return err
	}
	return nil
}

func (ce *Executor) storeOutput(ctx context.Context, w io.Writer, sql string) error {
	ce.log.Debug("Copying output to writer", "query", sql)
	err := ce.conn.RawCopyTo(ctx, w, sql)
	if err != nil {
		return fmt.Errorf("error exporting file: %+v", err)
	}
	return nil
}

func (ce *Executor) checkTableExistence(ctx context.Context, tableName string) (query string, err error) {

	explainQuery := fmt.Sprintf("select pg_get_viewdef('%s'::regclass::oid) ", tableName)
	rows, err := ce.conn.Query(ctx, explainQuery)
	if err != nil {
		return "", err
	}
	var s string
	for rows.Next() {

		if err := rows.Scan(&s); err != nil {
			ce.log.Error("error scanning into variable", "error", err)
		}
	}
	if err := rows.Err(); err != nil {
		ce.log.Error("Error fetching rows", "query", explainQuery, "error", err)
		return "", err
	}

	return s, err
}

// func (ce *Executor) CleanDatabase(ctx context.Context) error {

// 	return ce.conn.Exec(ctx, `DROP SCHEMA public CASCADE;
// 	CREATE SCHEMA public;
// 	GRANT ALL ON SCHEMA public TO postgres;
// 	GRANT ALL ON SCHEMA public TO public`)
// }
