package connector

import (
	"crypto/sha1"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/kyf/klein/session"
)

type ID [20]byte

func GeneralId() ID {
	seed := fmt.Sprintf("%d_%d", time.Now().UnixNano(), rand.Intn(1000))
	id := sha1.Sum([]byte(seed))
	return ID(id)
}

func NewRegReq() interface{} {
	return &session.RegisterRequest{
		UserId:        "",
		ConnId:        "",
		ConnectorHost: "",
	}
}

func NewUnRegReq() interface{} {
	return &session.UnRegisterRequest{
		UserId:        "",
		ConnId:        "",
		ConnectorHost: "",
	}
}

var (
	regReqPool   *sync.Pool = &sync.Pool{New: NewRegReq}
	unRegReqPool *sync.Pool = &sync.Pool{New: NewUnRegReq}
)

func getRegReqFromPool() *session.RegisterRequest {
	return regReqPool.Get().(*session.RegisterRequest)
}

func releaseRegReq(r *session.RegisterRequest) {
	r.Reset()
	regReqPool.Put(r)
}

func getUnRegReqFromPool() *session.UnRegisterRequest {
	return unRegReqPool.Get().(*session.UnRegisterRequest)
}

func releaseUnRegReq(r *session.UnRegisterRequest) {
	r.Reset()
	unRegReqPool.Put(r)
}
