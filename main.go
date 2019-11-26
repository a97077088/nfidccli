// 由res2go自动生成。
package main

import (
	"github.com/ying32/govcl/vcl"
)

func main() {
	vcl.Application.SetFormScaled(true)
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
    vcl.Application.CreateForm(formLoginBytes, &FormLogin)
    vcl.Application.CreateForm(formHomeBytes, &FormHome)
    vcl.Application.CreateForm(formjianceBytes, &Formjiance)
	vcl.Application.Run()
}
