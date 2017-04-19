package config

type LogicConfig struct {
	LogPrefix string `yaml:"LogPrefix"`
	LogPath   string `yaml:"LogPath"`
	RpcPort   string `yaml:"RpcPort"`
}

func (this *LogicConfig) Load(fpath string) error {
	return load(this, fpath)
}
