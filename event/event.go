/**
 * @Author: dingQingHui
 * @Description:将事件派发者与事件处理者解耦
 * @File: event
 * @Version: 1.0.0
 * @Date: 2022/9/26 10:47
 */

package event

import "sync"

var event = make(map[string][]func(interface{}))
var lock sync.RWMutex

func RegisterEvent(name string, callBack func(interface{})) {
	lock.Lock()
	defer lock.Unlock()
	//事件的列表
	list := event[name]
	/*
	   通过事件名（name）进行查询，返回回调列表（[]func(interface{}）
	*/
	//在列表切片中添加函数
	list = append(list, callBack)
	/*
	   为同一个事件名称在已经注册的事件回调的列表中再添加一个回调函数
	*/
	//将修改的事件列表切片保存回去
	event[name] = list
	/*
	   将修改后的函数列表设置到 map 的对应事件名中
	*/
}

//
// CallEvent
// @Description: 调用事件
// @param name
// @param param
//
func CallEvent(name string, param interface{}) {
	lock.RLock()
	defer lock.RUnlock()
	//找到事件map映射
	list := event[name]
	/*
	   通过注册事件回调的 event 和事件名字查询处理函数列表 list
	*/

	//遍历列表找到函数
	for _, callback := range list {
		/*
		   遍历这个事件列表，如果没有找到对应的事件，list 将是一个空切片
		*/

		//传入参数调用回调
		callback(param)
	}
}
