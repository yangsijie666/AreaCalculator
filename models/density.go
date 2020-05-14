package models

import (
	"AreaCalculator/types"
	"github.com/shopspring/decimal"
)

type Density struct {
	Value      decimal.Decimal
	WeightUnit types.WeightUnit
	VolumeUnit types.Unit
}
