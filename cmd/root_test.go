package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestInit(t *testing.T) {
	c := newRootCmd()
	b := bytes.NewBufferString("")
	c.SetOut(b)
	c.SetArgs([]string{"init", "aws"})
	err := c.Execute()
	if err != nil {
		t.Errorf("init command failed: %v", err)
	}
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(out))
}
