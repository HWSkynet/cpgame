// clock.go
package cpgame

import "time"

// 每回合时限3分钟
// 当所有人都耗尽行动点时，自动进入下回合
// 回合更新时，刷新行动点，DOT伤害生效，地图变更生效

var GameMinutes int = 0
var clockCtr = make(chan bool)
var roundSecond int = 180 // 回合剩余时间

func GameClockStart() {
	GameMinutes = 0
	go minute(clockCtr)
}

func GameClockEnd() {
	clockCtr <- true
}

func minute(input chan bool) {
	min := time.NewTicker(1 * time.Minute)
	tenSec := time.NewTicker(10 * time.Second)
	for {
		select {
		case <-tenSec.C:
			game10sHandler()
		case <-min.C:
			GameMinutes += 1
		case <-input:
			return
		}
	}
}

func GetRoundTime() int {
	return roundSecond
}

func energyFresh() {
	for _, v := range PL {
		if v.Alive {
			v.Energy = 5
		}
	}
}

func game10sHandler() {
	if PL.CountLive() <= 1 {
		GameClockEnd()
		return
	}
	var active int = 0 // 场上剩余行动点的玩家数
	for _, v := range PL {
		if v.Alive && v.Energy > 0 {
			active += 1
		}
	}
	roundSecond -= 10
	if roundSecond <= 0 || active == 0 {
		roundSecond = 180
		energyFresh()
	}
}
