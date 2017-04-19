package connector

import (
	"net/http"

	"github.com/gorilla/mux"
)

type HttpServer struct {
	ctx *Connector
	svr *http.Server
}

func syncMessage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("sync message ..."))
}

func NewHttpServer(port string, ctx *Connector) *HttpServer {
	handler := mux.NewRouter()
	handler.HandleFunc("/sync/msg", syncMessage)
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
