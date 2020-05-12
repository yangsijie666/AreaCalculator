package models

import (
    "AreaCalculator/types"
    "testing"
)

func TestPdfController_PrintWorkOrderInPdfFormat(t *testing.T) {
    p := NewPdfController("20.245", "13", "无备注", "", "/Users/yangsijie/Desktop/aaaaaaaaaaa.pdf", types.DM)
    p.PrintWorkOrderInPdfFormat()
}
