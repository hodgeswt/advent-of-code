package util

import "errors"

var Done = errors.New("Done")

type Queue struct {
	Data []rune
}

func (it *Queue) Enqueue(r rune) {
    it.Data = append(it.Data, r)
}

func (it *Queue) Push(r rune) {
	it.Data = append([]rune{r}, it.Data...)
}

func (it *Queue) Dequeue() (rune, error) {
	if len(it.Data) == 0 {
		return '#', Done
	}

	r := it.Data[0]

	it.Data = it.Data[1:]

	return r, nil
}
