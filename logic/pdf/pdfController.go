package pdf

import (
	"AreaCalculator/types"
	"fmt"
)

type PdfController struct {
	totalArea   string
	totalNumber string
	remark      string
	timeStamp   string
	useSysTime  bool
	timeFormt   string
	filePath    string
}

func NewPdfController(totalArea, totalNumber, remark, timeStamp, timeFormat, filePath string, useSysTime bool, areaUnit types.Unit) *PdfController {
	p := new(PdfController)
	switch areaUnit {
	case types.MM:
		p.totalArea = totalArea + " 平方毫米"
	case types.CM:
		p.totalArea = totalArea + " 平方厘米"
	case types.DM:
		p.totalArea = totalArea + " 平方分米"
	case types.M:
		p.totalArea = totalArea + " 平方米"
	default:
		p.totalArea = totalArea + " 平方厘米"
	}
	p.totalNumber = totalNumber
	p.remark = remark
	p.timeStamp = timeStamp
	p.useSysTime = useSysTime
	p.timeFormt = timeFormat
	p.filePath = filePath
	return p
}

func (p *PdfController) PrintWorkOrderInPdfFormat(fontPath string) (err error) {
	defer func() {
		if recover() != nil {
			err = fmt.Errorf(types.UnknownError)
		}
	}()

	PrintArea = true
	TotalArea = p.totalArea
	TotalNumber = p.totalNumber
	Remark = p.remark
	TimeStamp = p.timeStamp
	TimeFormat = p.timeFormt
	UseSysTime = p.useSysTime
	PrintWorkOrderInPdfFormat(fontPath, p.filePath)

	return err
}

type PdfWeightController struct {
	totalWeight string
	totalNumber string
	remark      string
	timeStamp   string
	useSysTime  bool
	timeFormt   string
	filePath    string
}

func NewPdfWeightController(totalWeight, totalNumber, remark, timeStamp, timeFormat, filePath string, useSysTime bool,
	weightUnit types.WeightUnit) *PdfWeightController {
	p := new(PdfWeightController)
	switch weightUnit {
	case types.MG:
		p.totalWeight = totalWeight + " 毫克"
	case types.G:
		p.totalWeight = totalWeight + " 克"
	case types.KG:
		p.totalWeight = totalWeight + " 千克"
	case types.T:
		p.totalWeight = totalWeight + " 吨"
	default:
		p.totalWeight = totalWeight + " 克"
	}
	p.totalNumber = totalNumber
	p.remark = remark
	p.timeStamp = timeStamp
	p.useSysTime = useSysTime
	p.timeFormt = timeFormat
	p.filePath = filePath
	return p
}

func (p *PdfWeightController) WeightPrintWorkOrderInPdfFormat(fontPath string) (err error) {
	defer func() {
		if recover() != nil {
			err = fmt.Errorf(types.UnknownError)
		}
	}()

	PrintArea = false
	TotalWeight = p.totalWeight
	TotalNumber = p.totalNumber
	Remark = p.remark
	TimeStamp = p.timeStamp
	TimeFormat = p.timeFormt
	UseSysTime = p.useSysTime
	PrintWorkOrderInPdfFormat(fontPath, p.filePath)

	return err
}
