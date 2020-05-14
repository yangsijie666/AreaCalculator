package main

import (
	"AreaCalculator/initial"
	"AreaCalculator/logic/pdf"
	"AreaCalculator/tools"
	mytypes "AreaCalculator/types"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"strings"
)

func (f *TExportPdfForm) OnFormCreate(sender vcl.IObject) {

	// ================== 窗口组件 ==================
	f.SetCaption("导出PDF")
	f.SetHeight(217)
	f.SetWidth(345)
	f.SetLeft(8)
	f.SetTop(8)
	f.SetFont(newFontBlack(-15, 11))
	f.WorkAreaCenter()
	f.EnabledMaximize(false)
	// 约束窗口组件大小
	f.SetOnConstrainedResize(func(sender vcl.IObject, minWidth, minHeight, maxWidth, maxHeight *int32) {
		*minWidth = f.Width()
		*maxWidth = f.Width()
		*minHeight = f.Height()
		*maxHeight = f.Height()
	})

	if MainForm.Caption() == mytypes.AreaCalculator {
		// ================== 总面积累积和 标签组件 ==================
		f.TotalAreaLabel = vcl.NewLabel(f)
		f.TotalAreaLabel.SetParent(f)
		f.TotalAreaLabel.SetCaption("总面积累积和")
		f.TotalAreaLabel.SetFont(newFontBlack(-15, 11))
		f.TotalAreaLabel.SetHeight(18)
		f.TotalAreaLabel.SetWidth(92)
		f.TotalAreaLabel.SetLeft(7)
		f.TotalAreaLabel.SetTop(11)
		f.TotalAreaLabel.SetParentColor(false)

		// ================== 总面积累积和 编辑框组件 ==================
		f.TotalAreaEdit = vcl.NewEdit(f)
		f.TotalAreaEdit.SetParent(f)
		f.TotalAreaEdit.SetFont(newFontBlack(-13, 10))
		f.TotalAreaEdit.SetHeight(22)
		f.TotalAreaEdit.SetWidth(71)
		f.TotalAreaEdit.SetLeft(101)
		f.TotalAreaEdit.SetTop(7)
		f.TotalAreaEdit.SetParentColor(false)
		f.TotalAreaEdit.SetTabOrder(0)
		f.TotalAreaEdit.SetText(initial.MyDisplay.SumArea(mytypes.DisplayUnitDefault))
	} else {
		// ================== 总重量累积和 标签组件 ==================
		f.TotalWeightLabel = vcl.NewLabel(f)
		f.TotalWeightLabel.SetParent(f)
		f.TotalWeightLabel.SetCaption("总重量累积和")
		f.TotalWeightLabel.SetFont(newFontBlack(-15, 11))
		f.TotalWeightLabel.SetHeight(18)
		f.TotalWeightLabel.SetWidth(92)
		f.TotalWeightLabel.SetLeft(7)
		f.TotalWeightLabel.SetTop(11)
		f.TotalWeightLabel.SetParentColor(false)

		// ================== 总重量累积和 编辑框组件 ==================
		f.TotalWeightEdit = vcl.NewEdit(f)
		f.TotalWeightEdit.SetParent(f)
		f.TotalWeightEdit.SetFont(newFontBlack(-13, 10))
		f.TotalWeightEdit.SetHeight(22)
		f.TotalWeightEdit.SetWidth(71)
		f.TotalWeightEdit.SetLeft(101)
		f.TotalWeightEdit.SetTop(7)
		f.TotalWeightEdit.SetParentColor(false)
		f.TotalWeightEdit.SetTabOrder(0)
		f.TotalWeightEdit.SetText(initial.MyWeightDisplayer.SumWeight(mytypes.DisplayWeightUnitDefault))
	}

	// ================== 显示单位 标签组件 ==================
	f.DisplayUnit = vcl.NewLabel(f)
	f.DisplayUnit.SetParent(f)
	f.DisplayUnit.SetCaption("显示单位")
	f.DisplayUnit.SetFont(newFontBlack(-15, 11))
	f.DisplayUnit.SetHeight(18)
	f.DisplayUnit.SetWidth(61)
	f.DisplayUnit.SetLeft(184)
	f.DisplayUnit.SetTop(11)
	f.DisplayUnit.SetParentColor(false)

	// ================== 显示单位 下拉框组件 ==================
	f.DisplayUnitComboBox = vcl.NewComboBox(f)
	f.DisplayUnitComboBox.SetParent(f)
	f.DisplayUnitComboBox.SetFont(newFontBlack(-13, 10))
	f.DisplayUnitComboBox.SetHeight(20)
	f.DisplayUnitComboBox.SetWidth(87)
	f.DisplayUnitComboBox.SetLeft(250)
	f.DisplayUnitComboBox.SetTop(8)
	f.DisplayUnitComboBox.SetItemHeight(26)
	f.DisplayUnitComboBox.SetStyle(types.CsDropDownList)
	if MainForm.Caption() == mytypes.AreaCalculator {
		f.DisplayUnitComboBox.Items().Add("平方毫米")
		f.DisplayUnitComboBox.Items().Add("平方厘米")
		f.DisplayUnitComboBox.Items().Add("平方分米")
		f.DisplayUnitComboBox.Items().Add("平方米")
		f.DisplayUnitComboBox.SetItemIndex(int32(mytypes.DisplayUnitDefault))
	} else {
		f.DisplayUnitComboBox.Items().Add("毫克")
		f.DisplayUnitComboBox.Items().Add("克")
		f.DisplayUnitComboBox.Items().Add("千克")
		f.DisplayUnitComboBox.Items().Add("吨")
		f.DisplayUnitComboBox.SetItemIndex(int32(mytypes.DisplayWeightUnitDefault))
	}
	f.DisplayUnitComboBox.SetParentColor(false)
	f.DisplayUnitComboBox.SetTabOrder(1)

	// ================== 物体总数 标签组件 ==================
	f.TotalNumberLabel = vcl.NewLabel(f)
	f.TotalNumberLabel.SetParent(f)
	f.TotalNumberLabel.SetCaption("物体总数")
	f.TotalNumberLabel.SetFont(newFontBlack(-15, 11))
	f.TotalNumberLabel.SetHeight(18)
	f.TotalNumberLabel.SetWidth(61)
	f.TotalNumberLabel.SetLeft(38)
	f.TotalNumberLabel.SetTop(39)
	f.TotalNumberLabel.SetParentColor(false)

	// ================== 物体总数 编辑框组件 ==================
	f.TotalNumberEdit = vcl.NewEdit(f)
	f.TotalNumberEdit.SetParent(f)
	f.TotalNumberEdit.SetEnabled(false)
	f.TotalNumberEdit.SetFont(newFontBlack(-13, 10))
	f.TotalNumberEdit.SetHeight(22)
	f.TotalNumberEdit.SetWidth(71)
	f.TotalNumberEdit.SetLeft(101)
	f.TotalNumberEdit.SetTop(35)
	f.TotalNumberEdit.SetParentColor(false)
	f.TotalNumberEdit.SetTabOrder(2)
	if MainForm.Caption() == mytypes.AreaCalculator {
		f.TotalNumberEdit.SetText(initial.MyDisplay.TotalNumber())
	} else {
		f.TotalNumberEdit.SetText(initial.MyWeightDisplayer.TotalNumber())
	}

	// ================== 导出物体总数 复选框组件 ==================
	f.TotalNumberCheckBox = vcl.NewCheckBox(f)
	f.TotalNumberCheckBox.SetParent(f)
	f.TotalNumberCheckBox.SetCaption("导出物体总数")
	f.TotalNumberCheckBox.SetFont(newFontBlack(-15, 11))
	f.TotalNumberCheckBox.SetHeight(19)
	f.TotalNumberCheckBox.SetWidth(114)
	f.TotalNumberCheckBox.SetLeft(184)
	f.TotalNumberCheckBox.SetTop(37)
	f.TotalNumberCheckBox.SetParentFont(false)
	f.TotalNumberCheckBox.SetTabOrder(3)

	// ================== 日期格式 标签组件 ==================
	f.DateFormat = vcl.NewLabel(f)
	f.DateFormat.SetParent(f)
	f.DateFormat.SetCaption("日期格式")
	f.DateFormat.SetFont(newFontBlack(-15, 11))
	f.DateFormat.SetHeight(18)
	f.DateFormat.SetWidth(61)
	f.DateFormat.SetLeft(38)
	f.DateFormat.SetTop(68)
	f.DateFormat.SetParentColor(false)

	// ================== 日期格式 下拉框组件 ==================
	f.DateFormatComboBox = vcl.NewComboBox(f)
	f.DateFormatComboBox.SetParent(f)
	f.DateFormatComboBox.SetFont(newFontBlack(-13, 10))
	f.DateFormatComboBox.SetHeight(20)
	f.DateFormatComboBox.SetWidth(187)
	f.DateFormatComboBox.SetLeft(119)
	f.DateFormatComboBox.SetTop(65)
	f.DateFormatComboBox.SetStyle(types.CsDropDownList)
	f.DateFormatComboBox.SetItemHeight(26)
	f.DateFormatComboBox.Items().Add("yyyy-MM-dd HH:mm:ss")
	f.DateFormatComboBox.Items().Add("yyyy-MM-dd HH:mm")
	f.DateFormatComboBox.Items().Add("yyyy-MM-dd")
	f.DateFormatComboBox.SetItemIndex(0)
	f.DateFormatComboBox.SetParentFont(false)
	f.DateFormatComboBox.SetTabOrder(4)

	// ================== 自定义日期 复选框组件 ==================
	f.CustomDate = vcl.NewCheckBox(f)
	f.CustomDate.SetParent(f)
	f.CustomDate.SetCaption("自定义日期")
	f.CustomDate.SetFont(newFontBlack(-15, 11))
	f.CustomDate.SetHeight(19)
	f.CustomDate.SetWidth(99)
	f.CustomDate.SetLeft(4)
	f.CustomDate.SetTop(95)
	f.CustomDate.SetParentFont(false)
	f.CustomDate.SetTabOrder(5)

	// ================== 年 标签组件 ==================
	f.YearLabel = vcl.NewLabel(f)
	f.YearLabel.SetParent(f)
	f.YearLabel.SetCaption("年")
	f.YearLabel.SetFont(newFontBlack(-15, 11))
	f.YearLabel.SetHeight(18)
	f.YearLabel.SetWidth(15)
	f.YearLabel.SetLeft(162)
	f.YearLabel.SetTop(97)
	f.YearLabel.SetParentColor(false)

	// ================== 月 标签组件 ==================
	f.MonthLabel = vcl.NewLabel(f)
	f.MonthLabel.SetParent(f)
	f.MonthLabel.SetCaption("月")
	f.MonthLabel.SetFont(newFontBlack(-15, 11))
	f.MonthLabel.SetHeight(18)
	f.MonthLabel.SetWidth(15)
	f.MonthLabel.SetLeft(243)
	f.MonthLabel.SetTop(97)
	f.MonthLabel.SetParentColor(false)

	// ================== 日 标签组件 ==================
	f.DayLabel = vcl.NewLabel(f)
	f.DayLabel.SetParent(f)
	f.DayLabel.SetCaption("日")
	f.DayLabel.SetFont(newFontBlack(-15, 11))
	f.DayLabel.SetHeight(18)
	f.DayLabel.SetWidth(15)
	f.DayLabel.SetLeft(322)
	f.DayLabel.SetTop(97)
	f.DayLabel.SetParentColor(false)

	// ================== 时 标签组件 ==================
	f.HourLabel = vcl.NewLabel(f)
	f.HourLabel.SetParent(f)
	f.HourLabel.SetCaption("时")
	f.HourLabel.SetFont(newFontBlack(-15, 11))
	f.HourLabel.SetHeight(18)
	f.HourLabel.SetWidth(15)
	f.HourLabel.SetLeft(162)
	f.HourLabel.SetTop(126)
	f.HourLabel.SetParentColor(false)

	// ================== 分 标签组件 ==================
	f.MinuteLabel = vcl.NewLabel(f)
	f.MinuteLabel.SetParent(f)
	f.MinuteLabel.SetCaption("分")
	f.MinuteLabel.SetFont(newFontBlack(-15, 11))
	f.MinuteLabel.SetHeight(18)
	f.MinuteLabel.SetWidth(15)
	f.MinuteLabel.SetLeft(243)
	f.MinuteLabel.SetTop(126)
	f.MinuteLabel.SetParentColor(false)

	// ================== 秒 标签组件 ==================
	f.SecondLabel = vcl.NewLabel(f)
	f.SecondLabel.SetParent(f)
	f.SecondLabel.SetCaption("秒")
	f.SecondLabel.SetFont(newFontBlack(-15, 11))
	f.SecondLabel.SetHeight(18)
	f.SecondLabel.SetWidth(15)
	f.SecondLabel.SetLeft(322)
	f.SecondLabel.SetTop(126)
	f.SecondLabel.SetParentColor(false)

	// ================== 年 编辑框组件 ==================
	f.YearEdit = vcl.NewEdit(f)
	f.YearEdit.SetParent(f)
	f.YearEdit.SetAlignment(types.TaRightJustify)
	f.YearEdit.SetEnabled(false)
	f.YearEdit.SetFont(newFontBlack(-13, 10))
	f.YearEdit.SetHeight(22)
	f.YearEdit.SetWidth(40)
	f.YearEdit.SetLeft(116)
	f.YearEdit.SetTop(93)
	f.YearEdit.SetParentColor(false)
	f.YearEdit.SetTabOrder(6)
	f.YearEdit.SetText("1999")

	// ================== 月 编辑框组件 ==================
	f.MonthEdit = vcl.NewEdit(f)
	f.MonthEdit.SetParent(f)
	f.MonthEdit.SetAlignment(types.TaRightJustify)
	f.MonthEdit.SetEnabled(false)
	f.MonthEdit.SetFont(newFontBlack(-13, 10))
	f.MonthEdit.SetHeight(22)
	f.MonthEdit.SetWidth(40)
	f.MonthEdit.SetLeft(196)
	f.MonthEdit.SetTop(93)
	f.MonthEdit.SetParentColor(false)
	f.MonthEdit.SetTabOrder(7)
	f.MonthEdit.SetText("01")

	// ================== 日 编辑框组件 ==================
	f.DayEdit = vcl.NewEdit(f)
	f.DayEdit.SetParent(f)
	f.DayEdit.SetAlignment(types.TaRightJustify)
	f.DayEdit.SetEnabled(false)
	f.DayEdit.SetFont(newFontBlack(-13, 10))
	f.DayEdit.SetHeight(22)
	f.DayEdit.SetWidth(40)
	f.DayEdit.SetLeft(274)
	f.DayEdit.SetTop(93)
	f.DayEdit.SetParentColor(false)
	f.DayEdit.SetTabOrder(8)
	f.DayEdit.SetText("01")

	// ================== 时 编辑框组件 ==================
	f.HourEdit = vcl.NewEdit(f)
	f.HourEdit.SetParent(f)
	f.HourEdit.SetAlignment(types.TaRightJustify)
	f.HourEdit.SetEnabled(false)
	f.HourEdit.SetFont(newFontBlack(-13, 10))
	f.HourEdit.SetHeight(22)
	f.HourEdit.SetWidth(40)
	f.HourEdit.SetLeft(116)
	f.HourEdit.SetTop(122)
	f.HourEdit.SetParentColor(false)
	f.HourEdit.SetTabOrder(9)
	f.HourEdit.SetText("01")

	// ================== 分 编辑框组件 ==================
	f.MinuteEdit = vcl.NewEdit(f)
	f.MinuteEdit.SetParent(f)
	f.MinuteEdit.SetAlignment(types.TaRightJustify)
	f.MinuteEdit.SetEnabled(false)
	f.MinuteEdit.SetFont(newFontBlack(-13, 10))
	f.MinuteEdit.SetHeight(22)
	f.MinuteEdit.SetWidth(40)
	f.MinuteEdit.SetLeft(196)
	f.MinuteEdit.SetTop(122)
	f.MinuteEdit.SetParentColor(false)
	f.MinuteEdit.SetTabOrder(10)
	f.MinuteEdit.SetText("01")

	// ================== 秒 编辑框组件 ==================
	f.SecondEdit = vcl.NewEdit(f)
	f.SecondEdit.SetParent(f)
	f.SecondEdit.SetAlignment(types.TaRightJustify)
	f.SecondEdit.SetEnabled(false)
	f.SecondEdit.SetFont(newFontBlack(-13, 10))
	f.SecondEdit.SetHeight(22)
	f.SecondEdit.SetWidth(40)
	f.SecondEdit.SetLeft(274)
	f.SecondEdit.SetTop(122)
	f.SecondEdit.SetParentColor(false)
	f.SecondEdit.SetTabOrder(11)
	f.SecondEdit.SetText("01")

	// ================== 备注 标签组件 ==================
	f.RemarkLabel = vcl.NewLabel(f)
	f.RemarkLabel.SetParent(f)
	f.RemarkLabel.SetCaption("备注")
	f.RemarkLabel.SetFont(newFontBlack(-15, 11))
	f.RemarkLabel.SetHeight(18)
	f.RemarkLabel.SetWidth(31)
	f.RemarkLabel.SetLeft(7)
	f.RemarkLabel.SetTop(156)
	f.RemarkLabel.SetParentColor(false)

	// ================== 备注 编辑框组件 ==================
	f.RemarkEdit = vcl.NewEdit(f)
	f.RemarkEdit.SetParent(f)
	f.RemarkEdit.SetFont(newFontBlack(-13, 10))
	f.RemarkEdit.SetHeight(22)
	f.RemarkEdit.SetWidth(297)
	f.RemarkEdit.SetLeft(40)
	f.RemarkEdit.SetTop(152)
	f.RemarkEdit.SetParentColor(false)
	f.RemarkEdit.SetTabOrder(12)
	f.RemarkEdit.SetText(" ")

	// ================== 导出 按钮组件 ==================
	f.ExportButton = vcl.NewButton(f)
	f.ExportButton.SetParent(f)
	f.ExportButton.SetCaption("导出")
	f.ExportButton.SetFont(newFontBlack(-15, 11))
	f.ExportButton.SetHeight(25)
	f.ExportButton.SetWidth(75)
	f.ExportButton.SetLeft(135)
	f.ExportButton.SetTop(183)
	f.ExportButton.SetTabOrder(13)

	// ================== 保存文件 ==================
	f.ExportSaveDialog = vcl.NewSaveDialog(f)
	f.ExportSaveDialog.SetDefaultExt("pdf")

	// ================== 选择文件 ==================
	f.FontOpenDialog = vcl.NewOpenDialog(f)

	f.DisplayUnitComboBox.SetOnSelect(func(sender vcl.IObject) {
		if MainForm.Caption() == mytypes.AreaCalculator {
			f.TotalAreaEdit.SetText(initial.MyDisplay.SumArea(mytypes.Unit(f.DisplayUnitComboBox.ItemIndex())))
		} else {
			f.TotalWeightEdit.SetText(initial.MyWeightDisplayer.SumWeight(mytypes.WeightUnit(f.DisplayUnitComboBox.ItemIndex())))
		}
	})
}

