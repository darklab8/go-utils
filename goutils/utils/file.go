package utils

import (
	"bufio"
	"fmt"
	"os"

	"github.com/darklab8/darklab_goutils/goutils/logus"
	"github.com/darklab8/darklab_goutils/goutils/logus/utils_logus"
)

type File struct {
	Filepath string
	file     *os.File
	lines    []string
}

func (f *File) OpenToReadF() *File {
	file, err := os.Open(f.Filepath)
	f.file = file

	utils_logus.Log.CheckFatal(err, "failed to open", logus.FilePath(f.Filepath))
	return f
}

func (f *File) Close() {
	f.file.Close()
}

func (f *File) ReadLines() []string {

	scanner := bufio.NewScanner(f.file)
	f.lines = []string{}
	for scanner.Scan() {
		f.lines = append(f.lines, scanner.Text())
	}
	return f.lines
}

func (f *File) FileReadLines() []string {
	f.OpenToReadF()
	defer f.Close()
	return f.ReadLines()
}

func (f *File) CreateToWriteF() *File {
	file, err := os.Create(f.Filepath)
	f.file = file
	utils_logus.Log.CheckFatal(err, "failed to open ", logus.FilePath(f.Filepath))

	return f
}
func (f *File) WritelnF(msg string) {
	_, err := f.file.WriteString(fmt.Sprintf("%v\n", msg))

	utils_logus.Log.CheckFatal(err, "failed to write string to file")
}
