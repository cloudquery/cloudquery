package client

import (
	"context"
	"encoding/json"
	"github.com/spf13/afero"

	"github.com/cloudquery/cloudquery/pkg/config"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog/log"
)

type PolicyResult struct {
	Name    string          `json:"name"`
	Columns []string        `json:"result_headers"`
	Data    [][]interface{} `json:"result_rows"`
	Passed  bool            `json:"check_passed"`
}

func createViews(ctx context.Context, conn *pgxpool.Conn, views []config.PolicyView) error {
	log.Info().Int("count", len(views)).Msg("Creating views")
	for _, view := range views {
		if _, err := conn.Exec(ctx, "DROP VIEW IF EXISTS %s", view.Name); err != nil {
			return err
		}
		log.Info().Str("name", view.Name).Msg("Creating view")
		_, err := conn.Exec(ctx, view.Query)
		if err != nil {
			return err
		}
	}
	return nil
}

func executePolicyQuery(ctx context.Context, conn *pgxpool.Conn, query config.PolicyQueries) (*PolicyResult, error) {
	data, err := conn.Query(ctx, query.Query)
	if err != nil {
		return nil, err
	}

	result := &PolicyResult{
		Name:    query.Name,
		Columns: make([]string, 0),
		Data:    make([][]interface{}, 0),
		Passed:  true,
	}
	for _, fd := range data.FieldDescriptions() {
		result.Columns = append(result.Columns, string(fd.Name))
	}

	for data.Next() {
		values, err := data.Values()
		if err != nil {
			return nil, err
		}
		result.Data = append(result.Data, values)
	}
	result.Passed = len(result.Data) == 0 && !query.Invert
	return result, nil
}

func createPolicyOutput(output string, result *PolicyExecutionResult) error {
	f, err := afero.NewOsFs().Open(output)
	defer f.Close()
	if err != nil {
		return err
	}

	data, err := json.Marshal(&result)
	if err != nil {
		return err
	}
	if _, err := f.Write(data); err != nil {
		return err
	}
	return nil
}
