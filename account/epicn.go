package account

import (
	"bufio"
	"os"
	"fmt"
	"jindouyun/jdyError"
	"strings"
	"jindouyun/util"
	"jindouyun/helper"
)

/**
太平洋配置
 */
type JinDouYunTpy struct {
	JinDouYunConfig
	User      string `json:"user"`
	Pass      string `json:"pass"`
	IsDefault bool   `json:"is_default"`
}

func (tpy *JinDouYunTpy) Run() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("选择操作:\n1.列表\n2.详情\n3.添加\n4.编辑\n5.删除\n6.取消操作")
	for {
		input, err := inputReader.ReadString('\n')
		jdyError.CheckError(err, true)
		input = strings.TrimSpace(input)

		switch input {
		case "3":
			tpy.add()
		case "4":
			tpy.update()
		case "5":
			tpy.delete()
		case "1":
			tpy.list()
		case "2":
			tpy.detail()
		case "6":
			println("取消操作")
			os.Exit(0)
		default:
			continue
		}
	}
}

/**
获取列表
 */
func (tpy *JinDouYunTpy) list() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("请选择:\n1.查看\n2.返回\n3.退出")
	for {
		input, err := inputReader.ReadString('\n')
		jdyError.CheckError(err, true)
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			tpy.getList()
			os.Exit(0)
		case "2":
			tpy.Run()
		case "3":
			fmt.Println("退出")
			os.Exit(0)
		default:
			continue
		}
	}
}

func (tpy *JinDouYunTpy) add()  {
	tpy.Set()
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
			tpy.postAccount()
			os.Exit(0)
		case "2":
			tpy.add()
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
func (tpy *JinDouYunTpy) update() {
	tpy.Set()
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
			tpy.updateAccount()
			os.Exit(0)
		case "2":
			tpy.update()
		case "3":
			fmt.Println("已取消操作")
			os.Exit(0)
		default:
			continue
		}
	}
}

/**
添加账号
 */
func (tpy *JinDouYunTpy) postAccount() {
	util := util.GetJinDouYunUtil(tpy.Address, tpy.AppId, tpy.AppKey)
	//fmt.Printf("%s\n", util)

	query := map[string]string{}
	query["apiUrl"] = "accounts"

	data := map[string]interface{}{}
	data["account"] = map[string]string{
		"user":tpy.User,
		"pass":tpy.Pass,
		//"belong_org":tpy.BelongOrg,
		//"belong_department":tpy.BelongDepartment,
		//"belong_person":tpy.BelongPerson,
		//"operator_no":tpy.OperatorNo,
		//"business_src":tpy.BusinessSrc,
		//"agent_point":tpy.AgentPoint,
		//"xsryzyzhm":tpy.Xsryzyzhm,
		//"channel":tpy.Channel,
		//"proxy_id":tpy.ProxyId,
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
	data["ic_code"] = "epicn"
	data["city_code"] = tpy.CityCode

	re, err := util.Request(query, "POST", data)
	jdyError.CheckError(err, true)
	helper.FormatOutPut(re)
}

/**
编辑账号
 */
func (tpy *JinDouYunTpy) updateAccount() {

}

/**
设置各种参数
 */
func (tpy *JinDouYunTpy) Set() {

}
