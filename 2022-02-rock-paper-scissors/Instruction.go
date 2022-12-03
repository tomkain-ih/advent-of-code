package main

type Instruction struct {
	Opponent *Play
	Result   string
}

func (i *Instruction) NewInstruction(opponentAbbrev string, result string) {
	opponent := new(Play)
	opponent.NewPlay(opponentAbbrev)
	i.Opponent = opponent
	i.Result = result
}

func (i *Instruction) getScore() int {
	switch i.Result {
	case "X": //lose
		return i.Opponent.losingValue()
	case "Y": //draw
		return i.Opponent.Value + 3
	case "Z": //win
		return i.Opponent.winningValue() + 6
	default:
		return 0
	}
}
