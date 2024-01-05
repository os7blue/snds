package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"snds/operation"
	"snds/option"
	"snds/util"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	var input string
	for {

		showMenu()
		fmt.Print("请输入选项：")

		scanner.Scan()
		input = scanner.Text()

		if scanner.Err() != nil {
			fmt.Println(scanner.Err().Error())
			os.Exit(1)
		}

		switch input {
		case "0":
			os.Exit(1)
		case "1":
			//todo start server
			break
		case "2":
			//todo stop server
			break
		case "3":
			//todo restart server
			break
		case "4":
			//todo show config info
			break
		case "5":
			//todo create new task

			operation.Create.Create()

			break
		case "6":
			//todo show help menu
			break
		default:

		}

	}

}

func showMenu() {
	util.RcpUtil.RandomColorPrintln("1、启动服务")
	util.RcpUtil.RandomColorPrintln("2、停止服务")
	util.RcpUtil.RandomColorPrintln("3、重启服务")
	util.RcpUtil.RandomColorPrintln("4、查看配置")
	util.RcpUtil.RandomColorPrintln("5、新建任务")
	util.RcpUtil.RandomColorPrintln("6、help")
	util.RcpUtil.RandomColorPrintln("0、退出")

}

type flags struct {
	//start service
	Start bool
	//stop service
	Stop bool
	//run
	Run bool
	//option
	Option bool
	//Create
	Create bool
}

func main1() {
	err := option.Init()

	if err != nil {
		log.Printf("snds初始化失败：%s", err.Error())
		os.Exit(1)
	}

	option.Logger.Infof("初始化完成，参数：%+v", option.Option)

	var f = new(flags)

	//start
	flag.BoolVar(&f.Start, "start", false, "start service")
	//stop
	flag.BoolVar(&f.Stop, "stop", false, "stop service")
	//Run
	flag.BoolVar(&f.Run, "run", false, "run service,private flag")
	//Option
	flag.BoolVar(&f.Option, "option", false, "run service,private flag")
	//create
	flag.BoolVar(&f.Create, "create", false, "create new config.go")

	flag.Parse()

	numFlags := flag.NFlag()
	if numFlags != 1 {
		log.Print("参数错误")
		os.Exit(1)
	}

	option.Logger.Infof("cli调用，flag：%+v", f)

	if f.Start {
		cmd := option.Option.Command.Run
		err := cmd.Start()
		if err != nil {
			os.Exit(1)
		}
		log.Println("启动后台启动")

	}

	if f.Run {
		err := operation.Start.StartTask()
		if err != nil {
			log.Printf("启动任务失败：%s", err.Error())
			return
		}
		log.Println("启动任务成功")
	}

	if f.Stop {
		fmt.Println("stop")
	}

	if f.Option {

		// 使用json.Marshal将结构体转换为JSON格式的字节切片
		jsonData, err := json.MarshalIndent(option.Option, "", "    ")
		if err != nil {
			log.Println("JSON编码失败:", err)
			return
		}

		// 将格式化的JSON格式的字节切片转换为字符串并打印
		log.Println(string(jsonData))

	}

	if f.Create {
		operation.Create.Create()
	}

}

func parseJson(t any) string {
	jsonData, err := json.MarshalIndent(t, "", "    ")
	if err != nil {
		log.Println("JSON编码失败:", err)
		return ""
	}

	// 将格式化的JSON格式的字节切片转换为字符串并打印
	return string(jsonData)
}
