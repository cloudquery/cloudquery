package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/cloudquery/cq-provider-sdk/cqproto"
	"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/hashicorp/go-hclog"

	env "github.com/caitlinelfring/go-env-default"
)

type Configuration struct {
	requestedFormat cqproto.ConfigFormat
}

type TestClient struct {
	L hclog.Logger
}

type fuzzConfig struct {
	numberOfResources int
	numberOfRelations int
	relationDepth     int
	minFetchDelay     int
	maxFetchDelay     int
}

type exampleData struct {
	Bool bool
	Int  int
	Str  string
	Json interface{}
}

const (
	defaultNumberOfResources           = 200
	defaultNumberOfRelations           = 10
	defaultRelationDepth               = 5
	defaultMinFetchDelayInMilliseconds = 100
	defaultMaxFetchDelayInMilliseconds = 60000
)

var (
	version = "Development"
)

func (t TestClient) Logger() hclog.Logger {
	return t.L
}

func (c Configuration) Example() string {
	switch c.requestedFormat {
	case cqproto.ConfigHCL:
		return `configuration {}`
	default:
		return ``
	}
}

func (c Configuration) Format() cqproto.ConfigFormat {
	return c.requestedFormat
}

func FuzzProvider() *provider.Provider {
	config := getFuzzingConfig()
	resourceMap := make(map[string]*schema.Table)

	for i := 0; i < config.numberOfResources; i++ {
		name := fmt.Sprintf("resource_%d", i)
		table := getTable(name, config, 1)
		resourceMap[name] = &table
	}

	return &provider.Provider{
		Name:    "fuzz",
		Version: version,
		Configure: func(logger hclog.Logger, i interface{}) (schema.ClientMeta, diag.Diagnostics) {
			return &TestClient{L: logger}, nil
		},
		ResourceMap: resourceMap,
		Config: func(f cqproto.ConfigFormat) provider.Config {
			return newConfiguration(f)
		},
		Logger: hclog.NewNullLogger(),
	}
}

func getFuzzingConfig() fuzzConfig {
	numberOfResources := env.GetIntDefault("CQ_FUZZING_NUMBER_OF_RESOURCES", defaultNumberOfResources)
	numberOfRelations := env.GetIntDefault("CQ_FUZZING_NUMBER_OF_RELATIONS", defaultNumberOfRelations)
	relationDepth := env.GetIntDefault("CQ_FUZZING_RELATION_DEPTH", defaultRelationDepth)
	minDelay := env.GetIntDefault("CQ_FUZZING_MIN_FETCH_DELAY", defaultMinFetchDelayInMilliseconds)
	maxDelay := env.GetIntDefault("CQ_FUZZING_MAX_FETCH_DELAY", defaultMaxFetchDelayInMilliseconds)

	config := fuzzConfig{
		numberOfResources: numberOfResources,
		numberOfRelations: numberOfRelations,
		relationDepth:     relationDepth,
		minFetchDelay:     minDelay,
		maxFetchDelay:     maxDelay,
	}

	return config
}

func getRandomInt(min int, max int) int {
	return rand.Intn(max-min) + min
}

func getRandomDuration(min int, max int) time.Duration {
	return time.Duration(getRandomInt(min, max)) * time.Millisecond
}

func getColumns() []schema.Column {
	return []schema.Column{
		{
			Name: "bool",
			Type: schema.TypeBool,
		},
		{
			Name: "int",
			Type: schema.TypeInt,
		},
		{
			Name: "string",
			Type: schema.TypeString,
		},
		{
			Name: "json",
			Type: schema.TypeJSON,
		},
	}
}

func getRelations(parentName string, config fuzzConfig, currentDepth int) []*schema.Table {
	relations := []*schema.Table{}
	if currentDepth > config.relationDepth {
		return relations
	}

	// decrease relations as we go deeper to avoid too many
	for i := 0; i < config.numberOfRelations/currentDepth; i++ {
		name := fmt.Sprintf("%s_%d", parentName, i)
		table := getTable(name, config, currentDepth+1)
		relations = append(relations, &table)
	}

	return relations
}

func getExampleData() exampleData {
	var value interface{}

	json.Unmarshal([]byte(`{"key":"value"}`), &value)
	return exampleData{
		Bool: true,
		Int:  1,
		Str:  "string",
		Json: value,
	}
}

func getTable(name string, config fuzzConfig, currentDepth int) schema.Table {
	return schema.Table{Name: name,
		Resolver:  getResolverFunc(config),
		Columns:   getColumns(),
		Relations: getRelations(name, config, currentDepth),
	}
}

func getResolverFunc(config fuzzConfig) schema.TableResolver {
	return func(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
		meta.Logger().Info("fetching")
		select {
		case <-ctx.Done():
			return nil
		case <-time.After(getRandomDuration(config.minFetchDelay, config.maxFetchDelay)):
			res <- getExampleData()
			return nil
		}
	}
}

func newConfiguration(f cqproto.ConfigFormat) *Configuration {
	return &Configuration{
		requestedFormat: f,
	}
}
