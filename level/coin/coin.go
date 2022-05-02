package coin

type Coin struct {
	Coins []*Coins
}

type Coins struct {
	X, Y        int
	Uncollected bool
}

func New() Coin {
	var c Coin

	c.Coins = make([]*Coins, 0)

	return c
}
