package persistentdata

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestReadOrder(t *testing.T) {
	const fn = "the-file"
	af := afero.Afero{Fs: afero.NewMemMapFs()}

	home, err := os.UserHomeDir()
	assert.NoError(t, err)

	err = af.WriteFile(filepath.Join(".", ".cq", fn), []byte("bar"), 0644)
	assert.NoError(t, err)

	// write file in home dir last because "." and $HOME could be the same path
	err = af.WriteFile(filepath.Join(home, ".cq", fn), []byte("foo"), 0644)
	assert.NoError(t, err)

	for i := 0; i < 2; i++ { // run it multiple times so we're sure it's not overwriting current files
		val, isNew, err := New(af, fn, func() string { return "boo" }).Get()
		assert.NoError(t, err)
		assert.False(t, isNew)
		assert.Equal(t, val, "foo")
	}
}

func TestReadDir(t *testing.T) {
	const fn = "the-file"
	af := afero.Afero{Fs: afero.NewMemMapFs()}

	home, err := os.UserHomeDir()
	assert.NoError(t, err)

	// write file in home dir last because "." and $HOME could be the same path
	err = af.WriteFile(filepath.Join(home, ".cq", fn, "inner-file"), []byte("we're in a directory!"), 0644)
	assert.NoError(t, err)

	for i := 0; i < 2; i++ { // run it multiple times so we're sure it's not overwriting current files
		val, isNew, err := New(af, fn, func() string { return "boo" }).Get()
		assert.False(t, isNew)
		assert.Equal(t, val, "")
		assert.Equal(t, ErrIsDirectory, err)
	}
}

func TestRegularRead(t *testing.T) {
	const fn = "the-file"
	af := afero.Afero{Fs: afero.NewMemMapFs()}

	err := af.WriteFile(filepath.Join(".", ".cq", fn), []byte("bar"), 0644)
	assert.NoError(t, err)

	for i := 0; i < 2; i++ { // run it multiple times so we're sure it's not overwriting current files
		val, isNew, err := New(af, fn, func() string { return "boo" }).Get()
		assert.NoError(t, err)
		assert.False(t, isNew)
		assert.Equal(t, val, "bar")
	}
}

func TestGen(t *testing.T) {
	const fn = "the-file"
	af := afero.Afero{Fs: afero.NewMemMapFs()}

	p := New(af, fn, func() string { return "hello" })
	val, isNew, err := p.Get()
	assert.NoError(t, err)
	assert.True(t, isNew)
	assert.Equal(t, val, "hello")

	val, isNew, err = New(af, fn, func() string { return "boo" }).Get()
	assert.NoError(t, err)
	assert.False(t, isNew)
	assert.Equal(t, val, "hello")
}
