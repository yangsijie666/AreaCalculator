package models

import (
	"AreaCalculator/types"
	"strings"
)

type HistoryDisplay struct {
	histories []*History
}

func NewHistoryDisplay() *HistoryDisplay {
	historyDisplay := new(HistoryDisplay)
	historyDisplay.histories = make([]*History, 0, 100)
	return historyDisplay
}

func (h *HistoryDisplay) AppendHistory(history *History) {
	// TODO: 增加灵活调整历史记录数功能，暂时最大数为 100
	if len(h.histories) == 100 {
		h.histories = append([]*History{}, h.histories[1:]...)
	}
	h.histories = append(h.histories, history)
}

func (h *HistoryDisplay) Display(unit types.Unit) string {
	if len(h.histories) == 0 {
		return types.HistoryDisplayDefaultStr
	}

	var s strings.Builder
	for i := len(h.histories) - 1; i >= 0; i-- {
		s.WriteString(h.histories[i].Display(unit))
	}
	return s.String()
}

type WeightHistoryDisplayer struct {
	histories []*WeightHistory
}

func NewWeightHistoryDisplayer() *WeightHistoryDisplayer {
	weightHistoryDisplay := new(WeightHistoryDisplayer)
	weightHistoryDisplay.histories = make([]*WeightHistory, 0, 100)
	return weightHistoryDisplay
}

func (h *WeightHistoryDisplayer) AppendWeightHistory(weightHistory *WeightHistory) {
	// TODO: 增加灵活调整历史记录数功能，暂时最大数为 100
	if len(h.histories) == 100 {
		h.histories = append([]*WeightHistory{}, h.histories[1:]...)
	}
	h.histories = append(h.histories, weightHistory)
}

func (h *WeightHistoryDisplayer) Display(unit types.Unit, weightUnit types.WeightUnit) string {
	if len(h.histories) == 0 {
		return types.HistoryDisplayDefaultStr
	}

	var s strings.Builder
	for i := len(h.histories) - 1; i >= 0; i-- {
		s.WriteString(h.histories[i].DisplayWeight(unit, weightUnit))
	}
	return s.String()
}
