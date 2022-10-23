package data

import (
	"fmt"
	"net/http"

	"github.com/hasicorp/go-hclog"
)

type ExchangeRates struct {
	log  hclog.Logger
	rate map[string]float64
}

func NewRate(l hclog.Logger) (*ExchangeRates, error) {
	err := &ExchangeRates{log: l, rate: map[string]float64{}}

	return err, nil
}

func (e *ExchangeRates) getRate() error {
	resp, err := http.DefaultClient.Get("https://www.ecb.europa.eu/status/eurofxref/eurofxref-daily.xml")

	if err != nil {
		return nil
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("expected error ocde 200 got %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	md := &Cubes{}

	xml

}
