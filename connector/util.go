package connector

import (
	"crypto/sha1"
	"math/rand"
	"time"
)

type ID string

func GeneralId() ID {
	seed := fmt.Sprintf("%d_%d", time.Now().UnixNano(), rand.Intn(1000))
	id := sha1.Sum([]byte(seed))
	return ID(id)
}
