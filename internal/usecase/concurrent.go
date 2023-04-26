package usecase

import (
	"chessKnight/internal/entity"
	"fmt"
	"runtime"
)

var size = runtime.GOMAXPROCS(0)
var One = make(chan Data, size)
var Second = make(chan Data, size)

type Data struct {
	Position entity.Position
	Moves    []entity.Position
}

func AddMoves() {
	for d := range Second {
		if len(d.Moves) < 13 {
			d.Moves = append(d.Moves, d.Position)
			for {
				select {
				case One <- d:
				default:
				}
			}
		} else {
			fmt.Println(d.Moves)
			break
		}
	}
	close(One)
}
