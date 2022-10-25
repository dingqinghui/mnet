/**
 * @Author: dingQingHui
 * @Description:
 * @File: logger
 * @Version: 1.0.0
 * @Date: 2022/10/6 16:06
 */

package actorNew

import "fmt"

func DebugLog(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func InfoLog(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}
