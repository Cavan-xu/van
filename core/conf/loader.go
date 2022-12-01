package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
)

type ILoader interface {
	LoadConfig(file string, v interface{}) error
}

type JsonLoader struct{}

func (c *JsonLoader) LoadConfig(file string, v interface{}) error {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	return json.Unmarshal(content, v)
}

var loaders = map[string]ILoader{
	".json": &JsonLoader{},
}

// LoadConfig load config from file, only support .json .yaml .yml file
func LoadConfig(file string, v interface{}) error {
	loader, ok := loaders[path.Ext(file)]
	if !ok {
		return fmt.Errorf("unrecognized config file type: %s", file)
	}

	return loader.LoadConfig(file, v)
}
