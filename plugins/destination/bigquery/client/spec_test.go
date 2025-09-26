package client

import (
	"testing"
	"time"

	"github.com/cloudquery/codegen/jsonschema"
	"github.com/cloudquery/plugin-sdk/v4/configtype"
)

func TestJSONSchema(t *testing.T) {
	jsonschema.TestJSONSchema(t, JSONSchema, []jsonschema.TestCase{
		{
			Name: "empty spec",
			Spec: `{}`,
			Err:  true,
		},
		{
			Name: "spec with project_id",
			Spec: `{"project_id": "value"}`,
			Err:  true,
		},
		{
			Name: "spec with dataset_id",
			Spec: `{"dataset_id": "value"}`,
			Err:  true,
		},
		{
			Name: "spec with project_id and dataset_id",
			Spec: `{"project_id": "foo", "dataset_id": "bar"}`,
		},
		{
			Name: "spec with bool project_id",
			Spec: `{"project_id": true, "dataset_id": "bar"}`,
			Err:  true,
		},
		{
			Name: "spec with null project_id",
			Spec: `{"project_id": null, "dataset_id": "bar"}`,
			Err:  true,
		},
		{
			Name: "spec with int project_id",
			Spec: `{"project_id": 123, "dataset_id": "bar"}`,
			Err:  true,
		},
		{
			Name: "spec with bool batch_size",
			Spec: `{"project_id": "foo", "dataset_id": "bar", "batch_size":false}`,
			Err:  true,
		},
		{
			Name: "spec with null batch_size",
			Spec: `{"project_id": "foo", "dataset_id": "bar", "batch_size":null}`,
			Err:  true,
		},
		{
			Name: "spec with string batch_size",
			Spec: `{"project_id": "foo", "dataset_id": "bar", "batch_size":"str"}`,
			Err:  true,
		},
		{
			Name: "spec with array batch_size",
			Spec: `{"project_id": "foo", "dataset_id": "bar", "batch_size":["abc"]}`,
			Err:  true,
		},
		{
			Name: "spec with proper batch_size",
			Spec: `{"project_id": "foo", "dataset_id": "bar", "batch_size": 7}`,
		},
		{
			Name: "spec with unknown field",
			Spec: `{"project_id": "foo", "dataset_id": "bar", "unknown": "test"}`,
			Err:  true,
		},
	})
}

