package ioutil

import (
	iu "io/ioutil"
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
	f, err := iu.TempFile("", "")
	if err != nil {
		return nil, func() {}
	}
	return f, func() { _ = os.Remove(f.Name()) }
}

// MustNewTmpFile returns a new temp dir and a function to remove the dir.
func MustNewTmpDir() (string, func()) {
	dirname, err := iu.TempDir("", "")
	if err != nil {
		return "", func() {}
	}
	return dirname, func() { _ = os.RemoveAll(dirname) }
}
