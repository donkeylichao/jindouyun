package account

import (
	"os"
	"jindouyun/jdyError"
	"fmt"
	"bufio"
	"strings"
	"jindouyun/util"
	"jindouyun/helper"
)

type JinDouYunConfig struct {
	Address   string `json:"address"`
	AppId     string `json:"app_id"`
	AppKey    string `json:"app_key"`
	CityCode  string `json:"city_code"`
	Id        string `json:"id"`
	IsDefault bool   `json:"is_default"`
}

/**
参数赋值函数类型
 */
type HandleFunc func(interface{})

/**
获取列表
 */
func (conf *JinDouYunConfig) getList() {
	util := util.GetJinDouYunUtil(conf.Address, conf.AppId, conf.AppKey)

	query := map[string]string{}
	query["apiUrl"] = "accounts"

	re, err := util.Request(query, "GET", nil)
	jdyError.CheckError(err, true)

	helper.FormatOutPut(re)
}

/**
获取详情
 */
func (conf *JinDouYunConfig) detail() {
	conf.Handle("id", conf.SetId)
	util := util.GetJinDouYunUtil(conf.Address, conf.AppId, conf.AppKey)
	query := map[string]string{}
	query["apiUrl"] = "accounts/"+ conf.Id

	re, err := util.Request(query, "GET", nil)
	jdyError.CheckError(err, true)
	helper.FormatOutPut(re)
}

/**
删除操作
 */
func (conf *JinDouYunConfig) delete() {
	conf.Handle("id", conf.SetId)
	util := util.GetJinDouYunUtil(conf.Address, conf.AppId, conf.AppKey)
	query := map[string]string{}
	query["apiUrl"] = "accounts/"+ conf.Id

	re, err := util.Request(query, "DELETE", nil)
	jdyError.CheckError(err, true)
	helper.FormatOutPut(re)
}

/**
处理参数赋值
 */
func (conf *JinDouYunConfig) Handle(lab string, handler HandleFunc) {

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

func (conf *JinDouYunConfig) SetCityCode() {
	fmt.Println("请选择城市编码的值:\n1.佛山\n2.东莞")
	inputReader := bufio.NewReader(os.Stdin)
	input, err := inputReader.ReadString('\n')
	jdyError.CheckError(err, false)
	input = strings.TrimSpace(input)

	switch input {
	case "1":
		conf.CityCode = "440600"
	case "2":
		conf.CityCode = "441900"
	default:
		fmt.Println("城市编码输入错误")
		conf.SetCityCode()
	}
}

func (conf *JinDouYunConfig) SetIsDefault() {
	fmt.Println("是否为设置为默认账号:\n1.是\n2.否")
	inputReader := bufio.NewReader(os.Stdin)
	input, err := inputReader.ReadString('\n')
	jdyError.CheckError(err, false)
	input = strings.TrimSpace(input)

	switch input {
	case "1":
		conf.IsDefault = true
	case "2":
		conf.IsDefault = false
	default:
		fmt.Println("输入错误")
		conf.SetIsDefault()
	}
}

func (conf *JinDouYunConfig) SetAddress(address interface{}) {
	conf.Address = address.(string)
}

func (conf *JinDouYunConfig) SetAppId(appid interface{}) {
	conf.AppId = appid.(string)
}

func (conf *JinDouYunConfig) SetAppKey(appKey interface{}) {
	conf.AppKey = appKey.(string)
}

func (conf *JinDouYunConfig) SetId(id interface{}) {
	conf.Id = id.(string)
}
