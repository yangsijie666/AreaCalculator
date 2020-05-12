package models

import (
    "AreaCalculator/types"
    "fmt"
    "github.com/shopspring/decimal"
)

type Cubes struct {
    lengths     [4]decimal.Decimal
    number      int
    areas       [4]decimal.Decimal
}

func NewCubes(length decimal.Decimal, lengthUnit types.Unit, number int) *Cubes {
    cube := new(Cubes)

    cube.number = number
    cube.lengths[lengthUnit] = length

    tenDecimal := decimal.NewFromInt(10)
    hundredDecimal := decimal.NewFromInt(100)

    // 设定各单位的长
    for i := lengthUnit - 1; i >= types.MM; i-- {
        cube.lengths[i] = cube.lengths[i + 1].Mul(tenDecimal)
    }
    for i := lengthUnit + 1; i <= types.M; i++ {
        cube.lengths[i] = cube.lengths[i - 1].Div(tenDecimal)
    }

    cube.areas[types.MM] = cube.lengths[types.MM].Mul(cube.lengths[types.MM]).
        Mul(decimal.NewFromInt(int64(6 * number)))
    // 设定各单位的面积
    for i := types.CM; i <= types.M; i++ {
        cube.areas[i] = cube.areas[i - 1].Div(hundredDecimal)
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
