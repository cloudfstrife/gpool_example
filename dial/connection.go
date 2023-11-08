package dial

import (
	"fmt"
	"net"
)

//Connection 连接池对象
type Connection struct {
	TCPConn net.Conn
}

//Initial 初始化
func (c *Connection) Initial(params map[string]string) error {
	con, err := net.Dial("tcp", params["host"]+":"+params["port"])
	if err != nil {
		return err
	}
	c.TCPConn = con
	return nil
}

//Destory 销毁连接
func (c *Connection) Destory(params map[string]string) error {
	return c.TCPConn.Close()
}

//Check 检查元素连接是否可用
func (c *Connection) Check(params map[string]string) error {
	fmt.Println("检查连接可用")
	return nil
}
