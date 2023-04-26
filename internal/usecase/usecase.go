package usecase

import (
	"chessKnight/internal/entity"
	"chessKnight/internal/utils"
	"errors"
	"math"
	"strconv"
)

func MakeData(i entity.Input) entity.Data {
	data := entity.Data{
		Length:   i.Len * i.Len,
		Position: i.Position,
		Board:    make(map[entity.Position]bool, i.Len*i.Len),
	}
	return data
}

func FindPath(data entity.Data) ([]string, error) {
	var err error
	for {
		if len(data.Moves) < data.Length*2-1 {
			data = Unavailable(data)
			data = Append(data)
			data, err = LessMoves(data)
			if err != nil {
				return nil, err
			}
		} else {
			return data.Moves, nil
		}
	}
}

func Append(data entity.Data) entity.Data {
	data.Moves = append(data.Moves, utils.Convert(data.Position))
	if len(data.Moves) < data.Length*2-2 {
		data.Moves = append(data.Moves, "->")
	}

	return data
}

func LessMoves(data entity.Data) (entity.Data, error) {
	var slice []entity.Position
	var m = make(map[entity.Position]int)
	moves := utils.Moves(data)
	if moves == nil && len(data.Moves) < data.Length*2-4 {
		str := strconv.Itoa(int(math.Sqrt(float64(data.Length)))) + "*" + strconv.Itoa(int(math.Sqrt(float64(data.Length))))
		return entity.Data{}, errors.New("It's not possible on " + str + " chess board.")
	}
	for _, pos := range moves {
		data.Position = pos
		m[pos] = len(utils.Moves(data))
	}
	for p, n := range m {
		var test = true
		for _, n2 := range m {
			if n > n2 {
				test = false
			}
		}
		if test {
			slice = append(slice, p)
		}
	}
	if len(slice) > 1 {
		data.Position = utils.Corner(slice, data.Length)
	} else if len(slice) != 0 {
		data.Position = slice[0]
	}

	return data, nil
}

func Unavailable(c entity.Data) entity.Data {
	c.Board[c.Position] = true
	return c
}
