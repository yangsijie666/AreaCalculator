package models

import (
	"AreaCalculator/types"
	"fmt"
	"github.com/shopspring/decimal"
	"strconv"
)

type Display struct {
	currentCuboids CuboidInterface
	totalNumber    int
	sumArea        [4]decimal.Decimal
}

func NewDisplay() *Display {
	display := new(Display)
	display.currentCuboids = nil
	return display
}

func (d *Display) DisplayResult(unit types.Unit) string {
	if d.currentCuboids == nil {
		return types.DisplayResultInitStr
	}
	return fmt.Sprintf(d.currentCuboids.Display(unit)+types.DisplayResultStr, d.totalNumber, d.sumArea[unit].String())
}

func (d *Display) SumArea(unit types.Unit) string {
	return d.sumArea[unit].String()
}

func (d *Display) TotalNumber() string {
	return strconv.Itoa(d.totalNumber)
}

type WeightDisplayer struct {
	currentCuboids CuboidWeightInterface
	totalNumber    int
	sumWeight      [4]decimal.Decimal
}

func NewWeightDisplayer() *WeightDisplayer {
	displayer := new(WeightDisplayer)
	displayer.currentCuboids = nil
	return displayer
}

func (d *WeightDisplayer) DisplayWeightResult(unit types.Unit, weightUnit types.WeightUnit) string {
	if d.currentCuboids == nil {
		return types.WeightDisplayResultInitStr
	}
	return fmt.Sprintf(d.currentCuboids.DisplayWeight(unit, weightUnit)+types.WeightDisplayResultStr, d.totalNumber, d.sumWeight[weightUnit].String())
}

func (d *WeightDisplayer) SumWeight(unit types.WeightUnit) string {
	return d.sumWeight[unit].String()
}

func (d *WeightDisplayer) TotalNumber() string {
	return strconv.Itoa(d.totalNumber)
}
