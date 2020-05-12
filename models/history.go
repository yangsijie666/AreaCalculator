package models

import (
    "AreaCalculator/types"
)

type History struct {
    operation OperationInterface
}

func NewHistory(operation OperationInterface) *History {
    history := new(History)
    history.operation = operation
    return history
}

func (h *History) Display(unit types.Unit) string {
    return h.operation.Display(unit)
}
