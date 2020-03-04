package model

import (
	"encoding/json"
	"io/ioutil"
)

func ReadJSON(fileName string, v interface{}) error {

	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, v)
	if err != nil {
		return err
	}

	return nil
}
