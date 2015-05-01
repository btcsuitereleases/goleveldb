package iterator_test

import (
	"testing"

	"github.com/btcsuitereleases/goleveldb/leveldb/testutil"
)

func TestIterator(t *testing.T) {
	testutil.RunSuite(t, "Iterator Suite")
}
