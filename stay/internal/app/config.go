package app


type Config struct {
    Port int
}

func NewAppConfig(port int)  *Config {
    return &Config{Port: port}
}
