package session

import (
	"golang.org/x/net/context"
)

type SessionServerSide struct{}

func (this *SessionServerSide) Register(context.Context, *RegisterRequest) (*RegisterReply, error) {

}

func (this *SessionServerSide) UnRegister(context.Context, *UnRegisterRequest) (*UnRegisterReply, error) {

}

func (this *SessionServerSide) GetClients(context.Context, *GetClientsRequest) (*GetClientsReply, error) {

}
