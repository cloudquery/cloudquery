package cloudqueryclient

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/cloudquery/cloudquery/database"
	//"github.com/cloudquery/cloudquery/providers/azure"
	//"github.com/cloudquery/cloudquery/providers/gcp"
	//"github.com/cloudquery/cloudquery/providers/k8s"
	//"github.com/cloudquery/cloudquery/providers/okta"
	//"github.com/cloudquery/cloudquery/providers/provider"
	"github.com/cloudquery/cloudquery/sdk"
	"github.com/olekukonko/tablewriter"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"strings"
)


type Config struct {
	Providers []struct {
		Name string
		Rest map[string]interface{} `yaml:",inline"`
	}
}

type PolicyConfig struct {
	Views []struct {
		Name  string
		Query string
	}
	Queries []struct {
		Name  string
		Invert bool
		Query string
	}
}

type Client struct {
	driver string
	dsn string
	verbose string
	db     *database.Database
	config Config
	log    *zap.Logger
}



func New(driver string, dsn string, verbose bool) (*Client, error) {
	client := Client{
		driver: driver,
		dsn: dsn,
	}
	var err error
	client.db, err = database.Open(driver, dsn)
	if err != nil {
		return nil, err
	}

	zapLogger, err := sdk.NewLogger(verbose)
	client.log = zapLogger
	if err != nil {
		return nil, err
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
	for _, provider := range config.Providers {
		if provider.Name == "" {
			return fmt.Errorf("provider must contain key: name")
		}
		//if !ProviderMap[provider.Name] {
		//	return fmt.Errorf("provider %s is not supported\n", provider.Name)
		//}

		p, err:= sdk.GetProviderPluginClient("./" + provider.Name)
		if err != nil {
			return err
		}

		err = p.Init(c.driver, c.dsn, true)
		if err != nil {
			return err
		}

		d, err := yaml.Marshal(&provider.Rest)
		if err != nil {
			return err
		}

		err = p.Fetch(d)
		if err != nil {
			return err
		}
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

	if c.db.Driver == "neo4j" {
		return fmt.Errorf("query command doesn't support neo4j driver yet")
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
	Name string
	CheckPassed bool
	ResultHeaders []string
	ResultRows [][]string
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



	c.log.Info("Executing queries", zap.Int("count", len(config.Queries)))
	for idx, query := range config.Queries {
		queryResult := QueryResult{
			Name:        query.Name,
			CheckPassed: true,
		}
		c.log.Info("Executing query", zap.String("name", query.Name))
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
		if resultsCount > 0 && !query.Invert{
			c.log.Info("Check failed. Query returned results.", zap.String("name", query.Name), zap.Int("count", resultsCount))
			table.Render()
			queryResult.CheckPassed = false
		} else {
			if query.Invert {
				c.log.Info("Check failed. Query returned no results.", zap.String("name", query.Name))
				queryResult.CheckPassed = false
			} else {
				c.log.Info("Check passed. Query returned no results.", zap.String("name", query.Name))
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
	c.log.Info("Creating views", zap.Int("count", len(config.Views)))
	for _, view := range config.Views {
		c.log.Info("Creating view", zap.String("name", view.Name))
		fmt.Println(view.Query)
		err := c.db.GormDB.Exec(view.Query).Error
		if err != nil {
			if strings.HasPrefix(err.Error(), "table") {
				c.log.Info("table already exist. skipping.", zap.String("name", view.Name))
				continue
			}
			return err
		}
	}

	return nil
}
