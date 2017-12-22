// map.go
package cpgame

type Coord struct {
	x      int8
	y      int8
	detail string
}

type Field struct {
	Players []Player
	Items   []Item
}

type Room struct {
	W Field
	A Field
	S Field
	D Field
}

type Map struct {
}
