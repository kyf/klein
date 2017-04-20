package config

type LogicConfig struct {
	LogPrefix   string `yaml:"LogPrefix"`
	LogPath     string `yaml:"LogPath"`
	RpcPort     string `yaml:"RpcPort"`
	RedisHost   string `yaml:"RedisHost"`
	RedisAuth   string `yaml:"RedisAuth"`
	SessionHost string `yaml:"SessionHost"`
}

func (this *LogicConfig) Load(fpath string) error {
	return load(this, fpath)
}
