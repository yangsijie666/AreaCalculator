package models

import (
    "AreaCalculator/types"
    "fmt"
    "github.com/shopspring/decimal"
)

type Cuboids struct {
    lengths     [4]decimal.Decimal
    widths      [4]decimal.Decimal
    heights     [4]decimal.Decimal
    number      int
    areas       [4]decimal.Decimal
}

func NewCuboids(length decimal.Decimal, lengthUnit types.Unit, width decimal.Decimal, widthUnit types.Unit,
    height decimal.Decimal, heightUnit types.Unit, number int) *Cuboids {
    cuboid := new(Cuboids)

    cuboid.number = number

    cuboid.lengths[lengthUnit] = length
    cuboid.widths[widthUnit] = width
    cuboid.heights[heightUnit] = height

    tenDecimal := decimal.NewFromInt(10)
    hundredDecimal := decimal.NewFromInt(100)

    // 设定各单位的长
    for i := lengthUnit - 1; i >= types.MM; i-- {
        cuboid.lengths[i] = cuboid.lengths[i + 1].Mul(tenDecimal)
    }
    for i := lengthUnit + 1; i <= types.M; i++ {
        cuboid.lengths[i] = cuboid.lengths[i - 1].Div(tenDecimal)
    }

    // 设定各单位的宽
    for i := widthUnit - 1; i >= types.MM; i-- {
        cuboid.widths[i] = cuboid.widths[i + 1].Mul(tenDecimal)
    }
    for i := widthUnit + 1; i <= types.M; i++ {
        cuboid.widths[i] = cuboid.widths[i - 1].Div(tenDecimal)
    }

    // 设定各单位的高
    for i := heightUnit - 1; i >= types.MM; i-- {
        cuboid.heights[i] = cuboid.heights[i + 1].Mul(tenDecimal)
    }
    for i := heightUnit + 1; i <= types.M; i++ {
        cuboid.heights[i] = cuboid.heights[i - 1].Div(tenDecimal)
    }

    cuboid.areas[types.MM] = ((cuboid.lengths[types.MM].Mul(cuboid.widths[types.MM])).
        Add(cuboid.widths[types.MM].Mul(cuboid.heights[types.MM])).
        Add(cuboid.lengths[types.MM].Mul(cuboid.heights[types.MM]))).
        Mul(decimal.NewFromInt(int64(2 * number)))
    // 设定各单位的面积
    for i := types.CM; i <= types.M; i++ {
        cuboid.areas[i] = cuboid.areas[i - 1].Div(hundredDecimal)
    }

    return cuboid
}

func (c *Cuboids) Length(unit types.Unit) decimal.Decimal {
    return c.lengths[unit]
}

func (c *Cuboids) Width(unit types.Unit) decimal.Decimal {
    return c.widths[unit]
}

func (c *Cuboids) Height(unit types.Unit) decimal.Decimal {
    return c.heights[unit]
}

func (c *Cuboids) Area(unit types.Unit) decimal.Decimal {
    return c.areas[unit]
}

func (c *Cuboids) Number() int {
    return c.number
}

func (c *Cuboids) Display(unit types.Unit) string {
    return fmt.Sprintf(types.CuboidsDisplayStr, c.lengths[unit].String(), c.widths[unit].String(),
        c.heights[unit].String(), c.number, c.areas[unit].String())
}
