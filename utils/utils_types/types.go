package utils_types

import (
	"crypto/md5"
	"embed"
	"fmt"
	"io"
	"io/fs"
	"log"
	"path"
	"path/filepath"
	"strings"
)

type FilePath string

func (f FilePath) ToString() string { return string(f) }

func (f FilePath) Base() FilePath { return FilePath(filepath.Base(string(f))) }

func (f FilePath) Dir() FilePath { return FilePath(filepath.Dir(string(f))) }

func (f FilePath) Join(paths ...string) FilePath {
	paths = append([]string{string(f)}, paths...)
	return FilePath(filepath.Join(paths...))
}

func (f FilePath) JoinEmbed(paths ...string) FilePath {
	paths = append([]string{string(f)}, paths...)
	return FilePath(path.Join(paths...))
}

type RegExp string

type TemplateExpression string

type File struct {
	Relpath   FilePath
	Name      string
	Extension string
	Content   string
}

type GetFilesParams struct {
	EmbeddedFilerName string   // required param
	RootFolder        FilePath // to exclude from RelPath
	IsNotRecursive    bool
	relFolder         FilePath
	AllowedExtensions []string
}

func GetFiles(filesystem embed.FS, params GetFilesParams) []File {
	if len(params.AllowedExtensions) == 0 {
		params.AllowedExtensions = []string{"js", "css", "png", "jpeg", "json"}
	}
	if params.RootFolder == "" {
		params.RootFolder = "."
	}

	files, err := filesystem.ReadDir(params.RootFolder.ToString())
	var result []File
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		if f.IsDir() {
			if params.IsNotRecursive {
				continue
			}
			params.relFolder = params.relFolder.JoinEmbed(f.Name())
			params.RootFolder = params.RootFolder.JoinEmbed(f.Name())
			result = append(result, GetFiles(filesystem, params)...)
		} else {
			splitted := strings.Split(f.Name(), ".")
			var extension string
			if len(splitted) > 0 {
				extension = splitted[len(splitted)-1]
			}

			requested := params.RootFolder.JoinEmbed(f.Name()).ToString()
			content, err := filesystem.ReadFile(requested)
			if err != nil {
				PrintFilesForDebug(filesystem)
				fmt.Println(err.Error(), "failed to read file from embedded fs",
					"requested=", requested,
				)
			}

			isAllowedExtension := false
			for _, allowed := range params.AllowedExtensions {
				if allowed == extension {
					isAllowedExtension = true
				}
			}
			if !isAllowedExtension {
				continue
			}

			result = append(result, File{
				Relpath:   params.relFolder.JoinEmbed(f.Name()),
				Name:      f.Name(),
				Extension: extension,
				Content:   string(content),
			})

		}

	}
	return result
}

func PrintFilesForDebug(filesystem embed.FS) {
	fs.WalkDir(filesystem, ".", func(p string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			st, _ := fs.Stat(filesystem, p)
			r, _ := filesystem.Open(p)
			defer r.Close()

			// Read prefix
			var buf [md5.Size]byte
			n, _ := io.ReadFull(r, buf[:])

			// Hash remainder
			h := md5.New()
			_, _ = io.Copy(h, r)
			s := h.Sum(nil)

			fmt.Printf("%s %d %x %x\n", p, st.Size(), buf[:n], s)
		}
		return nil
	})
}
