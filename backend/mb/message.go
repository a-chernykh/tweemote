package mb

import (
	"encoding/json"

	"bitbucket.org/andreychernih/tweemote/errors"
)

type Message interface {
	Serialize() []byte
	ContentType() string
	Headers() map[string]interface{}
	SetHeader(key string, value interface{})
}

type JsonMessage struct {
	Data    interface{}
	headers map[string]interface{}
}

func NewJsonMessage(data interface{}) *JsonMessage {
	return &JsonMessage{
		Data:    data,
		headers: make(map[string]interface{}),
	}
}

func (jm *JsonMessage) Serialize() []byte {
	d, err := json.Marshal(jm.Data)
	errors.Check(err)
	return d
}

func (jm *JsonMessage) ContentType() string {
	return "application/json"
}

func (jm *JsonMessage) Headers() map[string]interface{} {
	return jm.headers
}

func (jm *JsonMessage) SetHeader(key string, value interface{}) {
	jm.headers[key] = value
}

type SerializedMessage struct {
	Data             []byte
	ContentTypeValue string
	headers          map[string]interface{}
}

func (jm *SerializedMessage) Serialize() []byte {
	return jm.Data
}

func (jm *SerializedMessage) ContentType() string {
	return jm.ContentTypeValue
}

func (jm *SerializedMessage) Headers() map[string]interface{} {
	return jm.headers
}

func (jm *SerializedMessage) SetHeader(key string, value interface{}) {
	jm.headers[key] = value
}
