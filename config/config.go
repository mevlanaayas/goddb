package config

type Config struct {
	Port int `yaml:"port"`
}

func (receiver *Config) Init() error {
	// TODO: read config.yaml file
	receiver.Port = 8099
	return nil
}
