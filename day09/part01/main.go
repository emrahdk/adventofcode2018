package main

import (
	"container/ring"
	"fmt"
)

func main() {
	players := 404
	marble := 71852

	scores := make([]int, players)

	current := &ring.Ring{Value: 0}
	current.Value = 0

	for i := 0; i <= marble; i++ {
		if i%23 == 0 {
			current = current.Move(-8)   // move to one before seventh
			removed := current.Unlink(1) // remove it and give as score
			scores[i%players] += i + removed.Value.(int)

			current = current.Move(1) // now at seventh
		} else {
			current = current.Move(1)

			next := ring.New(1)
			next.Value = i

			current.Link(next)
			current = current.Move(1)
		}
	}

	highestScore := 0
	for _, v := range scores {
		if v > highestScore {
			highestScore = v
		}
	}

	fmt.Println(highestScore)
}
