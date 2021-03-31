package gcp

import (
	"context"
	"fmt"
	raw "google.golang.org/api/compute/v1"
	"google.golang.org/api/option"
)


type Client struct {
	projectID string
	svc *raw.Service
}

func NewClient(ctx context.Context, opts ...option.ClientOption)(*Client, error) {
	svc, err :=raw.NewService(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("constructing compute Client: %v", err)
	}
	c := &Client{
		svc: svc,
	}
	return c, nil
}
