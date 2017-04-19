package session

import (
	"log"
	"net"

	"github.com/kyf/klein/config"
	"google.golang.org/grpc"
)

type Server struct {
	ln     net.Listener
	svr    *grpc.Server
	conf   *config.SessionConfig
	logger *log.Logger
}

func NewServer(confpath string) (*Server, error) {
	conf := &config.SessionConfig{}
	err := conf.Load(confpath)
	if err != nil {
		return err
	}

	logger, err := NewLogger()

	ln, err := net.Listen("tcp", this.conf.RpcPort)
	if err != nil {
		return err
	}

	svr := grpc.NewServer()
	sessServer := &SessionServerSide{}
	RegisterSessionServer(svr, sessServer)
	return &Server{svr: svr, conf: conf, ln: ln}
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
