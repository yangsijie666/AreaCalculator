package models

import (
	"AreaCalculator/types"
	"fmt"
	"github.com/shopspring/decimal"
	"testing"
)

func TestNewCubes(t *testing.T) {
	defer func() {
		errMsg := recover().(string)
		fmt.Println(errMsg)
	}()
	c := NewCuboids(decimal.NewFromInt(5), types.M, decimal.NewFromInt(2), types.CM, decimal.NewFromInt(3), types.DM, 1, Density{
		Value:      decimal.Zero,
		WeightUnit: types.G,
		VolumeUnit: types.M,
	})
	fmt.Printf("长: %s mm, 宽 %s mm, 高 %s mm, 面积 %s mm^2\n",
		c.Length(types.MM).String(), c.Width(types.MM).String(), c.Height(types.MM).String(), c.Area(types.MM).String())
	fmt.Printf("长: %s cm, 宽 %s cm, 高 %s cm, 面积 %s cm^2\n",
		c.Length(types.CM).String(), c.Width(types.CM).String(), c.Height(types.CM).String(), c.Area(types.CM).String())
	fmt.Printf("长: %s dm, 宽 %s dm, 高 %s dm, 面积 %s dm^2\n",
		c.Length(types.DM).String(), c.Width(types.DM).String(), c.Height(types.DM).String(), c.Area(types.DM).String())
	fmt.Printf("长: %s m, 宽 %s m, 高 %s m, 面积 %s m^2\n",
		c.Length(types.M).String(), c.Width(types.M).String(), c.Height(types.M).String(), c.Area(types.M).String())
	fmt.Println(c.DisplayWeight(types.CM, types.G))
}

func TestCuboids_Display(t *testing.T) {
	c := NewCuboids(decimal.NewFromInt(0), types.M, decimal.NewFromInt(0), types.CM, decimal.NewFromInt(0), types.DM, 0, Density{
		Value:      decimal.Decimal{},
		WeightUnit: 0,
		VolumeUnit: 0,
	})
	fmt.Println(c.Display(types.M))
}
