package models

import (
	"AreaCalculator/types"
	"fmt"
	"github.com/shopspring/decimal"
	"testing"
)

func TestCubes_Display(t *testing.T) {
	value, _ := decimal.NewFromString("0.1")
	cubes := NewCubes(decimal.NewFromInt(2), types.CM, 2, Density{
		Value:      value,
		WeightUnit: types.G,
		VolumeUnit: types.CM,
	})
	fmt.Println(cubes.Display(types.MM))
	fmt.Println(cubes.Display(types.CM))
	fmt.Println(cubes.Display(types.DM))
	fmt.Println(cubes.Display(types.M))
	fmt.Println(cubes.DisplayWeight(types.CM, types.G))
	fmt.Println(cubes.DisplayWeight(types.CM, types.KG))
	fmt.Println(cubes.DisplayWeight(types.CM, types.T))
	fmt.Println(cubes.DisplayWeight(types.CM, types.MG))
}
