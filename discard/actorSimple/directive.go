/**
 * @Author: dingQingHui
 * @Description:收到子actor异常处理指令
 * @File: derective
 * @Version: 1.0.0
 * @Date: 2022/10/11 15:49
 */

package actorNew

type Directive int

const (
	ResumeDirective       Directive = iota //丢弃panic消息,继续运行
	RestartDirective                       //重启，未处理的消息全部丢弃
	StopDirective                          //终止
	EscalateStopDirective                  //向上传递异常
)
