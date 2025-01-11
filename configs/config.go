package configs

type Config struct {
	Debug bool   `doc:"enable debug logs" default:"true"`
	Host  string `help:"App host address to bind to" default:"0.0.0.0"`
	Port  int    `help:"App port to bind to" short:"p" default:"9292"`
}
