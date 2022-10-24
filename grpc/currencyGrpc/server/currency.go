package server

import (
	"context"
	"time"

	hclog "github.com/hashicorp/go-hclog"
	protos "github.com/upalchowdhury/golangcoding/grpc/currencyGrpc/protos/currency"
	data "github.com/upalchowdhury/prolog/golangcoding/grpc/productRest/data"
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
	for {
		err := src.Send(&protos.RateResponse{Rate: 12.1})
		if err != nil {
			return err
		}
		time.Sleep(5 * time.Second)
	}
}
