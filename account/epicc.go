package account

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
	User             string `json:"user"`
	Pass             string `json:"pass"`
	Lock             bool   `json:"lock"`
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
	Saler            string `json:"saler"`
	AgentType        string `json:"agent_type"`
	Comment          string `json:"comment"`
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

func (epicc *JinDouYunEpicc) add()  {
	epicc.Set()
	//fmt.Printf("%s\n", epicc)
	//return
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
			epicc.Clear()
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
	//fmt.Printf("%s\n", epicc)
	epicc.Handle("id", epicc.SetId)
	for epicc.Id == "" {
		epicc.Handle("id", epicc.SetId)
	}
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("请选择修改项:\n1.用户名(user)\n2.密码(pass)\n3.代理(proxy_id)\n4.城市(city_code)\n" +
		"5.出单机构(BelongOrg)\n6.归属部门(BelongDepartment)\n7.业务来源(BusinessSrc)\n" +
			"8.归属人(BelongPerson)\n9.经办人(OperatorNo)\n10.验车人(StaffCode)\n11.中介机构销售人员名称(Xsryzyzhm)\n" +
				"12.渠道代码(AgentPoint)\n13.项目代码(AgentType)\n14.销售渠道(SalesChannel)\n15.设置默认账号(IsDefault)")

	input, err := inputReader.ReadString('\n')
	jdyError.CheckError(err, true)
	input = strings.TrimSpace(input)

	switch input {
	case "1":
		epicc.Handle("用户名(User)", epicc.setUser)
	case "2":
		epicc.Handle("密码(Pass)",epicc.setPass)
	case "3":
		epicc.Handle("代理(ProxyId)",epicc.setProxyId)
	case "4":
		epicc.SetCityCode()
	case "5":
		epicc.SetCityCode()
		epicc.Handle("用户名(User)", epicc.setUser)
		epicc.Handle("密码(Pass)",epicc.setPass)
		epicc.Handle("代理(ProxyId)",epicc.setProxyId)
		epicc.Handle("出单机构(BelongOrg)",epicc.setBelongOrg)
	case "6":
		epicc.SetCityCode()
		epicc.Handle("用户名(User)", epicc.setUser)
		epicc.Handle("密码(Pass)",epicc.setPass)
		epicc.Handle("代理(ProxyId)",epicc.setProxyId)
		epicc.Handle("归属部门(BelongDepartment)",epicc.setBelongDepartment)
	case "7":
		epicc.SetCityCode()
		epicc.Handle("用户名(User)", epicc.setUser)
		epicc.Handle("密码(Pass)",epicc.setPass)
		epicc.Handle("代理(ProxyId)",epicc.setProxyId)
		epicc.Handle("归属部门(BelongDepartment)",epicc.setBelongDepartment)
		epicc.Handle("业务来源(BusinessSrc)",epicc.setBusinessSrc)
	case "8":
		epicc.SetCityCode()
		epicc.Handle("用户名(User)", epicc.setUser)
		epicc.Handle("密码(Pass)",epicc.setPass)
		epicc.Handle("代理(ProxyId)",epicc.setProxyId)
		epicc.Handle("归属部门(BelongDepartment)",epicc.setBelongDepartment)
		epicc.Handle("归属人(BelongPerson)",epicc.setBelongPerson)
	case "9":
		epicc.SetCityCode()
		epicc.Handle("用户名(User)", epicc.setUser)
		epicc.Handle("密码(Pass)",epicc.setPass)
		epicc.Handle("代理(ProxyId)",epicc.setProxyId)
		epicc.Handle("归属部门(BelongDepartment)",epicc.setBelongDepartment)
		epicc.Handle("经办人(OperatorNo)",epicc.setOperatorNo)
	case "10":
		epicc.SetCityCode()
		epicc.Handle("用户名(User)", epicc.setUser)
		epicc.Handle("密码(Pass)",epicc.setPass)
		epicc.Handle("代理(ProxyId)",epicc.setProxyId)
		epicc.Handle("归属部门(BelongDepartment)",epicc.setBelongDepartment)
		epicc.Handle("验车人(StaffCode)",epicc.setStaffCode)
	case "11":
		epicc.SetCityCode()
		epicc.Handle("用户名(User)", epicc.setUser)
		epicc.Handle("密码(Pass)",epicc.setPass)
		epicc.Handle("代理(ProxyId)",epicc.setProxyId)
		epicc.Handle("归属部门(BelongDepartment)",epicc.setBelongDepartment)
		epicc.Handle("业务来源(BusinessSrc)",epicc.setBusinessSrc)
		epicc.Handle("渠道代码(AgentPoint)",epicc.setAgentPoint)
		epicc.Handle("中介机构销售人员名称(Xsryzyzhm)",epicc.setXsryzyzhm)
	case "12":
		epicc.SetCityCode()
		epicc.Handle("用户名(User)", epicc.setUser)
		epicc.Handle("密码(Pass)",epicc.setPass)
		epicc.Handle("代理(ProxyId)",epicc.setProxyId)
		epicc.Handle("归属部门(BelongDepartment)",epicc.setBelongDepartment)
		epicc.Handle("业务来源(BusinessSrc)",epicc.setBusinessSrc)
		epicc.Handle("渠道代码(AgentPoint)",epicc.setAgentPoint)
	case "13":
		epicc.Handle("项目代码(AgentType)",epicc.setAgentType)
	case "14":
		epicc.SetCityCode()
		epicc.Handle("用户名(User)", epicc.setUser)
		epicc.Handle("密码(Pass)",epicc.setPass)
		epicc.Handle("代理(ProxyId)",epicc.setProxyId)
		epicc.Handle("销售渠道(SalesChannel)",epicc.setSalesChannel)
	case "15":
		epicc.setIsDefault()
		epicc.setDefault()
		os.Exit(0)
	case "20":
		fmt.Println("已取消操作")
		os.Exit(0)
	default:
		fmt.Println("非法操作")
		os.Exit(0)
	}
	epicc.updateAccount()
}

