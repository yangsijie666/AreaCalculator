package pdf

import (
	"github.com/tiechui1994/gopdf"
	"github.com/tiechui1994/gopdf/core"
	"math"
	"time"
)

const (
	TableMy = "微软雅黑"
)

var (
	PrintArea   = true
	TotalWeight = ""
	TotalArea   = ""
	TotalNumber = ""
	Remark      = ""
	TimeStamp   = ""
	UseSysTime  bool
	TimeFormat  = "2006-01-02 15:04:05"
)

func PrintWorkOrderInPdfFormat(fontPath, filePath string) {
	r := core.CreateReport()
	font := core.FontMap{
		FontName: TableMy,
		FileName: fontPath,
	}
	r.SetFonts([]*core.FontMap{&font})
	r.SetPage("A4", "P")

	r.RegisterExecutor(printWorkOrderInPdfFormatExecutor, core.Detail)

	r.Execute(filePath)
}

func printWorkOrderInPdfFormatExecutor(report *core.Report) {
	lineSpace := 1.0
	lineHeight := 20.0

	remarkRows := int(math.Ceil(float64(len(Remark))/10)) + 2

	table := gopdf.NewTable(4, remarkRows+3, 415, lineHeight, report)
	table.SetMargin(core.Scope{})

	// 先把当前的行设置完毕, 然后才能添加单元格内容.
	cTitle := table.NewCellByRange(4, 1)
	var (
		cTotalAreaOrWeight *gopdf.TableCell
		cTotalNumber       *gopdf.TableCell
		cRemark            *gopdf.TableCell
		totalNumberRows    int
		totalAreaRows      int
	)
	if TotalNumber != "" {
		totalAreaRows = remarkRows - remarkRows/2
		totalNumberRows = remarkRows - totalAreaRows
		cTotalAreaOrWeight = table.NewCellByRange(2, totalAreaRows)
		cRemark = table.NewCellByRange(2, remarkRows)
		cTotalNumber = table.NewCellByRange(2, totalNumberRows)
	} else {
		cTotalAreaOrWeight = table.NewCellByRange(2, remarkRows)
		cRemark = table.NewCellByRange(2, remarkRows)
	}
	cOrderMaker := table.NewCellByRange(2, 1)
	cOrderSigner := table.NewCellByRange(2, 1)
	cDate := table.NewCellByRange(4, 1)

	fTitle := core.Font{Family: TableMy, Size: 17, Style: ""}
	fNormal := core.Font{Family: TableMy, Size: 13, Style: ""}

	border := core.NewScope(4.0, 4.0, 4.0, 3.0)

	cTitle.SetElement(gopdf.NewTextCell(table.GetColWidth(0, 0), 25, lineSpace, report).SetFont(fTitle).SetBorder(border).HorizontalCentered().VerticalCentered().SetContent("工作量单"))
	if cTotalNumber != nil {
		if PrintArea {
			cTotalAreaOrWeight.SetElement(gopdf.NewTextCell(table.GetColWidth(1, 0), 15, lineSpace, report).SetFont(fNormal).SetBorder(border).SetContent("总面积累积和：" + TotalArea))
		} else {
			cTotalAreaOrWeight.SetElement(gopdf.NewTextCell(table.GetColWidth(1, 0), 15, lineSpace, report).SetFont(fNormal).SetBorder(border).SetContent("总重量累积和：" + TotalWeight))
		}
		cTotalNumber.SetElement(gopdf.NewTextCell(table.GetColWidth(totalAreaRows+1, 0), 15, lineSpace, report).SetFont(fNormal).SetBorder(border).SetContent("物体总数量：" + TotalNumber))
	} else {
		if PrintArea {
			cTotalAreaOrWeight.SetElement(gopdf.NewTextCell(table.GetColWidth(1, 0), 15, lineSpace, report).SetFont(fNormal).SetBorder(border).SetContent("总面积累积和：" + TotalArea))
		} else {
			cTotalAreaOrWeight.SetElement(gopdf.NewTextCell(table.GetColWidth(1, 0), 15, lineSpace, report).SetFont(fNormal).SetBorder(border).SetContent("总重量累积和：" + TotalWeight))
		}
	}
	cOrderMaker.SetElement(gopdf.NewTextCell(table.GetColWidth(remarkRows+1, 0), 15, lineSpace, report).SetFont(fNormal).SetBorder(border).VerticalCentered().SetContent("制单人："))
	cOrderSigner.SetElement(gopdf.NewTextCell(table.GetColWidth(remarkRows+1, 2), 15, lineSpace, report).SetFont(fNormal).SetBorder(border).VerticalCentered().SetContent("确认人："))
	if !UseSysTime {
		cDate.SetElement(gopdf.NewTextCell(table.GetColWidth(remarkRows+2, 0), 15, lineSpace, report).SetFont(fNormal).SetBorder(border).VerticalCentered().SetContent("日期：" + TimeStamp))
	} else {
		cDate.SetElement(gopdf.NewTextCell(table.GetColWidth(remarkRows+2, 0), 15, lineSpace, report).SetFont(fNormal).SetBorder(border).VerticalCentered().SetContent("日期：" + time.Now().Format(TimeFormat)))
	}
	cRemark.SetElement(gopdf.NewTextCell(table.GetColWidth(1, 2), 15, lineSpace, report).SetFont(fNormal).SetBorder(border).SetContent("备注：" + Remark))

	table.GenerateAtomicCell()
}
