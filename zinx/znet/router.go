package znet

import "zinx/zinx/ziface"

// 自定义实现时，可继承base，然后重写,base不需要实现，只需要提供实现所有接口的base，在自定义中实现需要的方法，这样自定义中就可以不需要全部实现。

type BaseRouter struct{}

func (br *BaseRouter) PreHandle(request ziface.IRequest)  {}
func (br *BaseRouter) MainHandle(request ziface.IRequest) {}
func (br *BaseRouter) PostHandle(request ziface.IRequest) {}
