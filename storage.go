package main

import (
	"encoding/json"
	"os"
)

type Storage[T any] struct {
	FileName string
}

func NewStorage[T any](fileName string) *Storage[T] {
	return &Storage[T]{FileName: fileName}
}

// save method
func (s *Storage[T]) save(data T) error {
	fileData, err := json.MarshalIndent(data, "", "    ")

	if err != nil{
		return nil
	}

	return os.WriteFile(s.FileName, fileData, 0644)

}

