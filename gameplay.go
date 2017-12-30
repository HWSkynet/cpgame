// game.go
package cpgame

//"fmt"

type PlayerList []*Player

var PL PlayerList

func NewPlayerList() *PlayerList {
	PL = make(PlayerList, 0, 100)
	return &PL
}

func (self *PlayerList) ID(id string) *Player {
	for _, v := range *self {
		if v.Id == id {
			return v
		}
	}
	return &Player{}
}

func (self *PlayerList) NAME(name string) *Player {
	for _, v := range *self {
		if v.Name == name {
			return v
		}
	}
	return &Player{}
}

func (self *PlayerList) IsExist(id string) bool {
	if self.ID(id).Id != "" {
		return true
	}
	return false
}
func (self *PlayerList) IsReady(id string) bool {
	p := self.ID(id)
	if p.Id != "" {
		if p.State == "ready" {
			return true
		}
	}
	return false
}

func (self *PlayerList) Add(id string) {
	*self = append(*self, NewPlayer(id))
}

func (self *PlayerList) Remove(id string) {
	for i := range *self {
		if (*self)[i].Id == id {
			(*self) = append((*self)[:i], (*self)[i+1:]...)
			return
		}
	}
}

func (self *PlayerList) ImReady(id string) {
	p := self.ID(id)
	if p.Id != "" {
		if p.State == "ready" {
			p.State = "!ready"
		} else {
			p.State = "ready"
		}
	}
}

func (self *PlayerList) CountReady() int {
	var cnt int = 0
	for _, v := range *self {
		if v.State == "ready" {
			cnt += 1
		}
	}
	return cnt
}

func (self *PlayerList) CountLive() int {
	var cnt int = 0
	for _, v := range *self {
		if v.Alive {
			cnt += 1
		}
	}
	return cnt
}
