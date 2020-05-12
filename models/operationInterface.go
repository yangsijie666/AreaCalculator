package models

import (
    "AreaCalculator/types"
)

type OperationInterface interface {
    Operate(*Display)
    Operator() types.Operator
    Display(types.Unit) string
}