func (f *TExportPdfForm) OnTotalNumberCheckBoxClick(sender vcl.IObject) {
	if f.TotalNumberCheckBox.Checked() {
		// 导出物体总数
		f.TotalNumberEdit.SetEnabled(true)
	} else {
		// 不导出物体总数
		f.TotalNumberEdit.SetEnabled(false)
	}
}

func (f *TExportPdfForm) OnCustomDateClick(sender vcl.IObject) {
	if f.CustomDate.Checked() {
		// 自定义日期格式
		f.DateFormatComboBox.SetEnabled(false)
		f.YearEdit.SetEnabled(true)
		f.MonthEdit.SetEnabled(true)
		f.DayEdit.SetEnabled(true)
		f.HourEdit.SetEnabled(true)
		f.MinuteEdit.SetEnabled(true)
		f.SecondEdit.SetEnabled(true)
	} else {
		// 预设格式
		f.DateFormatComboBox.SetEnabled(true)
		f.YearEdit.SetEnabled(false)
		f.MonthEdit.SetEnabled(false)
		f.DayEdit.SetEnabled(false)
		f.HourEdit.SetEnabled(false)
		f.MinuteEdit.SetEnabled(false)
		f.SecondEdit.SetEnabled(false)
	}
}

func (f *TExportPdfForm) OnExportButtonClick(sender vcl.IObject) {
	// 验证字体文件
	newFontPath := ""
	var fontErr error
	for {
		newFontPath, fontErr = tools.ValidateFontPath(newFontPath)
		if fontErr != nil {
			vcl.ShowMessage(fontErr.Error())
			if f.FontOpenDialog.Execute() {
				newFontPath = f.FontOpenDialog.FileName()
			} else {
				return
			}
		} else {
			break
		}
	}

	if f.ExportSaveDialog.Execute() {

		var totalArea string
		var totalWeight string
		if MainForm.Caption() == mytypes.AreaCalculator {
			totalArea = f.TotalAreaEdit.Text()
		} else {
			totalWeight = f.TotalWeightEdit.Text()
		}
		totalNumber := ""
		if f.TotalNumberCheckBox.Checked() {
			// 导出物体总数
			totalNumber = f.TotalNumberEdit.Text()
		}
		remark := f.RemarkEdit.Text()
		if remark == " " {
			remark = ""
		}
		timeStamp := ""
		timeFormt := ""
		var useSysTime bool
		if !f.CustomDate.Checked() {
			// 非自定义日期格式
			useSysTime = true
			timeFormt = mytypes.DateFormats[f.DateFormatComboBox.ItemIndex()]
		} else {
			// 自定义日期格式
			useSysTime = false
			temp1 := make([]string, 0)
			year := f.YearEdit.Text()
			month := f.MonthEdit.Text()
			day := f.DayEdit.Text()
			if year != "" {
				temp1 = append(temp1, year)
			}
			if month != "" {
				temp1 = append(temp1, month)
			}
			if day != "" {
				temp1 = append(temp1, day)
			}
			temp1Str := strings.Join(temp1, "-")

			temp2 := make([]string, 0)
			hour := f.HourEdit.Text()
			minute := f.MinuteEdit.Text()
			second := f.SecondEdit.Text()
			if hour != "" {
				temp2 = append(temp2, hour)
			}
			if minute != "" {
				temp2 = append(temp2, minute)
			}
			if second != "" {
				temp2 = append(temp2, second)
			}
			timeStamp = temp1Str + " " + strings.Join(temp2, ":")
		}

		var pdfController *pdf.PdfController
		var pdfWeightController *pdf.PdfWeightController
		if MainForm.Caption() == mytypes.AreaCalculator {
			pdfController = pdf.NewPdfController(totalArea, totalNumber, remark, timeStamp, timeFormt,
				f.ExportSaveDialog.FileName(), useSysTime, mytypes.Unit(f.DisplayUnitComboBox.ItemIndex()))
			if err := pdfController.PrintWorkOrderInPdfFormat(newFontPath); err != nil {
				vcl.ShowMessage(err.Error())
				f.Close()
				return
			}
		} else {
			pdfWeightController = pdf.NewPdfWeightController(totalWeight, totalNumber, remark, timeStamp, timeFormt,
				f.ExportSaveDialog.FileName(), useSysTime, mytypes.WeightUnit(f.DisplayUnitComboBox.ItemIndex()))
			if err := pdfWeightController.WeightPrintWorkOrderInPdfFormat(newFontPath); err != nil {
				vcl.ShowMessage(err.Error())
				f.Close()
				return
			}
		}

		vcl.ShowMessage(mytypes.SaveSuccess)
		f.Close()
	}
}
