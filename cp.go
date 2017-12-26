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
func NewGame() {
	rand.Seed(time.Now().UnixNano())
}

func RandInt(n int) int {
	return rand.Intn(n)
}
func Version() string {
	return "0.0.0"
}
