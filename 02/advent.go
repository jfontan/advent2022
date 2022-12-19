package advent02

import (
	"bufio"
	"io"
	"strings"
)

type Shape int

const (
	ShapeRock     Shape = 1
	ShapePaper    Shape = 2
	ShapeScissors Shape = 3

	ResultLose = "X"
	ResultDraw = "Y"
	ResultWin  = "Z"
)

func (s Shape) Value() int {
	return int(s)
}

var (
	translation = map[string]Shape{
		"A": ShapeRock,
		"B": ShapePaper,
		"C": ShapeScissors,
		"X": ShapeRock,
		"Y": ShapePaper,
		"Z": ShapeScissors,
	}

	wins = map[Shape]Shape{
		ShapeRock:     ShapeScissors,
		ShapePaper:    ShapeRock,
		ShapeScissors: ShapePaper,
	}

	loses = map[Shape]Shape{
		ShapeRock:     ShapePaper,
		ShapePaper:    ShapeScissors,
		ShapeScissors: ShapeRock,
	}
)

type Round struct {
	Me       Shape
	Opponent Shape
}

func (r Round) Score() int {
	score := r.Me.Value()
	if r.Me == r.Opponent {
		score += 3
	} else if r.Opponent == wins[r.Me] {
		score += 6
	}
	return score
}

func first(r io.Reader) int {
	scan := bufio.NewScanner(r)

	var score int
	for scan.Scan() {
		t := strings.TrimSpace(scan.Text())
		parts := strings.Split(t, " ")

		round := Round{
			Opponent: translation[parts[0]],
			Me:       translation[parts[1]],
		}
		score += round.Score()
	}

	return score
}

func second(r io.Reader) int {
	scan := bufio.NewScanner(r)

	var score int
	for scan.Scan() {
		t := strings.TrimSpace(scan.Text())
		parts := strings.Split(t, " ")

		opponent := translation[parts[0]]

		var me Shape
		switch parts[1] {
		case ResultLose:
			me = wins[opponent]
		case ResultDraw:
			me = opponent
		case ResultWin:
			me = loses[opponent]
		}

		round := Round{
			Opponent: opponent,
			Me:       me,
		}
		score += round.Score()

		// spew.Dump(round)
		// println(round.Score())
	}

	return score
}
