package main

type Storage[T any] struct{
	FileName string
}

func NewStorage[T any](fileName string) *Storage[T]{
	return &Storage[T]{FileName: fileName}
}