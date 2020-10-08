package time

import (
	"fmt"
	"time"
)

func main() {


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


	time.NewTimer(time.Second)  // 时间到了触发一次
	time.NewTicker(time.Second) // 每隔时间段触发一次


	achan := make(chan int, 1)
	go func() {
		ticker := time.Tick(time.Second)
		<-ticker
		fmt.Println("123")
	}()

	<-achan
	return
}

// timeer的使用、ticker的使用、time的构造。 timeer的实现原理，涉及到go runtime的调用了