package main

import (
	"context"
	"fmt"
)

type metricService struct {
	next PriceFetcher
}

func NewMetricServce(next  PriceFetcher) PriceFetcher{
	return &metricService{
		next: next,
	}
}
func (s *metricService) FetchPrice(ctx context.Context, ticker string)(price float64 , err error){
	fmt.Println("pushing metrics to prometheus ")
	return s.next.FetchPrice(ctx, ticker)
}