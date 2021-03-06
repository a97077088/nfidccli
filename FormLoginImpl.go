// 由res2go自动生成。
// 在这里写你的事件。

package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/a97077088/nifdc"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/rtl"
	"io/ioutil"
	"math/rand"
	"nfidccli/models"
	"nfidccli/proc"
	"os"
	"strconv"
	"time"
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
			req := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(999999989)
			rlogin, err := nfidcproc.Login(context.Background(), &proc.LoginReq{
				R: enmp(map[string]string{
					"user": u,
					"tm":   fmt.Sprintf("%d", time.Now().UnixNano()),
					"req":  fmt.Sprintf("%d", req),
				}),
			})
			if err != nil {
				return errors.New("授权失败")
			}
			mpr := demp(rlogin.R)
			if mpr == nil {
				return errors.New("授权失败")
			}
			if fmt.Sprintf("%d", req) != mpr["req"] {
				return errors.New("授权失败")
			}
			sth := mpr["th"]
			if sth == "" {
				return errors.New("授权失败")
			}
			thread, err = strconv.Atoi(sth)
			if err != nil {
				return errors.New("授权失败")
			}

			sw := mpr["w"]
			sr := mpr["r"]
			if sw == "" || sr == "" {
				return errors.New("授权失败")
			}
			nw, err := strconv.Atoi(sw)
			if err != nil {
				return errors.New("授权失败")
			}
			nr, err := strconv.Atoi(sr)
			if err != nil {
				return errors.New("授权失败")
			}
			if nw == 1 {
				w = true
			}
			if nr == 1 {
				r = true
			}

			srvaddr := mpr["sql_server"]
			if srvaddr == "" {
				return nil
			}
			port, _ := strconv.Atoi(mpr["sql_port"])
			err = models.InitDb(mpr["sql_server"], port, mpr["sql_user"], mpr["sql_pass"], mpr["sql_dbname"])
			if err != nil {
				return err
			}
			user = u
			return nil
		}()
		if err != nil {
			vcl.ThreadSync(func() {
				vcl.ShowMessage(err.Error())
			})
			return
		}
		vcl.ThreadSync(func() {
			if debug == true {
				ioutil.WriteFile("./ck", []byte(ck), os.ModePerm)
			}
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

	if debug == true {
		if rtl.FileExists("./ck") == true {
			byck, _ := ioutil.ReadFile("./ck")
			ck = string(byck)
			w = true
			r = true
			thread = 1
			models.InitDb("122.51.93.214", 1433, "sa", "haosql", "testdb")
			FormHome.Show()
			go func() {
				time.Sleep(time.Millisecond * 200)
				vcl.ThreadSync(func() {
					f.Hide()
				})
			}()
			return
		}
	}
}
