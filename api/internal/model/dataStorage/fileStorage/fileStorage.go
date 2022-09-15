package fileStorage

type FileStorage struct {
	*fileReader
	*fileWriter
}

func CreateFileStorage(filepath string) *FileStorage {
	return &FileStorage{
		CreateFileReader(filepath),
		CreateFileWriter(filepath),
	}
}
