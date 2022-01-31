package ioutil

import (
	"io/ioutil"
	"os"
)

// IsFile returns true if the file exists and is not a directory.
func IsFile(f string) bool {
	fi, err := os.Stat(f)
	if err != nil {
		return false
	}
	return !fi.IsDir()
}

// IsDir returns true if the dir exists and is a directory.
func IsDir(f string) bool {
	fi, err := os.Stat(f)
	if err != nil {
		return false
	}
	return fi.IsDir()
}

// MustNewTmpFile returns a new temp file and a function to remove the file.
func MustNewTmpFile() (*os.File, func()) {
	f, err := ioutil.TempFile("", "")
	if err != nil {
		return nil, func() {}
	}
	return f, func() { _ = os.Remove(f.Name()) }
}

// MustWriteFile writes the content into a file and returns a remover
func MustWriteFile(filename, content string) func() {
	if err := ioutil.WriteFile(filename, []byte(content), 0644); err != nil {
		return func() {}
	}
	return func() {
		os.Remove(filename)
	}
}

// MustNewTmpFile returns a new temp dir and a function to remove the dir.
func MustNewTmpDir() (string, func()) {
	dirname, err := ioutil.TempDir("", "")
	if err != nil {
		return "", func() {}
	}
	return dirname, func() { _ = os.RemoveAll(dirname) }
}
