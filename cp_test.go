// cp_test.go
package cpgame_test

import (
	"testing"

	. "github.com/HWSkynet/cpgame"
)

var test Player = Player{
	Id: "ayame",
}
var playerList *PlayerList = &PL

func Test0(t *testing.T) {
	t.Log("newgame")
	playerList, _ = NewGame()
	t.Log("123")
	playerList.Add("123")
	playerList.ID("123").Name = "test#123"
}

func TestA(t *testing.T) {
	t.Log(PL.ID("123").InputParse("atk test#123 "))
	t.Log("version:" + Version())
	t.Log(Dice("11d2").Value())
}
