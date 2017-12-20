// cp_test.go
package cpgame_test

import (
	"testing"

	. "github.com/HWSkynet/cpgame"
)

func TestVersion(t *testing.T) {
	t.Log("version:" + Version())
}
