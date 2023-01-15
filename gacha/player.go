package gacha

import (
	"errors"
)

type Player struct {
	tickets int
	coin    int
}

func NewPlayer(tickets, coin int) *Player {
	return &Player{tickets: tickets, coin: coin}
}

func (p *Player) DrawableNum() int {
	return p.tickets + p.coin/10
}

func (p *Player) draw(n int) error {

	if p.DrawableNum() < n {
		return errors.New("ガチャ券またはコインが不足しています")
	}

	// ガチャ券から優先的に使う
	if p.tickets > n {
		p.tickets -= n
		return nil
	}

	p.tickets = 0
	p.coin -= n * 10 // 1回あたり10枚消費する

	return nil
}
