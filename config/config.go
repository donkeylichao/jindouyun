package config

import (
	"os"
	"io/ioutil"
	"encoding/json"
	"jindouyun/jdyError"
)

type JinDouYunConfig struct {
	Id          string `json:"id"`
	Address     string `json:"address"`
	AppId       string `json:"app_id"`
	AppKey      string `json:"app_key"`
	User        string `json:"user"`
	Pass        string `json:"pass"`
	ProxyId     string `json:"proxy_id"`
	BusinessSrc string `json:"business_src"`
	CityCode    string `json:"city_code"`
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
		//conf, err := ReadAll("./config.json")
		//jdyError.CheckError(err, true)
		//err = json.Unmarshal(conf, &jdyConf)
		//jdyError.CheckError(err, true)
		return jdyConf
	case 3:
		var jdyConf JinDouYunPinAn
		//conf, err := ReadAll("./config.json")
		//jdyError.CheckError(err, true)
		//err = json.Unmarshal(conf, &jdyConf)
		//jdyError.CheckError(err, true)
		return jdyConf
	}
	return nil
}

