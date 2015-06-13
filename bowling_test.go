package main

import "testing"

func TestScoreBeforeAnyRollIs0(t *testing.T) {
	game := NewGame()

	assertScore(t, 0, game)
}

func TestScoreForFirstRollIsNumberOfHit(t *testing.T) {
	game := NewGame()
	hit := 4
	game.Roll(hit)
	game.Roll(0)

	assertScore(t, hit, game)
}

func TestScoreForAGameIsTotalCountOfHit(t *testing.T) {
	game := NewGame()
	for i := 0; i < 20; i++ {
		game.Roll(4)
	}

	assertScore(t, 80, game)
}

func TestSpare(t *testing.T) {
	game := NewGame()
	game.Roll(5)
	game.Roll(5)
	game.Roll(4)
	game.Roll(3)

	assertScore(t, 21, game)
}

func TestStrike(t *testing.T) {
	game := NewGame()
	game.Roll(10)
	game.Roll(4)
	game.Roll(3)

	assertScore(t, 24, game)
}

func TestFullSpares(t *testing.T) {
	game := NewGame()
	for i := 0; i < 21; i++ {
		game.Roll(5)
	}

	assertScore(t, 150, game)
}

func TestFullStrikes(t *testing.T) {
	game := NewGame()
	for i := 0; i < 12; i++ {
		game.Roll(10)
	}

	assertScore(t, 300, game)
}

func assertScore(t *testing.T, expected int, game *Game) {
	score := game.Score()

	if score != expected {
		t.Errorf("Expected %v, got %v", expected, score)
	}
}
