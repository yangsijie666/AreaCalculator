package models

import (
	"AreaCalculator/types"
	"github.com/shopspring/decimal"
)

type CuboidInterface interface {
	Area(types.Unit) decimal.Decimal
	Number() int
	Display(types.Unit) string
}

type CuboidWeightInterface interface {
	Number() int
	Weight(types.WeightUnit) decimal.Decimal
	DisplayWeight(types.Unit, types.WeightUnit) string
}
