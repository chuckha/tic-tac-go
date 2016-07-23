package player

import (
	"fmt"
)

type Player struct {
	id int
}

func New(id int) (*Player, error) {
	if id == 0 {
		return nil, fmt.Errorf("Player number cannot be 0.")
	}
	return &Player{
		id: id,
	}, nil
}

func (p *Player) Id() int {
	return p.id
}
