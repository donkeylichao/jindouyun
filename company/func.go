package company

import (
	"fmt"
	"os"
)

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
