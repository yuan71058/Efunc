package class

import (
	"bufio"
	"net"
	"sync"
	"time"
)

// L_TCP服务端 TCP 服务端，支持多客户端并发连接、消息收发和广播。
// 基于 net.ListenTCP 实现，每个客户端连接在独立 goroutine 中处理。
//
// 使用示例:
//
//	s := &class.L_TCP服务端{}
//	s.S收到数据回调 = func(客户端地址 string, 数据 []byte) {
//	    fmt.Println("收到:", string(数据))
//	}
//	s.Q启动(8888)
type L_TCP服务端 struct {
	mu             sync.Mutex
	listener       net.Listener
	客户端列表     map[string]net.Conn
	running        bool
	S收到数据回调  func(客户端地址 string, 数据 []byte)
	K客户端连接回调 func(客户端地址 string)
	K客户端断开回调 func(客户端地址 string)
}

// Q启动 启动 TCP 服务端，监听指定端口。
// 每个客户端连接在独立 goroutine 中处理，通过回调函数通知上层。
//
// 参数:
//   - 端口: 监听的端口号（如 8888）
//
// 返回:
//   - error: 启动失败时返回错误
func (s *L_TCP服务端) Q启动(端口 int) error {
	addr := &net.TCPAddr{Port: 端口}
	var err error
	s.listener, err = net.ListenTCP("tcp", addr)
	if err != nil {
		return err
	}
	s.客户端列表 = make(map[string]net.Conn)
	s.running = true

	go func() {
		for s.running {
			conn, err := s.listener.Accept()
			if err != nil {
				if !s.running {
					return
				}
				continue
			}

			addr := conn.RemoteAddr().String()
			s.mu.Lock()
			s.客户端列表[addr] = conn
			s.mu.Unlock()

			if s.K客户端连接回调 != nil {
				go s.K客户端连接回调(addr)
			}

			go s.handleClient(conn, addr)
		}
	}()

	return nil
}

func (s *L_TCP服务端) handleClient(conn net.Conn, addr string) {
	defer func() {
		conn.Close()
		s.mu.Lock()
		delete(s.客户端列表, addr)
		s.mu.Unlock()
		if s.K客户端断开回调 != nil {
			go s.K客户端断开回调(addr)
		}
	}()

	reader := bufio.NewReader(conn)
	for s.running {
		conn.SetReadDeadline(time.Now().Add(30 * time.Second))
		data, err := reader.ReadBytes('\n')
		if err != nil {
			return
		}
		if len(data) > 0 && data[len(data)-1] == '\n' {
			data = data[:len(data)-1]
		}
		if len(data) > 0 && data[len(data)-1] == '\r' {
			data = data[:len(data)-1]
		}
		if s.S收到数据回调 != nil {
			go s.S收到数据回调(addr, data)
		}
	}
}

// T停止 停止 TCP 服务端，关闭所有连接。
func (s *L_TCP服务端) T停止() {
	s.running = false
	if s.listener != nil {
		s.listener.Close()
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, conn := range s.客户端列表 {
		conn.Close()
	}
	s.客户端列表 = make(map[string]net.Conn)
}

// F发送数据 向指定客户端发送数据（自动追加换行符作为分隔符）。
//
// 参数:
//   - 客户端地址: 客户端地址（如 "192.168.1.1:12345"）
//   - 数据: 要发送的字节数据
//
// 返回:
//   - error: 发送失败时返回错误
func (s *L_TCP服务端) F发送数据(客户端地址 string, 数据 []byte) error {
	s.mu.Lock()
	conn, ok := s.客户端列表[客户端地址]
	s.mu.Unlock()
	if !ok {
		return net.ErrClosed
	}
	_, err := conn.Write(append(数据, '\n'))
	return err
}

// F发送文本 向指定客户端发送文本数据。
//
// 参数:
//   - 客户端地址: 客户端地址
//   - 文本: 要发送的文本
//
// 返回:
//   - error: 发送失败时返回错误
func (s *L_TCP服务端) F发送文本(客户端地址 string, 文本 string) error {
	return s.F发送数据(客户端地址, []byte(文本))
}

// G广播数据 向所有已连接客户端广播数据。
//
// 参数:
//   - 数据: 要广播的字节数据
func (s *L_TCP服务端) G广播数据(数据 []byte) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, conn := range s.客户端列表 {
		conn.Write(append(数据, '\n'))
	}
}

