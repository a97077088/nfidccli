// 由res2go自动生成。
// 在这里写你的事件。

package main

import (
    "errors"
    "fmt"
    "github.com/360EntSecGroup-Skylar/excelize"
    "github.com/a97077088/addrmgr"
    "github.com/a97077088/nettool"
    "github.com/a97077088/nifdc"
    "github.com/ying32/govcl/vcl"
)

//::private::
type TFormMainFields struct {
    ck string
    lt string
    execution string
    uid string
    chs []*nifdc.Channel
}
func (f *TFormMain) OnButton1Click(sender vcl.IObject) {
    f.Button1.SetEnabled(false)
    err:= func() error{
        var err error
        f.ck,f.lt,f.execution,err=nifdc.Initck(nil)
        if err!=nil{
            return err
        }
        pd:="1234567"
        pd+="8"
        scks,err:=nifdc.Login("15738889730",pd,f.ck,f.lt,f.execution,nil)
        if err!=nil{
            return err
        }
        f.ck=scks
        uid,chs,err:=nifdc.TaskIndex(f.ck,nil)
        if err!=nil{
            return err
        }
        f.uid=uid
        f.chs=chs
        f.Cbb1.Clear()
        for _,ch:=range chs{
            f.Cbb1.Items().Add(ch.Name)
        }
        //f.ck=fmt.Sprintf("%s;%s",f.ck,scks)
        return nil
    }()
    if err!=nil{
        vcl.ShowMessage(err.Error())
        f.Button1.SetEnabled(true)
    }else{
        f.Button1.SetCaption("登录成功")
    }

}


func (f *TFormMain) OnFormCreate(sender vcl.IObject) {

}


func (f *TFormMain) OnLabel2Click(sender vcl.IObject) {
}





