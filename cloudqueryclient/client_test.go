// +build integration

package cloudqueryclient

import (
	"github.com/cloudquery/cloudquery/providers/aws"
	"io/ioutil"
	"testing"
)

func TestAllResources(t *testing.T) {
	tmpFile, err := ioutil.TempFile("", "*.cloudquery.db")
	if err != nil {
		t.Fatal(err)
	}
	if err = tmpFile.Close(); err != nil {
		t.Fatal(err)
	}
	client, err := New("sqlite", tmpFile.Name())
	if err != nil {
		t.Fatal(err)
	}
	err = client.Run("./testdata/config.yml")
	if err != nil {
		t.Fatal(err)
	}

	var images []aws.Image
	tests := []struct {
		name     string
		resource interface{}
		count    int
	}{
		{"ec2.images", &images, 1},
	}

	// test ec2.images
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := client.db.Find(tc.resource)
			if res.RowsAffected != 1 {
				t.Errorf("count(images) should equal 1 but equals %d", res.RowsAffected)
			}
		})
	}

}
