package models

import (
	"AreaCalculator/types"
	"fmt"
	"github.com/shopspring/decimal"
)

type MinusOperation struct {
	operateNumber  int
	operator       types.Operator
	cuboids        CuboidInterface
	curTotalNumber int
	curSumArea     [4]decimal.Decimal
}

func NewMinusOperation(operateNumber int, cuboids CuboidInterface) *MinusOperation {
	minusOperation := new(MinusOperation)
	minusOperation.operateNumber = operateNumber
	minusOperation.operator = types.Minus
	minusOperation.cuboids = cuboids
	return minusOperation
}

func (o *MinusOperation) Operate(display *Display) {
	display.currentCuboids = o.cuboids
	display.totalNumber -= o.cuboids.Number()
	for i := types.MM; i <= types.M; i++ {
		display.sumArea[i] = display.sumArea[i].Sub(o.cuboids.Area(i))
		o.curSumArea[i] = display.sumArea[i]
	}
	o.curTotalNumber = display.totalNumber
}

func (o *MinusOperation) Operator() types.Operator {
	return o.operator
}

func (o *MinusOperation) Display(unit types.Unit) string {
	return fmt.Sprintf(types.OperationDisplayStr, o.operateNumber, "-",
		o.cuboids.Display(unit),
		o.curTotalNumber,
		o.curSumArea[unit].String())
}

type MinusWeightOperation struct {
	operateNumber  int
	operator       types.Operator
	cuboids        CuboidWeightInterface
	curTotalNumber int
	curSumWeight   [4]decimal.Decimal
}

func NewMinusWeightOperation(operateNumber int, cuboids CuboidWeightInterface) *MinusWeightOperation {
	minusWeightOperation := new(MinusWeightOperation)
	minusWeightOperation.operateNumber = operateNumber
	minusWeightOperation.operator = types.Minus
	minusWeightOperation.cuboids = cuboids
	return minusWeightOperation
}

func (o *MinusWeightOperation) Operate(displayer *WeightDisplayer) {
	displayer.currentCuboids = o.cuboids
	displayer.totalNumber -= o.cuboids.Number()
	for i := types.MG; i <= types.T; i++ {
		displayer.sumWeight[i] = displayer.sumWeight[i].Sub(o.cuboids.Weight(i))
		o.curSumWeight[i] = displayer.sumWeight[i]
	}
	o.curTotalNumber = displayer.totalNumber
}

func (o *MinusWeightOperation) Operator() types.Operator {
	return o.operator
}

func (o *MinusWeightOperation) DisplayWeight(unit types.Unit, weightUnit types.WeightUnit) string {
	return fmt.Sprintf(types.OperationWeightDisplayStr, o.operateNumber, "-",
		o.cuboids.DisplayWeight(unit, weightUnit),
		o.curTotalNumber,
		o.curSumWeight[weightUnit].String())
}
