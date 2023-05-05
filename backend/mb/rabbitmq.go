package mb

import (
	"fmt"
	"os"
	"sync"
	"time"

	"bitbucket.org/andreychernih/tweemote/errors"

	"github.com/golang/glog"
	"github.com/streadway/amqp"
)

const exchangeName = "tweemote"
const deadLetterExchangeName = "tweemote.dead"

type Rabbitmq struct {
	conn               *amqp.Connection
	defaultChannel     *amqp.Channel
	queueChannelsMutex *sync.Mutex
	queueChannels      map[string]*amqp.Channel
	prefetchCount      int
}

func NewConnection(prefetchCount int) (*Rabbitmq, error) {
	// TODO retry
	rabbit := &Rabbitmq{
		prefetchCount: prefetchCount,
	}
	err := rabbit.Connect()
	return rabbit, err
}

func (r *Rabbitmq) Connect() error {
	conn, err := amqp.Dial(os.Getenv("AMQP_URL"))
	if err != nil {
		return err
	}
	r.conn = conn

	defaultChannel, err := r.conn.Channel()
	if err != nil {
		return err
	}
	r.defaultChannel = defaultChannel

	r.queueChannelsMutex = &sync.Mutex{}
	r.queueChannels = make(map[string]*amqp.Channel)

	// Configure Exchange
	err = r.defaultChannel.ExchangeDeclare(exchangeName, "topic", true, false, false, false, nil)
	errors.Check(err)

	// Configure Dead-Letter Exchange
	err = r.defaultChannel.ExchangeDeclare(deadLetterExchangeName, "direct", true, false, false, false, nil)
	errors.Check(err)

	return nil
}

func (r *Rabbitmq) Disconnect() {
	if r.conn != nil {
		r.conn.Close()
	}
}

func (r *Rabbitmq) getQueueChannel(queue string) (*amqp.Channel, error) {
	r.queueChannelsMutex.Lock()
	defer r.queueChannelsMutex.Unlock()

	if _, ok := r.queueChannels[queue]; !ok {
		ch, err := r.conn.Channel()
		if err != nil {
			return nil, err
		}
		r.queueChannels[queue] = ch
	}

	return r.queueChannels[queue], nil
}

func (r *Rabbitmq) configureQueue(queue string) {
	// TODO extract this to a separate method and only call once
	// Add dead queue
	deadQueue := deadQueueName(queue)
	r.defaultChannel.QueueDeclare(deadQueue, true, false, false, false, nil)
	r.defaultChannel.QueueBind(deadQueue, queue, deadLetterExchangeName, false, nil)

	// Add queue
	args := make(amqp.Table)
	args["x-dead-letter-exchange"] = deadLetterExchangeName
	r.defaultChannel.QueueDeclare(queue, true, false, false, false, args)
	r.defaultChannel.QueueBind(queue, queue, exchangeName, false, nil)
	r.defaultChannel.Qos(r.prefetchCount, 0, false)
}

func (r *Rabbitmq) Publish(queue string, message Message) {
	r.configureQueue(queue)

	msg := amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		Timestamp:    time.Now(),
		ContentType:  message.ContentType(),
		Body:         message.Serialize(),
		Headers:      message.Headers(),
	}
	err := r.defaultChannel.Publish(exchangeName, queue, false, false, msg)
	if err != nil && err.(*amqp.Error).Code == amqp.ChannelError {
		// Channel was closed, make an attempt to re-establish connection
		glog.Errorf("Got error when publishing to RabbitMQ: %s. Re-connecting.", err)

		r.Disconnect()
		time.Sleep(10 * time.Second)
		r.Connect()
		r.Publish(queue, message)
	} else {
		errors.Check(err)
	}
}

func (r *Rabbitmq) Consume(queue string, prefetchCount int) (<-chan MessageDelivery, error) {
	r.configureQueue(queue)

	queueChannel, err := r.getQueueChannel(queue)
	if err != nil {
		return nil, err
	}
	queueChannel.Qos(prefetchCount, 0, false)

	c := make(chan MessageDelivery)
	messages, err := queueChannel.Consume(queue, "consumer", false, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	go func() {
		for message := range messages {
			delivery := Delivery{
				queue:            queue,
				OriginalDelivery: message,
			}
			c <- &delivery
		}
	}()

	return c, nil
}

func (r *Rabbitmq) QueueInspect(queue string) (*QueueInfo, error) {
	r.configureQueue(queue)

	q, err := r.defaultChannel.QueueInspect(queue)
	if err != nil {
		return nil, err
	}

	return &QueueInfo{
		Name:      q.Name,
		Messages:  q.Messages,
		Consumers: q.Consumers,
	}, nil
}

func (r *Rabbitmq) Nack(delivery MessageDelivery, reason string) {
	deadQueue := deadQueueName(delivery.GetQueue())

	msg := delivery.GetMessage()
	msg.SetHeader("x-reason", reason)

	mm := amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		Timestamp:    time.Now(),
		ContentType:  msg.ContentType(),
		Body:         msg.Serialize(),
		Headers:      msg.Headers(),
	}
	err := r.defaultChannel.Publish(deadLetterExchangeName, deadQueue, false, false, mm)
	errors.Check(err)

	delivery.Ack()
}

func deadQueueName(queue string) string {
	return fmt.Sprintf("%s-dead", queue)
}
