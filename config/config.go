package config

import (
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
)

type Configuration interface {
	Load(string) error
}

func load(obj interface{}, fpath string) error {
	content, err := ioutil.ReadFile(fpath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(content, obj)
	if err != nil {
		return err
	}

	return nil
}
