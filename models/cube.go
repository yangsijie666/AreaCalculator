package models

import (
	"AreaCalculator/types"
	"fmt"
	"github.com/shopspring/decimal"
)

type Cubes struct {
	lengths [4]decimal.Decimal
	number  int
	areas   [4]decimal.Decimal
	density Density
	weights [4]decimal.Decimal
}

func NewCubes(length decimal.Decimal, lengthUnit types.Unit, number int, density Density) *Cubes {
	cube := new(Cubes)

	cube.number = number
	cube.lengths[lengthUnit] = length

	tenDecimal := decimal.NewFromInt(10)
	hundredDecimal := decimal.NewFromInt(100)

	// 设定各单位的长
	for i := lengthUnit - 1; i >= types.MM; i-- {
		cube.lengths[i] = cube.lengths[i+1].Mul(tenDecimal)
	}
	for i := lengthUnit + 1; i <= types.M; i++ {
		cube.lengths[i] = cube.lengths[i-1].Div(tenDecimal)
	}

	cube.areas[types.MM] = cube.lengths[types.MM].Mul(cube.lengths[types.MM]).
		Mul(decimal.NewFromInt(int64(6 * number)))
	// 设定各单位的面积
	for i := types.CM; i <= types.M; i++ {
		cube.areas[i] = cube.areas[i-1].Div(hundredDecimal)
	}

	// 设定各单位重量
	if density.Value.GreaterThan(decimal.Zero) {
		cube.density = density
		thousandDecimal := decimal.NewFromInt(1000)

		cube.weights[density.WeightUnit] = cube.lengths[density.VolumeUnit].Pow(decimal.NewFromInt(3)).
			Mul(density.Value).Mul(decimal.NewFromInt(int64(number)))
		for i := density.WeightUnit - 1; i >= types.MG; i-- {
			cube.weights[i] = cube.weights[i+1].Mul(thousandDecimal)
		}
		for i := density.WeightUnit + 1; i <= types.T; i++ {
			cube.weights[i] = cube.weights[i-1].Div(thousandDecimal)
		}
	}

	return cube
}

func (c *Cubes) Length(unit types.Unit) decimal.Decimal {
	return c.lengths[unit]
}

func (c *Cubes) Area(unit types.Unit) decimal.Decimal {
	return c.areas[unit]
}

func (c *Cubes) Number() int {
	return c.number
}

func (c *Cubes) Display(unit types.Unit) string {
	return fmt.Sprintf(types.CubesDisplayStr, c.lengths[unit].String(), c.number, c.areas[unit].String())
}

// 重量相关输出
func (c *Cubes) Weight(unit types.WeightUnit) decimal.Decimal {
	if c.density.Value.GreaterThan(decimal.Zero) {
		return c.weights[unit]
	}
	panic(types.WeightInfoNonexistent)
}

func (c *Cubes) DisplayWeight(unit types.Unit, weightUnit types.WeightUnit) string {
	if c.density.Value.GreaterThan(decimal.Zero) {
		return fmt.Sprintf(types.CubesWeightDisplayStr, c.lengths[unit].String(), c.number, c.weights[weightUnit].String())
	}
	panic(types.WeightInfoNonexistent)
}
