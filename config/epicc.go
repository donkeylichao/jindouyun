package config

import (
	"jindouyun/jdyError"
	"bufio"
	"os"
	"fmt"
	"strings"
	"jindouyun/util"
	"encoding/json"
	"jindouyun/data"
)

/**
人保配置
 */
type JinDouYunEpicc struct {
	JinDouYunConfig
	BelongOrg        string
	BelongPerson     string
	Channel          string
	BelongDepartment string
	OperatorNo       string
	AgentPoint       string
	Xsryzyzhm        string
}

/**********************人保操作开始*************************/

func (epicc *JinDouYunEpicc) Run() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("选择操作:\n1.列表\n2.详情\n3.添加\n4.编辑\n5.删除\n6.取消操作")
	for {
		input, err := inputReader.ReadString('\n')
		jdyError.CheckError(err, true)
		input = strings.TrimSpace(input)

		switch input {
		case "3":
			epicc.add()
			os.Exit(0)
		case "4":
			epicc.update()
			os.Exit(0)
		case "5":
			epicc.delete()
			os.Exit(0)
		case "1":
			epicc.list()
			os.Exit(0)
		case "2":
			epicc.detail()
			os.Exit(0)
		case "6":
			println("取消操作")
			os.Exit(0)
		default:
			continue
		}
	}
}

/**
添加账号操作
 */
func (epicc *JinDouYunEpicc) add()  {
	epicc.Set()
	//fmt.Printf("%s\n", epicc)
	return
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("请确认:\n1.正确\n2.重新设置\n3.取消")
	for {
		input, err := inputReader.ReadString('\n')
		jdyError.CheckError(err, true)
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			fmt.Println("调用接口保存账号")
			epicc.postAccount()
			os.Exit(0)
		case "2":
			epicc.add()
		case "3":
			fmt.Println("已取消操作")
			os.Exit(0)
		default:
			continue
		}
	}
}

/**
编辑操作
 */
func (epicc *JinDouYunEpicc) update() {
	epicc.Set()
	//fmt.Printf("%s\n", epicc)
	return
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("请确认:\n1.正确\n2.重新设置\n3.取消")
	for {
		input, err := inputReader.ReadString('\n')
		jdyError.CheckError(err, true)
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			fmt.Println("调用接口保存账号")
			epicc.updateAccount()
			os.Exit(0)
		case "2":
			epicc.update()
		case "3":
			fmt.Println("已取消操作")
			os.Exit(0)
		default:
			continue
		}
	}
}

/**
获取列表
 */
func (epicc *JinDouYunEpicc) list() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("请选择:\n1.查看\n2.返回\n3.退出")
	for {
		input, err := inputReader.ReadString('\n')
		jdyError.CheckError(err, true)
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			epicc.getList()
			os.Exit(0)
		case "2":
			epicc.Run()
		case "3":
			fmt.Println("退出")
			os.Exit(0)
		default:
			continue
		}
	}
}

/**
获取详情
 */
func (epicc *JinDouYunEpicc) detail()  {
	
}

/**
删除操作
 */

func (epicc *JinDouYunEpicc) delete() {
	
}
 
/**
设置各种参数
 */
func (epicc *JinDouYunEpicc) Set() {
	if epicc.Address == "" {
		epicc.handle("Address", epicc.setAddress)
	}
	if epicc.AppId == "" {
		epicc.handle("AppId", epicc.setAppId)
	}
	if epicc.AppKey == "" {
		epicc.handle("AppKey", epicc.setAppKey)
	}
	if epicc.CityCode == "" {
		epicc.setCityCode()
	}
	if epicc.User == "" {
		epicc.handle("User", epicc.setUser)
	}
	if epicc.Pass == "" {
		epicc.handle("Pass", epicc.setPass)
	}
	if epicc.ProxyId == "" {
		epicc.handle("代理ID(ProxyId)", epicc.setProxyId)
	}
	if epicc.BelongOrg == "" {
		epicc.handle("出单机构(BelongOrg)", epicc.setBelongOrg)
	}

}

/**
处理参数赋值
 */
func (epicc *JinDouYunEpicc) handle(lab string, handler HandleFunc) {

	fmt.Printf("输入%s的值:\n", lab)
	inputReader := bufio.NewReader(os.Stdin)
	input, err := inputReader.ReadString('\n')
	jdyError.CheckError(err, false)
	input = strings.TrimSpace(input)

	switch input {
	case "":
		fmt.Printf("%s输入为空.\n", lab)
	default:
		fmt.Printf("%s输入为:%s\n", lab, input)
		handler(input)
	}

}

