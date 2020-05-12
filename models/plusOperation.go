package models

import (
    "AreaCalculator/types"
    "fmt"
    "github.com/shopspring/decimal"
)

type PlusOperation struct {
    operateNumber   int
    operator        types.Operator
    cuboids         CuboidInterface
    curTotalNumber  int
    curSumArea      [4]decimal.Decimal
}

func NewPlusOperation(operateNumber int, cuboids CuboidInterface) *PlusOperation {
    plusOperation := new(PlusOperation)
    plusOperation.operateNumber = operateNumber
    plusOperation.operator = types.Plus
    plusOperation.cuboids = cuboids
    return plusOperation
}

func (o *PlusOperation) Operate(display *Display) {
    display.currentCuboids = o.cuboids
    display.totalNumber += o.cuboids.Number()
    for i := types.MM; i <= types.M; i++ {
        display.sumArea[i] = display.sumArea[i].Add(o.cuboids.Area(i))
        o.curSumArea[i] = display.sumArea[i]
    }
    o.curTotalNumber = display.totalNumber
}

func (o *PlusOperation) Operator() types.Operator {
    return o.operator
}

func (o *PlusOperation) Display(unit types.Unit) string {
    return fmt.Sprintf(types.OperationDisplayStr, o.operateNumber, "+",
        o.cuboids.Display(unit),
        o.curTotalNumber,
        o.curSumArea[unit].String())
}
