package main

type Player struct {
	Person
	Scores []Score
}

func (p *Player) makeScore(scoreValue uint) {
	var score = Score{Value: scoreValue, player: *p}
	p.Scores = append(p.Scores, score)
}

func (p Player) calculateScore() uint {
	var result uint = 0

	for _, score := range p.Scores {
		result += score.Value
	}

	return result
}
