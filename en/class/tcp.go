package class

import (
	"bufio"
	"net"
	"sync"
	"time"
)

// TCPServer TCP 服务端，支持多客户端并发连接、消息收发和广播。
// 基于 net.ListenTCP 实现，每个客户端连接在独立 goroutine 中处理。
//
// 使用示例:
//
//	s := &class.TCPServer{}
//	s.OnReceiveData = func(addr string, data []byte) {
//	    fmt.Println("收到:", string(data))
//	}
//	s.Start(8888)
type TCPServer struct {
	mu            sync.Mutex
	listener      net.Listener
	clients       map[string]net.Conn
	running       bool
	OnReceiveData func(addr string, data []byte)
	OnConnect     func(addr string)
	OnDisconnect  func(addr string)
}

// Start 启动 TCP 服务端，监听指定端口。
// 每个客户端连接在独立 goroutine 中处理，通过回调函数通知上层。
//
// 参数:
//   - port: 监听的端口号（如 8888）
//
// 返回:
//   - error: 启动失败时返回错误
func (s *TCPServer) Start(port int) error {
	addr := &net.TCPAddr{Port: port}
	var err error
	s.listener, err = net.ListenTCP("tcp", addr)
	if err != nil {
		return err
	}
	s.clients = make(map[string]net.Conn)
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
			s.clients[addr] = conn
			s.mu.Unlock()

			if s.OnConnect != nil {
				go s.OnConnect(addr)
			}

			go s.handleClient(conn, addr)
		}
	}()

	return nil
}

func (s *TCPServer) handleClient(conn net.Conn, addr string) {
	defer func() {
		conn.Close()
		s.mu.Lock()
		delete(s.clients, addr)
		s.mu.Unlock()
		if s.OnDisconnect != nil {
			go s.OnDisconnect(addr)
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
		if s.OnReceiveData != nil {
			go s.OnReceiveData(addr, data)
		}
	}
}

// Stop 停止 TCP 服务端，关闭所有连接。
func (s *TCPServer) Stop() {
	s.running = false
	if s.listener != nil {
		s.listener.Close()
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, conn := range s.clients {
		conn.Close()
	}
	s.clients = make(map[string]net.Conn)
}

// SendBytes 向指定客户端发送数据（自动追加换行符作为分隔符）。
//
// 参数:
//   - addr: 客户端地址（如 "192.168.1.1:12345"）
//   - data: 要发送的字节数据
//
// 返回:
//   - error: 发送失败时返回错误
func (s *TCPServer) SendBytes(addr string, data []byte) error {
	s.mu.Lock()
	conn, ok := s.clients[addr]
	s.mu.Unlock()
	if !ok {
		return net.ErrClosed
	}
	_, err := conn.Write(append(data, '\n'))
	return err
}

// SendText 向指定客户端发送文本数据。
//
// 参数:
//   - addr: 客户端地址
//   - text: 要发送的文本
//
// 返回:
//   - error: 发送失败时返回错误
func (s *TCPServer) SendText(addr string, text string) error {
	return s.SendBytes(addr, []byte(text))
}

// BroadcastBytes 向所有已连接客户端广播数据。
//
// 参数:
//   - data: 要广播的字节数据
func (s *TCPServer) BroadcastBytes(data []byte) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, conn := range s.clients {
		conn.Write(append(data, '\n'))
	}
}

// BroadcastText 向所有已连接客户端广播文本。
//
// 参数:
//   - text: 要广播的文本
func (s *TCPServer) BroadcastText(text string) {
	s.BroadcastBytes([]byte(text))
}

// ClientCount 获取当前连接的客户端数量。
//
// 返回:
//   - int: 客户端数量
func (s *TCPServer) ClientCount() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.clients)
}

// ClientList 获取所有已连接客户端的地址列表。
//
// 返回:
//   - []string: 客户端地址列表
func (s *TCPServer) ClientList() []string {
	s.mu.Lock()
	defer s.mu.Unlock()
	list := make([]string, 0, len(s.clients))
	for addr := range s.clients {
		list = append(list, addr)
	}
	return list
}

// TCPClient TCP 客户端，支持连接服务端、收发数据。
// 使用 bufio 以换行符作为消息分隔符。
//
// 使用示例:
//
//	c := &class.TCPClient{}
//	c.OnReceiveData = func(data []byte) { fmt.Println(string(data)) }
//	c.Connect("127.0.0.1:8888")
//	c.SendText("hello")
type TCPClient struct {
	conn          net.Conn
	running       bool
	mu            sync.Mutex
	OnReceiveData func(data []byte)
	OnDisconnect  func()
}

// Connect 连接到 TCP 服务端。
// 连接成功后启动接收 goroutine，通过 OnReceiveData 回调通知上层。
//
// 参数:
//   - addr: 服务端地址（格式 "IP:端口"，如 "127.0.0.1:8888"）
//
// 返回:
//   - error: 连接失败时返回错误
func (c *TCPClient) Connect(addr string) error {
	var err error
	c.conn, err = net.DialTimeout("tcp", addr, 10*time.Second)
	if err != nil {
		return err
	}
	c.running = true
	go c.receive()
	return nil
}

func (c *TCPClient) receive() {
	defer func() {
		c.running = false
		c.conn.Close()
		if c.OnDisconnect != nil {
			c.OnDisconnect()
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
		if c.OnReceiveData != nil {
			go c.OnReceiveData(data)
		}
	}
}

// Disconnect 断开与服务端的连接。
func (c *TCPClient) Disconnect() {
	c.running = false
	if c.conn != nil {
		c.conn.Close()
	}
}

// SendBytes 向服务端发送字节数据（自动追加换行符）。
//
// 参数:
//   - data: 要发送的字节数据
//
// 返回:
//   - error: 发送失败时返回错误
func (c *TCPClient) SendBytes(data []byte) error {
	if c.conn == nil {
		return net.ErrClosed
	}
	_, err := c.conn.Write(append(data, '\n'))
	return err
}

// SendText 向服务端发送文本数据。
//
// 参数:
//   - text: 要发送的文本
//
// 返回:
//   - error: 发送失败时返回错误
func (c *TCPClient) SendText(text string) error {
	return c.SendBytes([]byte(text))
}

// IsConnected 检查客户端是否已连接。
//
// 返回:
//   - bool: true 表示已连接
func (c *TCPClient) IsConnected() bool {
	return c.running && c.conn != nil
}

// LocalAddr 获取本地地址。
//
// 返回:
//   - string: 本地地址
func (c *TCPClient) LocalAddr() string {
	if c.conn != nil {
		return c.conn.LocalAddr().String()
	}
	return ""
}

// RemoteAddr 获取远程地址。
//
// 返回:
//   - string: 远程地址
func (c *TCPClient) RemoteAddr() string {
	if c.conn != nil {
		return c.conn.RemoteAddr().String()
	}
	return ""
}