package helpers

import (
	"testing"
)

func TestFileExists(t *testing.T) {
	equal(t, FileExists("file.go"), true)
	equal(t, FileExists("file1.go"), false)
}
