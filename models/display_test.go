package models

import (
    "AreaCalculator/types"
    "fmt"
    "github.com/shopspring/decimal"
    "testing"
)

func TestDisplay_DisplayResult(t *testing.T)  {
    d := NewDisplay()
    d.currentCuboids = NewCuboids(decimal.NewFromInt(5), types.M, decimal.NewFromInt(2), types.CM, decimal.NewFromInt(3), types.DM, 1)
    fmt.Println(d.DisplayResult(types.M))
}

func Test1(t *testing.T)  {
    d := NewDisplay()
    fmt.Println(d.TotalNumber())
}
