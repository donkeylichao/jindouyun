package proxy

import (
	"jindouyun/account"
	"jindouyun/jdyError"
	"jindouyun/util"
	"fmt"
	"bufio"
	"os"
	"strings"
	"jindouyun/helper"
)

type JinDouYunProxy struct {
	account.JinDouYunConfig
	ProxyAddress string `json:"proxy"`
}

func (p *JinDouYunProxy) List() {
	util := util.GetJinDouYunUtil(p.Address, p.AppId, p.AppKey)

	query := map[string]string{}
	query["apiUrl"] = "proxies"

	re, err := util.Request(query, "GET", nil)
	jdyError.CheckError(err, true)
	helper.FormatOutPut(re)
}

func (p *JinDouYunProxy) Add()  {
	p.Set()
	//fmt.Printf("%s\n", p)
	//return
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Print("请确认:\n1.正确\n2.重新设置\n3.取消\n> ")
	for {
		input, err := inputReader.ReadString('\n')
		jdyError.CheckError(err, true)
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			fmt.Println("调用接口保存账号")
			p.add()
			os.Exit(0)
		case "2":
			p.Clear()
			p.Add()
		case "3":
			fmt.Println("已取消操作")
			os.Exit(0)
		default:
			continue
		}
	}

}

func (p *JinDouYunProxy) add()  {
	util := util.GetJinDouYunUtil(p.Address,p.AppId,p.AppKey)

	query := map[string]string{}
	query["apiUrl"] = "proxies"

	data := map[string]interface{}{}
	data["proxy"] = p.ProxyAddress

	re, err := util.Request(query, "POST", data)
	jdyError.CheckError(err,true)

	helper.FormatOutPut(re)
}

func (p *JinDouYunProxy) Update() {
	p.SetUpdate()
	//fmt.Printf("%s\n", p)
	//return
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Print("请确认:\n1.正确\n2.重新设置\n3.取消\n> ")
	for {
		input, err := inputReader.ReadString('\n')
		jdyError.CheckError(err, true)
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			fmt.Println("调用接口保存账号")
			p.update()
			os.Exit(0)
		case "2":
			p.Clear()
			p.Update()
		case "3":
			fmt.Println("已取消操作")
			os.Exit(0)
		default:
			continue
		}
	}
}

func (p *JinDouYunProxy) update()  {
	util := util.GetJinDouYunUtil(p.Address,p.AppId,p.AppKey)

	query := map[string]string{}
	query["apiUrl"] = "proxies/" + p.Id

	data := map[string]interface{}{}
	data["proxy"] = p.ProxyAddress

	re, err := util.Request(query, "PATCH", data)
	jdyError.CheckError(err,true)

	helper.FormatOutPut(re)
}

func (p *JinDouYunProxy) Set() {
	if p.Address == "" {
		p.Handle("Address", p.SetAddress)
	}
	if p.AppId == "" {
		p.Handle("AppId", p.SetAppId)
	}
	if p.AppKey == "" {
		p.Handle("AppKey", p.SetAppKey)
	}
	if p.ProxyAddress == "" {
		p.Handle("代理地址(Proxy)",p.setProxyAddress)
	}
}

func (p *JinDouYunProxy) SetUpdate() {
	if p.Address == "" {
		p.Handle("Address", p.SetAddress)
	}
	if p.AppId == "" {
		p.Handle("AppId", p.SetAppId)
	}
	if p.AppKey == "" {
		p.Handle("AppKey", p.SetAppKey)
	}
	if p.Id == "" {
		p.Handle("Id",p.SetId)
	}
	if p.ProxyAddress == "" {
		p.Handle("代理地址(Proxy)",p.setProxyAddress)
	}
}

func (p *JinDouYunProxy) setProxyAddress(address interface{}) {
	p.ProxyAddress = address.(string)
}

func (p *JinDouYunProxy) Clear() {
	p.ProxyAddress = ""
	p.Id = ""
}