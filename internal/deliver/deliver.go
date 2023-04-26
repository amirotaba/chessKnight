package deliver

import (
	"chessKnight/internal/entity"
	"chessKnight/internal/usecase"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func Run() {
	//get board length and start position
	input, err := GetArg()
	if err != nil {
		fmt.Println(err)
		return
	}

	//make data struct
	data := usecase.MakeData(input)

	//Find correct path
	result, err := usecase.FindPath(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}

func GetArg() (entity.Input, error) {
	if len(os.Args) != 4 {
		return entity.Input{}, errors.New("wrong input format -> go run main.go {Board length} {Start point row} {Start point column}")
	}

	length, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return entity.Input{}, err
	}

	row, err := strconv.Atoi(os.Args[2])
	if err != nil {
		return entity.Input{}, err
	}

	column, err := strconv.Atoi(os.Args[3])
	if err != nil {
		return entity.Input{}, err
	}

	p := entity.Position{
		Row:    row,
		Column: column,
	}

	out := entity.Input{
		Len:      length,
		Position: p,
	}

	return out, nil
}
