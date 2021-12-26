package server

import (
	"context"

	"github.com/hashicorp/go-hclog"
)

type Currency struct {
	log hclog.Logger
}

func (c *Currency) GetRate(ctx context.Context, rr *protos.RateRequest) (*protos.RateResponse, error) {

	c.log.Info("Handle GetRate", "base", rr.GetBase(), "destination", rr.GetDestination)
}
