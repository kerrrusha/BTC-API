package fileStorage

import (
	"io"
	"os"

	"github.com/kerrrusha/btc-api/api/internal/model/dataStorage"
)

type fileWriter struct {
	*dataStorage.FileAccessable
}

func (writer *fileWriter) Write(content string, append bool) int {
	file := writer.AccessFileWrite()

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	return length
}

func CreateFileWriter(filepath string) *fileWriter {
	return &fileWriter{
		FileAccessable: &dataStorage.FileAccessable{Path: filepath},
	}
}
