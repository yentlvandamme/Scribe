package storage

import (
	"io"
	"os"
)

func ReadFromFile(fileName string) (io.ReadCloser, error) {
	return os.Open(fileName)
}

func WriteToFile(fileName string, data []byte) error {
	return os.WriteFile(fileName, data, os.ModePerm)
}
