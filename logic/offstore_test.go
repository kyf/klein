package logic

import (
	"log"
	"testing"

	"github.com/kyf/klein/config"
	"github.com/kyf/klein/message"
)

const (
	testkey = "offline_msg_keyongfeng"
)

func initStore() *OfflineStore {
	svr := &Server{
		conf: &config.LogicConfig{
			RedisHost: "127.0.0.1:6379",
			RedisAuth: "6Renyou2016",
		},
	}

	return NewOfflineStore(svr)
}

func initMsgReq() *MessageRequest {
	m := &MessageRequest{
		MsgType:    int32(message.TextMessage),
		SequenceId: "sequenceid",
		ConnId:     "connid",
		Sender:     "sender",
		Receiver:   "receiver",
		Body:       "body is  .... ",
	}

	return m
}

func TestPut(t *testing.T) {
	store := initStore()
	m := initMsgReq()
	err := store.Put(testkey, m)
	if err != nil {
		t.Error(err)
	}
}

func TestGet(t *testing.T) {
	store := initStore()
	list, err := store.Get(testkey)
	if err != nil {
		t.Error(err)
	} else {
		for _, it := range list {
			log.Printf("%+v", it)
		}
	}
}

func TestRemove(t *testing.T) {
	store := initStore()
	err := store.Remove(testkey)
	if err != nil {
		t.Error(err)
	}
}