/**
设置各种参数
 */
func (epicc *JinDouYunEpicc) Set() {
	if epicc.CityCode == "" {
		epicc.SetCityCode()
	}
	epicc.SetIsDefault()
	if epicc.User == "" {
		epicc.Handle("User", epicc.setUser)
	}
	if epicc.Pass == "" {
		epicc.Handle("Pass", epicc.setPass)
	}
	if epicc.ProxyId == "" {
		epicc.Handle("代理ID(ProxyId)", epicc.setProxyId)
	}
	if epicc.SalesChannel == "" {
		epicc.Handle("销售渠道(SalesChannel)", epicc.setSalesChannel)
	}
	if epicc.BelongOrg == "" {
		epicc.Handle("出单机构(BelongOrg)", epicc.setBelongOrg)
	}
	if epicc.BelongDepartment == "" {
		epicc.Handle("归属部门(BelongDepartment)", epicc.setBelongDepartment)
	}
	if epicc.BelongPerson == "" {
		epicc.Handle("归属人(BelongPerson)", epicc.setBelongPerson)
	}
	if epicc.OperatorNo == "" {
		epicc.Handle("经办人(OperatorNo)", epicc.setOperatorNo)
	}
	if epicc.BusinessSrc == "" {
		epicc.Handle("业务来源(BusinessSrc)", epicc.setBusinessSrc)
	}
	if epicc.StaffCode == "" {
		epicc.Handle("验车人(StaffCode)", epicc.setStaffCode)
	}
	if epicc.AgentPoint == "" {
		epicc.Handle("渠道代码(AgentPoint)", epicc.setAgentPoint)
	}
	if epicc.Xsryzyzhm == "" {
		epicc.Handle("中介机构销售人员名称(Xsryzyzhm)",epicc.setXsryzyzhm)
	}
	if epicc.Saler == "" {
		epicc.Handle("推荐送修代码(Saler)", epicc.setSaler)
	}
	if epicc.AgentType == "" {
		epicc.Handle("项目代码(AgentType)", epicc.setAgentType)
	}
	if epicc.Comment == "" {
		epicc.Handle("备注(Comment)", epicc.setComment)
	}

}

func (epicc *JinDouYunEpicc) Clear()  {
	epicc.CityCode = ""
	epicc.User = ""
	epicc.Pass = ""
	epicc.ProxyId = ""
	epicc.SalesChannel = ""
	epicc.BelongOrg = ""
	epicc.BelongDepartment = ""
	epicc.BelongPerson = ""
	epicc.OperatorNo = ""
	epicc.BusinessSrc = ""
	epicc.StaffCode = ""
	epicc.AgentPoint = ""
	epicc.Xsryzyzhm = ""
	epicc.Id = ""
	epicc.Channel = ""
	epicc.Saler = ""
	epicc.AgentType = ""
	epicc.Comment = ""
}
/**********************参数赋值方法*************************/

