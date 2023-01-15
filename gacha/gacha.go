package gacha

import (
	"math/rand"
	"time"
)

// const baseURL = "https://gohandson-gacha.uc.r.appspot.com/"

func init() {
	rand.Seed(time.Now().Unix())
}

type Play struct {
	player  *Player
	results []*Card
	summary map[Rarity]int
	err     error
}

func NewPlay(p *Player) *Play {
	return &Play{
		player:  p,
		summary: make(map[Rarity]int),
	}
}

func (p *Play) Results() []*Card {
	return p.results
}

func (p *Play) Result() *Card {
	if len(p.results) == 0 {
		return nil
	}
	return p.results[len(p.results)-1]
}

func (p *Play) Summary() map[Rarity]int {
	return p.summary
}

func (p *Play) Err() error {
	return p.err
}

func (p *Play) Draw() bool {
	if p.err != nil {
		return false
	}

	if err := p.player.draw(1); err != nil {
		p.err = err
		return false
	}

	card, err := p.draw()
	if err != nil {
		p.err = err
		return false
	}
	p.results = append(p.results, card)
	p.summary[card.Rarity]++

	return p.player.DrawableNum() > 0
}

func (p *Play) draw() (*Card, error) {
	num := rand.Intn(100)

	switch {
	case num < 50:
		return &Card{Rarity: RarityN, Name: "木の枝"}, nil
	case num < 75:
		return &Card{Rarity: RarityR, Name: "こん棒"}, nil
	case num < 90:
		return &Card{Rarity: RaritySR, Name: "鉄の剣"}, nil
	case num < 95:
		return &Card{Rarity: RaritySSR, Name: "炎の剣"}, nil
	case num < 99:
		return &Card{Rarity: RarityUR, Name: "闇の剣"}, nil
	default:
		return &Card{Rarity: RarityLR, Name: "爆炎神龍剣"}, nil
	}

	// q := "スライム:80,オーク:15,ドラゴン:4,イフリート:1"
	// req, err := http.NewRequest(http.MethodGet, baseURL+"?q="+q, nil)
	// if err != nil {
	// 	return nil, fmt.Errorf("リクエスト作成:%w", err)
	// }

	// resp, err := http.DefaultClient.Do(req)
	// if err != nil {
	// 	return nil, fmt.Errorf("APIリクエスト:%w", err)
	// }
	// defer resp.Body.Close()

	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	return nil, fmt.Errorf("Bodyの読み込み:%w", err)
	// }

	// result := string(body)
	// switch result {
	// case "スライム":
	// 	return &Card{Rarity: RarityN, Name: "スライム"}, nil
	// case "オーク":
	// 	return &Card{Rarity: RarityR, Name: "オーク"}, nil
	// case "ドラゴン":
	// 	return &Card{Rarity: RaritySR, Name: "ドラゴン"}, nil
	// default:
	// 	return &Card{Rarity: RarityXR, Name: "イフリート"}, nil
	// }
}
