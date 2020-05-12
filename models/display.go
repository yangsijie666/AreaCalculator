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
    return fmt.Sprintf(d.currentCuboids.Display(unit) + types.DisplayResultStr, d.totalNumber, d.sumArea[unit].String())
}

func (d *Display) SumArea(unit types.Unit) string {
    return d.sumArea[unit].String()
}

func (d *Display) TotalNumber() string {
    return strconv.Itoa(d.totalNumber)
}
