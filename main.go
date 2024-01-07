package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"snds/operation"
	"snds/option"
)

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

func main() {
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

	//123

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
