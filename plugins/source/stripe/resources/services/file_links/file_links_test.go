package file_links_test

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/resources/services/file_links"
)

func TestFileLinks(t *testing.T) {
	client.MockTestHelper(t, file_links.FileLinks(), client.TestOptions{})
}
