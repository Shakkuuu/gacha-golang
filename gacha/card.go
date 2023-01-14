package gacha

type Rarity string

const (
	RarityN  Rarity = "ノーマル"
	RarityR  Rarity = "R"
	RaritySR Rarity = "SR"
	RarityXR Rarity = "XR"
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
