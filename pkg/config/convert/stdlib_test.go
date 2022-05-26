package convert

import (
	"testing"

	"github.com/cloudquery/cloudquery/internal/file"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/zclconf/go-cty/cty"
)

func TestFileFunc(t *testing.T) {
	osFs := file.NewOsFs()
	osFs.SetFSInstance(afero.NewMemMapFs())
	fileFunc := MakeFileFunc("")

	// Create file with content
	testFilePath := "/testfile"
	f, err := osFs.Create(testFilePath)
	assert.NoError(t, err)
	fileContent := "teststring"
	_, err = f.WriteString(fileContent)
	assert.NoError(t, err)
	assert.NoError(t, f.Close())

	val, err := fileFunc.Call([]cty.Value{
		cty.StringVal(testFilePath),
	})
	assert.NoError(t, err)
	assert.Equal(t, fileContent, val.AsString())
}
