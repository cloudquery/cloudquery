package specs

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/configtype"
	"github.com/ghodss/yaml"
	"github.com/rs/zerolog/log"
)

type SpecReader struct {
	sourcesMap      map[string]*Source
	destinationsMap map[string]*Destination
	transformersMap map[string]*Transformer

	sourceWarningsMap      map[string]Warnings
	destinationWarningsMap map[string]Warnings
	transformerWarningsMap map[string]Warnings

	Sources      []*Source
	Destinations []*Destination
	Transformers []*Transformer
}

var fileRegex = regexp.MustCompile(`\$\{file:([^}]+)\}`)
var envRegex = regexp.MustCompile(`\$\{([^}]+)\}`)
var timeRegex = regexp.MustCompile(`\$\{time:([^}]+)\}`)

func expandFileConfig(cfg []byte) ([]byte, error) {
	var expandErr error
	cfg = fileRegex.ReplaceAllFunc(cfg, func(match []byte) []byte {
		filename := fileRegex.FindSubmatch(match)[1]
		content, err := os.ReadFile(string(filename))
		if err != nil {
			expandErr = err
			return nil
		}
		if bytes.ContainsAny(content, "\n\r") && json.Valid(content) {
			// Values that should be treated as strings in YAML have leading and trailing quotes already
			// so we remove the one added by strconv.Quote
			quoted := strconv.Quote(string(content))
			return []byte(quoted[1 : len(quoted)-1])
		}
		return content
	})
	return cfg, expandErr
}

// expand environment variables in the format ${ENV_VAR}
func expandEnv(cfg []byte) ([]byte, error) {
	var expandErr error
	cfg = envRegex.ReplaceAllFunc(cfg, func(match []byte) []byte {
		envVar := envRegex.FindSubmatch(match)[1]
		content, ok := os.LookupEnv(string(envVar))
		if !ok {
			expandErr = fmt.Errorf("env variable %s not found", envVar)
			return nil
		}
		return []byte(content)
	})

	return cfg, expandErr
}

func expandTime(cfg []byte) ([]byte, error) {
	var expandErr error
	cfg = timeRegex.ReplaceAllFunc(cfg, func(match []byte) []byte {
		relativeTime := timeRegex.FindSubmatch(match)[1]
		parsedTime, err := configtype.ParseTime(string(relativeTime))
		if err != nil {
			expandErr = errors.Join(errors.New("failed to substitute time"), err)
			return nil
		}
		return []byte(parsedTime.AsTime(time.Now()).Format(time.RFC3339))
	})
	return cfg, expandErr
}

