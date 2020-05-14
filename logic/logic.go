package logic

import (
	"AreaCalculator/initial"
	"AreaCalculator/models"
	"AreaCalculator/tools"
	"AreaCalculator/types"
	"fmt"
	"github.com/shopspring/decimal"
)

// ---------------------------------------------面积相关计算（开始）---------------------------------------------------
// 加减长方体
func PlusCuboids(lengthStr, widthStr, heightStr string, lengthUnit, widthUnit, heightUnit int32, numberStr string) error {
	length, width, height, number, err := validateCuboidsInput(lengthStr, widthStr, heightStr, numberStr)
	if err != nil {
		return err
	}

	cuboids := models.NewCuboids(length, types.Unit(lengthUnit), width, types.Unit(widthUnit),
		height, types.Unit(heightUnit), number, models.Density{
			Value:      decimal.Zero,
			WeightUnit: 0,
			VolumeUnit: 0,
		})
	plusOperation := models.NewPlusOperation(initial.CurOperateNumber, cuboids)
	initial.CurOperateNumber += 1
	plusOperation.Operate(initial.MyDisplay)
	initial.MyHistoryDisplay.AppendHistory(models.NewHistory(plusOperation))

	return nil
}

func MinusCuboids(lengthStr, widthStr, heightStr string, lengthUnit, widthUnit, heightUnit int32, numberStr string) error {
	length, width, height, number, err := validateCuboidsInput(lengthStr, widthStr, heightStr, numberStr)
	if err != nil {
		return err
	}

	cuboids := models.NewCuboids(length, types.Unit(lengthUnit), width, types.Unit(widthUnit),
		height, types.Unit(heightUnit), number, models.Density{
			Value:      decimal.Zero,
			WeightUnit: 0,
			VolumeUnit: 0,
		})
	minusOperation := models.NewMinusOperation(initial.CurOperateNumber, cuboids)
	initial.CurOperateNumber += 1
	minusOperation.Operate(initial.MyDisplay)
	initial.MyHistoryDisplay.AppendHistory(models.NewHistory(minusOperation))

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

	cubes := models.NewCubes(length, types.Unit(unit), number, models.Density{
		Value:      decimal.Zero,
		WeightUnit: 0,
		VolumeUnit: 0,
	})
	plusOperation := models.NewPlusOperation(initial.CurOperateNumber, cubes)
	initial.CurOperateNumber += 1
	plusOperation.Operate(initial.MyDisplay)
	initial.MyHistoryDisplay.AppendHistory(models.NewHistory(plusOperation))

	return nil
}

func MinusCubes(lengthStr string, unit int32, numberStr string) error {
	length, number, err := validateCubesInput(lengthStr, numberStr)
	if err != nil {
		return err
	}

	cubes := models.NewCubes(length, types.Unit(unit), number, models.Density{
		Value:      decimal.Zero,
		WeightUnit: 0,
		VolumeUnit: 0,
	})
	minusOperation := models.NewMinusOperation(initial.CurOperateNumber, cubes)
	initial.CurOperateNumber += 1
	minusOperation.Operate(initial.MyDisplay)
	initial.MyHistoryDisplay.AppendHistory(models.NewHistory(minusOperation))

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

// ---------------------------------------------面积相关计算（结束）---------------------------------------------------

// ---------------------------------------------体积相关计算（开始）---------------------------------------------------
// 加减长方体
func PlusCuboidsWeight(lengthStr, widthStr, heightStr string, lengthUnit, widthUnit, heightUnit int32,
	numberStr, densityStr string, weightUnit, volumeUnit int32) error {
	length, width, height, number, err := validateCuboidsInput(lengthStr, widthStr, heightStr, numberStr)
	if err != nil {
		return err
	}
	density, err := validateWeightInput(densityStr, weightUnit, volumeUnit)
	if err != nil {
		return err
	}

	cuboids := models.NewCuboids(length, types.Unit(lengthUnit), width, types.Unit(widthUnit),
		height, types.Unit(heightUnit), number, density)
	plusOperation := models.NewPlusWeightOperation(initial.CurWeightOperateNumber, cuboids)
	initial.CurWeightOperateNumber += 1
	plusOperation.Operate(initial.MyWeightDisplayer)
	initial.MyWeightHistoryDisplayer.AppendWeightHistory(models.NewWeightHistory(plusOperation))

	return nil
}

func MinusCuboidsWeight(lengthStr, widthStr, heightStr string, lengthUnit, widthUnit, heightUnit int32,
	numberStr, densityStr string, weightUnit, volumeUnit int32) error {
	length, width, height, number, err := validateCuboidsInput(lengthStr, widthStr, heightStr, numberStr)
	if err != nil {
		return err
	}
	density, err := validateWeightInput(densityStr, weightUnit, volumeUnit)
	if err != nil {
		return err
	}

	cuboids := models.NewCuboids(length, types.Unit(lengthUnit), width, types.Unit(widthUnit),
		height, types.Unit(heightUnit), number, density)
	minusOperation := models.NewMinusWeightOperation(initial.CurWeightOperateNumber, cuboids)
	initial.CurWeightOperateNumber += 1
	minusOperation.Operate(initial.MyWeightDisplayer)
	initial.MyWeightHistoryDisplayer.AppendWeightHistory(models.NewWeightHistory(minusOperation))

	return nil
}

// 加减正方体
func PlusCubesWeight(lengthStr string, unit int32, numberStr, densityStr string, weightUnit, volumeUnit int32) error {
	length, number, err := validateCubesInput(lengthStr, numberStr)
	if err != nil {
		return err
	}
	density, err := validateWeightInput(densityStr, weightUnit, volumeUnit)
	if err != nil {
		return err
	}

	cubes := models.NewCubes(length, types.Unit(unit), number, density)
	plusOperation := models.NewPlusWeightOperation(initial.CurWeightOperateNumber, cubes)
	initial.CurWeightOperateNumber += 1
	plusOperation.Operate(initial.MyWeightDisplayer)
	initial.MyWeightHistoryDisplayer.AppendWeightHistory(models.NewWeightHistory(plusOperation))

	return nil
}

func MinusCubesWeight(lengthStr string, unit int32, numberStr, densityStr string, weightUnit, volumeUnit int32) error {
	length, number, err := validateCubesInput(lengthStr, numberStr)
	if err != nil {
		return err
	}
	density, err := validateWeightInput(densityStr, weightUnit, volumeUnit)
	if err != nil {
		return err
	}

	cubes := models.NewCubes(length, types.Unit(unit), number, density)
	minusOperation := models.NewMinusWeightOperation(initial.CurWeightOperateNumber, cubes)
	initial.CurWeightOperateNumber += 1
	minusOperation.Operate(initial.MyWeightDisplayer)
	initial.MyWeightHistoryDisplayer.AppendWeightHistory(models.NewWeightHistory(minusOperation))

	return nil
}

func validateWeightInput(densityStr string, weightUnit, volumeUnit int32) (density models.Density, errMsg error) {
	value, err := decimal.NewFromString(densityStr)
	if err != nil {
		errMsg = fmt.Errorf(types.ErrMsgTemplate, "密度", densityStr)
		return
	}

	if value.LessThanOrEqual(decimal.Zero) {
		errMsg = fmt.Errorf(types.ErrMsgTemplate, "密度", densityStr)
		return
	}

	density.Value = value
	density.WeightUnit = types.WeightUnit(weightUnit)
	density.VolumeUnit = types.Unit(volumeUnit)
	return
}

// ---------------------------------------------面积相关计算（结束）---------------------------------------------------