// G广播文本 向所有已连接客户端广播文本。
//
// 参数:
//   - 文本: 要广播的文本
func (s *L_TCP服务端) G广播文本(文本 string) {
	s.G广播数据([]byte(文本))
}

// Q取客户端数量 获取当前连接的客户端数量。
//
// 返回:
//   - int: 客户端数量
func (s *L_TCP服务端) Q取客户端数量() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.客户端列表)
}

// Q取客户端列表 获取所有已连接客户端的地址列表。
//
// 返回:
//   - []string: 客户端地址列表
func (s *L_TCP服务端) Q取客户端列表() []string {
	s.mu.Lock()
	defer s.mu.Unlock()
	list := make([]string, 0, len(s.客户端列表))
	for addr := range s.客户端列表 {
		list = append(list, addr)
	}
	return list
}

// L_TCP客户端 TCP 客户端，支持连接服务端、收发数据。
// 使用 bufio 以换行符作为消息分隔符。
//
// 使用示例:
//
//	c := &class.L_TCP客户端{}
//	c.S收到数据回调 = func(数据 []byte) { fmt.Println(string(数据)) }
//	c.L连接("127.0.0.1:8888")
//	c.F发送文本("hello")
type L_TCP客户端 struct {
	conn           net.Conn
	running        bool
	mu             sync.Mutex
	S收到数据回调  func(数据 []byte)
	D断开回调      func()
}

// L连接 连接到 TCP 服务端。
// 连接成功后启动接收 goroutine，通过 S收到数据回调 通知上层。
//
// 参数:
//   - 地址: 服务端地址（格式 "IP:端口"，如 "127.0.0.1:8888"）
//
// 返回:
//   - error: 连接失败时返回错误
func (c *L_TCP客户端) L连接(地址 string) error {
	var err error
	c.conn, err = net.DialTimeout("tcp", 地址, 10*time.Second)
	if err != nil {
		return err
	}
	c.running = true
	go c.receive()
	return nil
}

func (c *L_TCP客户端) receive() {
	defer func() {
		c.running = false
		c.conn.Close()
		if c.D断开回调 != nil {
			c.D断开回调()
		}
	}()

	reader := bufio.NewReader(c.conn)
	for c.running {
		c.conn.SetReadDeadline(time.Now().Add(30 * time.Second))
		data, err := reader.ReadBytes('\n')
		if err != nil {
			return
		}
		if len(data) > 0 && data[len(data)-1] == '\n' {
			data = data[:len(data)-1]
		}
		if len(data) > 0 && data[len(data)-1] == '\r' {
			data = data[:len(data)-1]
		}
		if c.S收到数据回调 != nil {
			c.S收到数据回调(data)
		}
	}
}

// D断开 断开与 TCP 服务端的连接。
func (c *L_TCP客户端) D断开() {
	c.running = false
	if c.conn != nil {
		c.conn.Close()
	}
}

// F发送数据 向服务端发送数据（自动追加换行符）。
//
// 参数:
//   - 数据: 要发送的字节数据
//
// 返回:
//   - error: 发送失败时返回错误
func (c *L_TCP客户端) F发送数据(数据 []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.conn == nil {
		return net.ErrClosed
	}
	_, err := c.conn.Write(append(数据, '\n'))
	return err
}

// F发送文本 向服务端发送文本数据。
//
// 参数:
//   - 文本: 要发送的文本
//
// 返回:
//   - error: 发送失败时返回错误
func (c *L_TCP客户端) F发送文本(文本 string) error {
	return c.F发送数据([]byte(文本))
}

// S是否已连接 检查是否已连接服务端。
//
// 返回:
//   - bool: true 表示已连接
func (c *L_TCP客户端) S是否已连接() bool {
	return c.conn != nil && c.running
}

// Q取本地地址 获取本地连接地址。
//
// 返回:
//   - string: 本地地址
func (c *L_TCP客户端) Q取本地地址() string {
	if c.conn != nil {
		return c.conn.LocalAddr().String()
	}
	return ""
}