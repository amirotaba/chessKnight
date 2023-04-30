package entity

var DataChan = make(chan Data)

var AllConditions [][]string

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
	Position Position
	Used     []Position
}
