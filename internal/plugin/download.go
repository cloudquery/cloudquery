package plugin

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/pkg/errors"
	"github.com/schollz/progressbar/v3"
)

func downloadFile(filepath string, url string) (err error) {

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return errors.Wrapf(err, "failed to create file: %s", filepath)
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return errors.Wrapf(err, "failed to get url: %s", url)
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s. downloading %s", resp.Status, url)
	}
	bar := progressbar.DefaultBytes(
		resp.ContentLength,
		"downloading",
	)
	// Writer the body to file
	_, err = io.Copy(io.MultiWriter(out, bar), resp.Body)
	if err != nil {
		return errors.Wrap(err, "failed to copy body to file")
	}

	return nil
}
