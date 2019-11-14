// 由res2go自动生成。
// 在这里写你的事件。

package main

import (
	"fmt"
	"github.com/a97077088/nifdc"
	"github.com/ying32/govcl/vcl"
)

//::private::
type TFormjianceFields struct {
	Td *nifdc.UploadData
}

func (f *TFormjiance) OnFormShow(sender vcl.IObject) {
	if f.Td == nil {
		return
	}
	f.ListView1.Clear()

	items := f.ListView1.Items()
	tditem := f.Td.Subitem()
	for idx, tdit := range tditem {
		item := items.Add()
		item.SetCaption(fmt.Sprintf("%d", idx+1))
		subitem := item.SubItems()
		subitem.Add(tdit["检验项目"])
		subitem.Add(tdit["检验结果"])
		subitem.Add(tdit["结果单位"])
		subitem.Add(tdit["结果判定"])
		subitem.Add(tdit["检验依据"])
		subitem.Add(tdit["判定依据"])
		subitem.Add(tdit["最大允许限"])
		subitem.Add(tdit["最小允许限"])
		subitem.Add(tdit["允许限单位"])
		subitem.Add(tdit["方法检出限"])
		subitem.Add(tdit["检出限单位"])
	}

}

func (f *TFormjiance) OnListView1Resize(sender vcl.IObject) {
	go vcl.ThreadSync(func() {
		lastitem := f.ListView1.Column(f.ListView1.Columns().Count() - 1)
		lastitem.SetWidth(lastitem.Width() - 10)
	})
}
