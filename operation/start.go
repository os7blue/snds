package operation

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
	"snds/option"
)

type start struct {
}

func (s *start) StartTask() {

	// 创建一个Cron调度器
	c := cron.New()

	for _, t := range option.Option.Tasks {

		// 添加定时任务
		_, err := c.AddFunc(t.Cron, func() {

			fmt.Println(t)

		})
		if err != nil {
			log.Println("添加定时任务失败:", err)
			return
		}

	}

	// 启动Cron调度器
	c.Start()

	select {}

}
