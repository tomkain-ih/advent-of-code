package main

// A / B / C opponent picks
// X / Y / Z my picks
// rock (1)  / paper (2)  / scissors (3)
// loss (0) / draw (3) / win (6)

type Play struct {
	Choice string
	Value  int
}

func (p *Play) NewPlay(abbrev string) {
	switch abbrev {
	case "A":
		p.Choice = "rock"
		p.Value = 1
	case "B":
		p.Choice = "paper"
		p.Value = 2
	case "C":
		p.Choice = "scissors"
		p.Value = 3
	case "X":
		p.Choice = "rock"
		p.Value = 1
	case "Y":
		p.Choice = "paper"
		p.Value = 2
	case "Z":
		p.Choice = "scissors"
		p.Value = 3
	}
}

func (p *Play) winningValue() int {
	if p.Value == 3 {
		return 1
	}
	return p.Value + 1
}

func (p *Play) losingValue() int {
	if p.Value == 1 {
		return 3
	}
	return p.Value - 1
}

func (p *Play) Versus(opponent *Play) int {
	if p.Value == opponent.Value {
		return p.Value + 3
	}
	if (p.Value == 1 && opponent.Value == 3) ||
		(p.Value == 2 && opponent.Value == 1) ||
		(p.Value == 3 && opponent.Value == 2) {
		return p.Value + 6
	}
	return p.Value
}
