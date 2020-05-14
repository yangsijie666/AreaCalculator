package main

import (
	_ "AreaCalculator/initial"
	"github.com/ying32/govcl/vcl"
)

func main() {

	//vcl.Application.SetOnException(func(sender vcl.IObject, e *vcl.Exception) {
	//    // 忽略异常
	//})

	vcl.Application.Initialize()
	vcl.Application.CreateForm(&MainForm)
	vcl.Application.Run()
}
