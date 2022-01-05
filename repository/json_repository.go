package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/marcellribeiro/odoscope_attrrange3/entities"
)

type ResponseError struct {
	Message string `json:"message"`
}

func GetDataFromFile(filePath string) (entities.Datasets, error) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}

	data := entities.Datasets{}

	_ = json.Unmarshal([]byte(file), &data)

	return data, nil
}
