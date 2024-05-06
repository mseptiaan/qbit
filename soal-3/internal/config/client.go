package config

import "fmt"

type ServerConfig struct {
	HostPort string
}

type ReverseProxyConfig struct {
	Address string
}

var DefaultServerConfig = ServerConfig{
	HostPort: fmt.Sprintf("localhost%s", DefaultGRPCServerConfig.Address),
}

var DefaultReverseProxyConfig = ReverseProxyConfig{
	Address: ":8080",
}
