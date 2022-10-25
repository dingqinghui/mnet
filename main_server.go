/**
 * @Author: dingQingHui
 * @Description:
 * @File: main_server
 * @Version: 1.0.0
 * @Date: 2022/7/12 15:56
 */

package main

import (
	_ "github.com/dingqinghui/mz/service"
	"time"
)

func main() {

	//service := actor.NewService("watchDog")

	time.Sleep(time.Hour * 2)
}
