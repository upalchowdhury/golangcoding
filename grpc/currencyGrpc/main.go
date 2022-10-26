package main

import (
	"net"
	"os"

	"github.com/hashicorp/go-hclog"
	data "github.com/upalchowdhury/golangcoding/grpc/currencyGrpc/data"
	protos "github.com/upalchowdhury/golangcoding/grpc/currencyGrpc/protos/currency"
	server "github.com/upalchowdhury/golangcoding/grpc/currencyGrpc/server"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log := hclog.Default()

	rates, err := data.NewRates(log)

	if err != nil {
		log.Error("Unable to listen", "error", err)
		os.Exit(1)
	}

	gs := grpc.NewServer()

	cs := server.NewCurrency(rates, log)

	protos.RegisterCurrencyServer(gs, cs)

	reflection.Register(gs) // dont use in production

	l, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Error("Unable to listen", "error", err)
		os.Exit(1)
	}

	gs.Serve(l)

}