func (epicc *JinDouYunEpicc) setProxyId(proxyId interface{}) {
	epicc.ProxyId = proxyId.(string)
}

func (epicc *JinDouYunEpicc) setUser(user interface{}) {
	epicc.User = user.(string)
}

func (epicc *JinDouYunEpicc) setPass(pass interface{}) {
	epicc.Pass = pass.(string)
}
func (epicc *JinDouYunEpicc) setComment(comment interface{}) {
	epicc.Comment = comment.(string)
}
func (epicc *JinDouYunEpicc) setSaler(saler interface{}) {
	epicc.Saler = saler.(string)
}
func (epicc *JinDouYunEpicc) setAgentType(agentType interface{}) {
	epicc.AgentType = agentType.(string)
}

func (epicc *JinDouYunEpicc) setBelongOrg(belong interface{}) {
	util := util.GetJinDouYunUtil(epicc.Address, epicc.AppId, epicc.AppKey)

	query := map[string]string{}
	query["apiUrl"] = "ic-accounts/" + epicc.CityCode + "/epicc/belong_org/options"
	query["user"] = epicc.User
	query["pass"] = epicc.Pass
	query["proxy_id"] = epicc.ProxyId

	re, err := util.Request(query,"GET",nil)
	jdyError.CheckError(err, true)

	reData := data.OptionsData{}
	json.Unmarshal(re, &reData)

	for _,v := range reData.Options {
		if v.Value == belong.(string) || v.Text == belong.(string){
			epicc.BelongOrg = v.Value
			fmt.Printf("出单机构(BelongOrg)的值为:%s\n",epicc.BelongOrg)
			return
		}
	}
	fmt.Println("出单机构(BelongOrg)的值未获取到")
}

func (epicc *JinDouYunEpicc) setSalesChannel(salesChannel interface{})  {
	util := util.GetJinDouYunUtil(epicc.Address, epicc.AppId, epicc.AppKey)

	query := map[string]string{}
	query["apiUrl"] = "ic-accounts/" + epicc.CityCode + "/epicc/sales_channel/options"
	query["user"] = epicc.User
	query["pass"] = epicc.Pass
	query["proxy_id"] = epicc.ProxyId

	re, err := util.Request(query,"GET",nil)
	jdyError.CheckError(err, true)

	reData := data.OptionsData{}
	json.Unmarshal(re, &reData)

	//os.Exit(0)
	for _, v := range reData.Options {
		if v.Value == salesChannel.(string) || v.Text == salesChannel.(string) {
			epicc.SalesChannel = v.Value
			fmt.Printf("销售渠道(SalesChannel)的值为:%s\n", epicc.SalesChannel)
			return
		}
	}
	fmt.Println("销售渠道(SalesChannel)的值未获取到")
}

func (epicc *JinDouYunEpicc) setBelongDepartment(belongDepartment interface{})  {
	util := util.GetJinDouYunUtil(epicc.Address, epicc.AppId, epicc.AppKey)

	query := map[string]string{}
	query["apiUrl"] = "ic-accounts/" + epicc.CityCode + "/epicc/belong_department/options"
	query["user"] = epicc.User
	query["pass"] = epicc.Pass
	query["proxy_id"] = epicc.ProxyId

	re, err := util.Request(query,"GET",nil)
	jdyError.CheckError(err, true)

	reData := data.OptionsData{}
	json.Unmarshal(re, &reData)
	for _, v := range reData.Options {
		if v.Value == belongDepartment.(string) || v.Text == belongDepartment.(string) {
			epicc.BelongDepartment = v.Value
			fmt.Printf("归属部门(BelongDepartment)的值为:%s\n", epicc.BelongDepartment)
			return
		}
	}
	fmt.Println("归属部门(BelongDepartment)的值未获取到")
}

func (epicc *JinDouYunEpicc) setBelongPerson(belongPerson interface{}) {
	util := util.GetJinDouYunUtil(epicc.Address, epicc.AppId, epicc.AppKey)

	query := map[string]string{}
	query["apiUrl"] = "ic-accounts/" + epicc.CityCode + "/epicc/belong_person/options"
	query["user"] = epicc.User
	query["pass"] = epicc.Pass
	query["proxy_id"] = epicc.ProxyId
	query["belong_department"] = epicc.BelongDepartment

	re, err := util.Request(query,"GET",nil)
	jdyError.CheckError(err, true)

	reData := data.OptionsData{}
	json.Unmarshal(re, &reData)
	for _, v := range reData.Options {
		if v.Value == belongPerson.(string) || v.Text == belongPerson.(string){
			epicc.BelongPerson = v.Value
			fmt.Printf("归属人(BelongPerson)的值为:%s\n", epicc.BelongPerson)
			return
		}
	}
	fmt.Println("归属人(BelongPerson)的值未获取到")
}

