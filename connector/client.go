package connector

import (
	"encoding/binary"
	"errors"
	"net"
	"time"
)

var (
	ErrCodehandshakeInvalidSize = 101
	ErrCodehandshakeInvalidBody = 102

	SuccessCodehandshake = 200
)

const (
	SessionRegisterTimeout   = time.Second * 10
	SessionUnRegisterTimeout = time.Second * 10
)

type Client struct {
	conn       net.Conn
	ConnId     ID
	UserId     [32]byte
	DeviceInfo []byte
	Ch         chan *MessageRequest
}

func NewClient(c net.Conn) *Client {
	return &Client{conn: c, ConnId: GeneralId()}
}

func (this *Client) Read(p []byte) (int, error) {
	return this.conn.Read(p)
}

func (this *Client) Write(p []byte) (int, error) {
	return this.conn.Write(p)
}

func (this *Client) ReadFull(p []byte) error {
	num, err := this.conn.Read(p)
	if err != nil {
		return err
	}

	if num != len(p) {
		return errors.New("client.ReadFull: not full")
	}

	return nil
}

func (this *Client) ReadMessageHeader(p []byte) error {
	err := this.ReadFull(p)
	if err != nil {
		return errors.New("message header is invalid")
	}
	return nil
}

func (this *Client) WriteError(err error) {
	this.Write([]byte(err.Error()))
}

func (this *Client) Id() ID {
	return this.ConnId
}

func (this *Client) handshake() error {
	length := make([]byte, 4)
	err := this.ReadFull(length)
	if err != nil {
		this.Write([]byte{byte(ErrCodehandshakeInvalidSize)})
		return err
	}

	_len := binary.BigEndian.Uint64(length)
	buf := make([]byte, _len)
	err = this.ReadFull(buf)
	if err != nil {
		this.Write([]byte{byte(ErrCodehandshakeInvalidBody)})
		return err
	}

	userid, deviceInfo := buf[:32], buf[32:]
	copy(this.UserId[:], userid)
	this.DeviceInfo = deviceInfo

	this.Write([]byte{byte(SuccessCodehandshake)})
	return nil
}

func (this *Client) Wait() {
	for m := range this.Ch {
		msg := &message.Message{
			Type:       m.MsgType,
			SequenceId: m.SequenceId,
			ConnId:     m.ConnId,
			Sender:     m.Sender,
			Receiver:   m.Receiver,
			Body:       m.Body,
		}

		data, err := msg.Encode()
		_, err := this.Write(data)
		if err != nil {

		}
	}
}
