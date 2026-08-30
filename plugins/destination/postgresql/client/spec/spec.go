package spec

import (
	_ "embed"
	"errors"
	"slices"
	"strings"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/configtype"
	"github.com/invopop/jsonschema"
	"github.com/jackc/pgx/v5/tracelog"
)

const (
	defaultBatchSize        = 10000
	defaultBatchSizeBytes   = 100000000
	defaultBatchTimeout     = 60 * time.Second
	defaultWriteConcurrency = 1
	CQIDColumn              = "_cq_id"
)

type Spec struct {
	// Connection string to connect to the database. This can be a DSN or a URI, for example:
	//
	// - `"user=user password=pass host=localhost port=5432 dbname=mydb sslmode=disable"` DSN format
	// - `"postgres://user:pass@localhost:5432/mydb?sslmode=prefer"` connect with tcp and prefer TLS
	// - `"postgres://user:pass@localhost:5432/mydb?sslmode=disable&search_path=myschema"` connect with tcp, disable TLS and use a custom schema
	ConnectionString string `json:"connection_string,omitempty" jsonschema:"required,minLength=1,example=${POSTGRESQL_CONNECTION_STRING}"`

	// Available: `error`, `warn`, `info`, `debug`, `trace`.
	// Defines what [`pgx`](https://github.com/jackc/pgx) call events should be logged.
	PgxLogLevel LogLevel `json:"pgx_log_level,omitempty" jsonschema:"default=error"`

	// Maximum number of items that may be grouped together to be written in a single write.
	BatchSize int64 `json:"batch_size,omitempty" jsonschema:"minimum=1,default=10000"`

	// Maximum size of items that may be grouped together to be written in a single write.
	BatchSizeBytes int64 `json:"batch_size_bytes,omitempty" jsonschema:"minimum=1,default=100000000"`

	// Maximum interval between batch writes.
	BatchTimeout configtype.Duration `json:"batch_timeout,omitempty"`

	// Number of insert batches to apply concurrently, each on its own connection.
	//
	// The default of 1 writes one batch at a time, which is what the plugin has
	// always done. Raising it lets a database with spare capacity apply several
	// batches at once; `connection_string` should carry a `pool_max_conns` at
	// least this large, and `retry_on_deadlock` should be non-zero, because
	// concurrent upserts over overlapping keys deadlock far more often.
	//
	// Batches in flight together may be applied in any order, so a primary key
	// repeated across batches is no longer guaranteed to resolve to the last one
	// written. Leave this at 1 for sources that emit updates to a row within a
	// single sync.
	WriteConcurrency int64 `json:"write_concurrency,omitempty" jsonschema:"minimum=1,default=1"`

	// Option to create specific indexes to improve deletion performance
	CreatePerformanceIndexes bool `json:"create_performance_indexes,omitempty" jsonschema:"default=false"`

	// Write rows with the PostgreSQL `COPY` protocol instead of `INSERT` statements.
	//
	// `COPY` sends a batch as a single statement rather than one statement per row,
	// which is substantially faster against a remote database. Tables with a
	// primary key are copied into a temporary staging table and merged into the
	// target, since `COPY` has no `ON CONFLICT` of its own.
	//
	// Rows repeating a primary key within one batch are collapsed before the merge,
	// keeping the last one, so a database-side trigger fires once rather than once
	// per row.
	//
	// It has no effect on CockroachDB or CrateDB, which always use `INSERT`.
	UseCopyFrom bool `json:"use_copy_from,omitempty" jsonschema:"default=false"`

	// Optional configuration to enable PgVector embedding support.
	PgVectorConfig *PgVectorConfig `json:"pgvector_config,omitempty"`

	// Number of times to retry a transaction if a deadlock is detected by Postgres.
	RetryOnDeadlock int64 `json:"retry_on_deadlock,omitempty" jsonschema:"default=0"`

	// Optional configuration to connect to [Databricks Lakebase](https://docs.databricks.com/aws/en/oltp),
	// a PostgreSQL-compatible managed database.
	Lakebase *LakebaseSpec `json:"lakebase,omitempty"`

	// Optional configuration to connect to an AWS-managed, PostgreSQL-compatible
	// database with [IAM database authentication](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.IAMDBAuth.html).
	AWSIAMAuth *AWSIAMAuthSpec `json:"aws_iam_auth,omitempty"`
}

