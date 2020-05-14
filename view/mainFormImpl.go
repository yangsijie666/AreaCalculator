package main

import (
	"AreaCalculator/initial"
	"AreaCalculator/logic"
	"AreaCalculator/models"
	mytypes "AreaCalculator/types"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/types/colors"
	"runtime"
)

func newFontBold(height, size int32) *vcl.TFont {
	f := vcl.NewFont()

	f.SetHeight(height)
	f.SetSize(size)
	f.SetStyle(types.FsBold)

	return f
}

func newFontBlack(height, size int32) *vcl.TFont {
	f := vcl.NewFont()

	f.SetHeight(height)
	f.SetSize(size)
	f.SetColor(colors.ClBlack)

	return f
}

func (f *TMainForm) OnFormCreate(sender vcl.IObject) {
	// ================== 窗口组件 ==================
	f.SetCaption(mytypes.AreaCalculator)
	if runtime.GOOS == "darwin" {
		f.SetHeight(345)
	} else {
		f.SetHeight(360)
	}
	f.SetWidth(673)
	f.SetLeft(8)
	f.SetTop(8)
	f.SetFont(newFontBold(-20, 15))
	f.WorkAreaCenter()
	f.EnabledMaximize(false)
	// 约束窗口组件大小
	f.SetOnConstrainedResize(func(sender vcl.IObject, minWidth, minHeight, maxWidth, maxHeight *int32) {
		*minWidth = f.Width()
		*maxWidth = f.Width()
		*minHeight = f.Height()
		*maxHeight = f.Height()
	})

	// ================== 计算结果显示组件 ==================
	f.DisplayMemo = vcl.NewMemo(f)
	f.DisplayMemo.SetParent(f)
	f.DisplayMemo.SetAlignment(types.TaRightJustify)
	f.DisplayMemo.SetColor(colors.ClWhite)
	f.DisplayMemo.SetFont(newFontBlack(-16, 12))
	f.DisplayMemo.SetHeight(145)
	f.DisplayMemo.SetWidth(374)
	f.DisplayMemo.SetLeft(10)
	f.DisplayMemo.SetTop(10)
	f.DisplayMemo.Lines().SetText(initial.MyDisplay.DisplayResult(mytypes.DisplayUnitDefault))
	f.DisplayMemo.SetParentFont(false)
	f.DisplayMemo.SetReadOnly(true)
	f.DisplayMemo.SetTabOrder(0)
	f.DisplayMemo.SetTabStop(false)
	f.DisplayMemo.SetWantReturns(false)

	// ================== 历史记录显示组件 ==================
	f.HistoryMemo = vcl.NewMemo(f)
	f.HistoryMemo.SetParent(f)
	f.HistoryMemo.SetAlignment(types.TaCenter)
	f.HistoryMemo.SetColor(colors.ClWhite)
	f.HistoryMemo.SetFont(newFontBlack(-13, 10))
	f.HistoryMemo.SetHeight(320)
	f.HistoryMemo.SetWidth(272)
	f.HistoryMemo.SetLeft(390)
	f.HistoryMemo.SetTop(10)
	f.HistoryMemo.Lines().SetText(initial.MyHistoryDisplay.Display(mytypes.DisplayUnitDefault))
	f.HistoryMemo.SetParentFont(false)
	f.HistoryMemo.SetReadOnly(true)
	f.HistoryMemo.SetScrollBars(types.SsAutoBoth)
	f.HistoryMemo.SetTabOrder(1)
	f.HistoryMemo.SetTabStop(false)
	f.DisplayMemo.SetWantReturns(false)

	// ================== 长 标签组件 ==================
	f.HeightLabel = vcl.NewLabel(f)
	f.HeightLabel.SetParent(f)
	f.HeightLabel.SetAlignment(types.TaCenter)
	f.HeightLabel.SetCaption("长")
	f.HeightLabel.SetFont(newFontBlack(-15, 11))
	f.HeightLabel.SetHeight(18)
	f.HeightLabel.SetWidth(15)
	f.HeightLabel.SetLeft(12)
	f.HeightLabel.SetTop(165)
	f.HeightLabel.SetParentColor(false)
	f.HeightLabel.SetParentFont(false)

	// ================== 长 编辑框组件 ==================
	f.LengthEdit = vcl.NewEdit(f)
	f.LengthEdit.SetParent(f)
	f.LengthEdit.SetAlignment(types.TaRightJustify)
	f.LengthEdit.SetFont(newFontBlack(-13, 10))
	f.LengthEdit.SetHeight(22)
	f.LengthEdit.SetWidth(33)
	f.LengthEdit.SetLeft(33)
	f.LengthEdit.SetTop(161)
	f.LengthEdit.SetParentFont(false)
	f.LengthEdit.SetTabOrder(2)
	f.LengthEdit.SetText(mytypes.DefaultLength)

	// ================== 长 列表选择框组件 ==================
	f.LengthUnit = vcl.NewComboBox(f)
	f.LengthUnit.SetParent(f)
	f.LengthUnit.SetFont(newFontBlack(-13, 10))
	f.LengthUnit.SetHeight(20)
	f.LengthUnit.SetWidth(54)
	f.LengthUnit.SetLeft(73)
	f.LengthUnit.SetTop(162)
	f.LengthUnit.SetItemHeight(26)
	f.LengthUnit.Items().Append("毫米")
	f.LengthUnit.Items().Append("厘米")
	f.LengthUnit.Items().Append("分米")
	f.LengthUnit.Items().Append("米")
	f.LengthUnit.SetItemIndex(int32(mytypes.DisplayUnitDefault))
	f.LengthUnit.SetParentFont(false)
	f.LengthUnit.SetStyle(types.CsDropDownList)
	f.LengthUnit.SetTabOrder(3)

	// ================== 宽 标签组件 ==================
	f.WidthLabel = vcl.NewLabel(f)
	f.WidthLabel.SetParent(f)
	f.WidthLabel.SetAlignment(types.TaCenter)
	f.WidthLabel.SetCaption("宽")
	f.WidthLabel.SetFont(newFontBlack(-15, 11))
	f.WidthLabel.SetHeight(18)
	f.WidthLabel.SetWidth(15)
	f.WidthLabel.SetLeft(140)
	f.WidthLabel.SetTop(165)
	f.WidthLabel.SetParentColor(false)
	f.WidthLabel.SetParentFont(false)

	// ================== 宽 编辑框组件 ==================
	f.WidthEdit = vcl.NewEdit(f)
	f.WidthEdit.SetParent(f)
	f.WidthEdit.SetAlignment(types.TaRightJustify)
	f.WidthEdit.SetFont(newFontBlack(-13, 10))
	f.WidthEdit.SetHeight(22)
	f.WidthEdit.SetWidth(33)
	f.WidthEdit.SetLeft(161)
	f.WidthEdit.SetTop(161)
	f.WidthEdit.SetParentFont(false)
	f.WidthEdit.SetTabOrder(4)
	f.WidthEdit.SetText(mytypes.DefaultWidth)

	// ================== 宽 列表选择框组件 ==================
	f.WidthUnit = vcl.NewComboBox(f)
	f.WidthUnit.SetParent(f)
	f.WidthUnit.SetFont(newFontBlack(-13, 10))
	f.WidthUnit.SetHeight(20)
	f.WidthUnit.SetWidth(54)
	f.WidthUnit.SetLeft(201)
	f.WidthUnit.SetTop(162)
	f.WidthUnit.SetItemHeight(26)
	f.WidthUnit.Items().Append("毫米")
	f.WidthUnit.Items().Append("厘米")
	f.WidthUnit.Items().Append("分米")
	f.WidthUnit.Items().Append("米")
	f.WidthUnit.SetItemIndex(int32(mytypes.DisplayUnitDefault))
	f.WidthUnit.SetParentFont(false)
	f.WidthUnit.SetStyle(types.CsDropDownList)
	f.WidthUnit.SetTabOrder(5)

	// ================== 高 标签组件 ==================
	f.HeightLabel = vcl.NewLabel(f)
	f.HeightLabel.SetParent(f)
	f.HeightLabel.SetAlignment(types.TaCenter)
	f.HeightLabel.SetCaption("高")
	f.HeightLabel.SetFont(newFontBlack(-15, 11))
	f.HeightLabel.SetHeight(18)
	f.HeightLabel.SetWidth(15)
	f.HeightLabel.SetLeft(268)
	f.HeightLabel.SetTop(165)
	f.HeightLabel.SetParentColor(false)
	f.HeightLabel.SetParentFont(false)

	// ================== 高 编辑框组件 ==================
	f.HeightEdit = vcl.NewEdit(f)
	f.HeightEdit.SetParent(f)
	f.HeightEdit.SetAlignment(types.TaRightJustify)
	f.HeightEdit.SetFont(newFontBlack(-13, 10))
	f.HeightEdit.SetHeight(22)
	f.HeightEdit.SetWidth(33)
	f.HeightEdit.SetLeft(289)
	f.HeightEdit.SetTop(161)
	f.HeightEdit.SetParentFont(false)
	f.HeightEdit.SetTabOrder(6)
	f.HeightEdit.SetText(mytypes.DefaultHeight)

	// ================== 高 列表选择框组件 ==================
	f.HeightUnit = vcl.NewComboBox(f)
	f.HeightUnit.SetParent(f)
	f.HeightUnit.SetFont(newFontBlack(-13, 10))
	f.HeightUnit.SetHeight(20)
	f.HeightUnit.SetWidth(54)
	f.HeightUnit.SetLeft(329)
	f.HeightUnit.SetTop(162)
	f.HeightUnit.SetItemHeight(26)
	f.HeightUnit.Items().Append("毫米")
	f.HeightUnit.Items().Append("厘米")
	f.HeightUnit.Items().Append("分米")
	f.HeightUnit.Items().Append("米")
	f.HeightUnit.SetItemIndex(int32(mytypes.DisplayUnitDefault))
	f.HeightUnit.SetParentFont(false)
	f.HeightUnit.SetStyle(types.CsDropDownList)
	f.HeightUnit.SetTabOrder(7)

	// ================== 立方体数量 标签组件 ==================
	f.NumberLabel = vcl.NewLabel(f)
	f.NumberLabel.SetParent(f)
	f.NumberLabel.SetAlignment(types.TaCenter)
	f.NumberLabel.SetCaption("立方体数量")
	f.NumberLabel.SetFont(newFontBlack(-15, 11))
	f.NumberLabel.SetHeight(18)
	f.NumberLabel.SetWidth(76)
	f.NumberLabel.SetLeft(12)
	f.NumberLabel.SetTop(196)
	f.NumberLabel.SetParentColor(false)
	f.NumberLabel.SetParentFont(false)

	// ================== 立方体数量 编辑框组件 ==================
	f.NumberEdit = vcl.NewEdit(f)
	f.NumberEdit.SetParent(f)
	f.NumberEdit.SetAlignment(types.TaRightJustify)
	f.NumberEdit.SetFont(newFontBlack(-13, 10))
	f.NumberEdit.SetHeight(22)
	f.NumberEdit.SetWidth(33)
	f.NumberEdit.SetLeft(96)
	f.NumberEdit.SetTop(192)
	f.NumberEdit.SetParentFont(false)
	f.NumberEdit.SetTabOrder(8)
	f.NumberEdit.SetText(mytypes.DefaultNumber)

	// ================== 正方体 选择框组件 ==================
	f.CubeCheckBox = vcl.NewCheckBox(f)
	f.CubeCheckBox.SetParent(f)
	f.CubeCheckBox.SetCaption("正方体")
	f.CubeCheckBox.SetFont(newFontBlack(-15, 11))
	f.CubeCheckBox.SetHeight(19)
	f.CubeCheckBox.SetWidth(68)
	f.CubeCheckBox.SetLeft(150)
	f.CubeCheckBox.SetTop(195)
	f.CubeCheckBox.SetParentFont(false)
	f.CubeCheckBox.SetTabOrder(9)

	// ================== 显示单位 标签组件 ==================
	f.UnitDisplayLabel = vcl.NewLabel(f)
	f.UnitDisplayLabel.SetParent(f)
	f.UnitDisplayLabel.SetAlignment(types.TaCenter)
	f.UnitDisplayLabel.SetCaption("显示单位")
	f.UnitDisplayLabel.SetFont(newFontBlack(-15, 11))
	f.UnitDisplayLabel.SetHeight(18)
	f.UnitDisplayLabel.SetWidth(61)
	f.UnitDisplayLabel.SetLeft(260)
	f.UnitDisplayLabel.SetTop(196)
	f.UnitDisplayLabel.SetParentFont(false)
	f.UnitDisplayLabel.SetParentColor(false)

	// ================== 显示单位 列表选择框组件 ==================
	f.DisplayUnitComboBox = vcl.NewComboBox(f)
	f.DisplayUnitComboBox.SetParent(f)
	f.DisplayUnitComboBox.SetFont(newFontBlack(-13, 10))
	f.DisplayUnitComboBox.SetHeight(20)
	f.DisplayUnitComboBox.SetWidth(54)
	f.DisplayUnitComboBox.SetLeft(329)
	f.DisplayUnitComboBox.SetTop(194)
	f.DisplayUnitComboBox.SetItemHeight(26)
	f.DisplayUnitComboBox.Items().Append("毫米")
	f.DisplayUnitComboBox.Items().Append("厘米")
	f.DisplayUnitComboBox.Items().Append("分米")
	f.DisplayUnitComboBox.Items().Append("米")
	f.DisplayUnitComboBox.SetItemIndex(int32(mytypes.DisplayUnitDefault))
	f.DisplayUnitComboBox.SetParentFont(false)
	f.DisplayUnitComboBox.SetStyle(types.CsDropDownList)
	f.DisplayUnitComboBox.SetTabOrder(10)

	// ================== 密度 标签组件 ==================
	f.DensityLabel = vcl.NewLabel(f)
	f.DensityLabel.SetParent(f)
	f.DensityLabel.SetAlignment(types.TaCenter)
	f.DensityLabel.SetCaption("密度")
	f.DensityLabel.SetFont(newFontBlack(-15, 11))
	f.DensityLabel.SetLeft(12)
	f.DensityLabel.SetTop(228)
	f.DensityLabel.SetParentFont(false)
	f.DensityLabel.SetParentColor(false)

	// ================== 密度 编辑框组件 ==================
	f.DensityEdit = vcl.NewEdit(f)
	f.DensityEdit.SetParent(f)
	f.DensityEdit.SetAlignment(types.TaRightJustify)
	f.DensityEdit.SetFont(newFontBlack(-13, 10))
	f.DensityEdit.SetHeight(22)
	f.DensityEdit.SetWidth(33)
	f.DensityEdit.SetLeft(43)
	f.DensityEdit.SetTop(225)
	f.DensityEdit.SetParentFont(false)
	f.DensityEdit.SetTabOrder(11)
	f.DensityEdit.SetText(mytypes.DefaultDensity)
	f.DensityEdit.SetEnabled(false)

	// ================== 密度（重量） 选择框组件 ==================
	f.DensityWeightUnit = vcl.NewComboBox(f)
	f.DensityWeightUnit.SetParent(f)
	f.DensityWeightUnit.SetFont(newFontBlack(-13, 10))
	f.DensityWeightUnit.SetHeight(20)
	f.DensityWeightUnit.SetWidth(54)
	f.DensityWeightUnit.SetLeft(78)
	f.DensityWeightUnit.SetTop(226)
	f.DensityWeightUnit.SetItemHeight(26)
	f.DensityWeightUnit.Items().Append("毫克")
	f.DensityWeightUnit.Items().Append("克")
	f.DensityWeightUnit.Items().Append("千克")
	f.DensityWeightUnit.Items().Append("吨")
	f.DensityWeightUnit.SetItemIndex(int32(mytypes.DisplayWeightUnitDefault))
	f.DensityWeightUnit.SetParentFont(false)
	f.DensityWeightUnit.SetStyle(types.CsDropDownList)
	f.DensityWeightUnit.SetTabOrder(12)
	f.DensityWeightUnit.SetEnabled(false)

	// ================== 密度单位分隔符 编辑框组件 ==================
	f.DensityEveryLabel = vcl.NewLabel(f)
	f.DensityEveryLabel.SetParent(f)
	f.DensityEveryLabel.SetAlignment(types.TaCenter)
	f.DensityEveryLabel.SetCaption("/")
	f.DensityEveryLabel.SetFont(newFontBlack(-15, 11))
	f.DensityEveryLabel.SetLeft(134)
	f.DensityEveryLabel.SetTop(228)
	f.DensityEveryLabel.SetParentFont(false)
	f.DensityEveryLabel.SetParentColor(false)

	// ================== 密度（体积） 选择框组件 ==================
	f.DensityVolumeUnit = vcl.NewComboBox(f)
	f.DensityVolumeUnit.SetParent(f)
	f.DensityVolumeUnit.SetFont(newFontBlack(-13, 10))
	f.DensityVolumeUnit.SetHeight(20)
	f.DensityVolumeUnit.SetWidth(85)
	f.DensityVolumeUnit.SetLeft(143)
	f.DensityVolumeUnit.SetTop(226)
	f.DensityVolumeUnit.SetItemHeight(26)
	f.DensityVolumeUnit.Items().Append("立方毫米")
	f.DensityVolumeUnit.Items().Append("立方厘米")
	f.DensityVolumeUnit.Items().Append("立方分米")
	f.DensityVolumeUnit.Items().Append("立方米")
	f.DensityVolumeUnit.SetItemIndex(int32(mytypes.DisplayUnitDefault))
	f.DensityVolumeUnit.SetParentFont(false)
	f.DensityVolumeUnit.SetStyle(types.CsDropDownList)
	f.DensityVolumeUnit.SetTabOrder(13)
	f.DensityVolumeUnit.SetEnabled(false)

	// ================== 重量显示单位 标签组件 ==================
	f.WeightUnitDisplayLabel = vcl.NewLabel(f)
	f.WeightUnitDisplayLabel.SetParent(f)
	f.WeightUnitDisplayLabel.SetAlignment(types.TaCenter)
	f.WeightUnitDisplayLabel.SetCaption("重量显示单位")
	f.WeightUnitDisplayLabel.SetFont(newFontBlack(-15, 11))
	f.WeightUnitDisplayLabel.SetLeft(230)
	f.WeightUnitDisplayLabel.SetTop(228)
	f.WeightUnitDisplayLabel.SetParentFont(false)
	f.WeightUnitDisplayLabel.SetParentColor(false)

	// ================== 重量显示单位 列表选择框组件 ==================
	f.WeightUnitDisplayComboBox = vcl.NewComboBox(f)
	f.WeightUnitDisplayComboBox.SetParent(f)
	f.WeightUnitDisplayComboBox.SetFont(newFontBlack(-13, 10))
	f.WeightUnitDisplayComboBox.SetHeight(20)
	f.WeightUnitDisplayComboBox.SetWidth(54)
	f.WeightUnitDisplayComboBox.SetLeft(329)
	f.WeightUnitDisplayComboBox.SetTop(226)
	f.WeightUnitDisplayComboBox.SetItemHeight(26)
	f.WeightUnitDisplayComboBox.Items().Append("毫克")
	f.WeightUnitDisplayComboBox.Items().Append("克")
	f.WeightUnitDisplayComboBox.Items().Append("千克")
	f.WeightUnitDisplayComboBox.Items().Append("吨")
	f.WeightUnitDisplayComboBox.SetItemIndex(int32(mytypes.DisplayWeightUnitDefault))
	f.WeightUnitDisplayComboBox.SetParentFont(false)
	f.WeightUnitDisplayComboBox.SetStyle(types.CsDropDownList)
	f.WeightUnitDisplayComboBox.SetTabOrder(14)
	f.WeightUnitDisplayComboBox.SetEnabled(false)

	// ================== "+" 按钮组件 ==================
	f.PlusSpeedButton = vcl.NewSpeedButton(f)
	f.PlusSpeedButton.SetParent(f)
	f.PlusSpeedButton.SetCaption("+")
	f.PlusSpeedButton.SetFont(newFontBold(-27, 20))
	f.PlusSpeedButton.SetHeight(75)
	f.PlusSpeedButton.SetWidth(103)
	f.PlusSpeedButton.SetLeft(12)
	f.PlusSpeedButton.SetTop(256)
	f.PlusSpeedButton.SetParentFont(false)

	// ================== "-" 按钮组件 ==================
	f.MinusSpeedButton = vcl.NewSpeedButton(f)
	f.MinusSpeedButton.SetParent(f)
	f.MinusSpeedButton.SetCaption("-")
	f.MinusSpeedButton.SetFont(newFontBold(-27, 20))
	f.MinusSpeedButton.SetHeight(75)
	f.MinusSpeedButton.SetWidth(103)
	f.MinusSpeedButton.SetLeft(145)
	f.MinusSpeedButton.SetTop(256)
	f.MinusSpeedButton.SetParentFont(false)

	// ================== "AC" 按钮组件 ==================
	f.ClearSpeedButton = vcl.NewSpeedButton(f)
	f.ClearSpeedButton.SetParent(f)
	f.ClearSpeedButton.SetCaption(mytypes.AllClearCaption)
	f.ClearSpeedButton.SetFont(newFontBold(-27, 20))
	f.ClearSpeedButton.SetHeight(75)
	f.ClearSpeedButton.SetWidth(103)
	f.ClearSpeedButton.SetLeft(280)
	f.ClearSpeedButton.SetTop(256)
	f.ClearSpeedButton.SetParentFont(false)

	// ================== 主菜单 ==================
	f.MainMenu = vcl.NewMainMenu(f)
	f.CalculatorTypeMenuItem = vcl.NewMenuItem(f)
	f.CalculatorTypeMenuItem.SetCaption("类型")
	f.MainMenu.Items().Add(f.CalculatorTypeMenuItem)
	f.ExportMenuItem = vcl.NewMenuItem(f)
	f.ExportMenuItem.SetCaption("导出")
	f.MainMenu.Items().Add(f.ExportMenuItem)

	// 计算器类型菜单添加 添加两种计算器类型子菜单
	f.AreaCalculatorTypeMenuItem = vcl.NewMenuItem(f)
	f.AreaCalculatorTypeMenuItem.SetCaption(mytypes.AreaCalculator)
	f.CalculatorTypeMenuItem.Add(f.AreaCalculatorTypeMenuItem)
	if runtime.GOOS == "windows" {
		f.AreaCalculatorTypeMenuItem.SetShortCutFromString("Ctrl+A")
	} else if runtime.GOOS == "darwin" {
		f.AreaCalculatorTypeMenuItem.SetShortCutFromString("Meta+A")
	}
	f.WeightCalculatorTypeMenuItem = vcl.NewMenuItem(f)
	f.WeightCalculatorTypeMenuItem.SetCaption(mytypes.WeightCalculator)
	f.CalculatorTypeMenuItem.Add(f.WeightCalculatorTypeMenuItem)
	if runtime.GOOS == "windows" {
		f.WeightCalculatorTypeMenuItem.SetShortCutFromString("Ctrl+W")
	} else if runtime.GOOS == "darwin" {
		f.WeightCalculatorTypeMenuItem.SetShortCutFromString("Meta+W")
	}

	// 导出主菜单添加 导出pdf子菜单
	f.ExportPdfMenuItem = vcl.NewMenuItem(f)
	f.ExportPdfMenuItem.SetCaption("PDF")
	f.ExportMenuItem.Add(f.ExportPdfMenuItem)
	if runtime.GOOS == "windows" {
		f.ExportPdfMenuItem.SetShortCutFromString("Ctrl+P")
	} else if runtime.GOOS == "darwin" {
		f.ExportPdfMenuItem.SetShortCutFromString("Meta+P")
	}
}

