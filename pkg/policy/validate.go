package policy

import (
	"context"
	"fmt"
	"path"

	"github.com/cloudquery/cloudquery/pkg/core/database"

	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v4"
	"github.com/rs/zerolog/log"
	"github.com/thoas/go-funk"
)

type ValidateRequest struct {
	// Policy we want to validate
	Policy *Policy
	// Directory is where policies reside
	Directory string
}

func Validate(ctx context.Context, storage database.Storage, req *ValidateRequest) diag.Diagnostics {
	log.Info().Str("policy", req.Policy.Name).Str("version", req.Policy.Version()).Str("subPath", req.Policy.SubPolicy()).Msg("preparing to run policy")
	loadedPolicy, err := Load(ctx, req.Directory, req.Policy)
	if err != nil {
		return diag.FromError(err, diag.INTERNAL)
	}
	return validatePolicy(ctx, storage, loadedPolicy, nil, "")
}

func validatePolicy(ctx context.Context, storage database.Storage, policy *Policy, identifiers []string, policyPath string) diag.Diagnostics {
	var diags diag.Diagnostics
	if len(policy.Identifiers) > 0 {
		log.Debug().Strs("previous_identifiers", identifiers).Strs("identifiers", policy.Identifiers).Msg("overriding parent policy identifiers")
		identifiers = policy.Identifiers
	}
	if identifiers == nil {
		diags = diags.Add(diag.FromError(fmt.Errorf("policy %s has no identifiers set", policy.Name), diag.USER, diag.WithSeverity(diag.WARNING)))
	}

	if len(policy.Checks) > 0 {
		diags = diags.Add(validateChecks(ctx, storage, policy.Identifiers, policy.Checks, path.Join(policyPath, policy.Name)))
	}
	for _, p := range policy.Policies {
		diags = diags.Add(validatePolicy(ctx, storage, p, identifiers, path.Join(policyPath, policy.Name)))
	}
	return diags
}

func validateChecks(ctx context.Context, storage database.Storage, identifiers []string, checks []*Check, policyPath string) diag.Diagnostics {
	conn, err := pgx.Connect(ctx, storage.DSN())
	if err != nil {
		return diag.FromError(err, diag.DATABASE)
	}
	var diags diag.Diagnostics

	for _, c := range checks {
		columns, dd := getQueryColumns(ctx, conn, c.Query)
		if dd.HasErrors() {
			return diags.Add(dd)
		}
		// Check identifiers exist
		for _, id := range identifiers {
			if funk.InStrings(columns, id) {
				continue
			}
			diags = diags.Add(diag.FromError(fmt.Errorf("check %s is missing identifier %s", path.Join(policyPath, c.Name), id), diag.USER, diag.WithSeverity(diag.WARNING)))
		}
		// Check for cq_meta columns
		if c.Reason == "" {
			if funk.InStrings(columns, "cq-reason") {
				diags = diags.Add(diag.FromError(fmt.Errorf("check %s doesn't define reason in configuration or query", path.Join(policyPath, c.Name)), diag.USER, diag.WithSeverity(diag.WARNING)))
			}
		}
	}
	return diags
}

func getQueryColumns(ctx context.Context, conn *pgx.Conn, query string) ([]string, diag.Diagnostics) {
	pStmt, err := conn.Prepare(ctx, "query", query)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok && pgerrcode.IsSyntaxErrororAccessRuleViolation(pgErr.Code) {
			return nil, diag.FromError(err, diag.USER, diag.WithSeverity(diag.WARNING))
		}
		return nil, diag.FromError(err, diag.DATABASE)
	}

	columns := make([]string, len(pStmt.Fields))
	for i, f := range pStmt.Fields {
		columns[i] = string(f.Name)
	}
	return columns, nil
}
