package utils

import (
	"chessKnight/internal/entity"
	"math"
	"strconv"
)

func Moves(d entity.Data) []entity.Position {
	var position entity.Position
	var slice []entity.Position

	length := int(math.Sqrt(float64(d.Length)))

	position = entity.Position{
		Row:    d.Position.Row + 2,
		Column: d.Position.Column + 1,
	}
	if position.Row < length && position.Row > -1 && position.Column < length && position.Column > -1 {
		if !Exist(d.Used, position) {
			slice = append(slice, position)
		}
	}
	position = entity.Position{
		Row:    d.Position.Row + 2,
		Column: d.Position.Column - 1,
	}
	if position.Row < length && position.Row > -1 && position.Column < length && position.Column > -1 {
		if !Exist(d.Used, position) {
			slice = append(slice, position)
		}
	}
	position = entity.Position{
		Row:    d.Position.Row - 2,
		Column: d.Position.Column + 1,
	}
	if position.Row < length && position.Row > -1 && position.Column < length && position.Column > -1 {
		if !Exist(d.Used, position) {
			slice = append(slice, position)
		}
	}
	position = entity.Position{
		Row:    d.Position.Row - 2,
		Column: d.Position.Column - 1,
	}
	if position.Row < length && position.Row > -1 && position.Column < length && position.Column > -1 {
		if !Exist(d.Used, position) {
			slice = append(slice, position)
		}
	}
	position = entity.Position{
		Row:    d.Position.Row + 1,
		Column: d.Position.Column + 2,
	}
	if position.Row < length && position.Row > -1 && position.Column < length && position.Column > -1 {
		if !Exist(d.Used, position) {
			slice = append(slice, position)
		}
	}
	position = entity.Position{
		Row:    d.Position.Row + 1,
		Column: d.Position.Column - 2,
	}
	if position.Row < length && position.Row > -1 && position.Column < length && position.Column > -1 {
		if !Exist(d.Used, position) {
			slice = append(slice, position)
		}
	}
	position = entity.Position{
		Row:    d.Position.Row - 1,
		Column: d.Position.Column + 2,
	}
	if position.Row < length && position.Row > -1 && position.Column < length && position.Column > -1 {
		if !Exist(d.Used, position) {
			slice = append(slice, position)
		}
	}
	position = entity.Position{
		Row:    d.Position.Row - 1,
		Column: d.Position.Column - 2,
	}
	if position.Row < length && position.Row > -1 && position.Column < length && position.Column > -1 {
		if !Exist(d.Used, position) {
			slice = append(slice, position)
		}
	}

	return slice
}

func Corner(slice []entity.Position, c int) []entity.Position {
	var test bool
	var cSlice []int
	var conditions []entity.Position
	length := math.Sqrt(float64(c)) - 1
	for _, p := range slice {
		if math.Abs(float64(p.Row)-length/2) > math.Abs(float64(p.Column)-length/2) {
			if float64(p.Row) > length/2 {
				cSlice = append(cSlice, int(length)-p.Row)
			} else {
				cSlice = append(cSlice, p.Row)
			}
		} else {
			if float64(p.Column) > length/2 {
				cSlice = append(cSlice, int(length)-p.Column)
			} else {
				cSlice = append(cSlice, p.Column)
			}
		}
	}

	for i, n1 := range cSlice {
		test = true
		for _, n2 := range cSlice {
			if n1 > n2 {
				test = false
				break
			}
		}
		if test {
			conditions = append(conditions, slice[i])
		}
	}

	return conditions
}

func Convert(position entity.Position) string {
	return toStr(position.Column) + strconv.Itoa(position.Row+1)
}

func toStr(i int) string {
	return string(rune('A' + i))
}

func Exist(slice []entity.Position, position entity.Position) bool {
	var out bool

	for _, p := range slice {
		if p == position {
			out = true
		}
	}

	return out
}

func ExistCon(slice [][]string, moves []string) bool {
	var out bool

	for _, m := range slice {
		out = true
		for i := range moves {
			if moves[i] != m[i] {
				out = false
			}
		}
	}

	return out
}
