package model

type BarItem struct {
	Time   int     `json:"time"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Open   float64 `json:"open"`
	Close  float64 `json:"close"`
	Volume int     `json:"volume"`
}

type Bar struct {
	ID uint `gorm:"primary_key"`
	BarItem
	SymbolID     uint
	ResolutionID uint
}

type Bars struct {
	Symbol     string    `json:"symbol"`
	Resolution string    `json:"resolution"`
	Items      []BarItem `json:"items"`
}
