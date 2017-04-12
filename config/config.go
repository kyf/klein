package config

type Configuration interface {
	Load(string) error
}
