package utils

import "net"

// 获取动态启动的端口号
func GetFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	listen, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}

	// *net.TCPAddr 实现了接口Addr，因此可以使用下述断言，将Addr()返回的Addr接口，断言为具体类型TCPAddr
	return listen.Addr().(*net.TCPAddr).Port, nil
}
