package types

type Unit int

const (
	MM Unit = iota
	CM
	DM
	M
)

type Operator int

const (
	Plus Operator = iota
	Minus
)

type WeightUnit int

const (
	MG WeightUnit = iota
	G
	KG
	T
)
