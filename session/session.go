package session

import (
	"golang.org/x/net/context"
)

type SessionServerSide struct{}

func (this *SessionServerSide) Register(context.Context, *RegisterRequest) (*RegisterReply, error) {
	reply := &RegisterReply{Status: true, Message: "success"}

	return reply, nil
}

func (this *SessionServerSide) UnRegister(context.Context, *UnRegisterRequest) (*UnRegisterReply, error) {
	reply := &UnRegisterReply{Status: true, Message: "success"}

	return reply, nil
}

func (this *SessionServerSide) GetClients(context.Context, *GetClientsRequest) (*GetClientsReply, error) {
	connhosts := make([]*ConnHost, 1)
	connhosts = append(connhosts, &ConnHost{ConnId: "541245454", Host: "127.0.0.1"})
	reply := &GetClientsReply{Hosts: connhosts}
	return reply, nil
}
