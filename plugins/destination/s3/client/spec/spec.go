package spec

import (
	"errors"
	"fmt"
	"path"
	"path/filepath"
	"slices"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/cloudquery/filetypes/v4"
	"github.com/cloudquery/plugin-sdk/v4/configtype"
)

const (
	varFormat      = "{{FORMAT}}"
	varTable       = "{{TABLE}}"
	varUUID        = "{{UUID}}"
	varYear        = "{{YEAR}}"
	varMonth       = "{{MONTH}}"
	varDay         = "{{DAY}}"
	varHour        = "{{HOUR}}"
	varMinute      = "{{MINUTE}}"
	varSyncID      = "{{SYNC_ID}}"
	varTableHyphen = "{{TABLE_HYPHEN}}"
)

type Spec struct {
	filetypes.FileSpec

	// Bucket where to sync the files.
	Bucket string `json:"bucket,omitempty" jsonschema:"required,minLength=1"`

	// Region where bucket is located.
	Region string `json:"region,omitempty" jsonschema:"required,minLength=1"`

	// [Local profile](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html) to use to authenticate this account with.
	// Should be set to the name of the profile.
	//
	// For example, with the following credentials file:
	//
	//   ```ini copy
	//   [default]
	//   aws_access_key_id=xxxx
	//   aws_secret_access_key=xxxx
	//
	//   [user1]
	//   aws_access_key_id=xxxx
	//   aws_secret_access_key=xxxx
	//   ```
	//
	// `local_profile` should be set to either `default` or `user1`.
	LocalProfile string `json:"local_profile,omitempty" jsonschema:"example=my_aws_profile"`

	Credentials *Credentials `json:"credentials,omitempty"`

	//    Path to where the files will be uploaded in the above bucket, for example `path/to/files/{{TABLE}}/{{UUID}}.parquet`.
	//    The path supports the following placeholder variables:
	//
	// - `{{TABLE}}` will be replaced with the table name
	// - `{{TABLE_HYPHEN}}` will be replaced with the table name with hyphens instead of underscores
	// - `{{FORMAT}}` will be replaced with the file format, such as `csv`, `json` or `parquet`. If compression is enabled, the format will be `csv.gz`, `json.gz` etc.
	// - `{{UUID}}` will be replaced with a random UUID to uniquely identify each file
	// - `{{YEAR}}` will be replaced with the current year in `YYYY` format
	// - `{{MONTH}}` will be replaced with the current month in `MM` format
	// - `{{DAY}}` will be replaced with the current day in `DD` format
	// - `{{HOUR}}` will be replaced with the current hour in `HH` format
	// - `{{MINUTE}}` will be replaced with the current minute in `mm` format
	//
	// **Note** that timestamps are in `UTC` and will be the current time at the time the file is written, not when the sync started.
	Path string `json:"path,omitempty" jsonschema:"required,pattern=^[^/].*$,example=path/to/files/{{TABLE}}/{{UUID}}.parquet" jsonschema_extras:"errorMessage=value should not start with /"` // other cases (//, ./, ../) are covered in extended part

	// If set to `true`, the plugin will write to one file per table.
	// Otherwise, for every batch a new file will be created with a different `.<UUID>` suffix.
	NoRotate bool `json:"no_rotate,omitempty" jsonschema:"default=false"`

	// When `athena` is set to `true`, the S3 plugin will sanitize keys in JSON columns to be compatible with the Hive Metastore / Athena.
	// This allows tables to be created with a Glue Crawler and then queried via Athena, without changes to the table schema.
	Athena bool `json:"athena,omitempty" jsonschema:"default=false"`

	// Ensure write access to the given bucket and path by writing a test object on each sync.
	// If you are sure that the bucket and path are writable, you can set this to `false` to skip the test.
	TestWrite *bool `json:"test_write,omitempty" jsonschema:"default=true"`

	// This allows you to set the Content Type of objects uploaded to S3. This will override the default the content type set based on the file format
	// "csv": "text/csv"
	// "json": "application/json"
	// "parquet":" "application/vnd.apache.parquet"
	ContentType string `json:"content_type,omitempty" jsonschema:"default="`

	// Endpoint to use for S3 API calls. This is useful for S3-compatible storage services such as MinIO.
	// **Note**: if you want to use path-style addressing, i.e., `https://s3.amazonaws.com/BUCKET/KEY`, `use_path_style` should be enabled, too.
	Endpoint string `json:"endpoint,omitempty"  jsonschema:"default="`

	// Server-side encryption settings.
	ServerSideEncryptionConfiguration *ServerSideEncryptionConfiguration `json:"server_side_encryption_configuration,omitempty"`

	// Allows to use path-style addressing in the `endpoint` option, i.e., `https://s3.amazonaws.com/BUCKET/KEY`.
	// By default, the S3 client will use virtual hosted bucket addressing when possible (`https://BUCKET.s3.amazonaws.com/KEY`).
	UsePathStyle bool `json:"use_path_style,omitempty" jsonschema:"default=false"`

	// Disable TLS verification for requests to your S3 endpoint.
	//
	// This option is intended to be used when using a custom endpoint using the `endpoint` option.
	EndpointSkipTLSVerify bool `json:"endpoint_skip_tls_verify,omitempty" jsonschema:"default=false"`

	ACL string `json:"acl,omitempty" jsonschema:"default="`

	// If set to `true`, the plugin will create empty parquet files with the table headers and data types for those tables that have no data.
	GenerateEmptyObjects bool `json:"write_empty_objects_for_empty_tables,omitempty" jsonschema:"default=false"`

	// Maximum number of items that may be grouped together to be written in a single write.
	//
	// Defaults to `10000` unless `no_rotate` is `true` (will be `0` then).
	BatchSize *int64 `json:"batch_size" jsonschema:"minimum=1,default=10000"`

	// Maximum size of items that may be grouped together to be written in a single write.
	//
	// Defaults to `52428800` (50 MiB) unless `no_rotate` is `true` (will be `0` then).
	BatchSizeBytes *int64 `json:"batch_size_bytes" jsonschema:"minimum=1,default=52428800"`

	// Maximum interval between batch writes.
	//
	// Defaults to `30s` unless `no_rotate` is `true` (will be `0s` then).
	BatchTimeout *configtype.Duration `json:"batch_timeout" jsonschema:"default=30s"`

	// If `true`, will log AWS debug logs, including retries and other request/response metadata. Requires passing `--log-level debug` to the CloudQuery CLI.
	AWSDebug bool `json:"aws_debug,omitempty" jsonschema:"default=false"`

	// Defines the maximum number of times an API request will be retried.
	MaxRetries *int `json:"max_retries,omitempty" jsonschema:"default=3"`

	// Defines the duration between retry attempts.
	MaxBackoff *int `json:"max_backoff,omitempty" jsonschema:"default=30"`

	// Defines the maximum size of each part in the multipart upload.
	PartSize *int64 `json:"part_size,omitempty" jsonschema:"default=5242880"` // 5 MiB
}