// 面积计算器 被点击时
func (f *TMainForm) OnAreaCalculatorTypeMenuItemClick(sender vcl.IObject) {
	f.SetCaption(mytypes.AreaCalculator)
	f.UnitDisplayLabel.SetCaption("显示单位")
	f.UnitDisplayLabel.SetLeft(260)
	f.DisplayMemo.Lines().SetText(initial.MyDisplay.DisplayResult(mytypes.Unit(f.DisplayUnitComboBox.ItemIndex())))
	f.HistoryMemo.Lines().SetText(initial.MyHistoryDisplay.Display(mytypes.Unit(f.DisplayUnitComboBox.ItemIndex())))
	f.DensityEdit.SetEnabled(false)
	f.DensityWeightUnit.SetEnabled(false)
	f.DensityVolumeUnit.SetEnabled(false)
	f.WeightUnitDisplayComboBox.SetEnabled(false)
}

// 重量计算器 被点击时
func (f *TMainForm) OnWeightCalculatorTypeMenuItemClick(sender vcl.IObject) {
	f.SetCaption(mytypes.WeightCalculator)
	f.UnitDisplayLabel.SetCaption("长度显示单位")
	f.UnitDisplayLabel.SetLeft(230)
	f.DisplayMemo.Lines().SetText(initial.MyWeightDisplayer.DisplayWeightResult(
		mytypes.Unit(f.DisplayUnitComboBox.ItemIndex()), mytypes.WeightUnit(f.WeightUnitDisplayComboBox.ItemIndex())))
	f.HistoryMemo.Lines().SetText(initial.MyWeightHistoryDisplayer.Display(
		mytypes.Unit(f.DisplayUnitComboBox.ItemIndex()), mytypes.WeightUnit(f.WeightUnitDisplayComboBox.ItemIndex())))
	f.DensityEdit.SetEnabled(true)
	f.DensityWeightUnit.SetEnabled(true)
	f.DensityVolumeUnit.SetEnabled(true)
	f.WeightUnitDisplayComboBox.SetEnabled(true)
}

