package message

import (
	"log"
	"testing"
)

func TestMessageType(t *testing.T) {
	log.Print(TextMessage)
	log.Print(ImageMessage)
}
