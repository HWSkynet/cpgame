// cp_test.go
package cpgame_test

import (
	"testing"

	. "github.com/HWSkynet/cpgame"
)

var playerList *PlayerList = &PL

var zjd Knife = Knife{Item{Name: "指甲刀", Weight: 1}}

func Test0(t *testing.T) {
	t.Log("newgame")
	playerList, _ = NewGame()
	playerList.Add("123")
	playerList.ID("123").Name = "test#123"
	playerList.Add("007")
	playerList.ID("007").Name = "target#007"
}

func TestA(t *testing.T) {
	PL.ID("123").Pocket = append(PL.ID("123").Pocket, &zjd)
	t.Log(PL.ID("123").InputParse("equip 指甲刀"))
	t.Log(PL.ID("123"))
	t.Log(PL.ID("007"))
	t.Log(PL.ID("123").InputParse("throw target#007"))
	t.Log(PL.ID("123"))
	t.Log(PL.ID("007"))
	t.Log(PL.ID("123").InputParse("atk test#123 "))
	t.Log("version:" + Version())
	t.Log(Dice("11d2").Value())
}