func (epicc *JinDouYunEpicc) setOperatorNo(operatorNo interface{})  {
	util := util.GetJinDouYunUtil(epicc.Address, epicc.AppId, epicc.AppKey)

	query := map[string]string{}
	query["apiUrl"] = "ic-accounts/" + epicc.CityCode + "/epicc/operator_no/options"
	query["user"] = epicc.User
	query["pass"] = epicc.Pass
	query["proxy_id"] = epicc.ProxyId
	query["belong_department"] = epicc.BelongDepartment

	re, err := util.Request(query,"GET",nil)
	jdyError.CheckError(err, true)

	reData := data.OptionsData{}
	json.Unmarshal(re, &reData)
	for _, v := range reData.Options {
		if v.Value == operatorNo.(string) || v.Text == operatorNo.(string) {
			epicc.OperatorNo = v.Value
			fmt.Printf("经办人(OperatorNo)的值为:%s\n", epicc.OperatorNo)
			return
		}
	}
	fmt.Println("经办人(OperatorNo)的值未获取到")
}

func (epicc *JinDouYunEpicc) setBusinessSrc(value interface{})  {
	util := util.GetJinDouYunUtil(epicc.Address, epicc.AppId, epicc.AppKey)

	query := map[string]string{}
	query["apiUrl"] = "ic-accounts/" + epicc.CityCode + "/epicc/business_src/options"
	query["user"] = epicc.User
	query["pass"] = epicc.Pass
	query["proxy_id"] = epicc.ProxyId
	query["belong_department"] = epicc.BelongDepartment

	re, err := util.Request(query,"GET",nil)
	jdyError.CheckError(err, true)

	reData := data.OptionsData{}
	json.Unmarshal(re, &reData)
	for _, v := range reData.Options {
		if v.Value == value.(string) || v.Text == value.(string) {
			epicc.BusinessSrc = v.Value
			fmt.Printf("业务来源(BusinessSrc)的值为:%s\n", epicc.BusinessSrc)
			return
		}
	}
	fmt.Println("业务来源(BusinessSrc)的值未获取到")
}

func (epicc *JinDouYunEpicc) setStaffCode(value interface{})  {
	util := util.GetJinDouYunUtil(epicc.Address, epicc.AppId, epicc.AppKey)

	query := map[string]string{}
	query["apiUrl"] = "ic-accounts/" + epicc.CityCode + "/epicc/staff_code/options"
	query["user"] = epicc.User
	query["pass"] = epicc.Pass
	query["proxy_id"] = epicc.ProxyId
	query["belong_department"] = epicc.BelongDepartment

	re, err := util.Request(query,"GET",nil)
	jdyError.CheckError(err, true)

	reData := data.OptionsData{}
	json.Unmarshal(re, &reData)
	for _, v := range reData.Options {
		if v.Value == value.(string) || v.Text == value.(string){
			epicc.StaffCode = v.Value
			fmt.Printf("验车人(StaffCode)的值为:%s\n", epicc.StaffCode)
			return
		}
	}
	fmt.Println("验车人(StaffCode)的值未获取到")
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

	re, err := util.Request(query,"GET",nil)
	jdyError.CheckError(err, true)

	reData := data.OptionsData{}
	json.Unmarshal(re, &reData)
	for _, v := range reData.Options {
		if v.Value == input.(string) || v.Text == input.(string){
			epicc.AgentPoint = v.Value
			fmt.Printf("渠道代码(AgentPoint)的值为:%s\n", epicc.AgentPoint)
			return
		}
	}
	fmt.Println("渠道代码(AgentPoint)的值未获取到")
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

	re, err := util.Request(query,"GET",nil)
	jdyError.CheckError(err, true)

	reData := data.OptionsData{}
	json.Unmarshal(re, &reData)
	for _, v := range reData.Options {
		if v.Value == input.(string) || v.Text == input.(string){
			epicc.Xsryzyzhm = v.Value
			fmt.Printf("中介机构销售人员名称(Xsryzyzhm)的值为:%s\n", epicc.Xsryzyzhm)
			return
		}
	}
	fmt.Println("中介机构销售人员名称(Xsryzyzhm)的值未获取到")
}

