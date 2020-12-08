package server

import (
	"context"

	protos "github.com/gmgale/go_micro/13_Intro_to_gRPC_and_Protocol_Buffers/protos/currency"
	"github.com/hashicorp/go-hclog"
)

type Currency struct {
	log hclog.Logger
}

// type CurrencyServer interface {
// 	GetRate(context.Context, *RateRequest) (*RateResponse, error)
// 	mustEmbedUnimplementedCurrencyServer()
// }

func NewCurrency(l hclog.Logger) *Currency {
	return &Currency{l}
}

func (c *Currency) GetRate(ctx context.Context, rr *protos.RateRequest) (*protos.RateResponse, error) {
	c.log.Info("Handle request for GetRate", "base", rr.GetBase(), "dest", rr.GetDestination())
	return &protos.RateResponse{Rate: 0.5}, nil
}
