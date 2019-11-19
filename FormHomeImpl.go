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
	"github.com/a97077088/threadpool"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"regexp"
	"sync"
	"sync/atomic"
	"time"
)

//::private::
type TFormHomeFields struct {
	sample_uuid  string
	sample_type  string
	sample_chs   []*nifdc.Channel
	sample_ds    *nifdc.Download_Data_r
	sample_ds_lk sync.Mutex
	sample_init  bool
	sample_ck    string

	uploaddatas    []*nifdc.UploadData
	uploaddatas_lk sync.Mutex

	test_platform_init bool
	test_platform_ck   string
}

func (f *TFormHome) OnFormClose(sender vcl.IObject, action *types.TCloseAction) {
	vcl.Application.Terminate()
}
func (f *TFormHome) OnFormCreate(sender vcl.IObject) {
	FormHome.SetShowInTaskBar(types.StAlways)
	f.Cbbt1s2.SetItemIndex(0)
	f.Cbbt2s1.SetItemIndex(0)
	f.Cbbt1s3.SetItemIndex(0)
	f.Dtpt1s1.SetDate(time.Now().AddDate(0, 0, -1))
	f.Dtpt1s2.SetDate(time.Now())
}
func (f *TFormHome) OnTss1Show(sender vcl.IObject) {
	err := func() error {
		var err error
		if f.sample_init == true {
			return nil
		}
		tmpck := ""
		f.sample_uuid, f.sample_chs, tmpck, err = nifdc.Sample_login(ck, nil)
		if err != nil {
			return err
		}
		f.sample_ck = fmt.Sprintf("%s;%s", ck, tmpck)
		f.Cbbt1s1.Items().Clear()
		for _, ch := range f.sample_chs {
			f.Cbbt1s1.Items().Add(ch.Name)
		}
		f.Cbbt1s1.SetItemIndex(0)
		f.sample_init = true
		return nil
	}()
	if err != nil {
		vcl.ShowMessage(err.Error())
	}
}
func (f *TFormHome) OnFormShow(sender vcl.IObject) {
	//fmt.Println(ck)
	f.SetCaption(fmt.Sprintf("数据同步组件 当前账号:%s ", user))
}
func (f *TFormHome) OnListView1Data(sender vcl.IObject, item *vcl.TListItem) {
	f.sample_ds_lk.Lock()
	defer f.sample_ds_lk.Unlock()
	idx := item.Index()
	if len(f.sample_ds.Data) < int(idx) {
		return
	}
	d := f.sample_ds.Data[idx]
	item.SetCaption(fmt.Sprintf("%d", idx+1))
	sitem := item.SubItems()
	sitem.Add(d.Update_time)
	sitem.Add(d.Sample_code)
	sitem.Add(d.New_sample_name)
	sitem.Add(d.Sp_d_38)
	sitem.Add(d.Resource_org_name)
	sitem.Add(d.Check_user_name)
}
func (f *TFormHome) OnButtonp1s1Click(sender vcl.IObject) {
	state := 0
	if f.Cbbt1s2.ItemIndex() == 0 {
		state = 4
	}
	if f.Cbbt1s2.ItemIndex() == 1 {
		state = 5
	}
	if f.Cbbt1s2.ItemIndex()==2{
		state = 12
	}
	f.Buttonp1s1.SetEnabled(false)
	go func() {
		defer vcl.ThreadSync(func() {
			f.Buttonp1s1.SetEnabled(true)
		})
		err := func() error {
			if r == false {
				return nil
			}
			var err error
			if len(f.sample_chs) != 0 {
				ch := f.sample_chs[f.Cbbt1s1.ItemIndex()].Type
				f.sample_ck, err = nifdc.Sample_switchchannel(f.sample_uuid, ch, ck, nil)
				if err != nil {
					return err
				}
			} else {
				tmpck := ""
				f.sample_uuid, f.sample_chs, tmpck, err = nifdc.Sample_login(ck, nil)
				if err != nil {
					return err
				}
				f.sample_ck = fmt.Sprintf("%s;%s", ck, tmpck)
			}

			sd := f.Dtpt1s1.Date().Format("2006-01-02")
			ed := f.Dtpt1s2.Date().Format("2006-01-02")
			tmpds, err := nifdc.DownData(state, sd, ed, f.sample_ck, nil)
			if err != nil {
				return err
			}
			vcl.ThreadSync(func() {
				f.sample_ds_lk.Lock()
				f.sample_ds = tmpds
				f.sample_ds_lk.Unlock()

				f.ListView1.Items().SetCount(int32(len(f.sample_ds.Data)))
			})
			return nil
		}()
		if err != nil {
			vcl.ThreadSyncVcl(func() {
				vcl.ShowMessage(err.Error())
			})
			return
		}
	}()
}

