// map.go
package cpgame

type Xy struct {
	X int
	Y int
}

type Coord struct {
	X      int8
	Y      int8
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

type Map map[Xy]Field

func NewMap() Map {
	x := Map{
		{0, 0}: {},
		{1, 0}: {},
		{2, 0}: {},
		{0, 1}: {},
		{1, 1}: {},
		{2, 1}: {},
		{0, 2}: {},
		{1, 2}: {},
		{2, 2}: {},
	}
	return x
}
