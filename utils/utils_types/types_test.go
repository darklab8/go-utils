package utils_types

import (
	"embed"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func GetCurrentFolder() FilePath {
	_, filename, _, _ := runtime.Caller(1)
	directory := filepath.Dir(filename)
	return FilePath(directory)
}

//go:embed testdata/*
var testdata embed.FS

func TestGrabFiles(t *testing.T) {
	files := GetFiles(testdata, GetFilesParams{
		AllowedExtensions: []string{"txt"},
		RootFolder:        FilePath("testdata"),
	})

	var file1 File
	var file2infolder1 File

	for _, file := range files {
		if file.Relpath == FilePath("file1.txt") {
			file1 = file
		} else if file.Relpath == FilePath("folder1").Join("file2.txt") {
			file2infolder1 = file
		} else {
			continue
		}
	}

	assert.NotEmpty(t, file1.Relpath)
	assert.NotEmpty(t, file2infolder1.Relpath)
}