func TestTextEmbeddingsValidation(t *testing.T) {
	tests := []struct {
		name    string
		spec    Spec
		wantErr bool
	}{
		{
			name: "valid text embeddings spec",
			spec: Spec{
				ProjectID:        "test-project",
				DatasetID:        "test-dataset",
				TimePartitioning: TimePartitioningOptionNone,
				TextEmbeddings: &TextEmbeddingsSpec{
					RemoteModelName: "textembedding-gecko@001",
					Tables: []TableConfig{
						{
							SourceTableName: "source_table",
							TargetTableName: "target_table",
							EmbedColumns:    []string{"content"},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "text embeddings without remote model name",
			spec: Spec{
				ProjectID:        "test-project",
				DatasetID:        "test-dataset",
				TimePartitioning: TimePartitioningOptionNone,
				TextEmbeddings: &TextEmbeddingsSpec{
					Tables: []TableConfig{
						{
							SourceTableName: "source_table",
							TargetTableName: "target_table",
							EmbedColumns:    []string{"content"},
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "text embeddings without tables",
			spec: Spec{
				ProjectID:        "test-project",
				DatasetID:        "test-dataset",
				TimePartitioning: TimePartitioningOptionNone,
				TextEmbeddings: &TextEmbeddingsSpec{
					RemoteModelName: "textembedding-gecko@001",
				},
			},
			wantErr: true,
		},
		{
			name: "text embeddings with empty source table name",
			spec: Spec{
				ProjectID:        "test-project",
				DatasetID:        "test-dataset",
				TimePartitioning: TimePartitioningOptionNone,
				TextEmbeddings: &TextEmbeddingsSpec{
					RemoteModelName: "textembedding-gecko@001",
					Tables: []TableConfig{
						{
							SourceTableName: "",
							TargetTableName: "target_table",
							EmbedColumns:    []string{"content"},
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "text embeddings with empty embed columns",
			spec: Spec{
				ProjectID:        "test-project",
				DatasetID:        "test-dataset",
				TimePartitioning: TimePartitioningOptionNone,
				TextEmbeddings: &TextEmbeddingsSpec{
					RemoteModelName: "textembedding-gecko@001",
					Tables: []TableConfig{
						{
							SourceTableName: "source_table",
							TargetTableName: "target_table",
							EmbedColumns:    []string{},
						},
					},
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Call SetDefaults first to set default values
			tt.spec.SetDefaults()
			err := tt.spec.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Spec.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTextEmbeddingsSetDefaults(t *testing.T) {
	spec := &Spec{
		ProjectID:        "test-project",
		DatasetID:        "test-dataset",
		TimePartitioning: TimePartitioningOptionNone,
		TextEmbeddings: &TextEmbeddingsSpec{
			RemoteModelName: "textembedding-gecko@001",
			Tables: []TableConfig{
				{
					SourceTableName: "source_table",
					TargetTableName: "target_table",
					EmbedColumns:    []string{"content"},
				},
			},
		},
	}

	spec.SetDefaults()

	// Check that TextSplitter defaults are set
	if spec.TextEmbeddings.TextSplitter.RecursiveText.ChunkSize != 1000 {
		t.Errorf("Expected ChunkSize to be 1000, got %d", spec.TextEmbeddings.TextSplitter.RecursiveText.ChunkSize)
	}
	if spec.TextEmbeddings.TextSplitter.RecursiveText.ChunkOverlap != 100 {
		t.Errorf("Expected ChunkOverlap to be 100, got %d", spec.TextEmbeddings.TextSplitter.RecursiveText.ChunkOverlap)
	}

	// Check that _cq_id is added to metadata columns
	if len(spec.TextEmbeddings.Tables[0].MetadataColumns) != 1 {
		t.Errorf("Expected 1 metadata column, got %d", len(spec.TextEmbeddings.Tables[0].MetadataColumns))
	}
	if spec.TextEmbeddings.Tables[0].MetadataColumns[0] != "_cq_id" {
		t.Errorf("Expected _cq_id in metadata columns, got %s", spec.TextEmbeddings.Tables[0].MetadataColumns[0])
	}
}

func TestTimePartitioningValidation(t *testing.T) {
	tests := []struct {
		name    string
		spec    Spec
		wantErr bool
	}{
		{
			name: "valid spec no partitioning",
			spec: Spec{
				ProjectID:                  "test-project",
				DatasetID:                  "test-dataset",
				TimePartitioning:           TimePartitioningOptionNone,
				TimePartitioningExpiration: configtype.NewDuration(0),
			},
			wantErr: false,
		},
		{
			name: "valid spec",
			spec: Spec{
				ProjectID:                  "test-project",
				DatasetID:                  "test-dataset",
				TimePartitioning:           TimePartitioningOptionMonth,
				TimePartitioningExpiration: configtype.NewDuration(0),
			},
			wantErr: false,
		},
		{
			name: "valid spec with partition expiration",
			spec: Spec{
				ProjectID:                  "test-project",
				DatasetID:                  "test-dataset",
				TimePartitioning:           TimePartitioningOptionMonth,
				TimePartitioningExpiration: configtype.NewDuration(time.Hour * 30 * 24),
			},
			wantErr: false,
		},
		{
			name: "no partitioning set but expiration set",
			spec: Spec{
				ProjectID:                  "test-project",
				DatasetID:                  "test-dataset",
				TimePartitioningExpiration: configtype.NewDuration(time.Hour * 30 * 24),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Call SetDefaults first to set default values
			tt.spec.SetDefaults()
			err := tt.spec.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Spec.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
