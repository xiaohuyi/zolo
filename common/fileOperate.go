package common

import (
	"fmt"
	"os"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func ReadFile(path string) []byte {
	fileInfo, _ := os.Stat(path)
	fileSize := fileInfo.Size()
	f, err := os.OpenFile((path), os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	buffer := make([]byte, fileSize)
	f.Read(buffer)
	f.Close()
	return buffer
}
