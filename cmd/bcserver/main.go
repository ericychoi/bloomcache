package main

import (
	"fmt"
	"log"
	"net"

	"github.com/caarlos0/env"
	"github.com/ericychoi/bloomcache"
	"github.com/ericychoi/bloomcache/protobuf"
	"github.com/willf/bloom"

	"google.golang.org/grpc"
)

type config struct {
	Port int `env:"BLOOMCACHE_PORT" envDefault:"58080"`
	M    int `env:"BLOOMCACHE_M"`
	K    int `env:"BLOOMCACHE_K"`
}

func main() {
	cfg := config{}

	err := env.Parse(&cfg)
	if err != nil {
		log.Fatalf("could not parse config: %s\n", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := &bloomcache.Server{BF: bloom.New(uint(cfg.M), uint(cfg.K))}

	gs := grpc.NewServer()
	protobuf.RegisterBloomcacheServer(gs, s)
	log.Printf("bcserver listening on :%d\n", cfg.Port)
	if err := gs.Serve(lis); err != nil {
		log.Fatalf("couldn't serve. err: %s\n", err)
	}
}
