package main

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

func express() {
	/* time 的格式来源 */
	time.Now()
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 按照指定时区和指定格式解析字符串时间
	timeObj, err := time.ParseInLocation("2006/01/02 15:04:05", "2019/08/04 14:15:20", loc)
	fmt.Println(timeObj)
}

// ticker 里每隔多长时间。ticker.C里多一个元素。
func ticker() {
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			fmt.Println("this is a ticker")
		}
	}
}

func tick() {
	tick := time.Tick(time.Second)
	for range tick {
		fmt.Println("time tick")
	}
}

// ticker 里每到了某某时间错。ticker.C里多一个元素。
func timer() {
	//for {
	timer := time.NewTimer(time.Second)

	for range timer.C {
		fmt.Println("tick ")
	}
	//select {
	//case <-timer.C:
	//	fmt.Println("this is a timer")
	//}
	//}
}

func main() {
	aChan := make(chan int, 1)
	go ticker()
	<-aChan
	return
}

// timeer的使用、ticker的使用、time的构造。 timeer的实现原理，涉及到go runtime的调用了
