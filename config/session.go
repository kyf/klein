package config

type SessionConfig struct {
	LogPrefix string `yaml:"logprefix"`
	LogPath   string `yaml:"logpath"`
	RpcPort   string `yaml:"rpcport"`
}

func (this *SessionConfig) Load(fpath string) error {
	return load(this, fpath)
}
