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

	return GetObjFromJson(string(data), obj)
}

func GetObjFromJson(data string, obj any) any {
	err := json.Unmarshal([]byte(data), obj)
	if err != nil {
		log.Print("unmarshal json error")
		log.Println(err)
		return nil
	}

	return obj
}

func GetJsonFromObj(obj any) string {
	json_str, err := json.MarshalIndent(obj, "", "    ")
	if err != nil {
		log.Panicln(err)
		panic(err)
	}

	return string(json_str)
}
