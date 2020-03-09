// 由res2go自动生成。
// 在这里写你的事件。

package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/a97077088/addrmgr"
	"github.com/a97077088/chinese-holidays-go/holidays"
	"github.com/a97077088/nettool"
	"github.com/a97077088/nifdc"
	"github.com/a97077088/threadpool"
	"github.com/tealeg/xlsx"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"io/ioutil"
	"net/url"
	"nfidccli/models"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"text/template"
	"time"
)

//::private::
type TFormHomeFields struct {
	sample_uuid string
	sample_type string

	sample_chs   []*nifdc.Channel
	sample_ds    []*nifdc.Data_o
	sample_ds_lk sync.Mutex
	sample_init  bool
	sample_ck    string

	uploaddatas    []*nifdc.UploadData
	uploaddatas_lk sync.Mutex

	getFood_ds    []*nifdc.Api_food_getFood_o
	getFood_ds_lk sync.Mutex

	test_platform_init bool
	test_platform_ck   string

	renwudapingtaisql_rule    [][]string
	jianyanjieguosql_rule    [][]string
	jianyanjieguoexcel_rule  [][]string
	renwudapingtaiexcel_rule [][]string
}

func (f *TFormHome) readrule(rulename string) ([][]string, error) {
	rrule := [][]string{}
	allbt, err := ioutil.ReadFile(rulename)
	if err != nil {
		return nil, err
	}
	spbf := strings.Split(string(allbt), "\n")
	for _, itbf := range spbf {
		sitbf := strings.TrimSpace(itbf)
		if sitbf != "" {
			spcbf := strings.Split(sitbf, "=")
			if len(spcbf) != 2 {
				continue
			}
			rrule = append(rrule, []string{
				spcbf[0],
				spcbf[1],
			})
		}
	}
	return rrule, nil
}