// 导出 PDF 被点击时
func (f *TMainForm) OnExportPdfMenuItemClick(sender vcl.IObject) {
	vcl.Application.CreateForm(&ExportPdfForm)
	ExportPdfForm.ShowModal()
	ExportPdfForm.Free()
}

// 长 编辑框修改时，AC 变 C
func (f *TMainForm) OnLengthEditChange(sender vcl.IObject) {
	if f.LengthEdit.Text() != mytypes.DefaultLength {
		f.ClearSpeedButton.SetCaption(mytypes.ClearCaption)
	}
}

// 宽 编辑框修改时，AC 变 C
func (f *TMainForm) OnWidthEditChange(sender vcl.IObject) {
	if f.WidthEdit.Text() != mytypes.DefaultWidth {
		f.ClearSpeedButton.SetCaption(mytypes.ClearCaption)
	}
}

// 高 编辑框修改时，AC 变 C
func (f *TMainForm) OnHeightEditChange(sender vcl.IObject) {
	if f.HeightEdit.Text() != mytypes.DefaultHeight {
		f.ClearSpeedButton.SetCaption(mytypes.ClearCaption)
	}
}

// 立方体数量 编辑框修改时，AC 变 C
func (f *TMainForm) OnNumberEditChange(sender vcl.IObject) {
	if f.NumberEdit.Text() != mytypes.DefaultNumber {
		f.ClearSpeedButton.SetCaption(mytypes.ClearCaption)
	}
}

