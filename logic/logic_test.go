package logic

import (
    "AreaCalculator/types"
    "fmt"
    "log"
    "testing"
)

func TestPlus(t *testing.T) {
    err := PlusCubes("2", int32(types.CM), "1")
    if err != nil {
        log.Fatal(err)
    }

    err = PlusCuboids("2", "1", "3", int32(types.M), int32(types.DM), int32(types.CM), "3")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(MyDisplay.DisplayResult(types.M))
    fmt.Println()
    fmt.Println()
    fmt.Println(MyHistoryDisplay.Display(types.M))
}

func TestMinus(t *testing.T)  {
    err := PlusCuboids("2", "1", "3", int32(types.M), int32(types.CM), int32(types.DM), "1")
    if err != nil {
        log.Fatal(err)
    }

    err = MinusCubes("2", int32(types.CM), "2")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(MyDisplay.DisplayResult(types.M))
    fmt.Println()
    fmt.Println()
    fmt.Println(MyHistoryDisplay.Display(types.M))
}
