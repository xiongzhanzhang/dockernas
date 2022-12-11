package utils

import (
	"log"
	"os"
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
