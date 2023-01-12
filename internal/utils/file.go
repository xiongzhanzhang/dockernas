package utils

import (
	"io"
	"io/ioutil"
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

func GetFileModTimestamp(filePath string) (int64, error) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	modTime := fileInfo.ModTime()
	return modTime.UnixMilli(), nil
}

func IsFileExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	return true
}

func ReadFile(filePath string) string {
	data, error := ioutil.ReadFile(filePath)
	if error != nil {
		panic(error)
	}

	return string(data)
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

func CopyFile(srcFile, destFile string) (int64, error) {
	file1, err := os.Open(srcFile)
	if err != nil {
		return 0, err
	}
	file2, err := os.OpenFile(destFile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer file1.Close()
	defer file2.Close()

	return io.Copy(file2, file1)
}
