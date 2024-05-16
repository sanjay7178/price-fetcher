package main

import (
	"context"
	"fmt"
	"time"
)

// PriceFetcher is an interface that can  fetch a price
type PriceFetcher interface {
	FetchPrice(context.Context, string) (float64, error)
}
// priceFetcher  implements the PriceFetcher inferace
type priceFecther struct{}



func (s *priceFecther) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	price, err := MockPriceFetcher(ctx, ticker)
	if err != nil {
		return price, err
	}
	return price , nil 

}

var priceMocks = map[string]float64{
	"BTC": 20_000,
	"ETH": 2000_0,
}

func MockPriceFetcher(ctx context.Context, ticker string) (float64, error) {
	// mimic the HTTP roundtrip
	time.Sleep(100*time.Millisecond)
	price, ok := priceMocks[ticker]
	if !ok {
		return price, fmt.Errorf("the given ticker (%s) is not supported", ticker)
	}
	return price, nil
}
