package policy

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/hashicorp/go-hclog"
	"github.com/jeremywohl/flatten"
)

func persistSnapshot(ctx context.Context, e *Executor, path string, table string) error {
	ef, err := os.OpenFile(filepath.Join(path, fmt.Sprintf("table_%s.csv", table)), os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return fmt.Errorf("error opening file %q: %w", table, err)
	}
	defer ef.Close()
	query := fmt.Sprintf("COPY (select * from %s) TO STDOUT DELIMITER '|' CSV HEADER", table)
	e.log.Debug("exporting", "table", table, "query", query)

	err = storeOutput(ctx, e, ef, query)
	if err != nil {
		return fmt.Errorf("error exporting file: %w", err)
	}
	return nil

}
func StoreSnapshot(ctx context.Context, e *Executor, path string, tables []string) error {
	if len(tables) == 0 {
		return errors.New("no tables to snapshot")
	}

	for _, table := range tables {
		err := persistSnapshot(ctx, e, path, table)
		if err != nil {
			return err
		}

	}

	return nil
}

func RestoreSnapshot(ctx context.Context, conn LowLevelQueryExecer, log hclog.Logger, source string) error {
	ef, err := os.OpenFile(source, os.O_RDONLY, 0777)
	if err != nil {
		return fmt.Errorf("error opening file for restore: %w", err)
	}
	defer ef.Close()
	fileName := path.Base(source)
	if !strings.HasPrefix(fileName, "table_") || !strings.HasSuffix(fileName, ".csv") {
		log.Error("truncating", "source", source, "file", fileName)
		return fmt.Errorf("invalid filename: %q", fileName)
	}
	source_name := strings.TrimPrefix(strings.TrimSuffix(fileName, ".csv"), "table_")
	err = deleteFKs(ctx, conn, log, source_name)
	if err != nil {
		return fmt.Errorf("error removing fks from %q: %w", source_name, err)
	}
	truncQuery := fmt.Sprintf("TRUNCATE %s CASCADE", source_name)
	log.Debug("truncating", "table", source_name, "query", truncQuery)
	err = conn.Exec(ctx, truncQuery)
	if err != nil {
		return fmt.Errorf("error importing file %q: %w", fileName, err)
	}

	query := fmt.Sprintf("copy \"%s\" from stdin DELIMITER '|' CSV HEADER;", source_name)
	log.Debug("importing", "table", source_name, "query", query)

	err = conn.RawCopyFrom(ctx, ef, query)
	if err != nil {
		return fmt.Errorf("error importing file: %w", err)
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

func (ce *Executor) extractTableNames(ctx context.Context, query string) ([]string, error) {
	tableNames := make([]string, 0)
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
	if err := rows.Err(); err != nil {
		ce.log.Error("Error fetching rows", "query", query, "error", err)
		return tableNames, err
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
				viewTables, err := ce.extractTableNames(ctx, viewQuery)
				if err != nil {
					return []string{}, err
				}
				ce.log.Error("Added tables from view", "resp", viewTables)
				tableNames = append(tableNames, viewTables...)
			}
		}
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

func StoreOutput(ctx context.Context, e *Executor, pol *Policy, destination string) (err error) {
	if !pol.HasChecks() {
		return errors.New("no checks")
	}
	err = e.createViews(ctx, pol)
	if err != nil {
		return fmt.Errorf("failed to create views: %w", err)
	}

	ef, err := os.OpenFile(filepath.Join(destination, "snapshot_data.csv"), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		e.log.Error("error opening file:", err)
		return err
	}
	defer ef.Close()
	query := fmt.Sprintf("COPY (%s)  TO STDOUT DELIMITER '|' CSV HEADER", cleanQuery(pol.Checks[0].Query))
	err = storeOutput(ctx, e, ef, query)
	if err != nil {
		e.log.Error("Storing output error", "query", query, "error", err)
		return err
	}
	return nil
}

func storeOutput(ctx context.Context, e *Executor, w io.Writer, sql string) error {
	e.log.Debug("Copying output to writer", "query", sql)
	err := e.conn.RawCopyTo(ctx, w, sql)
	if err != nil {
		return fmt.Errorf("error exporting file: %w", err)
	}
	return nil
}

func (ce *Executor) checkTableExistence(ctx context.Context, tableName string) (query string, err error) {

	explainQuery := fmt.Sprintf("select coalesce(pg_get_viewdef('%s'::regclass::oid),'') ", tableName)
	rows, err := ce.conn.Query(ctx, explainQuery)
	if err != nil {
		ce.log.Error("error running explain", "tableName", tableName, "err", err)
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

func deleteFKs(ctx context.Context, conn LowLevelQueryExecer, log hclog.Logger, tableName string) error {
	fkQuery := fmt.Sprintf(`SELECT conname AS foreign_key
FROM   pg_constraint 
WHERE  contype = 'f' 
AND    connamespace = 'public'::regnamespace
AND conrelid::regclass = '%s'::regclass;`, tableName)
	rows, err := conn.Query(ctx, fkQuery)
	if err != nil {
		log.Error("error finding fks for table", "query", fkQuery, "err", err)
		return err
	}

	for rows.Next() {
		var fkName string
		if err := rows.Scan(&fkName); err != nil {
			log.Error("error scanning into variable", "error", err)
		}
		deletionQuery := fmt.Sprintf("ALTER TABLE %s DROP CONSTRAINT  IF EXISTS %s;", tableName, fkName)
		err = conn.Exec(ctx, deletionQuery)
		if err != nil {
			log.Error("error deleting fks for table", "query", deletionQuery, "err", err)
			return err
		}

	}
	if err := rows.Err(); err != nil {
		log.Error("Error fetching rows", "query", fkQuery, "error", err)
		return err
	}

	return nil
}
