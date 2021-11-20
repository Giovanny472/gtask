package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Configurator interface {
	Load(namefile string, object interface{}) error
	Save(namefile string, object interface{}) error
}

type ConfigApp struct {
}

func New() Configurator {
	return &ConfigApp{}
}

// загрузка данных json из определенного файла
func (configapp *ConfigApp) Load(namefile string, object interface{}) error {

	// путь + файл
	fullNameFile, err := configapp.getPathFile(namefile)
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
func (configapp *ConfigApp) Save(namefile string, object interface{}) error {

	// путь + файл
	fullNameFile, err := configapp.getPathFile(namefile)
	if err != nil {
		return err
	}

	// загружаем данные json
	jsonbyte, _ := json.MarshalIndent(object, "", " ")

	// сохраняем данные
	return ioutil.WriteFile(fullNameFile, jsonbyte, 0644)
}

func (configapp *ConfigApp) getPathFile(namefile string) (string, error) {

	app, err := os.Executable()
	if err != nil {
		return "", err
	}

	appPathData := strings.Replace(app, FolderBin, FolderData, 1)
	pathDataFile := filepath.Dir(appPathData) + FolderSeparator + namefile

	return pathDataFile, nil
}
