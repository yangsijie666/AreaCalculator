package types

const (
	AreaCalculator   = "面积计算器"
	WeightCalculator = "重量计算器"

	AllClearCaption = "AC"
	ClearCaption    = "C"

	DisplayUnitDefault       = CM
	DisplayWeightUnitDefault = G

	DefaultLength  = "0.0"
	DefaultWidth   = "0.0"
	DefaultHeight  = "0.0"
	DefaultNumber  = "1"
	DefaultDensity = "0.0"

	DisplayResultInitStr = `长  0    宽  0    高  0
数量  0
面积  0
------------------------------------------------
物体总数  0
总面积  0`

	DisplayResultStr = `
------------------------------------------------
物体总数  %d
总面积  %s`

	WeightDisplayResultInitStr = `长  0    宽  0    高  0
数量  0
重量  0
------------------------------------------------
物体总数  0
总重量  0`

	WeightDisplayResultStr = `
------------------------------------------------
物体总数  %d
总重量  %s`

	HistoryDisplayDefaultStr = "历史记录"

	OperationDisplayStr = `编号 No.%d
"%s" 操作
-----------------------
%s
-----------------------
物体总数  %d
总面积  %s
========================
`

	OperationWeightDisplayStr = `编号 No.%d
"%s" 操作
-----------------------
%s
-----------------------
物体总数  %d
总体积  %s
========================
`

	CuboidsDisplayStr = `长  %s    宽  %s    高  %s
数量  %d
面积  %s`

	CuboidsWeightDisplayStr = `长  %s    宽  %s    高  %s
数量  %d
重量  %s`

	CubesDisplayStr = `长  %s
数量  %d
面积  %s`

	CubesWeightDisplayStr = `长  %s
数量  %d
重量  %s`

	ErrMsgTemplate       = "\"%s\" 输入有误! 输入的值为\"%s\", 应为大于\"0\"的数"
	NumberErrMsgTemplate = "\"立方体数量\" 输入有误! 输入的值为\"%s\", 应为大于等于\"1\"的整数"

	UnknownError = "未知错误"
	SaveSuccess  = "导出成功"

	DefaultFontDir        = "./simhei.ttf"
	DefaultWindowsFontDir = `C:\Windows\Fonts\simhei.ttf`
	DefaultMacFontDir     = "/System/Library/Fonts/simhei.ttf"

	CanNotFindFontFile = "未找到字体文件"
	FontFileExtError   = "仅支持 .ttf 格式的字体文件"

	WeightInfoNonexistent = "重量信息不存在"
)
