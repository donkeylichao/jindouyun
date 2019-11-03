package main

import (
	//"jindouyun/config"
	"fmt"
	//"encoding/json"
	"jindouyun/jdyError"
	"bufio"
	"os"
	"strings"
	"jindouyun/company"
	"jindouyun/config"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("请选择公司:\n1.人保\n2.太平洋\n3.平安")
	input,err := inputReader.ReadString('\n')
	jdyError.CheckError(err,true)
	input = strings.TrimSpace(input)

	company := config.GetCompany(company.SetCompany(input))
	switch company.(type) {
	case config.JinDouYunEpicc:
		com := company.(config.JinDouYunEpicc)
		com.Run()
	case config.JinDouYunPinAn:
		com := company.(config.JinDouYunPinAn)
		com.Run()
	case config.JinDouYunTpy:
		com := company.(config.JinDouYunTpy)
		com.Run()
	default:
		jdyError.CheckError(jdyError.ErrCompanyNotExists, true)
	}
}
