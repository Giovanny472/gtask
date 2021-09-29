package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

type ManagerJson interface {
	Load(namefile string, object interface{}) error
	Save()
}

type ManJson struct {
}

// загрузка данных json из определенного файла
func (manJson *ManJson) Load(namefile string, object interface{}) error {

	// путь + файл
	fullNameFile, err := manJson.getPathFile(namefile)
	if err != nil {
		return err
	}

	// читаем файл
	jsonFile, err := os.Open(fullNameFile)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	// файл -> []byte
	jsonByte, _ := ioutil.ReadAll(jsonFile)

	// загружаем данных в object
	return json.Unmarshal(jsonByte, &object)
}

// сохранение данных json из определенного файла
func (manJson *ManJson) Save(namefile string, object interface{}) error {

	// путь + файл
	fullNameFile, err := manJson.getPathFile(namefile)
	if err != nil {
		return err
	}

	// загружаем данные json
	jsonbyte, _ := json.MarshalIndent(object, "", " ")

	// сохраняем данные
	return ioutil.WriteFile(fullNameFile, jsonbyte, 0644)
}

func (manJson *ManJson) getPathFile(namefile string) (string, error) {

	app, err := os.Executable()
	if err != nil {
		return "", err
	}
	pathfile := filepath.Dir(app) + namefile
	return pathfile, err
}
