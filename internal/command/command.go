package command

type CliCommand struct {
	Name        string
	Description string
	Callback    func(config *Config, location string)
}

type Config struct {
	Next string
	Prev string
}

func NewConfig() *Config {
	return &Config{}
}
