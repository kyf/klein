package connector

import (
	"net/http"
)

type HttpServer struct {
	ctx *Connector
	svr *http.Server
}

func NewHttpServer(port string, ctx *Connector) *HttpServer {
	svr := &http.Server{Addr: ":" + port, Handler: handler}
	return &HttpServer{ctx, svr}
}

func (this *HttpServer) run() {
	err := this.svr.ListenAndServe()
	if err != nil {
		this.ctx.logger.Fatal(err)
	}
}

func (this *HttpServer) stop() {
	this.svr.Close()
}
