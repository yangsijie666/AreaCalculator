package models

import (
	"AreaCalculator/logic/pdf"
	"AreaCalculator/types"
	"testing"
)

func TestPdfController_PrintWorkOrderInPdfFormat(t *testing.T) {
	p := pdf.NewPdfController("20.245", "13", "无备注", "", "", "/Users/yangsijie/Desktop/aaaaaaaaaaa.pdf", true, types.DM)
	p.PrintWorkOrderInPdfFormat("./simhei.ttf")
}
