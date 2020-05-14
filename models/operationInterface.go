package models

import (
	"AreaCalculator/types"
)

type OperationInterface interface {
	Operate(*Display)
	Operator() types.Operator
	Display(types.Unit) string
}

type WeightOperationInterface interface {
	Operate(*WeightDisplayer)
	Operator() types.Operator
	DisplayWeight(types.Unit, types.WeightUnit) string
}
