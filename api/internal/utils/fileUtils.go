package utils

import (
	"errors"
	"os"
	"path/filepath"
)

func FileNotExist(filepath string) bool {
	_, err := os.Stat(filepath)
	return errors.Is(err, os.ErrNotExist)
}

func FileIsEmpty(filepath string) bool {
	fileStat, err := os.Stat(filepath)
	CheckForError(err)
	return fileStat.Size() == 0
}

func GetGoSrcPath() string {
	ex, err := os.Getwd()
	CheckForError(err)
	return filepath.Dir(ex)
}

func GetProjPath(projectName string) string {
	sep := "\\"
	return GetGoSrcPath() + sep + projectName + sep
}
