package leveldb

import (
	"testing"

	"github.com/tsandl/skvdb/leveldb/testutil"
)

func TestLevelDB(t *testing.T) {
	testutil.RunSuite(t, "LevelDB Suite")
}
