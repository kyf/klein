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
	ln *net.TCPListener
}

func (this *TcpKeepAliveListener) Accept() (net.Conn, error) {
	conn, err := this.ln.AcceptTCP()
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

	return &TcpServer{&TcpKeepAliveListener{ln.(*net.TCPListener)}, ctx}, nil
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

	header := make([]byte, 4)
	var body []byte
	for {
		err := cli.ReadMessageHeader(header)
		if err != nil {
			cli.WriteError(err)
			break
		}
		body = make([]byte, binary.BigEndian.Uint64(header))
		err = cli.ReadFull(body)
		if err != nil {
			cli.WriteError(err)
			break
		}
		msg := &message.Message{}
		err = msg.Decode(body)
		if err != nil {
			cli.WriteError(err)
			break
		}
		switch msg.Type {
		case message.TextImage:
		default:
		}
	}
}
