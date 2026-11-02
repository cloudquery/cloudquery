package cmd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/cloudquery/plugin-pb-go/pb/plugin/v3"
	"github.com/rs/zerolog/log"
	"github.com/santhosh-tekuri/jsonschema/v6"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// initPlugin is a simple wrapper that will try to validate the spec before actually passing it to Init.
func initPlugin(ctx context.Context, client plugin.PluginClient, spec map[string]any, noConnection bool, syncID string) error {
	_, err := initPluginWithResponse(ctx, client, spec, noConnection, syncID)
	return err
}

func initPluginWithResponse(ctx context.Context, client plugin.PluginClient, spec map[string]any, noConnection bool, syncID string) (*plugin.Init_Response, error) {
	specBytes, err := marshalSpec(spec)
	if err != nil {
		return nil, err
	}

	return client.Init(ctx, &plugin.Init_Request{Spec: specBytes, NoConnection: noConnection, InvocationId: syncID})
}

// validatePluginSpec encompasses spec validation against a plugin obtained
// over gRPC. Schema acquisition and validation are kept lenient so that a
// buggy or out-of-date plugin cannot block validation entirely:
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
	return validateSpecAgainstSchema(schema.GetJsonSchema(), spec)
}

// validateSpecAgainstSchema validates spec against a JSON schema string.
// Mirrors the lenient gRPC-side semantics: an empty or unparseable schema
// is logged and treated as "skip validation". Used both by validatePluginSpec
// (after a successful gRPC GetSpecSchema) and by the Hub-API path in
// validateConfig.
func validateSpecAgainstSchema(jsonSchema string, spec any) error {
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
	c.DefaultDraft(jsonschema.Draft2020)
	c.AssertFormat()

	s, err := jsonschema.UnmarshalJSON(strings.NewReader(jsonSchema))
	if err != nil {
		return nil, err
	}

	if err := c.AddResource("schema.json", s); err != nil {
		return nil, err
	}

	sc, err := c.Compile("schema.json")
	if err != nil {
		var se *jsonschema.SchemaValidationError
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
