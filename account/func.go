package account

import (
	"fmt"
	"os"
	"io/ioutil"
	"jindouyun/jdyError"
	"encoding/json"
)

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
设置公司
 */
func SetCompany(company string) int {
	switch company {
	case "1":
		fmt.Println("人保")
		return 1
	case "2":
		fmt.Println("太平洋")
		return 2
	case "3":
		fmt.Println("平安")
		return 3
	default:
		fmt.Println("不支持的保险公司")
		os.Exit(0)
		return 0
	}
}

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

		if jdyConf.Address == "" {
			jdyConf.Handle("Address", jdyConf.SetAddress)
		}
		if jdyConf.AppId == "" {
			jdyConf.Handle("AppId", jdyConf.SetAppId)
		}
		if jdyConf.AppKey == "" {
			jdyConf.Handle("AppKey", jdyConf.SetAppKey)
		}

		return jdyConf
	case 2:
		var jdyConf JinDouYunTpy
		conf, err := ReadAll("./config.json")
		jdyError.CheckError(err, true)
		err = json.Unmarshal(conf, &jdyConf)
		jdyError.CheckError(err, true)

		if jdyConf.Address == "" {
			jdyConf.Handle("Address", jdyConf.SetAddress)
		}
		if jdyConf.AppId == "" {
			jdyConf.Handle("AppId", jdyConf.SetAppId)
		}
		if jdyConf.AppKey == "" {
			jdyConf.Handle("AppKey", jdyConf.SetAppKey)
		}

		return jdyConf
	case 3:
		var jdyConf JinDouYunPinAn
		conf, err := ReadAll("./config.json")
		jdyError.CheckError(err, true)
		err = json.Unmarshal(conf, &jdyConf)
		jdyError.CheckError(err, true)

		if jdyConf.Address == "" {
			jdyConf.Handle("Address", jdyConf.SetAddress)
		}
		if jdyConf.AppId == "" {
			jdyConf.Handle("AppId", jdyConf.SetAppId)
		}
		if jdyConf.AppKey == "" {
			jdyConf.Handle("AppKey", jdyConf.SetAppKey)
		}

		return jdyConf
	}
	return nil
}
