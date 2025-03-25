package database

import (
	"fmt"
	"os"
)

type IFileHandler interface {
	Write([]byte) (int, error)
	Read() ([]byte, error)
}

const (
	OwnerRWPermission = 0644
)

type FileHandler struct {
	FileName string
}

func NewFileHandler(fileName string) *FileHandler {
	return &FileHandler{
		FileName: fileName,
	}
}

func (dbc *FileHandler) Write(p []byte) (int, error) {
	if err := os.WriteFile(dbc.FileName, p, OwnerRWPermission); err != nil {
		return 0, fmt.Errorf("error while trying to write to file: %s", err.Error())
	}
	return len(p), nil
}

func (dbc *FileHandler) Read() ([]byte, error) {
	var err error
	data, err := os.ReadFile(dbc.FileName)
	if err != nil {
		return nil, fmt.Errorf("error while trying to read all file data: %s", err.Error())
	}
	return data, nil
}
