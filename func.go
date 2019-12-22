package main

import (
	"encoding/base64"
	"encoding/json"
	"regexp"
)

func enmp(mp map[string]string) string {
	enb, err := json.Marshal(mp)
	if err != nil {
		return ""
	}
	for i, b := range enb {
		enb[i] = b ^ 0x30
	}
	return base64.StdEncoding.EncodeToString(enb)
}
func demp(s string) map[string]string {
	deb, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return nil
	}
	for i, b := range deb {
		deb[i] = b ^ 0x30
	}
	mp := make(map[string]string, 0)
	err = json.Unmarshal(deb, &mp)
	if err != nil {
		return nil
	}
	return mp
}

func replaceex(s string,old string,new string)string{
	ex:=regexp.MustCompile(old)
	return ex.ReplaceAllString(s,new)
}