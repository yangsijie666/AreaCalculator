package models

import (
    "AreaCalculator/logic/pdf"
    "AreaCalculator/types"
    "fmt"
)

type PdfController struct {
    totalArea       string
    totalNumber     string
    remark          string
    timeStamp       string
    useSysTime      bool
    timeFormt       string
    filePath        string
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
            err =  fmt.Errorf(types.UnknownError)
        }
    }()

    pdf.TotalArea = p.totalArea
    pdf.TotalNumber = p.totalNumber
    pdf.Remark = p.remark
    pdf.TimeStamp = p.timeStamp
    pdf.TimeFormat = p.timeFormt
    pdf.UseSysTime = p.useSysTime
    pdf.PrintWorkOrderInPdfFormat(fontPath, p.filePath)

    return err
}
