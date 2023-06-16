package logging

import (
	"bufio"
	"io"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func longStr(len int) string {
	b := make([]byte, len)
	for i := 0; i < len; i++ {
		b[i] = byte(65 + (i % 26)) // cycle through letters A to Z
	}
	return string(b)
}

func genLogs(num, lineLen int) string {
	s := make([]string, num)
	for i := 0; i < num; i++ {
		s[i] = longStr(lineLen)
	}
	return strings.Join(s, "\n")
}

func Test_LogReader(t *testing.T) {
	cases := []struct {
		name      string
		text      string
		wantLines []string
		wantErr   bool
	}{
		{
			name: "basic case",
			text: `{"k": "v"}
{"k2": "v2"}`,
			wantErr: false,
			wantLines: []string{
				`{"k": "v"}`,
				`{"k2": "v2"}`,
			}},
		{
			name: "very long line",
			text: longStr(10000000),
			wantLines: []string{
				longStr(logReaderPrefixLen) + "...",
			},
			wantErr: true,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			r := io.NopCloser(strings.NewReader(tc.text))
			lr := NewLogReader(r)
			var gotErr error
			gotLines := make([]string, 0)
			for i := 0; i < len(tc.wantLines)+1; i++ {
				line, err := lr.NextLine()
				if err == io.EOF {
					break
				} else if err != nil {
					gotErr = err
				}
				gotLines = append(gotLines, string(line))
			}
			if gotErr == nil && tc.wantErr {
				t.Fatal("NextLine() was expected to return error, but didn't")
			}
			if len(gotLines) != len(tc.wantLines) {
				t.Fatalf("NextLine() calls got %d lines, want %d", len(gotLines), len(tc.wantLines))
			}
			if diff := cmp.Diff(gotLines, tc.wantLines); diff != "" {
				t.Errorf("NextLine() lines differ from expected. Diff (-got, +want): %s", diff)
			}
		})
	}
}

// we store these package-level variables so that the compiler cannot eliminate the Benchmarks themselves
var (
	bufScannerResult []byte
	logReaderResult  []byte
)

func Benchmark_BufferedScanner(b *testing.B) {
	logs := genLogs(10, 10000)
	bs := bufio.NewScanner(io.NopCloser(strings.NewReader(logs)))
	b.ResetTimer()
	var got []byte
	for n := 0; n < b.N; n++ {
		for bs.Scan() {
			got = bs.Bytes()
		}
	}
	bufScannerResult = got
}

func Benchmark_LogReader(b *testing.B) {
	logs := genLogs(10, 10000)
	lr := NewLogReader(io.NopCloser(strings.NewReader(logs)))
	b.ResetTimer()
	var got []byte
	for n := 0; n < b.N; n++ {
		for {
			line, err := lr.NextLine()
			if err == io.EOF {
				break
			}
			got = line
		}
	}
	logReaderResult = got
}
