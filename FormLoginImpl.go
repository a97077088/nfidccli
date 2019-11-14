// 由res2go自动生成。
// 在这里写你的事件。

package main

import (
	"errors"
	"github.com/a97077088/nifdc"
	"github.com/ying32/govcl/vcl"
)

//::private::
type TFormLoginFields struct {
}

func (f *TFormLogin) OnButton1Click(sender vcl.IObject) {
	f.Button1.SetEnabled(false)
	u := f.LabelEdit1.Text()
	p := f.LabelEdit2.Text()
	go func() {
		defer vcl.ThreadSync(func() {
			f.Button1.SetEnabled(true)
		})
		err := func() error {
			if f.LabelEdit1.Text() == "" {
				return errors.New("用户名不能为空")
			}
			if f.LabelEdit2.Text() == "" {
				return errors.New("密码不能为空")
			}
			lt, execution, rck, err := nifdc.InitLoginck(nil)
			if err != nil {
				return err
			}
			ck, err = nifdc.Login(u, p, lt, execution, rck, nil)
			if err != nil {
				return err
			}
			err = nifdc.Index(ck, nil)
			if err != nil {
				return err
			}
			user = f.LabelEdit1.Text()
			return nil
		}()
		if err != nil {
			vcl.ThreadSync(func() {
				vcl.ShowMessage(err.Error())
			})
			return
		}
		vcl.ThreadSync(func() {
			//ioutil.WriteFile("./ck", []byte(ck), os.ModePerm)
			FormLogin.Hide()
			FormHome.Show()
		})
	}()
}

func (f *TFormLogin) OnButton2Click(sender vcl.IObject) {
	f.Close()
}

func (f *TFormLogin) OnFormCreate(sender vcl.IObject) {

}

func (f *TFormLogin) OnFormShow(sender vcl.IObject) {
	//if rtl.FileExists("./ck") == true {
	//	byck, _ := ioutil.ReadFile("./ck")
	//	ck = string(byck)
	//	FormHome.Show()
	//	go func() {
	//		time.Sleep(time.Millisecond * 200)
	//		vcl.ThreadSync(func() {
	//			f.Hide()
	//		})
	//	}()
	//	return
	//}
}
