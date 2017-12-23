// player.go
package cpgame

import "fmt"

type Player struct {
	Id       string
	Name     string
	State    string
	Life     int8   // 生命值0-50 显示为0.0-5.0
	Energy   int8   // 行动点数
	Weight   int8   // 当前负重(最大100)
	Killed   int8   // 杀敌数
	Assist   int8   // 助攻数
	Pos      Coord  // 坐标
	Mainhand Item   // 物品
	Pocket   []Item // 包裹
}

func NewPlayer(id string) *Player {
	return &Player{Id: id, Life: 50, Energy: 5, State: "init"}
}

func (self *Player) GetDamage(damage int8) {
	self.Life -= damage
	if self.Life < 0 {
		self.Life = 0
	}
}

func (self *Player) Recover(v int8) {
	self.Life += v
	if self.Life > 50 {
		self.Life = 50
	}
}

func (self *Player) refreshWeight() {
	// 重新计算负重
}

func (self *Player) Pick(item Item) Info {
	self.refreshWeight()
	if self.Weight+item.Weight > 100 {
		// 返回失败，超过负重
	} else {
		// 放入包裹(当主手为空时，持于主手)
	}
	return Info{}
}

func (self *Player) Use(item Item) {

}

func (self *Player) Handhold(item Item) {

}

func (self *Player) Throw(item Item) {

}

// 好想用个map直接做啊，不会啊！！！

var DefaultActionList []string = []string{
	"mv", "atk",
}

func (self *Player) DefaultAction(action string, parameter []string) Info {
	switch action {
	case "mv":
	case "atk":
	}
	return Info{}
}

func (self *Player) attack(parameter []string) Info {
	fmt.Println(self.Id, "attack")
	return Info{}
}

func (self *Player) move(direction string) Info {
	fmt.Println(self.Id, "move")
	return Info{}
}
