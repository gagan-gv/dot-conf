package grpc

import (
	"context"
	"dot_conf/proto"
	"dot_conf/services"
)

type Config struct {
	proto.ConfigServiceServer
	config services.IConfigService
}

func NewConfigRpc() *Config {
	return &Config{
		config: services.NewConfigService(),
	}
}

func (c *Config) Fetch(ctx context.Context, request *proto.ConfigRequest) (*proto.ConfigResponse, error) {
	return c.config.Fetch(ctx, request)
}
