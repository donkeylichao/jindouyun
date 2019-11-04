package config

import (
	"../jdyError"
	"bufio"
	"os"
	"fmt"
	"strings"
	"../util"
	"encoding/json"
	"../data"
)

/**
人保配置
 */
type JinDouYunEpicc struct {
	JinDouYunConfig
	Id               string `json:"id"`
	ProxyId          string `json:"proxy_id"`
	BusinessSrc      string `json:"business_src"`
	SalesChannel     string `json:"business_src"`
	BelongOrg        string `json:"belong_org"`
	BelongPerson     string `json:"belong_person"`
	Channel          string `json:"channel"`
	BelongDepartment string `json:"belong_department"`
	OperatorNo       string `json:"operator_no"`
	AgentPoint       string `json:"agent_point"`
	Xsryzyzhm        string `json:"xsryzyzhm"`
	StaffCode        string `json:"staff_code"`
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
		case "4":
			epicc.update()
		case "5":
			epicc.delete()
		case "1":
			epicc.list()
		case "2":
			epicc.detail()
		case "6":
			println("取消操作")
			os.Exit(0)
		default:
			continue
		}
	}
}

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
	if epicc.SalesChannel == "" {
		epicc.handle("销售渠道(SalesChannel)", epicc.setSalesChannel)
	}
	if epicc.BelongOrg == "" {
		epicc.handle("出单机构(BelongOrg)", epicc.setBelongOrg)
	}
	if epicc.BelongDepartment == "" {
		epicc.handle("归属部门(BelongDepartment)", epicc.setBelongDepartment)
	}
	if epicc.BelongPerson == "" {
		epicc.handle("归属人(BelongPerson)", epicc.setBelongPerson)
	}
	if epicc.OperatorNo == "" {
		epicc.handle("经办人(OperatorNo)", epicc.setOperatorNo)
	}
	if epicc.BusinessSrc == "" {
		epicc.handle("业务来源(BusinessSrc)", epicc.setBusinessSrc)
	}
	if epicc.StaffCode == "" {
		epicc.handle("验车人(StaffCode)", epicc.setStaffCode)
	}
	if epicc.AgentPoint == "" {
		epicc.handle("渠道代码(AgentPoint)", epicc.setAgentPoint)
	}
	if epicc.Xsryzyzhm == "" {
		epicc.handle("中介机构销售人员名称(Xsryzyzhm)",epicc.setXsryzyzhm)
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

func (epicc *JinDouYunEpicc) setProxyId(proxyId interface{}) {
	epicc.ProxyId = proxyId.(string)
}

func (epicc *JinDouYunEpicc) setUser(user interface{}) {
	epicc.User = user.(string)
}

func (epicc *JinDouYunEpicc) setPass(pass interface{}) {
	epicc.Pass = pass.(string)
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

func (epicc *JinDouYunEpicc) setSalesChannel(salesChannel interface{})  {
	util := util.GetJinDouYunUtil(epicc.Address, epicc.AppId, epicc.AppKey)

	query := map[string]string{}
	query["apiUrl"] = "ic-accounts/" + epicc.CityCode + "/epicc/sales_channel/options"
	query["user"] = epicc.User
	query["pass"] = epicc.Pass
	query["proxy_id"] = epicc.ProxyId

	re, err := util.IcAccountsOptions(query)
	jdyError.CheckError(err, true)

	reData := data.OptionsData{}
	json.Unmarshal(re, &reData)
	fmt.Printf("data:%s\n",reData)
	os.Exit(0)
	for _, v := range reData.Options {
		if v.Value == salesChannel.(string) || v.Text == salesChannel.(string) {
			epicc.SalesChannel = v.Value
			fmt.Printf("销售渠道(SalesChannel)的值为:%s\n", epicc.SalesChannel)
			return
		}
	}
}

func (epicc *JinDouYunEpicc) setBelongDepartment(belongDepartment interface{})  {
	util := util.GetJinDouYunUtil(epicc.Address, epicc.AppId, epicc.AppKey)

	query := map[string]string{}
	query["apiUrl"] = "ic-accounts/" + epicc.CityCode + "/epicc/belong_department/options"
	query["user"] = epicc.User
	query["pass"] = epicc.Pass
	query["proxy_id"] = epicc.ProxyId

	re, err := util.IcAccountsOptions(query)
	jdyError.CheckError(err, true)

	reData := data.OptionsData{}
	json.Unmarshal(re, &reData)
	for _, v := range reData.Options {
		if v.Value == belongDepartment.(string) {
			epicc.BelongDepartment = v.Value
			fmt.Printf("归属部门(BelongDepartment)的值为:%s\n", epicc.BelongDepartment)
			return
		}
		if v.Text == belongDepartment.(string) {
			epicc.BelongDepartment = v.Value
			fmt.Printf("归属部门(BelongDepartment)的值为:%s\n", epicc.BelongDepartment)
			return
		}
	}
}

func (epicc *JinDouYunEpicc) setBelongPerson(belongPerson interface{}) {
	util := util.GetJinDouYunUtil(epicc.Address, epicc.AppId, epicc.AppKey)

	query := map[string]string{}
	query["apiUrl"] = "ic-accounts/" + epicc.CityCode + "/epicc/belong_person/options"
	query["user"] = epicc.User
	query["pass"] = epicc.Pass
	query["proxy_id"] = epicc.ProxyId
	query["belong_department"] = epicc.BelongDepartment

	re, err := util.IcAccountsOptions(query)
	jdyError.CheckError(err, true)

	reData := data.OptionsData{}
	json.Unmarshal(re, &reData)
	for _, v := range reData.Options {
		if v.Value == belongPerson.(string) {
			epicc.BelongPerson = v.Value
			fmt.Printf("归属人(BelongPerson)的值为:%s\n", epicc.BelongPerson)
			return
		}
		if v.Text == belongPerson.(string) {
			epicc.BelongPerson = v.Value
			fmt.Printf("归属人(BelongPerson)的值为:%s\n", epicc.BelongPerson)
			return
		}
	}
}

func (epicc *JinDouYunEpicc) setOperatorNo(operatorNo interface{})  {
	util := util.GetJinDouYunUtil(epicc.Address, epicc.AppId, epicc.AppKey)

	query := map[string]string{}
	query["apiUrl"] = "ic-accounts/" + epicc.CityCode + "/epicc/operator_no/options"
	query["user"] = epicc.User
	query["pass"] = epicc.Pass
	query["proxy_id"] = epicc.ProxyId
	query["belong_department"] = epicc.BelongDepartment

	re, err := util.IcAccountsOptions(query)
	jdyError.CheckError(err, true)

	reData := data.OptionsData{}
	json.Unmarshal(re, &reData)
	for _, v := range reData.Options {
		if v.Value == operatorNo.(string) {
			epicc.OperatorNo = v.Value
			fmt.Printf("归属人(OperatorNo)的值为:%s\n", epicc.OperatorNo)
			return
		}
		if v.Text == operatorNo.(string) {
			epicc.OperatorNo = v.Value
			fmt.Printf("归属人(OperatorNo)的值为:%s\n", epicc.OperatorNo)
			return
		}
	}
}

func (epicc *JinDouYunEpicc) setBusinessSrc(value interface{})  {
	util := util.GetJinDouYunUtil(epicc.Address, epicc.AppId, epicc.AppKey)

	query := map[string]string{}
	query["apiUrl"] = "ic-accounts/" + epicc.CityCode + "/epicc/business_src/options"
	query["user"] = epicc.User
	query["pass"] = epicc.Pass
	query["proxy_id"] = epicc.ProxyId
	query["belong_department"] = epicc.BelongDepartment

	re, err := util.IcAccountsOptions(query)
	jdyError.CheckError(err, true)

	reData := data.OptionsData{}
	json.Unmarshal(re, &reData)
	for _, v := range reData.Options {
		if v.Value == value.(string) {
			epicc.BusinessSrc = v.Value
			fmt.Printf("业务来源(BusinessSrc)的值为:%s\n", epicc.BusinessSrc)
			return
		}
		if v.Text == value.(string) {
			epicc.BusinessSrc = v.Value
			fmt.Printf("业务来源(BusinessSrc)的值为:%s\n", epicc.BusinessSrc)
			return
		}
	}
}

func (epicc *JinDouYunEpicc) setStaffCode(value interface{})  {
	util := util.GetJinDouYunUtil(epicc.Address, epicc.AppId, epicc.AppKey)

	query := map[string]string{}
	query["apiUrl"] = "ic-accounts/" + epicc.CityCode + "/epicc/staff_code/options"
	query["user"] = epicc.User
	query["pass"] = epicc.Pass
	query["proxy_id"] = epicc.ProxyId
	query["belong_department"] = epicc.BelongDepartment

	re, err := util.IcAccountsOptions(query)
	jdyError.CheckError(err, true)

	reData := data.OptionsData{}
	json.Unmarshal(re, &reData)
	for _, v := range reData.Options {
		if v.Value == value.(string) {
			epicc.StaffCode = v.Value
			fmt.Printf("验车人(StaffCode)的值为:%s\n", epicc.StaffCode)
			return
		}
		if v.Text == value.(string) {
			epicc.StaffCode = v.Value
			fmt.Printf("验车人(StaffCode)的值为:%s\n", epicc.StaffCode)
			return
		}
	}
}

func (epicc *JinDouYunEpicc) setAgentPoint(input interface{}) {
	util := util.GetJinDouYunUtil(epicc.Address, epicc.AppId, epicc.AppKey)

	query := map[string]string{}
	query["apiUrl"] = "ic-accounts/" + epicc.CityCode + "/epicc/agent_point/options"
	query["user"] = epicc.User
	query["pass"] = epicc.Pass
	query["proxy_id"] = epicc.ProxyId
	query["belong_department"] = epicc.BelongDepartment
	query["business_src"] = epicc.BusinessSrc

	re, err := util.IcAccountsOptions(query)
	jdyError.CheckError(err, true)

	reData := data.OptionsData{}
	json.Unmarshal(re, &reData)
	for _, v := range reData.Options {
		if v.Value == input.(string) {
			epicc.AgentPoint = v.Value
			fmt.Printf("渠道代码(AgentPoint)的值为:%s\n", epicc.AgentPoint)
			return
		}
		if v.Text == input.(string) {
			epicc.AgentPoint = v.Value
			fmt.Printf("渠道代码(AgentPoint)的值为:%s\n", epicc.AgentPoint)
			return
		}
	}
}

func (epicc *JinDouYunEpicc) setXsryzyzhm(input interface{}) {
	util := util.GetJinDouYunUtil(epicc.Address, epicc.AppId, epicc.AppKey)

	query := map[string]string{}
	query["apiUrl"] = "ic-accounts/" + epicc.CityCode + "/epicc/xsryzyzhm/options"
	query["user"] = epicc.User
	query["pass"] = epicc.Pass
	query["proxy_id"] = epicc.ProxyId
	query["agent_point"] = epicc.AgentPoint
	query["business_src"] = epicc.BusinessSrc
	query["belong_org"] = epicc.BelongOrg

	re, err := util.IcAccountsOptions(query)
	jdyError.CheckError(err, true)

	reData := data.OptionsData{}
	json.Unmarshal(re, &reData)
	for _, v := range reData.Options {
		if v.Value == input.(string) {
			epicc.Xsryzyzhm = v.Value
			fmt.Printf("中介机构销售人员名称(Xsryzyzhm)的值为:%s\n", epicc.Xsryzyzhm)
			return
		}
		if v.Text == input.(string) {
			epicc.Xsryzyzhm = v.Value
			fmt.Printf("中介机构销售人员名称(Xsryzyzhm)的值为:%s\n", epicc.Xsryzyzhm)
			return
		}
	}
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

