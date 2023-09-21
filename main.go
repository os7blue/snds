package main

import (
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
}

func main() {
	err := option.Init()

	if err != nil {
		log.Printf("初始化失败：%s", err.Error())
		os.Exit(1)
	}

	var f = new(flags)

	//start
	flag.BoolVar(&f.Start, "start", false, "start service")
	//stop
	flag.BoolVar(&f.Stop, "stop", false, "stop service")
	//Run
	flag.BoolVar(&f.Run, "run", false, "run service,private flag")

	flag.Parse()

	fmt.Println(option.Option)
	fmt.Println(f)

	if f.Start {
		cmd := option.Option.Command.Run
		err := cmd.Start()
		if err != nil {
			os.Exit(1)
		}
	}

	if f.Run {
		operation.Operations.Start.StartTask()
	}

	if f.Stop {
		fmt.Println("stop")
	}

}
