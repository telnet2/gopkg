package ioutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTmpFile(t *testing.T) {
	f, remove := MustNewTmpFile()
	assert.True(t, IsFile(f.Name()))
	assert.False(t, IsDir(f.Name()))
	remove()
	assert.False(t, IsFile(f.Name()))
}

func TestTmpDir(t *testing.T) {
	f, remove := MustNewTmpDir()
	assert.True(t, IsDir(f))
	assert.False(t, IsFile(f))
	remove()
	assert.False(t, IsDir(f))
	assert.False(t, IsFile(f))
}
