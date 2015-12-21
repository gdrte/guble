package store

import (
	"testing"
)

func TestMemoryPutGetDelete(t *testing.T) {
	CommonTestPutGetDelete(t, NewMemoryKVStore())
}

func TestMemoryIterateKeys(t *testing.T) {
	CommonTestIterateKeys(t, NewMemoryKVStore())
}

func BenchmarkMemoryPutGet(b *testing.B) {
	CommonBenchPutGet(b, NewMemoryKVStore())
}
