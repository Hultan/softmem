package mainWindow

import (
	"errors"
	"fmt"
	memoDatabase "github.com/hultan/softmemo/database"
	"math/rand"
	"time"
)

type NumberPicker struct {
	database *memoDatabase.Database
	current  *memoDatabase.NumberTable
	numbers  []memoDatabase.NumberTable
}

func NewNumberPicker() *NumberPicker {
	return new(NumberPicker)
}

func (n *NumberPicker) Initialize() {
	// Load PEGs
	n.database = new(memoDatabase.Database)
	numbers, err := n.database.GetAllNumbers()
	errorCheck(err)

	n.numbers = numbers

	// Set a seed for RNG
	rand.Seed(time.Now().UTC().UnixNano())
}

func (n *NumberPicker) UpdateStatistics() {
	for _, number := range n.numbers {
		if number.HasChanged {
			err := n.database.UpdateStatistics(number)
			if err != nil {
				fmt.Printf("Failed to update statistics for : %s", number.Number)
			}
		}
	}
}

func (n *NumberPicker) limit(number int) int {
	if number<1 {
		return 1
	}
	if number>5 {
		return 5
	}

	return number
}

func (n *NumberPicker) GetNextNumber() (string, error) {
	_, low := n.getCorrectStatistics()
	var values []string
	for _, item := range n.numbers {
		count := item.Correct - low + 1
		for i := 0; i < n.limit(count); i++ {
			values = append(values, item.Number)
		}
	}

	// get next number
	length := len(values)
	for {
		next := rand.Intn(length)
		nextNumber := values[next]
		var currentNumber string
		if n.current == nil {
			currentNumber = ""
		} else {
			currentNumber = n.current.Number
		}
		var index = 0
		for idx, value := range n.numbers {
			if value.Number == nextNumber && nextNumber != currentNumber {
				n.current = &n.numbers[idx]
				return value.Number, nil
			}
			index++
		}
	}

	return "", errors.New("GetNextNumber failed")
}

func (n *NumberPicker) getCorrectStatistics() (int, int) {
	highest := 0
	lowest := MaxInt
	for _, item := range n.numbers {
		if item.Correct > highest {
			highest = item.Correct
		}
		if item.Correct < lowest {
			lowest = item.Correct
		}
	}
	return highest, lowest
}
