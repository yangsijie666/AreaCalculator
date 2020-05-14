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

type WeightHistory struct {
	operation WeightOperationInterface
}

func NewWeightHistory(operation WeightOperationInterface) *WeightHistory {
	weightHistory := new(WeightHistory)
	weightHistory.operation = operation
	return weightHistory
}

func (h *WeightHistory) DisplayWeight(unit types.Unit, weightUnit types.WeightUnit) string {
	return h.operation.DisplayWeight(unit, weightUnit)
}