// 长单位 选框的选择
func (f *TMainForm) OnLengthUnitSelect(sender vcl.IObject) {
	f.ClearSpeedButton.SetCaption(mytypes.ClearCaption)
}

// 宽单位 选框的选择
func (f *TMainForm) OnWidthUnitSelect(sender vcl.IObject) {
	f.ClearSpeedButton.SetCaption(mytypes.ClearCaption)
}

// 高单位 选框的选择
func (f *TMainForm) OnHeightUnitSelect(sender vcl.IObject) {
	f.ClearSpeedButton.SetCaption(mytypes.ClearCaption)
}

// 当正方体选择框被选中，宽 和 高的编辑框设为不可用
func (f *TMainForm) OnCubeCheckBoxClick(sender vcl.IObject) {
	f.ClearSpeedButton.SetCaption(mytypes.ClearCaption)
	if f.CubeCheckBox.Checked() {
		f.WidthEdit.SetEnabled(false)
		f.HeightEdit.SetEnabled(false)
	} else {
		f.WidthEdit.SetEnabled(true)
		f.HeightEdit.SetEnabled(true)
	}
}

// 显示单位 选框的选择
func (f *TMainForm) OnDisplayUnitComboBoxSelect(sender vcl.IObject) {
	f.ClearSpeedButton.SetCaption(mytypes.ClearCaption)
	displayUnit := mytypes.Unit(f.DisplayUnitComboBox.ItemIndex())
	if f.Caption() == mytypes.AreaCalculator {
		f.DisplayMemo.Lines().SetText(initial.MyDisplay.DisplayResult(displayUnit))
		f.HistoryMemo.Lines().SetText(initial.MyHistoryDisplay.Display(displayUnit))
	} else {
		displayWeightUnit := mytypes.WeightUnit(f.WeightUnitDisplayComboBox.ItemIndex())
		f.DisplayMemo.Lines().SetText(initial.MyWeightDisplayer.DisplayWeightResult(displayUnit, displayWeightUnit))
		f.HistoryMemo.Lines().SetText(initial.MyWeightHistoryDisplayer.Display(displayUnit, displayWeightUnit))
	}
}

