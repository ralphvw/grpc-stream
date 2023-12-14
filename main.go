package main

import (
	"context"
	"log"
	"net"

	currency "github.com/ralphvw/grpc-stream/protos"
	"google.golang.org/grpc"
)

type CurrencyServer struct {
	currency.UnimplementedCurrencyServer
}

func (c *CurrencyServer) GetRate(ctx context.Context, req *currency.RateRequest) (*currency.RateResponse, error) {
	//baseCurrency := req.GetBase()
	//destinationCurrency := req.GetDestination()

	rate := float32(1.25)

	// Construct the response
	response := &currency.RateResponse{
		Rate: rate,
	}

	return response, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8089")

	if err != nil {
		log.Fatalf("Cannot create listener... %s", err)
	}

	s := grpc.NewServer()

	service := &CurrencyServer{}

	currency.RegisterCurrencyServer(s, service)

	err = s.Serve(lis)

	if err != nil {
		log.Fatalf("Cannot serve %s", err)
	}
}
