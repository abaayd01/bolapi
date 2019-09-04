package bolpy_client

import (
	"bolapi/internal/pkg/bolproto"
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	defaultAddress = "localhost:50051"
)

type BolpyClient struct {
	Context *BolpyContext
}

type BolpyContext struct {
	PriceEvaluatorClient bolproto.PriceEvaluatorClient
}

func NewBolpyContext(address string) (*BolpyContext, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return nil, err
	}

	c := bolproto.NewPriceEvaluatorClient(conn)

	return &BolpyContext{
		PriceEvaluatorClient: c,
	}, nil
}

func (bC *BolpyClient) EvaluatePrice(req *bolproto.PriceEvaluationRequest) (*bolproto.PriceEvaluationResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := bC.Context.PriceEvaluatorClient.EvaluatePrice(ctx, &bolproto.PriceEvaluationRequest{
		CurrentPrice:     nil,
		HistoricalPrices: nil,
	})

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return r, nil
}
