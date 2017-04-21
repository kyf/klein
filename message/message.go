package message

import (
	"errors"
)

type MessageType byte

func (this MessageType) String() string {
	switch this {
	case TextMessage:
		return "TextMessage"
	case ImageMessage:
		return "ImageMessage"
	case AckMessage:
		return "AckMessage"
	case PingMessage:
		return "PingMessage"
	case PongMessage:
		return "PongMessage"
	default:
	}

	return "unknwon message type"
}

var (
	ErrInvalidMessage = errors.New("message is invalid")
)

const (
	_ MessageType = iota
	TextMessage
	ImageMessage
	AckMessage
	PingMessage
	PongMessage
)

type Message struct {
	Type       MessageType
	SequenceId [16]byte
	ConnId     [20]byte
	Sender     [16]byte
	Receiver   [16]byte
	Body       []byte
}

func (this *Message) Decode(body []byte) error {
	skip := 0
	this.Type = MessageType(body[skip])
	skip++

	length := len(this.SequenceId)
	num := copy(this.SequenceId[:], body[skip:length])
	if num < length {
		return ErrInvalidMessage
	}
	skip += length

	length = len(this.ConnId)
	num = copy(this.ConnId[:], body[skip:length])
	if num < length {
		return ErrInvalidMessage
	}
	skip += length

	length = len(this.Sender)
	num = copy(this.Sender[:], body[skip:length])
	if num < length {
		return ErrInvalidMessage
	}
	skip += length

	length = len(this.Receiver)
	num = copy(this.Receiver[:], body[skip:length])
	if num < length {
		return ErrInvalidMessage
	}
	skip += length

	copy(this.Body[:], body[skip:])
	return nil
}

func (this *Message) Encode() ([]byte, error) {
	return nil, nil
}