type ServerSideEncryptionConfiguration struct {
	// ServerSideEncryptionConfiguration KMS Key ID appended to S3 API calls header. Used in conjunction with server_side_encryption.
	SSEKMSKeyId string `json:"sse_kms_key_id,omitempty" jsonschema:"required, minLength=1"`

	// Server Side Encryption header which declares encryption type in S3 API calls header: x-amz-server-side-encryption.
	ServerSideEncryption types.ServerSideEncryption `json:"server_side_encryption,omitempty" jsonschema:"required,enum=AES256,enum=aws:kms,enum=aws:kms:dsse"`
}

func (s *Spec) SetDefaults() {
	if !strings.Contains(s.Path, varTable) {
		// for backwards-compatibility, default to given path plus /{{TABLE}}.[format].{{UUID}} if
		// no {{TABLE}} value is found in the path string
		s.Path += fmt.Sprintf("/%s.%s", varTable, s.Format)
		if !s.NoRotate {
			s.Path += "." + varUUID
		}
	}
	if s.TestWrite == nil {
		b := true
		s.TestWrite = &b
	}
	if s.BatchSize == nil {
		if s.NoRotate {
			s.BatchSize = ptr(int64(0))
		} else {
			s.BatchSize = ptr(int64(10000))
		}
	}
	if s.BatchSizeBytes == nil {
		if s.NoRotate {
			s.BatchSizeBytes = ptr(int64(0))
		} else {
			s.BatchSizeBytes = ptr(int64(50 * 1024 * 1024)) // 50 MiB
		}
	}
	if s.BatchTimeout == nil {
		if s.NoRotate {
			d := configtype.NewDuration(0)
			s.BatchTimeout = &d
		} else {
			d := configtype.NewDuration(30 * time.Second)
			s.BatchTimeout = &d
		}
	}

	if s.LocalProfile != "" {
		s.Credentials = &Credentials{
			LocalProfile: s.LocalProfile,
		}
	}

	if s.MaxRetries == nil {
		maxRetries := 3
		s.MaxRetries = &maxRetries
	}

	if s.MaxBackoff == nil {
		maxBackoff := 30
		s.MaxBackoff = &maxBackoff
	}
	if s.PartSize == nil {
		maxPartSize := manager.DefaultUploadPartSize
		s.PartSize = &maxPartSize
	}
}

