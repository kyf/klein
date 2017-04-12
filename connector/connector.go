package connector

import (
	"os"
	"os/signal"
	"sync"

	"github.com/kyf/klein/config"
)

type Connector struct {
	clientPool map[ID]*Client
	conf       *config.ConnectorConfig
	logger     *Logger
	sync.Mutex
}

func NewConnector() *Connector {
	return &Connector{clientPool: make(map[ID]*Client)}
}

func (this *Connector) Add(cli *Client) error {
	this.Lock()
	defer this.Unlock()

	this.clientPool[cli.Id()] = cli
}

func (this *Connector) Remove(cli *Client) error {
	this.Lock()
	defer this.Unlock()

	delete(this.clientPool, cli.Id())
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
	return nil
}

func (this *Connector) Run() {
	exit := make(chan os.Signal, 1)
	go handleSignal(exit)
	signal.Notify(exit, os.Interrupt, os.Kill)

	tcpserver, err := NewTcpServer(this.conf.TcpPort)
	if err != nil {
		this.logger.Fatal(err)
	}

	httpserver, err := NewHttpServer(this.conf.HttpPort)
	if err != nil {
		this.logger.Fatal(err)
	}

	go tcpserver.run()

	go httpserver.run()

	<-exit

	tcpserver.stop()
	httpserver.stop()
	this.Stop()
}

func (this *Connector) Stop() {
	this.logger.Close()
}
