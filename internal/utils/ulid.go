package utils

import (
	cryptoRand "crypto/rand"
	"io"
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

// NewULIDString returns a new ULID string with entropy from crypto/rand and a monotonic source.
func NewULIDString() string {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	return ulid.MustNew(ulid.Timestamp(t), entropy).String()
}

// NewULID returns raw ULID type
func NewULID() ulid.ULID {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	return ulid.MustNew(ulid.Timestamp(t), entropy)
}

// NewCryptoULIDString pure crypto-random ULID (not monotonic, but more secure)
func NewCryptoULIDString() string {
	t := time.Now()
	return ulid.MustNew(ulid.Timestamp(t), cryptoEntropy()).String()
}

func cryptoEntropy() io.Reader {
	return cryptoRand.Reader
}