func (s *Spec) Validate() error {
	if len(s.Bucket) == 0 {
		return errors.New("`bucket` is required")
	}
	if len(s.Region) == 0 {
		return errors.New("`region` is required")
	}

	if len(s.Path) == 0 {
		return errors.New("`path` is required")
	}
	if path.IsAbs(s.Path) {
		return errors.New("`path` should not start with a \"/\"")
	}
	if s.Path != path.Clean(s.Path) {
		return errors.New("`path` should not contain relative paths or duplicate slashes")
	}

	if s.GenerateEmptyObjects && s.Format != filetypes.FormatTypeParquet {
		return errors.New("`write_empty_objects_for_empty_tables` can only be used with `parquet` format")
	}

	if s.LocalProfile != "" && (s.Credentials != nil && (s.Credentials.RoleARN != "" || s.Credentials.RoleSessionName != "" || s.Credentials.ExternalID != "" || s.Credentials.LocalProfile != "")) {
		return errors.New("`local_profile` cannot be used with `credentials`")
	}
	if s.NoRotate {
		if strings.Contains(s.Path, varUUID) {
			return fmt.Errorf("`path` should not contain %s when `no_rotate` = true", varUUID)
		}

		if (s.BatchSize != nil && *s.BatchSize > 0) || (s.BatchSizeBytes != nil && *s.BatchSizeBytes > 0) || (s.BatchTimeout != nil && s.BatchTimeout.Duration() > 0) {
			return errors.New("`no_rotate` cannot be used with non-zero `batch_size`, `batch_size_bytes` or `batch_timeout_ms`")
		}
	}

	if !strings.Contains(s.Path, varUUID) && s.batchingEnabled() {
		return fmt.Errorf("`path` should contain %s when using a non-zero `batch_size`, `batch_size_bytes` or `batch_timeout_ms`", varUUID)
	}

	if s.ACL != "" {
		acl := types.ObjectCannedACL(s.ACL)
		cannedACLS := acl.Values()

		if !slices.Contains(cannedACLS, acl) {
			return fmt.Errorf("invalid `acl` value: %s", s.ACL)
		}
	}

	// required for s.FileSpec.Validate call
	err := s.FileSpec.UnmarshalSpec()
	if err != nil {
		return err
	}
	s.FileSpec.SetDefaults()

	return s.FileSpec.Validate()
}

func (s *Spec) ReplacePathVariables(table string, fileIdentifier string, t time.Time, syncID string) string {
	name := strings.ReplaceAll(s.Path, varTable, table)
	if strings.Contains(name, varFormat) {
		e := string(s.Format) + s.Compression.Extension()
		name = strings.ReplaceAll(name, varFormat, e)
	}
	name = strings.ReplaceAll(name, varUUID, fileIdentifier)
	name = strings.ReplaceAll(name, varYear, t.Format("2006"))
	name = strings.ReplaceAll(name, varMonth, t.Format("01"))
	name = strings.ReplaceAll(name, varDay, t.Format("02"))
	name = strings.ReplaceAll(name, varHour, t.Format("15"))
	name = strings.ReplaceAll(name, varMinute, t.Format("04"))
	name = strings.ReplaceAll(name, varSyncID, syncID)
	name = strings.ReplaceAll(name, varTableHyphen, strings.ReplaceAll(table, "_", "-"))
	return filepath.Clean(name)
}

func (s *Spec) PathContainsUUID() bool {
	return strings.Contains(s.Path, varUUID)
}

func (s *Spec) PathContainsSyncID() bool {
	return strings.Contains(s.Path, varSyncID)
}

func (s *Spec) batchingEnabled() bool {
	if s.NoRotate {
		// if that's set we don't allow batching
		return false
	}

	return (s.BatchSize == nil || *s.BatchSize > 0) ||
		(s.BatchSizeBytes == nil || *s.BatchSizeBytes > 0) ||
		(s.BatchTimeout == nil || s.BatchTimeout.Duration() > 0)
}

func ptr[A any](a A) *A {
	return &a
}

func (s *Spec) GetContentType() string {
	if s.ContentType != "" {
		return s.ContentType
	}
	switch {
	case s.Compression == filetypes.CompressionTypeGZip:
		// https://www.iana.org/assignments/media-types/application/gzip
		return "application/gzip"
	case s.Compression != "":
		// https://www.iana.org/assignments/media-types/application/octet-stream
		return "application/octet-stream"
	}

	switch s.Format {
	case "json":
		// https://www.iana.org/assignments/media-types/application/json
		return "application/json"
	case "csv":
		// https://www.iana.org/assignments/media-types/text/csv
		return "text/csv"
	case "parquet":
		// https://www.iana.org/assignments/media-types/application/vnd.apache.parquet
		return "application/vnd.apache.parquet"
	}
	// This is the default content type for all unknown files
	return "application/octet-stream"
}
