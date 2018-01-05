// items.go
package cpgame

type Item struct {
	Name    string
	Weight  int
	Belongs string // 所属主体，为玩家或房间
}

func (self *Item) moveto(p string) {
	if PL.ID(p).Alive {
		self.Belongs = p
	}
}

func (self *Item) pickedby(p string) {
	if PL.ID(p).Alive {
		self.Belongs = p
		PL.ID(p).Pocket = append(PL.ID(p).Pocket, self)
	}
}

func (self *Item) Throw(py *Player, input []string) Info {
	if len(input) == 0 {
		return Info{Statu: -1, Text: "没有有效参数,请用help查询用法"}
	}
	target := PL.NAME(input[0])
	if target.Alive {
		// 检测目标是否在交互范围内
		// s := TargetTest(target)
		// 投掷后主手清空
		py.Mainhand = nil
		r := Dice("2d3").Value()
		if r > 4 {
			// 改变物品归属主体，背包物品由背包管理刷新
			self.pickedby(target.Id)
			return Info{Statu: 0, Text: "你向" + target.Name + "投掷了" + self.Name + "，并正中眉心"}
		} else {
			self.moveto(target.Id)
			return Info{Statu: 0, Text: "你向" + target.Name + "投掷了" + self.Name + "，但是没有击中"}
		}
	} else {
		return Info{Statu: -1, Text: "无效的目标名称"}
	}
	return Info{Text: "test"}
}

type Knife struct {
	Item
}

func (self *Knife) Slice(py *Player, input []string) Info {
	if len(input) == 0 {
		return Info{Statu: -1, Text: "未识别到有效输入"}
	}
	return Info{Text: "test"}
}

type Gun struct {
	Item
	ExUsages map[string]Usage
}

func (self *Item) Use(py *Player, input []string) Info {
	if len(input) == 0 {
		return Info{Statu: -1, Text: "未识别到有效输入"}
	}
	return Info{}
}

type Usage struct {
	text string
	use  func([]string) Info
}

type Info struct {
	Statu int // <0表示错误
	Text  string
}

// name:时间胶囊
// desc:一个小小的胶囊，里面有一小撮时间
// efect:使用后，增加1D3个行动点

// name:弓
// 可使用一个行动点蓄力，增加伤害

// type:枪
// 可使用一个行动点瞄准，增加命中
