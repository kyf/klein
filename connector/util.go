package connector

import (
	"crypto/sha1"
	"fmt"
	"math/rand"
	"time"
)

type ID [20]byte

func GeneralId() ID {
	seed := fmt.Sprintf("%d_%d", time.Now().UnixNano(), rand.Intn(1000))
	id := sha1.Sum([]byte(seed))
	return ID(id)
}
