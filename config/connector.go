package config

type ConnectorConfig struct {
	LogPrefix   string `yaml:"LogPrefix"`
	LogPath     string `yaml:"LogPath"`
	TcpPort     string `yaml:"TcpPort"`
	HttpPort    string `yaml:"HttpPort"`
	GrpcHost    string `yaml:"GrpcHost"`
	SessionHost string `yaml:"SessionHost"`
}

func (this *ConnectorConfig) Load(fpath string) error {
	return load(this, fpath)
}
