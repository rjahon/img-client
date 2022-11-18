package config

type Config struct {
	ServiceName    string
	HTTPPort       string
	HTTPScheme     string
	ImgServiceHost string
	ImgGRPCPort    string
	DefaultOffset  string
	DefaultLimit   string
}

func Load() Config {
	config := Config{}

	config.ServiceName = "img-client"
	config.HTTPPort = ":9001"
	config.HTTPScheme = "http"
	config.ImgServiceHost = "localhost"
	config.ImgGRPCPort = ":9000"
	config.DefaultOffset = "0"
	config.DefaultLimit = "10"

	return config
}
