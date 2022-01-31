package e2etest

import (
	"encoding/json"
	"io"
	"os"
	"path"
	"strings"
	"testing"
	"text/template"

	"github.com/cloudquery/cloudquery/internal/s3"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

type tfState struct {
	Resources []tfResource
}

type tfResource struct {
	Mode      string
	Type      string
	Instances []tfInstance
}

type tfInstance struct {
	Attributes map[string]interface{}
}

type jsonObject map[string]interface{}

// loadKnownResources will extract known resources from terraform state file.
func loadKnownResources(t *testing.T, filename string) map[string][]string {
	f, err := os.Open(filename)
	require.Nil(t, err)

	var state tfState
	dec := json.NewDecoder(f)
	err = dec.Decode(&state)
	_ = f.Close()
	require.Nil(t, err)

	resources := make(map[string][]jsonObject)
	for _, r := range state.Resources {
		if r.Mode != "managed" {
			continue
		}
		for _, i := range r.Instances {
			resources[r.Type] = append(resources[r.Type], i.Attributes)
		}
	}

	tmpl, err := template.ParseFiles("known_template.yaml")
	require.Nil(t, err)

	r, w := io.Pipe()
	defer r.Close()
	go func() {
		err := tmpl.Execute(w, resources)
		w.CloseWithError(err)
	}()
	yd := yaml.NewDecoder(r)
	var result map[string][]string
	require.Nil(t, yd.Decode(&result))
	return result
}

func possiblySaveFromS3(t *testing.T, filename string) string {
	if !strings.HasPrefix(filename, "s3://") {
		return filename
	}
	// prepare a temporary file
	dir := t.TempDir()
	temp := path.Join(dir, "known.yaml")
	f, err := os.Create(temp)
	require.Nil(t, err)
	defer f.Close()

	// start loading s3 object
	bucket, key, err := parseS3URI(filename)
	require.Nil(t, err)
	svc, err := s3.Session(bucket, "", "")
	require.Nil(t, err)
	body, err := s3.GetObject(svc, bucket, key)
	require.Nil(t, err)
	defer body.Close()

	// write s3 object into temporary file
	_, err = io.Copy(f, body)
	require.Nil(t, err)
	return temp
}
