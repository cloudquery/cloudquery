//go:build tools
// +build tools

package main

import (
	_ "github.com/cloudquery/cq-gen"
	_ "github.com/golang/mock/mockgen"
	_ "github.com/golang/mock/mockgen/model"
)
