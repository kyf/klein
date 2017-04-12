package connector

import (
	"net"
	"time"
)

type TcpServer struct {
	ln  *TcpKeepAliveListener
	ctx *Connector
}

type TcpKeepAliveListener struct {
	ln *net.TcpListener
}

func (this *TcpKeepAliveListener) Accept() (net.Conn, error) {
	conn, err := this.ln.AcceptTcp()
	if err != nil {
		return nil, err
	}

	conn.SetKeepAlive(true)
	conn.SetKeepAlivePeriod(time.Minute * 1)
	return conn, nil
}

func (this *TcpKeepAliveListener) Stop() {
	this.ln.Close()
}

func NewTcpServer(port string, ctx *Connector) (*TcpServer, error) {
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return nil, err
	}

	return &TcpServer{&TcpKeepAliveListener{ln}, ctx}
}

func (this *TcpServer) run() {
	for {
		conn, err := this.ln.Accept()
		if err != nil {
			if ne, ok := err.(net.Error); ok && ne.Temporary() {
				time.Sleep(time.Millisecond * 100)
				continue
			}
			this.ctx.logger.Fatal(err)
		}

		handleTcpConn(conn, this)
	}
}

func (this *TcpServer) stop() {
	this.ln.Stop()
}

func handleTcpConn(conn net.Conn, server *TcpServer) {
	defer conn.Close()
	cli := NewClient(conn)
	err := cli.handshake()
	if err != nil {
		return
	}

	server.ctx.Add(cli)
	defer server.ctx.Remove(cli)

	for {
		cli.Read()
	}
}
