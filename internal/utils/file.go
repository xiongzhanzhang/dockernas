package utils

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
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

func MakePathReadAble(path string) {
	os.Chmod(path, 0777)
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

// func CheckCaExist(path string) bool {
// 	if IsRunOnWindows() {
// 		return CheckPathByList(path)
// 	}
// 	return IsFileExist(path)
// }

// func CheckPathByList(path string) bool {
// 	strList := strings.Split(path, "/")
// 	name := strList[len(strList)-1]
// 	return CheckFileByList(path[:len(path)-len(name)], name)
// }

// func CheckFileByList(path string, file string) bool {
// 	dirs, err := ioutil.ReadDir(path)
// 	if err != nil {
// 		return false
// 	}

// 	for _, fi := range dirs {
// 		if !fi.IsDir() {
// 			if fi.Name() == file {
// 				return true
// 			}
// 		}
// 	}

// 	return false
// }

func ReadFile(filePath string) string {
	data, error := ioutil.ReadFile(filePath)
	if error != nil {
		panic(error)
	}

	return string(data)
}

func WriteFile(filePath string, data string) {
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
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
		if info != nil && !info.IsDir() {
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

func TryFixCaPathOnWindows(path string) string {
	if IsRunOnWindows() {
		return strings.ReplaceAll(path, "\uf02a", "*")
	}
	return path
}

func GetCaFilePathOnHost(caFileDir string, domain string) (string, string, string) {
	cer := ""
	key := ""
	msg := ""

	for _, prefix := range []string{"", "*.", "\uf02a."} {
		for _, subfix := range []string{".cer", ".crt", "_bundle.crt"} {
			cer, key, msg = tryGetCaFilePathOnHost(caFileDir, domain, prefix, subfix)
			if msg == "" {
				break
			}
		}
	}

	return TryFixCaPathOnWindows(cer), TryFixCaPathOnWindows(key), msg
}

func tryGetCaFilePathOnHost(caFileDir string, domain string, prefix string, subfix string) (string, string, string) {
	if caFileDir == "" {
		return "", "", "ca file dir is not set"
	}
	if domain == "" {
		return "", "", "domain is not set"
	}

	cer := caFileDir + "/" + prefix + domain + subfix
	key := caFileDir + "/" + prefix + domain + ".key"
	if IsFileExist(cer) && IsFileExist(key) {
		return cer, key, ""
	}

	cer = caFileDir + "/" + prefix + domain + "/" + prefix + domain + subfix
	key = caFileDir + "/" + prefix + domain + "/" + prefix + domain + ".key"
	if IsFileExist(cer) && IsFileExist(key) {
		return cer, key, ""
	}

	cer = caFileDir + "/" + prefix + domain + "_ecc/" + prefix + domain + subfix
	key = caFileDir + "/" + prefix + domain + "_ecc/" + prefix + domain + ".key"
	if IsFileExist(cer) && IsFileExist(key) {
		return cer, key, ""
	}

	return "", "", "can't find ca file under " + caFileDir
}
