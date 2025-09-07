package config

type LoggingConfig struct {
	Level    string `json:"level"`
	Type     string `json:"type"`
	Colorize bool   `json:"colorize"`
}

type KeyGeneratorConfig struct {
	Type     string `json:"type"`
	KeySpace string `json:"keyspace"`
}

type RateLimitCategory struct {
	TotalRequests int `json:"totalRequests"`
	Every         int `json:"every"`
}

type RateLimitsConfig struct {
	Categories map[string]RateLimitCategory `json:"categories"`
}

type StorageConfig struct {
	Type string `json:"type"`
}

type DocumentsConfig struct {
	About string `json:"about"`
}

type Config struct {
	Host                   string             `json:"host"`
	Port                   int                `json:"port"`
	KeyLength              int                `json:"keyLength"`
	MaxLength              int                `json:"maxLength"`
	StaticMaxAge           int                `json:"staticMaxAge"`
	RecompressStaticAssets bool               `json:"recompressStaticAssets"`
	Logging                []LoggingConfig    `json:"logging"`
	KeyGenerator           KeyGeneratorConfig `json:"keyGenerator"`
	RateLimits             RateLimitsConfig   `json:"rateLimits"`
	Storage                StorageConfig      `json:"storage"`
	Documents              DocumentsConfig    `json:"documents"`
}