// LakebaseSpec enables connecting to Databricks Lakebase, a PostgreSQL-compatible
// OLTP database. When set, the plugin uses the Databricks SDK to generate a
// short-lived OAuth database credential before each new connection and uses it as
// the connection password. The `connection_string` still supplies the host, port,
// database name and user (the service principal client ID), and must use
// `sslmode=require` (or `verify-ca`/`verify-full`); TLS is required and enforced.
type LakebaseSpec struct {
	// The Lakebase database endpoint resource name, in the format
	// `projects/{project_id}/branches/{branch_id}/endpoints/{endpoint_id}`.
	Endpoint string `json:"endpoint" jsonschema:"required,minLength=1"`

	// Databricks workspace host, for example `https://your-workspace.cloud.databricks.com`.
	// If empty, the Databricks SDK resolves it from the `DATABRICKS_HOST` environment
	// variable (or other default Databricks configuration sources).
	Host string `json:"host,omitempty"`

	// Databricks service principal OAuth client ID.
	// If empty, the Databricks SDK resolves it from the `DATABRICKS_CLIENT_ID`
	// environment variable (or other default Databricks configuration sources).
	ClientID string `json:"client_id,omitempty"`

	// Databricks service principal OAuth client secret.
	// If empty, the Databricks SDK resolves it from the `DATABRICKS_CLIENT_SECRET`
	// environment variable (or other default Databricks configuration sources).
	ClientSecret string `json:"client_secret,omitempty"`
}

// AWSIAMAuthService is the AWS database service to authenticate to.
type AWSIAMAuthService string

const (
	// AWSIAMAuthServiceRDS is Amazon RDS and Aurora PostgreSQL.
	AWSIAMAuthServiceRDS AWSIAMAuthService = "rds"
)

// Keep the `service` field's jsonschema enum in sync when adding to this.
var awsIAMAuthServices = []AWSIAMAuthService{AWSIAMAuthServiceRDS}

func supportedAWSIAMAuthServices() string {
	names := make([]string, len(awsIAMAuthServices))
	for i, service := range awsIAMAuthServices {
		names[i] = string(service)
	}
	return strings.Join(names, ", ")
}

// AWSIAMAuthSpec enables IAM database authentication for AWS-managed,
// PostgreSQL-compatible databases. The plugin signs a short-lived token before each
// new connection and uses it as the password; `connection_string` still supplies
// the host, port, database name and user, and must use TLS, which is enforced.
type AWSIAMAuthSpec struct {
	// The AWS database service to authenticate to.
	Service AWSIAMAuthService `json:"service,omitempty" jsonschema:"enum=rds,default=rds"`

	// AWS region the database is in, for example `us-east-1`. If empty, the region
	// is resolved from the standard AWS configuration sources (the `AWS_REGION`
	// environment variable, the shared config file, ...).
	Region string `json:"region,omitempty" jsonschema:"example=us-east-1"`

	// The endpoint the token is signed for, as `host` or `host:port`. If empty, the
	// host and port from `connection_string` are used; if only a host is given, the
	// port from `connection_string` is used. Set this when the plugin connects
	// through a different address than the database endpoint itself, such as an SSH
	// tunnel or a CNAME.
	Endpoint string `json:"endpoint,omitempty" jsonschema:"example=mydb.123456789012.us-east-1.rds.amazonaws.com:5432"`

	// [Local profile](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html)
	// to use to authenticate with. This should be set to the name of the profile.
	LocalProfile string `json:"local_profile,omitempty" jsonschema:"example=my_aws_profile"`

	// If specified, this role is assumed and its credentials are used to sign the
	// authentication token.
	RoleARN string `json:"role_arn,omitempty" jsonschema:"pattern=^(arn(:[^:\n]*){5}([:/].*)?)?$"`

	// If specified, this session name is used when assuming `role_arn`.
	RoleSessionName string `json:"role_session_name,omitempty" jsonschema:"example=my_aws_role_session_name"`

	// If specified, this external ID is used when assuming `role_arn`.
	ExternalID string `json:"external_id,omitempty" jsonschema:"example=external_id"`
}

