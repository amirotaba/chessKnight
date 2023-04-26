package entity

type Position struct {
	Row    int
	Column int
}

type Input struct {
	Len      int
	Position Position
}

type Data struct {
	Length   int
	Moves    []string
	Board    map[Position]bool
	Position Position
}
