package cmd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/cloudquery/cloudquery/cli/v6/internal/specs/v0"
	"github.com/cloudquery/plugin-pb-go/pb/plugin/v3"
	pbSpecs "github.com/cloudquery/plugin-pb-go/specs"
	"github.com/rs/zerolog/log"
	"github.com/santhosh-tekuri/jsonschema/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func CLIRegistryToPbRegistry(registry specs.Registry) pbSpecs.Registry {
	switch registry {
	case specs.RegistryGitHub:
		return pbSpecs.RegistryGithub
	case specs.RegistryLocal:
		return pbSpecs.RegistryLocal
	case specs.RegistryGRPC:
		return pbSpecs.RegistryGrpc
	case specs.RegistryCloudQuery:
		return pbSpecs.RegistryCloudQuery
	default:
		panic(fmt.Sprintf("unknown registry %q", registry.String()))
	}
}

// This converts CLI configuration to a source spec prior to V3 version
// when our spec wasn't decoupled from the over the wire protocol
func CLISourceSpecToPbSpec(spec specs.Source) pbSpecs.Source {
	return pbSpecs.Source{
		Name:                spec.Name,
		Version:             spec.Version,
		Path:                spec.Path,
		Registry:            CLIRegistryToPbRegistry(spec.Registry),
		Tables:              spec.Tables,
		SkipTables:          spec.SkipTables,
		SkipDependentTables: *spec.SkipDependentTables,
		Destinations:        spec.Destinations,
		Spec:                spec.Spec,
		DeterministicCQID:   spec.DeterministicCQID,
	}
}

func CLIWriteModeToPbWriteMode(writeMode specs.WriteMode) pbSpecs.WriteMode {
	switch writeMode {
	case specs.WriteModeAppend:
		return pbSpecs.WriteModeAppend
	case specs.WriteModeOverwrite:
		return pbSpecs.WriteModeOverwrite
	case specs.WriteModeOverwriteDeleteStale:
		return pbSpecs.WriteModeOverwriteDeleteStale
	default:
		panic(fmt.Sprintf("unknown write mode %q", writeMode.String()))
	}
}

func CLIMigrateModeToPbMigrateMode(migrateMode specs.MigrateMode) pbSpecs.MigrateMode {
	switch migrateMode {
	case specs.MigrateModeSafe:
		return pbSpecs.MigrateModeSafe
	case specs.MigrateModeForced:
		return pbSpecs.MigrateModeForced
	default:
		panic(fmt.Sprintf("unknown migrate mode %q", migrateMode.String()))
	}
}

func CLIPkModeToPbPKMode(pkMode specs.PKMode) pbSpecs.PKMode {
	switch pkMode {
	case specs.PKModeCQID:
		return pbSpecs.PKModeCQID
	case specs.PKModeDefaultKeys:
		return pbSpecs.PKModeDefaultKeys
	default:
		panic(fmt.Sprintf("unknown pk mode %q", pkMode.String()))
	}
}

func CLIDestinationSpecToPbSpec(spec specs.Destination) pbSpecs.Destination {
	return pbSpecs.Destination{
		Name:        spec.Name,
		Version:     spec.Version,
		Path:        spec.Path,
		Registry:    CLIRegistryToPbRegistry(spec.Registry),
		WriteMode:   CLIWriteModeToPbWriteMode(spec.WriteMode),
		MigrateMode: CLIMigrateModeToPbMigrateMode(spec.MigrateMode),
		PKMode:      CLIPkModeToPbPKMode(spec.PKMode),
		Spec:        spec.Spec,
	}
}

// initPlugin is a simple wrapper that will try to validate the spec before actually passing it to Init.
func initPlugin(ctx context.Context, client plugin.PluginClient, spec map[string]any, noConnection bool, syncID string) error {
	specBytes, err := marshalSpec(spec)
	if err != nil {
		return err
	}

	_, err = client.Init(ctx, &plugin.Init_Request{Spec: specBytes, NoConnection: noConnection, InvocationId: syncID})
	return err
}

// validatePluginSpec encompasses spec validation only:
//  1. Get spec schema from the plugin.
//     If the call isn't implemented, just skip the validation.
//  2. Validate that the provided JSON schema is valid & can be used for spec validation.
//     If the spec is empty (i.e., the plugin didn't supply the schema) just skip.
//  3. If the schema isn't empty but not valid, print the error message & skip the validation.
//  4. Finally, return the validation result.
func validatePluginSpec(ctx context.Context, client plugin.PluginClient, spec any) error {
	schema, err := client.GetSpecSchema(ctx, &plugin.GetSpecSchema_Request{})
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			// not a gRPC-compatible error
			log.Err(err).Msg("failed to get spec schema")
			return err
		}
		if st.Code() != codes.Unimplemented {
			// unimplemented is OK, treat as empty schema
			log.Err(err).Msg("failed to get spec schema")
			return err
		}
	}

	jsonSchema := schema.GetJsonSchema()
	if len(jsonSchema) == 0 {
		// This will also be true for Unimplemented response (schema = nil => schema.GetJsonSchema() = "")
		log.Info().Msg("empty JSON schema for plugin spec, skipping validation")
		return nil
	}

	sc, err := parseJSONSchema(jsonSchema)
	if err != nil {
		log.Err(err).Msg("failed to parse spec schema, skipping validation")
		return nil
	}

	return sc.Validate(spec)
}

func parseJSONSchema(jsonSchema string) (*jsonschema.Schema, error) {
	c := jsonschema.NewCompiler()
	c.Draft = jsonschema.Draft2020
	c.AssertFormat = true

	if err := c.AddResource("schema.json", strings.NewReader(jsonSchema)); err != nil {
		return nil, err
	}

	sc, err := c.Compile("schema.json")
	if err != nil {
		var se *jsonschema.SchemaError
		if errors.As(err, &se); se != nil && se.Err != nil {
			// We add resource as `file`, but there's none, actually.
			// So, we need to prettify message a bit.
			return nil, fmt.Errorf("jsonschema compilation failed: %w",
				errors.New(strings.Replace(se.Err.Error(), "jsonschema: '' ", "", 1)))
		}
		return nil, err
	}

	return sc, nil
}

func marshalSpec(spec map[string]any) ([]byte, error) {
	// All nil or empty values to be marshaled as null
	if len(spec) == 0 {
		return []byte(`null`), nil
	}

	return json.Marshal(spec)
}