/**********************参数赋值方法*************************/

func (epicc *JinDouYunEpicc) setAddress(address interface{}) {
	epicc.Address = address.(string)
}

func (epicc *JinDouYunEpicc) setAppId(appid interface{}) {
	epicc.AppId = appid.(string)
}

func (epicc *JinDouYunEpicc) setAppKey(appKey interface{}) {
	epicc.AppKey = appKey.(string)
}

func (epicc *JinDouYunEpicc) setBelongOrg(belong interface{}) {
	util := util.GetJinDouYunUtil(epicc.Address, epicc.AppId, epicc.AppKey)

	query := map[string]string{}
	query["apiUrl"] = "ic-accounts/" + epicc.CityCode + "/epicc/belong_org/options"
	query["user"] = epicc.User
	query["pass"] = epicc.Pass
	query["proxy_id"] = epicc.ProxyId

	re, err := util.IcAccountsOptions(query)
	jdyError.CheckError(err, true)

	reData := data.OptionsData{}
	json.Unmarshal(re, &reData)

	for _,v := range reData.Options {
		if v.Value == belong.(string) {
			epicc.BelongOrg = v.Value
			fmt.Printf("出单机构(BelongOrg)的值为:%s\n",epicc.BelongOrg)
			return
		}
		if v.Text == belong.(string) {
			epicc.BelongOrg = v.Value
			fmt.Printf("出单机构(BelongOrg)的值为:%s\n",epicc.BelongOrg)
			return
		}
	}
}

func (epicc *JinDouYunEpicc) setCityCode() {
	fmt.Println("请选择城市编码的值:\n1.佛山\n2.东莞")
	inputReader := bufio.NewReader(os.Stdin)
	input, err := inputReader.ReadString('\n')
	jdyError.CheckError(err, false)
	input = strings.TrimSpace(input)

	switch input {
	case "1":
		epicc.CityCode = "440600"
	case "2":
		epicc.CityCode = "441900"
	default:
		fmt.Println("城市编码输入错误")
		epicc.setCityCode()
	}
}

func (epicc *JinDouYunEpicc) setProxyId(proxyId interface{}) {
	epicc.ProxyId = proxyId.(string)
}



func (epicc *JinDouYunEpicc) setUser(user interface{}) {
	epicc.User = user.(string)
}

func (epicc *JinDouYunEpicc) setPass(pass interface{}) {
	epicc.Pass = pass.(string)
}

/**
添加账号
 */
func (epicc *JinDouYunEpicc) postAccount() {
	util := util.GetJinDouYunUtil(epicc.Address, epicc.AppId, epicc.AppKey)
	//fmt.Printf("%s\n", util)

	query := map[string]string{}
	query["apiUrl"] = "accounts"

	data := map[string]interface{}{}
	data["account"] = map[string]string{
		"user":epicc.User,
		"pass":epicc.Pass,
		"belong_org":epicc.BelongOrg,
		"belong_department":epicc.BelongDepartment,
		"belong_person":epicc.BelongPerson,
		"operator_no":epicc.OperatorNo,
		"business_src":epicc.BusinessSrc,
		"agent_point":epicc.AgentPoint,
		"xsryzyzhm":epicc.Xsryzyzhm,
		"channel":epicc.Channel,
		"proxy_id":epicc.ProxyId,
	}
	data["is_default"] = true
	data["lock"] = false
	contant := map[string]string{
		"name":  "米米佛山人保测试",
		"email": "529755212@qq.com",
		"phone": "18201695833",
	}
	contactList := make([]map[string]string, 0)
	contactList = append(contactList,contant)
	data["contacts"] = contactList
	data["ic_code"] = "epicc"
	data["city_code"] = epicc.CityCode

	re, err := util.Account(query, "POST", data)
	jdyError.CheckError(err, true)
	fmt.Printf("%s\n", re)
}

/**
编辑账号
 */
func (epicc *JinDouYunEpicc) updateAccount() {

}

/**
获取列表
 */
func (epicc *JinDouYunEpicc) getList() {
	util := util.GetJinDouYunUtil(epicc.Address, epicc.AppId, epicc.AppKey)

	query := map[string]string{}
	query["apiUrl"] = "accounts"

	re, err := util.Account(query, "GET", nil)
	jdyError.CheckError(err, true)
	fmt.Printf("%s\n", re)
}
