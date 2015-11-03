package rmq

type Consumer interface {
	Consume(delivery Delivery)
}

type ClosableConsumer interface {
	Closer() chan bool
	Close() error
	ConsumeOneTime() bool
}
