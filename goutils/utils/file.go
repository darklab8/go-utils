package utils

import (
	"bufio"
	"fmt"
	"os"

	"github.com/darklab8/darklab_goutils/goutils/logus"
	"github.com/darklab8/darklab_goutils/goutils/logus/utils_logus"
)

type file struct {
	filepath string

	file  *os.File
	lines []string
}

type fileRead struct {
	file
}

func NewReadFile(filepath string, callback func(*fileRead)) {
	f := &fileRead{file{filepath: filepath}}

	file, err := os.Open(f.filepath)
	f.file.file = file

	utils_logus.Log.CheckFatal(err, "failed to open", logus.FilePath(f.filepath))
	defer f.file.file.Close()

	callback(f)
}

func (f *fileRead) ReadLines() []string {

	scanner := bufio.NewScanner(f.file.file)
	f.lines = []string{}
	for scanner.Scan() {
		f.lines = append(f.lines, scanner.Text())
	}
	return f.lines
}

type fileWrite struct {
	file
}

func NewWriteFile(filepath string, callback func(*fileWrite)) {
	f := &fileWrite{file{filepath: filepath}}

	file, err := os.Create(f.filepath)
	f.file.file = file
	utils_logus.Log.CheckFatal(err, "failed to open ", logus.FilePath(f.filepath))
	defer f.file.file.Close()
	callback(f)
}

func (f *fileWrite) WritelnF(msg string) {
	_, err := f.file.file.WriteString(fmt.Sprintf("%v\n", msg))

	utils_logus.Log.CheckFatal(err, "failed to write string to file")
}
