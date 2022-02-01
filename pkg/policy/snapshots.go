package policy

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
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
			fmt.Println("error opening file:", err)
			return err
		}
		ce.log.Info("exporting", "table", table)

		err = ce.conn.RawCopyTo(ctx, ef, fmt.Sprintf("COPY (select * from %s) TO STDOUT DELIMITER '|' CSV HEADER", table))
		if err != nil {
			return fmt.Errorf("error exporting file: %+v", err)
		}
		ef.Close()

	}

	return nil
}

func (ce *Executor) RestoreSnapshot(ctx context.Context, source string) error {

	ef, err := os.OpenFile(source, os.O_RDONLY, 0777)
	if err != nil {
		fmt.Println("error opening file:", err)
		return err
	}
	fileName := strings.Split(source, "/")

	source_name := strings.TrimPrefix(strings.TrimSuffix(fileName[len(fileName)-1], ".csv"), "table_")
	truncQuery := fmt.Sprintf("TRUNCATE %s CASCADE", source_name)
	ce.log.Info("truncating", "table", source_name, "query", truncQuery)
	err = ce.conn.Exec(ctx, truncQuery)
	if err != nil {

		return fmt.Errorf("error importing file: %+v", err)
	}

	query := fmt.Sprintf("copy \"%s\" from stdin DELIMITER '|' CSV HEADER;", source_name)
	ce.log.Info("importing", "table", source_name, "query", query)
	defer ef.Close()
	err = ce.conn.RawCopyFrom(ctx, ef, query)
	if err != nil {
		return fmt.Errorf("error importing file: %+v", err)
	}

	return nil
}

func cleanQuery(query string) string {
	var re = regexp.MustCompile(`(?s)\/\*.*?\*\/|--.*?\n`)
	query = re.ReplaceAllString(query, "")
	query = strings.TrimSuffix(query, ";")
	query = strings.TrimSpace(query)
	query = strings.TrimSuffix(query, ";")

	return strings.TrimSpace(query)
}
func (ce *Executor) ExtractTableNames(ctx context.Context, query string) (tableNames []string, err error) {
	ce.log.Debug("extracting Table names-raw", "raw query", query)
	cleanedQuery := cleanQuery(query)
	ce.log.Debug("extracting Table names-cleaned", "cleaned query", cleanedQuery)
	explainQuery := fmt.Sprintf("EXPLAIN (FORMAT JSON) %s", cleanedQuery)

	rows, err := ce.conn.Query(ctx, explainQuery)
	if err != nil {
		return tableNames, err
	}
	var s string
	for rows.Next() {

		if err := rows.Scan(&s); err != nil {
			ce.log.Error("error scanning into variable", "error", err)
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
			ce.log.Info("Found relation Name", "Table", val)
			tableNames = append(tableNames, val.(string))
		}
		if strings.HasSuffix(key, "Alias") {
			viewQuery, err := ce.checkTableExistence(ctx, val.(string))
			if err != nil && strings.Contains(err.Error(), "does not exist (SQLSTATE 42P01)") {
				ce.log.Error("failed to grab query", "err", err, "resp", viewQuery)
				continue
			} else if err != nil {
				return tableNames, err
			}
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
		fmt.Println("error creating views:", err)
		return err
	}

	ef, err := os.OpenFile(fmt.Sprintf("%s/data.csv", destination), os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println("error opening file:", err)
		return err
	}
	query := fmt.Sprintf("COPY (%s)  TO STDOUT DELIMITER '|' CSV HEADER", cleanQuery(pol.Checks[0].Query))
	ce.log.Debug("Copying output", "cleaned query", query)
	err = ce.conn.RawCopyTo(ctx, ef, query)
	if err != nil {
		return fmt.Errorf("error exporting file: %+v", err)
	}
	ef.Close()

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
