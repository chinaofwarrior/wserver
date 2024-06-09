package config

// HandlerConfig 和 Config 用于解析配置文件
type HandlerConfig struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Enabled bool   `json:"enabled"`
}

type Config struct {
	Handlers []HandlerConfig `json:"handlers"`
}