//导出检验完成全部
func (f *TFormHome) Exportjianyanwancheng_full(thread int, data []*nifdc.Data_o, fname string) error {
	xlsxsheet := "Sheet1"
	Sheetidx := 1
	xlsx := excelize.NewFile()
	xlsx.SetCellValue(xlsxsheet, "A1", "任务来源")
	xlsx.SetCellValue(xlsxsheet, "B1", "报送分类")
	xlsx.SetCellValue(xlsxsheet, "C1", "检验机构名称")
	xlsx.SetCellValue(xlsxsheet, "D1", "部署机构")
	xlsx.SetCellValue(xlsxsheet, "E1", "抽样类型")
	xlsx.SetCellValue(xlsxsheet, "F1", "抽样环节")
	xlsx.SetCellValue(xlsxsheet, "G1", "抽样地点")
	xlsx.SetCellValue(xlsxsheet, "H1", "食品分类")
	xlsx.SetCellValue(xlsxsheet, "I1", "抽样单编号")
	xlsx.SetCellValue(xlsxsheet, "J1", "检验目的/任务类别")
	xlsx.SetCellValue(xlsxsheet, "K1", "单位名称")
	xlsx.SetCellValue(xlsxsheet, "L1", "单位地址")
	xlsx.SetCellValue(xlsxsheet, "M1", "所在省份")
	xlsx.SetCellValue(xlsxsheet, "N1", "抽样人员")
	xlsx.SetCellValue(xlsxsheet, "O1", "联系人")
	xlsx.SetCellValue(xlsxsheet, "P1", "电子邮箱")
	xlsx.SetCellValue(xlsxsheet, "Q1", "电话")
	xlsx.SetCellValue(xlsxsheet, "R1", "传真")
	xlsx.SetCellValue(xlsxsheet, "S1", "邮编")
	xlsx.SetCellValue(xlsxsheet, "T1", "所在地")
	xlsx.SetCellValue(xlsxsheet, "U1", "区域类型")
	xlsx.SetCellValue(xlsxsheet, "V1", "单位名称")
	xlsx.SetCellValue(xlsxsheet, "W1", "单位地址")
	xlsx.SetCellValue(xlsxsheet, "X1", "营业执照/社会信用代码")
	xlsx.SetCellValue(xlsxsheet, "Y1", "许可证类型")
	xlsx.SetCellValue(xlsxsheet, "Z1", "经营许可证号")
	xlsx.SetCellValue(xlsxsheet, "AA1", "年销售额")
	xlsx.SetCellValue(xlsxsheet, "AB1", "单位法人")
	xlsx.SetCellValue(xlsxsheet, "AC1", "联系人")
	xlsx.SetCellValue(xlsxsheet, "AD1", "电话")
	xlsx.SetCellValue(xlsxsheet, "AE1", "传真")
	xlsx.SetCellValue(xlsxsheet, "AF1", "邮编")
	xlsx.SetCellValue(xlsxsheet, "AG1", "摊位号或姓名")
	xlsx.SetCellValue(xlsxsheet, "AH1", "身份证号")
	xlsx.SetCellValue(xlsxsheet, "AI1", "所在地")
	xlsx.SetCellValue(xlsxsheet, "AJ1", "企业地址")
	xlsx.SetCellValue(xlsxsheet, "AK1", "企业名称")
	xlsx.SetCellValue(xlsxsheet, "AL1", "生产许可证编号")
	xlsx.SetCellValue(xlsxsheet, "AM1", "生产单位联系人")
	xlsx.SetCellValue(xlsxsheet, "AN1", "生产单位电话")
	xlsx.SetCellValue(xlsxsheet, "AO1", "是否存在第三方企业信息")
	xlsx.SetCellValue(xlsxsheet, "AP1", "第三方企业省份")
	xlsx.SetCellValue(xlsxsheet, "AQ1", "第三方企业市区")
	xlsx.SetCellValue(xlsxsheet, "AR1", "第三方企业县区")
	xlsx.SetCellValue(xlsxsheet, "AS1", "第三方企业地址")
	xlsx.SetCellValue(xlsxsheet, "AT1", "第三方企业名称")
	xlsx.SetCellValue(xlsxsheet, "AU1", "第三方企业许可证编号")
	xlsx.SetCellValue(xlsxsheet, "AV1", "第三方企业联系人")
	xlsx.SetCellValue(xlsxsheet, "AW1", "第三方企业电话")
	xlsx.SetCellValue(xlsxsheet, "AX1", "第三方企业性质")
	xlsx.SetCellValue(xlsxsheet, "AY1", "样品条码")
	xlsx.SetCellValue(xlsxsheet, "AZ1", "样品商标")
	xlsx.SetCellValue(xlsxsheet, "BA1", "样品类型")
	xlsx.SetCellValue(xlsxsheet, "BB1", "样品来源")
	xlsx.SetCellValue(xlsxsheet, "BC1", "样品属性")
	xlsx.SetCellValue(xlsxsheet, "BD1", "包装分类")
	xlsx.SetCellValue(xlsxsheet, "BE1", "样品名称")
	xlsx.SetCellValue(xlsxsheet, "BF1", "购进日期")
	xlsx.SetCellValue(xlsxsheet, "BG1", "保质期")
	xlsx.SetCellValue(xlsxsheet, "BH1", "样品批号")
	xlsx.SetCellValue(xlsxsheet, "BI1", "规格型号")
	xlsx.SetCellValue(xlsxsheet, "BJ1", "质量等级")
	xlsx.SetCellValue(xlsxsheet, "BK1", "单价")
	xlsx.SetCellValue(xlsxsheet, "BL1", "是否进口")
	xlsx.SetCellValue(xlsxsheet, "BM1", "原产地")
	xlsx.SetCellValue(xlsxsheet, "BN1", "抽样日期")
	xlsx.SetCellValue(xlsxsheet, "BO1", "抽样方式")
	xlsx.SetCellValue(xlsxsheet, "BP1", "样品形态")
	xlsx.SetCellValue(xlsxsheet, "BQ1", "样品包装")
	xlsx.SetCellValue(xlsxsheet, "BR1", "抽样工具")
	xlsx.SetCellValue(xlsxsheet, "BS1", "抽样时样品储存条件")
	xlsx.SetCellValue(xlsxsheet, "BT1", "抽样基数")
	xlsx.SetCellValue(xlsxsheet, "BU1", "抽样数量")
	xlsx.SetCellValue(xlsxsheet, "BV1", "备样数量")
	xlsx.SetCellValue(xlsxsheet, "BW1", "抽样数量单位")
	xlsx.SetCellValue(xlsxsheet, "BX1", "执行标准/技术文件")
	xlsx.SetCellValue(xlsxsheet, "BY1", "备注")

	th := threadpool.NewThreadPool(thread, len(data))
	for idx, d := range data {
		fmt.Println(idx)
		_d := d
		th.Req(func() interface{} {
			itr, err := nettool.RNet_Call_1(&nettool.RNetOptions{}, func(source *addrmgr.AddrSource) (i interface{}, e error) {
				tr, err := nifdc.Viewrefusedsample_full(_d.Sample_code, f.sample_ck, nil)
				if err != nil {
					return nil, err
				}
				return tr, nil
			})
			if err != nil {
				fmt.Println(err)
			}
			vcl.ThreadSync(func() {
				Sheetidx++
				xlsx.SetRowHeight(xlsxsheet, Sheetidx, 30)
				tr := itr.(map[string]string)
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("A%d", Sheetidx), tr["抽样基础信息_任务来源"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("B%d", Sheetidx), tr["抽样基础信息_报送分类"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("C%d", Sheetidx), tr["抽样基础信息_检验机构名称"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("D%d", Sheetidx), tr["抽样基础信息_部署机构"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("E%d", Sheetidx), tr["抽样基础信息_抽样类型"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("F%d", Sheetidx), tr["抽样基础信息_抽样环节"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("G%d", Sheetidx), tr["抽样基础信息_抽样地点"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("H%d", Sheetidx), tr["抽样基础信息_食品分类"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("I%d", Sheetidx), tr["抽样基础信息_抽样单编号"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("J%d", Sheetidx), tr["抽样基础信息_检验目的/任务类别"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("K%d", Sheetidx), tr["抽样单位信息_单位名称"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("L%d", Sheetidx), tr["抽样单位信息_单位地址"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("M%d", Sheetidx), tr["抽样单位信息_所在省份"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("N%d", Sheetidx), tr["抽样单位信息_抽样人员"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("O%d", Sheetidx), tr["抽样单位信息_联系人"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("P%d", Sheetidx), tr["抽样单位信息_电子邮箱"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("Q%d", Sheetidx), tr["抽样单位信息_电话"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("R%d", Sheetidx), tr["抽样单位信息_传真"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("S%d", Sheetidx), tr["抽样单位信息_邮编"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("T%d", Sheetidx), tr["抽检场所信息_所在地"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("U%d", Sheetidx), tr["抽检场所信息_区域类型"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("V%d", Sheetidx), tr["抽检场所信息_单位名称"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("W%d", Sheetidx), tr["抽检场所信息_单位地址"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("X%d", Sheetidx), tr["抽检场所信息_营业执照/社会信用代码"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("Y%d", Sheetidx), tr["抽检场所信息_许可证类型"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("Z%d", Sheetidx), tr["抽检场所信息_经营许可证号"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AA%d", Sheetidx), tr["抽检场所信息_年销售额"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AB%d", Sheetidx), tr["抽检场所信息_单位法人"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AC%d", Sheetidx), tr["抽检场所信息_联系人"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AD%d", Sheetidx), tr["抽检场所信息_电话"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AE%d", Sheetidx), tr["抽检场所信息_传真"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AF%d", Sheetidx), tr["抽检场所信息_邮编"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AG%d", Sheetidx), tr["抽检场所信息_摊位号或姓名"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AH%d", Sheetidx), tr["抽检场所信息_身份证号"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AI%d", Sheetidx), tr["抽样生产企业信息_所在地"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AJ%d", Sheetidx), tr["抽样生产企业信息_企业地址"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AK%d", Sheetidx), tr["抽样生产企业信息_企业名称"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AL%d", Sheetidx), tr["抽样生产企业信息_生产许可证编号"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AM%d", Sheetidx), tr["抽样生产企业信息_生产单位联系人"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AN%d", Sheetidx), tr["抽样生产企业信息_生产单位电话"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AO%d", Sheetidx), tr["抽样生产企业信息_是否存在第三方企业信息"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AP%d", Sheetidx), tr["抽样生产企业信息_第三方企业省份"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AQ%d", Sheetidx), tr["抽样生产企业信息_第三方企业市区"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AR%d", Sheetidx), tr["抽样生产企业信息_第三方企业县区"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AS%d", Sheetidx), tr["抽样生产企业信息_第三方企业地址"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AT%d", Sheetidx), tr["抽样生产企业信息_第三方企业名称"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AU%d", Sheetidx), tr["抽样生产企业信息_第三方企业许可证编号"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AV%d", Sheetidx), tr["抽样生产企业信息_第三方企业联系人"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AW%d", Sheetidx), tr["抽样生产企业信息_第三方企业电话"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AX%d", Sheetidx), tr["抽样生产企业信息_第三方企业性质"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AY%d", Sheetidx), tr["抽检样品信息_样品条码"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AZ%d", Sheetidx), tr["抽检样品信息_样品商标"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BA%d", Sheetidx), tr["抽检样品信息_样品类型"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BB%d", Sheetidx), tr["抽检样品信息_样品来源"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BC%d", Sheetidx), tr["抽检样品信息_样品属性"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BD%d", Sheetidx), tr["抽检样品信息_包装分类"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BE%d", Sheetidx), tr["抽检样品信息_样品名称"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BF%d", Sheetidx), tr["抽检样品信息_购进日期"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BG%d", Sheetidx), tr["抽检样品信息_保质期"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BH%d", Sheetidx), tr["抽检样品信息_样品批号"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BI%d", Sheetidx), tr["抽检样品信息_规格型号"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BJ%d", Sheetidx), tr["抽检样品信息_质量等级"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BK%d", Sheetidx), tr["抽检样品信息_单价"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BL%d", Sheetidx), tr["抽检样品信息_是否进口"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BM%d", Sheetidx), tr["抽检样品信息_原产地"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BN%d", Sheetidx), tr["抽检样品信息_抽样日期"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BO%d", Sheetidx), tr["抽检样品信息_抽样方式"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BP%d", Sheetidx), tr["抽检样品信息_样品形态"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BQ%d", Sheetidx), tr["抽检样品信息_样品包装"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BR%d", Sheetidx), tr["抽检样品信息_抽样工具"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BS%d", Sheetidx), tr["抽检样品信息_抽样时样品储存条件"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BT%d", Sheetidx), tr["抽检样品信息_抽样基数"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BU%d", Sheetidx), tr["抽检样品信息_抽样数量"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BV%d", Sheetidx), tr["抽检样品信息_备样数量"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BW%d", Sheetidx), tr["抽检样品信息_抽样数量单位"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BX%d", Sheetidx), tr["抽检样品信息_执行标准/技术文件"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BY%d", Sheetidx), tr["抽检样品信息_备注"])
				f.Gauge1.SetProgress(f.Gauge1.Progress() + 1)
			})
			return nil
		})

	}
	th.Wait()

	err := xlsx.SaveAs(fname)
	if err != nil {
		return err
	}
	return nil
}


//导出已接收
func (f *TFormHome) Exportyijieshou(thread int, data []*nifdc.Data_o, fname string) error {
	xlsxsheet := "Sheet1"
	Sheetidx := 1
	xlsx := excelize.NewFile()
	xlsx.SetCellValue(xlsxsheet, "A1", "任务来源")
	xlsx.SetCellValue(xlsxsheet, "B1", "报送分类")
	xlsx.SetCellValue(xlsxsheet, "C1", "检验机构名称")
	xlsx.SetCellValue(xlsxsheet, "D1", "部署机构")
	xlsx.SetCellValue(xlsxsheet, "E1", "抽样类型")
	xlsx.SetCellValue(xlsxsheet, "F1", "抽样环节")
	xlsx.SetCellValue(xlsxsheet, "G1", "抽样地点")
	xlsx.SetCellValue(xlsxsheet, "H1", "食品分类")
	xlsx.SetCellValue(xlsxsheet, "I1", "抽样单编号")
	xlsx.SetCellValue(xlsxsheet, "J1", "检验目的/任务类别")
	xlsx.SetCellValue(xlsxsheet, "K1", "单位名称")
	xlsx.SetCellValue(xlsxsheet, "L1", "单位地址")
	xlsx.SetCellValue(xlsxsheet, "M1", "所在省份")
	xlsx.SetCellValue(xlsxsheet, "N1", "抽样人员")
	xlsx.SetCellValue(xlsxsheet, "O1", "联系人")
	xlsx.SetCellValue(xlsxsheet, "P1", "电子邮箱")
	xlsx.SetCellValue(xlsxsheet, "Q1", "电话")
	xlsx.SetCellValue(xlsxsheet, "R1", "传真")
	xlsx.SetCellValue(xlsxsheet, "S1", "邮编")
	xlsx.SetCellValue(xlsxsheet, "T1", "所在地")
	xlsx.SetCellValue(xlsxsheet, "U1", "区域类型")
	xlsx.SetCellValue(xlsxsheet, "V1", "单位名称")
	xlsx.SetCellValue(xlsxsheet, "W1", "单位地址")
	xlsx.SetCellValue(xlsxsheet, "X1", "营业执照/社会信用代码")
	xlsx.SetCellValue(xlsxsheet, "Y1", "许可证类型")
	xlsx.SetCellValue(xlsxsheet, "Z1", "经营许可证号")
	xlsx.SetCellValue(xlsxsheet, "AA1", "年销售额")
	xlsx.SetCellValue(xlsxsheet, "AB1", "单位法人")
	xlsx.SetCellValue(xlsxsheet, "AC1", "联系人")
	xlsx.SetCellValue(xlsxsheet, "AD1", "电话")
	xlsx.SetCellValue(xlsxsheet, "AE1", "传真")
	xlsx.SetCellValue(xlsxsheet, "AF1", "邮编")
	xlsx.SetCellValue(xlsxsheet, "AG1", "摊位号或姓名")
	xlsx.SetCellValue(xlsxsheet, "AH1", "身份证号")
	xlsx.SetCellValue(xlsxsheet, "AI1", "所在地")
	xlsx.SetCellValue(xlsxsheet, "AJ1", "企业地址")
	xlsx.SetCellValue(xlsxsheet, "AK1", "企业名称")
	xlsx.SetCellValue(xlsxsheet, "AL1", "生产许可证编号")
	xlsx.SetCellValue(xlsxsheet, "AM1", "生产单位联系人")
	xlsx.SetCellValue(xlsxsheet, "AN1", "生产单位电话")
	xlsx.SetCellValue(xlsxsheet, "AO1", "是否存在第三方企业信息")
	xlsx.SetCellValue(xlsxsheet, "AP1", "第三方企业省份")
	xlsx.SetCellValue(xlsxsheet, "AQ1", "第三方企业市区")
	xlsx.SetCellValue(xlsxsheet, "AR1", "第三方企业县区")
	xlsx.SetCellValue(xlsxsheet, "AS1", "第三方企业地址")
	xlsx.SetCellValue(xlsxsheet, "AT1", "第三方企业名称")
	xlsx.SetCellValue(xlsxsheet, "AU1", "第三方企业许可证编号")
	xlsx.SetCellValue(xlsxsheet, "AV1", "第三方企业联系人")
	xlsx.SetCellValue(xlsxsheet, "AW1", "第三方企业电话")
	xlsx.SetCellValue(xlsxsheet, "AX1", "第三方企业性质")
	xlsx.SetCellValue(xlsxsheet, "AY1", "样品条码")
	xlsx.SetCellValue(xlsxsheet, "AZ1", "样品商标")
	xlsx.SetCellValue(xlsxsheet, "BA1", "样品类型")
	xlsx.SetCellValue(xlsxsheet, "BB1", "样品来源")
	xlsx.SetCellValue(xlsxsheet, "BC1", "样品属性")
	xlsx.SetCellValue(xlsxsheet, "BD1", "包装分类")
	xlsx.SetCellValue(xlsxsheet, "BE1", "样品名称")
	xlsx.SetCellValue(xlsxsheet, "BF1", "购进日期")
	xlsx.SetCellValue(xlsxsheet, "BG1", "保质期")
	xlsx.SetCellValue(xlsxsheet, "BH1", "样品批号")
	xlsx.SetCellValue(xlsxsheet, "BI1", "规格型号")
	xlsx.SetCellValue(xlsxsheet, "BJ1", "质量等级")
	xlsx.SetCellValue(xlsxsheet, "BK1", "单价")
	xlsx.SetCellValue(xlsxsheet, "BL1", "是否进口")
	xlsx.SetCellValue(xlsxsheet, "BM1", "原产地")
	xlsx.SetCellValue(xlsxsheet, "BN1", "抽样日期")
	xlsx.SetCellValue(xlsxsheet, "BO1", "抽样方式")
	xlsx.SetCellValue(xlsxsheet, "BP1", "样品形态")
	xlsx.SetCellValue(xlsxsheet, "BQ1", "样品包装")
	xlsx.SetCellValue(xlsxsheet, "BR1", "抽样工具")
	xlsx.SetCellValue(xlsxsheet, "BS1", "抽样时样品储存条件")
	xlsx.SetCellValue(xlsxsheet, "BT1", "抽样基数")
	xlsx.SetCellValue(xlsxsheet, "BU1", "抽样数量")
	xlsx.SetCellValue(xlsxsheet, "BV1", "备样数量")
	xlsx.SetCellValue(xlsxsheet, "BW1", "抽样数量单位")
	xlsx.SetCellValue(xlsxsheet, "BX1", "执行标准/技术文件")
	xlsx.SetCellValue(xlsxsheet, "BY1", "备注")

	th := threadpool.NewThreadPool(thread, len(data))
	for _, d := range data {
		_d := d
		th.Req(func() interface{} {
			itr, err := nettool.RNet_Call_1(&nettool.RNetOptions{}, func(source *addrmgr.AddrSource) (i interface{}, e error) {
				tr, err := nifdc.Viewcheckedsample_full(_d.Sample_code, f.sample_ck, nil)
				if err != nil {
					return nil, err
				}
				return tr, nil
			})
			if err != nil {
				fmt.Println(err)
			}
			vcl.ThreadSync(func() {
				Sheetidx++
				xlsx.SetRowHeight(xlsxsheet, Sheetidx, 30)
				tr := itr.(map[string]string)
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("A%d", Sheetidx), tr["抽样基础信息_任务来源"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("B%d", Sheetidx), tr["抽样基础信息_报送分类"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("C%d", Sheetidx), tr["抽样基础信息_检验机构名称"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("D%d", Sheetidx), tr["抽样基础信息_部署机构"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("E%d", Sheetidx), tr["抽样基础信息_抽样类型"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("F%d", Sheetidx), tr["抽样基础信息_抽样环节"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("G%d", Sheetidx), tr["抽样基础信息_抽样地点"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("H%d", Sheetidx), tr["抽样基础信息_食品分类"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("I%d", Sheetidx), tr["抽样基础信息_抽样单编号"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("J%d", Sheetidx), tr["抽样基础信息_检验目的/任务类别"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("K%d", Sheetidx), tr["抽样单位信息_单位名称"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("L%d", Sheetidx), tr["抽样单位信息_单位地址"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("M%d", Sheetidx), tr["抽样单位信息_所在省份"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("N%d", Sheetidx), tr["抽样单位信息_抽样人员"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("O%d", Sheetidx), tr["抽样单位信息_联系人"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("P%d", Sheetidx), tr["抽样单位信息_电子邮箱"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("Q%d", Sheetidx), tr["抽样单位信息_电话"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("R%d", Sheetidx), tr["抽样单位信息_传真"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("S%d", Sheetidx), tr["抽样单位信息_邮编"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("T%d", Sheetidx), tr["抽检场所信息_所在地"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("U%d", Sheetidx), tr["抽检场所信息_区域类型"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("V%d", Sheetidx), tr["抽检场所信息_单位名称"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("W%d", Sheetidx), tr["抽检场所信息_单位地址"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("X%d", Sheetidx), tr["抽检场所信息_营业执照/社会信用代码"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("Y%d", Sheetidx), tr["抽检场所信息_许可证类型"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("Z%d", Sheetidx), tr["抽检场所信息_经营许可证号"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AA%d", Sheetidx), tr["抽检场所信息_年销售额"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AB%d", Sheetidx), tr["抽检场所信息_单位法人"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AC%d", Sheetidx), tr["抽检场所信息_联系人"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AD%d", Sheetidx), tr["抽检场所信息_电话"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AE%d", Sheetidx), tr["抽检场所信息_传真"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AF%d", Sheetidx), tr["抽检场所信息_邮编"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AG%d", Sheetidx), tr["抽检场所信息_摊位号或姓名"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AH%d", Sheetidx), tr["抽检场所信息_身份证号"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AI%d", Sheetidx), tr["抽样生产企业信息_所在地"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AJ%d", Sheetidx), tr["抽样生产企业信息_企业地址"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AK%d", Sheetidx), tr["抽样生产企业信息_企业名称"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AL%d", Sheetidx), tr["抽样生产企业信息_生产许可证编号"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AM%d", Sheetidx), tr["抽样生产企业信息_生产单位联系人"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AN%d", Sheetidx), tr["抽样生产企业信息_生产单位电话"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AO%d", Sheetidx), tr["抽样生产企业信息_是否存在第三方企业信息"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AP%d", Sheetidx), tr["抽样生产企业信息_第三方企业省份"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AQ%d", Sheetidx), tr["抽样生产企业信息_第三方企业市区"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AR%d", Sheetidx), tr["抽样生产企业信息_第三方企业县区"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AS%d", Sheetidx), tr["抽样生产企业信息_第三方企业地址"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AT%d", Sheetidx), tr["抽样生产企业信息_第三方企业名称"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AU%d", Sheetidx), tr["抽样生产企业信息_第三方企业许可证编号"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AV%d", Sheetidx), tr["抽样生产企业信息_第三方企业联系人"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AW%d", Sheetidx), tr["抽样生产企业信息_第三方企业电话"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AX%d", Sheetidx), tr["抽样生产企业信息_第三方企业性质"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AY%d", Sheetidx), tr["抽检样品信息_样品条码"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AZ%d", Sheetidx), tr["抽检样品信息_样品商标"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BA%d", Sheetidx), tr["抽检样品信息_样品类型"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BB%d", Sheetidx), tr["抽检样品信息_样品来源"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BC%d", Sheetidx), tr["抽检样品信息_样品属性"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BD%d", Sheetidx), tr["抽检样品信息_包装分类"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BE%d", Sheetidx), tr["抽检样品信息_样品名称"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BF%d", Sheetidx), tr["抽检样品信息_购进日期"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BG%d", Sheetidx), tr["抽检样品信息_保质期"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BH%d", Sheetidx), tr["抽检样品信息_样品批号"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BI%d", Sheetidx), tr["抽检样品信息_规格型号"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BJ%d", Sheetidx), tr["抽检样品信息_质量等级"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BK%d", Sheetidx), tr["抽检样品信息_单价"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BL%d", Sheetidx), tr["抽检样品信息_是否进口"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BM%d", Sheetidx), tr["抽检样品信息_原产地"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BN%d", Sheetidx), tr["抽检样品信息_抽样日期"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BO%d", Sheetidx), tr["抽检样品信息_抽样方式"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BP%d", Sheetidx), tr["抽检样品信息_样品形态"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BQ%d", Sheetidx), tr["抽检样品信息_样品包装"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BR%d", Sheetidx), tr["抽检样品信息_抽样工具"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BS%d", Sheetidx), tr["抽检样品信息_抽样时样品储存条件"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BT%d", Sheetidx), tr["抽检样品信息_抽样基数"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BU%d", Sheetidx), tr["抽检样品信息_抽样数量"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BV%d", Sheetidx), tr["抽检样品信息_备样数量"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BW%d", Sheetidx), tr["抽检样品信息_抽样数量单位"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BX%d", Sheetidx), tr["抽检样品信息_执行标准/技术文件"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BY%d", Sheetidx), tr["抽检样品信息_备注"])
				f.Gauge1.SetProgress(f.Gauge1.Progress() + 1)
			})
			return nil
		})

	}
	th.Wait()

	err := xlsx.SaveAs(fname)
	if err != nil {
		return err
	}
	return nil
}

//导出抽样完成全部
func (f *TFormHome) Exportchouyangwancheng_full(thread int, data []*nifdc.Data_o, fname string) error {
	xlsxsheet := "Sheet1"
	Sheetidx := 1
	xlsx := excelize.NewFile()
	xlsx.SetCellValue(xlsxsheet, "A1", "任务来源")
	xlsx.SetCellValue(xlsxsheet, "B1", "报送分类")
	xlsx.SetCellValue(xlsxsheet, "C1", "检验机构名称")
	xlsx.SetCellValue(xlsxsheet, "D1", "部署机构")
	xlsx.SetCellValue(xlsxsheet, "E1", "抽样类型")
	xlsx.SetCellValue(xlsxsheet, "F1", "抽样环节")
	xlsx.SetCellValue(xlsxsheet, "G1", "抽样地点")
	xlsx.SetCellValue(xlsxsheet, "H1", "食品分类")
	xlsx.SetCellValue(xlsxsheet, "I1", "抽样单编号")
	xlsx.SetCellValue(xlsxsheet, "J1", "检验目的/任务类别")
	xlsx.SetCellValue(xlsxsheet, "K1", "单位名称")
	xlsx.SetCellValue(xlsxsheet, "L1", "单位地址")
	xlsx.SetCellValue(xlsxsheet, "M1", "所在省份")
	xlsx.SetCellValue(xlsxsheet, "N1", "抽样人员")
	xlsx.SetCellValue(xlsxsheet, "O1", "联系人")
	xlsx.SetCellValue(xlsxsheet, "P1", "电子邮箱")
	xlsx.SetCellValue(xlsxsheet, "Q1", "电话")
	xlsx.SetCellValue(xlsxsheet, "R1", "传真")
	xlsx.SetCellValue(xlsxsheet, "S1", "邮编")
	xlsx.SetCellValue(xlsxsheet, "T1", "所在地")
	xlsx.SetCellValue(xlsxsheet, "U1", "区域类型")
	xlsx.SetCellValue(xlsxsheet, "V1", "单位名称")
	xlsx.SetCellValue(xlsxsheet, "W1", "单位地址")
	xlsx.SetCellValue(xlsxsheet, "X1", "营业执照/社会信用代码")
	xlsx.SetCellValue(xlsxsheet, "Y1", "许可证类型")
	xlsx.SetCellValue(xlsxsheet, "Z1", "经营许可证号")
	xlsx.SetCellValue(xlsxsheet, "AA1", "年销售额")
	xlsx.SetCellValue(xlsxsheet, "AB1", "单位法人")
	xlsx.SetCellValue(xlsxsheet, "AC1", "联系人")
	xlsx.SetCellValue(xlsxsheet, "AD1", "电话")
	xlsx.SetCellValue(xlsxsheet, "AE1", "传真")
	xlsx.SetCellValue(xlsxsheet, "AF1", "邮编")
	xlsx.SetCellValue(xlsxsheet, "AG1", "摊位号或姓名")
	xlsx.SetCellValue(xlsxsheet, "AH1", "身份证号")
	xlsx.SetCellValue(xlsxsheet, "AI1", "所在地")
	xlsx.SetCellValue(xlsxsheet, "AJ1", "企业地址")
	xlsx.SetCellValue(xlsxsheet, "AK1", "企业名称")
	xlsx.SetCellValue(xlsxsheet, "AL1", "生产许可证编号")
	xlsx.SetCellValue(xlsxsheet, "AM1", "生产单位联系人")
	xlsx.SetCellValue(xlsxsheet, "AN1", "生产单位电话")
	xlsx.SetCellValue(xlsxsheet, "AO1", "是否存在第三方企业信息")
	xlsx.SetCellValue(xlsxsheet, "AP1", "第三方企业省份")
	xlsx.SetCellValue(xlsxsheet, "AQ1", "第三方企业市区")
	xlsx.SetCellValue(xlsxsheet, "AR1", "第三方企业县区")
	xlsx.SetCellValue(xlsxsheet, "AS1", "第三方企业地址")
	xlsx.SetCellValue(xlsxsheet, "AT1", "第三方企业名称")
	xlsx.SetCellValue(xlsxsheet, "AU1", "第三方企业许可证编号")
	xlsx.SetCellValue(xlsxsheet, "AV1", "第三方企业联系人")
	xlsx.SetCellValue(xlsxsheet, "AW1", "第三方企业电话")
	xlsx.SetCellValue(xlsxsheet, "AX1", "第三方企业性质")
	xlsx.SetCellValue(xlsxsheet, "AY1", "样品条码")
	xlsx.SetCellValue(xlsxsheet, "AZ1", "样品商标")
	xlsx.SetCellValue(xlsxsheet, "BA1", "样品类型")
	xlsx.SetCellValue(xlsxsheet, "BB1", "样品来源")
	xlsx.SetCellValue(xlsxsheet, "BC1", "样品属性")
	xlsx.SetCellValue(xlsxsheet, "BD1", "包装分类")
	xlsx.SetCellValue(xlsxsheet, "BE1", "样品名称")
	xlsx.SetCellValue(xlsxsheet, "BF1", "购进日期")
	xlsx.SetCellValue(xlsxsheet, "BG1", "保质期")
	xlsx.SetCellValue(xlsxsheet, "BH1", "样品批号")
	xlsx.SetCellValue(xlsxsheet, "BI1", "规格型号")
	xlsx.SetCellValue(xlsxsheet, "BJ1", "质量等级")
	xlsx.SetCellValue(xlsxsheet, "BK1", "单价")
	xlsx.SetCellValue(xlsxsheet, "BL1", "是否进口")
	xlsx.SetCellValue(xlsxsheet, "BM1", "原产地")
	xlsx.SetCellValue(xlsxsheet, "BN1", "抽样日期")
	xlsx.SetCellValue(xlsxsheet, "BO1", "抽样方式")
	xlsx.SetCellValue(xlsxsheet, "BP1", "样品形态")
	xlsx.SetCellValue(xlsxsheet, "BQ1", "样品包装")
	xlsx.SetCellValue(xlsxsheet, "BR1", "抽样工具")
	xlsx.SetCellValue(xlsxsheet, "BS1", "抽样时样品储存条件")
	xlsx.SetCellValue(xlsxsheet, "BT1", "抽样基数")
	xlsx.SetCellValue(xlsxsheet, "BU1", "抽样数量")
	xlsx.SetCellValue(xlsxsheet, "BV1", "备样数量")
	xlsx.SetCellValue(xlsxsheet, "BW1", "抽样数量单位")
	xlsx.SetCellValue(xlsxsheet, "BX1", "执行标准/技术文件")
	xlsx.SetCellValue(xlsxsheet, "BY1", "备注")

	th := threadpool.NewThreadPool(thread, len(data))
	for idx, d := range data {
		fmt.Println(idx)
		_d := d
		th.Req(func() interface{} {
			itr, err := nettool.RNet_Call_1(&nettool.RNetOptions{}, func(source *addrmgr.AddrSource) (i interface{}, e error) {
				tr, err := nifdc.Viewnormalsample_full(_d.Sample_code, f.sample_ck, nil)
				if err != nil {
					return nil, err
				}
				return tr, nil
			})
			if err != nil {
				fmt.Println(err)
			}
			vcl.ThreadSync(func() {
				Sheetidx++
				xlsx.SetRowHeight(xlsxsheet, Sheetidx, 30)
				tr := itr.(map[string]string)
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("A%d", Sheetidx), tr["抽样基础信息_任务来源"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("B%d", Sheetidx), tr["抽样基础信息_报送分类"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("C%d", Sheetidx), tr["抽样基础信息_检验机构名称"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("D%d", Sheetidx), tr["抽样基础信息_部署机构"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("E%d", Sheetidx), tr["抽样基础信息_抽样类型"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("F%d", Sheetidx), tr["抽样基础信息_抽样环节"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("G%d", Sheetidx), tr["抽样基础信息_抽样地点"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("H%d", Sheetidx), tr["抽样基础信息_食品分类"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("I%d", Sheetidx), tr["抽样基础信息_抽样单编号"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("J%d", Sheetidx), tr["抽样基础信息_检验目的/任务类别"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("K%d", Sheetidx), tr["抽样单位信息_单位名称"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("L%d", Sheetidx), tr["抽样单位信息_单位地址"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("M%d", Sheetidx), tr["抽样单位信息_所在省份"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("N%d", Sheetidx), tr["抽样单位信息_抽样人员"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("O%d", Sheetidx), tr["抽样单位信息_联系人"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("P%d", Sheetidx), tr["抽样单位信息_电子邮箱"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("Q%d", Sheetidx), tr["抽样单位信息_电话"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("R%d", Sheetidx), tr["抽样单位信息_传真"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("S%d", Sheetidx), tr["抽样单位信息_邮编"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("T%d", Sheetidx), tr["抽检场所信息_所在地"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("U%d", Sheetidx), tr["抽检场所信息_区域类型"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("V%d", Sheetidx), tr["抽检场所信息_单位名称"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("W%d", Sheetidx), tr["抽检场所信息_单位地址"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("X%d", Sheetidx), tr["抽检场所信息_营业执照/社会信用代码"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("Y%d", Sheetidx), tr["抽检场所信息_许可证类型"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("Z%d", Sheetidx), tr["抽检场所信息_经营许可证号"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AA%d", Sheetidx), tr["抽检场所信息_年销售额"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AB%d", Sheetidx), tr["抽检场所信息_单位法人"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AC%d", Sheetidx), tr["抽检场所信息_联系人"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AD%d", Sheetidx), tr["抽检场所信息_电话"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AE%d", Sheetidx), tr["抽检场所信息_传真"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AF%d", Sheetidx), tr["抽检场所信息_邮编"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AG%d", Sheetidx), tr["抽检场所信息_摊位号或姓名"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AH%d", Sheetidx), tr["抽检场所信息_身份证号"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AI%d", Sheetidx), tr["抽样生产企业信息_所在地"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AJ%d", Sheetidx), tr["抽样生产企业信息_企业地址"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AK%d", Sheetidx), tr["抽样生产企业信息_企业名称"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AL%d", Sheetidx), tr["抽样生产企业信息_生产许可证编号"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AM%d", Sheetidx), tr["抽样生产企业信息_生产单位联系人"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AN%d", Sheetidx), tr["抽样生产企业信息_生产单位电话"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AO%d", Sheetidx), tr["抽样生产企业信息_是否存在第三方企业信息"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AP%d", Sheetidx), tr["抽样生产企业信息_第三方企业省份"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AQ%d", Sheetidx), tr["抽样生产企业信息_第三方企业市区"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AR%d", Sheetidx), tr["抽样生产企业信息_第三方企业县区"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AS%d", Sheetidx), tr["抽样生产企业信息_第三方企业地址"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AT%d", Sheetidx), tr["抽样生产企业信息_第三方企业名称"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AU%d", Sheetidx), tr["抽样生产企业信息_第三方企业许可证编号"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AV%d", Sheetidx), tr["抽样生产企业信息_第三方企业联系人"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AW%d", Sheetidx), tr["抽样生产企业信息_第三方企业电话"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AX%d", Sheetidx), tr["抽样生产企业信息_第三方企业性质"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AY%d", Sheetidx), tr["抽检样品信息_样品条码"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AZ%d", Sheetidx), tr["抽检样品信息_样品商标"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BA%d", Sheetidx), tr["抽检样品信息_样品类型"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BB%d", Sheetidx), tr["抽检样品信息_样品来源"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BC%d", Sheetidx), tr["抽检样品信息_样品属性"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BD%d", Sheetidx), tr["抽检样品信息_包装分类"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BE%d", Sheetidx), tr["抽检样品信息_样品名称"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BF%d", Sheetidx), tr["抽检样品信息_购进日期"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BG%d", Sheetidx), tr["抽检样品信息_保质期"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BH%d", Sheetidx), tr["抽检样品信息_样品批号"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BI%d", Sheetidx), tr["抽检样品信息_规格型号"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BJ%d", Sheetidx), tr["抽检样品信息_质量等级"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BK%d", Sheetidx), tr["抽检样品信息_单价"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BL%d", Sheetidx), tr["抽检样品信息_是否进口"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BM%d", Sheetidx), tr["抽检样品信息_原产地"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BN%d", Sheetidx), tr["抽检样品信息_抽样日期"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BO%d", Sheetidx), tr["抽检样品信息_抽样方式"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BP%d", Sheetidx), tr["抽检样品信息_样品形态"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BQ%d", Sheetidx), tr["抽检样品信息_样品包装"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BR%d", Sheetidx), tr["抽检样品信息_抽样工具"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BS%d", Sheetidx), tr["抽检样品信息_抽样时样品储存条件"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BT%d", Sheetidx), tr["抽检样品信息_抽样基数"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BU%d", Sheetidx), tr["抽检样品信息_抽样数量"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BV%d", Sheetidx), tr["抽检样品信息_备样数量"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BW%d", Sheetidx), tr["抽检样品信息_抽样数量单位"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BX%d", Sheetidx), tr["抽检样品信息_执行标准/技术文件"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("BY%d", Sheetidx), tr["抽检样品信息_备注"])
				f.Gauge1.SetProgress(f.Gauge1.Progress() + 1)
			})
			return nil
		})

	}
	th.Wait()

	err := xlsx.SaveAs(fname)
	if err != nil {
		return err
	}
	return nil
}

//导出抽样完成一半
func (f *TFormHome) Exportchouyangwancheng_half(thread int, data []*nifdc.Data_o, fname string) error {
	xlsxsheet := "Sheet1"
	Sheetidx := 1
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

	th := threadpool.NewThreadPool(thread, len(data))
	for _, d := range data {
		_d := d
		th.Req(func() interface{} {
			itr, err := nettool.RNet_Call_1(&nettool.RNetOptions{}, func(source *addrmgr.AddrSource) (i interface{}, e error) {
				tr, err := nifdc.Viewnormalsample_mode1(_d.Sample_code, f.sample_ck, nil)
				if err != nil {
					return nil, err
				}
				return tr, nil
			})
			if err != nil {
				fmt.Println(err)
			}
			vcl.ThreadSync(func() {
				Sheetidx++
				xlsx.SetRowHeight(xlsxsheet, Sheetidx, 30)
				tr := itr.(map[string]string)
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("A%d", Sheetidx), tr["委托单位"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("B%d", Sheetidx), tr["抽样地点"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("C%d", Sheetidx), tr["抽样单号"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("D%d", Sheetidx), tr["检验类型"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("E%d", Sheetidx), tr["抽送样人"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("F%d", Sheetidx), tr["受检单位"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("G%d", Sheetidx), tr["地址"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("H%d", Sheetidx), tr["联系人"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("I%d", Sheetidx), tr["电话"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("J%d", Sheetidx), tr["生产单位地址"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("K%d", Sheetidx), tr["生产单位"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("L%d", Sheetidx), tr["生产单位联系人"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("M%d", Sheetidx), tr["生产单位电话"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("N%d", Sheetidx), tr["商标"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("O%d", Sheetidx), tr["样品名称br"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("P%d", Sheetidx), tr["生产日期"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("Q%d", Sheetidx), tr["保质期"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("R%d", Sheetidx), tr["生产批号"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("S%d", Sheetidx), tr["规格型号"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("T%d", Sheetidx), tr["样品等级"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("U%d", Sheetidx), tr["抽到样日期"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("V%d", Sheetidx), tr["抽样方式"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("W%d", Sheetidx), tr["样品状态"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("X%d", Sheetidx), tr["样品状态2"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("Y%d", Sheetidx), tr["保存条件"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("Z%d", Sheetidx), tr["抽样基数"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AA%d", Sheetidx), tr["样品数"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AB%d", Sheetidx), tr["检验依据"])
				xlsx.SetCellValue(xlsxsheet, fmt.Sprintf("AC%d", Sheetidx), tr["备注"])
				f.Gauge1.SetProgress(f.Gauge1.Progress() + 1)
			})
			return nil
		})

	}
	th.Wait()

	err := xlsx.SaveAs(fname)
	if err != nil {
		return err
	}
	return nil
}
func (f *TFormHome) OnButtonp1s2Click(sender vcl.IObject) {

	sel := int(f.Cbbt1s2.ItemIndex())  //任务状态
	ssel := int(f.Cbbt1s3.ItemIndex()) //导出模式
	if f.SaveDialog1.Execute() == false {
		return
	}
	fname := f.SaveDialog1.FileName()

	f.sample_ds_lk.Lock()
	tmpds := f.sample_ds
	f.sample_ds_lk.Unlock()

	f.Buttonp1s2.SetEnabled(false)
	go func() {
		defer vcl.ThreadSync(func() {
			f.Buttonp1s2.SetEnabled(true)
		})
		err := func() error {
			if w == false {
				return nil
			}
			if tmpds == nil || len(tmpds.Data) == 0 {
				return errors.New("数据不能为空")
			}
			vcl.ThreadSync(func() {
				f.Gauge1.SetProgress(0)
				f.Gauge1.SetMaxValue(int32(len(tmpds.Data)))
			})
			if sel == 0 { //导出抽样完成
				if ssel == 0 { //导出全部字段
					err := f.Exportchouyangwancheng_full(thread, tmpds.Data, fname)
					if err != nil {
						return err
					}
				} else { //导出半字段
					err := f.Exportchouyangwancheng_half(thread, tmpds.Data, fname)
					if err != nil {
						return err
					}
				}
			}
			if sel == 1 { //导出
				err := f.Exportyijieshou(thread, tmpds.Data, fname)
				if err != nil {
					return err
				}
			}
			if sel == 2 { //导出
				err := f.Exportjianyanwancheng_full(thread, tmpds.Data, fname)
				if err != nil {
					return err
				}
			}

			vcl.ThreadSync(func() {
				vcl.ShowMessage(fmt.Sprintf("%s 已保存", fname))
			})

			return nil
		}()
		if err != nil {
			vcl.ThreadSyncVcl(func() {
				vcl.ShowMessage(err.Error())
			})
			return
		}
	}()
}
func (f *TFormHome) OnListView1Resize(sender vcl.IObject) {
	go vcl.ThreadSync(func() {
		lastitem := f.ListView1.Column(f.ListView1.Columns().Count() - 1)
		lastitem.SetWidth(lastitem.Width() - 10)
	})
}
func (f *TFormHome) OnTimer1Timer(sender vcl.IObject) {
	f.ListView1.Invalidate()
	f.ListView2.Invalidate()
}
func (f *TFormHome) GetUploadData(k string) *nifdc.UploadData {
	f.uploaddatas_lk.Lock()
	defer f.uploaddatas_lk.Unlock()
	var rit *nifdc.UploadData
	for _, it := range f.uploaddatas {
		if it.SEV("抽样单编号") == k {
			rit = it
			break
		}
	}
	return rit
}
func (f *TFormHome) GetUploadDataOrCreate(k string) *nifdc.UploadData {
	f.uploaddatas_lk.Lock()
	defer f.uploaddatas_lk.Unlock()
	var rit *nifdc.UploadData
	for _, it := range f.uploaddatas {
		if it.SEV("抽样单编号") == k {
			rit = it
			break
		}
	}
	if rit == nil {
		rit = &nifdc.UploadData{}
		rit.SSEV("抽样单编号", k)
		f.uploaddatas = append(f.uploaddatas, rit)
	}
	return rit
}
func (f *TFormHome) OnButtont2s1Click(sender vcl.IObject) {
	if f.DlgOpen1.Execute() == false {
		return
	}
	fname := f.DlgOpen1.FileName()
	f.Buttont2s1.SetEnabled(false)
	go func() {
		defer vcl.ThreadSync(func() {
			f.Buttont2s1.SetEnabled(true)
		})
		err := func() error {
			if r == false {
				return nil
			}
			f.uploaddatas_lk.Lock()
			f.uploaddatas = nil
			f.uploaddatas_lk.Unlock()
			xlsx, err := excelize.OpenFile(fname)
			if err != nil {
				return err
			}
			if xlsx.SheetCount == 0 {
				return errors.New("excel是空数据")
			}
			sheetname := xlsx.GetSheetMap()[1]
			rows := xlsx.GetRows(sheetname)
			for _, row := range rows {

				ok, err := regexp.MatchString(`NCP\d+|DC\d+`, row[0])
				if err != nil {
					return err
				}
				if ok == false {
					continue
				}

				d := f.GetUploadDataOrCreate(row[0])
				d.SSEV("样品匹配", "否")
				d.SSEV("抽样单编号", row[0])
				d.SSEV("报告书编号", row[1])
				d.SSEV("检测项目", "双击查看")
				d.SSEV("监督抽检报告备注", row[14])
				d.SSEV("风险监测报告备注", row[15])
				d.SSEV("上传状态", "否")
				d.SSEV("上传结果", "")
				d.SSEV("样品名称", row[16])

				d.AddSubitem(map[string]string{
					"检验项目":  row[2],
					"检验结果":  row[3],
					"结果单位":  row[4],
					"结果判定":  row[5],
					"检验依据":  row[6],
					"判定依据":  row[7],
					"最小允许限": row[8],
					"最大允许限": row[9],
					"允许限单位": row[10],
					"方法检出限": row[11],
					"检出限单位": row[12],
					"说明":    row[13],
				})
			}
			vcl.ThreadSync(func() {
				f.ListView2.Items().SetCount(int32(len(f.uploaddatas)))
			})
			return nil
		}()
		if err != nil {
			vcl.ThreadSync(func() {
				vcl.ShowMessage(err.Error())
			})
		}
	}()
}
func (f *TFormHome) OnListView2Data(sender vcl.IObject, item *vcl.TListItem) {
	f.uploaddatas_lk.Lock()
	defer f.uploaddatas_lk.Unlock()
	idx := item.Index()
	if len(f.uploaddatas) < int(idx) {
		return
	}
	d := f.uploaddatas[idx]
	item.SetCaption(fmt.Sprintf("%d", idx+1))
	sitem := item.SubItems()
	sitem.Add(d.SEV("样品名称"))
	sitem.Add(fmt.Sprintf("%d", len(d.Subitem())))
	sitem.Add(d.SEV("样品匹配"))
	sitem.Add(d.SEV("抽样单编号"))
	sitem.Add(d.SEV("报告书编号"))
	sitem.Add(d.SEV("检测项目"))
	sitem.Add(d.SEV("监督抽检报告备注"))
	sitem.Add(d.SEV("风险监测报告备注"))
	sitem.Add(d.SEV("上传状态"))
	sitem.Add(d.SEV("上传结果"))
}
func (f *TFormHome) OnListView2Resize(sender vcl.IObject) {
	go vcl.ThreadSync(func() {
		lastitem := f.ListView2.Column(f.ListView2.Columns().Count() - 1)
		lastitem.SetWidth(lastitem.Width() - 10)
	})
}
func (f *TFormHome) OnButtont2s2Click(sender vcl.IObject) {
	tp := 0
	if f.Cbbt2s1.ItemIndex() == 0 {
		tp = 0
	}
	if f.Cbbt2s1.ItemIndex() == 1 {
		tp = 1
	}

	f.Buttont2s2.SetEnabled(false)
	go func() {
		defer vcl.ThreadSync(func() {
			f.Buttont2s2.SetEnabled(true)
		})
		err := func() error {
			if r == false {
				return nil
			}
			enddate := time.Now()
			startdate := enddate.AddDate(-1, 0, 0)
			var dds *nifdc.Api_food_getFood_r
			var err error
			if tp == 0 {
				dds, err = nifdc.Test_platform_api_food_getFood(startdate.Format("2006-01-02"), enddate.Format("2006-01-02"), f.test_platform_ck, nil)
				if err != nil {
					return err
				}
			}
			if tp == 1 {
				dds, err = nifdc.Test_platform_api_agriculture_getAgriculture(startdate.Format("2006-01-02"), enddate.Format("2006-01-02"), f.test_platform_ck, nil)
				if err != nil {
					return err
				}
			}

			updpipei := 0
			updcount := dds.Total
			for _, d := range dds.Rows {
				upd := f.GetUploadData(d.Sp_s_16)
				if upd == nil {
					continue
				}
				updpipei++
				upd.Set_env_value("id", d.Id)
				upd.SSEV("样品匹配", "是")
			}

			vcl.ThreadSync(func() {
				f.Labelt2s2.SetCaption(fmt.Sprintf("匹配结果: 查询到%d条记录,匹配到%d条", updcount, updpipei))
			})
			return nil
		}()
		if err != nil {
			vcl.ThreadSync(func() {
				vcl.ShowMessage(err.Error())
			})
		}
	}()
}
func (f *TFormHome) OnTss2Show(sender vcl.IObject) {
	err := func() error {
		var err error
		if f.test_platform_init == true {
			return nil
		}
		f.test_platform_ck, err = nifdc.Test_platform_login(ck, nil)
		if err != nil {
			return err
		}
		f.test_platform_init = true
		return nil
	}()
	if err != nil {
		vcl.ShowMessage(err.Error())
	}
}
func (f *TFormHome) OnButtont2s3Click(sender vcl.IObject) {
	f.Buttont2s3.SetEnabled(false)
	go func() {
		defer vcl.ThreadSync(func() {
			f.Buttont2s3.SetEnabled(true)
		})
		err := func() error {
			if w == false {
				return nil
			}
			var tmpds []*nifdc.UploadData
			f.uploaddatas_lk.Lock()
			tmpds = f.uploaddatas
			f.uploaddatas_lk.Unlock()

			vcl.ThreadSync(func() {
				f.Gauge2.SetProgress(0)
				f.Gauge2.SetMaxValue(int32(len(tmpds)))
			})
			var nok, nerr int32
			th := threadpool.NewThreadPool(thread, len(tmpds))
			for _, td := range tmpds {
				_td := td
				th.Req(func() interface{} {
					err := func() error {
						if _td.SEV("样品匹配") == "否" {
							return errors.New("没有匹配数据")
						}
						err := nettool.RNet_Call(nil, func(source *addrmgr.AddrSource) error {
							fddetail, err := nifdc.Test_platform_foodTest_foodDetail(td.Env_for_key("id").(int), f.test_platform_ck, nil)
							if err != nil {
								return err
							}
							testinfo, err := nifdc.Test_platform_api_food_getTestInfo(fddetail["sd"], f.test_platform_ck, nil)
							if err != nil {
								return err
							}
							subitem := _td.Subitem()
							unqualifieds := nifdc.Getunqualified(subitem)
							jielun := "纯抽检合格样品"
							baogaoleibie := "合格报告"
							if len(unqualifieds) != 0 {
								jielun = "纯抽检不合格样品"
								baogaoleibie = "一般不合格报告"
							}
							jiancejielun := nifdc.Buildbaogao(subitem)

							nifdc.Fill_item(map[string]string{
								"报告书编号":    _td.SEV("报告书编号"),
								"监督抽检报告备注": _td.SEV("监督抽检报告备注"),
								"风险监测报告备注": _td.SEV("风险监测报告备注"),
								"结论":       jielun,
								"报告类别":     baogaoleibie,
								"检验结论":     jiancejielun,
							}, fddetail)
							nifdc.Fill_subitem(subitem, testinfo.Rows)
							err = nifdc.Test_platform_api_food_save(fddetail, testinfo.Rows, f.test_platform_ck, nil)
							if err != nil {
								return err
							}

							atomic.AddInt32(&nok, 1)
							_td.SSEV("上传结果", "成功")
							return nil
						})
						if err != nil {
							return err
						}
						return nil
					}()
					vcl.ThreadSync(func() {
						f.Gauge2.SetProgress(f.Gauge2.Progress() + 1)
					})
					_td.SSEV("上传状态", "是")
					if err != nil {
						atomic.AddInt32(&nerr, 1)
						_td.SSEV("上传结果", err.Error())
					}
					return nil
				})
			}
			th.Wait()
			vcl.ThreadSync(func() {
				vcl.ShowMessage(fmt.Sprintf("成功:%d\n\n错误:%d", nok, nerr))
			})
			return nil
		}()
		if err != nil {
			vcl.ThreadSync(func() {
				vcl.ShowMessage(err.Error())
			})
		}
	}()
}
func (f *TFormHome) OnListView2DblClick(sender vcl.IObject) {
	if f.ListView2.Selected().IsValid() == false {
		return
	}
	var td *nifdc.UploadData
	idx := f.ListView2.ItemIndex()
	f.uploaddatas_lk.Lock()
	td = f.uploaddatas[idx]
	f.uploaddatas_lk.Unlock()
	Formjiance.Td = td
	Formjiance.ShowModal()
}

func (f *TFormHome) OnCbbt1s2Change(sender vcl.IObject) {
	f.Cbbt1s3.Clear()
	if f.Cbbt1s2.Text() == "抽样完成" {
		f.Cbbt1s3.Items().Add("全部导出")
		f.Cbbt1s3.Items().Add("模式1")
	} else if f.Cbbt1s2.Text() == "已接收" {
		f.Cbbt1s3.Items().Add("全部导出")
	}else if f.Cbbt1s2.Text() == "检验完成" {
		f.Cbbt1s3.Items().Add("全部导出")
	}
	f.Cbbt1s3.SetItemIndex(0)
}
