package usecase

import (
	"chessKnight/internal/entity"
	"chessKnight/internal/utils"
	"errors"
	"fmt"
	"sync"
)

func MakeData(i entity.Input) {
	data := entity.Data{
		Position: i.Position,
		Length:   i.Len * i.Len,
	}

	entity.DataChan = make(chan entity.Data, data.Length)

	entity.DataChan <- data
	return
}

func FindPath() {
	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
}

func Append(data entity.Data) entity.Data {
	var out entity.Data
	out = data
	out.Moves = append(out.Moves, utils.Convert(out.Position))
	if len(data.Moves) < data.Length*2-2 {
		out.Moves = append(out.Moves, "->")
	}

	return out
}

func LessMoves(data entity.Data) ([]entity.Data, error) {
	var slice []entity.Position
	var dataSlice []entity.Data
	var m = make(map[entity.Position]int)
	moves := utils.Moves(data)
	if moves == nil && len(data.Moves) < data.Length*2-4 {
		//todo move this to all condtions
		//str := strconv.Itoa(int(math.Sqrt(float64(data.Length)))) + "*" + strconv.Itoa(int(math.Sqrt(float64(data.Length))))
		return nil, errors.New("there's no more move from " + utils.Convert(data.Position))
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
		cons := utils.Corner(slice, data.Length)
		slice = nil
		for _, c := range cons {
			data.Position = c
			dataSlice = append(dataSlice, data)
		}
	} else if len(slice) != 0 {
		data.Position = slice[0]
		dataSlice = append(dataSlice, data)
	}

	return dataSlice, nil
}

func Unavailable(data entity.Data) entity.Data {
	data.Used = append(data.Used, data.Position)
	return data
}

func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	var err error
	var newData entity.Data
	var dataSlice []entity.Data
	for {
		select {
		case data := <-entity.DataChan:
			newData = Unavailable(data)
			newData = Append(newData)
			test := entity.AllConditions
			_ = test
			if len(newData.Moves) < newData.Length*2-1 {
				dataSlice, err = LessMoves(newData)
				if err != nil {
					fmt.Println(err)
					break
				}

				for _, d := range dataSlice {
					select {
					case entity.DataChan <- d:
					default:
					}
				}
			} else {
				if !utils.ExistCon(entity.AllConditions, newData.Moves) || len(entity.AllConditions) < 1 {
					entity.AllConditions = append(entity.AllConditions, newData.Moves)
				}
			}
		default:
			return
		}
	}
}
