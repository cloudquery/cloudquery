package testdata

import (
	_ "embed"
)

//go:embed fixtures_gen.json
var OpenAPIFixtures []byte

//go:embed spec3.json
var OpenAPISpec []byte
