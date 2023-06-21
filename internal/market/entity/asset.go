package entity

type Asset struct{
	ID 		string
	Name 	string
	MarketVolume int
}

func NewAsset  (ID string, Name string, MarketVolume int) *Asset{
	return &Asset{
		ID:         ID,
		Name:     Name,
		MarketVolume: MarketVolume,
	}
}
