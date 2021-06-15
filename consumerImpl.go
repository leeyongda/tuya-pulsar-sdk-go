package pulsarx

import (
	"context"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/apache/pulsar-client-go/pulsar/log"
)

type consumerImpl struct {
	csm    pulsar.Consumer
	logger log.Logger
}

func (c *consumerImpl) ReceiveAndHandle(ctx context.Context, handler PayloadHandler) {
	for {
		msg, err := c.csm.Receive(ctx)
		if err != nil {
			c.logger.Errorf("ReceiveAndHandle Error: %s", err.Error())
			return
		}
		if err := handler.HandlePayload(context.Background(), msg); err != nil {
			c.logger.Errorf("ReceiveAndHandle HandlePayload Error: %s", err.Error())
		} else {
			c.csm.Ack(msg)
		}
	}

}

func (c *consumerImpl) Close() {
	if c.csm != nil {
		c.csm.Close()
	}
}
