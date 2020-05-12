package models

import (
    "AreaCalculator/types"
    "fmt"
    "github.com/shopspring/decimal"
    "testing"
)

func TestCubes_Display(t *testing.T) {
    cubes := NewCubes(decimal.NewFromInt(2),  types.CM, 1)
    fmt.Println(cubes.Display(types.MM))
    fmt.Println(cubes.Display(types.CM))
    fmt.Println(cubes.Display(types.DM))
    fmt.Println(cubes.Display(types.M))
}
