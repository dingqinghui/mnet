/**
 * @Author: dingQingHui
 * @Description:
 * @File: main_server
 * @Version: 1.0.0
 * @Date: 2022/7/12 15:56
 */

package main

import "github.com/dingqinghui/mz/service"

func main() {
	watchDog := service.NewWatchdog()
	watchDog.Init()
	watchDog.Run()
	select {}
}
