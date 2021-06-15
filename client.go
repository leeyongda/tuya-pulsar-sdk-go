package pulsarx

import (
	"context"
	"errors"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/apache/pulsar-client-go/pulsar/log"
	"github.com/sirupsen/logrus"

	"github.com/leeyongda/tuya-pulsar-sdk-go/auth"
)

var (
	client           pulsar.Client
	TuYaPulsarAddrCN = "pulsar+ssl://mqe.tuyacn.com:7285"
	TuYaPulsarAddrEU = "pulsar+ssl://mqe.tuyaeu.com:7285"
	TuYaPulsarAddrUS = "pulsar+ssl://mqe.tuyaus.com:7285"
)

type Client struct {
	client            pulsar.Client
	logger            log.Logger
	connectionTimeout time.Duration
	operationTimeout  time.Duration
}

func (c *Client) CreateProducer(options pulsar.ProducerOptions) (pulsar.Producer, error) {
	return c.client.CreateProducer(options)
}

func (c *Client) Subscribe(options pulsar.ConsumerOptions) (pulsar.Consumer, error) {
	csm, err := c.client.Subscribe(options)
	if err != nil {
		return nil, err
	}
	return csm, nil
}

func (c *Client) CreateReader(options pulsar.ReaderOptions) (pulsar.Reader, error) {
	return c.client.CreateReader(options)
}

func (c *Client) TopicPartitions(topic string) ([]string, error) {
	return c.client.TopicPartitions(topic)
}

func (c *Client) Close() {
	if c.client != nil {
		c.client.Close()
	}
}

type Consumer interface {
	ReceiveAndHandle(ctx context.Context, handler PayloadHandler)
	Close()
}

func (c *Client) NewConsumer(options pulsar.ConsumerOptions) (Consumer, error) {
	csm, err := c.Subscribe(options)
	if err != nil {
		return nil, err
	}
	return &consumerImpl{csm: csm, logger: c.logger}, nil
}

type Options func(*Client)

func WithConnectionTimeout(t time.Duration) Options {

	return func(c *Client) {
		if t > 0 {
			c.connectionTimeout = time.Second * t
		}
	}
}

func WithOperationTimeout(t time.Duration) Options {

	return func(c *Client) {
		if t > 0 {
			c.operationTimeout = time.Second * t
		}
	}
}

func NewClientWithTuYa(URL, accessID, accessKey string, opts ...Options) (client *Client, err error) {
	if accessID == "" && accessKey == "" {
		return nil, errors.New("accessID or accessKey is null")
	}
	authentication, err := auth.NewAuthenticationDataProviderWithTuYa(accessID, accessKey)
	if err != nil {
		return
	}
	var c = &Client{
		connectionTimeout: time.Second * 60,
		operationTimeout:  time.Second * 60,
	}
	for _, opt := range opts {
		opt(c)
	}
	logger := logrus.StandardLogger()
	logger.SetLevel(logrus.ErrorLevel)
	logs := log.NewLoggerWithLogrus(logger)
	cli, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:                        URL,
		ConnectionTimeout:          c.connectionTimeout,
		OperationTimeout:           c.operationTimeout,
		Authentication:             authentication,
		TLSAllowInsecureConnection: true,
		Logger:                     logs,
	})
	return &Client{client: cli, logger: logs}, err
}
