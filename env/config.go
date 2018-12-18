package env

type Config struct {
	Name        string `yaml:"name"`
	Service     string `yaml:"service"`
	Environment string `yaml:"environment"`
}
