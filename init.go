package main

import (
	"fmt"
	"google.golang.org/grpc"
	"nfidccli/proc"
)

var thread int
var ck = "JSESSIONID=B982784226324A539BF7C49B16D08CE7-n1; Path=/; HttpOnly;SSOWarnCookie=; Path=/; Expires=Thu, 01 Jan 1970 00:00:10 GMT; Max-Age=0;SSOTGTCookie=TGT-2217131-aktp7cVjlboYyKfVy0nAbv7CsE6Bs2uay7YCeG6tbnOeIg9iPm-sso.lims.gettec.com; Path=/;lims_service_sid=23f34d22-121a-495c-95ed-d8213055a5a1-key-username; Path=/; Expires=Thu, 28 Nov 2019 12:55:30 GMT; Max-Age=604800"
var user = ""
var addr = "122.51.93.214:14662"
var w, r bool

//打开调试模式
var debug = false

var nfidcproc proc.NifdcrpcClient

func initrpc() error {
	grpccli, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return err
	}
	nfidcproc = proc.NewNifdcrpcClient(grpccli)
	return nil
}
func init() {
	err := initrpc()
	if err != nil {
		fmt.Println(err)
	}
}
