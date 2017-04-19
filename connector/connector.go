package connector

import (
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/kyf/klein/config"
	"github.com/kyf/klein/session"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Connector struct {
	clientPool    map[ID]*Client
	conf          *config.ConnectorConfig
	logger        *Logger
	sessionClient session.SessionClient
	sync.Mutex
}

func NewConnector() *Connector {
	return &Connector{clientPool: make(map[ID]*Client)}
}

func (this *Connector) Add(cli *Client) error {
	this.Lock()
	defer this.Unlock()

	req := getRegReqFromPool()
	defer releaseRegReq(req)
	req.ConnectorHost = this.conf.GrpcHost
	req.UserId = string(cli.UserId[:])
	req.ConnId = string(cli.ConnId[:])

	ctx, _ := context.WithTimeout(context.Background(), SessionRegisterTimeout)
	_, err := this.sessionClient.Register(ctx, req)
	if err != nil {
		return err
	}

	this.clientPool[cli.Id()] = cli

	return nil
}

func (this *Connector) Remove(cli *Client) error {
	this.Lock()
	defer this.Unlock()

	req := getUnRegReqFromPool()
	defer releaseUnRegReq(req)
	req.ConnectorHost = this.conf.GrpcHost
	req.UserId = string(cli.UserId[:])
	req.ConnId = string(cli.ConnId[:])

	ctx, _ := context.WithTimeout(context.Background(), SessionUnRegisterTimeout)
	_, err := this.sessionClient.UnRegister(ctx, req)
	if err != nil {
		return err
	}

	delete(this.clientPool, cli.Id())
	return nil
}

func (this *Connector) Init(confPath string) error {
	conf := &config.ConnectorConfig{}
	err := conf.Load(confPath)
	if err != nil {
		return err
	}

	logger, err := NewLogger(conf.LogPath, conf.LogPrefix)
	if err != nil {
		return err
	}

	this.conf = conf
	this.logger = logger

	ctx, _ := context.WithTimeout(context.Background(), time.Second*30)
	conn, err := grpc.DialContext(ctx, conf.SessionHost, grpc.WithUserAgent(CONNECTOR_USERAGENT), grpc.WithInsecure())
	if err != nil {
		return err
	}
	secli := session.NewSessionClient(conn)
	this.sessionClient = secli
	return nil
}

func (this *Connector) Run() {
	exit := make(chan os.Signal, 1)

	tcpserver, err := NewTcpServer(this.conf.TcpPort, this)
	if err != nil {
		this.logger.Fatal(err)
	}

	httpserver := NewHttpServer(this.conf.HttpPort, this)

	go tcpserver.run()

	go httpserver.run()

	signal.Notify(exit, os.Interrupt, os.Kill)
	<-exit

	tcpserver.stop()
	httpserver.stop()
	this.Stop()
}

func (this *Connector) Stop() {
	this.logger.Close()
}
