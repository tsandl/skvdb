package iterator_test

import (
	"testing"

	"github.com/tsandl/skvdb/leveldb/testutil"
)

func TestIterator(t *testing.T) {
	testutil.RunSuite(t, "Iterator Suite")
}
