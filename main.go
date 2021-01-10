package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"jindouyun/account"
	"jindouyun/jdyError"
	"jindouyun/prompt"
	"jindouyun/proxy"
	"os"
	"strings"
)

func main() {
	prompt.Info("请选择操作:\n1.代理\n2.账号\n3.元数据")
	inputScan := bufio.NewScanner(os.Stdin)
	for inputScan.Scan() {
		input := inputScan.Text()
		switch input {
		case "1":
			proxyHandle()
		case "2":
			accountHandle()
		case "3":
			metaHandle()
		default:
			jdyError.CheckError(jdyError.ErrNotSupportHandle, true)
		}
	}
}

func proxyHandle() {
	inputReader := bufio.NewScanner(os.Stdin)
	prompt.Info("请选择操作:\n1.列表\n2.添加\n3.编辑")

	var jdyProxy proxy.JinDouYunProxy
	conf, err := account.ReadAll("config.json")
	jdyError.CheckError(err, true)
	err = json.Unmarshal(conf, &jdyProxy)
	jdyError.CheckError(err, true)

	if jdyProxy.Address == "" {
		jdyProxy.Handle("Address", jdyProxy.SetAddress)
	}
	if jdyProxy.AppId == "" {
		jdyProxy.Handle("AppId", jdyProxy.SetAppId)
	}
	if jdyProxy.AppKey == "" {
		jdyProxy.Handle("AppKey", jdyProxy.SetAppKey)
	}

	for inputReader.Scan() {
		input := inputReader.Text()
		switch input {
		case "1":
			jdyProxy.List()
		case "2":
			jdyProxy.Add()
		case "3":
			jdyProxy.Update()
		default:
			jdyError.CheckError(jdyError.ErrNotSupportHandle, true)
		}
	}
}

func accountHandle() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Print("请选择公司:\n1.人保\n2.太平洋\n3.平安\n> ")
	input, err := inputReader.ReadString('\n')
	jdyError.CheckError(err, true)
	input = strings.TrimSpace(input)

	company := account.GetCompany(account.SetCompany(input))
	switch company.(type) {
	case account.JinDouYunEpicc:
		com := company.(account.JinDouYunEpicc)
		com.Run()
	case account.JinDouYunPinAn:
		com := company.(account.JinDouYunPinAn)
		com.Run()
	case account.JinDouYunTpy:
		com := company.(account.JinDouYunTpy)
		com.Run()
	default:
		jdyError.CheckError(jdyError.ErrCompanyNotExists, true)
	}
}

func metaHandle()  {
	
}
