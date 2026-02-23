package command

type CliCommand struct {
	Name        string
	Description string
	Callback    func(config *Config, args []string)
}

type Config struct {
	Next string
	Prev string
}

func NewConfig() *Config {
	return &Config{}
}