// 密度 编辑框修改时，AC 变 C
func (f *TMainForm) OnDensityEditChange(sender vcl.IObject) {
	if f.DensityEdit.Text() != mytypes.DefaultDensity {
		f.ClearSpeedButton.SetCaption(mytypes.ClearCaption)
	}
}

// 密度 重量单位 选框的选择
func (f *TMainForm) OnDensityWeightUnitSelect(sender vcl.IObject) {
	f.ClearSpeedButton.SetCaption(mytypes.ClearCaption)
}

// 密度 体积单位 选框的选择
func (f *TMainForm) OnDensityVolumeUnitSelect(sender vcl.IObject) {
	f.ClearSpeedButton.SetCaption(mytypes.ClearCaption)
}

// 重量显示单位 选框的选择
func (f *TMainForm) OnWeightUnitDisplayComboBoxSelect(sender vcl.IObject) {
	f.ClearSpeedButton.SetCaption(mytypes.ClearCaption)
	displayUnit := mytypes.Unit(f.DisplayUnitComboBox.ItemIndex())
	displayWeightUnit := mytypes.WeightUnit(f.WeightUnitDisplayComboBox.ItemIndex())
	f.DisplayMemo.Lines().SetText(initial.MyWeightDisplayer.DisplayWeightResult(displayUnit, displayWeightUnit))
	f.HistoryMemo.Lines().SetText(initial.MyWeightHistoryDisplayer.Display(displayUnit, displayWeightUnit))
}

