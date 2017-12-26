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
var playerList *PlayerList = &PL

func Test0(t *testing.T) {
	playerList = NewPlayerList()
	playerList.Add("123")
	playerList.ID("123").Name = "test#123"
}

func TestA(t *testing.T) {
	t.Log(PL.ID("123").Life)
	t.Log(PL.ID("123").InputParse("atk test#123 "))
	t.Log(PL.ID("123").Life)
	t.Log("version:" + Version())
	t.Log(Dice("11d2").Value())
}
