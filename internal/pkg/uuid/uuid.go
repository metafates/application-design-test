package uuid

import (
	"crypto/rand"
	"fmt"
)

// NewPseudoUUID generates a new pseudo uuid.
// Note - NOT RFC4122 compliant
func NewPseudoUUID() string {
	b := make([]byte, 16)

	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}
