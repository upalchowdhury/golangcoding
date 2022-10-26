package server

import (
	"context"
	"io"
	"time"

	hclog "github.com/hashicorp/go-hclog"
	data "github.com/upalchowdhury/golangcoding/grpc/currencyGrpc/data"
	protos "github.com/upalchowdhury/golangcoding/grpc/currencyGrpc/protos/currency"
)

type Currency struct {
	rates *data.ExchangeRates
	log   hclog.Logger
}

func NewCurrency(r *data.ExchangeRates, l hclog.Logger) *Currency {
	return &Currency{r, l}
}

func (c *Currency) GetRate(ctx context.Context, rr *protos.RateRequest) (*protos.RateResponse, error) {
	c.log.Info("Handle getrate", "base", rr.GetBase(), "destination", rr.GetDestination())

	//c.rates.GetRate(rr.GetBase().String(), rr.GetDestination().String())
	rate, err := c.rates.GetRate(rr.GetBase().String(), rr.GetDestination().String())

	if err != nil {
		return nil, err
	}
	return &protos.RateResponse{Rate: rate}, nil
}

func (c *Currency) SubscribeRates(src protos.Currency_SubscribeRatesServer) error {

	go func() {
		for {
			rr, err := src.Recv()
			if err == io.EOF {
				c.log.Info("client has closed connection")
			}

			if err != nil {
				c.log.Error("unable to read from client", "error", err)
				break
			}
			c.log.Info("handle client request", "request", rr)
		}
	}()
	for {
		err := src.Send(&protos.RateResponse{Rate: 12.1})
		if err != nil {
			return err
		}
		time.Sleep(5 * time.Second)
	}
}