func (s *Spec) HasPgVectorConfig() bool {
	return s.PgVectorConfig != nil
}

func (s *Spec) HasLakebaseConfig() bool {
	return s.Lakebase != nil
}

func (s *Spec) HasAWSIAMAuthConfig() bool {
	return s.AWSIAMAuth != nil
}

func (s *AWSIAMAuthSpec) ServiceOrDefault() AWSIAMAuthService {
	if s.Service == "" {
		return AWSIAMAuthServiceRDS
	}
	return s.Service
}

func (s *Spec) SetDefaults() {
	if s.BatchSize <= 0 {
		s.BatchSize = defaultBatchSize
	}
	if s.BatchSizeBytes <= 0 {
		s.BatchSizeBytes = defaultBatchSizeBytes
	}
	if s.WriteConcurrency <= 0 {
		s.WriteConcurrency = defaultWriteConcurrency
	}
	if s.BatchTimeout.Duration() <= 0 {
		s.BatchTimeout = configtype.NewDuration(defaultBatchTimeout)
	}
	if s.PgxLogLevel == 0 {
		s.PgxLogLevel = LogLevel(tracelog.LogLevelError)
	}
	if s.PgVectorConfig != nil {
		for i := range s.PgVectorConfig.Tables {
			// Ensure MetadataColumns is an empty slice when unset (not nil)
			if s.PgVectorConfig.Tables[i].MetadataColumns == nil {
				s.PgVectorConfig.Tables[i].MetadataColumns = []string{}
			}
			// Always ensure _cq_id is present in metadata columns
			s.PgVectorConfig.Tables[i].MetadataColumns = ensureCQIDPresent(s.PgVectorConfig.Tables[i].MetadataColumns)
		}
		if s.PgVectorConfig.TextSplitter == nil {
			s.PgVectorConfig.TextSplitter = &PgVectorTextSplitter{
				RecursiveText: PgVectorRecursiveText{
					ChunkSize:    1000,
					ChunkOverlap: 500,
				},
			}
		}
	}
}

func ensureCQIDPresent(metadataColumns []string) []string {
	if slices.Contains(metadataColumns, CQIDColumn) {
		return metadataColumns
	}
	return append([]string{CQIDColumn}, metadataColumns...)
}

func embeddingDimensionsForModel(model string) (int, error) {
	switch model {
	case "text-embedding-3-small":
		return 1536, nil
	case "text-embedding-3-large":
		return 3072, nil
	default:
		return 0, errors.New("`pgvector_config.openai_embedding.model_name` must be one of: text-embedding-3-small, text-embedding-3-large")
	}
}

func (s *Spec) Validate() error {
	if len(s.ConnectionString) == 0 {
		return errors.New("`connection_string` is required")
	}
	if s.Lakebase != nil && len(s.Lakebase.Endpoint) == 0 {
		return errors.New("`lakebase.endpoint` is required when `lakebase` is set")
	}
	// Both set the connection password, so only one of them can be used at a time.
	if s.Lakebase != nil && s.AWSIAMAuth != nil {
		return errors.New("`lakebase` and `aws_iam_auth` are mutually exclusive")
	}
	if s.AWSIAMAuth != nil && !slices.Contains(awsIAMAuthServices, s.AWSIAMAuth.ServiceOrDefault()) {
		return errors.New("`aws_iam_auth.service` must be one of: " + supportedAWSIAMAuthServices())
	}
	if s.PgVectorConfig != nil {
		if len(s.PgVectorConfig.Tables) == 0 {
			return errors.New("`pgvector_config.tables` must contain at least 1 table")
		}
		seenSourceNames := make(map[string]struct{}, len(s.PgVectorConfig.Tables))
		seenTargetNames := make(map[string]struct{}, len(s.PgVectorConfig.Tables))
		for _, tbl := range s.PgVectorConfig.Tables {
			if len(tbl.SourceTableName) == 0 {
				return errors.New("`pgvector_config.tables.source_table_name` is required")
			}
			if len(tbl.TargetTableName) == 0 {
				return errors.New("`pgvector_config.tables.target_table_name` is required")
			}
			if _, ok := seenSourceNames[tbl.SourceTableName]; ok {
				return errors.New("`pgvector_config.tables` contains duplicate source table names: " + tbl.SourceTableName)
			}
			if _, ok := seenTargetNames[tbl.TargetTableName]; ok {
				return errors.New("`pgvector_config.tables` contains duplicate target table names: " + tbl.TargetTableName)
			}
			seenSourceNames[tbl.SourceTableName] = struct{}{}
			seenTargetNames[tbl.TargetTableName] = struct{}{}
			if len(tbl.EmbedColumns) == 0 {
				return errors.New("`pgvector_config.tables.embed_columns` must contain at least 1 column")
			}
		}
		emb := s.PgVectorConfig.OpenAIEmbedding
		if emb.Dimensions <= 0 || len(emb.APIKey) == 0 || len(emb.ModelName) == 0 {
			return errors.New("`pgvector_config.openai_embedding` must have `dimensions`, `api_key`, and `model_name` set")
		}
		// Enforce model support and sync dimensions to the selected model
		dims, err := embeddingDimensionsForModel(emb.ModelName)
		if err != nil {
			return err
		}
		s.PgVectorConfig.OpenAIEmbedding.Dimensions = dims
		if s.PgVectorConfig.TextSplitter != nil {
			ts := s.PgVectorConfig.TextSplitter
			if ts.RecursiveText.ChunkSize <= 0 {
				return errors.New("`pgvector_config.text_splitter.recursive_text.chunk_size` must be > 0")
			}
			if ts.RecursiveText.ChunkOverlap < 0 {
				return errors.New("`pgvector_config.text_splitter.recursive_text.chunk_overlap` must be >= 0")
			}
		}
	}
	return nil
}