func (f *TFormHome) OnFormClose(sender vcl.IObject, action *types.TCloseAction) {
	vcl.Application.Terminate()
}
func (f *TFormHome) OnFormCreate(sender vcl.IObject) {
	FormHome.SetShowInTaskBar(types.StAlways)
	f.Cbbt1s2.SetItemIndex(0)
	f.Cbbt2s1.SetItemIndex(0)
	f.Cbbt3s1.SetItemIndex(0)
	f.Cbbt3s2.SetItemIndex(0)
	f.Cbbt3s3.SetItemIndex(0)
	f.Cbbt3s4.SetItemIndex(0)
	f.Dtpt1s1.SetDate(time.Now().AddDate(0, -1, 0))
	f.Dtpt1s2.SetDate(time.Now())

	f.Dtpt2s1.SetDate(time.Now().AddDate(0, -1, 0))
	f.Dtpt2s2.SetDate(time.Now())

	f.Dtpt3s1.SetDate(time.Now().AddDate(0, -1, 0))
	f.Dtpt3s2.SetDate(time.Now())
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

	f.SetCaption(fmt.Sprintf("数据同步组件 当前账号:%s ", user))
}
func (f *TFormHome) OnListView1Data(sender vcl.IObject, item *vcl.TListItem) {
	f.sample_ds_lk.Lock()
	defer f.sample_ds_lk.Unlock()
	idx := item.Index()
	if f.sample_ds == nil || len(f.sample_ds) < int(idx) {
		return
	}
	d := f.sample_ds[idx]
	item.SetCaption(fmt.Sprintf("%d", idx+1))
	sitem := item.SubItems()
	sitem.Add(d.Update_time)
	sitem.Add(d.Sample_code)
	sitem.Add(d.New_sample_name)
	sitem.Add(d.Sp_d_38)
	sitem.Add(d.Resource_org_name)
	sitem.Add(d.Check_user_name)
	sitem.Add(d.User.SEV("处理状态"))
	sitem.Add(d.User.SEV("处理结果"))
}
func (f *TFormHome) OnButtonp1s1Click(sender vcl.IObject) {
	state := 0
	if f.Cbbt1s2.ItemIndex() == 0 {
		state = 3
	} else if f.Cbbt1s2.ItemIndex() == 1 {
		state = 4
	} else if f.Cbbt1s2.ItemIndex() == 2 {
		state = 5
	} else if f.Cbbt1s2.ItemIndex() == 3 {
		state = 6
	} else if f.Cbbt1s2.ItemIndex() == 4 {
		state = 7
	} else if f.Cbbt1s2.ItemIndex() == 5 {
		state = 12
	}

	resource_org_id := f.Edtt1s1.Text()
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
			tmpds, err := nifdc.DownData(resource_org_id, state, sd, ed, f.sample_ck, nil)
			if err != nil {
				return err
			}
			vcl.ThreadSync(func() {
				f.sample_ds_lk.Lock()
				f.sample_ds = tmpds.Data
				f.sample_ds_lk.Unlock()

				f.ListView1.Items().SetCount(int32(len(f.sample_ds)))
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

//导出任务大平台excel
func (f *TFormHome) Exportrenwudapingtai_excel(thread int, data []*nifdc.Data_o, fname string) error {
	for _, d := range data {
		d.User.SSEV("处理状态", "")
		d.User.SSEV("处理结果", "")
	}
	nerr := int32(0)
	nok := int32(0)
	nrey := int32(0)
	xlsxf := xlsx.NewFile()
	sheet, err := xlsxf.AddSheet("数据报告")
	if err != nil {
		return err
	}
	sheet_lk := sync.Mutex{}
	row := sheet.AddRow()
	for _, it := range f.renwudapingtaiexcel_rule {
		dbk := it[0]
		cell := row.AddCell()
		cell.SetString(dbk)
	}
	sheet.SetColWidth(0, len(f.renwudapingtaiexcel_rule), 15)

	th := threadpool.NewThreadPool(thread, len(data))
	for _, d := range data {
		_d := d
		th.Req(func() interface{} {
			defer vcl.ThreadSync(func() {
				f.Gauge1.SetProgress(f.Gauge1.Progress() + 1)
				_d.User.SSEV("处理状态", "完成")
			})
			err := func() error {
				itr, err := nettool.RNet_Call_1(&nettool.RNetOptions{}, func(source *addrmgr.AddrSource) (i interface{}, e error) {
					tr, err := nifdc.Viewcheckedsample_full(_d.Sample_code, f.sample_ck, nil)
					if err != nil {
						return nil, err
					}
					return tr, nil
				})
				if err != nil {
					return err
				}
				tr := itr.(map[string]string)
				fmt.Println(tr)

				sheet_lk.Lock()
				row := sheet.AddRow()
				tmj := template.New("tmj")
				tmj.Funcs(map[string]interface{}{
					"replace": strings.ReplaceAll,
					"replaceex": replaceex,
				})
				for _, it := range f.renwudapingtaiexcel_rule {
					webk := it[1]

					_, err = tmj.Parse(webk)
					if err != nil {
						fmt.Println(err)
						return err
					}
					var tmpwebv bytes.Buffer
					err := tmj.Execute(&tmpwebv, tr)
					if err != nil {
						fmt.Println(err)
						return err
					}

					cl := row.AddCell()
					cl.SetString(tmpwebv.String())
				}
				sheet_lk.Unlock()

				atomic.AddInt32(&nok, 1)
				return nil
			}()
			if err != nil {
				atomic.AddInt32(&nerr, 1)
				_d.User.SSEV("处理结果", err.Error())
				return err
			}
			_d.User.SSEV("处理结果", "完成")
			return nil
		})

	}
	th.Wait()
	err = xlsxf.Save(fname)
	if err != nil {
		return err
	}
	vcl.ThreadSync(func() {
		vcl.ShowMessage(fmt.Sprintf("成功:%d\n\n失败:%d\n\n已存在:%d", atomic.LoadInt32(&nok), atomic.LoadInt32(&nerr), atomic.LoadInt32(&nrey)))
	})

	return nil
}
func (f *TFormHome) OnButtonp1s2Click(sender vcl.IObject) {
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
			if tmpds == nil || len(tmpds) == 0 {
				return errors.New("数据不能为空")
			}

			var err error
			f.renwudapingtaiexcel_rule, err = f.readrule("./下载任务大平台excel规则.txt")
			if err != nil {
				return err
			}

			vcl.ThreadSync(func() {
				f.Gauge1.SetProgress(0)
				f.Gauge1.SetMaxValue(int32(len(tmpds)))
			})
			err = f.Exportrenwudapingtai_excel(thread, tmpds, fname)
			if err != nil {
				return err
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
	f.ListView3.Invalidate()
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
			xsfile, err := xlsx.OpenFile(fname)
			if err != nil {
				return err
			}
			if len(xsfile.Sheets)== 0 {
				return errors.New("excel是空数据")
			}

			rows := xsfile.Sheets[0].Rows
			for idx, row := range rows {
				if len(row.Cells)<17||idx==0{
					continue
				}
				d := f.GetUploadDataOrCreate(strings.TrimSpace(row.Cells[0].Value))
				d.SSEV("样品匹配", "否")
				d.SSEV("抽样单编号", strings.TrimSpace(row.Cells[0].Value))
				d.SSEV("报告书编号", strings.TrimSpace(row.Cells[1].Value))
				d.SSEV("检测项目", "双击查看")
				d.SSEV("监督抽检报告备注", strings.TrimSpace(row.Cells[14].Value))
				d.SSEV("风险监测报告备注", strings.TrimSpace(row.Cells[15].Value))
				d.SSEV("上传状态", "否")
				d.SSEV("上传结果", "")
				d.SSEV("样品名称", strings.TrimSpace(row.Cells[16].Value))

				d.AddSubitem(map[string]string{
					"检验项目":  strings.TrimSpace(row.Cells[2].Value),
					"检验结果":  strings.TrimSpace(row.Cells[3].Value),
					"结果单位":  strings.TrimSpace(row.Cells[4].Value),
					"结果判定":  strings.TrimSpace(row.Cells[5].Value),
					"检验依据":  strings.TrimSpace(row.Cells[6].Value),
					"判定依据":  strings.TrimSpace(row.Cells[7].Value),
					"最小允许限": strings.TrimSpace(row.Cells[8].Value),
					"最大允许限": strings.TrimSpace(row.Cells[9].Value),
					"允许限单位": strings.TrimSpace(row.Cells[10].Value),
					"方法检出限": strings.TrimSpace(row.Cells[11].Value),
					"检出限单位": strings.TrimSpace(row.Cells[12].Value),
					"说明":    strings.TrimSpace(row.Cells[13].Value),
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
	if f.uploaddatas == nil || len(f.uploaddatas) < int(idx) {
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
	sd := f.Dtpt2s1.Date().Format("2006-01-02")
	ed := f.Dtpt2s2.Date().Format("2006-01-02")
	taskfrom := url.QueryEscape(f.Edtt2s1.Text())
	f.Buttont2s2.SetEnabled(false)
	go func() {
		defer vcl.ThreadSync(func() {
			f.Buttont2s2.SetEnabled(true)
		})
		err := func() error {
			if r == false {
				return nil
			}
			var dds *nifdc.Api_food_getFood_r
			var err error
			if tp == 0 {
				dds, err = nifdc.Test_platform_api_food_getFood(taskfrom, 8, sd, ed, 0, 10000, "", "desc", f.test_platform_ck, nil)
				if err != nil {
					return err
				}
			}
			if tp == 1 {
				dds, err = nifdc.Test_platform_api_agriculture_getAgriculture(taskfrom, 8, sd, ed, 0, 10000, "", "desc", f.test_platform_ck, nil)
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
	tp := 0
	if f.Cbbt2s1.ItemIndex() == 0 {
		tp = 0
	}
	if f.Cbbt2s1.ItemIndex() == 1 {
		tp = 1
	}
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
							if tp==0{
								fddetail, err := nifdc.Test_platform_foodTest_foodDetail(_td.Env_for_key("id").(int), f.test_platform_ck, nil)
								if err != nil {
									return err
								}
								testitems,err:=nifdc.Test_platform_api_food_getTestItems(fddetail,f.test_platform_ck,nil)
								if err != nil {
									return err
								}
								testinfo, err := nifdc.Test_platform_api_food_getTestInfo(fddetail["sd"], f.test_platform_ck, nil)
								if err != nil {
									return err
								}
								updatas:=nifdc.Build_agriculture_updata(testitems.Rows,testinfo.Rows,_td.Subitem())

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
								err = nifdc.Test_platform_api_food_save(fddetail, updatas, f.test_platform_ck, nil)
								if err != nil {
									return err
								}

								atomic.AddInt32(&nok, 1)
								_td.SSEV("上传结果", "成功")
								return nil
							}else if tp==1{
								fddetail, err := nifdc.Test_platform_agricultureTest_agricultureDetail(_td.Env_for_key("id").(int), f.test_platform_ck, nil)
								if err != nil {
									return err
								}
								//fmt.Println(fddetail["sample_code"])

								testitems,err:=nifdc.Test_platform_api_agriculture_getTestItems(fddetail,f.test_platform_ck,nil)
								if err != nil {
									return err
								}
								testinfo, err := nifdc.Test_platform_api_agriculture_getTestInfo(fddetail["sd"], f.test_platform_ck, nil)
								if err != nil {
									return err
								}
								updatas:=nifdc.Build_agriculture_updata(testitems.Rows,testinfo.Rows,_td.Subitem())

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

								err = nifdc.Test_platform_api_agriculture_save(fddetail, updatas, f.test_platform_ck, nil)
								if err != nil {
									return err
								}

								atomic.AddInt32(&nok, 1)
								_td.SSEV("上传结果", "成功")
								return nil
							}

							return errors.New("不支持这个模式")
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

}
func timeSub(t1, t2 time.Time) int {
	t1 = time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, time.Local)
	t2 = time.Date(t2.Year(), t2.Month(), t2.Day(), 0, 0, 0, 0, time.Local)
	return int(t1.Sub(t2).Hours() / 24)
}
func (f *TFormHome) getworkday(tm time.Time) string {
	d := timeSub(time.Now(), tm)
	nd := 0
	for i := 0; i < d; i++ {
		b, err := holidays.IsWorkingday(tm.AddDate(0, 0, (i + 1)))
		if err != nil {
			return "-"
		}
		if b == true {
			nd++
		}
	}
	if nd > 25 {
		return "-"
	}
	return fmt.Sprintf("%d", nd)
}
func (f *TFormHome) OnListView3Data(sender vcl.IObject, item *vcl.TListItem) {
	f.getFood_ds_lk.Lock()
	defer f.getFood_ds_lk.Unlock()
	idx := item.Index()
	if f.getFood_ds == nil || len(f.getFood_ds) < int(idx) {
		return
	}
	d := f.getFood_ds[idx]
	item.SetCaption(fmt.Sprintf("%d", idx+1))
	sitem := item.SubItems()
	sitem.Add(time.Unix(d.Sp_d_38/1000, 0).Format("2006-01-02"))
	sitem.Add(time.Unix(d.Sp_d_46/1000, 0).Format("2006-01-02"))
	sitem.Add(f.getworkday(time.Unix(d.Sp_d_46/1000, 0)))
	sitem.Add(d.Sp_s_16)
	sitem.Add(time.Unix(d.Updated_at/1000, 0).Format("2006-01-02 15:04:05"))
	sitem.Add(d.Sp_s_3)
	sitem.Add(d.Sp_s_14)
	sitem.Add(d.Sp_s_2_1)
	sitem.Add(d.Sp_s_44)
	sitem.Add(d.Sp_s_43)
	sitem.Add(d.Sp_s_35)
	sitem.Add(d.Sp_s_71)
	sitem.Add(d.User.SEV("处理状态"))
	sitem.Add(d.User.SEV("处理结果"))
}
func (f *TFormHome) OnListView3Resize(sender vcl.IObject) {
	go vcl.ThreadSync(func() {
		lastitem := f.ListView3.Column(f.ListView3.Columns().Count() - 1)
		lastitem.SetWidth(lastitem.Width() - 10)
	})
}

//搜索按钮
func (f *TFormHome) OnButtonp3s1Click(sender vcl.IObject) {
	tp := 0
	if f.Cbbt3s1.ItemIndex() == 0 {
		tp = 0
	}
	if f.Cbbt3s1.ItemIndex() == 1 {
		tp = 1
	}

	tasktype := int(f.Cbbt3s2.ItemIndex())
	order := "desc"
	if f.Cbbt3s4.ItemIndex() == 0 {
		order = "desc"
	} else {
		order = "asc"
	}
	sort := ""
	switch f.Cbbt3s3.ItemIndex() {
	case 1:
		sort = "sp_d_38"
	case 2:
		sort = "sp_s_16"
	case 3:
		sort = "updated_at"
	case 4:
		sort = "sp_s_71"
	}
	sd := f.Dtpt3s1.Date().Format("2006-01-02")
	tmsd := f.Dtpt3s1.Date()
	ed := f.Dtpt3s2.Date().Format("2006-01-02")
	taskfrom := url.QueryEscape(f.Edtt3s1.Text())
	f.Buttonp3s1.SetEnabled(false)
	go func() {
		defer vcl.ThreadSync(func() {
			f.Buttonp3s1.SetEnabled(true)
		})
		err := func() error {
			if r == false {
				return nil
			}

			var tmpds *nifdc.Api_food_getFood_r
			var err error
			if tp == 0 {
				tmpds, err = nifdc.Test_platform_api_food_getFood(taskfrom, tasktype, sd, ed, 0, 20000, sort, order, f.test_platform_ck, nil)
				if err != nil {
					return err
				}
			}
			if tp == 1 {
				tmpds, err = nifdc.Test_platform_api_agriculture_getAgriculture(taskfrom, tasktype, sd, ed, 0, 20000, sort, order, f.test_platform_ck, nil)
				if err != nil {
					return err
				}
			}
			tmpds1 := make([]*nifdc.Api_food_getFood_o, 0)
			for _, it := range tmpds.Rows {
				if it.Created_at/1000 >= tmsd.Unix() {
					tmpds1 = append(tmpds1, it)
				}
			}

			vcl.ThreadSync(func() {
				f.getFood_ds_lk.Lock()
				f.getFood_ds = tmpds1
				f.getFood_ds_lk.Unlock()

				f.ListView3.Items().SetCount(int32(len(f.getFood_ds)))
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
func (f *TFormHome) OnTss3Show(sender vcl.IObject) {
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
func inArray(s string, arr []string) bool {
	for _, it := range arr {
		if it == s {
			return true
		}
	}
	return false
}
func (f *TFormHome) cv_jianyanrenwu_dbv(dbk string, v string) string {
	intarr := []string{
		"进度",
		"type",
		"年度",
		"季度",
		"money",
		"企业性质",
		"样品来源",
		"检验次数",
		"产商品",
		"报告页数",
		"打印",
		"样品处理",
		"返工",
		"流转方式",
		"业务委托id",
		"委托排除",
		"委检",
		"检验结束日期",
		"项目收费折扣",
		"委托时付费",
		"打印份数",
		"累计参照",
		"结算",
		"数",
		"数2",
		"天数",
		"发票号",
		"是否加急",
		"实收费",
		"发票号金额",
		"项目数",
		"不合格项目数",
		"样品编号数",
		"当前页数",
		"起付检验费",
	}
	datearr := []string{
		"委托日期",
		"生产日期",
		"抽到样日期",
		"下达日期",
		"要求完成日期",
		"检验日期",
		"签发日期",
		"样品入库日期",
		"分派日期",
		"报告发出日期",
		"领样日期",
		"返样日期",
		"退样日期",
		"创建日期",
		"完成日期",
		"检验结束日期",
		"封样日期",
		"二审日期",
		"入库时间",
		"收费日期",
		"业务分派日期",
		"送样日期",
		"收样日期",
		"打印日期",
	}
	if inArray(dbk, intarr) == true {
		return v
	}
	if inArray(dbk, datearr) == true {
		fmt.Sprintf("'%s'", v)
	}
	return fmt.Sprintf("'%s'", v)
}

//下载检验任务导出到sql
func (f *TFormHome) Exportxiazaijianyanrenwu_sql(thread int, data []*nifdc.Api_food_getFood_o, tp int) error {
	if models.Ctx() == nil {
		return errors.New("数据库未配置")
	}
	for _, d := range data {
		d.User.SSEV("处理状态", "")
		d.User.SSEV("处理结果", "")
	}
	nerr := int32(0)
	nok := int32(0)
	nrey := int32(0)
	th := threadpool.NewThreadPool(thread, len(data))
	for _, d := range data {
		_d := d
		th.Req(func() interface{} {
			defer vcl.ThreadSync(func() {
				f.Gauge3.SetProgress(f.Gauge3.Progress() + 1)
				_d.User.SSEV("处理状态", "完成")
			})
			err := func() error {
				itr, err := nettool.RNet_Call_1(&nettool.RNetOptions{}, func(source *addrmgr.AddrSource) (i interface{}, e error) {
					if tp == 0 { //普通食品
						tr, err := nifdc.Test_platform_foodTest_foodDetail(_d.Id, f.test_platform_ck, nil)
						if err != nil {
							return nil, err
						}
						return tr, nil
					} else if tp == 1 { //农产品
						tr, err := nifdc.Test_platform_agricultureTest_agricultureDetail(_d.Id, f.test_platform_ck, nil)
						if err != nil {
							return nil, err
						}
						return tr, nil
					}
					return nil, errors.New("不支持的模式")
				})
				if err != nil {
					return err
				}
				tr := itr.(map[string]string)
				cn := 0
				err = models.Ctx().Model(&models.Jianyanrenwu{}).Where("抽样单号=?", tr["抽样基础信息_抽样单编号"]).Count(&cn).Error
				if err != nil {
					return err
				}
				if cn != 0 {
					//已导入的跳过
					fds := []string{

					}
					vds := []string{
					}

					tmj := template.New("tmj")
					tmj.Funcs(map[string]interface{}{
						"replace": strings.ReplaceAll,
						"replaceex": replaceex,
					})
					for _, it := range f.jianyanjieguosql_rule {
						dbk := it[0]
						webk := it[1]
						fds = append(fds, dbk)

						_, err = tmj.Parse(webk)
						if err != nil {
							fmt.Println(err)
							return err
						}

						var tmpwebv bytes.Buffer

						err := tmj.Execute(&tmpwebv, tr)
						if err != nil {
							fmt.Println(err)
							return err
						}

						webv := f.cv_jianyanrenwu_dbv(dbk, tmpwebv.String())
						vds = append(vds, webv)
					}

					forupdates:=[]string{}
					for idx,fd:=range fds{
						forupdates=append(forupdates,fmt.Sprintf("%s=%s",fd,vds[idx],))
					}

					err = models.Ctx().Exec(fmt.Sprintf("update 检验任务 set %s where 抽样单号=?", strings.Join(forupdates,",")),fmt.Sprintf("'%s'", tr["抽样基础信息_抽样单编号"]),).Error
					if err != nil {
						return err
					}
					atomic.AddInt32(&nrey, 1)
					_d.User.SSEV("处理结果", "更新")
					return nil
				}else{
					fds := []string{
						"id",
						"任务编号",
						"抽样单号",
					}
					vds := []string{
						fmt.Sprintf("'%s'", models.Build_id()),
						fmt.Sprintf("'%s'", models.Build_taskid()),
						fmt.Sprintf("'%s'", tr["抽样基础信息_抽样单编号"]),
					}

					tmj := template.New("tmj")
					tmj.Funcs(map[string]interface{}{
						"replace": strings.ReplaceAll,
						"replaceex": replaceex,
					})
					for _, it := range f.jianyanjieguosql_rule {
						dbk := it[0]
						webk := it[1]
						fds = append(fds, dbk)

						_, err = tmj.Parse(webk)
						if err != nil {
							fmt.Println(err)
							return err
						}

						var tmpwebv bytes.Buffer

						err := tmj.Execute(&tmpwebv, tr)
						if err != nil {
							fmt.Println(err)
							return err
						}

						webv := f.cv_jianyanrenwu_dbv(dbk, tmpwebv.String())
						vds = append(vds, webv)
					}
					err = models.Ctx().Exec(fmt.Sprintf("insert into 检验任务 (%s) values (%s)", strings.Join(fds, ","), strings.Join(vds, ","))).Error
					if err != nil {
						return err
					}
					atomic.AddInt32(&nok, 1)
					_d.User.SSEV("处理结果", "完成")
					return nil
				}
				return nil
			}()
			if err != nil {
				atomic.AddInt32(&nerr, 1)
				_d.User.SSEV("处理结果", err.Error())
				return err
			}
			return nil
		})

	}
	th.Wait()
	vcl.ThreadSync(func() {
		vcl.ShowMessage(fmt.Sprintf("成功:%d\n\n失败:%d\n\n更新:%d", atomic.LoadInt32(&nok), atomic.LoadInt32(&nerr), atomic.LoadInt32(&nrey)))
	})
	return nil
}
func (f *TFormHome) OnButtonp3s2Click(sender vcl.IObject) {
	tp := 0
	if f.Cbbt3s1.ItemIndex() == 0 {
		tp = 0
	}
	if f.Cbbt3s1.ItemIndex() == 1 {
		tp = 1
	}
	f.getFood_ds_lk.Lock()
	tmpds := f.getFood_ds
	f.getFood_ds_lk.Unlock()
	f.Buttonp3s2.SetEnabled(false)
	go func() {
		defer vcl.ThreadSync(func() {
			f.Buttonp3s2.SetEnabled(true)
		})
		err := func() error {
			if w == false {
				return nil
			}
			var err error
			f.jianyanjieguosql_rule, err = f.readrule("./下载检验结果sql规则.txt")
			if err != nil {
				return err
			}

			if tmpds == nil || len(tmpds) == 0 {
				return errors.New("数据不能为空")
			}
			vcl.ThreadSync(func() {
				f.Gauge3.SetProgress(0)
				f.Gauge3.SetMaxValue(int32(len(tmpds)))
			})
			err = f.Exportxiazaijianyanrenwu_sql(thread, tmpds, tp)
			if err != nil {
				return err
			}
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

//删除检验结果导出到sql
func (f *TFormHome) Deletexiazaijianyanrenwu_sql(thread int, data []*nifdc.Api_food_getFood_o) error {
	if models.Ctx() == nil {
		return errors.New("数据库未配置")
	}
	for _, d := range data {
		d.User.SSEV("处理状态", "")
		d.User.SSEV("处理结果", "")
	}
	nerr := int32(0)
	nok := int32(0)
	th := threadpool.NewThreadPool(thread, len(data))
	for _, d := range data {
		_d := d
		th.Req(func() interface{} {
			defer vcl.ThreadSync(func() {
				f.Gauge3.SetProgress(f.Gauge3.Progress() + 1)
				_d.User.SSEV("处理状态", "完成")
			})
			err := func() error {
				err := models.Ctx().Delete(&models.Jianyanrenwu{}, "抽样单号=?", _d.Sp_s_16).Error
				if err != nil {
					return err
				}
				atomic.AddInt32(&nok, 1)
				return nil
			}()
			if err != nil {
				atomic.AddInt32(&nerr, 1)
				_d.User.SSEV("处理结果", err.Error())
				return err
			}
			_d.User.SSEV("处理结果", "完成")
			return nil
		})

	}
	th.Wait()
	vcl.ThreadSync(func() {
		vcl.ShowMessage(fmt.Sprintf("成功:%d\n\n失败:%d\n\n", atomic.LoadInt32(&nok), atomic.LoadInt32(&nerr)))
	})
	return nil
}
func (f *TFormHome) OnButtonp3s3Click(sender vcl.IObject) {
	f.getFood_ds_lk.Lock()
	tmpds := f.getFood_ds
	f.getFood_ds_lk.Unlock()
	f.Buttonp3s3.SetEnabled(false)
	go func() {
		defer vcl.ThreadSync(func() {
			f.Buttonp3s3.SetEnabled(true)
		})
		err := func() error {
			if w == false {
				return nil
			}
			if tmpds == nil || len(tmpds) == 0 {
				return errors.New("数据不能为空")
			}
			vcl.ThreadSync(func() {
				f.Gauge3.SetProgress(0)
				f.Gauge3.SetMaxValue(int32(len(tmpds)))
			})
			err := f.Deletexiazaijianyanrenwu_sql(thread, tmpds)
			if err != nil {
				return err
			}
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

//下载检验结果导出到excel
func (f *TFormHome) Exportxiazaijianyanrenwu_excel(thread int, data []*nifdc.Api_food_getFood_o, fname string, tp int) error {
	for _, d := range data {
		d.User.SSEV("处理状态", "")
		d.User.SSEV("处理结果", "")
	}
	nerr := int32(0)
	nok := int32(0)
	nrey := int32(0)
	xlsxf := xlsx.NewFile()
	sheet, err := xlsxf.AddSheet("数据报告")
	if err != nil {
		return err
	}
	sheet_lk := sync.Mutex{}
	row := sheet.AddRow()
	for _, it := range f.jianyanjieguoexcel_rule {
		dbk := it[0]
		cell := row.AddCell()
		cell.SetString(dbk)
	}
	sheet.SetColWidth(0, len(f.jianyanjieguoexcel_rule), 15)

	th := threadpool.NewThreadPool(thread, len(data))
	for _, d := range data {
		_d := d
		th.Req(func() interface{} {
			defer vcl.ThreadSync(func() {
				f.Gauge3.SetProgress(f.Gauge3.Progress() + 1)
				_d.User.SSEV("处理状态", "完成")
			})
			err := func() error {
				itr, err := nettool.RNet_Call_1(&nettool.RNetOptions{}, func(source *addrmgr.AddrSource) (i interface{}, e error) {
					if tp == 0 { //普通食品
						tr, err := nifdc.Test_platform_foodTest_foodDetail(_d.Id, f.test_platform_ck, nil)
						if err != nil {
							return nil, err
						}
						return tr, nil
					} else if tp == 1 { //农产品
						tr, err := nifdc.Test_platform_agricultureTest_agricultureDetail(_d.Id, f.test_platform_ck, nil)
						if err != nil {
							return nil, err
						}
						return tr, nil
					}
					return nil, errors.New("不支持的模式")
				})
				if err != nil {
					return err
				}
				tr := itr.(map[string]string)

				sheet_lk.Lock()
				row := sheet.AddRow()
				tmj := template.New("tmj")
				tmj.Funcs(map[string]interface{}{
					"replace": strings.ReplaceAll,
					"replaceex": replaceex,
				})
				for _, it := range f.jianyanjieguoexcel_rule {
					webk := it[1]

					_, err = tmj.Parse(webk)
					if err != nil {
						fmt.Println(err)
						return err
					}
					var tmpwebv bytes.Buffer
					err := tmj.Execute(&tmpwebv, tr)
					if err != nil {
						fmt.Println(err)
						return err
					}

					cl := row.AddCell()
					cl.SetString(tmpwebv.String())
				}
				sheet_lk.Unlock()

				atomic.AddInt32(&nok, 1)
				return nil
			}()
			if err != nil {
				atomic.AddInt32(&nerr, 1)
				_d.User.SSEV("处理结果", err.Error())
				return err
			}
			_d.User.SSEV("处理结果", "完成")
			return nil
		})

	}
	th.Wait()
	err = xlsxf.Save(fname)
	if err != nil {
		return err
	}
	vcl.ThreadSync(func() {
		vcl.ShowMessage(fmt.Sprintf("成功:%d\n\n失败:%d\n\n已存在:%d", atomic.LoadInt32(&nok), atomic.LoadInt32(&nerr), atomic.LoadInt32(&nrey)))
	})
	return nil
}
func (f *TFormHome) OnButtonp3s4Click(sender vcl.IObject) {
	tp := 0
	if f.Cbbt3s1.ItemIndex() == 0 {
		tp = 0
	}
	if f.Cbbt3s1.ItemIndex() == 1 {
		tp = 1
	}
	if f.SaveDialog1.Execute() == false {
		return
	}
	fname := f.SaveDialog1.FileName()
	f.getFood_ds_lk.Lock()
	tmpds := f.getFood_ds
	f.getFood_ds_lk.Unlock()
	f.Buttonp3s4.SetEnabled(false)
	go func() {
		defer vcl.ThreadSync(func() {
			f.Buttonp3s4.SetEnabled(true)
		})
		err := func() error {
			if w == false {
				return nil
			}
			var err error
			f.jianyanjieguoexcel_rule, err = f.readrule("./下载检验结果excel规则.txt")
			if err != nil {
				return err
			}

			if tmpds == nil || len(tmpds) == 0 {
				return errors.New("数据不能为空")
			}
			vcl.ThreadSync(func() {
				f.Gauge3.SetProgress(0)
				f.Gauge3.SetMaxValue(int32(len(tmpds)))
			})
			err = f.Exportxiazaijianyanrenwu_excel(thread, tmpds, fname, tp)
			if err != nil {
				return err
			}
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

//下载检验项目导出到sql
func (f *TFormHome) Exportxiazaijianyanxiangmu_sql(thread int, data []*nifdc.Api_food_getFood_o, tp int) error {
	if models.Ctx() == nil {
		return errors.New("数据库未配置")
	}
	for _, d := range data {
		d.User.SSEV("处理状态", "")
		d.User.SSEV("处理结果", "")
	}
	nerr := int32(0)
	nok := int32(0)
	nrey := int32(0)
	th := threadpool.NewThreadPool(thread, len(data))
	for _, d := range data {
		_d := d
		th.Req(func() interface{} {
			defer vcl.ThreadSync(func() {
				f.Gauge3.SetProgress(f.Gauge3.Progress() + 1)
				_d.User.SSEV("处理状态", "完成")
			})
			err := func() error {
				itr, err := nettool.RNet_Call_1(&nettool.RNetOptions{}, func(source *addrmgr.AddrSource) (i interface{}, e error) {
					if tp == 0 { //普通食品
						tr, err := nifdc.Test_platform_foodTest_foodDetail(_d.Id, f.test_platform_ck, nil)
						if err != nil {
							return nil, err
						}
						return tr, nil
					} else if tp == 1 { //农产品
						tr, err := nifdc.Test_platform_agricultureTest_agricultureDetail(_d.Id, f.test_platform_ck, nil)
						if err != nil {
							return nil, err
						}
						return tr, nil
					}
					return nil, errors.New("不支持的模式")
				})
				if err != nil {
					return err
				}
				tr := itr.(map[string]string)
				itsubtr, err := nettool.RNet_Call_1(&nettool.RNetOptions{}, func(source *addrmgr.AddrSource) (i interface{}, err error) {
					if tp == 0 { //普通食品
						testinfor, err := nifdc.Test_platform_api_food_getTestInfo(tr["sd"], f.test_platform_ck, nil)
						if err != nil {
							return nil, err
						}
						itemsr,err:=nifdc.Test_platform_api_food_getTestItems(tr,f.test_platform_ck,nil)
						if err != nil {
							return nil, err
						}
						rmp:=nifdc.TestInfotoMap(testinfor.Rows,itemsr.Rows)
						return rmp, nil
					} else if tp == 1 { //农产品
						testinfor, err := nifdc.Test_platform_api_agriculture_getTestInfo(tr["sd"], f.test_platform_ck, nil)
						if err != nil {
							return nil, err
						}
						itemsr,err:=nifdc.Test_platform_api_agriculture_getTestItems(tr,f.test_platform_ck,nil)
						if err != nil {
							return nil, err
						}
						rmp:=nifdc.TestInfotoMap(testinfor.Rows,itemsr.Rows)
						return rmp, nil
					}
					return nil, errors.New("不支持的模式")
				})
				if err != nil {
					return err
				}
				subtr := itsubtr.([]map[string]string)
				renwu := &models.Jianyanrenwu{}
				err = models.Ctx().Model(&models.Jianyanrenwu{}).Where("抽样单号=?", tr["抽样基础信息_抽样单编号"]).Find(&renwu).Error
				if err != nil {
					return err
				}


				jianyanshi:="GC"
				jianyanyuan:="检验员"
				jindu:="20"

				if user=="15738889730"{
					jianyanshi="YJ"
					jianyanyuan=""
				}

				if  user=="18039661206"{
					jianyanshi="YJ"
					jianyanyuan="检验员"

				}
				if user!="13483719195" {
					jindu="0"
				}

				for idx, subr := range subtr {
					subidx := idx + 1
					rn := 0
					err = models.Ctx().Model(&models.Jianyanxiangmu{}).Where("任务编号=? and 项目名称=?", renwu.V任务编号, subr["检验项目*"]).Count(&rn).Error
					if err != nil {
						return err
					}
					if rn == 0 { //插入
						err = models.Ctx().Model(&models.Jianyanxiangmu{}).Exec("insert into 检验项目 (序号,任务编号,显示序号,项目名称,样品名称,单位,检验方法,实测值,单项结论,判定依据,最小允许限,最大允许限,检出限,标准值,备注,检验室,检验员,进度,返工) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
							strconv.Itoa(subidx),
							renwu.V任务编号,
							strconv.Itoa(subidx),
							subr["检验项目*"],
							tr["抽检样品信息_样品名称"],
							subr["结果单位*"],
							subr["检验依据*"],
							subr["检验结果*"],
							subr["结果判定*"],
							subr["判定依据*"],
							subr["最小允许限*"],
							subr["最大允许限*"],
							subr["方法检出限*"],
							subr["最大允许限*"],
							subr["备注"],
							jianyanshi,
							jianyanyuan,
							jindu,
							"0",
						).Error
						if err != nil {
							return err
						}
					} else { //添加
						err = models.Ctx().Model(&models.Jianyanxiangmu{}).Exec("update 检验项目 set 项目名称=?,单位=?,检验方法=?,实测值=?,单项结论=?,判定依据=?,最小允许限=?,最大允许限=?,检出限=?,标准值=?,备注=? where 任务编号=? and 项目名称=?",
							subr["检验项目*"],
							subr["结果单位*"],
							subr["检验依据*"],
							subr["检验结果*"],
							subr["结果判定*"],
							subr["判定依据*"],
							subr["最小允许限*"],
							subr["最大允许限*"],
							subr["方法检出限*"],
							subr["最大允许限*"],
							subr["备注"],
							renwu.V任务编号,
							subr["检验项目*"],
						).Error
						if err != nil {
							return err
						}
					}
				}
				atomic.AddInt32(&nok, 1)
				return nil
			}()
			if err != nil {
				atomic.AddInt32(&nerr, 1)
				_d.User.SSEV("处理结果", err.Error())
				return err
			}
			_d.User.SSEV("处理结果", "完成")
			return nil
		})

	}
	th.Wait()
	vcl.ThreadSync(func() {
		vcl.ShowMessage(fmt.Sprintf("成功:%d\n\n失败:%d\n\n已存在:%d", atomic.LoadInt32(&nok), atomic.LoadInt32(&nerr), atomic.LoadInt32(&nrey)))
	})
	return nil
}
func (f *TFormHome) OnButtonp3s5Click(sender vcl.IObject) {
	tp := 0
	if f.Cbbt3s1.ItemIndex() == 0 {
		tp = 0
	}
	if f.Cbbt3s1.ItemIndex() == 1 {
		tp = 1
	}
	f.getFood_ds_lk.Lock()
	tmpds := f.getFood_ds
	f.getFood_ds_lk.Unlock()
	f.Buttonp3s5.SetEnabled(false)
	go func() {
		defer vcl.ThreadSync(func() {
			f.Buttonp3s5.SetEnabled(true)
		})
		err := func() error {
			if w == false {
				return nil
			}

			if tmpds == nil || len(tmpds) == 0 {
				return errors.New("数据不能为空")
			}
			vcl.ThreadSync(func() {
				f.Gauge3.SetProgress(0)
				f.Gauge3.SetMaxValue(int32(len(tmpds)))
			})
			err := f.Exportxiazaijianyanxiangmu_sql(thread, tmpds, tp)
			if err != nil {
				return err
			}
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

//下载检验项目导出到web
func (f *TFormHome) Exportxiazaijianyanxiangmu_web(thread int, data []*nifdc.Api_food_getFood_o, tp int) error {
	if models.Ctx() == nil {
		return errors.New("数据库未配置")
	}
	for _, d := range data {
		d.User.SSEV("处理状态", "")
		d.User.SSEV("处理结果", "")
	}
	nerr := int32(0)
	nok := int32(0)
	nrey := int32(0)
	th := threadpool.NewThreadPool(thread, len(data))
	for _, d := range data {
		_d := d
		th.Req(func() interface{} {
			defer vcl.ThreadSync(func() {
				f.Gauge3.SetProgress(f.Gauge3.Progress() + 1)
				_d.User.SSEV("处理状态", "完成")
			})
			err := func() error {
				itr, err := nettool.RNet_Call_1(&nettool.RNetOptions{}, func(source *addrmgr.AddrSource) (i interface{}, e error) {
					if tp == 0 { //普通食品
						tr, err := nifdc.Test_platform_foodTest_foodDetail(_d.Id, f.test_platform_ck, nil)
						if err != nil {
							return nil, err
						}
						return tr, nil
					} else if tp == 1 { //农产品
						tr, err := nifdc.Test_platform_agricultureTest_agricultureDetail(_d.Id, f.test_platform_ck, nil)
						if err != nil {
							return nil, err
						}
						return tr, nil
					}
					return nil, errors.New("不支持的模式")
				})
				if err != nil {
					return err
				}
				tr := itr.(map[string]string)
				itsubtr, err := nettool.RNet_Call_1(&nettool.RNetOptions{}, func(source *addrmgr.AddrSource) (i interface{}, err error) {
					if tp == 0 { //普通食品
						testinfor, err := nifdc.Test_platform_api_food_getTestInfo(tr["sd"], f.test_platform_ck, nil)
						if err != nil {
							return nil, err
						}
						itemsr,err:=nifdc.Test_platform_api_food_getTestItems(tr,f.test_platform_ck,nil)
						if err != nil {
							return nil, err
						}
						rmp:=nifdc.TestInfotoMap(testinfor.Rows,itemsr.Rows)
						return rmp, nil
					} else if tp == 1 { //农产品
						testinfor, err := nifdc.Test_platform_api_agriculture_getTestInfo(tr["sd"], f.test_platform_ck, nil)
						if err != nil {
							return nil, err
						}
						itemsr,err:=nifdc.Test_platform_api_agriculture_getTestItems(tr,f.test_platform_ck,nil)
						if err != nil {
							return nil, err
						}
						rmp:=nifdc.TestInfotoMap(testinfor.Rows,itemsr.Rows)
						return rmp, nil
					}
					return nil, errors.New("不支持的模式")
				})
				if err != nil {
					return err
				}
				subtr := itsubtr.([]map[string]string)
				renwu := &models.Jianyanrenwu{}
				err = models.Ctx().Model(&models.Jianyanrenwu{}).Where("抽样单号=?", tr["抽样基础信息_抽样单编号"]).Find(&renwu).Error
				if err != nil {
					return err
				}

				for idx, subr := range subtr {
					subidx := idx + 1
					rn := 0
					err = models.Ctx().Model(&models.Jianyanxiangmu{}).Where("任务编号=? and 项目名称=?", renwu.V任务编号, subr["检验项目*"]).Count(&rn).Error
					if err != nil {
						return err
					}
					if rn == 0 { //插入
						err = models.Ctx().Model(&models.Jianyanxiangmu{}).Exec("insert into 检验项目 (序号,任务编号,显示序号,项目名称,样品名称,单位,检验方法,实测值,单项结论,判定依据) values (?,?,?,?,?,?,?,?,?,?)",
							strconv.Itoa(subidx),
							renwu.V任务编号,
							strconv.Itoa(subidx),
							subr["检验项目*"],
							tr["抽检样品信息_样品名称"],
							subr["结果单位*"],
							subr["检验依据*"],
							subr["检验结果*"],
							subr["结果判定*"],
							subr["判定依据*"],
						).Error
						if err != nil {
							return err
						}
					} else { //添加

					}
				}
				atomic.AddInt32(&nok, 1)
				return nil
			}()
			if err != nil {
				atomic.AddInt32(&nerr, 1)
				_d.User.SSEV("处理结果", err.Error())
				return err
			}
			_d.User.SSEV("处理结果", "完成")
			return nil
		})

	}
	th.Wait()
	vcl.ThreadSync(func() {
		vcl.ShowMessage(fmt.Sprintf("成功:%d\n\n失败:%d\n\n已存在:%d", atomic.LoadInt32(&nok), atomic.LoadInt32(&nerr), atomic.LoadInt32(&nrey)))
	})
	return nil
}
func (f *TFormHome) OnButtonp3s6Click(sender vcl.IObject) {

	vcl.ShowMessage("功能开发中")
	return
	tp := 0
	if f.Cbbt3s1.ItemIndex() == 0 {
		tp = 0
	}
	if f.Cbbt3s1.ItemIndex() == 1 {
		tp = 1
	}
	f.getFood_ds_lk.Lock()
	tmpds := f.getFood_ds
	f.getFood_ds_lk.Unlock()
	f.Buttonp3s6.SetEnabled(false)
	go func() {
		defer vcl.ThreadSync(func() {
			f.Buttonp3s6.SetEnabled(true)
		})
		err := func() error {
			if w == false {
				return nil
			}

			if tmpds == nil || len(tmpds) == 0 {
				return errors.New("数据不能为空")
			}
			vcl.ThreadSync(func() {
				f.Gauge3.SetProgress(0)
				f.Gauge3.SetMaxValue(int32(len(tmpds)))
			})
			err := f.Exportxiazaijianyanxiangmu_web(thread, tmpds, tp)
			if err != nil {
				return err
			}
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


//下载任务大平台导出到sql
func (f *TFormHome) Exportxiazairenwudapingtai_sql(thread int, data []*nifdc.Data_o, ) error {
	if models.Ctx() == nil {
		return errors.New("数据库未配置")
	}
	for _, d := range data {
		d.User.SSEV("处理状态", "")
		d.User.SSEV("处理结果", "")
	}
	nerr := int32(0)
	nok := int32(0)
	nrey := int32(0)
	th := threadpool.NewThreadPool(thread, len(data))
	for _, d := range data {
		_d := d
		th.Req(func() interface{} {
			defer vcl.ThreadSync(func() {
				f.Gauge1.SetProgress(f.Gauge3.Progress() + 1)
				_d.User.SSEV("处理状态", "完成")
			})
			err := func() error {
				itr, err := nettool.RNet_Call_1(&nettool.RNetOptions{}, func(source *addrmgr.AddrSource) (i interface{}, e error) {
					tr, err := nifdc.Viewcheckedsample_full(_d.Sample_code, f.sample_ck, nil)
					if err != nil {
						return nil, err
					}
					return tr, nil
				})
				if err != nil {
					return err
				}
				tr := itr.(map[string]string)
				cn := 0
				err = models.Ctx().Model(&models.Jianyanrenwu{}).Where("抽样单号=?", tr["抽样基础信息_抽样单编号"]).Count(&cn).Error
				if err != nil {
					return err
				}
				if cn != 0 {
					//已导入的跳过
					fds := []string{

					}
					vds := []string{
					}

					tmj := template.New("tmj")
					tmj.Funcs(map[string]interface{}{
						"replace": strings.ReplaceAll,
						"replaceex": replaceex,
					})
					for _, it := range f.renwudapingtaisql_rule {
						dbk := it[0]
						webk := it[1]
						fds = append(fds, dbk)

						_, err = tmj.Parse(webk)
						if err != nil {
							fmt.Println(err)
							return err
						}

						var tmpwebv bytes.Buffer

						err := tmj.Execute(&tmpwebv, tr)
						if err != nil {
							fmt.Println(err)
							return err
						}

						webv := f.cv_jianyanrenwu_dbv(dbk, tmpwebv.String())
						vds = append(vds, webv)
					}

					forupdates:=[]string{}
					for idx,fd:=range fds{
						forupdates=append(forupdates,fmt.Sprintf("%s=%s",fd,vds[idx],))
					}

					err = models.Ctx().Exec(fmt.Sprintf("update 检验任务 set %s where 抽样单号=?", strings.Join(forupdates,",")),fmt.Sprintf("'%s'", tr["抽样基础信息_抽样单编号"]),).Error
					if err != nil {
						return err
					}
					atomic.AddInt32(&nrey, 1)
					_d.User.SSEV("处理结果", "更新")
					return nil
				}else{
					fds := []string{
						"id",
						"任务编号",
						"抽样单号",
					}
					vds := []string{
						fmt.Sprintf("'%s'", models.Build_id()),
						fmt.Sprintf("'%s'", models.Build_taskid()),
						fmt.Sprintf("'%s'", tr["抽样基础信息_抽样单编号"]),
					}

					tmj := template.New("tmj")
					tmj.Funcs(map[string]interface{}{
						"replace": strings.ReplaceAll,
						"replaceex": replaceex,
					})
					for _, it := range f.renwudapingtaisql_rule {
						dbk := it[0]
						webk := it[1]
						fds = append(fds, dbk)

						_, err = tmj.Parse(webk)
						if err != nil {
							fmt.Println(err)
							return err
						}

						var tmpwebv bytes.Buffer

						err := tmj.Execute(&tmpwebv, tr)
						if err != nil {
							fmt.Println(err)
							return err
						}

						webv := f.cv_jianyanrenwu_dbv(dbk, tmpwebv.String())
						vds = append(vds, webv)
					}
					err = models.Ctx().Exec(fmt.Sprintf("insert into 检验任务 (%s) values (%s)", strings.Join(fds, ","), strings.Join(vds, ","))).Error
					if err != nil {
						return err
					}
					atomic.AddInt32(&nok, 1)
					_d.User.SSEV("处理结果", "完成")
					return nil
				}
				return nil
			}()
			if err != nil {
				atomic.AddInt32(&nerr, 1)
				_d.User.SSEV("处理结果", err.Error())
				return err
			}
			return nil
		})

	}
	th.Wait()
	vcl.ThreadSync(func() {
		vcl.ShowMessage(fmt.Sprintf("成功:%d\n\n失败:%d\n\n更新:%d", atomic.LoadInt32(&nok), atomic.LoadInt32(&nerr), atomic.LoadInt32(&nrey)))
	})
	return nil
}
//导出任务大平台到sql
func (f *TFormHome) OnButtonp1s3Click(sender vcl.IObject) {
	f.sample_ds_lk.Lock()
	tmpds := f.sample_ds
	f.sample_ds_lk.Unlock()
	f.Buttonp1s3.SetEnabled(false)
	go func() {
		defer vcl.ThreadSync(func() {
			f.Buttonp1s3.SetEnabled(true)
		})
		err := func() error {
			if w == false {
				return nil
			}
			var err error
			f.renwudapingtaisql_rule, err = f.readrule("./下载任务大平台sql规则.txt")
			if err != nil {
				return err
			}

			if tmpds == nil || len(tmpds) == 0 {
				return errors.New("数据不能为空")
			}
			vcl.ThreadSync(func() {
				f.Gauge1.SetProgress(0)
				f.Gauge1.SetMaxValue(int32(len(tmpds)))
			})
			err = f.Exportxiazairenwudapingtai_sql(thread, tmpds,)
			if err != nil {
				return err
			}
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





//删除任务大平台导出到sql
func (f *TFormHome) Deleterenwudapingtai_sql(thread int, data []*nifdc.Data_o) error {
	if models.Ctx() == nil {
		return errors.New("数据库未配置")
	}
	for _, d := range data {
		d.User.SSEV("处理状态", "")
		d.User.SSEV("处理结果", "")
	}
	nerr := int32(0)
	nok := int32(0)
	th := threadpool.NewThreadPool(thread, len(data))
	for _, d := range data {
		_d := d
		th.Req(func() interface{} {
			defer vcl.ThreadSync(func() {
				f.Gauge1.SetProgress(f.Gauge3.Progress() + 1)
				_d.User.SSEV("处理状态", "完成")
			})
			err := func() error {
				err := models.Ctx().Delete(&models.Jianyanrenwu{}, "抽样单号=?", _d.Sample_code).Error
				if err != nil {
					return err
				}
				atomic.AddInt32(&nok, 1)
				return nil
			}()
			if err != nil {
				atomic.AddInt32(&nerr, 1)
				_d.User.SSEV("处理结果", err.Error())
				return err
			}
			_d.User.SSEV("处理结果", "完成")
			return nil
		})

	}
	th.Wait()
	vcl.ThreadSync(func() {
		vcl.ShowMessage(fmt.Sprintf("成功:%d\n\n失败:%d\n\n", atomic.LoadInt32(&nok), atomic.LoadInt32(&nerr)))
	})
	return nil
}
func (f *TFormHome) OnButtonp1s4Click(sender vcl.IObject) {
	f.sample_ds_lk.Lock()
	tmpds := f.sample_ds
	f.sample_ds_lk.Unlock()
	f.Buttonp1s4.SetEnabled(false)
	go func() {
		defer vcl.ThreadSync(func() {
			f.Buttonp1s4.SetEnabled(true)
		})
		err := func() error {
			if w == false {
				return nil
			}
			if tmpds == nil || len(tmpds) == 0 {
				return errors.New("数据不能为空")
			}
			vcl.ThreadSync(func() {
				f.Gauge1.SetProgress(0)
				f.Gauge1.SetMaxValue(int32(len(tmpds)))
			})
			err := f.Deleterenwudapingtai_sql(thread, tmpds)
			if err != nil {
				return err
			}
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

