package main

import (
	"fmt"
	"google.golang.org/grpc"
	"nfidccli/proc"
)

var thread int
var ck = "JSESSIONID=2306528F3FB00AE2F2F3117CA97F0616-n1; Path=/; HttpOnly;SSOWarnCookie=; Path=/; Expires=Thu, 01 Jan 1970 00:00:10 GMT; Max-Age=0;SSOTGTCookie=TGT-2008020-VkunLq5rtetsUn5MlImedzoVjBuLKgKRTZcp6drxmPbWeUSPch-sso.lims.gettec.com; Path=/;lims_service_sid=9021a567-8a67-48ea-80e0-83b84f293dcd-key-username; Path=/; Expires=Wed, 20 Nov 2019 15:33:12 GMT; Max-Age=604800"
var user = ""
var addr = "122.51.93.214:14662"
var w, r bool

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
