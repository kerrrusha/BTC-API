package dataStorage

import (
	"os"

	"github.com/kerrrusha/btc-api/api/internal/utils"
)

type FileAccessable struct {
	Path string
}

func (f *FileAccessable) AccessFileRead() *os.File {
	file, err := os.Open(f.Path)
	utils.CheckForError(err)
	return file
}

func (f *FileAccessable) AccessFileWrite() *os.File {
	file, err := os.Create(f.Path)
	utils.CheckForError(err)
	return file
}
