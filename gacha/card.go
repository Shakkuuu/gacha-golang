package gacha

type Rarity string

const (
	RarityN   Rarity = "N"
	RarityR   Rarity = "R"
	RaritySR  Rarity = "SR"
	RaritySSR Rarity = "SSR"
	RarityUR  Rarity = "UR"
	RarityLR  Rarity = "LR"
)

func (r Rarity) String() string {
	return string(r)
}

type Card struct {
	Rarity Rarity
	Name   string
}

func (c *Card) String() string {
	return c.Rarity.String() + ":" + c.Name
}
