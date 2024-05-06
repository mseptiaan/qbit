package config

type GRPCServerConfig struct {
	Network, Address string
}

var DefaultGRPCServerConfig = GRPCServerConfig{
	Network: "tcp",
	Address: ":50000",
}
