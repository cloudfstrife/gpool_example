package dial

import (
	"errors"
	"net"
	"sync"

	"github.com/cloudfstrife/gpool"
)

var (
	pool *gpool.Pool
	once sync.Once
)

func init() {
	once.Do(func() {
		pool = gpool.DefaultPool(NewConnection)
		err := pool.Config.LoadToml("general.toml")
		if err != nil {
			panic(err)
		}
		pool.Initial()
	})
}

//NewConnection 获取新连接
func NewConnection() gpool.Item {
	return &Connection{}
}

//GetConnection 获取连接
func GetConnection() (net.Conn, error) {
	item, err := pool.GetOne()
	if err != nil {
		return nil, err
	}
	con, ok := item.(*Connection)
	if ok {
		return con.TCPConn, nil
	}
	return nil, errors.New("类型转换错误")
}

//CloseConnection 关闭连接
func CloseConnection(conn net.Conn) {
	pool.BackOne(&Connection{
		TCPConn: conn,
	})
}

//ClosePool 关闭连接池
func ClosePool() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	pool.Shutdown(wg)
	wg.Wait()
}
