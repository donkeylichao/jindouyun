package account

import (
	"os"
	"../jdyError"
	"fmt"
	"bufio"
	"strings"
	"../util"
)

type JinDouYunConfig struct {
	Address  string `json:"address"`
	AppId    string `json:"app_id"`
	AppKey   string `json:"app_key"`
	CityCode string `json:"city_code"`
	Id       string `json:"id"`
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
	fmt.Printf("%s\n", re)
}

/**
获取详情
 */
func (conf *JinDouYunConfig) detail() {

}

/**
删除操作
 */

func (conf *JinDouYunConfig) delete() {

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
