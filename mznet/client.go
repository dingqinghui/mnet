/**
 * @Author: dingQingHui
 * @Description:
 * @File: client
 * @Version: 1.0.0
 * @Date: 2022/10/28 16:56
 */

package mznet

type (
	IClient interface {
		ICloser
		Connect() error
	}
)
