package config

type ConnectorConfig struct {
	LogPrefix string `yaml:"logprefix"`
	LogPath   string `yaml:"logpath"`
	TcpPort   string `yaml:"tcpport"`
	HttpPort  string `yaml:"httpport"`
}

func (this *ConnectorConfig) Load(fpath string) error {
	return load(this, fpath)
}
