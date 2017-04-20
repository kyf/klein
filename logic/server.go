package logic

import (
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/kyf/klein/config"
	//"github.com/kyf/klein/connector"
	"github.com/kyf/klein/session"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Server struct {
	conf          *config.LogicConfig
	svr           *grpc.Server
	logger        *Logger
	ln            net.Listener
	sessionClient session.SessionClient
	//connClients   map[string]connector.ConnectorClient
}

func NewServer(confpath string) (*Server, error) {
	conf := &config.LogicConfig{}
	err := conf.Load(confpath)
	if err != nil {
		return nil, err
	}

	logger, err := NewLogger(conf.LogPath, conf.LogPrefix)
	if err != nil {
		return nil, err
	}

	return &Server{conf: conf, logger: logger}, nil
}

func (this *Server) Run() error {
	svr := grpc.NewServer()
	serverSide := &LogicServerSide{this}
	RegisterLogicServer(svr, serverSide)

	ctx, _ := context.WithTimeout(context.Background(), time.Second*30)

	sessionConn, err := grpc.DialContext(ctx, this.conf.SessionHost, grpc.WithUserAgent(LOGIC_USERAGENT), grpc.WithInsecure())
	if err != nil {
		return err
	}
	sessionClient := session.NewSessionClient(sessionConn)

	this.svr = svr
	this.sessionClient = sessionClient

	ln, err := net.Listen("tcp", ":"+this.conf.RpcPort)
	if err != nil {
		return err
	}

	this.ln = ln
	exit := make(chan os.Signal, 1)

	go func() {
		err := this.svr.Serve(ln)
		if err != nil {
			this.logger.Fatal(err)
		}
	}()

	signal.Notify(exit, os.Kill, os.Interrupt)
	<-exit
	this.Stop()
	return nil
}

func (this *Server) Stop() {
	if this.ln != nil {
		this.ln.Close()
	}
	this.logger.Close()
}
