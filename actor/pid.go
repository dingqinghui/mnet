/**
 * @Author: dingQingHui
 * @Description:
 * @File: pid
 * @Version: 1.0.0
 * @Date: 2022/9/27 11:39
 */

package actor

type (
	IPid interface {
		Id() int64
		iProcess
	}

	Pid struct {
		id int64
		iProcess
	}
)

func (p *Pid) Id() int64 {
	return p.id
}
