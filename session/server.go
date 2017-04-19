package session

import (
	"net"
	"os"
	"os/signal"

	"github.com/kyf/klein/config"
	"google.golang.org/grpc"
)

type Server struct {
	ln     net.Listener
	svr    *grpc.Server
	conf   *config.SessionConfig
	logger *Logger
}

func NewServer(confpath string) (*Server, error) {
	conf := &config.SessionConfig{}
	err := conf.Load(confpath)
	if err != nil {
		return nil, err
	}

	logger, err := NewLogger(conf.LogPath, conf.LogPrefix)

	ln, err := net.Listen("tcp", conf.RpcPort)
	if err != nil {
		return nil, err
	}

	svr := grpc.NewServer()
	sessServer := &SessionServerSide{}
	RegisterSessionServer(svr, sessServer)
	return &Server{svr: svr, conf: conf, ln: ln, logger: logger}, nil
}

func (this *Server) Run() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Kill, os.Interrupt)

	go func() {
		for _ = range ch {
			this.Stop()
		}
	}()

	err := this.svr.Serve(this.ln)
	if err != nil {
		this.logger.Fatal(err)
	}
}

func (this *Server) Stop() {
	this.ln.Close()
	this.logger.Close()
}