func (Spec) JSONSchemaExtend(sc *jsonschema.Schema) {
	sc.Properties.Value("batch_timeout").Default = "60s"
}

//go:embed schema.json
var JSONSchema string

// PgVectorConfig holds configuration for creating embeddings and storing them with pgvector.
type PgVectorConfig struct {
	// Tables to create embeddings for.
	Tables []PgVectorTableConfig `json:"tables,omitempty" jsonschema:"required,minItems=1"`
	// Optional text splitting configuration. If set, all sub-configurations must be set.
	TextSplitter *PgVectorTextSplitter `json:"text_splitter,omitempty"`
	// OpenAI embedding provider configuration. Required if PgVectorConfig is set.
	OpenAIEmbedding OpenAIEmbedding `json:"openai_embedding" jsonschema:"required"`
}

// PgVectorTableConfig defines per-source-table embedding configuration.
// SourceTableName is the base/source table from which text columns will be embedded.
// TargetTableName is the destination table that will be created to store embeddings
// and metadata columns.
type PgVectorTableConfig struct {
	SourceTableName string   `json:"source_table_name" jsonschema:"required,minLength=1"`
	TargetTableName string   `json:"target_table_name" jsonschema:"required,minLength=1"`
	EmbedColumns    []string `json:"embed_columns" jsonschema:"required,minItems=1"`
	MetadataColumns []string `json:"metadata_columns,omitempty"`
}

// PgVectorTextSplitter defines how source text should be split into chunks for embedding.
type PgVectorTextSplitter struct {
	RecursiveText PgVectorRecursiveText `json:"recursive_text" jsonschema:"required"`
}

type PgVectorRecursiveText struct {
	ChunkSize    int `json:"chunk_size" jsonschema:"required,minimum=1"`
	ChunkOverlap int `json:"chunk_overlap" jsonschema:"required,minimum=0"`
}

// OpenAIEmbedding holds embedding provider settings.
type OpenAIEmbedding struct {
	APIKey     string `json:"api_key" jsonschema:"required,minLength=1,title=OpenAI API Key"`
	ModelName  string `json:"model_name" jsonschema:"required,minLength=1"`
	Dimensions int    `json:"dimensions" jsonschema:"minimum=1"`
}

// GetPgVectorTableConfig returns the pgvector table configuration for the given source table name.
func (s *Spec) GetPgVectorTableConfig(tableName string) *PgVectorTableConfig {
	if s.PgVectorConfig == nil {
		return nil
	}
	for i := range s.PgVectorConfig.Tables {
		if s.PgVectorConfig.Tables[i].SourceTableName == tableName {
			return &s.PgVectorConfig.Tables[i]
		}
	}
	return nil
}
