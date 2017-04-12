package config

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type ConnectorConfig struct {
	LogPrefix string `yaml:"logprefix"`
	LogPath   string `yaml:"logpath"`
	TcpPort   string `yaml:"tcpport"`
	HttpPort  string `yaml:"httpport"`
}

func (this *ConnectorConfig) Load(fpath string) error {
	content, err := ioutil.ReadFile(fpath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(content, this)
	if err != nil {
		return err
	}
}
