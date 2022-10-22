package main

import (
	"net"
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/upalchowdhury/protos/currency"
	"github.com/upalchowdhury/server"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log := hclog.Default()

	gs := grpc.NewServer()

	cs := server.NewCurrency(log)

	currency.RegisterCurrencyServer(gs, cs)

	reflection.Register(gs) // dont use in production

	l, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Error("Unable to listen", "error", err)
		os.Exit(1)
	}

	gs.Serve(l)

}
