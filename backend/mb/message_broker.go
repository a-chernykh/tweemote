package mb

type MessageBroker interface {
	Connect() error
	Disconnect()
	Publish(queue string, message Message)
	Consume(queue string, prefetchCount int) (<-chan MessageDelivery, error)
	QueueInspect(queue string) (*QueueInfo, error)
	Nack(delivery MessageDelivery, reason string)
}
