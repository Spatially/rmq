package rmq

type Consumer interface {
	Consume(delivery Delivery)
}

type ClosableConsumer interface {
	Consumer
	Close() error
	Closer() chan bool
	ConsumeOneTime() bool
}
