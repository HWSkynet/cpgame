// items.go
package cpgame

type Item struct {
	Name   string
	Weight int
	Usage  map[string]Method
}

var Throw Method = func([]string) Info {
	return Info{Text: "test"}
}

var Knife Item = Item{
	Name:   "铅笔刀",
	Weight: 1,
	Usage: map[string]Method{
		"投掷": Throw,
	},
}

// name:时间胶囊
// desc:一个小小的胶囊，里面有一小撮时间
// efect:使用后，增加1D3个行动点
