package rps

import (
	"math/rand"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2

	PLAYERWINS   = 1
	COMPUTERWINS = 2
	DRAW         = 3
)

type Round struct {
	Winner         int    `json:"winner"`
	ComputerChoice string `json:"computerChoice"`
	RoundResult    string `json:"roundResult"`
}

func PlayRound(playerValue int) Round {
	computerValue := rand.Intn(3)
	computerChoice := ""
	roundResult := ""
	winner := 0

	switch computerValue {
	case ROCK:
		computerChoice = "Computer chose ROCK"
	case PAPER:
		computerChoice = "Computer chose PAPER"
	case SCISSORS:
		computerChoice = "Computer chose SCISSORS"
	default:
	}

	if playerValue == computerValue {
		roundResult = "It's a draw!"
		winner = DRAW
	} else if playerValue == (computerValue+1)%3 {
		roundResult = "Player wins!"
		winner = PLAYERWINS
	} else {
		roundResult = "Computer wins!"
		winner = COMPUTERWINS
	}

	result := Round{
		Winner:         winner,
		ComputerChoice: computerChoice,
		RoundResult:    roundResult,
	}

	return result
}
