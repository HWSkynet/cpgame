// method.go
package cpgame

import (
	//"fmt"
	"strings"
)

type Info struct {
	Statu int // <0表示错误 >0表示时间(秒)
	Text  string
}

type Method func([]string) Info

func InputParse(input string) []string {
	return strings.Fields(input)
}

func (self *Player) InputParse(input string) Info {
	parameter := InputParse(input)
	if len(parameter) <= 0 {
		return Info{-1, "invalid parameter"}
	}
	// 先识别第一个参数是否是默认动作
	// 如果是，则尝试执行动作
	for _, v := range DefaultActionList {
		if parameter[0] == v {
			return self.DefaultAction(parameter[0], parameter[1:])
		}
	}
	// 否则，查找第一个参数是否在玩家背包中
	// 如果是，则将剩余参数作为输入，调用物品的接口
	return Info{}
}
