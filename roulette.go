package roulette

import (
	"math/rand"
	"time"
)

type pocket struct {
	Number int
	Color  string
}

type wheel struct {
	pockets []pocket
}

func (w wheel) spin(s int64) pocket {
	rand.Seed(s)
	n := rand.Intn(len(w.pockets))
	// return w.pockets[n].Number, w.pockets[n].Color
	return w.pockets[n]
}

func (w wheel) Spin(t int) chan pocket {
	c := make(chan pocket)
	//
	time.AfterFunc(time.Duration(t)*time.Second, func() {
		c <- w.spin(time.Now().UTC().UnixNano())
	})
	return c
}

func NewZWheel() wheel {
	// https://upload.wikimedia.org/wikipedia/en/a/ac/French_Layout-Single_Zero_Wheel.jpg
	order := []int{0, 32, 15, 19, 4, 21, 2, 25, 17, 34, 6, 27, 13, 36,
		11, 30, 8, 23, 10, 5, 24, 16, 33, 1, 20, 14, 31, 9, 22, 18, 29,
		7, 28, 12, 35, 3, 26}

	pp := make([]pocket, 37)
	for k, v := range order {
		pp[k].Number = v

		if v == 0 {
			pp[k].Color = "green"
		}

		// In number ranges from 1 to 10 and 19 to 28,
		//   odd numbers are red and even are black.
		if (v >= 1 && v <= 10) || (v >= 19 && v <= 28) {
			switch v % 2 {
			case 0:
				pp[k].Color = "black"
			case 1:
				pp[k].Color = "red"
			}
		}

		// In ranges from 11 to 18 and 29 to 36,
		//   odd numbers are black and even are red.
		if (v >= 11 && v <= 18) || (v >= 29 && v <= 36) {
			switch v % 2 {
			case 0:
				pp[k].Color = "red"
			case 1:
				pp[k].Color = "black"
			}
		}
	}
	w := wheel{pockets: pp}
	return w
}
