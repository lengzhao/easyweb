package util

import (
	"fmt"
	"sync/atomic"
)

var lastId int64
var idPrefix string = "id"

func SetIDPrefix(prefix string) {
	idPrefix = prefix
}

// GetID returns a unique identifier.
//
// It does this by incrementing the value of lastId by 1 using the atomic
// package's AddInt64 function, assigning the result to val. It then formats
// the val value as a string with leading zeros using the fmt package's
// Sprintf function and returns the result.
//
// Returns:
//
//	string: The unique identifier.
func GetID() string {
	val := atomic.AddInt64(&lastId, 1)
	return fmt.Sprintf("%s_%04d", idPrefix, val)
}
