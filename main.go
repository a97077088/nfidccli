// 由res2go自动生成。
package main

import (
	"github.com/ying32/govcl/vcl"
)

func main() {
	vcl.Application.SetFormScaled(true)
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	vcl.Application.CreateForm(&FormLogin)
	vcl.Application.CreateForm(&FormHome)
	vcl.Application.CreateForm(&Formjiance)
	vcl.Application.Run()
}
