package entity

type CharNum struct {
	ChCount    int
	NumCount   int
	SpaceCount int
	OtherCount int
}

type Student struct {
	Name  string  `json:"name"`
	Age   int     `json:"age"`
	Grade float64 `json:"grade"`
}
