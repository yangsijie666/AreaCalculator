package models

import (
	"AreaCalculator/types"
	"fmt"
	"github.com/shopspring/decimal"
	"testing"
)

func TestDisplay_DisplayResult(t *testing.T) {
	d := NewDisplay()
	d.currentCuboids = NewCuboids(decimal.NewFromInt(5), types.M, decimal.NewFromInt(2), types.CM, decimal.NewFromInt(3), types.DM, 2, Density{
		Value:      decimal.NewFromFloat(3.85),
		WeightUnit: types.G,
		VolumeUnit: types.MM,
	})
	fmt.Println(d.DisplayResult(types.M))
}

func Test1(t *testing.T) {
	d := NewDisplay()
	fmt.Println(d.TotalNumber())
}

func TestWeightDisplayer_DisplayWeightResult(t *testing.T) {
	d := NewWeightDisplayer()
	d.currentCuboids = NewCuboids(decimal.NewFromInt(5), types.M, decimal.NewFromInt(2), types.CM, decimal.NewFromInt(3), types.DM, 2, Density{
		Value:      decimal.NewFromFloat(3.85),
		WeightUnit: types.G,
		VolumeUnit: types.MM,
	})
	fmt.Println(d.DisplayWeightResult(types.M, types.KG))
}
