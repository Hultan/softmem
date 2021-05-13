package softmem

import (
	"errors"
	"fmt"
	memoDatabase "github.com/hultan/softmem/internal/database"
	"math/rand"
	"time"
)

type NumberPicker struct {
	database *memoDatabase.Database
	numbers  map[string]memoDatabase.NumberTable
}

func NewNumberPicker() *NumberPicker {
	n := new(NumberPicker)

	// Load PEGs
	n.database = new(memoDatabase.Database)
	numbers, err := n.database.GetAllNumbers()
	errorCheck(err)
	n.numbers = numbers

	// Set a seed for RNG
	rand.Seed(time.Now().UTC().UnixNano())

	return n
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

func (n *NumberPicker) GetNextNumber() (memoDatabase.NumberTable, error) {
	_, low := n.getCorrectStatistics()
	var values []string
	for _, item := range n.numbers {
		count := item.Correct - low + 1
		for i := 0; i < n.limit(count); i++ {
			values = append(values, item.Number)
		}
	}

	// get next number
	next := rand.Intn(len(values))
	if len(values)==0 {
		return memoDatabase.NumberTable{}, errors.New("slice is empty")
	}
	number := n.numbers[values[next]]
	return number, nil
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
