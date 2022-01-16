package main

import (
	"fmt"
	"time"

	"github.com/gorhill/cronexpr"
)

func main() {
	var (
		expr     *cronexpr.Expression
		err      error
		now      time.Time
		nextTime time.Time
	)

	// 那一分钟(0-59), 那一小时(0-23), 那天(1-31), 那月(1-12), 星期几(0-6)

	// 每分钟执行一次
	// if expr, err = cronexpr.Parse("* * * * *"); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// 每个5分钟执行一次
	if expr, err = cronexpr.Parse("*/5 * * * * * *"); err != nil {
		fmt.Println(err)
		return
	}

	// 获取当前时间
	now = time.Now()
	// 下次调度时间(下次调度时间并不是按照启动时间， 有内部规则)
	nextTime = expr.Next(now)

	// 等待定时器超时的函数
	time.AfterFunc(nextTime.Sub(now), func() {
		fmt.Println("被调度了:", nextTime)
	})

	time.Sleep(5 * time.Second)
}
