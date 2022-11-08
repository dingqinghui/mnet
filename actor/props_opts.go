/**
 * @Author: dingQingHui
 * @Description:
 * @File: props_opts
 * @Version: 1.0.0
 * @Date: 2022/9/27 10:54
 */

package actor

type PropsOption func(props *Props)

func WithOnInit(onInit ...InitFunc) PropsOption {
	return func(props *Props) {
		props.onInits = onInit
	}
}

func NewPropsWithProducer(producer Producer, options ...PropsOption) *Props {
	props := &Props{
		producer: producer,
	}
	for _, opt := range options {
		opt(props)
	}
	return props
}

func NewPropsWithFunc(receiver ReceiveFunc, options ...PropsOption) *Props {
	producer := func() IActor { return receiver }
	props := &Props{
		producer: producer,
	}
	for _, opt := range options {
		opt(props)
	}
	return props
}
