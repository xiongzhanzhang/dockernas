package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func GetObjFromJsonFile(filePath string, obj any) any {
	data, error := ioutil.ReadFile(filePath)
	if error != nil {
		log.Printf("read %s error", filePath)
		log.Println(error)
		return nil
	}

	err := json.Unmarshal([]byte(data), obj)
	if err != nil {
		log.Print("unmarshal json error")
		log.Println(err)
		return nil
	}

	return obj
}
