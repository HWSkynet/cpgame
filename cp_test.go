// cp_test.go
package cpgame_test

import (
	"testing"

	. "github.com/HWSkynet/cpgame"
)

var test Player = Player{
	Id:   "ayame",
	Life: 50,
}

func TestA(t *testing.T) {
	InputParse("  铅笔刀 投掷  ayame   ")
	t.Log("version:" + Version())
	t.Log(test.Life)
	test.GetDamage(15)
	t.Log(test.Life)
}
