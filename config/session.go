package config

type SessionConfig struct {
	LogPrefix string `yaml:"LogPrefix"`
	LogPath   string `yaml:"LogPath"`
	RpcPort   string `yaml:"RpcPort"`
}

func (this *SessionConfig) Load(fpath string) error {
	return load(this, fpath)
}
