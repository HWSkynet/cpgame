// player.go
package cpgame

import "fmt"

type Player struct {
	Id       string
	Name     string
	State    string
	Life     int    // 生命值0-50 显示为0.0-5.0
	Energy   int    // 行动点数
	Weight   int    // 当前负重(最大100)
	Killed   int    // 杀敌数
	Assist   int    // 助攻数
	Pos      Coord  // 坐标
	Mainhand Item   // 物品
	Pocket   []Item // 包裹
}

func NewPlayer(id string) *Player {
	return &Player{Id: id, Life: 50, Energy: 5, State: "init"}
}

func (self *Player) GetDamage(damage int) {
	self.Life -= damage
	if self.Life < 0 {
		self.Life = 0
	}
}

func (self *Player) Recover(v int) {
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
	if len(parameter) == 0 {
		return Info{-1, fmt.Sprintf("错误的指令格式")}
	}
	switch action {
	case "mv":
		return self.move(parameter[0])
	case "atk":
		return self.attack(parameter)
	}
	return Info{}
}

func (self *Player) attack(parameter []string) Info {
	if len(parameter) == 0 {
		return Info{-1, fmt.Sprintf("没有指定atk目标")}
	}
	p := PL.NAME(parameter[0])
	if self.Energy < 1 {
		return Info{-1, fmt.Sprintf("已经没有足够的行动点来进行攻击了")}
	}
	if p.Life > 0 {
		dmg := Dice("3d3").Value()
		p.GetDamage(dmg)
		self.Energy -= 1
		if p.Life > 0 {
			return Info{5, fmt.Sprintf("`%s`使用小拳拳锤击了`%s`的胸口，造成`%d`点伤害", self.Name, p.Name, dmg)}
		} else {
			return Info{5, fmt.Sprintf("`%s`使用小拳拳锤击`%s`的胸口,造成`%s`五脏六腑爆裂而亡", self.Name, p.Name, p.Name)}
		}
	}
	return Info{-1, fmt.Sprintf("尝试攻击的目标`%s`不在这个时间线，或者已没有生命迹象", parameter[0])}
}

func (self *Player) move(direction string) Info {
	if self.Energy < 1 {
		return Info{-1, "本回合你已经没有可用的行动点数了"}
	}
	switch direction {
	case "w":
		// if Map.Valid(self.Pos.x,self.Pos.y-1)
		self.Pos.y -= 1
		self.Pos.detail = "s"
	case "a":
		self.Pos.x -= 1
		self.Pos.detail = "d"
	case "s":
		self.Pos.y += 1
		self.Pos.detail = "w"
	case "d":
		self.Pos.x += 1
		self.Pos.detail = "a"
	default:
		return Info{-1, "无效的移动参数[w,a,s,d]"}
	}
	self.Energy -= 1
	return Info{5, fmt.Sprintf("成功移动到坐标[%d,%d]", self.Pos.x, self.Pos.y)}
}
