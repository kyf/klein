package connector

import (
	"google.golang.org/grpc"
)

type ConnectorServerSide struct{}

func (this *ConnectorServerSide) ReceiveMessage(mcontext context.Context, m *MessageRequest) (*ReceiveReply, error) {

}
