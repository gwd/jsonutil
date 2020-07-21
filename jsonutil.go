package jsonutil

import (
	"encoding/json"
	"io/ioutil"
)

// Converts v into JSON and saves it into fname.  Returns errors from
// json.Marshal and ioutil.WriteFile unmodified.
func JsonSave(v interface{}, filename string) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, b, 0666)
	if err != nil {
		return err
	}

	return nil
}

// Reads data from filename and unmarshals into v.  Returns errors
// from json.Marshal and ioutil.WriteFile unmodified.
func JsonLoad(v interface{}, filename string) error {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, v)
	if err != nil {
		return err
	}

	return nil
}
