package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 准备从标准输入读取数据
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("请输入你的名字:")
	// 读取数值知道碰到\n  为止
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occurred:", err)
		// 异常退出
		os.Exit(1)
	} else {
		// 用切片操作删除最后的\n
		name := input[:len(input)-1]
		fmt.Println("Hello," + name + "我可以为你做什么事情吗?")

	}
	for {
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occurred:", err)
			// 继续执行
			continue
		}
		input = input[:len(input)-1]
		input = strings.ToLower(input)
		switch input {
		case "":
			continue
		case "nothting", "bye":
			fmt.Println("Bye！")
			os.Exit(0)
		default:
			fmt.Println("Sorry,I didnt catch you.")
		}

	}
}
