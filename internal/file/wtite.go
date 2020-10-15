package file

import (
	"fmt"
	"os"
	"time"
)

const timeFormat = "2006-01-02T15:04:05"

func NewFile(ext string) *File {
	return &File{Ext: ext}
}

type File struct {
	Ext string
}

func (f File) Write(b []byte) (n int, err error) {
	fileName := fmt.Sprintf("%s.%s", time.Now().Format(timeFormat), f.Ext)
	path := fmt.Sprint("./files/", fileName)

	file, err := os.Create(path)
	if err != nil {
		return 0, fmt.Errorf("%w: can't create file", err)
	}
	defer file.Close()

	return file.Write(b)
}
