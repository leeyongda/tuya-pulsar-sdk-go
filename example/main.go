package main

import (
	"context"
	"fmt"
	"os"

	"github.com/apache/pulsar-client-go/pulsar"
	_ "github.com/joho/godotenv/autoload"

	pulsarx "github.com/leeyongda/tuya-pulsar-sdk-go"
)

func main() {

	ak := os.Getenv("AK")
	sk := os.Getenv("SK")

	client, err := pulsarx.NewClientWithTuYa(pulsarx.TuYaPulsarAddrCN, ak, sk)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	options := pulsar.ConsumerOptions{
		Topic:            pulsarx.GetTopicTestForAccessID(ak), // test env
		SubscriptionName: pulsarx.GetSubscriptionName(pulsarx.GetTopicForAccessID(ak)),
		Type:             2,
		RetryEnable:      false,
	}

	consumer, err := client.NewConsumer(options)
	if err != nil {
		panic(err)

	}

	defer consumer.Close()

	consumer.ReceiveAndHandle(context.TODO(), testHandler{})
}

type testHandler struct {
}

func (testHandler) HandlePayload(ctx context.Context, message pulsar.Message) error {
	// todo process HandlePayload
	fmt.Println("msg", string(message.Payload()))
	return nil
}
