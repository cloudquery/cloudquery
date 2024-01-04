package publish

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindMarkdownImages(t *testing.T) {
	cases := []struct {
		name        string
		contents    string
		expect      map[string][]imageReference
		expectError string
	}{
		{
			name: "no images",
			contents: `# Title
`,
			expect: nil,
		},
		{
			name: "basic",
			contents: `# Title
![](image.png)`,
			expect: map[string][]imageReference{"a94a8fe5ccb19ba61c4c0873d391e987982fbbd3:image.png": {{mdFull: "![](image.png)", mdPartial: "image.png", url: ""}}},
		},
		{
			name: "basic with alt",
			contents: `# Title
![Alt text](image.png)`,
			expect: map[string][]imageReference{"a94a8fe5ccb19ba61c4c0873d391e987982fbbd3:image.png": {{mdFull: "![Alt text](image.png)", mdPartial: "image.png", url: ""}}},
		},
		{
			name: "basic with alt and title",
			contents: `# Title
![Alt text](image.png "Title Here")
`,
			expect: map[string][]imageReference{"a94a8fe5ccb19ba61c4c0873d391e987982fbbd3:image.png": {{mdFull: "![Alt text](image.png \"Title Here\")", mdPartial: "image.png", url: ""}}},
		},
		{
			name: "basic with title",
			contents: `# Title
![](image.png "Title Here")
`,
			expect: map[string][]imageReference{"a94a8fe5ccb19ba61c4c0873d391e987982fbbd3:image.png": {{mdFull: "![](image.png \"Title Here\")", mdPartial: "image.png", url: ""}}},
		},
		{
			name: "basic with title or alt, multiple",
			contents: `# Title
![](image.png "Title Here")
More test
![alt](image.png)
`,
			expect: map[string][]imageReference{"a94a8fe5ccb19ba61c4c0873d391e987982fbbd3:image.png": {{mdFull: "![](image.png \"Title Here\")", mdPartial: "image.png", url: ""}, {mdFull: "![alt](image.png)", mdPartial: "image.png", url: ""}}},
		},
		{
			name: "ref",
			contents: `# Title
![Alt text][image-id]

[image-id]: image.png "Optional Title Here"
`,
			expect: nil, // unsupported
		},
		{
			name: "href",
			contents: `# Title
[![Alt text](image.png)](http://example.com/)
`,
			expect: map[string][]imageReference{"a94a8fe5ccb19ba61c4c0873d391e987982fbbd3:image.png": {{mdFull: "![Alt text](image.png)", mdPartial: "image.png", url: ""}}},
		},
		{
			name: "subdir",
			contents: `# Title
![Alt text](assets/images/image.png)

`,
			expect: map[string][]imageReference{"a94a8fe5ccb19ba61c4c0873d391e987982fbbd3:image.png": {{mdFull: "![Alt text](assets/images/image.png)", mdPartial: "assets/images/image.png", url: ""}}},
		},
	}

	tempdir := t.TempDir()
	data := []byte("test")
	require.NoError(t, os.WriteFile(filepath.Join(tempdir, "image.png"), data, 0644))
	require.NoError(t, os.MkdirAll(filepath.Join(tempdir, "assets", "images"), 0755))
	require.NoError(t, os.WriteFile(filepath.Join(tempdir, "assets", "images", "image.png"), data, 0644))

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			out, err := findMarkdownImages(tc.contents, tempdir)
			if tc.expectError != "" {
				require.ErrorContains(t, err, tc.expectError)
				return
			}
			require.NoError(t, err)

			// reset absFile to empty string for comparison
			for k, v := range out {
				for i := range v {
					v[i].absFile = ""
				}
				out[k] = v
			}

			require.EqualValues(t, tc.expect, out)
		})
	}

}