func (r *SpecReader) loadSpecsFromFile(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %w", path, err)
	}

	// support multiple yamls in one file
	// this should work both on Windows and Unix
	normalizedConfig := bytes.ReplaceAll(data, []byte("\r\n"), []byte("\n"))

	sections := bytes.Split(normalizedConfig, []byte("\n---\n"))
	for i, doc := range sections {
		doc, err = stripYamlComments(doc)
		if err != nil {
			return fmt.Errorf("failed to strip yaml comments in file %s (section %d): %w", path, i+1, err)
		}
		doc, err = expandFileConfig(doc)
		if err != nil {
			return fmt.Errorf("failed to expand file variable in file %s (section %d): %w", path, i+1, err)
		}
		doc, err = expandTime(doc)
		if err != nil {
			return fmt.Errorf("failed to expand time variable in file %s (section %d): %w", path, i+1, err)
		}
		doc, err = expandEnv(doc)
		if err != nil {
			return fmt.Errorf("failed to expand environment variable in file %s (section %d): %w", path, i+1, err)
		}
		var s Spec
		if err := SpecUnmarshalYamlStrict(doc, &s); err != nil {
			return fmt.Errorf("failed to unmarshal file %s: %w", path, err)
		}
		switch s.Kind {
		case KindSource:
			source := s.Spec.(*Source)
			if r.sourcesMap[source.Name] != nil {
				return fmt.Errorf("duplicate source name %s", source.Name)
			}
			r.sourceWarningsMap[source.Name] = source.GetWarnings()
			source.SetDefaults()
			if err := source.Validate(); err != nil {
				return fmt.Errorf("failed to validate source %s: %w", source.Name, err)
			}
			if source.Registry == RegistryGitHub {
				log.Warn().
					Str("name", source.Name).
					Str("kind", "source").
					Msg("registry: github is deprecated & will be removed in future releases")
			}
			r.sourcesMap[source.Name] = source
			r.Sources = append(r.Sources, source)
		case KindDestination:
			destination := s.Spec.(*Destination)
			if r.destinationsMap[destination.Name] != nil {
				return fmt.Errorf("duplicate destination name %s", destination.Name)
			}
			r.destinationWarningsMap[destination.Name] = destination.GetWarnings()
			destination.SetDefaults()
			if err := destination.Validate(); err != nil {
				return fmt.Errorf("failed to validate destination %s: %w", destination.Name, err)
			}
			if destination.Registry == RegistryGitHub {
				log.Warn().
					Str("name", destination.Name).
					Str("kind", "destination").
					Msg("registry: github is deprecated & will be removed in future releases")
			}
			r.destinationsMap[destination.Name] = destination
			r.Destinations = append(r.Destinations, destination)
		case KindTransformer:
			transformer := s.Spec.(*Transformer)
			if r.transformersMap[transformer.Name] != nil {
				return fmt.Errorf("duplicate transformer name %s", transformer.Name)
			}
			r.transformerWarningsMap[transformer.Name] = transformer.GetWarnings()
			transformer.SetDefaults()
			if err := transformer.Validate(); err != nil {
				return fmt.Errorf("failed to validate transformer %s: %w", transformer.Name, err)
			}
			if transformer.Registry == RegistryGitHub {
				log.Warn().
					Str("name", transformer.Name).
					Str("kind", "transformer").
					Msg("registry: github is deprecated & will be removed in future releases")
			}
			r.transformersMap[transformer.Name] = transformer
			r.Transformers = append(r.Transformers, transformer)
		default:
			return fmt.Errorf("unknown kind %s", s.Kind)
		}
	}
	return nil
}

func (r *SpecReader) loadSpecsFromDir(path string) error {
	files, err := os.ReadDir(path)
	if err != nil {
		return fmt.Errorf("failed to read directory %s: %w", path, err)
	}
	for _, file := range files {
		if !file.IsDir() && !strings.HasPrefix(file.Name(), ".") &&
			(strings.HasSuffix(file.Name(), ".yml") || strings.HasSuffix(file.Name(), ".yaml")) {
			if err := r.loadSpecsFromFile(filepath.Join(path, file.Name())); err != nil {
				return err
			}
		}
	}
	return nil
}

func (r *SpecReader) validate() error {
	if len(r.Sources) == 0 {
		return errors.New("expecting at least one source")
	}
	if len(r.Destinations) == 0 {
		return errors.New("expecting at least one destination")
	}

	// here we check if source with different versions use the same destination and error out if yes
	var destinationSourceMap = make(map[string]string)
	for _, source := range r.Sources {
		for _, destination := range source.Destinations {
			if r.destinationsMap[destination] == nil {
				return fmt.Errorf("source %s references unknown destination %s", source.Name, destination)
			}
			destinationToSourceKey := fmt.Sprintf("%s-%s", destination, source.Path)
			if destinationSourceMap[destinationToSourceKey] == "" {
				destinationSourceMap[destinationToSourceKey] = source.Path + "@" + source.Version
			} else if destinationSourceMap[destinationToSourceKey] != source.Path+"@"+source.Version {
				return fmt.Errorf("destination %s is used by multiple sources %s with different versions", destination, source.Path)
			}
		}
	}

	var err error
	for _, destination := range r.Destinations {
		if destination.SyncGroupId != "" && destination.WriteMode == WriteModeOverwriteDeleteStale {
			err = errors.Join(err, fmt.Errorf("destination %s: sync_group_id is not supported with write_mode: %s", destination.Name, destination.WriteMode))
		}
		for _, transformer := range destination.Transformers {
			if r.transformersMap[transformer] == nil {
				err = errors.Join(err, fmt.Errorf("destination %s references unknown transformer %s", destination.Name, transformer))
			}
		}
	}

	return err
}

