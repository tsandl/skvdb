package table

import (
	"testing"

	"github.com/tsandl/skvdb/leveldb/testutil"
)

func TestTable(t *testing.T) {
	testutil.RunSuite(t, "Table Suite")
}
