package main

import (
	c "jarvis/config"
	server "jarvis/in_bound/http"
	"jarvis/out_bound/grpc"
)

func main() {
	conf := c.ReadConfig()
	grpc.CreateClient(conf.Gpt2.Host, conf.Gpt2.Port)
	server.Run(conf.Server.Host, conf.Server.Port)
}
