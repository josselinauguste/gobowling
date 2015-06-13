package main

type Game struct {
	hits []int
}

func NewGame() *Game {
	return &Game{make([]int, 0, 22)}
}

func (g *Game) Roll(hit int) {
	g.hits = append(g.hits, hit)
}

func (g *Game) Score() int {
	score := 0
	frame := 0
	for i := 0; i < len(g.hits) && frame < 10; frame += 1 {
		if g.isStrike(i) {
			i, score = g.calculateStrike(i, score)
		} else if g.isSpare(i) {
			i, score = g.calculateSpare(i, score)
		} else {
			i, score = g.calculate(i, score)
		}
	}
	return score
}

func (g *Game) isStrike(index int) bool {
	return g.hits[index] == 10
}

func (g *Game) calculateStrike(index int, score int) (int, int) {
	score += 10 + g.frameHits(index+1, 2)
	index += 1
	return index, score
}

func (g *Game) isSpare(index int) bool {
	return g.frameHits(index, 2) == 10
}

func (g *Game) calculateSpare(index int, score int) (int, int) {
	index, score = g.calculate(index, score)
	score += g.frameHits(index, 1)
	return index, score
}

func (g *Game) calculate(index int, score int) (int, int) {
	score += g.frameHits(index, 2)
	index += 2
	return index, score
}

func (g *Game) frameHits(index int, size int) int {
	hits := 0
	for i := index; i < index+size; i++ {
		if i < len(g.hits) {
			hits += g.hits[i]
		}
	}
	return hits
}
