package convert

import (
	"io/ioutil"
	"path/filepath"

	"github.com/cloudquery/cloudquery/internal/file"
	hcl "github.com/hashicorp/hcl/v2"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
	"github.com/zclconf/go-cty/cty/function/stdlib"
)

// GetEvalContext returns the hcl eval context with a subset of functions used in terraform
// that can be used when simplifying during conversion.
func GetEvalContext(basePath string) *hcl.EvalContext {
	return &hcl.EvalContext{
		Functions: map[string]function.Function{
			// numeric
			"abs":      stdlib.AbsoluteFunc,
			"ceil":     stdlib.CeilFunc,
			"floor":    stdlib.FloorFunc,
			"log":      stdlib.LogFunc,
			"max":      stdlib.MaxFunc,
			"min":      stdlib.MinFunc,
			"parseint": stdlib.ParseIntFunc,
			"pow":      stdlib.PowFunc,
			"signum":   stdlib.SignumFunc,

			// string
			"chomp":      stdlib.ChompFunc,
			"format":     stdlib.FormatFunc,
			"formatlist": stdlib.FormatListFunc,
			"indent":     stdlib.IndentFunc,
			"join":       stdlib.JoinFunc,
			"split":      stdlib.SplitFunc,
			"strrev":     stdlib.ReverseFunc,
			"trim":       stdlib.TrimFunc,
			"trimprefix": stdlib.TrimPrefixFunc,
			"trimsuffix": stdlib.TrimSuffixFunc,
			"trimspace":  stdlib.TrimSpaceFunc,

			// collections
			"chunklist": stdlib.ChunklistFunc,
			"concat":    stdlib.ConcatFunc,
			"distinct":  stdlib.DistinctFunc,
			"flatten":   stdlib.FlattenFunc,
			"length":    stdlib.LengthFunc,
			"merge":     stdlib.MergeFunc,
			"reverse":   stdlib.ReverseListFunc,
			"sort":      stdlib.SortFunc,

			// encoding
			"csvdecode":  stdlib.CSVDecodeFunc,
			"jsondecode": stdlib.JSONDecodeFunc,
			"jsonencode": stdlib.JSONEncodeFunc,

			// time
			"formatdate": stdlib.FormatDateFunc,
			"timeadd":    stdlib.TimeAddFunc,

			// file
			"file": MakeFileFunc(basePath),
		},
	}
}

func MakeFileFunc(basePath string) function.Function {
	return function.New(&function.Spec{
		Params: []function.Parameter{
			{
				Name: "path",
				Type: cty.String,
			},
		},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			path := args[0].AsString()
			osFs := file.NewOsFs()

			// Allow if the given path is complete or is it relative?
			if _, err := osFs.Stat(path); err != nil {
				path = filepath.Join(basePath, path)
				if _, err = osFs.Stat(path); err != nil {
					err = function.NewArgError(0, err)
					return cty.UnknownVal(cty.String), err
				}
			}

			f, err := osFs.Open(path)
			if err != nil {
				err = function.NewArgError(0, err)
				return cty.UnknownVal(cty.String), err
			}
			defer f.Close()
			src, err := ioutil.ReadAll(f)
			if err != nil {
				err = function.NewArgError(0, err)
				return cty.UnknownVal(cty.String), err
			}
			return cty.StringVal(string(src)), nil
		},
	})
}
