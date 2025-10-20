package frontend

import "context"

type IDemo interface {
	Index(ctx context.Context) (string, error)
}

var localDemo IDemo

func Demo() IDemo {
	return localDemo
}

func RegisterDemo(d IDemo) {
	localDemo = d
}
