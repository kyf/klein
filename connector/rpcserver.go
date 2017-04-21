package connector

import (
	"google.golang.org/grpc"
)

type ConnectorServerSide struct {
	ctx *Connector
}

func (this *ConnectorServerSide) ReceiveMessage(mcontext context.Context, m *MessageRequest) (*ReceiveReply, error) {
	times := 0
	this.ctx.Find(m.ConnId)

BEGIN:
	ch := make(chan bool, 1)
	pkg := &MessagePackage{ch, m}

	select {
	case <-ch:
		close(ch)
	case <-time.After(time.Second * 10):
		close(ch)
		times++
		if times < 3 {
			goto BEGIN
		} else {
			return errors.New("device not ack")
		}
	}

	return nil, nil
}
