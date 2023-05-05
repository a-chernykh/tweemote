package mb

import (
	"encoding/json"

	"bitbucket.org/andreychernih/tweemote/errors"
	"github.com/streadway/amqp"
)

type MessageDelivery interface {
	Ack()
	Nack() error
	UnmarshalTo(out interface{})
	GetMessage() Message
	GetQueue() string
}

type Delivery struct {
	OriginalDelivery amqp.Delivery
	queue            string
}

func (d *Delivery) UnmarshalTo(out interface{}) {
	err := json.Unmarshal(d.OriginalDelivery.Body, out)
	errors.Check(err)
}

func (d *Delivery) Ack() {
	//glog.Info("Ack: %s", string(d.GetMessage().Serialize()))
	err := d.OriginalDelivery.Ack(false)
	errors.Check(err)
}

func (d *Delivery) Nack() error {
	//glog.Info("Nack: %s", string(d.GetMessage().Serialize()))
	return d.OriginalDelivery.Nack(false, false)
}

func (d *Delivery) GetMessage() Message {
	headers := d.OriginalDelivery.Headers
	if headers == nil {
		headers = make(map[string]interface{})
	}

	return &SerializedMessage{
		Data:             d.OriginalDelivery.Body,
		ContentTypeValue: d.OriginalDelivery.ContentType,
		headers:          headers,
	}
}

func (d *Delivery) GetQueue() string {
	return d.queue
}
