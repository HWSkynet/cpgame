// method.go
package cpgame

import (
	//"fmt"
	"reflect"
	"strings"
)

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
	// 否则，查找第一个参数是否为玩家主手物品的接口
	// 如果是，则将剩余参数作为输入，调用物品的接口
	if self.Mainhand != nil {
		item := reflect.ValueOf(self.Mainhand)
		parameter[0] = strings.Title(parameter[0])
		if item.MethodByName(parameter[0]).IsValid() {
			method := item.MethodByName(parameter[0])
			params := make([]reflect.Value, 0)
			params = append(params, reflect.ValueOf(self))
			params = append(params, reflect.ValueOf(parameter[1:]))
			ret := method.Call(params)[0]
			return Info{int(ret.FieldByName("Statu").Int()), ret.FieldByName("Text").String()}
		}
	}
	return Info{-1, "YUI不知道你想要什么呢。。。"}
}