func (r *SpecReader) relaxedValidate() error {
	if len(r.Sources) == 0 && len(r.Destinations) == 0 {
		return errors.New("expecting at least one source or destination")
	}

	var err error
	for _, destination := range r.Destinations {
		if destination.SyncGroupId != "" && destination.WriteMode == WriteModeOverwriteDeleteStale {
			err = errors.Join(err, fmt.Errorf("destination %s: sync_group_id is not supported with write_mode: %s", destination.Name, destination.WriteMode))
		}
	}

	return err
}

func (r *SpecReader) GetSourceByName(name string) *Source {
	return r.sourcesMap[name]
}

func (r *SpecReader) GetDestinationByName(name string) *Destination {
	return r.destinationsMap[name]
}

func (r *SpecReader) GetSourceWarningsByName(name string) Warnings {
	return r.sourceWarningsMap[name]
}

func (r *SpecReader) GetDestinationWarningsByName(name string) Warnings {
	return r.destinationWarningsMap[name]
}

func (r *SpecReader) GetTransformerWarningsByName(name string) Warnings {
	return r.transformerWarningsMap[name]
}

func (r *SpecReader) GetDestinationNamesForSource(name string) []string {
	var destinations []string
	source := r.sourcesMap[name]
	for _, destinationName := range source.Destinations {
		if slices.Contains(source.Destinations, destinationName) {
			destinations = append(destinations, r.destinationsMap[destinationName].Name)
		}
	}
	return destinations
}

func NewSpecReader(paths []string) (*SpecReader, error) {
	reader, err := newSpecReader(paths)
	if err != nil {
		return nil, err
	}

	if err := reader.validate(); err != nil {
		return nil, err
	}

	return reader, nil
}

func NewRelaxedSpecReader(paths []string) (*SpecReader, error) {
	reader, err := newSpecReader(paths)
	if err != nil {
		return nil, err
	}

	if err := reader.relaxedValidate(); err != nil {
		return nil, err
	}

	return reader, nil
}

func newSpecReader(paths []string) (*SpecReader, error) {
	reader := &SpecReader{
		sourcesMap:             make(map[string]*Source),
		destinationsMap:        make(map[string]*Destination),
		transformersMap:        make(map[string]*Transformer),
		Sources:                make([]*Source, 0),
		Destinations:           make([]*Destination, 0),
		Transformers:           make([]*Transformer, 0),
		sourceWarningsMap:      make(map[string]Warnings),
		destinationWarningsMap: make(map[string]Warnings),
		transformerWarningsMap: make(map[string]Warnings),
	}
	for _, path := range paths {
		file, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		fileInfo, err := file.Stat()
		if err != nil {
			file.Close()
			return nil, err
		}
		file.Close()
		if fileInfo.IsDir() {
			if err := reader.loadSpecsFromDir(path); err != nil {
				return nil, err
			}
		} else {
			if err := reader.loadSpecsFromFile(path); err != nil {
				return nil, err
			}
		}
	}

	return reader, nil
}

// strip yaml comments from the given yaml document by converting to JSON and back :)
func stripYamlComments(b []byte) ([]byte, error) {
	// replace placeholder variables with valid yaml, otherwise it cannot be parsed
	// in some cases. Short of writing our own yaml parser to remove comments,
	// this seems like the best we can do.
	// We replace placeholder variables with random numbers, because numbers in quotes
	// will then remain quoted in the final yaml. If we replace with strings, they will
	// be unquoted in the final yaml.
	r := rand.New(rand.NewSource(1))
	placeholders := map[string]string{}
	b = envRegex.ReplaceAllFunc(b, func(match []byte) []byte {
		content := envRegex.FindSubmatch(match)[1]
		k := strconv.Itoa(r.Int())
		for bytes.Contains(content, []byte(k)) {
			k = strconv.Itoa(r.Int())
		}
		placeholders[k] = string(content)
		return []byte(k)
	})
	j, err := yaml.YAMLToJSON(b)
	if err != nil {
		return nil, err
	}
	b, err = yaml.JSONToYAML(j)
	if err != nil {
		return nil, err
	}
	// place back placeholder variables
	for k, v := range placeholders {
		b = bytes.ReplaceAll(b, []byte(k), []byte(fmt.Sprintf("${%s}", v)))
	}
	return b, nil
}
