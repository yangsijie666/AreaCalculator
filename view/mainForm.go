package main

import (
    _ "github.com/ying32/govcl/pkgs/macapp"
    "github.com/ying32/govcl/vcl"
    // 如果你使用自定义的syso文件则不要引用此包
    _ "github.com/ying32/govcl/pkgs/winappres"
)

type TMainForm struct {
    *vcl.TForm

    DisplayMemo         *vcl.TMemo
    HistoryMemo         *vcl.TMemo
    LengthLabel         *vcl.TLabel
    LengthEdit          *vcl.TEdit
    LengthUnit          *vcl.TComboBox
    WidthLabel          *vcl.TLabel
    WidthEdit           *vcl.TEdit
    WidthUnit           *vcl.TComboBox
    HeightLabel         *vcl.TLabel
    HeightEdit          *vcl.TEdit
    HeightUnit          *vcl.TComboBox
    NumberLabel         *vcl.TLabel
    NumberEdit          *vcl.TEdit
    CubeCheckBox        *vcl.TCheckBox
    UnitDisplayLabel    *vcl.TLabel
    DisplayUnitComboBox *vcl.TComboBox
    PlusSpeedButton     *vcl.TSpeedButton
    MinusSpeedButton    *vcl.TSpeedButton
    ClearSpeedButton    *vcl.TSpeedButton

    MainMenu            *vcl.TMainMenu
    ExportMenuItem      *vcl.TMenuItem
    ExportPdfMenuItem   *vcl.TMenuItem

    ExportPdfForm       *TExportPdfForm
}

var MainForm *TMainForm
