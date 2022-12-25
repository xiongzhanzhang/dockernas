package utils

import (
	"log"
	"os"
	"path/filepath"
)

func CheckCreateDir(path string) {
	_, err := os.Stat(path)
	if err != nil {
		log.Println(err)
		error := os.MkdirAll(path, os.ModePerm)
		if error != nil {
			log.Println(error)
			panic(error)
		}
	}
}

func IsFileExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	return true
}

func WriteFile(filePath string, data string) {
	f, err := os.OpenFile(filePath, os.O_RDONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Println("open file error :", err)
		panic(err)
	}

	defer f.Close()
	_, err = f.WriteString(data)
	if err != nil {
		log.Println(err)
		panic(err)
	}
}

func GetDirectorySize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size, err
}
