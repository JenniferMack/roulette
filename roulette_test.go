package roulette

import (
	"testing"
	"time"
)

func TestExpSpin(t *testing.T) {
	t.Skip()
	w := NewZWheel()
	tm := time.Now()

	c := w.Spin(2)
	v := <-c
	t.Error(time.Since(tm), v.Color, v.Number)
}

func TestSpin(t *testing.T) {
	w := NewZWheel()
	p := w.spin(12345678)
	if p.Number != 15 && p.Color != "black" {
		t.Error(p)
	}
}
func TestGreen(t *testing.T) {
	w := NewZWheel()
	g, b, r := 0, 0, 0
	for _, v := range w.pockets {
		switch v.Color {
		case "green":
			g++
		case "red":
			r++
		case "black":
			b++
		}
	}
	t.Run("green", func(t *testing.T) {
		if g != 1 {
			t.Error(g)
		}
	})
	t.Run("red", func(t *testing.T) {
		if r != 18 {
			t.Error(g)
		}
	})
	t.Run("black", func(t *testing.T) {
		if b != 18 {
			t.Error(g)
		}
	})
	//spot checks
	t.Run("pocket 0", func(t *testing.T) {
		if p := w.pockets[0]; p.Color != "green" && p.Number != 0 {
			t.Error(w.pockets[0])
		}
	})
	t.Run("pocket 7", func(t *testing.T) {
		if p := w.pockets[7]; p.Color != "red" && p.Number != 25 {
			t.Error(w.pockets[7])
		}
	})
	t.Run("pocket 14", func(t *testing.T) {
		if p := w.pockets[14]; p.Color != "black" && p.Number != 11 {
			t.Error(w.pockets[14])
		}
	})
	t.Run("pocket 2", func(t *testing.T) {
		if p := w.pockets[21]; p.Color != "red" && p.Number != 16 {
			t.Error(w.pockets[21])
		}
	})
	t.Run("pocket 28", func(t *testing.T) {
		if p := w.pockets[28]; p.Color != "black" && p.Number != 22 {
			t.Error(w.pockets[28])
		}
	})
	t.Run("pocket 35", func(t *testing.T) {
		if p := w.pockets[35]; p.Color != "red" && p.Number != 3 {
			t.Error(w.pockets[35])
		}
	})
}

func TestSpinChart(t *testing.T) {
	w := NewZWheel()
	resCol := make(map[string]int64)
	resNum := make(map[int]int64)

	rounds := float64(10000)
	for i := 0; i < int(rounds); i++ {
		p := w.spin(time.Now().UnixNano())
		resNum[p.Number] += 1
		resCol[p.Color] += 1
	}
	sumf := 0.0
	t.Logf("results for %.f rounds:", rounds)
	for k, v := range resCol {
		pct := (float64(v) / rounds) * 100
		t.Logf("%6s : %5d : %5.2f%%\n", k, v, pct)
		sumf += pct
	}
	t.Logf("Total: %.3f", sumf)
	if sumf < 99.999 {
		t.Error("Color sums low:", sumf)
	}
	sumf = 0.0
	for i := 0; i < len(resNum); i++ {
		pct := (float64(resNum[i]) / rounds) * 100
		t.Logf("%3d : %d : %5.2f%%\n", i, resNum[i], pct)
		sumf += pct
	}
	t.Logf("Total: %.3f", sumf)
	if sumf < 99.999 {
		t.Error("Number sums low:", sumf)
	}
}
