package logic

import (
	//"net"

	//"github.com/kyf/klein/session"
	"golang.org/x/net/context"
	"time"
)

type LogicServerSide struct {
	ctx *Server
}

func (this *LogicServerSide) ReceiveMessage(ctx context.Context, req *MessageRequest) (*ReceiveReply, error) {
	mcontext, _ := context.WithTimeout(context.Background(), time.Second*5)
	clients, err := this.ctx.sessionClient.GetClients(mcontext, nil)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
