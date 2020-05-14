package models

import (
	"AreaCalculator/types"
	"fmt"
	"github.com/shopspring/decimal"
)

type PlusOperation struct {
	operateNumber  int
	operator       types.Operator
	cuboids        CuboidInterface
	curTotalNumber int
	curSumArea     [4]decimal.Decimal
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

type PlusWeightOperation struct {
	operateNumber  int
	operator       types.Operator
	cuboids        CuboidWeightInterface
	curTotalNumber int
	curSumWeight   [4]decimal.Decimal
}

func NewPlusWeightOperation(operateNumber int, cuboids CuboidWeightInterface) *PlusWeightOperation {
	plusWeightOperation := new(PlusWeightOperation)
	plusWeightOperation.operateNumber = operateNumber
	plusWeightOperation.operator = types.Plus
	plusWeightOperation.cuboids = cuboids
	return plusWeightOperation
}

func (o *PlusWeightOperation) Operate(displayer *WeightDisplayer) {
	displayer.currentCuboids = o.cuboids
	displayer.totalNumber += o.cuboids.Number()
	for i := types.MG; i <= types.T; i++ {
		displayer.sumWeight[i] = displayer.sumWeight[i].Add(o.cuboids.Weight(i))
		o.curSumWeight[i] = displayer.sumWeight[i]
	}
	o.curTotalNumber = displayer.totalNumber
}

func (o *PlusWeightOperation) Operator() types.Operator {
	return o.operator
}

func (o *PlusWeightOperation) DisplayWeight(unit types.Unit, weightUnit types.WeightUnit) string {
	return fmt.Sprintf(types.OperationWeightDisplayStr, o.operateNumber, "+",
		o.cuboids.DisplayWeight(unit, weightUnit),
		o.curTotalNumber,
		o.curSumWeight[weightUnit].String())
}