// AC 按键的点击
func (f *TMainForm) OnClearSpeedButtonClick(sender vcl.IObject) {
	f.LengthEdit.SetText(mytypes.DefaultLength)
	f.LengthUnit.SetItemIndex(int32(mytypes.DisplayUnitDefault))
	f.WidthEdit.SetText(mytypes.DefaultWidth)
	f.WidthUnit.SetItemIndex(int32(mytypes.DisplayUnitDefault))
	f.HeightEdit.SetText(mytypes.DefaultHeight)
	f.HeightUnit.SetItemIndex(int32(mytypes.DisplayUnitDefault))
	f.NumberEdit.SetText(mytypes.DefaultNumber)
	f.CubeCheckBox.SetChecked(false)
	f.DisplayUnitComboBox.SetItemIndex(int32(mytypes.DisplayUnitDefault))

	switch f.Caption() {
	case mytypes.AreaCalculator:
		switch f.ClearSpeedButton.Caption() {
		case mytypes.AllClearCaption:
			// 全部复位
			initial.MyDisplay = models.NewDisplay()
			initial.MyHistoryDisplay = models.NewHistoryDisplay()
			initial.CurOperateNumber = 1
			f.DisplayMemo.Lines().SetText(initial.MyDisplay.DisplayResult(mytypes.DisplayUnitDefault))
			f.HistoryMemo.Lines().SetText(initial.MyHistoryDisplay.Display(mytypes.DisplayUnitDefault))
		case mytypes.ClearCaption:
			f.ClearSpeedButton.SetCaption(mytypes.AllClearCaption)
		}
	case mytypes.WeightCalculator:
		switch f.ClearSpeedButton.Caption() {
		case mytypes.AllClearCaption:
			// 全部复位
			initial.MyWeightDisplayer = models.NewWeightDisplayer()
			initial.MyWeightHistoryDisplayer = models.NewWeightHistoryDisplayer()
			initial.CurWeightOperateNumber = 1
			f.DisplayMemo.Lines().SetText(initial.MyWeightDisplayer.DisplayWeightResult(mytypes.DisplayUnitDefault, mytypes.DisplayWeightUnitDefault))
			f.HistoryMemo.Lines().SetText(initial.MyWeightHistoryDisplayer.Display(mytypes.DisplayUnitDefault, mytypes.DisplayWeightUnitDefault))
		case mytypes.ClearCaption:
			f.ClearSpeedButton.SetCaption(mytypes.AllClearCaption)
		}

		f.DensityEdit.SetText(mytypes.DefaultDensity)
		f.DensityWeightUnit.SetItemIndex(int32(mytypes.DisplayWeightUnitDefault))
		f.DensityVolumeUnit.SetItemIndex(int32(mytypes.DisplayUnitDefault))
		f.WeightUnitDisplayComboBox.SetItemIndex(int32(mytypes.DisplayWeightUnitDefault))
	default:
	}
}

