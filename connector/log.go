package connector

import (
	"log"
	"os"
)

type Logger struct {
	*log.Logger
	io.ReadWriteCloser
}

func NewLogger(path, prefix string) (*Logger, error) {
	fp, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.A_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	logger := log.New(fp, prefix, log.LstdFlags)
	return &Logger{logger, fp}
}
