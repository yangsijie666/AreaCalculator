package logic

import (
    "AreaCalculator/models"
    "AreaCalculator/tools"
    "AreaCalculator/types"
    "fmt"
    "github.com/shopspring/decimal"
)

var (
    MyDisplay        = models.NewDisplay()
    MyHistoryDisplay = models.NewHistoryDisplay()
    curOperateNumber = 1
)

// 加减长方体
func PlusCuboids(lengthStr, widthStr, heightStr string, lengthUnit, widthUnit, heightUnit int32, numberStr string) error {
    length, width, height, number, err := validateCuboidsInput(lengthStr, widthStr, heightStr, numberStr)
    if err != nil {
        return err
    }

    cuboids := models.NewCuboids(length, types.Unit(lengthUnit), width, types.Unit(widthUnit), height, types.Unit(heightUnit), number)
    plusOperation := models.NewPlusOperation(curOperateNumber, cuboids)
    curOperateNumber += 1
    plusOperation.Operate(MyDisplay)
    MyHistoryDisplay.AppendHistory(models.NewHistory(plusOperation))

    return nil
}

func MinusCuboids(lengthStr, widthStr, heightStr string, lengthUnit, widthUnit, heightUnit int32, numberStr string) error {
    length, width, height, number, err := validateCuboidsInput(lengthStr, widthStr, heightStr, numberStr)
    if err != nil {
        return err
    }

    cuboids := models.NewCuboids(length, types.Unit(lengthUnit), width, types.Unit(widthUnit), height, types.Unit(heightUnit), number)
    minusOperation := models.NewMinusOperation(curOperateNumber, cuboids)
    curOperateNumber += 1
    minusOperation.Operate(MyDisplay)
    MyHistoryDisplay.AppendHistory(models.NewHistory(minusOperation))

    return nil
}

func validateCuboidsInput(length, width, height, number string) (
    lengthFloat, widthFloat, heightFloat decimal.Decimal, numberInt int, errMsg error) {
     lengthFloat, err := decimal.NewFromString(length)
     if err != nil {
         errMsg = fmt.Errorf(types.ErrMsgTemplate, "长", length)
         return
     }
     if lengthFloat.LessThanOrEqual(decimal.NewFromInt(0)) {
         errMsg = fmt.Errorf(types.ErrMsgTemplate, "长", length)
         return
     }

    widthFloat, err = decimal.NewFromString(width)
    if err != nil {
        errMsg = fmt.Errorf(types.ErrMsgTemplate, "宽", width)
        return
    }
    if widthFloat.LessThanOrEqual(decimal.NewFromInt(0)) {
        errMsg = fmt.Errorf(types.ErrMsgTemplate, "宽", width)
        return
    }

    heightFloat, err = decimal.NewFromString(height)
    if err != nil {
        errMsg = fmt.Errorf(types.ErrMsgTemplate, "高", height)
        return
    }
    if heightFloat.LessThanOrEqual(decimal.NewFromInt(0)) {
        errMsg = fmt.Errorf(types.ErrMsgTemplate, "高", height)
        return
    }

    numberInt, err = tools.ValidateIntStr(number)
    if err != nil {
        errMsg = fmt.Errorf(types.NumberErrMsgTemplate, number)
        return
    }
    if numberInt < 1 {
        errMsg = fmt.Errorf(types.NumberErrMsgTemplate, number)
        return
    }

    return
}

// 加减正方体
func PlusCubes(lengthStr string, unit int32, numberStr string) error {
    length, number, err := validateCubesInput(lengthStr, numberStr)
    if err != nil {
        return err
    }

    cubes := models.NewCubes(length, types.Unit(unit), number)
    plusOperation := models.NewPlusOperation(curOperateNumber, cubes)
    curOperateNumber += 1
    plusOperation.Operate(MyDisplay)
    MyHistoryDisplay.AppendHistory(models.NewHistory(plusOperation))

    return nil
}

func MinusCubes(lengthStr string, unit int32, numberStr string) error {
    length, number, err := validateCubesInput(lengthStr, numberStr)
    if err != nil {
        return err
    }

    cubes := models.NewCubes(length, types.Unit(unit), number)
    minusOperation := models.NewMinusOperation(curOperateNumber, cubes)
    curOperateNumber += 1
    minusOperation.Operate(MyDisplay)
    MyHistoryDisplay.AppendHistory(models.NewHistory(minusOperation))

    return nil
}

func validateCubesInput(length, number string) (lengthFloat decimal.Decimal, numberInt int, errMsg error) {
    lengthFloat, err := decimal.NewFromString(length)
    if err != nil {
        errMsg = fmt.Errorf(types.ErrMsgTemplate, "长", length)
        return
    }
    if lengthFloat.LessThanOrEqual(decimal.NewFromInt(0)) {
        errMsg = fmt.Errorf(types.ErrMsgTemplate, "长", length)
        return
    }

    numberInt, err = tools.ValidateIntStr(number)
    if err != nil {
        errMsg = fmt.Errorf(types.NumberErrMsgTemplate, number)
        return
    }
    if numberInt < 1 {
        errMsg = fmt.Errorf(types.NumberErrMsgTemplate, number)
        return
    }

    return
}
