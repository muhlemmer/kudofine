package config

import "github.com/muhlemmer/kudofine/internal/api"

type Config struct {
	API            api.Config
	AppName        string
	ExternalDomain string
	ListenAddr     string
}

var DefaultConfig = Config{
	API:            api.DefaultConfig,
	AppName:        "kudofine",
	ExternalDomain: "127.0.0.1",
	ListenAddr:     "127.0.0.1:8888",
}
