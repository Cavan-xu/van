package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
)

var loaders = map[string]func([]byte, interface{}) error{
	".json": loadFromJson,
	".yaml": loadFromYaml,
	".yml":  loadFromYaml,
}

// LoadConfig load config from file, only support .json .yaml .yml file
func LoadConfig(file string, v interface{}) error {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	loader, ok := loaders[path.Ext(file)]
	if !ok {
		return fmt.Errorf("unrecognized config file type: %s", file)
	}

	return loader(content, v)
}

// load config from json file
func loadFromJson(content []byte, v interface{}) error {
	return json.Unmarshal(content, v)
}

// load config from yaml or yml file
func loadFromYaml(content []byte, v interface{}) error {
	return nil
}
