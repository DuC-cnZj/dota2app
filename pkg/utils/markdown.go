package utils

import (
	"bytes"
	"regexp"
)

type MDFileType string

const (
	TypeFile  MDFileType = "file"
	TypeImage MDFileType = "image"
)

type File struct {
	FullText string
	Name     string
	Path     string
	Type     MDFileType
}

type FileList []File

func (fl FileList) GetAllPath() []string {
	var allPaths = make([]string, 0, len(fl))
	for _, file := range fl {
		allPaths = append(allPaths, file.Path)
	}

	return allPaths
}

var FileRegexp = regexp.MustCompile(`[!]?\[(.*?)]\((.*?)\)`)

func ParseFiles(markdownText string) FileList {
	submatch := FileRegexp.FindAllSubmatch([]byte(markdownText), -1)
	res := make([]File, 0, len(submatch))
	for _, sub := range submatch {
		if len(sub) == 3 {
			var t MDFileType = TypeFile
			if bytes.Index(sub[0], []byte("!")) != -1 {
				t = TypeImage
			}
			res = append(res, File{
				FullText: string(sub[0]),
				Name:     string(sub[1]),
				Path:     string(sub[2]),
				Type:     t,
			})
		}
	}

	return res
}
