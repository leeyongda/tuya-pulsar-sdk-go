package pulsarx

import (
	"context"

	"github.com/apache/pulsar-client-go/pulsar"
)

type PayloadHandler interface {
	HandlePayload(ctx context.Context, message pulsar.Message) error
}

type handler struct{}

func newHandler() PayloadHandler {

	return &handler{}
}

func (m *handler) HandlePayload(ctx context.Context, message pulsar.Message) error {
	return nil
}
