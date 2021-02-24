package client

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cloudquery/cloudquery/database"
	"github.com/cloudquery/cloudquery/plugin"
	"github.com/olekukonko/tablewriter"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"strings"
)

type PolicyConfig struct {
	Views []struct {
		Name  string
		Query string
	}
	Queries []struct {
		Name   string
		Invert bool
		Query  string
	}
}

type Client struct {
	driver  string
	dsn     string
	verbose string
	db      *database.Database
	runSelf bool
	config  Config
}

func New(driver string, dsn string) (*Client, error) {
	client := Client{
		driver:  driver,
		dsn:     dsn,
	}
	return &client, nil
}

func (c *Client) Run(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("%s doesn't exist. you can create one via 'gen config' command", path)
		}
		return err
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	config := Config{}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return err
	}
	// New code
	manager := plugin.GetManager()
	// Shutdown all plugins once we finish our run
	defer plugin.GetManager().Shutdown()
	errGroup, _ := errgroup.WithContext(context.Background())
	for _, provider := range config.Providers {

		if provider.Name == "" {
			log.Error().Msg("provider must contain key 'name' in configuration")
			return errors.New("provider must contain key 'name' in configuration")
		}
		version := provider.Version
		if provider.Version == "" {
			version = "latest"
		}

		log.Debug().Str("provider", provider.Name).Str("version", version).Msg("getting or creating provider")
		cqProvider, err := manager.GetOrCreateProvider(provider.Name, version)
		if err != nil {
			log.Error().Err(err).Str("provider", provider.Name).Str("version", version).Msg("failed to create provider plugin")
			continue
		}
		// create intermediate variable
		provider := provider
		errGroup.Go(func() error {

			log.Info().Str("provider", provider.Name).Str("version", provider.Version).Msg("requesting provider initialize")
			err = cqProvider.Init(c.driver, c.dsn, true)
			if err != nil {
				return err
			}
			d, err := yaml.Marshal(&provider.Rest)
			if err != nil {
				return err
			}
			log.Info().Str("provider", provider.Name).Str("version", provider.Version).Msg("requesting provider fetch")
			err = cqProvider.Fetch(d)
			if err != nil {
				return err
			}
			return nil
		})
	}

	if err := errGroup.Wait(); err != nil {
		return err
	}
	return nil
}


func (c *Client) RunQuery(path string, outputPath string) error {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("%s doesn't exist. you can create one via 'gen policy' command", path)
		}
		return err
	}

	c.db, err = database.Open(c.driver, c.dsn)
	if err != nil {
		return err
	}

	if c.db.Driver == "neo4j" {
		return fmt.Errorf("query command doesn't support neo4j driver yet")
	}

	c.db, err = database.Open(c.driver, c.dsn)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	config := PolicyConfig{}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return err
	}

	err = c.createViews(&config)
	if err != nil {
		return err
	}

	err = c.runQueries(&config, outputPath)
	if err != nil {
		return err
	}
	return nil
}

type QueryResult struct {
	Name          string
	CheckPassed   bool
	ResultHeaders []string
	ResultRows    [][]string
}

func (c *Client) runQueries(config *PolicyConfig, outputPath string) error {
	var f *os.File
	var err error
	if outputPath != "" {
		f, err = os.Create(outputPath)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = f.WriteString("[")
		if err != nil {
			return err
		}
	}

	log.Info().Int("count", len(config.Queries)).Msg("Executing queries")
	for idx, query := range config.Queries {
		queryResult := QueryResult{
			Name:        query.Name,
			CheckPassed: true,
		}
		log.Info().Str("query", query.Name).Msg("Executing query")
		rows, err := c.db.GormDB.Raw(query.Query).Rows()
		if err != nil {
			return err
		}
		columns, err := rows.Columns()
		if err != nil {
			return err
		}
		table := tablewriter.NewWriter(os.Stdout)
		table.SetAutoFormatHeaders(false)
		table.SetHeader(columns)
		nc := len(columns)
		queryResult.ResultHeaders = columns
		prettyRow := make([]string, nc)
		res := make([]sql.NullString, nc)
		resPtrs := make([]interface{}, nc)
		for i := 0; i < nc; i++ {
			resPtrs[i] = &res[i]
		}
		resultsCount := 0
		for rows.Next() {
			err := rows.Scan(resPtrs...)
			if err != nil {
				return err
			}
			resultsCount += 1
			for i, v := range res {
				prettyRow[i] = v.String
			}
			table.Append(prettyRow)
			queryResult.ResultRows = append(queryResult.ResultRows, prettyRow)
		}
		err = rows.Close()
		if err != nil {
			return err
		}
		if resultsCount > 0 && !query.Invert {
			log.Info().Str("query", query.Name).Msg("Check failed. Query returned results.")
			table.Render()
			queryResult.CheckPassed = false
		} else {
			if query.Invert {
				log.Warn().Str("query", query.Name).Msg("Check failed. Query returned no results.")
				queryResult.CheckPassed = false
			} else {
				log.Info().Str("query", query.Name).Msg("Check passed. Query returned no results.")
			}
		}

		if outputPath != "" {
			b, err := json.Marshal(&queryResult)
			if err != nil {
				return err
			}
			outputStr := string(b)
			if idx != len(config.Queries)-1 {
				outputStr = outputStr + ","
			}
			_, err = f.WriteString(outputStr)
			if err != nil {
				return err
			}
		}
	}

	if outputPath != "" {
		_, err = f.WriteString("]")
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Client) createViews(config *PolicyConfig) error {
	log.Info().Int("count", len(config.Views)).Msg("Creating views")
	for _, view := range config.Views {
		log.Info().Str("name", view.Name).Msg("Creating view")
		fmt.Println(view.Query)
		err := c.db.GormDB.Exec(view.Query).Error
		if err != nil {
			if strings.HasPrefix(err.Error(), "table") {
				log.Info().Str("name", view.Name).Msg("table already exist. skipping.")
				continue
			}
			return err
		}
	}
	return nil
}
