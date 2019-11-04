package config

import (
	"os"
	"io/ioutil"
	"encoding/json"
	"../jdyError"
	"fmt"
	"bufio"
	"strings"
	"../util"
)

type JinDouYunConfig struct {
	Address   string `json:"address"`
	AppId     string `json:"app_id"`
	AppKey    string `json:"app_key"`
	User      string `json:"user"`
	Pass      string `json:"pass"`
	CityCode  string `json:"city_code"`
	IsDefault bool   `json:"is_default"`
}

/**
获取配置文件方法
 */
func ReadAll(filePath string) ([]byte, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(f)
}

/**
参数赋值函数类型
 */
type HandleFunc func(interface{})

/**
获取保险公司方法
 */
func GetCompany(company int) interface{} {
	switch company {
	case 1:
		var jdyConf JinDouYunEpicc
		conf, err := ReadAll("./config.json")
		jdyError.CheckError(err, true)
		err = json.Unmarshal(conf, &jdyConf)
		jdyError.CheckError(err, true)
		return jdyConf
	case 2:
		var jdyConf JinDouYunTpy
		conf, err := ReadAll("./config.json")
		jdyError.CheckError(err, true)
		err = json.Unmarshal(conf, &jdyConf)
		jdyError.CheckError(err, true)
		return jdyConf
	case 3:
		var jdyConf JinDouYunPinAn
		conf, err := ReadAll("./config.json")
		jdyError.CheckError(err, true)
		err = json.Unmarshal(conf, &jdyConf)
		jdyError.CheckError(err, true)
		return jdyConf
	}
	return nil
}

/**
获取列表
 */
func (conf *JinDouYunConfig) getList() {
	util := util.GetJinDouYunUtil(conf.Address, conf.AppId, conf.AppKey)

	query := map[string]string{}
	query["apiUrl"] = "accounts"

	re, err := util.Account(query, "GET", nil)
	jdyError.CheckError(err, true)
	fmt.Printf("%s\n", re)
}

/**
获取详情
 */
func (conf *JinDouYunConfig) detail()  {

}

/**
删除操作
 */

func (conf *JinDouYunConfig) delete() {

}

/**
处理参数赋值
 */
func (conf *JinDouYunConfig) handle(lab string, handler HandleFunc) {

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