func (epicc *JinDouYunEpicc) setIsDefault() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("是否设置默认:\n1.是\n2.否")

	input, err := inputReader.ReadString('\n')
	jdyError.CheckError(err, true)
	input = strings.TrimSpace(input)

	switch input {
	case "1":
		epicc.IsDefault = true
	case "2":
		epicc.IsDefault = false
	default:
		fmt.Println("输入值不合法")
		epicc.setIsDefault()
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
		"user":              epicc.User,
		"pass":              epicc.Pass,
		"belong_org":        epicc.BelongOrg,
		"belong_department": epicc.BelongDepartment,
		"belong_person":     epicc.BelongPerson,
		"operator_no":       epicc.OperatorNo,
		"business_src":      epicc.BusinessSrc,
		"agent_point":       epicc.AgentPoint,
		"xsryzyzhm":         epicc.Xsryzyzhm,
		"channel":           epicc.Channel,
		"proxy_id":          epicc.ProxyId,
		"sale_channel":      epicc.SalesChannel,
		"comment":           epicc.Comment,
		"agent_type":        epicc.AgentType,
		"saler":             epicc.Saler,
	}
	data["is_default"] = epicc.IsDefault
	data["lock"] = false
	data["disabled"] = false
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

	re, err := util.Request(query, "POST", data)
	jdyError.CheckError(err, true)
	fmt.Printf("%s\n", re)
}

/**
编辑账号
 */
func (epicc *JinDouYunEpicc) updateAccount() {
	util := util.GetJinDouYunUtil(epicc.Address, epicc.AppId, epicc.AppKey)
	//fmt.Printf("%s\n", util)
	query := map[string]string{}
	query["apiUrl"] = "accounts/"+ epicc.Id

	data := map[string]interface{}{}
	account := map[string]string{}

	if epicc.User != "" {
		account["user"] = epicc.User
	}
	if epicc.Pass != "" {
		account["pass"] = epicc.Pass
	}
	if epicc.BelongOrg != "" {
		account["belong_org"] = epicc.BelongOrg
	}
	if epicc.BelongDepartment != "" {
		account["belong_department"] = epicc.BelongDepartment
	}
	if epicc.BelongPerson != "" {
		account["belong_person"] = epicc.BelongPerson
	}
	if epicc.OperatorNo != "" {
		account["operator_no"] = epicc.OperatorNo
	}
	if epicc.BusinessSrc != "" {
		account["business_src"] = epicc.BusinessSrc
	}
	if epicc.AgentPoint != "" {
		account["agent_point"] = epicc.AgentPoint
	}
	if epicc.Xsryzyzhm != "" {
		account["xsryzyzhm"] = epicc.Xsryzyzhm
	}
	if epicc.Channel != "" {
		account["channel"] = epicc.Channel
	}
	if epicc.ProxyId != "" {
		account["proxy_id"] = epicc.ProxyId
	}
	if epicc.SalesChannel != "" {
		account["sale_channel"] = epicc.SalesChannel
	}
	if epicc.Comment != "" {
		account["comment"] = epicc.Comment
	}
	if epicc.AgentType != "" {
		account["agent_type"] = epicc.AgentType
	}
	if epicc.Saler != "" {
		account["saler"] = epicc.Saler
	}

	if len(account) != 0 {
		data["account"] = account
	}
	data["ic_code"] = "epicc"
	data["lock"] = false
	data["disabled"] = false
	//contant := map[string]string{
	//	"name":  "米米佛山人保测试",
	//	"email": "529755212@qq.com",
	//	"phone": "18201695833",
	//}
	//contactList := make([]map[string]string, 0)
	//contactList = append(contactList,contant)
	//data["contacts"] = contactList
	//data["ic_code"] = "epicc"
	if epicc.CityCode != "" {
		data["city_code"] = epicc.CityCode
	}

	re, err := util.Request(query, "PATCH", data)
	jdyError.CheckError(err, true)
	fmt.Printf("%s\n", re)
}

func (epicc *JinDouYunEpicc) setDefault() {
	util := util.GetJinDouYunUtil(epicc.Address, epicc.AppId, epicc.AppKey)
	//fmt.Printf("%s\n", util)
	query := map[string]string{}
	query["apiUrl"] = "accounts/"+ epicc.Id

	data := map[string]interface{}{}
	data["is_default"] = epicc.IsDefault

	re, err := util.Request(query, "PATCH", data)
	jdyError.CheckError(err, true)
	fmt.Printf("%s\n", re)
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

