// cp.go
package cpgame

import (
	//"fmt"
	"math/rand"
	"strconv"
	//"strings"
	"time"
)
import "regexp"

// GN for GameNotice
const (
	GN_Success int = iota
)

type Dice string

var diceexp *regexp.Regexp

func NewGame() {
	diceexp, _ = regexp.Compile("(\\d+)(?:D|d)(\\d+)")
	rand.Seed(time.Now().UnixNano())
}

func RandInt(n int) int {
	return rand.Intn(n)
}

func (self Dice) Value() (sum int) {
	strs := diceexp.FindStringSubmatch(string(self))
	// strs[0]似乎是占位
	if len(strs) < 3 {
		panic("xDy格式错误")
	}
	x, _ := strconv.Atoi(strs[1])
	y, _ := strconv.Atoi(strs[2])
	//fmt.Printf("%d|%dD%d\n", len(strs), x, y)
	sum = 0
	for i := 0; i < x; i++ {
		sum += rand.Intn(y) + 1
	}
	return
}
func (self Dice) Add(a int) (sum int) {
	sum = self.Value() + a
	return
}

func Version() string {
	return "0.0.0"
}