func (f *TFormMain) OnButton3Click(sender vcl.IObject) {
    f.Button3.SetEnabled(false)

    go func() {
        defer f.Button3.SetEnabled(true)

        err:= func() error{
            if f.Cbb1.ItemIndex()==-1{
                return errors.New("必须先选择通道")
            }
            var err error
            f.ck,err=nifdc.Switchchannel(f.uid,f.chs[f.Cbb1.ItemIndex()].Type,f.ck,nil)
            if err!=nil{
                return err
            }
            dt,err:=nifdc.DownData(f.ck,nil)
            if err!=nil{
                return err
            }
            vcl.ThreadSync(func() {
                f.Gauge1.SetProgress(0)
                f.Gauge1.SetMaxValue(int32(len(dt.Data)))
            })

            xlsxsheet:="Sheet1"
            Sheetidx:=1
            xlsx := excelize.NewFile()
            xlsx.SetCellValue(xlsxsheet, "A1", "委托单位")
            xlsx.SetCellValue(xlsxsheet, "B1", "抽样地点")
            xlsx.SetCellValue(xlsxsheet, "C1", "抽样单号")
            xlsx.SetCellValue(xlsxsheet, "D1", "检验类型")
            xlsx.SetCellValue(xlsxsheet, "E1", "抽送样人")
            xlsx.SetCellValue(xlsxsheet, "F1", "受检单位")
            xlsx.SetCellValue(xlsxsheet, "G1", "地址")
            xlsx.SetCellValue(xlsxsheet, "H1", "联系人")
            xlsx.SetCellValue(xlsxsheet, "I1", "电话")
            xlsx.SetCellValue(xlsxsheet, "J1", "生产单位地址")
            xlsx.SetCellValue(xlsxsheet, "K1", "生产单位")
            xlsx.SetCellValue(xlsxsheet, "L1", "生产单位联系人")
            xlsx.SetCellValue(xlsxsheet, "M1", "生产单位电话")
            xlsx.SetCellValue(xlsxsheet, "N1", "商标")
            xlsx.SetCellValue(xlsxsheet, "O1", "样品名称br")
            xlsx.SetCellValue(xlsxsheet, "P1", "生产日期")
            xlsx.SetCellValue(xlsxsheet, "Q1", "保质期")
            xlsx.SetCellValue(xlsxsheet, "R1", "生产批号")
            xlsx.SetCellValue(xlsxsheet, "S1", "规格型号")
            xlsx.SetCellValue(xlsxsheet, "T1", "样品等级")
            xlsx.SetCellValue(xlsxsheet, "U1", "抽到样日期")
            xlsx.SetCellValue(xlsxsheet, "V1", "抽样方式")
            xlsx.SetCellValue(xlsxsheet, "W1", "样品状态")
            xlsx.SetCellValue(xlsxsheet, "X1", "样品状态2")
            xlsx.SetCellValue(xlsxsheet, "Y1", "保存条件")
            xlsx.SetCellValue(xlsxsheet, "Z1", "抽样基数")
            xlsx.SetCellValue(xlsxsheet, "AA1", "样品数")
            xlsx.SetCellValue(xlsxsheet, "AB1", "检验依据")
            xlsx.SetCellValue(xlsxsheet, "AC1", "备注")
            //
            //xlsx.SetColWidth(xlsxsheet,"A","AC",20)

            for _,d:=range dt.Data{
                Sheetidx++
                xlsx.SetRowHeight(xlsxsheet,Sheetidx,30)
                itr,err:=nettool.RNet_Call_1(&nettool.RNetOptions{
                }, func(source *addrmgr.AddrSource) (i interface{}, e error) {
                    tr,err:=nifdc.Viewnormalsample(d.Sample_code,f.ck,nil)
                    if err!=nil{
                        return nil,err
                    }
                    return tr,nil
                })
                if err!=nil{
                    return err
                }
                tr:=itr.(map[string]string)
                xlsx.SetCellValue(xlsxsheet,fmt.Sprintf("A%d",Sheetidx),tr["委托单位"])
                xlsx.SetCellValue(xlsxsheet,fmt.Sprintf("B%d",Sheetidx),tr["抽样地点"])
                xlsx.SetCellValue(xlsxsheet,fmt.Sprintf("C%d",Sheetidx),tr["抽样单号"])
                xlsx.SetCellValue(xlsxsheet,fmt.Sprintf("D%d",Sheetidx),tr["检验类型"])
                xlsx.SetCellValue(xlsxsheet,fmt.Sprintf("E%d",Sheetidx),tr["抽送样人"])
                xlsx.SetCellValue(xlsxsheet,fmt.Sprintf("F%d",Sheetidx),tr["受检单位"])
                xlsx.SetCellValue(xlsxsheet,fmt.Sprintf("G%d",Sheetidx),tr["地址"])
                xlsx.SetCellValue(xlsxsheet,fmt.Sprintf("H%d",Sheetidx),tr["联系人"])
                xlsx.SetCellValue(xlsxsheet,fmt.Sprintf("I%d",Sheetidx),tr["电话"])
                xlsx.SetCellValue(xlsxsheet,fmt.Sprintf("J%d",Sheetidx),tr["生产单位地址"])
                xlsx.SetCellValue(xlsxsheet,fmt.Sprintf("K%d",Sheetidx),tr["生产单位"])
                xlsx.SetCellValue(xlsxsheet,fmt.Sprintf("L%d",Sheetidx),tr["生产单位联系人"])
                xlsx.SetCellValue(xlsxsheet,fmt.Sprintf("M%d",Sheetidx),tr["生产单位电话"])
                xlsx.SetCellValue(xlsxsheet,fmt.Sprintf("N%d",Sheetidx),tr["商标"])
                xlsx.SetCellValue(xlsxsheet,fmt.Sprintf("O%d",Sheetidx),tr["样品名称br"])
                xlsx.SetCellValue(xlsxsheet,fmt.Sprintf("P%d",Sheetidx),tr["生产日期"])
                xlsx.SetCellValue(xlsxsheet,fmt.Sprintf("Q%d",Sheetidx),tr["保质期"])
                xlsx.SetCellValue(xlsxsheet,fmt.Sprintf("R%d",Sheetidx),tr["生产批号"])
                xlsx.SetCellValue(xlsxsheet,fmt.Sprintf("S%d",Sheetidx),tr["规格型号"])
                xlsx.SetCellValue(xlsxsheet,fmt.Sprintf("T%d",Sheetidx),tr["样品等级"])
                xlsx.SetCellValue(xlsxsheet,fmt.Sprintf("U%d",Sheetidx),tr["抽到样日期"])
                xlsx.SetCellValue(xlsxsheet,fmt.Sprintf("V%d",Sheetidx),tr["抽样方式"])
                xlsx.SetCellValue(xlsxsheet,fmt.Sprintf("W%d",Sheetidx),tr["样品状态"])
                xlsx.SetCellValue(xlsxsheet,fmt.Sprintf("X%d",Sheetidx),tr["样品状态2"])
                xlsx.SetCellValue(xlsxsheet,fmt.Sprintf("Y%d",Sheetidx),tr["保存条件"])
                xlsx.SetCellValue(xlsxsheet,fmt.Sprintf("Z%d",Sheetidx),tr["抽样基数"])
                xlsx.SetCellValue(xlsxsheet,fmt.Sprintf("AA%d",Sheetidx),tr["样品数"])
                xlsx.SetCellValue(xlsxsheet,fmt.Sprintf("AB%d",Sheetidx),tr["检验依据"])
                xlsx.SetCellValue(xlsxsheet,fmt.Sprintf("AC%d",Sheetidx),tr["备注"])
                vcl.ThreadSync(func() {
                    f.Gauge1.SetProgress(f.Gauge1.Progress()+1)
                })
            }
            xlsx.SaveAs("./导出.xlsx")
            vcl.ThreadSync(func() {
                vcl.ShowMessage("导出.xlsx 已保存")
            })
            return nil
        }()
        if err!=nil{
            vcl.ThreadSync(func() {
                vcl.ShowMessage(err.Error())
            })
        }
    }()
}
