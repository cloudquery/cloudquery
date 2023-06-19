package specs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/ghodss/yaml"
	"golang.org/x/exp/slices"
)

type SpecReader struct {
	sourcesMap      map[string]*Source
	destinationsMap map[string]*Destination

	Sources      []*Source
	Destinations []*Destination
}

var fileRegex = regexp.MustCompile(`\$\{file:([^}]+)\}`)
var envRegex = regexp.MustCompile(`\$\{([^}]+)\}`)

// escapeExternalContent escapes the given content if it contains newlines or is a JSON object or array, to satisfy YAML requirements
// It will suppress any JSON unmarshalling errors and may return the original content.
func escapeExternalContent(content []byte) ([]byte, error) {
	var isJSON any
	if err := json.Unmarshal(content, &isJSON); err != nil {
		if bytes.ContainsAny(content, "\n\r") {
			return []byte(strconv.Quote(string(content))), nil
		}
		return content, nil
	}

	k := reflect.TypeOf(isJSON).Kind()
	switch k {
	case reflect.Map, reflect.Slice:
		buffer := &bytes.Buffer{}
		encoder := json.NewEncoder(buffer)
		encoder.SetEscapeHTML(false)
		if err := encoder.Encode(string(content)); err != nil {
			return content, err
		}

		return bytes.TrimSuffix(buffer.Bytes(), []byte{'\n'}), nil
	}

	return content, nil
}

func expandFileConfig(cfg []byte) ([]byte, error) {
	var expandErr error
	cfg = fileRegex.ReplaceAllFunc(cfg, func(match []byte) []byte {
		filename := fileRegex.FindSubmatch(match)[1]
		content, err := os.ReadFile(string(filename))
		if err != nil {
			expandErr = err
			return nil
		}
		content, err = escapeExternalContent(content)
		if expandErr == nil {
			expandErr = err
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
		newcontent, err := escapeExternalContent([]byte(content))
		if expandErr == nil {
			expandErr = err
		}
		return newcontent
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
			source.SetDefaults()
			if err := source.Validate(); err != nil {
				return fmt.Errorf("failed to validate source %s: %w", source.Name, err)
			}
			r.sourcesMap[source.Name] = source
			r.Sources = append(r.Sources, source)
		case KindDestination:
			destination := s.Spec.(*Destination)
			if r.destinationsMap[destination.Name] != nil {
				return fmt.Errorf("duplicate destination name %s", destination.Name)
			}
			// We set the default value to 0, so it can be overridden later by plugins' defaults
			destination.SetDefaults(0, 0)
			if err := destination.Validate(); err != nil {
				return fmt.Errorf("failed to validate destination %s: %w", destination.Name, err)
			}
			r.destinationsMap[destination.Name] = destination
			r.Destinations = append(r.Destinations, destination)
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
		return fmt.Errorf("expecting at least one source")
	}
	if len(r.Destinations) == 0 {
		return fmt.Errorf("expecting at least one destination")
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

	return nil
}

func (r *SpecReader) GetSourceByName(name string) *Source {
	return r.sourcesMap[name]
}

func (r *SpecReader) GetDestinationByName(name string) *Destination {
	return r.destinationsMap[name]
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
	reader := &SpecReader{
		sourcesMap:      make(map[string]*Source),
		destinationsMap: make(map[string]*Destination),
		Sources:         make([]*Source, 0),
		Destinations:    make([]*Destination, 0),
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

	if err := reader.validate(); err != nil {
		return nil, err
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
		k := fmt.Sprintf("%d", r.Int())
		for bytes.Contains(content, []byte(k)) {
			k = fmt.Sprintf("%d", r.Int())
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
