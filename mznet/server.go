/**
 * @Author: dingQingHui
 * @Description:
 * @File: server
 * @Version: 1.0.0
 * @Date: 2022/10/28 17:25
 */

package mznet

type (
	IServer interface {
		RunEventLoop() error
		ICloser
	}
)
