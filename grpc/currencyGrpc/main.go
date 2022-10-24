package main

import (
	"net"
	"os"

	"github.com/hashicorp/go-hclog"
	currency "github.com/upalchowdhury/golangcoding/grpc/currencyGrpc/server"
	"github.com/upalchowdhury/prolog/golangcoding/grpc/productRest/data"
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

	cs := currency.NewCurrency(rates, log)

	currency.RegisterCurrencyServer(gs, cs)

	reflection.Register(gs) // dont use in production

	l, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Error("Unable to listen", "error", err)
		os.Exit(1)
	}

	gs.Serve(l)

}
