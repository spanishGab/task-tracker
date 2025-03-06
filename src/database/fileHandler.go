package database

import (
	"fmt"
	"os"
)

const (
	OwnerRWPermission = 0644
)

type FileHandler struct {
	FileName string
	instance *os.File
}

func NewFileHandler(fileName string) *FileHandler {
	return &FileHandler{
		FileName: fileName,
	}
}

func (dbc *FileHandler) Open() error {
	var err error
	if dbc.instance != nil {
		return nil
	}
	if dbc.instance, err = os.OpenFile(dbc.FileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, OwnerRWPermission); err != nil {
		return fmt.Errorf("error while trying to open file: %s", err.Error())
	}
	return nil
}

func (dbc *FileHandler) Close() error {
	if dbc.instance == nil {
		return nil
	}
	if err := dbc.instance.Close(); err != nil {
		return fmt.Errorf("error while trying to close file: %s", err.Error())
	}
	dbc.instance = nil
	return nil
}

func (dbc *FileHandler) Write(p []byte) (int, error) {
	if err := os.WriteFile(dbc.FileName, p, OwnerRWPermission); err != nil {
		return 0, fmt.Errorf("error while trying to write to file: %s", err.Error())
	}
	return len(p), nil
}

func (dbc *FileHandler) Read(p []byte) (int, error) {
	var err error
	var data []byte = p
	data, err = os.ReadFile(dbc.FileName)
	if err != nil {
		return 0, fmt.Errorf("error while trying to read all file data: %s", err.Error())
	}
	return len(data), nil
}
