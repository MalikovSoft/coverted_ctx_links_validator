package xml

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// GetAllOpenCMSNews возвращает слайс всех новостей в заданной директории
func GetAllOpenCMSNews(sourceDirPath string) map[string]*OpenCMSNewsBlocks {
	news := make(map[string]*OpenCMSNewsBlocks)
	files, err := ioutil.ReadDir(sourceDirPath)
	if err != nil {
		os.Mkdir(sourceDirPath, 0777)
		fmt.Printf("Директория \"%s\" не существует!!!.............\n СОЗДАНИЕ ДИРЕКТОРИИ ЗАВЕРШЕНО!!!\n", sourceDirPath)
		return news
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		tmpFile, _ := os.Open(sourceDirPath + file.Name())
		defer tmpFile.Close()
		data, _ := ioutil.ReadAll(tmpFile)
		var tmpNews OpenCMSNewsBlocks
		parseError := xml.Unmarshal(data, &tmpNews)
		if parseError != nil {
			fmt.Println("Файл: ", file.Name(), parseError)
			panic(parseError)
		}
		news[sourceDirPath+file.Name()] = &tmpNews
	}
	return news
}

// ChangeXMLFiles вносит изменения в XML-файлы
func ChangeXMLFiles(changedXMLContentFiles map[string]*OpenCMSNewsBlocks) {
	for filename, xmlFile := range changedXMLContentFiles {
		outPath := filename[0:strings.LastIndex(filename, "/")] + "/output" + filename[strings.LastIndex(filename, "/"):]
		file, _ := xml.Marshal(&xmlFile)
		file = []byte(xml.Header + string(file))
		err := ioutil.WriteFile(outPath, file, 0644)
		if err != nil {
			fmt.Printf(`Error writing file:%s`, filename)
		}
	}
}