// + 按键的点击
func (f *TMainForm) OnPlusSpeedButtonClick(sender vcl.IObject) {
	switch f.Caption() {
	case mytypes.AreaCalculator:
		if f.CubeCheckBox.Checked() {
			// 正方体
			err := logic.PlusCubes(f.LengthEdit.Text(), f.LengthUnit.ItemIndex(), f.NumberEdit.Text())
			if err != nil {
				vcl.ShowMessage(err.Error())
				return
			}
		} else {
			// 长方体
			err := logic.PlusCuboids(f.LengthEdit.Text(), f.WidthEdit.Text(), f.HeightEdit.Text(),
				f.LengthUnit.ItemIndex(), f.WidthUnit.ItemIndex(), f.HeightUnit.ItemIndex(), f.NumberEdit.Text())
			if err != nil {
				vcl.ShowMessage(err.Error())
				return
			}
		}
		displayUnit := mytypes.Unit(f.DisplayUnitComboBox.ItemIndex())
		f.DisplayMemo.Lines().SetText(initial.MyDisplay.DisplayResult(displayUnit))
		f.HistoryMemo.Lines().SetText(initial.MyHistoryDisplay.Display(displayUnit))
	case mytypes.WeightCalculator:
		if f.CubeCheckBox.Checked() {
			// 正方体
			err := logic.PlusCubesWeight(f.LengthEdit.Text(), f.LengthUnit.ItemIndex(), f.NumberEdit.Text(),
				f.DensityEdit.Text(), f.DensityWeightUnit.ItemIndex(), f.DensityVolumeUnit.ItemIndex())
			if err != nil {
				vcl.ShowMessage(err.Error())
				return
			}
		} else {
			// 长方体
			err := logic.PlusCuboidsWeight(f.LengthEdit.Text(), f.WidthEdit.Text(), f.HeightEdit.Text(),
				f.LengthUnit.ItemIndex(), f.WidthUnit.ItemIndex(), f.HeightUnit.ItemIndex(), f.NumberEdit.Text(),
				f.DensityEdit.Text(), f.DensityWeightUnit.ItemIndex(), f.DensityVolumeUnit.ItemIndex())
			if err != nil {
				vcl.ShowMessage(err.Error())
				return
			}
		}
		lengthDisplayUnit := mytypes.Unit(f.DisplayUnitComboBox.ItemIndex())
		weightDisplayUnit := mytypes.WeightUnit(f.WeightUnitDisplayComboBox.ItemIndex())
		f.DisplayMemo.Lines().SetText(initial.MyWeightDisplayer.DisplayWeightResult(lengthDisplayUnit, weightDisplayUnit))
		f.HistoryMemo.Lines().SetText(initial.MyWeightHistoryDisplayer.Display(lengthDisplayUnit, weightDisplayUnit))
	default:
	}
	f.ClearSpeedButton.SetCaption(mytypes.ClearCaption)
}

