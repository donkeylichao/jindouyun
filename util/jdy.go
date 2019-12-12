package util

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"time"
	"jindouyun/jdyError"
	"net/url"
	"sort"
	"crypto/md5"
	"io"
	"fmt"
	"crypto/hmac"
	"strings"
	"strconv"
	"crypto/sha256"
	"os"
)

type JinDouYunUtil struct {
	Url    string
	AppId  string
	AppKey string
	Accept string
}

var util JinDouYunUtil

func GetJinDouYunUtil(url, appId, appKey string) *JinDouYunUtil {
	if util.Url == "" ||
		util.AppId == "" ||
		util.AppKey == "" ||
		util.Accept == "" {
		return &JinDouYunUtil{url, appId, appKey, "application/vnd.botpy.v2+json"}
	}
	return &util
}

func (u *JinDouYunUtil) Request(query map[string]string, method string, data map[string]interface{}) ([]byte, error) {
	uri := u.Url + "/" + u.getApiQueryString(query)
	//println(uri[:len(uri)-1])
	//return nil,nil
	if u.AppKey == "" || u.AppId == "" || u.Url == "" {
		fmt.Println("address,appid,appkey 没有设置")
		os.Exit(0)
	}
	jsonStr, err := json.Marshal(data)

	req, err := http.NewRequest(method, uri[:len(uri)-1], bytes.NewBuffer(jsonStr))

	req.Header.Set("Accept", u.Accept)
	req.Header.Set("Authorization", "appid " + u.AppId)
	req.Header.Set("X-Yobee-Timestamp", strconv.FormatInt(time.Now().Unix(),10))
	req.Header.Set("X-Yobee-Signature", u.getStringToSign(query, data))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		jdyError.CheckError(err, false)
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func (u *JinDouYunUtil) getStringToSign(query map[string]string, data map[string]interface{}) string {
	timeStamp := time.Now().Unix()
	//fmt.Printf("timestamp:%s\n",strconv.FormatInt(timeStamp,10))
	paramString := u.getParamsString(query)
	//fmt.Printf("paramString:%s\n",paramString)
	paramMD5 := u.getParamsMD5(data)
	//fmt.Printf("paramMD5:%s\n",paramMD5)
	//fmt.Printf("query:%s\n",query)
	//获取签名串
	stringToSign := strconv.FormatInt(timeStamp, 10) + u.AppKey + u.Accept + u.Url + "/" + query["apiUrl"] + paramString + paramMD5 + u.AppId
	//fmt.Printf("stringToSign:%s\n",stringToSign)

	//sha256加密
	mac := hmac.New(sha256.New, []byte(u.AppKey))
	io.WriteString(mac,stringToSign)
	stringToSign = fmt.Sprintf("%x", mac.Sum(nil))
	//fmt.Printf("sha256:%s\n",stringToSign)

	//bin2hex
	//hex,err := hex.DecodeString(stringToSign)
	//jdyError.CheckError(err,false)
	//fmt.Printf("hex:%s\n",string(hex))

	//strtolower
	stringToSign = strings.ToLower(stringToSign)
	return stringToSign
}

func (u *JinDouYunUtil) getParamsMD5(data map[string]interface{}) string {
	jsonStr, err := json.Marshal(data)
	jdyError.CheckError(err,false)

	w := md5.New()
	io.WriteString(w, string(jsonStr))
	return  strings.ToLower(fmt.Sprintf("%x", w.Sum(nil)))
}

func (u *JinDouYunUtil) getParamsString(query map[string]string) string  {
	var newMp = make([]string, 0)
	for k, _ := range query {
		if k == "apiUrl" {
			continue
		}
		newMp = append(newMp, k)
	}
	sort.Strings(newMp)
	paramString := ""
	for _, v := range newMp {
		if query[v] == "" {
			continue
		}
		paramString += v + query[v]
	}
	return paramString
}

func (u *JinDouYunUtil) getApiQueryString(query map[string]string) string {
	queryParam := query["apiUrl"] + "?"
	for k,v := range query  {
		if v == "" {
			continue
		}
		if k == "apiUrl" {
			continue
		}
		str := k + "=" + url.QueryEscape(v)+"&"
		queryParam += str
	}
	return queryParam
}

