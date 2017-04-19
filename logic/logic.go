package logic

import (
	"net"

	"github.com/kyf/klein/session"
	"golang.org/x/net/context"
)

type LogicServerSide struct {
	ctx *Server
}

func (this *LogicServerSide) SendMessage(ctx context.Context, req *MessageRequest) (*SendReply, error) {

}

func (this *LogicServerSide) ReceiveMessage(ctx context.Context, req *MessageRequest) (*ReceiveReply, error) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	rp, err := this.ctx.sessionClient.GetClients(ctx, r)
	if err != nil {
		return nil, err
	}

}