// - 按键的点击
func (f *TMainForm) OnMinusSpeedButtonClick(sender vcl.IObject) {
	switch f.Caption() {
	case mytypes.AreaCalculator:
		if f.CubeCheckBox.Checked() {
			// 正方体
			err := logic.MinusCubes(f.LengthEdit.Text(), f.LengthUnit.ItemIndex(), f.NumberEdit.Text())
			if err != nil {
				vcl.ShowMessage(err.Error())
				return
			}
		} else {
			// 长方体
			err := logic.MinusCuboids(f.LengthEdit.Text(), f.WidthEdit.Text(), f.HeightEdit.Text(),
				f.LengthUnit.ItemIndex(), f.WidthUnit.ItemIndex(), f.HeightUnit.ItemIndex(), f.NumberEdit.Text())
			if err != nil {
				vcl.ShowMessage(err.Error())
				return
			}
		}
		displayUnit := mytypes.Unit(f.DisplayUnitComboBox.ItemIndex())
		f.DisplayMemo.Lines().SetText(initial.MyDisplay.DisplayResult(displayUnit))
		f.HistoryMemo.Lines().SetText(initial.MyHistoryDisplay.Display(displayUnit))
	case mytypes.WeightCalculator:
		if f.CubeCheckBox.Checked() {
			// 正方体
			err := logic.MinusCubesWeight(f.LengthEdit.Text(), f.LengthUnit.ItemIndex(), f.NumberEdit.Text(),
				f.DensityEdit.Text(), f.DensityWeightUnit.ItemIndex(), f.DensityVolumeUnit.ItemIndex())
			if err != nil {
				vcl.ShowMessage(err.Error())
				return
			}
		} else {
			// 长方体
			err := logic.MinusCuboidsWeight(f.LengthEdit.Text(), f.WidthEdit.Text(), f.HeightEdit.Text(),
				f.LengthUnit.ItemIndex(), f.WidthUnit.ItemIndex(), f.HeightUnit.ItemIndex(), f.NumberEdit.Text(),
				f.DensityEdit.Text(), f.DensityWeightUnit.ItemIndex(), f.DensityVolumeUnit.ItemIndex())
			if err != nil {
				vcl.ShowMessage(err.Error())
				return
			}
		}
	default:
	}

	f.ClearSpeedButton.SetCaption(mytypes.ClearCaption)
}
