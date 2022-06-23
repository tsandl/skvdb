package memdb

import (
	"testing"

	"github.com/tsandl/skvdb/leveldb/testutil"
)

func TestMemDB(t *testing.T) {
	testutil.RunSuite(t, "MemDB Suite")
}
