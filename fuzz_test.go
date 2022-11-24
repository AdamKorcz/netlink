package netlink

import (
	"testing"
)

func FuzzDeserializeRoute(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		_, _ = deserializeRoute(data)
	})
}
